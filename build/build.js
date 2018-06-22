var shell = require('shelljs');
var path = require('path');
var sprintf = require('sprintf-js').sprintf
var fs = require('fs-extra')
var {argv, name, pluginFolder, packageFolder,  uiFolder, filesFolder, modConfig, deploymentFolder, nodeModulesFolder, buildFolder, tmpFolder} = require('./buildconfig');
var entity = require('./entity')
var {compileJSWebUI, getJSUIModules} = require('./buildjsui')
var {buildGoObjects} = require('./buildgoserver');
var {log} = require('./utils');
var {compileDartUI} = require('./builddart');

function buildModule() {
  createTempDirectory(!(argv.skipUI || argv.skipObjects))
  startTask("copyconfig")()
}

function buildUI(nextTask) {
  if (argv.skipUI || !fs.pathExistsSync(uiFolder)) {
    nextTask()
    return
  }

  let jsUIconfig = modConfig.ui? modConfig.ui.js: null;

  getJSUIModules(jsUIconfig);
  //getUIModules();

  let copyUIFiles = function() {
    log("Copying UI files")
    fs.mkdirsSync(filesFolder)
    if (fs.pathExistsSync(path.join(uiFolder, 'dist'))) {
      fs.copySync(path.join(uiFolder, "dist"), filesFolder)
    }
    nextTask()
  }
  
  compileJSWebUI(jsUIconfig, function() { 
    compileDartUI(copyUIFiles);
  });
}


function copyproperties(nextTask) {
  let propsSrcFolder = path.join(pluginFolder, "properties")
  if (fs.pathExistsSync(propsSrcFolder)) {
    let propsDestFolder = path.join("/plugins", "tmp", name, "properties")
    fs.mkdirsSync(propsDestFolder)
    log("Copying properties", "dest", propsDestFolder, "src", propsSrcFolder)

    fs.removeSync(propsDestFolder)
    fs.copySync(propsSrcFolder, propsDestFolder)

  }

  nextTask()
}

function copyUIRegistry(nextTask) {
  let regSrcFolder = path.join(uiFolder, "registry")
  if (fs.pathExistsSync(regSrcFolder)) {
    let regDestFolder = path.join("/plugins", "tmp", name, "ui")
    fs.mkdirsSync(regDestFolder)
    log("Copying registered items", "dest", regDestFolder, "src", regSrcFolder)
    fs.removeSync(regDestFolder)
    fs.copySync(regSrcFolder, regDestFolder)
  }
  nextTask()
}



function createTempDirectory(removeDir) {
  let tmpFolder = path.join("/plugins", "tmp", name)
  if (removeDir) {
    fs.removeSync(tmpFolder)
  }
  log("Ensuring temp folder ", tmpFolder)
  fs.mkdirsSync(tmpFolder)
}

function buildObjects(nextTask) {
  if (argv.skipObjects) {
    nextTask()
    return
  }

  buildGoObjects(nextTask)
}

function copyConfig(nextTask) {
  log("Copying config")
  let configDestFolder = path.join("/plugins", "tmp", name)
  let configSrcFolder = path.join(pluginFolder, "config")
  log("Copying config", "dest", configDestFolder, "src", configSrcFolder)

  fs.removeSync(configDestFolder)
  fs.copySync(configSrcFolder, configDestFolder)

  nextTask()
}

function autoGen(nextTask) {
  let entities = {}
  if (fs.pathExistsSync(path.join(pluginFolder, 'build'))) {
    let entitiesFolder = path.join(pluginFolder, 'build', "entities")
    if (fs.pathExistsSync(entitiesFolder)) {
      let files = fs.readdirSync(entitiesFolder)
      for(var i=0;i<files.length; i++) {
        if(files[i].endsWith('.json')) {
          let jsonF = path.join(entitiesFolder, files[i])
          let jsonContent = require(jsonF)
          entities[jsonContent["name"]] = jsonContent
          entity.createEntity(jsonContent, pluginFolder, files[i])
        }
      }
    }
  }
  if(entities && Object.keys(entities).length >0) {
    entity.createManifest(entities, pluginFolder)
  }
  nextTask()
}

function copyFiles(nextTask) {
  let filesSrcFolder = path.join(pluginFolder, "files")
  if (!fs.pathExistsSync(filesSrcFolder)) {
    nextTask()
    return
  }
  log("Copying config")
  let filesDestFolder = path.join("/plugins", "tmp", name, "files")
  fs.removeSync(filesDestFolder)

  fs.copySync(filesSrcFolder, filesDestFolder)
  nextTask()
}

function bundleModule(nextTask) {
  let verbose = argv.verbose? "-v":""
  let tarfilepath = path.join("/plugins", "tmp", name+".tar.gz")
  let command = sprintf('tar %s -czf %s -C %s %s', verbose, tarfilepath, path.join("/plugins", "tmp"), name)
  log("Bundle module: ", command)
  if (shell.exec(command).code > 1) { //ignore the exit code for file changed
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
    nextTask = "autogen"
  }
  if (taskName === "autogen") {
    func = autoGen
    nextTask = "objcompile"
  }
  if ( taskName === "objcompile" ){
    func = buildObjects
    nextTask = "copyproperties"
  }
  if ( taskName === "copyproperties" ){
    func = copyproperties
    nextTask = "copyuiregistry"
  }
  if ( taskName === "copyuiregistry" ){
    func = copyUIRegistry
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
      console.log("Tasks complete. Module Built ", argv.name)
    }
  }
  return function() {
    nextTaskFunc = startTask(nextTask)
    log("Starting task ", taskName)
    func(nextTaskFunc)
  }
}

buildModule()
