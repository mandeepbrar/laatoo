package main

import (
	"bytes"
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"path"
	"strings"
	"text/template"
)

const (
	ENTITY_MODULE   = "EntityModule"
	ENTITY_OBJECT   = "object"
	ENTITY_INSTANCE = "instance"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: ENTITY_MODULE, Object: EntityModule{}}}
}

type EntityModule struct {
	core.Module
	object       string
	entityConf   config.Config
	instance     string
	templatesDir string
}

/*
func (adapter *DataAdapterModule) Describe(ctx core.ServerContext) {
	adapter.AddStringConfiguration(ctx, CONF_DATASERVICE_FACTORY)
	adapter.AddStringConfiguration(ctx, data.CONF_DATA_OBJECT)
	adapter.AddStringConfigurations(ctx, []string{DATA_ADAPTER_INSTANCE, MIDDLEWARE, CONF_PARENT_CHANNEL}, []string{"", "", "root"})
}*/
func (entity *EntityModule) MetaInfo(ctx core.ServerContext) map[string]interface{} {
	return map[string]interface{}{}
}

func (entity *EntityModule) Initialize(ctx core.ServerContext, conf config.Config) error {
	ctx = ctx.SubContext("Initializing entity module")
	entity.object, _ = entity.GetStringConfiguration(ctx, ENTITY_OBJECT)
	entity.instance, _ = entity.GetStringConfiguration(ctx, ENTITY_INSTANCE)
	md, err := ctx.GetObjectMetadata(entity.object)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if md == nil {
		return errors.ThrowError(ctx, errors.CORE_ERROR_RES_NOT_FOUND, "Object metadata not found", entity.object)
	}

	baseDir, _ := ctx.GetString(config.MODULEDIR)
	entity.templatesDir = path.Join(baseDir, "files", "templates")

	desc := md.GetProperty("descriptor")
	if desc != nil {
		str, ok := desc.(string)
		if ok {
			conf, err := ctx.ReadConfigData([]byte(str), nil)
			if err != nil {
				return errors.WrapError(ctx, err)
			}
			entity.entityConf = conf
		}
	}

	if entity.instance == "" {
		entity.instance = entity.object
	}
	return nil
}

func (entity *EntityModule) createName(ctx core.ServerContext, svc string) string {
	if entity.instance != "" {
		return fmt.Sprintf("dataadapter.%s.%s", svc, entity.instance)
	} else {
		return fmt.Sprintf("dataadapter.%s.%s", svc, entity.object)
	}
}

func (entity *EntityModule) LoadingComplete(ctx core.ServerContext) map[string]config.Config {
	return nil
}
func (entity *EntityModule) UILoad(ctx core.ServerContext) map[string]config.Config {
	reg := ctx.CreateConfig()
	forms := entity.createForms(ctx)
	reg.Set(ctx, "Forms", forms)
	blocks, err := entity.createBlocks(ctx)
	if err != nil {
		log.Error(ctx, "Error createing entity registry", "Error", err)
	}
	reg.Set(ctx, "Blocks", blocks)
	return map[string]config.Config{"registry": reg}
}

