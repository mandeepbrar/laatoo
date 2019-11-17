var path = require('path');
var sprintf = require('sprintf-js').sprintf
var fs = require('fs-extra')
var Handlebars = require('handlebars')
var {log, listDir} = require('./utils');
var {buildFolder, name, pluginFolder} = require('./buildconfig');
var {createGoModule} = require('./utils')

var hasSdk = false;

function collection(collection, name) {
  return collection? collection: name
}

function cacheable(cacheable) {
  return cacheable? cacheable: false
}

function multitenant(multitenant) {
  return multitenant? multitenant: false
}
function postsave(postsave) {
  return postsave? postsave: false
}
function presave(presave) {
  return presave? presave: false
}
function postload(postload) {
  return postload? postload: false
}
function sdkinclude(sdkinclude) {
  return sdkinclude? sdkinclude: false
}

function imports(imports) {
  var importsStr = ""
  if(imports) {
    imports.forEach(function(pkg) {
      if(pkg) {
        importsStr = importsStr + "\r\n\t\""+ pkg +"\" "
      }
    });  
  }
  return importsStr
}

function type(name) {
  return name? name: "entity"
}

function titleField(titleField) {
  return titleField? titleField: "Title"
}
function modulename() {
  return name
}

function fields(fields) {
  var fieldsStr = ""
  Object.keys(fields).forEach(function(fieldName) {
    let field = fields[fieldName]
    let jsonF = field.json?field.json:fieldName
    let bsonF = field.bson?field.bson:fieldName
    let datastoreF = field.datastore? field.datastore : fieldName
    let refentity = field.entity? field.entity: "";
    let fieldType = getFieldType(field)    
    fieldsStr = fieldsStr + sprintf("\r\n\t%s\t%s `json:\"%s\" bson:\"%s\" datastore:\"%s\"`", fieldName, fieldType, jsonF, bsonF, datastoreF)
  });
  return fieldsStr
}

function goFolder(name) {
  let autogenFolder = path.join(pluginFolder, "server", "go")
  if(!fs.existsSync(autogenFolder)) {
    fs.mkdirsSync(autogenFolder);
    createGoModule( name , autogenFolder)
  }
  return autogenFolder
}

function sdkFolder(name) {
  let sdkFolder = path.join(pluginFolder, "sdk", "go")
  if(!fs.existsSync(sdkFolder)) {
    fs.mkdirsSync(sdkFolder);
    //createGoModule( name +"/sdk" , sdkFolder)
  }
  return sdkFolder
}

function goFolder() {
  return path.join(pluginFolder, "server", "go")  
}

function createEntity(entityJson, filename) {
  let entityname = entityJson["name"];
  entityname = entityname? entityname: filename.substring(0, filename.length-5);
  let sdkInclude = entityJson["sdkinclude"];
  if(sdkInclude) {
    //copyEntityToSDK(filename);
    createEntityImpl(entityJson, entityname, sdkFolder(name));
    hasSdk = true
  } else {
    createEntityImpl(entityJson, entityname, goFolder(name));
  }
}

function addImport(entityJson, importName) {
  let imports = entityJson.imports
  if (imports == null) {
    let imports = []
    entityJson.imports = imports
  }
  imports.push(importName)
}

function createEntityImpl(entityJson, entityname, folder) {
  entityname = entityname +".go";
  fs.mkdirsSync(folder);
  let filepath = path.join(folder, "autogen_" + entityname)
  let tplpath = path.join(buildFolder, 'tpl/entitygocode.go.tpl');
  if(entityJson.collection == "<nocollection>") {
    tplpath = path.join(buildFolder, 'tpl/entitynodb.go.tpl');
  }
  var buf = fs.readFileSync(tplpath);
  Handlebars.registerHelper('cacheable', cacheable);
  Handlebars.registerHelper('postsave', postsave);
  Handlebars.registerHelper('presave', presave);
  Handlebars.registerHelper('modulename', modulename);
  Handlebars.registerHelper('sdkinclude', sdkinclude);  
  Handlebars.registerHelper('postload', postload);
  Handlebars.registerHelper('multitenant', multitenant);
  Handlebars.registerHelper('imports', imports);
  Handlebars.registerHelper('type', type);
  Handlebars.registerHelper('titleField', titleField);
  Handlebars.registerHelper('fields', fields);
  Handlebars.registerHelper('collection', collection);
  Handlebars.registerHelper('fieldFuncs', fieldFuncs);
  Handlebars.registerHelper('fieldReadAlls', fieldReadAlls);
  Handlebars.registerHelper('fieldWriteAlls', fieldWriteAlls);
  /*let requireBytes = bytesRequired(entityJson.fields) 
  if(requireBytes) {
    addImport(entityJson,"bytes")
  }*/

  var template = Handlebars.compile(buf.toString());
  let gofile = template(entityJson)
  if (fs.pathExistsSync(filepath)) {
    fs.removeSync(filepath)
  }
  fs.writeFileSync(filepath, gofile)
}


