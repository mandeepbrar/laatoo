
var path = require('path');
var sprintf = require('sprintf-js').sprintf
var fs = require('fs-extra')
var yamlparser = require('js-yaml')

var argv = require('minimist')(process.argv.slice(2), {boolean:["norust", "verbose", "nobundle", "release", "skipObjects", "skipUI", "overwriteJSMods", "forceUIModules", "printUIConfig", "getBuildPackages", "uionly"]});

//, default: {skipObjects: false, skipUI: false, skipUIModules: false, printUIConfig: false}

var name = argv.name

var pluginFolder

if(argv.packageFolder.startsWith("/")) {
  pluginFolder = argv.packageFolder
} else {
  pluginFolder = path.join("/modulesrepo", argv.packageFolder)
}

var uiFolder = path.join(pluginFolder, "ui")
var uiBuildFolder = path.join(pluginFolder, "uibuild")

var filesFolder = path.join(pluginFolder, "files")

var release = argv.release

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

var buildFolder = argv.buildFolder? argv.buildFolder: "/build"

var tmpFolder = "/tmp/compiletmp"

module.exports = {
    argv: argv,
    name: name,
    pluginFolder: pluginFolder,
    release: release,
    uiFolder: uiFolder,
    uiBuildFolder: uiBuildFolder,
    filesFolder: filesFolder,
    modConfig: modConfig,
    deploymentFolder: deploymentFolder,
    nodeModulesFolder: nodeModulesFolder,
    modulesRepo: deploymentFolder,
    goModulesRepo: "/laatoo/sdk/modules", 
    buildFolder: buildFolder,
    tmpFolder: tmpFolder
}