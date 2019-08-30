var {argv} = require('./buildconfig');
var fs = require('fs-extra')
var shell = require('shelljs');
var {execSync, execFileSync, spawnSync} = require('child_process');
var path = require('path');


function log() {
    let args = Array.prototype.slice.call(arguments);
    if(argv.verbose) {
      console.log(args.join(' '));
    }
}

function listDir(dirpath) {
  log("listing directory contents", dirpath)
  fs.readdirSync(dirpath).forEach(file => {
    let fullPath = path.join(dirpath, file);
    if (fs.lstatSync(fullPath).isDirectory()) {
       log(fullPath);
       listDir(fullPath);
     } else {
       log(fullPath);
     }  
  });
}

function clearDirectory(path) {
  var filesToDel
  filesToDel = fs.readdirSync(path)
  for (file of filesToDel) {
    fs.removeSync(file);
  }
}

function createGoModule(name, path) {
  let optionsArr = ["mod", "init", name]
  let res = spawnSync("go", optionsArr, {cwd: path})
  if(res.status !== 0) {
    shell.echo('Golang sdk creation unsuccessful');
    shell.echo("Res", res.stdout.toString())
    shell.echo("Err", res.stderr.toString())      
    shell.exit(1);
  } else {
    log("Go module created ", name, " at ", path)
  }              
}


function tidyGoModule(path) {
  let optionsArr = ["mod", "tidy"]
  let res = spawnSync("go", optionsArr, {cwd: path})
  if(res.status !== 0) {
    shell.echo('Golang tidy unsuccessful');
    shell.echo("Res", res.stdout.toString())
    shell.echo("Err", res.stderr.toString())      
    shell.exit(1);
  } else {
    log("Go module tidy complete ", path)
  }              
}

module.exports = {
    log: log,
    listDir: listDir,
    createGoModule: createGoModule,
    clearDirectory: clearDirectory,
    tidyGoModule: tidyGoModule
}