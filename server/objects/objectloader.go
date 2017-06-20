package objects

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"laatoo/sdk/utils"
	"os"
	"path"
	"path/filepath"
	"plugin"
)

type objectLoader struct {
	objectsFactoryRegister   map[string]core.ObjectFactory
	invokableMethodsRegister map[string]core.ServiceFunc
	name                     string
	parentElem               core.ServerElement
}

func (objLoader *objectLoader) Initialize(ctx core.ServerContext, conf config.Config) error {

	objectsBaseFolder, ok := conf.GetString(config.CONF_OBJECTS_BASE_DIR)
	if ok {
		ctx.Set(config.CONF_OBJECTS_BASE_DIR, objectsBaseFolder)
		log.Logger.Info(ctx, "Setting base directory for objects", "Directory", objectsBaseFolder)
	}

	err := objLoader.loadPlugins(ctx, conf)
	if err != nil {
		return err
	}

	for objectName, objFactory := range objectsFactoryRegister {
		objLoader.registerObjectFactory(ctx, objectName, objFactory)
	}

	for methodName, method := range invokableMethodsRegister {
		objLoader.registerInvokableMethod(ctx, methodName, method)
	}

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

func (objLoader *objectLoader) loadPluginsFolder(ctx core.ServerContext, folder string) error {
	log.Logger.Trace(ctx, "Loading plugins", "plugins folder", folder)
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.RethrowError(ctx, errors.CORE_ERROR_PLUGIN_NOT_LOADED, err, "Path", path)
		}
		if !info.IsDir() {
			log.Logger.Trace(ctx, "Loading plugin", "path", path)
			p, err := plugin.Open(path)
			if err != nil {
				return errors.RethrowError(ctx, errors.CORE_ERROR_PLUGIN_NOT_LOADED, err, "Path", path)
			}
			sym, err := p.Lookup("Manifest")
			if err != nil {
				return errors.RethrowError(ctx, errors.CORE_ERROR_PLUGIN_NOT_LOADED, err, "Path", path)
			}
			manifest, ok := sym.(func() []core.PluginComponent)
			if !ok {
				return errors.ThrowError(ctx, errors.CORE_ERROR_PLUGIN_NOT_LOADED, "Path", path, "Err", "Manifest incorrect")
			}
			if manifest == nil {
				return errors.ThrowError(ctx, errors.CORE_ERROR_PLUGIN_NOT_LOADED, "Path", path, "Err", "Manifest incorrect")
			}
			components := manifest()
			if components != nil {
				for _, comp := range components {
					if comp.Name != "" {
						if comp.ServiceFunc != nil {
							RegisterInvokableMethod(comp.Name, comp.ServiceFunc)
						} else if comp.ObjectFactory != nil {
							RegisterObjectFactory(comp.Name, comp.ObjectFactory)
						} else if comp.ObjectCreator != nil {
							RegisterObject(comp.Name, comp.ObjectCreator, comp.ObjectCollectionCreator)
						} else if comp.Object != nil {
							Register(comp.Name, comp.Object)
						} else {
							log.Logger.Info(ctx, "No component registered", "Component", comp.Name, "Path", path)
						}
					} else {
						log.Logger.Info(ctx, "No component registered for empty name", "Path", path)
					}
				}
				log.Logger.Info(ctx, "Loaded plugin", "Path", path)
				log.Logger.Trace(ctx, "Objects in the plugin", "Components", components)
			} else {
				log.Logger.Info(ctx, "No objects in the plugin", "Path", path)
				return errors.ThrowError(ctx, errors.CORE_ERROR_PLUGIN_NOT_LOADED, "Path", path, "Err", "No objects found")
			}
		}
		return nil
	})
	return err
}

func (objLoader *objectLoader) loadPluginsFolderIfExists(ctx core.ServerContext, folder string) error {
	exists, _, _ := utils.FileExists(folder)
	if exists {
		if err := objLoader.loadPluginsFolder(ctx, folder); err != nil {
			return errors.WrapError(ctx, err)
		}
	} else {
		log.Logger.Trace(ctx, "Folder does not exist", "Folder", folder)
	}
	return nil
}

