let webpack = require('webpack')
var shell = require('shelljs');
var path = require('path');
var sprintf = require('sprintf-js').sprintf
var fs = require('fs-extra')

var argv = require('minimist')(process.argv.slice(2), {boolean:["verbose", "skipObjects", "skipUI", "skipUIModules", "printUIConfig"]});

//, default: {skipObjects: false, skipUI: false, skipUIModules: false, printUIConfig: false}

//console.log(argv);

let name = argv.name

let packageFolder = argv.packageFolder

let pluginFolder = path.join("/plugins", "src", packageFolder)

let uiFolder = path.join(pluginFolder, "ui")

let filesFolder = path.join(pluginFolder, "files")

let modConfig = require(path.join(pluginFolder, "config", "config.json"));

let deploymentFolder = "/deploy/"

let nodeModulesFolder = "/nodemodules"

let buildFolder = "/build"

let tmpFolder = "/tmp"

function log() {
  let args = Array.prototype.slice.call(arguments);
  if(argv.verbose) {
    console.log(args.join(' '));
  }
}

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
    if (!modConfig.ui || !modConfig.ui.skipUIModules) {
      getUIModules();
    }
  }

  compileWebUI(function() {
    log("Copying UI files")
    fs.mkdirsSync(filesFolder)
    if (fs.pathExistsSync(path.join(uiFolder, 'dist/scripts/index.js'))) {
      fs.copySync(path.join(uiFolder, "dist/scripts/index.js"), path.join(filesFolder, "webui.js"))
    }
    if (fs.pathExistsSync(path.join(uiFolder, 'dist/scripts/vendor.js'))) {
      fs.copySync(path.join(uiFolder, "dist/scripts/vendor.js"), path.join(filesFolder, "vendor.js"))
    }
    nextTask()
  });
}
/*
function buildDll(nextTask) {
  if (modConfig.ui && modConfig.ui.buildDll) {
    let config = require(path.join(buildFolder, 'cfg/webpack.dll'))

    let compiler = webpack(config)

    fs.removeSync(path.join(nodeModulesFolder,'dll'))

    log("Removed directory dll")

    fs.mkdirsSync(path.join(nodeModulesFolder,'dll'))

    compiler.run(function(err, stats) {
      if(stats.compilation.errors.length != 0) {
        console.log("Errors: ", stats.compilation.errors);
      }
      nextTask()
    });

  } else {
    nextTask()
  }
}*/

function compileWebUI(nextTask) {

  let configFunc = require(path.join(buildFolder, 'cfg/webpack.lib'))

  let externals = modConfig.ui? modConfig.ui.externals:null

  let options = {
    library: name,
    uifolder: uiFolder,
    externals: externals
  }

  if(modConfig.ui && modConfig.ui.packages && modConfig.ui.buildDependencies ) {
    options.dependencies = modConfig.ui.packages
  }

  let config = configFunc(options)
  if (fs.pathExistsSync(path.join(uiFolder, 'build.js'))) {
    let custom = require(path.join(uiFolder, 'build'))
    if (custom.config!=nil) {
      config = custom.config(config)
    }
  } else {
    log("No custom ui build file")
  }

  if(argv.printUIConfig) {
    log("UI Config", config)
  }

  if(modConfig.dependencies!=null ) {
    let verbose = argv.verbose?"-v":"";
    log("Processing required laatoo modules")
    Object.keys(modConfig.dependencies).forEach(function(pkg) {
      let found = false
      let laatooModulePath = path.join(deploymentFolder, pkg)
      let modPath = laatooModulePath
      if (fs.pathExistsSync(laatooModulePath)) {
        found = true
      }
      if (!found && fs.pathExistsSync(laatooModulePath + '.tar.gz')) {
        let command = sprintf("cd %s && tar %s -xf %s", tmpFolder, verbose, laatooModulePath + '.tar.gz')
        if (shell.exec(command).code !== 0) {
          shell.echo('Package extraction failed');
          shell.exit(1);
        } else {
          modPath = path.join(tmpFolder, pkg)
          if (fs.pathExistsSync(modPath)) {
            found = true
          }
        }
      }
      if(found) {
        let uiSrc = path.join(modPath, "files", "webui.js")
        if (fs.pathExistsSync(uiSrc)) {
          log("Dependency " + pkg + " for ui found")
          let dest = path.join(nodeModulesFolder, "node_modules", pkg)
          fs.mkdirsSync(dest)
          fs.copySync(uiSrc, path.join(dest, "index.js"))
        }
      }
    });
  }

  let compiler = webpack(config)

  fs.removeSync(path.join(uiFolder,'dist'))

  log("Removed directory dist")

  fs.mkdirsSync(path.join(uiFolder,'dist/scripts'))

  log("Starting compilation", __dirname)
  compiler.run(function(err, stats) {
    if(stats && stats.compilation ) {
      console.log("Errors: ", stats.compilation.errors);
      //console.log(stats.compilation)
    } else {
      if(stats.stats) {
        stats.stats.forEach(function(stat) {
          console.log("Errors: ", stat.compilation.errors);
        })
      }
    }

    nextTask()
  });
}

function getUIModules() {
  let silent = argv.verbose?"":"-s";
  if (fs.pathExistsSync(path.join(nodeModulesFolder,'package.json'))) {
    log("Installing package json from nodemodules")
    var command = sprintf('cd %s && npm i %s --prefix %s', nodeModulesFolder, silent, nodeModulesFolder)
    log("Command ", command)
    if (shell.exec(command).code !== 0) {
      shell.echo('Get package failed');
      shell.exit(1);
    }
  } else {
    console.log("No package json in nodemodules")
  }

  if(modConfig.ui!=null && modConfig.ui.packages !=null) {
    log("Installing package json from plugin", modConfig.ui.packages)
    Object.keys(modConfig.ui.packages).forEach(function(pkg) {
      let ver = modConfig.ui.packages[pkg]
      let command = sprintf('npm i %s --prefix %s %s@%s', silent, nodeModulesFolder, pkg, ver)
      console.log(command)
      if (shell.exec(command).code !== 0) {
        shell.echo('Get package failed '+pkg);
        shell.exit(1);
      }
    });
  }

  if (fs.pathExistsSync(path.join(buildFolder,'package.json'))) {
    log("Installing package json from build folder")
    log(sprintf)
    let command = sprintf('cd %s && npm i %s --prefix %s', buildFolder, silent, nodeModulesFolder)
    console.log(command)
    if (shell.exec(command).code !== 0) {
      shell.echo('Get package failed');
      shell.exit(1);
    }
  } else {
    log("No package json in build")
  }
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

  log("Compiling golang")

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
  log("Copying config")
  let configDestFolder = path.join("/plugins", "tmp", name)
  let configSrcFolder = path.join(pluginFolder, "config")
  log("Copying config", "dest", configDestFolder, "src", configSrcFolder)

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
  log("Copying config")
  let filesDestFolder = path.join("/plugins", "tmp", name, "files")
  fs.removeSync(filesDestFolder)

  fs.copySync(filesSrcFolder, filesDestFolder)
  nextTask()
}

function bundleModule(nextTask) {
  let verbose = argv.verbose? "-v":""
  let command = sprintf('tar %s -czf %s -C %s %s', verbose, path.join("/plugins", "tmp", name+".tar.gz"), path.join("/plugins", "tmp"), name)
  log("Bundle module: ", command)
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
  /*if ( taskName === "builddll" ){
    func = buildDll
    nextTask = "uicompile"
  }*/
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
