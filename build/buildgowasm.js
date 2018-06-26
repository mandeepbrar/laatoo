var shell = require('shelljs');
var path = require('path');
var fs = require('fs-extra');
var {log} = require('./utils');
var {argv, name, pluginFolder, packageFolder,  uiFolder, filesFolder, modConfig, deploymentFolder, nodeModulesFolder, buildFolder, tmpFolder} = require('./buildconfig');
var sprintf = require('sprintf-js').sprintf

function compileGoWASMUI(nextTask) {
    log("Compiling go wasm")

    let wasmGoSrcFolder = path.join(pluginFolder, "ui", "go")
    if(!fs.pathExistsSync(wasmGoSrcFolder)) {
      nextTask()
      return
    }
  
    let tmpObjsFolder = path.join(pluginFolder, "ui", "dist", "wasm")
  
    fs.removeSync(tmpObjsFolder)
  
    fs.mkdirsSync(tmpObjsFolder)

    console.log("folder exists", fs.pathExistsSync(tmpObjsFolder))
  
    let command = sprintf('GOOS=js GOARCH=wasm go build -o %s/%s.wasm %s/*.go', tmpObjsFolder, name, wasmGoSrcFolder)
    if (shell.exec(command).code !== 0) {
      shell.echo('Golang build failed');
      shell.exit(1);
    } else {
      shell.echo('Golang compilation successfull');
      nextTask()
    }    
    nextTask()
}

module.exports = {
    compileGoWASMUI: compileGoWASMUI
}
