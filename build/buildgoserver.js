var shell = require('shelljs');
var path = require('path');
var fs = require('fs-extra');
var tar = require('tar-stream');
var zlib = require('zlib');
var rimraf = require('rimraf');
var {log, listDir} = require('./utils');
var {argv, name, pluginFolder, goModulesRepo, modulesRepo, uiFolder, filesFolder, modConfig, deploymentFolder, nodeModulesFolder, buildFolder, tmpFolder} = require('./buildconfig');
var sprintf = require('sprintf-js').sprintf
var {execSync, execFileSync, spawnSync} = require('child_process');
var {tidyGoModule} = require('./utils')

function buildGoObjects(nextTask) {
    if(argv.uionly) {
      nextTask()
      return
    }
    log("Beginning golang compile")

    let serverGoSrcFolder = path.join(pluginFolder, "server", "go")
    if(!fs.pathExistsSync(serverGoSrcFolder)) {
      nextTask()
      return
    }
    //tidyGoModule(serverGoSrcFolder)
    var outputFolder;

    if(argv.nobundle) {
      outputFolder = path.join(pluginFolder, "objects")
    } else {
      outputFolder = path.join(tmpFolder, name, "objects")
    }

    fs.removeSync(outputFolder)
  
    fs.mkdirsSync(outputFolder)

    let fileToBuild = sprintf('%s/%s.so', outputFolder, name)

    let optionsArr = ["build", "-buildmode=plugin", "-o", fileToBuild]
    /*let srcfileslist = fs.readdirSync(serverGoSrcFolder)
    srcfileslist.forEach((file)=> {
      optionsArr.push(path.join(serverGoSrcFolder, file));
    })*/
    //optionsArr.push();
    log("running go compile command")
    let res = spawnSync("go", optionsArr, {cwd: serverGoSrcFolder})
    if(res.status !== 0) {
      shell.echo('Golang build unsuccessful');
      shell.echo("Res", res.stdout.toString())
      shell.echo("Err", res.stderr.toString())      
      shell.exit(1);
    } else {
      let fileBuilt = fs.pathExistsSync(fileToBuild)
      log('Golang compilation successfull. Output to', outputFolder, "File built", fileBuilt);
      log(res.stdout.toString())
      log(res.stderr.toString())
      fs.pathExistsSync(fileToBuild)
      nextTask()
    }      

    /*let res = shell.exec(command, function(code, stdout, stderr) {
      log('Code4:', code);
      log('Program output:', stdout);
      log('Program stderr:', stderr);
      if(code !== 0) {
        shell.echo('Golang build failed');
        shell.exit(1);
      } else {
        let fileBuilt = fs.pathExistsSync(fileToBuild)
        shell.echo('Golang compilation successfull. Output to', outputFolder, "File built", fileBuilt);
        fs.pathExistsSync(fileToBuild)
        nextTask()
      }      
    }) */
}

function cleanGoFolders(nextTask) {
  let goSrcFolder = path.join(pluginFolder, "server", "go")
  if (fs.pathExistsSync(goSrcFolder)) {
      fs.readdirSync(goSrcFolder).forEach((item)=> {
          if(item.startsWith("autogen_")) {
              let fileName = path.join(goSrcFolder, item);
              fs.removeSync(fileName)
          }
      })
  }
  let gosdkSrcFolder = path.join(pluginFolder, "sdk", "go")
  if (fs.pathExistsSync(gosdkSrcFolder)) {
      fs.readdirSync(gosdkSrcFolder).forEach((item)=> {
          if(item.startsWith("autogen_")) {
              let fileName = path.join(gosdkSrcFolder, item);
              fs.removeSync(fileName)
          }
      })
  }
  nextTask()
}

function extractFile(modName, modtar, modSDK) {
  log("Extracting module file", modtar)
  var extract = tar.extract();
  var files={}
  var pathToExtract = modName+"/files/sdk/go/"

  extract.on('entry', function(header, stream, cb) {
    stream.on('data', function(chunk) {
        var fileinZip = header.name
        if (fileinZip.startsWith(pathToExtract)) {
          var filename = ''
          filename = fileinZip.substring(pathToExtract.length - 1)
          var data = files[filename]
          if(data) {
            data += chunk;
          } else {
            if(chunk) {
              data = chunk
            } else {
              data = ''
            }
          }
          files[filename] = data
        }
      });
    
      stream.on('end', function() {
          cb();
      });
    
      stream.resume();
  });
  
  return new Promise(function(resolve, reject) {
    extract.on('finish', function() {

      if(files) {
        Object.keys(files).forEach(function(file) {         
          var f = path.parse(file)     
          var dir = path.join(modSDK, f.dir)
          fs.mkdirsSync(dir)
          let fileToWrite = path.join(dir, f.base)
          fs.writeFileSync(fileToWrite, files[file])
        })
      }
      resolve()
    });
    fs.createReadStream(modtar)
      .pipe(zlib.createGunzip())
      .pipe(extract);        
  });
}

function installGoDependencies(nextTask) {
  var depPromises = new Array()
  if(modConfig.dependencies) {
    Object.keys(modConfig.dependencies).forEach(function(modName) {
      log("Analysing module", modName)
      let modSDK = path.join(goModulesRepo, modName)
      if (fs.pathExistsSync(modSDK)) {
        log("Mod SDK exists", modName)
      } else {
        var modtar = path.join(modulesRepo, modName + ".tar.gz")
        if (fs.pathExistsSync(modtar)) {
          let pr = extractFile(modName, modtar, modSDK)            
          depPromises.push(pr);
        } else {
          shell.echo("Err: Could not find dependency ", modName)      
          shell.exit(1);
        }
      }
    })  
  }
  Promise.all(depPromises).then(function(vals) {
    nextTask()
  })
}



function copySDK(nextTask) {
  let sdkSrcFolder = path.join(pluginFolder, "sdk")
  let filesFolder = path.join(pluginFolder, "files")
  let sdkDestFolder = path.join(filesFolder, "sdk")
  if (fs.pathExistsSync(sdkDestFolder)) {
      rimraf.sync(sdkDestFolder)
  }
  if(fs.pathExistsSync(sdkSrcFolder)) {
    fs.mkdirsSync(sdkDestFolder)
    fs.copySync(sdkSrcFolder, sdkDestFolder)
    let goModulesSrc = path.join(sdkSrcFolder, "go")
    if (fs.pathExistsSync(goModulesSrc)) {
      let goModulesDest = path.join(goModulesRepo, name)
      if (fs.pathExistsSync(goModulesDest)) {
        rimraf.sync(goModulesDest)  
      }
      log("Copying SDK to go modules Repo", goModulesDest)
      fs.copySync(goModulesSrc, goModulesDest)
    }  
  }
  nextTask()
}




module.exports = {
  cleanGoFolders: cleanGoFolders,
  copySDK: copySDK,
  installGoDependencies: installGoDependencies,
  buildGoObjects: buildGoObjects
}