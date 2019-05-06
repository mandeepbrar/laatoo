var shell = require('shelljs');
var path = require('path');
let webpack = require('webpack')
var sprintf = require('sprintf-js').sprintf
var fs = require('fs-extra');
var {argv, name, pluginFolder, packageFolder,  uiFolder, filesFolder, modConfig, deploymentFolder, nodeModulesFolder, buildFolder, tmpFolder} = require('./buildconfig');
var {log} = require('./utils');

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
  
    //only build vendor file if asked explicitly ie buildDependencies is set
    if(jsUIconfig && jsUIconfig.packages && jsUIconfig.buildDependencies) {
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
              shell.echo("Mod path not found: ", modPath)
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
      if(stats && stats.compilation && stats.compilation.errors && stats.compilation.errors.length !=0 ) {
        console.log("Errors: ", stats.compilation.errors);
        //console.log(stats.compilation)
      } else {
        if(stats.stats && stats.stats.length !=0 ) {
          stats.stats.forEach(function(stat) {
            console.log("Errors: ", stat.compilation.errors);
          })
        }
      }
      log("*UI compilation ends*")
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
        log("Getting package " + pkg)
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
        log(command)
        if (shell.exec(command).code !== 0) {
        shell.echo('Get package failed');
        shell.exit(1);
        }
    } else {
        log("No package json in build")
    }
}

  

module.exports = {
    compileJSWebUI: compileJSWebUI,
    getJSUIModules: getJSUIModules
}