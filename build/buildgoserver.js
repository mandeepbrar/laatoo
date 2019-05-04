var shell = require('shelljs');
var path = require('path');
var fs = require('fs-extra');
var {log} = require('./utils');
var {argv, name, pluginFolder, packageFolder,  uiFolder, filesFolder, modConfig, deploymentFolder, nodeModulesFolder, buildFolder, tmpFolder} = require('./buildconfig');
var sprintf = require('sprintf-js').sprintf

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

module.exports = {
    buildGoObjects: buildGoObjects
}