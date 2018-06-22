var shell = require('shelljs');
var path = require('path');
var fs = require('fs-extra');
var {log} = require('./utils');
var {argv, name, pluginFolder, packageFolder,  uiFolder, filesFolder, modConfig, deploymentFolder, nodeModulesFolder, buildFolder, tmpFolder} = require('./buildconfig');
var sprintf = require('sprintf-js').sprintf

function compileDartUI(nextTask) {
    nextTask()
}

module.exports = {
    compileDartUI: compileDartUI
}