function createEntityInterface(entityJson, entityname, folder) {
  let destFile = path.join(folder, "autogen_I" + entityname + ".go")
  let tplpath = path.join(buildFolder, 'tpl/entityinterface.go.tpl');
  var buf = fs.readFileSync(tplpath);
  Handlebars.registerHelper('cacheable', cacheable);
  Handlebars.registerHelper('postsave', postsave);
  Handlebars.registerHelper('presave', presave);
  Handlebars.registerHelper('postload', postload);
  Handlebars.registerHelper('modulename', modulename);
  Handlebars.registerHelper('multitenant', multitenant);
  Handlebars.registerHelper('imports', imports);
  Handlebars.registerHelper('type', type);
  Handlebars.registerHelper('titleField', titleField);
  Handlebars.registerHelper('fields', fields);
  Handlebars.registerHelper('collection', collection);
  Handlebars.registerHelper('fieldFuncDefs', fieldFuncDefs);
  Handlebars.registerHelper('fieldReadAlls', fieldReadAlls);
  Handlebars.registerHelper('fieldWriteAlls', fieldWriteAlls);
  /*let requireBytes = bytesRequired(entityJson.fields) 
  if(requireBytes) {
    addImport(entityJson,"bytes")
  }*/
  
  var template = Handlebars.compile(buf.toString());
  let gofile = template(entityJson)
  if (fs.pathExistsSync(destFile)) {
    fs.removeSync(destFile)
  }
  fs.writeFileSync(destFile, gofile)
}


function genfiles(entities) {
  let str = ""
  Object.keys(entities).forEach(function(entity) {
      str = sprintf("%s autogen_%s.go ", str, entity)
  });
  return str
}


