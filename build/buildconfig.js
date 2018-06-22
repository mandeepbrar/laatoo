
var path = require('path');
var sprintf = require('sprintf-js').sprintf
var fs = require('fs-extra')
var yamlparser = require('js-yaml')

var argv = require('minimist')(process.argv.slice(2), {boolean:["verbose", "skipObjects", "skipUI", "overwriteJSMods", "forceUIModules", "printUIConfig", "getBuildPackages"]});

//, default: {skipObjects: false, skipUI: false, skipUIModules: false, printUIConfig: false}

//console.log(argv);

var name = argv.name

var packageFolder = argv.packageFolder

var pluginFolder = path.join("/plugins", "src", packageFolder)

var uiFolder = path.join(pluginFolder, "ui")

var filesFolder = path.join(pluginFolder, "files")

var jsonConfigName = path.join(pluginFolder, "config", "config.json")
var yamlConfigName = path.join(pluginFolder, "config", "config.yml")

var modConfig = null

if (fs.pathExistsSync(jsonConfigName)) {
  let modConfig = require(jsonConfigName);
} else {
  if (fs.pathExistsSync(yamlConfigName)) {
    try {
      modConfig = yamlparser.safeLoad(fs.readFileSync(yamlConfigName, 'utf8'));
    } catch (e) {
      console.log("Could not load module config ", e);
      process.exit(1)
    }
  }
}


var deploymentFolder = "/deploy/"

var nodeModulesFolder = "/nodemodules"

var buildFolder = "/build"

var tmpFolder = "/tmp"

module.exports = {
    argv: argv,
    name: name,
    pluginFolder: pluginFolder,
    packageFolder: packageFolder,
    uiFolder: uiFolder,
    filesFolder: filesFolder,
    modConfig: modConfig,
    deploymentFolder: deploymentFolder,
    nodeModulesFolder: nodeModulesFolder,
    buildFolder: buildFolder,
    tmpFolder: tmpFolder
}