func (entity *EntityModule) createForms(ctx core.ServerContext) config.Config {
	forms := ctx.CreateConfig()

	newEntityForm := ctx.CreateConfig()
	var entityFormInfo config.Config
	formInfo, ok := entity.entityConf.GetSubConfig(ctx, "form")
	if ok {
		entityFormInfo = formInfo.Clone()
	} else {
		entityFormInfo = ctx.CreateConfig()
	}
	entityFormInfo.Set(ctx, "entity", entity.object)
	entityFormInfo.Set(ctx, "className", fmt.Sprint(" entityform ", strings.ToLower(entity.instance+"_form")))
	entityFormInfo.Set(ctx, "successRedirectPage", fmt.Sprint("list_", strings.ToLower(entity.instance)))
	formNewArgs, ok := ctx.Get("form_new_args")
	if ok {
		log.Error(ctx, " Setting pre assigned params for new form", "form_new_args", formNewArgs)
		entityFormInfo.Set(ctx, "preAssigned", formNewArgs)
	}
	newEntityForm.Set(ctx, "info", entityFormInfo)

	entityFormFields := ctx.CreateConfig()

	fields, ok := entity.entityConf.GetSubConfig(ctx, "fields")
	if ok {
		fieldNames := fields.AllConfigurations(ctx)
		for _, field := range fieldNames {
			//fieldToBeAdded := ctx.CreateConfig()
			fieldConf, _ := fields.GetSubConfig(ctx, field)
			fieldToBeAdded := fieldConf.Clone()
			widgetConf, ok := fieldToBeAdded.GetSubConfig(ctx, "widget")
			if !ok {
				widgetConf = ctx.CreateConfig()
				fieldToBeAdded.Set(ctx, "widget", widgetConf)
			}
			fieldProps, ok := widgetConf.GetSubConfig(ctx, "props")
			if ok {
				fieldProps = fieldProps.Clone()
			} else {
				fieldProps = ctx.CreateConfig()
			}
			className, _ := fieldProps.GetString(ctx, "className")
			className = fmt.Sprintf(" %s entityformfield ", className)
			fieldProps.Set(ctx, "className", className)
			widgetConf.Set(ctx, "props", fieldProps)
			entityFormFields.Set(ctx, field, fieldToBeAdded)
		}
	}

	newEntityForm.Set(ctx, "fields", entityFormFields)

	forms.Set(ctx, "new_form_"+strings.ToLower(entity.instance), newEntityForm)

	updateEntityForm := ctx.CreateConfig()
	updateFormInfo := entityFormInfo.Clone()
	//updateFormInfo.Set(ctx, "successRedirect", "/list_"+strings.ToLower(entity.instance))
	updateEntityForm.Set(ctx, "info", updateFormInfo)
	updateEntityForm.Set(ctx, "fields", entityFormFields)

	forms.Set(ctx, "update_form_"+strings.ToLower(entity.instance), updateEntityForm)

	return forms
}

func (entity *EntityModule) createBlocks(ctx core.ServerContext) (config.Config, error) {
	blocks := ctx.CreateConfig()

	type TemplateData struct {
		EntityName string
		LabelField string
	}

	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"lower": strings.ToLower,
	}

	viewtableHeaderTemp, err := template.New("viewtableheader.tpl").Delims("<<", ">>").Funcs(funcMap).ParseFiles(path.Join(entity.templatesDir, "viewtableheader.tpl"))
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	tableHeaderStr := new(bytes.Buffer)
	data := TemplateData{entity.object, "Name"}

	err = viewtableHeaderTemp.Execute(tableHeaderStr, data)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	tableHeader, err := ctx.ReadConfigData(tableHeaderStr.Bytes(), nil)
	if err == nil {
		blocks.Set(ctx, strings.ToLower(entity.object)+"_viewheader", tableHeader)
	} else {
		log.Error(ctx, "Error writing entity block", "Err", err)
		return blocks, nil
	}

	viewtableRowTemp, err := template.New("viewtablerow.tpl").Delims("<<", ">>").Funcs(funcMap).ParseFiles(path.Join(entity.templatesDir, "viewtablerow.tpl"))
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	tableRowStr := new(bytes.Buffer)

	err = viewtableRowTemp.Execute(tableRowStr, data)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}

	tableRow, err := ctx.ReadConfigData(tableRowStr.Bytes(), nil)
	if err == nil {
		blocks.Set(ctx, entity.object+"_listtablerow", tableRow)
	} else {
		log.Error(ctx, "Error writing entity block", "Err", err)
		return blocks, nil
	}

	defaultEntityTemp, err := template.New("defaultentity.tpl").Delims("<<", ">>").Funcs(funcMap).ParseFiles(path.Join(entity.templatesDir, "defaultentity.tpl"))
	defaultEntityStr := new(bytes.Buffer)

	err = defaultEntityTemp.Execute(defaultEntityStr, data)
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	log.Error(ctx, " defaultEntityTemp ", "file", string(defaultEntityStr.Bytes()))

	defEntity, err := ctx.ReadConfigData(defaultEntityStr.Bytes(), nil)
	if err == nil {
		blocks.Set(ctx, entity.object+"_default", defEntity)
	} else {
		log.Error(ctx, "Error writing entity block", "Err", err)
		return blocks, nil
	}

	/*defaultBlk := ctx.CreateConfig()
	blkDiv := ctx.CreateConfig()

	blkDiv.Set(ctx, "body", `{{jsreplace "Window.displayDefaultEntity(ctx, desc, uikit)"`)
	blkDiv.Set(ctx, "className", "entity default "+entity.object)
	defaultBlk.Set(ctx, "Block", blkDiv)
	blocks.Set(ctx, entity.object+"_default", defaultBlk)
	*/
	return blocks, nil
}
