var shell = require('shelljs');
var path = require('path');
var fs = require('fs-extra');
var {log} = require('./utils');
var {argv, name, pluginFolder, packageFolder, release, uiFolder, filesFolder, modConfig, deploymentFolder, nodeModulesFolder, buildFolder, tmpFolder} = require('./buildconfig');
var sprintf = require('sprintf-js').sprintf

function compileRustWASMUI(nextTask) {
    log("Compiling rust wasm")

    let wasmRustSrcFolder = path.join(pluginFolder, "ui", "rust")
    if(!fs.pathExistsSync(wasmRustSrcFolder)) {
      nextTask()
      return
    }
  
    let tmpObjsFolder = path.join(pluginFolder, "ui", "rust", "build")
  
    fs.removeSync(tmpObjsFolder)
  
    fs.mkdirsSync(tmpObjsFolder)

    let cargoHome = path.join(wasmRustSrcFolder, "cargohome")    

    let target = "wasm32-unknown-unknown";
    console.log("folder exists", fs.pathExistsSync(tmpObjsFolder))
    let mode = "debug"
    let command = sprintf('CARGO_HOME=%s CARGO_INCREMENTAL=1 cargo +nightly build --target-dir %s --manifest-path %s/Cargo.toml --target %s', cargoHome, tmpObjsFolder, wasmRustSrcFolder, target)
    if(release) {
      mode = "release";
      command = sprintf('CARGO_HOME=%s cargo +nightly build --release --target-dir %s --manifest-path %s/Cargo.toml --target %s', cargoHome, tmpObjsFolder, wasmRustSrcFolder, target)
    }
    console.log("Executing command ", command)
    if (shell.exec(command).code !== 0) {
      shell.echo('Rust wasm build failed');
      shell.exit(1);
    } else {
      shell.echo('Rust wasm compilation successfull');
      let wasmFolder = path.join(tmpObjsFolder, target, mode)
      let wasmFile = path.join(wasmFolder, name + ".wasm")
      if(fs.pathExistsSync(wasmFile)) {
        console.log("starting wasm bindgen",wasmFile)
        let bindgenCmd = sprintf("wasm-bindgen %s --no-modules --no-modules-global %s_wasm --out-dir %s ", wasmFile, name, wasmFolder)
        if (shell.exec(bindgenCmd).code !== 0) {
          shell.echo('Rust wasm-bindgen failed');
          shell.exit(1);
        } else {
          let wasmBGFile = path.join(wasmFolder, name + "_bg.wasm")
          if(fs.pathExistsSync(wasmBGFile)) {
            log("Copying Wasm files")
            let distWasmFolder = path.join(pluginFolder, "ui", "dist", "wasm")
            let distFolderFile = path.join(distWasmFolder, name + ".wasm")
            fs.mkdirsSync(distWasmFolder)
            fs.copySync(wasmBGFile, distFolderFile)
            let wasmJsFile = path.join(wasmFolder, name + ".js")
            let distJSFile = path.join(distWasmFolder, name + ".js")
            fs.copySync(wasmJsFile, distJSFile)
          } else {
            log("Wasm file not found")
            shell.exit(1);
          }
        }
      }  

/*      console.log("starting wasm bindgen")
      let wasmFolder = path.join(tmpObjsFolder, target, mode)
      let wasmFile = path.join(wasmFolder, name + ".wasm")
      console.log("starting wasm bindgen",wasmFile)

      let bindgenCmd = sprintf("wasm-bindgen %s --out-dir %s ", wasmFile, wasmFolder)
      if (shell.exec(bindgenCmd).code !== 0) {
        shell.echo('Rust wasm-bindgen failed');
        shell.exit(1);
      } else {
        shell.echo('Rust wasm compilation successfull');
        let wasmBGFile = path.join(wasmFolder, name + "_bg.wasm")
      }*/
    }    
    nextTask()
}

module.exports = {
    compileRustWASMUI: compileRustWASMUI
}
