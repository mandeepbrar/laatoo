let webpack = require('webpack')
var shell = require('shelljs');
var path = require('path');
var sprintf = require('sprintf-js').sprintf
var fs = require('fs-extra')
var entity = require('./entity')
var yamlparser = require('js-yaml')
const merge = require('webpack-merge');

var argv = require('minimist')(process.argv.slice(2), {boolean:["verbose", "skipObjects", "skipUI", "overwriteJSMods", "forceUIModules", "printUIConfig", "getBuildPackages"]});

//, default: {skipObjects: false, skipUI: false, skipUIModules: false, printUIConfig: false}

//console.log(argv);

let name = argv.name

let packageFolder = argv.packageFolder

let pluginFolder = path.join("/plugins", "src", packageFolder)

let uiFolder = path.join(pluginFolder, "ui")

let filesFolder = path.join(pluginFolder, "files")

let jsonConfigName = path.join(pluginFolder, "config", "config.json")
let yamlConfigName = path.join(pluginFolder, "config", "config.yml")

let modConfig = null

if (fs.pathExistsSync(jsonConfigName)) {
  let modConfig = require(jsonConfigName);
} else {
  if (fs.pathExistsSync(yamlConfigName)) {
    try {
      modConfig = yamlparser.safeLoad(fs.readFileSync(yamlConfigName, 'utf8'));
    } catch (e) {
      console.log("Could not load module config ", e);
      exit(1)
    }
  }
}


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

  let jsUIconfig = modConfig.ui? modConfig.ui.js: null;

  getJSUIModules(jsUIconfig);
  //getUIModules();

  let copyUIFiles = function() {
    log("Copying UI files")
    fs.mkdirsSync(filesFolder)
    if (fs.pathExistsSync(path.join(uiFolder, 'dist'))) {
      fs.copySync(path.join(uiFolder, "dist"), filesFolder)
    }
    /*if (fs.pathExistsSync(path.join(uiFolder, 'dist/scripts/index.js'))) {
      fs.copySync(path.join(uiFolder, "dist/scripts/index.js"), path.join(filesFolder, "webui.js"))
    }uiFolderuiFolder
    if (fs.pathExistsSync(path.join(uiFolder, 'dist/scripts/vendor.js'))) {
      fs.copySync(path.join(uiFolder, "dist/scripts/vendor.js"), path.join(filesFolder, "vendor.js"))
    }
    if (fs.pathExistsSync(path.join(uiFolder, 'dist/css/app.css'))) {
      fs.copySync(path.join(uiFolder, "dist/css/app.css"), path.join(filesFolder, "app.css"))
    }*/
    nextTask()
  }
  
  compileJSWebUI(jsUIconfig, function() { 
    compileDartUI(copyUIFiles);
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

function copyproperties(nextTask) {
  /*
  let propsSrcFolder = path.join(uiFolder, "src", "properties")
  if (fs.pathExistsSync(propsSrcFolder)) {
    let propsDstFolder = path.join(filesFolder, "properties")
    fs.mkdirsSync(propsDstFolder)
    let defaultProps={}
    if(fs.pathExistsSync(path.join(propsSrcFolder, "default.json"))) {
      defaultProps = require(path.join(propsSrcFolder, "default.json"))
    }
    let files = fs.readdirSync(propsSrcFolder)
    for(var i=0;i<files.length; i++) {
      if(files[i].endsWith('.json')) {
        let fileCont = require(path.join(propsSrcFolder, files[i]))
        fileCont = merge(defaultProps, fileCont)
        let contToWrite = JSON.stringify(fileCont);
        fs.writeFileSync(path.join(propsDstFolder, files[i]), contToWrite)
      }
    }
  }*/
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

function compileDartUI(nextTask) {
  nextTask()
}

function compileJSWebUI(jsUIconfig, nextTask) {
  let jsuiFolder = path.join(uiFolder, 'js')
  if(!fs.pathExistsSync(jsuiFolder)) {
    nextTask()
    return
  }

  let configFunc = require(path.join(buildFolder, 'cfg/webpack.lib'))

  let externals = jsUIconfig? jsUIconfig.externals:null

  let options = {
    library: name,
    uifolder: uiFolder,
    externals: externals
  }

  if(jsUIconfig && jsUIconfig.packages ) {
    options.dependencies = jsUIconfig.packages
  }

  let config = configFunc(options)
  
  if (fs.pathExistsSync(path.join(jsuiFolder, 'build.js'))) {
    let custom = require(path.join(jsuiFolder, 'build'))
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
      /*if (fs.pathExistsSync(laatooModulePath)) {
        found = true
      }*/
      if (!found && fs.pathExistsSync(laatooModulePath + '.tar.gz')) {
        let command = sprintf("cd %s && tar %s -xzf %s", tmpFolder, verbose, laatooModulePath + '.tar.gz')
        log("Extracting pkg", command)
        if (shell.exec(command).code !== 0) {
          shell.echo('Package extraction failed');
          shell.exit(1);
        } else {
          modPath = path.join(tmpFolder, pkg)
          if (fs.pathExistsSync(modPath)) {
            log("Mod found: ", pkg)
            found = true
          } else {
            console.log("Mod path not found: ", modPath)
          }
        }
      }
      if(found) {
        let uiSrc = path.join(modPath, "files", "scripts", "index.js")
        if (fs.pathExistsSync(uiSrc)) {
          log("Dependency " + pkg + " for ui found")
          let dest = path.join(nodeModulesFolder, "node_modules", pkg)
          fs.mkdirsSync(dest)
          fs.copySync(uiSrc, path.join(dest, "index.js"))
        }
        let uicss = path.join(modPath, "files", "css", "app.css")
        if (fs.pathExistsSync(uicss)) {
          let dest = path.join(nodeModulesFolder, "node_modules", pkg)
          fs.mkdirsSync(dest)
          fs.copySync(uicss, path.join(dest, "css", "app.css"))
          log("Css being copied for pkg", pkg)
        }
      } else {
        log("No ui found for Dependency ", pkg)
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

function processPackageJson() {
  let silent = argv.verbose?"":"-s";
  if (fs.pathExistsSync(path.join(nodeModulesFolder,'package.json'))) {
    log("Installing package json from nodemodules")
    var command = sprintf('npm i %s', silent)
    try {
      let wd = process.cwd();
      process.chdir(nodeModulesFolder);
      process.chdir(wd);
    } catch(ex){}
    log("Command ", command)
    if (shell.exec(command).code !== 0) {
      shell.echo('Get package failed');
      shell.exit(1);
    }
  }
}

function getJSUIModules(jsUIconfig) {
  let silent = argv.verbose?"":"-s";
  if (argv.getBuildPackages) {
    let wd = process.cwd(); 
    process.chdir(nodeModulesFolder);
    let pkgsrc = path.join(buildFolder,'package.json')
    fs.copySync(pkgsrc, path.join(nodeModulesFolder,'package.json'))
    processPackageJson()
    process.chdir(wd);
  }

  if(argv.getBuildPackages) {
    try {
      let wd = process.cwd();
      process.chdir(nodeModulesFolder);
      if (shell.exec("npm rebuild node-sass").code !== 0) {
        shell.echo('npm rebuild failed');
        shell.exit(1);
      } else {
        log("Npm rebuild sass successfull");
      }
      process.chdir(wd);
    } catch(ex){}  
  }

  if(jsUIconfig && jsUIconfig.packages !=null) {
    log("Installing package json from plugin", jsUIconfig.packages)
    Object.keys(jsUIconfig.packages).forEach(function(pkg) {
      console.log("Getting package " + pkg)
      let ver = jsUIconfig.packages[pkg]
      let existingPkg = path.join(nodeModulesFolder, "node_modules", pkg)
      if(!fs.pathExistsSync(existingPkg) || argv.overwriteJSMods) {
        let command = sprintf('npm i %s --prefix %s %s@%s', silent, nodeModulesFolder, pkg, ver)
        console.log(command)
        if (shell.exec(command).code !== 0) {
          shell.echo('Get package failed '+pkg);
          shell.exit(1);
        }  
      }
    });
  }

  if (argv.getBuildPackages && fs.pathExistsSync(path.join(buildFolder,'package.json'))) {
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

  let serverGoSrcFolder = path.join(pluginFolder, "server", "go")
  if(!fs.pathExistsSync(serverGoSrcFolder)) {
    nextTask()
    return
  }

  let tmpObjsFolder = path.join("/plugins", "tmp", name, "objects")

  fs.removeSync(tmpObjsFolder)

  fs.mkdirsSync(tmpObjsFolder)

  let command = sprintf('go build -buildmode=plugin -o %s/%s.so %s/*.go', tmpObjsFolder, name, serverGoSrcFolder)
  if (shell.exec(command).code !== 0) {
    shell.echo('Golang build failed');
    shell.exit(1);
  } else {
    shell.echo('Golang compilation successfull');
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

function autoGen(nextTask) {
  let entities = {}
  if (fs.pathExistsSync(path.join(pluginFolder, 'build'))) {
    let entitiesFolder = path.join(pluginFolder, 'build', "entities")
    if (fs.pathExistsSync(entitiesFolder)) {
      let files = fs.readdirSync(pluginFolderentitiesFolder)
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
