package core

import (
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/ctx"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"laatoo/server/constants"
	"os"
	"path"
	"path/filepath"
	"plugin"
)

type objectLoader struct {
	objectsFactoryRegister map[string]core.ObjectFactory
	objModMap              utils.StringMap
	name                   string
	parentElem             core.ServerElement
	provider               *metadataProvider
}

func (objLoader *objectLoader) Initialize(ctx core.ServerContext, conf config.Config) error {

	objectsBaseFolder, ok := conf.GetString(ctx, constants.CONF_OBJECTS_BASE_DIR)
	if ok {
		ctx.Set(constants.CONF_OBJECTS_BASE_DIR, objectsBaseFolder)
		log.Info(ctx, "Setting base directory for objects", "Directory", objectsBaseFolder)
	}

	err := objLoader.loadObjects(ctx, conf)
	if err != nil {
		return err
	}
	/*
		if objLoader.parentElem == nil {
			for objectName, objFactory := range objectsFactoryRegister {
				objLoader.registerObjectFactory(ctx, objectName, objFactory)
			}

			for methodName, method := range invokableMethodsRegister {
				objLoader.registerInvokableMethod(ctx, methodName, method)
			}
		}
	*/
	/*
		objectNames, ok := conf.GetStringArray(config.CONF_OBJECTLDR_OBJECTS)
		if ok {
			objectNames = append(objectNames, "string")
			for _, objectName := range objectNames {
				fac, ok := objectsFactoryRegister[objectName]
				if ok {
					objLoader.registerObjectFactory(ctx, objectName, fac)
				} else {
					return errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)
				}
			}
		}
		methodNames, ok := conf.GetStringArray(config.CONF_OBJECTLDR_METHODS)
		if ok {
			for _, methodName := range methodNames {
				method, ok := invokableMethodsRegister[methodName]
				if ok {
					objLoader.registerInvokableMethod(ctx, methodName, method)
				} else {
					return errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Method Name", methodName)
				}
			}
		}*/
	return nil
}

func (objLoader *objectLoader) Start(ctx core.ServerContext) error {
	return nil
}

func (objLoader *objectLoader) loadObjectsFolder(ctx core.ServerContext, folder string) error {
	log.Trace(ctx, "Loading objects", "objects folder", folder)
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.RethrowError(ctx, errors.CORE_ERROR_PLUGIN_NOT_LOADED, err, "Path", path)
		}
		if !info.IsDir() {
			log.Trace(ctx, "Loading plugin", "path", path)
			p, err := plugin.Open(path)
			if err != nil {
				return errors.RethrowError(ctx, errors.CORE_ERROR_PLUGIN_NOT_LOADED, err, "Path", path)
			}
			sym, err := p.Lookup("Manifest")
			if err != nil {
				return errors.RethrowError(ctx, errors.CORE_ERROR_PLUGIN_NOT_LOADED, err, "Path", path)
			}
			manifest, ok := sym.(func(core.MetaDataProvider) []core.PluginComponent)
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_PLUGIN_NOT_LOADED, "Path", path, "Err", "Manifest incorrect")
			}
			if manifest == nil {
				return errors.ThrowError(ctx, errors.CORE_ERROR_PLUGIN_NOT_LOADED, "Path", path, "Err", "Manifest incorrect")
			}
			components := manifest(objLoader.provider)
			if components != nil {
				for _, comp := range components {
					if comp.Name != "" {
						if comp.ObjectFactory != nil {
							objLoader.registerObjectFactory(ctx, comp.Name, comp.ObjectFactory)
						} else if comp.ObjectCreator != nil {
							objLoader.registerObject(ctx, comp.Name, comp.ObjectCreator, comp.ObjectCollectionCreator, comp.Metadata)
						} else if comp.Object != nil {
							objLoader.register(ctx, comp.Name, comp.Object, comp.Metadata)
						} else {
							log.Info(ctx, "No component registered", "Component", comp.Name, "Path", path)
						}
					} else {
						log.Info(ctx, "No component registered for empty name", "Path", path)
					}
				}
				log.Info(ctx, "Loaded plugin", "Path", path)
				log.Trace(ctx, "Objects in the plugin", "Components", components)
			} else {
				log.Info(ctx, "No objects in the plugin", "Path", path)
				return errors.ThrowError(ctx, errors.CORE_ERROR_PLUGIN_NOT_LOADED, "Path", path, "Err", "No objects found")
			}
		}
		return nil
	})
	return err
}