function plugins(entities) {
  let str = ""
  Object.keys(entities).forEach(function(entity) {
      let entityJson = entities[entity]
      let entityDesc = JSON.stringify(entityJson).replace(/\"/g,'\\"')
      let objectName = entityJson.sdkinclude? name + "." + entity: entity
      str = str + sprintf("core.PluginComponent{Name: \"%s\", Object: %s{}, Metadata: core.NewInfo(\"\",\"%s\", map[string]interface{}{\"descriptor\":\"%s\"})},", entity, objectName, entity, entityDesc)
  });
  return str
}

function createManifest(entities, name, pluginFolder) {
  let manifestpath = goFolder(name)
  let manifestFile = path.join(manifestpath, "manifest.go")
  if (!fs.pathExistsSync(manifestFile)) {
    fs.mkdirsSync(manifestpath)
    var buf = fs.readFileSync(path.join(buildFolder, '/tpl/manifest.go.tpl'));
    var template = Handlebars.compile(buf.toString());
    let gofile = template({})
    fs.writeFileSync(manifestFile, gofile)
  }

  let objectspath = path.join(goFolder(name), "autogen_objectsmanifest.go")
  if (!fs.pathExistsSync(objectspath)) {
    fs.removeSync(objectspath)
  }
  var buf = fs.readFileSync(path.join(buildFolder,'/tpl/objects.go.tpl'));
  Handlebars.registerHelper('plugins', plugins);
  Handlebars.registerHelper('genfiles', genfiles);
  Handlebars.registerHelper('goFolder', goFolder);
  var template = Handlebars.compile(buf.toString());
  let gofile = template({"entities": entities, "name": name, "hasSDK": hasSdk})
  fs.writeFileSync(objectspath, gofile)
  
}

/*
function bytesRequired(fields) {
  let requireBytes = false
  Object.values(fields).forEach(function(field) {
    let func = getFieldSerializationFunc(field)    
    if(func == "Object") {
      requireBytes = true
    }
  });
  return requireBytes
}*/


function fieldReadAlls(fields) {
  var fieldsStr = ""
  Object.keys(fields).forEach(function(fieldName) {
    let field = fields[fieldName]
    let func = getFieldSerializationFunc(field)    
    let fieldType = getFieldType(field)    
    let entity = field.entity? field.entity: "";
    let readAll = `
    if err = rdr.Read%s(c, cdc, "%s", &ent.%s); err != nil {
      return err
    }
    `  
    if(func != "Object") {
      fieldsStr = fieldsStr + sprintf(readAll, func, fieldName, fieldName)
    } else {
      readAll = `
      {
        ent.%s := &%s{}
        if err = rdr.ReadObject(c, cdc, "%s", &ent.%s); err != nil {
          return err
        }
      }
      `  
      fieldsStr = fieldsStr + sprintf(readAll, fieldName, entity, fieldName, fieldName)
    }
  });

  return fieldsStr
}

function fieldWriteAlls(fields) {
  var fieldsStr = ""
  Object.keys(fields).forEach(function(fieldName) {
    let field = fields[fieldName]
    let func = getFieldSerializationFunc(field)    
    let writeAll = `
    if err = wtr.Write%s(c, cdc, "%s", &ent.%s); err != nil {
      return err
    }
    `  
    fieldsStr = fieldsStr + sprintf(writeAll, func, fieldName, fieldName)
  });
  return fieldsStr
}


function getFieldSerializationFunc(field) {
  let fieldType = field.type
  if(field.list) {
    return "Array"
  }
  switch (fieldType) {
    case "string":
      fieldType = "String"
      break;
    case "bool":
      fieldType = "Bool"
      break;
    case "time":
      fieldType = "Time"
      break;
    case "int":
      fieldType = "Int"
      break;
    case "int64":
      fieldType = "Int64"
      break;
    case "float32":
      fieldType = "Float32"
      break;
    case "float64":
      fieldType = "Float64"
      break;
    case "any":
      fieldType = "Object"
      break;
    case "storable":
      fieldType = "Object"
      break;
    case "storableref":
      fieldType = "Object"
      break;
    break;
    case "entity":
      fieldType = "Object"
    break;
    case "subentity":
      fieldType = "Object"
      break;
    case "stringmap":
      fieldType = "Map"
      break;
    case "stringsmap":
      fieldType = "Map"
    break;
  }
  return fieldType  
}

function fieldFuncDefs(fields) {
  var fieldsStr = ""
  Object.keys(fields).forEach(function(fieldName) {
    let field = fields[fieldName]
    let refentity = field.entity? field.entity: "";
    let fieldType = getFieldType(field)    
    fieldsStr = fieldsStr + sprintf("\r\n\tGet%s()%s", fieldName, fieldType)
    fieldsStr = fieldsStr + sprintf("\r\n\tSet%s(%s)", fieldName, fieldType)
  });
  return fieldsStr
}

function fieldFuncs(fields, name) {
  var fieldsStr = ""
  Object.keys(fields).forEach(function(fieldName) {
    let field = fields[fieldName]
    let fieldType = getFieldType(field)    
    fieldsStr = fieldsStr + sprintf("\r\nfunc (ent *%s) Get%s()%s {\r\n\treturn ent.%s\r\n}", name, fieldName, fieldType, fieldName)
    fieldsStr = fieldsStr + sprintf("\r\nfunc (ent *%s) Set%s(val %s) {\r\n\tent.%s=val\r\n}", name, fieldName, fieldType, fieldName)
  });
  return fieldsStr
}

function getFieldType(field) {
  let fieldType = field.type
  switch (fieldType) {
    case "any":
      fieldType = "interface{}"
      break;
    case "storable":
      fieldType = "data.Storable"
      break;
    case "storableref":
      fieldType = "data.StorableRef"
      break;
    break;
    case "entity":
      fieldType = "*"+field.entity
    break;
    case "subentity":
      fieldType = field.entity
      break;
    case "stringmap":
      if(field.mappedElement) {
        fieldType = "map[string]"+field.mappedElement
      } else {
        fieldType = "map[string]interface{}"
      }
      break;
    case "stringsmap":
      fieldType = "map[string]string"
    break;
  }
  if(field.list) {
    fieldType = "[]"+fieldType
  }
  return fieldType  
}

module.exports = {
  createEntity: createEntity,
  createManifest: createManifest
}
