var {argv} = require('./buildconfig');


function log() {
    let args = Array.prototype.slice.call(arguments);
    if(argv.verbose) {
      console.log(args.join(' '));
    }
}

module.exports = {
    log: log
}