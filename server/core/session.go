package core

import (
	"laatoo/sdk/ctx"
	"laatoo/sdk/utils"
	"time"

	"github.com/twinj/uuid"
)

var (
	SESSION_OBJ = "__session"
)

type session struct {
	id           string
	userid       string
	creationTime time.Time
	data         utils.StringMap
	mgr          *sessionManager
}

func newSession(id string) *session {
	if id == "" {
		id = uuid.NewV1().String()
	}
	return &session{id: id, creationTime: time.Now(), data: make(utils.StringMap)}
}

func (sess *session) GetId() string {
	return sess.id
}
func (sess *session) CreationTime() time.Time {
	return sess.creationTime

}
func (sess *session) GetUser() string {
	return sess.userid
}
func (sess *session) GetString(key string) (string, bool) {
	return sess.data.GetString(key)
}
func (sess *session) GetBool(key string) (bool, bool) {
	return sess.data.GetBool(key)
}
func (sess *session) GetInt(key string) (int, bool) {
	return sess.data.GetInt(key)
}
func (sess *session) GetStringArray(key string) ([]string, bool) {
	return sess.data.GetStringArray(key)
}
func (sess *session) AllKeys() []string {
	return sess.data.AllKeys()
}
func (sess *session) GetStringMap(key string) (utils.StringMap, bool) {
	return sess.data.GetStringMap(key)
}
func (sess *session) GetStringsMap(key string) (utils.StringsMap, bool) {
	return sess.data.GetStringsMap(key)
}
func (sess *session) Set(key string, val interface{}) {
	sess.data.Set(key, val)
}
func (sess *session) SetVals(vals utils.StringMap) {
	sess.data.SetVals(vals)
}

func (sess *session) Save(ctx ctx.Context) error {
	if sess.mgr != nil {
		return sess.mgr.Save(ctx, sess)
	}
	return nil
}
