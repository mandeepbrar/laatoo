var shell = require('shelljs');
var path = require('path');
var sprintf = require('sprintf-js').sprintf
var fs = require('fs-extra')
var {argv, name, pluginFolder, uiFolder, uiBuildFolder, filesFolder, modConfig, deploymentFolder, nodeModulesFolder, buildFolder, tmpFolder} = require('./buildconfig');
var entity = require('./entity')
var {compileJSWebUI, getJSUIModules} = require('./buildjsui')
var {buildGoObjects, cleanGoFolders, copySDK} = require('./buildgoserver');
var {log, clearDirectory} = require('./utils');
var {compileDartUI} = require('./builddart');
var {compileGoWASMUI} = require('./buildgowasm');
var {compileRustWASMUI} = require('./buildrustwasm');

function buildModule() {
  if(argv.verbose) {
    console.log(argv)
  }
  if(!(argv.skipUI || argv.skipObjects)) {
    clearDirectory(tmpFolder)
  }
  startTask("cleango")()
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
    if (fs.pathExistsSync(uiBuildFolder)) {
      fs.copySync(uiBuildFolder, filesFolder)
    }
    nextTask()
  }
  
  compileJSWebUI(jsUIconfig, function() {     
    compileRustWASMUI(function() { 
      compileGoWASMUI(copyUIFiles)
    });
  });
}


function copyproperties(nextTask) {
  if(argv.nobundle) {
    nextTask()
    return
  }
  let propsSrcFolder = path.join(pluginFolder, "properties")
  if (fs.pathExistsSync(propsSrcFolder)) {
    let propsDestFolder = path.join(tmpFolder, name, "properties")
    fs.mkdirsSync(propsDestFolder)
    log("Copying properties", "dest", propsDestFolder, "src", propsSrcFolder)

    fs.removeSync(propsDestFolder)
    fs.copySync(propsSrcFolder, propsDestFolder)

  }

  nextTask()
}

function copyUIRegistry(nextTask) {
  if(argv.nobundle) {
    nextTask()
    return
  }
  let regSrcFolder = path.join(uiFolder, "registry")
  if (fs.pathExistsSync(regSrcFolder)) {
    let regDestFolder = path.join(tmpFolder, name, "ui", "registry")
    fs.mkdirsSync(regDestFolder)
    log("Copying registered items", "dest", regDestFolder, "src", regSrcFolder)
    fs.removeSync(regDestFolder)
    fs.copySync(regSrcFolder, regDestFolder)
  }
  nextTask()
}


function buildObjects(nextTask) {
  if (argv.skipObjects) {
    nextTask()
    return
  }

  buildGoObjects(nextTask)
}

function copyConfig(nextTask) {
  if(argv.nobundle) {
    nextTask()
    return
  }
  let configDestFolder = path.join(tmpFolder, name, "config")
  let configSrcFolder = path.join(pluginFolder, "config")
  log("Copying config", "dest", configDestFolder, "src", configSrcFolder)

  fs.removeSync(configDestFolder)
  fs.copySync(configSrcFolder, configDestFolder)

  nextTask()
}

function autoGen(nextTask) {
  if(argv.uionly) {
    nextTask()
    return
  }
  let entities = {}
  if (fs.pathExistsSync(path.join(pluginFolder, 'build'))) {
    let entitiesFolder = path.join(pluginFolder, 'build', "entities")
    if (fs.pathExistsSync(entitiesFolder)) {
      let files = fs.readdirSync(entitiesFolder)
      /*fs.removeSync(autogenFolder)
      if(files.length>0) {
        fs.mkdirsSync(autogenFolder)
      }*/
      for(var i=0;i<files.length; i++) {
        if(files[i].endsWith('.json')) {
          let jsonF = path.join(entitiesFolder, files[i])
          let jsonContent = require(jsonF)
          entities[jsonContent["name"]] = jsonContent
          entity.createEntity(jsonContent, files[i])
        }
      }
      if(entities && Object.keys(entities).length >0) {
        entity.createManifest(entities, name, pluginFolder)
      }    
    }
  }
  nextTask()
}

function copyFiles(nextTask) {
  if(argv.nobundle) {
    nextTask()
    return
  }
  let filesSrcFolder = path.join(pluginFolder, "files")
  if (!fs.pathExistsSync(filesSrcFolder)) {
    nextTask()
    return
  }
  log("Copying files")
  let filesDestFolder = path.join(tmpFolder, name, "files")
  fs.removeSync(filesDestFolder)

  fs.copySync(filesSrcFolder, filesDestFolder)
  nextTask()
}

function bundleModule(nextTask) {
  if(argv.nobundle) {
    nextTask()
    return
  }
  let verbose = argv.verbose? "-v":""
  let tarfilepath = path.join(tmpFolder, name+".tar.gz")
  let command = sprintf('tar %s -czf %s -C %s %s', verbose, tarfilepath, tmpFolder, name)
  log("Bundle module: ", command)
  if (shell.exec(command).code > 1) { //ignore the exit code for file changed
    shell.echo('Could not compress module failed');
    shell.exit(1);
  } else {
    nextTask()
  }
}

function deployModule(nextTask) {
  if(!argv.nobundle) {
    fs.copySync(path.join(tmpFolder, name +".tar.gz"), path.join(deploymentFolder, name +".tar.gz"))
  }
  nextTask()
}


function startTask(taskName) {
  var func = function(nt){}
  var nextTask = ""
  if (taskName === "cleango") {
    func = cleanGoFolders
    nextTask = "copyconfig"    
  }
  if (taskName === "copyconfig") {
    func = copyConfig
    nextTask = "autogen"
  }
  if (taskName === "autogen" ) {
    func = autoGen
    nextTask = "objcompile"
  }
  if ( taskName === "objcompile"){
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
    nextTask = "copysdk"
  }
  if ( taskName === "copysdk" ){
    func = copySDK
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
    log("-----------------------------")
    log("Starting task ", taskName)
    log("-----------------------------")
    func(nextTaskFunc)
  }
}

buildModule()
