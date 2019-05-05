var shell = require('shelljs');
var path = require('path');
var fs = require('fs-extra');
var {log, listDir} = require('./utils');
var {argv, name, pluginFolder, packageFolder,  uiFolder, filesFolder, modConfig, deploymentFolder, nodeModulesFolder, buildFolder, tmpFolder} = require('./buildconfig');
var sprintf = require('sprintf-js').sprintf
var {execSync, execFileSync, spawnSync} = require('child_process');

function buildGoObjects(nextTask) {
    if(argv.uionly) {
      nextTask()
      return
    }
    log("Compiling golang")

    let serverGoSrcFolder = path.join(pluginFolder, "server", "go")
    if(!fs.pathExistsSync(serverGoSrcFolder)) {
      nextTask()
      return
    }
    var outputFolder;

    if(argv.nobundle) {
      outputFolder = path.join(pluginFolder, "objects")
    } else {
      outputFolder = path.join("/plugins", "tmp", name, "objects")
    }

    fs.removeSync(outputFolder)
  
    fs.mkdirsSync(outputFolder)

    let fileToBuild = sprintf('%s/%s.so', outputFolder, name)

    let optionsArr = ["build", "-buildmode=plugin", "-o", fileToBuild]
    let srcfileslist = fs.readdirSync(serverGoSrcFolder)
    srcfileslist.forEach((file)=> {
      optionsArr.push(path.join(serverGoSrcFolder, file));
    })
    let res = spawnSync("go", optionsArr)
    if(res.status !== 0) {
      shell.echo('Golang build failed');
      shell.exit(1);
    } else {
      let fileBuilt = fs.pathExistsSync(fileToBuild)
      shell.echo('Golang compilation successfull. Output to', outputFolder, "File built", fileBuilt);
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

module.exports = {
    buildGoObjects: buildGoObjects
}