func (objLoader *objectLoader) loadObjectsFolderIfExists(ctx core.ServerContext, folder string) error {
	exists, _, _ := utils.FileExists(folder)
	if exists {
		log.Trace(ctx, "Loading objects folder", "Folder", folder)
		if err := objLoader.loadObjectsFolder(ctx, folder); err != nil {
			return errors.WrapError(ctx, err)
		}
	} else {
		log.Trace(ctx, "Folder does not exist", "Folder", folder)
	}
	return nil
}

func (objLoader *objectLoader) loadObjects(ctx core.ServerContext, conf config.Config) error {
	objsFolder, ok := conf.GetString(ctx, constants.CONF_OBJECTLDR_OBJECTS)
	if ok {
		if err := objLoader.loadObjectsFolderIfExists(ctx, objsFolder); err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	baseDir, _ := ctx.GetString(config.BASEDIR)
	baseFolder := path.Join(baseDir, constants.CONF_OBJECTLDR_OBJECTS)
	if err := objLoader.loadObjectsFolderIfExists(ctx, baseFolder); err != nil {
		return err
	}

	objectsbaseDir, found := ctx.GetString(constants.CONF_OBJECTS_BASE_DIR)
	if found {
		relDir, _ := ctx.GetString(constants.RELATIVE_DIR)
		baseFolder = path.Join(objectsbaseDir, relDir)
		if err := objLoader.loadObjectsFolderIfExists(ctx, baseFolder); err != nil {
			return err
		}
	}
	return nil
}

func (objLoader *objectLoader) setObjectInfo(ctx ctx.Context, objectName string, metadata core.Info) {
	factory, ok := objLoader.objectsFactoryRegister[objectName]
	if ok {
		factory.(*objectFactory).metadata = metadata
	}
}

func (objLoader *objectLoader) register(ctx ctx.Context, objectName string, object interface{}, metadata core.Info) {
	objLoader.registerObjectFactory(ctx, objectName, newObjectType(object, metadata))
}
func (objLoader *objectLoader) registerObject(ctx ctx.Context, objectName string, objectCreator core.ObjectCreator, objectCollectionCreator core.ObjectCollectionCreator, metadata core.Info) {
	objLoader.registerObjectFactory(ctx, objectName, newObjectFactory(objectCreator, objectCollectionCreator, metadata))
}
func (objLoader *objectLoader) registerObjectFactory(ctx ctx.Context, objectName string, factory core.ObjectFactory) {
	_, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		log.Info(ctx, "Registering object factory ", "Object Name", objectName)
		objLoader.objectsFactoryRegister[objectName] = factory
		mod, modok := ctx.GetString(constants.CONF_MODULE)
		if modok {
			objLoader.objModMap[objectName] = mod
		}
	}
}

//returns a collection of the object type
func (objLoader *objectLoader) createCollection(ctx ctx.Context, objectName string, length int) (interface{}, error) {
	//get the factory object from the register
	factory, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)

	}
	return factory.CreateObjectCollection(length), nil
}

//Provides an object with a given name
func (objLoader *objectLoader) createObject(ctx ctx.Context, objectName string) (interface{}, error) {
	//get the factory object from the register
	factory, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		log.Trace(ctx, "Objects in the register", "Map", objLoader.objectsFactoryRegister)
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)
	}
	return factory.CreateObject(), nil
}

func (objLoader *objectLoader) getObjectCreator(ctx ctx.Context, objectName string) (core.ObjectCreator, error) {
	//get the factory object from the register
	factory, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)

	}
	return factory.CreateObject, nil
}

func (objLoader *objectLoader) getObjectCollectionCreator(ctx ctx.Context, objectName string) (core.ObjectCollectionCreator, error) {
	//get the factory object from the register
	factory, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)

	}
	return factory.CreateObjectCollection, nil
}

func (objLoader *objectLoader) getMetaData(ctx ctx.Context, objectName string) (core.Info, error) {
	if objectName == "" {
		return nil, nil
	}
	//get the factory object from the register
	factory, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		log.Trace(ctx, "Object not found by loader", "Object Name", objectName)
		return nil, nil
	}
	return factory.Info(), nil
}

func (objLoader *objectLoader) unloadModuleObjects(ctx ctx.Context, modName string) error {
	for obj, mod := range objLoader.objModMap {
		if mod == modName {
			log.Info(ctx, "Unloaded Object", "Name:", obj, "Module", mod)
			delete(objLoader.objModMap, obj)
			delete(objLoader.objectsFactoryRegister, obj)
		}
	}
	return nil
}
