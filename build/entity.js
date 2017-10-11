var path = require('path');
var sprintf = require('sprintf-js').sprintf
var fs = require('fs-extra')
var Handlebars = require('handlebars')


function collection(collection, name) {
  return collection? collection: name
}

function cacheable(cacheable) {
  return cacheable? cacheable: false
}

function imports(imports) {
  var importsStr = ""
  imports.forEach(function(pkg) {
    if(pkg) {
      importsStr = importsStr + "\r\n\t\""+ pkg +"\" "
    }
  });
  return importsStr
}

function type(name) {
  return name? name: "entity"
}

function fields(fields) {
  var fieldsStr = ""
  Object.keys(fields).forEach(function(fieldName) {
    let field = fields[fieldName]
    let jsonF = field.json?field.json:fieldName
    let bsonF = field.bson?field.bson:fieldName
    let datastoreF = field.datastore? "datastore:\""+field.datastore+"\"": ""
    switch (field.type) {
      default:
      fieldsStr = fieldsStr + sprintf("\r\n\t%s\t%s `json:\"%s\" bson:\"%s\"  %s`", fieldName, field.type, jsonF, bsonF, datastoreF)
    }
  });
  return fieldsStr
}

function createEntity(entityJson, pluginFolder, filename) {
  let name = entityJson["name"]
  name = name? name +" .go": filename.substring(0, filename.length-5)+".go"
  let filepath = path.join(pluginFolder, name)
  var buf = fs.readFileSync('./tpl/entitygocode.go.tpl');
  Handlebars.registerHelper('cacheable', cacheable);
  Handlebars.registerHelper('imports', imports);
  Handlebars.registerHelper('type', type);
  Handlebars.registerHelper('fields', fields);
  Handlebars.registerHelper('collection', collection);
  var template = Handlebars.compile(buf.toString());
  let gofile = template(entityJson)
  if (fs.pathExistsSync(filepath)) {
    fs.removeSync(filepath)
  }
  fs.writeFileSync(filepath, gofile)
}

function plugins(entities) {
  let str = ""
  entities.forEach(function(entity) {
      str = str + sprintf("core.PluginComponent{Name: \"%s\", Object: %s{}},", entity, entity)
  });
  return str
}

function createManifest(entities, pluginFolder) {
  let filepath = path.join(pluginFolder, "manifest__.go")
  if (!fs.pathExistsSync(filepath)) {
    fs.removeSync(filepath)
  }
  var buf = fs.readFileSync('./tpl/manifest.go.tpl');
  Handlebars.registerHelper('plugins', plugins);
  var template = Handlebars.compile(buf.toString());
  let gofile = template({"entities": entities})
  fs.writeFileSync(filepath, gofile)
}

module.exports = {
  createEntity: createEntity,
  createManifest: createManifest
}
