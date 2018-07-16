var shell = require('shelljs');
var path = require('path');
var fs = require('fs-extra');
var {log} = require('./utils');
var {argv, name, pluginFolder, packageFolder,  uiFolder, filesFolder, modConfig, deploymentFolder, nodeModulesFolder, buildFolder, tmpFolder} = require('./buildconfig');
var sprintf = require('sprintf-js').sprintf

function compileRustWASMUI(nextTask) {
    log("Compiling rust wasm")

    let wasmRustSrcFolder = path.join(pluginFolder, "ui", "rust")
    if(!fs.pathExistsSync(wasmRustSrcFolder)) {
      nextTask()
      return
    }
  
    let tmpObjsFolder = path.join(pluginFolder, "ui", "dist", "wasm")
  
    fs.removeSync(tmpObjsFolder)
  
    fs.mkdirsSync(tmpObjsFolder)

    console.log("folder exists", fs.pathExistsSync(tmpObjsFolder))
  
    let command = sprintf('cargo +nightly build --target-dir %s --manifest-path %s/Cargo.toml --target wasm32-unknown-unknown', tmpObjsFolder, wasmRustSrcFolder)
    console.log("Executing command ", command)
    if (shell.exec(command).code !== 0) {
      shell.echo('Rust wasm build failed');
      shell.exit(1);
    } else {
      shell.echo('Rust wasm compilation successfull');
      nextTask()
    }    
    nextTask()
}

module.exports = {
    compileRustWASMUI: compileRustWASMUI
}
