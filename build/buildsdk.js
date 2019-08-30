var fs = require('fs-extra');
var path = require('path');
var rimraf = require('rimraf');
var {name, pluginFolder} = require('./buildconfig');

function cleanGoFolders(nextTask) {
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


function copySDK(nextTask) {
    let sdkSrcFolder = path.join(pluginFolder, "sdk")
    let filesFolder = path.join(pluginFolder, "files")
    let sdkDestFolder = path.join(filesFolder, "sdk")
    if (fs.pathExistsSync(sdkDestFolder)) {
        rimraf.sync(sdkDestFolder)
    }
    fs.copySync(sdkSrcFolder, sdkDestFolder)
    nextTask()
}




module.exports = {
    cleanGoFolders: cleanGoFolders,
    copySDK: copySDK,
}