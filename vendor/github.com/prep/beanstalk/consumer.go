package beanstalk

import (
	"strings"
	"sync"
	"time"
)

// Consumer reserves jobs from a beanstalk server and keeps those jobs alive
// until an external consumer has either buried, deleted or released it.
type Consumer struct {
	tubes     []string
	jobC      chan<- *Job
	pause     chan bool
	stop      chan struct{}
	isPaused  bool
	isStopped bool
	options   *Options
	client    *Client
	sync.Mutex
}

// NewConsumer returns a new Consumer object.
func NewConsumer(socket string, tubes []string, jobC chan<- *Job, options *Options) *Consumer {
	if options == nil {
		options = DefaultOptions()
	}

	consumer := &Consumer{
		tubes:    tubes,
		jobC:     jobC,
		pause:    make(chan bool, 1),
		stop:     make(chan struct{}, 1),
		isPaused: true,
		options:  options,
	}

	go consumer.connectionManager(socket)
	return consumer
}

// Play allows this consumer to start reserving jobs. Returns true on success
// and false if this consumer was stopped.
func (consumer *Consumer) Play() bool {
	consumer.Lock()
	defer consumer.Unlock()

	if consumer.isStopped {
		return false
	}

	select {
	case <-consumer.pause:
	default:
	}

	consumer.pause <- false
	return true
}

// Pause stops this consumer from reserving jobs. Returns true on success and
// false if this consumer was stopped.
func (consumer *Consumer) Pause() bool {
	consumer.Lock()
	defer consumer.Unlock()

	if consumer.isStopped {
		return false
	}

	select {
	case <-consumer.pause:
	default:
	}

	consumer.pause <- true
	return true
}

// Stop this consumer. Returns true on success and false if this consumer was
// already stopped.
func (consumer *Consumer) Stop() bool {
	consumer.Lock()
	defer consumer.Unlock()

	if consumer.isStopped {
		return false
	}

	consumer.isStopped = true
	consumer.stop <- struct{}{}
	return true
}

// connectionManager is responsible for setting up a connection to the
// beanstalk server and wrapping it in a Client, which on success is passed
// to the clientManager function.
func (consumer *Consumer) connectionManager(socket string) {
	var (
		err     error
		options = consumer.options
	)

	// Start a new connection.
	newConnection, abortConnect := Connect(socket, consumer.options)

	for {
		select {
		// This case triggers whenever a new connection was established.
		case conn := <-newConnection:
			client := NewClient(conn, options)

			options.LogInfo("Watching tubes: %s", strings.Join(consumer.tubes, ", "))
			for _, tube := range consumer.tubes {
				if err = client.Watch(tube); err != nil {
					options.LogError("%s: Error watching tube: %s", tube, err)
					break
				}
			}

			if err == nil && !includesString(consumer.tubes, "default") {
				if err := client.Ignore("default"); err != nil {
					options.LogError("default: Unable to ignore tube: %s", err)
				}
			}

			if err := consumer.clientManager(client); err != nil {
				newConnection, abortConnect = Connect(socket, consumer.options)
			} else {
				return
			}

		// Keep track of the pause state.
		case consumer.isPaused = <-consumer.pause:

		// Abort this connection and stop this consumer all together when the
		// stop signal was received.
		case <-consumer.stop:
			abortConnect <- struct{}{}
			return
		}
	}
}

// clientManager is responsible for reserving beanstalk jobs and offering them
// up the the job channel.
func (consumer *Consumer) clientManager(client *Client) (err error) {
	var (
		job          *Job
		jobOffer     chan<- *Job
		jobCommandC  = make(chan *JobCommand)
		jobsOutThere int
	)

	// This is used to pause the select-statement for a bit when the job queue
	// is full or when "reserve-with-timeout 0" yields no job.
	timeout := time.NewTimer(time.Second)
	timeout.Stop()

	// Set up a touch timer that fires whenever the pending reserved job needs
	// to be touched to keep the reservation on that job.
	touchTimer := time.NewTimer(time.Second)
	touchTimer.Stop()

	// Whenever this function returns, clean up the pending job and close the
	// client connection.
	defer func() {
		if job != nil {
			client.Release(job, job.Priority, 0)
		}

		client.Close()
		touchTimer.Stop()
		close(jobCommandC)
	}()

	// isFatalErr is a convenience function that checks if the returned error
	// from a beanstalk command is fatal, or can be ignored.
	isFatalErr := func() bool {
		if err == ErrNotFound {
			err = nil
		}
		return err != nil
	}

	for {
		// Attempt to reserve a job if the state allows for it.
		if job == nil && !consumer.isPaused {
			if jobsOutThere == 0 {
				job, err = client.Reserve(consumer.options.ReserveTimeout)
			} else {
				job, err = client.Reserve(0)
			}

			switch {
			case err == ErrDraining:
				timeout.Reset(time.Minute)
			case err == ErrDeadlineSoon:
				timeout.Reset(time.Second)
			case err != nil:
				consumer.options.LogError("Error reserving job: %s", err)
				return

			// A new job was reserved.
			case job != nil:
				job.commandC = jobCommandC
				jobOffer = consumer.jobC
				touchTimer.Reset(job.TouchAt())

			// With jobs out there and no successful reserve, wait a bit before
			// attempting another reserve.
			case jobsOutThere != 0:
				timeout.Reset(time.Second)

			// No job reserved and no jobs out there, perform another reserve almost
			// immediately.
			default:
				timeout.Reset(0)
			}
		} else {
			timeout.Reset(time.Second)
		}

		select {
		// Offer the job up on the shared jobs channel.
		case jobOffer <- job:
			job, jobOffer = nil, nil
			touchTimer.Stop()
			jobsOutThere++

		// Wait a bit before trying to reserve a job again, or just fall through.
		case <-timeout.C:

			// Touch the pending job to make sure it doesn't expire.
		case <-touchTimer.C:
			if job != nil {
				if err = client.Touch(job); err != nil {
					consumer.options.LogError("Unable to touch job %d: %s", job.ID, err)
					if isFatalErr() {
						return
					}

					job, jobOffer = nil, nil
				} else {
					touchTimer.Reset(job.TouchAt())
				}
			}

		// Bury, delete or release a reserved job.
		case req := <-jobCommandC:
			if req.Command == Touch {
				if err = client.Touch(req.Job); err != nil {
					consumer.options.LogError("Unable to touch job %d: %s", req.Job.ID, err)
				}
			} else {
				switch req.Command {
				case Bury:
					err = client.Bury(req.Job, req.Priority)
				case Delete:
					err = client.Delete(req.Job)
				case Release:
					err = client.Release(req.Job, req.Priority, req.Delay)
				}

				jobsOutThere--
				if err != nil {
					consumer.options.LogError("Unable to finish job %d: %s", req.Job.ID, err)
				}
			}

			req.Err <- err
			if isFatalErr() {
				return
			}

		// Pause or unpause this connection.
		case consumer.isPaused = <-consumer.pause:
			if consumer.isPaused && job != nil {
				if err = client.Release(job, job.Priority, 0); err != nil {
					consumer.options.LogError("Unable to release job %d: %s", job.ID, err)
					if isFatalErr() {
						return
					}
				}

				job, jobOffer = nil, nil
			}

		// Stop this connection and close this consumer down.
		case <-consumer.stop:
			return nil
		}
	}
}