func (objLoader *objectLoader) loadPlugins(ctx core.ServerContext, conf config.Config) error {
	pluginsFolder, ok := conf.GetString(config.CONF_OBJECTLDR_OBJECTS)
	if ok {
		if err := objLoader.loadPluginsFolder(ctx, pluginsFolder); err != nil {
			return errors.WrapError(ctx, err)
		}
	}
	baseDir, _ := ctx.GetString(config.CONF_BASE_DIR)
	baseFolder := path.Join(baseDir, config.CONF_OBJECTLDR_OBJECTS)
	if err := objLoader.loadPluginsFolderIfExists(ctx, baseFolder); err != nil {
		return err
	}

	objectsbaseDir, found := ctx.GetString(config.CONF_OBJECTS_BASE_DIR)
	if found {
		svrElementType := ctx.GetElementType()
		switch svrElementType {
		case core.ServerElementServer:
			if err := objLoader.loadPluginsFolderIfExists(ctx, path.Join(objectsbaseDir, config.CONF_APP_SERVER)); err != nil {
				return err
			}
			break
		case core.ServerElementEnvironment:
			if err := objLoader.loadPluginsFolderIfExists(ctx, path.Join(objectsbaseDir, config.CONF_ENVIRONMENTS, objLoader.name)); err != nil {
				return err
			}
			break
		case core.ServerElementApplication:
			if err := objLoader.loadPluginsFolderIfExists(ctx, path.Join(objectsbaseDir, config.CONF_APPLICATIONS, objLoader.name)); err != nil {
				return err
			}
			break
		}
	}
	return nil
}

func (objLoader *objectLoader) register(ctx core.Context, objectName string, object interface{}) {
	objLoader.registerObjectFactory(ctx, objectName, NewObjectType(object))
}
func (objLoader *objectLoader) registerObject(ctx core.Context, objectName string, objectCreator core.ObjectCreator, objectCollectionCreator core.ObjectCollectionCreator) {
	objLoader.registerObjectFactory(ctx, objectName, NewObjectFactory(objectCreator, objectCollectionCreator))
}
func (objLoader *objectLoader) registerObjectFactory(ctx core.Context, objectName string, factory core.ObjectFactory) {
	_, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		log.Logger.Info(ctx, "Registering object factory ", "Object Name", objectName)
		objLoader.objectsFactoryRegister[objectName] = factory
	}
}
func (objLoader *objectLoader) registerInvokableMethod(ctx core.Context, methodName string, method core.ServiceFunc) {
	_, ok := objLoader.invokableMethodsRegister[methodName]
	if !ok {
		log.Logger.Debug(ctx, "Registering method ", "Method Name", methodName)
		objLoader.invokableMethodsRegister[methodName] = method
	}
}

//returns a collection of the object type
func (objLoader *objectLoader) createCollection(ctx core.Context, objectName string, length int) (interface{}, error) {
	//get the factory object from the register
	factory, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)

	}
	return factory.CreateObjectCollection(length), nil
}

//Provides an object with a given name
func (objLoader *objectLoader) createObject(ctx core.Context, objectName string) (interface{}, error) {
	//get the factory object from the register
	factory, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)
	}
	return factory.CreateObject(), nil
}

func (objLoader *objectLoader) getObjectCreator(ctx core.Context, objectName string) (core.ObjectCreator, error) {
	//get the factory object from the register
	factory, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)

	}
	return factory.CreateObject, nil
}

func (objLoader *objectLoader) getObjectCollectionCreator(ctx core.Context, objectName string) (core.ObjectCollectionCreator, error) {
	//get the factory object from the register
	factory, ok := objLoader.objectsFactoryRegister[objectName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Object Name", objectName)

	}
	return factory.CreateObjectCollection, nil
}

func (objLoader *objectLoader) getMethod(ctx core.Context, methodName string) (core.ServiceFunc, error) {
	method, ok := objLoader.invokableMethodsRegister[methodName]
	if !ok {
		return nil, errors.ThrowError(ctx, errors.CORE_ERROR_PROVIDER_NOT_FOUND, "Method Name", methodName)

	}
	return method, nil
}
