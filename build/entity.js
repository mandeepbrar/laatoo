var path = require('path');
var sprintf = require('sprintf-js').sprintf
var fs = require('fs-extra')
var Handlebars = require('handlebars')
var {log} = require('./utils');
var {buildFolder} = require('./buildconfig');

function collection(collection, name) {
  return collection? collection: name
}

function cacheable(cacheable) {
  return cacheable? cacheable: false
}

function multitenant(multitenant) {
  return multitenant? multitenant: false
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

function titleField(titleField) {
  return titleField? titleField: "Title"
}

function fields(fields) {
  var fieldsStr = ""
  Object.keys(fields).forEach(function(fieldName) {
    let field = fields[fieldName]
    let jsonF = field.json?field.json:fieldName
    let bsonF = field.bson?field.bson:fieldName
    let datastoreF = field.datastore? field.datastore : fieldName
    switch (field.type) {
      case "storable":
        if(field.list) {
          fieldsStr = fieldsStr + sprintf("\r\n\t%s\t[]%s `json:\"%s\" bson:\"%s\" datastore: \"%s\"`", fieldName, "data.Storable", jsonF, bsonF, datastoreF)
        } else {
          fieldsStr = fieldsStr + sprintf("\r\n\t%s\t%s `json:\"%s\" bson:\"%s\" datastore: \"%s\"`", fieldName, "data.Storable", jsonF, bsonF, datastoreF)
        }
        break;
      case "storableref":
      if(field.list) {
        fieldsStr = fieldsStr + sprintf("\r\n\t%s\t[]%s `json:\"%s\" bson:\"%s\" datastore: \"%s\"`", fieldName, "data.StorableRef", jsonF, bsonF, datastoreF)
      } else {
        fieldsStr = fieldsStr + sprintf("\r\n\t%s\t%s `json:\"%s\" bson:\"%s\" datastore: \"%s\"`", fieldName, "data.StorableRef", jsonF, bsonF, datastoreF)
      }
      break;
      case "entity":
        if(field.list) {
          fieldsStr = fieldsStr + sprintf("\r\n\t%s\t[]%s `json:\"%s\" bson:\"%s\" datastore:\"%s\"`", fieldName, "*" + field.entity, jsonF, bsonF, datastoreF)
        } else {
          fieldsStr = fieldsStr + sprintf("\r\n\t%s\t%s `json:\"%s\" bson:\"%s\" datastore:\"%s\"`", fieldName, "*" + field.entity, jsonF, bsonF, datastoreF)
        }
        break;
      case "subentity":
        if(field.list) {
          fieldsStr = fieldsStr + sprintf("\r\n\t%s\t[]%s `json:\"%s\" bson:\"%s\" datastore: \"%s\"`", fieldName, field.entity, jsonF, bsonF, datastoreF)
        } else {
          fieldsStr = fieldsStr + sprintf("\r\n\t%s\t%s `json:\"%s\" bson:\"%s\" datastore: \"%s\"`", fieldName, field.entity, jsonF, bsonF, datastoreF)
        }
        break;
      case "map":
        if(field.mappedElement) {
          fieldsStr = fieldsStr + sprintf("\r\n\t%s\tmap[string]%s `json:\"%s\" bson:\"%s\" datastore:\"%s\"`", fieldName, field.mappedElement, jsonF, bsonF, datastoreF)
        } else {
          fieldsStr = fieldsStr + sprintf("\r\n\t%s\tmap[string]interface{} `json:\"%s\" bson:\"%s\" datastore:\"%s\"`", fieldName, jsonF, bsonF, datastoreF)
        }
        break;
      case "stringmap":
        fieldsStr = fieldsStr + sprintf("\r\n\t%s\tmap[string]string `json:\"%s\" bson:\"%s\" datastore:\"%s\"`", fieldName, jsonF, bsonF, datastoreF)
        break;
      default:
        if(field.list) {
          fieldsStr = fieldsStr + sprintf("\r\n\t%s\t[]%s `json:\"%s\" bson:\"%s\" datastore:\"%s\"`", fieldName, field.type, jsonF, bsonF, datastoreF)
        } else {
          fieldsStr = fieldsStr + sprintf("\r\n\t%s\t%s `json:\"%s\" bson:\"%s\" datastore:\"%s\"`", fieldName, field.type, jsonF, bsonF, datastoreF)
        }
    }
  });
  return fieldsStr
}

function autogenFolder(pluginFolder) {
  return path.join(pluginFolder, "server", "go")
}


function createEntity(entityJson, autogenFolder, filename) {
  let name = entityJson["name"]
  name = name? name +".go": filename.substring(0, filename.length-5)+".go"
  fs.mkdirsSync(autogenFolder)
  let filepath = path.join(autogenFolder, "autogen_" + name)
  let tplpath = path.join(buildFolder, 'tpl/entitygocode.go.tpl');
  if(entityJson.collection == "<nocollection>") {
    tplpath = path.join(buildFolder, 'tpl/entitynodb.go.tpl');
  }
  var buf = fs.readFileSync(tplpath);
  Handlebars.registerHelper('cacheable', cacheable);
  Handlebars.registerHelper('multitenant', multitenant);
  Handlebars.registerHelper('imports', imports);
  Handlebars.registerHelper('type', type);
  Handlebars.registerHelper('titleField', titleField);
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
  Object.keys(entities).forEach(function(entity) {
      let entityJson = entities[entity]
      let entityDesc = JSON.stringify(entityJson).replace(/\"/g,'\\"')
      str = str + sprintf("core.PluginComponent{Name: \"%s\", Object: %s{}, Metadata: core.NewInfo(\"\",\"%s\", map[string]interface{}{\"descriptor\":\"%s\"})},", entity, entity, entity, entityDesc)
  });
  return str
}

function createManifest(entities, autogenFolder, pluginFolder) {
  let manifestpath = path.join(pluginFolder, "server", "go", "manifest.go")
  if (!fs.pathExistsSync(manifestpath)) {
    var buf = fs.readFileSync(path.join(buildFolder, '/tpl/manifest.go.tpl'));
    var template = Handlebars.compile(buf.toString());
    let gofile = template({})
    fs.writeFileSync(manifestpath, gofile)
  }
  let objectspath = path.join(autogenFolder, "autogen_objectsmanifest.go")
  if (!fs.pathExistsSync(objectspath)) {
    fs.removeSync(objectspath)
  }
  var buf = fs.readFileSync(path.join(buildFolder,'/tpl/objects.go.tpl'));
  Handlebars.registerHelper('plugins', plugins);
  var template = Handlebars.compile(buf.toString());
  let gofile = template({"entities": entities})
  fs.writeFileSync(objectspath, gofile)
}

module.exports = {
  createEntity: createEntity,
  createManifest: createManifest,
  autogenFolder: autogenFolder
}
