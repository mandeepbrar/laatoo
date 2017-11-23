package main

import (
	"fmt"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
	"strings"
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
	object     string
	entityConf config.Config
	instance   string
}

/*
func (adapter *DataAdapterModule) Describe(ctx core.ServerContext) {
	adapter.AddStringConfiguration(ctx, CONF_DATASERVICE_FACTORY)
	adapter.AddStringConfiguration(ctx, data.CONF_DATA_OBJECT)
	adapter.AddStringConfigurations(ctx, []string{DATA_ADAPTER_INSTANCE, MIDDLEWARE, CONF_PARENT_CHANNEL}, []string{"", "", "root"})
}*/

func (entity *EntityModule) Initialize(ctx core.ServerContext, conf config.Config) error {
	ctx = ctx.SubContext("Initializing entity module")
	entity.object, _ = entity.GetStringConfiguration(ctx, ENTITY_OBJECT)
	entity.instance, _ = entity.GetStringConfiguration(ctx, ENTITY_INSTANCE)

	md, err := ctx.GetObjectMetadata(entity.object)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
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

func (entity *EntityModule) GetRegistry(ctx core.ServerContext) config.Config {
	reg := ctx.CreateConfig()
	forms := entity.createForms(ctx)
	reg.Set(ctx, "Forms", forms)
	blocks := entity.createBlocks(ctx)
	reg.Set(ctx, "Blocks", blocks)
	return reg
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
	newEntityForm.Set(ctx, "info", entityFormInfo)
	entityFormFields := ctx.CreateConfig()

	fields, ok := entity.entityConf.GetSubConfig(ctx, "fields")
	if ok {
		fieldNames := fields.AllConfigurations(ctx)
		for _, field := range fieldNames {
			fieldToBeAdded := ctx.CreateConfig()

			fieldConf, _ := fields.GetSubConfig(ctx, field)
			fType, ok := fieldConf.GetString(ctx, "type")
			if ok {
				fieldToBeAdded.Set(ctx, "type", fType)
			}
			fLabel, ok := fieldConf.GetString(ctx, "label")
			if ok {
				fieldToBeAdded.Set(ctx, "label", fLabel)
			}
			fieldToBeAdded.Set(ctx, "className", " entityformfield "+field)
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

func (entity *EntityModule) createBlocks(ctx core.ServerContext) config.Config {
	blocks := ctx.CreateConfig()

	tableHeaderStr := `{
			div: {
				className: "%s_list_header tableheading ",
				children: [
					{
						div:	{
							className: "tablecell",
							body: "Name"
						}
					},
					{
						div:	{
							className: "tablecell",
							body: "Last Updated"
						}
					},
					{
						div:	{
							className: "tablecell",
							body: ""
						}
					}
				]
			}
		}`

	tableHeader, err := ctx.ReadConfigData([]byte(fmt.Sprintf(tableHeaderStr, entity.object)), nil)
	if err == nil {
		blocks.Set(ctx, strings.ToLower(entity.object)+"_viewheader", tableHeader)
	} else {
		log.Error(ctx, "Error writing entity block", "Err", err)
		return blocks
	}

	tableRowStr := `{
			config: {
				skip: false
			},
			div: {
				className: "%s_default tablerow javascript%s ",
			  children: [
					{
						div: {
							className: "tablecell field",
							body: "javascript%s"
						}
					},
					{
						div: {
							className: "tablecell field",
							body: "javascript%s"
						}
					},
					{
						div: {
							className: "tablecell field",
							children: [
								{
									Action: {
											module: "reactwebcommon",
											name: "update_page_%s",
											params: {
												entityId: "javascript%s"
											},
											body: "View %s"
									}
								}
							]
						}
					}
				]
			}
		}	`
	labelField := "Name"
	tableRowStr = fmt.Sprintf(tableRowStr, entity.object, "#@#ctx.className#@#", fmt.Sprintf("#@#ctx.data.%s#@#", labelField), "#@#ctx.data.UpdatedAt#@#", strings.ToLower(entity.object), "#@#ctx.data.Id#@#", entity.object)
	tableRow, err := ctx.ReadConfigData([]byte(tableRowStr), nil)
	if err == nil {
		blocks.Set(ctx, entity.object+"_listtablerow", tableRow)
	} else {
		log.Error(ctx, "Error writing entity block", "Err", err)
		return blocks
	}

	/*fields, ok := entity.entityConf.GetSubConfig(ctx, "fields")
	if ok {
		fieldNames := fields.AllConfigurations(ctx)
		for _, field := range fieldNames {
			fieldToBeAdded := ctx.CreateConfig()

			fieldConf, _ := fields.GetSubConfig(ctx, field)
			fType, ok := fieldConf.GetString(ctx, "type")
			if ok {
				fieldToBeAdded.Set(ctx, "type", fType)
			}
			fLabel, ok := fieldConf.GetString(ctx, "label")
			if ok {
				fieldToBeAdded.Set(ctx, "label", fLabel)
			}
			fieldToBeAdded.Set(ctx, "className", " entityfield "+field)
			newEntityFormFields.Set(ctx, field, fieldToBeAdded)

		}
	}*/

	return blocks
}
