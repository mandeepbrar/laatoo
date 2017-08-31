let webpack = require('webpack')
var shell = require('shelljs');
var path = require('path');
var sprintf = require('sprintf-js').sprintf
var fs = require('fs-extra')

var argv = require('minimist')(process.argv.slice(2), {boolean:["skipObjects", "skipUI", "skipUIModules", "printUIConfig"]});

//, default: {skipObjects: false, skipUI: false, skipUIModules: false, printUIConfig: false}

//console.log(argv);

let name = argv.name

let packageFolder = argv.packageFolder

let pluginFolder = path.join("/plugins", "src", packageFolder)

let uiFolder = path.join(pluginFolder, "ui")

let filesFolder = path.join(pluginFolder, "files")

//let modConfig = require(path.join(pluginFolder, "config", "config.json"));

let deploymentFolder = "/deploy/"

let nodeModulesFolder = "/nodemodules"

let buildFolder = "/build"


function buildModule() {
  createTempDirectory(!(argv.skipUI || argv.skipObjects))
  startTask("copyconfig")()
}

function buildUI(nextTask) {
  if (argv.skipUI || !fs.pathExistsSync(uiFolder)) {
    nextTask()
    return
  }

  if (!argv.skipUIModules) {
    getUIModules();
  }

  compileWebUI(function() {
    console.log("Copying UI files")
    fs.mkdirsSync(filesFolder)
    fs.copySync(path.join(uiFolder, "dist/scripts/index.js"), path.join(filesFolder, "webui.js"))
    nextTask()
  });
}


function compileWebUI(nextTask) {
  let config = {}

  let configFunc = require(path.join(buildFolder, 'cfg/webpack.lib'))

  config = configFunc({
    library: name,
    uifolder: uiFolder
  })
  if (fs.pathExistsSync(path.join(uiFolder, 'build.js'))) {
    let custom = require(path.join(uiFolder, 'build'))
    if (custom.config!=nil) {
      config = custom.config(config)
    }
  } else {
    console.log("No custom ui build file")
  }

  if(argv.printUIConfig) {
    console.log("config", config)
  }

  let compiler = webpack(config)

  fs.removeSync(path.join(uiFolder,'dist'))

  console.log("Removed directory dist")

  fs.mkdirsSync(path.join(uiFolder,'dist/scripts'))

  console.log("Starting compilation")
  compiler.run(function(err, stats) {
    console.log("Errors: ", stats.compilation.errors);
    nextTask()
  });
}

function getUIModules() {
  if (fs.pathExistsSync(path.join(nodeModulesFolder,'package.json'))) {
    console.log("Installing package json from nodemodules")
    let command = sprintf('cd %s && npm i --prefix %s', nodeModulesFolder, nodeModulesFolder)
    if (shell.exec(command).code !== 0) {
      shell.echo('Get package failed');
      shell.exit(1);
    }
  } else {
    console.log("No package json in nodemodules")
  }

  if (fs.pathExistsSync(path.join(uiFolder,'package.json'))) {
    console.log("Installing package json from plugin")
    let command = sprintf('cd %s && npm i --prefix %s', uiFolder, nodeModulesFolder)
    if (shell.exec(command).code !== 0) {
      shell.echo('Get package failed');
      shell.exit(1);
    }
  } else {
    console.log("No package json file for plugin ui")
  }

  if (fs.pathExistsSync(path.join(buildFolder,'package.json'))) {
    console.log("Installing package json from build folder")
    let command = sprintf('cd %s && npm i --prefix %s', buildFolder, nodeModulesFolder)
    if (shell.exec(command).code !== 0) {
      shell.echo('Get package failed');
      shell.exit(1);
    }
  } else {
    console.log("No package json in build")
  }
}

function createTempDirectory(removeDir) {
  let tmpFolder = path.join("/plugins", "tmp", name)
  if (removeDir) {
    fs.removeSync(tmpFolder)
  }
  fs.mkdirsSync(tmpFolder)
}

function buildObjects(nextTask) {
  if (argv.skipObjects) {
    nextTask()
    return
  }

  console.log("Compiling golang")

  let files = fs.readdirSync(pluginFolder)
  let goFilesPresent = false
  for(var i=0;i<files.length; i++) {
    if(files[i].endsWith('.go')) {
      goFilesPresent = true
      break;
    }
  }

  if(!goFilesPresent) {
    nextTask()
    return
  }

  let tmpObjsFolder = path.join("/plugins", "tmp", name, "objects")


  fs.removeSync(tmpObjsFolder)

  fs.mkdirsSync(tmpObjsFolder)

  let command = sprintf('go build -buildmode=plugin -o %s/%s.so %s/*.go', tmpObjsFolder, name, pluginFolder)
  if (shell.exec(command).code !== 0) {
    shell.echo('Golang build failed');
    shell.exit(1);
  } else {
    nextTask()
  }
}

function copyConfig(nextTask) {
  console.log("Copying config")
  let configDestFolder = path.join("/plugins", "tmp", name, "config")
  let configSrcFolder = path.join(pluginFolder, "config")
  console.log("Copying config", "dest", configDestFolder, "src", configSrcFolder)

  fs.removeSync(configDestFolder)
  fs.copySync(configSrcFolder, configDestFolder)

  nextTask()
}

function copyFiles(nextTask) {
  let filesSrcFolder = path.join(pluginFolder, "files")
  if (!fs.pathExistsSync(filesSrcFolder)) {
    nextTask()
    return
  }
  console.log("Copying config")
  let filesDestFolder = path.join("/plugins", "tmp", name, "files")
  fs.removeSync(filesDestFolder)

  fs.copySync(filesSrcFolder, filesDestFolder)
  nextTask()
}

function bundleModule(nextTask) {
  let command = sprintf('tar -czvf %s -C %s %s', path.join("/plugins", "tmp", name+".tar.gz"), path.join("/plugins", "tmp"), name)
  if (shell.exec(command).code !== 0) {
    shell.echo('Could not compress module failed');
    shell.exit(1);
  } else {
    nextTask()
  }
}

function deployModule(nextTask) {
  fs.copySync(path.join("/plugins", "tmp", name +".tar.gz"), path.join(deploymentFolder, name +".tar.gz"))
  nextTask()
}


function startTask(taskName) {
  var func = function(nt){}
  var nextTask = ""
  if (taskName === "copyconfig") {
    func = copyConfig
    nextTask = "objcompile"
  }
  if ( taskName === "objcompile" ){
    func = buildObjects
    nextTask = "uicompile"
  }
  if (taskName === "uicompile" ){
    func = buildUI
    nextTask = "copyfiles"
  }
  if ( taskName === "copyfiles" ){
    func = copyFiles
    nextTask = "bundlemodule"
  }
  if ( taskName === "bundlemodule" ){
    func = bundleModule
    nextTask = "deploymodule"
  }
  if ( taskName === "deploymodule" ){
    func = deployModule
  }
  if ( taskName === null || taskName === "" ){
    return function() {
      console.log("Tasks complete")
    }
  }
  return function() {
    nextTaskFunc = startTask(nextTask)
    console.log("Starting task ", taskName)
    func(nextTaskFunc)
  }
}

buildModule()
