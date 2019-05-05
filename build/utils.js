var {argv} = require('./buildconfig');
var fs = require('fs-extra')


function log() {
    let args = Array.prototype.slice.call(arguments);
    if(argv.verbose) {
      console.log(args.join(' '));
    }
}

function listDir(path) {
  console.log("listing directory contents", path)
  fs.readdir(path, function(err, items) {
    for (var i=0; i<items.length; i++) {
        log(items[i]);
    }
  });  
}


module.exports = {
    log: log,
    listDir: listDir
}