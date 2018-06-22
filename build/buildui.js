/*let webpack = require('webpack')
let rimraf = require('rimraf')
let fileExists = require('file-exists')
var shell = require('shelljs');
var mkdirp = require('mkdirp');

let args = process.argv
console.log("arguments ", args)

let skipModules = "false"

let printconfig = "false"

if (args.length > 3) {
  skipModules = args[3]
}

if (args.length > 4) {
  printconfig = args[4]
}


if (skipModules === "false") {
  if (fileExists.sync('/nodemodules/package.json')) {
    console.log("Installing package json from nodemodules")
    if (shell.exec('cd /nodemodules && npm i --prefix /nodemodules ').code !== 0) {
      shell.echo('Get package failed');
      shell.exit(1);
    }
  } else {
    console.log("No package json in nodemodules")
  }

  if (fileExists.sync('/plugin/package.json')) {
    console.log("Installing package json from plugin")
    if (shell.exec('cd /plugin && npm i --prefix /nodemodules ').code !== 0) {
      shell.echo('Get package failed');
      shell.exit(1);
    }
  } else {
    console.log("No package json file for plugin")
  }

  if (fileExists.sync('/build/package.json')) {
    if (shell.exec('cd /build && npm i --prefix /nodemodules ').code !== 0) {
      shell.echo('Get package failed');
      shell.exit(1);
    }
  } else {
    console.log("No package json in build")
  }
}

let config = {}

let configFunc = require('./cfg/webpack.lib')

config = configFunc({
  library: args[2]
})

if (fileExists.sync('/plugin/build.js')) {
  let custom = require('/plugin/build')
  if (custom.config!=nil) {
    config = custom.config(config)
  }
} else {
  console.log("No custom build js file")
}

console.log("printconfig", printconfig)
if(printconfig === "true") {
  console.log("config", config)
}


let compiler = webpack(config)

rimraf('/plugin/dist', function() {
  console.log("Removed directory dist")
  mkdirp('/plugin/dist/scripts', function() {
    console.log("starting compilation")
    compiler.run(function(err, stats) {
      console.log("Errors: ", err);
      console.log("Stats: ", stats)
    });
  })
})


#!/bin/sh

rm -fr tmp/$1
mkdir -p tmp/$1/objects
go build -buildmode=plugin -o tmp/$1/objects/$1.so $2/*.go
if [ -e $2/config ]
 then cp -R $2/config/* tmp/$1/
fi
if [ -e $2/files ]
 then cp -R $2/files tmp/$1/
fi
tar -czvf tmp/$1.tar.gz -C tmp/ $1
mkdir -p $3
cp tmp/$1.tar.gz  $3
*/