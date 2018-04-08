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
			//fieldToBeAdded := ctx.CreateConfig()
			fieldConf, _ := fields.GetSubConfig(ctx, field)
			fieldToBeAdded := fieldConf.Clone()

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
				className: "%s_tablerow tablerow javascript%s ",
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
	/*defaultStr := `{
		config: {
			skip: false
		},
		div: {
			className: "%s_default javascript%s ",
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
	}	`*/
	defaultBlk := ctx.CreateConfig()
	blkDiv := ctx.CreateConfig()
	/*fields := make([]config.Config, 0)
	fieldsConf, ok := entity.entityConf.GetSubConfig(ctx, "fields")
	if ok {
		fieldNames := fieldsConf.AllConfigurations(ctx)
		for _, field := range fieldNames {
			//fieldConf, _ := fieldsConf.GetSubConfig(ctx, field)
			fieldConf := ctx.CreateConfig()

			fieldDiv := ctx.CreateConfig()

			fieldElems := make([]config.Config, 0)

			fieldNameDiv := ctx.CreateConfig()
			fieldNameElems := ctx.CreateConfig()
			fieldNameElems.Set(ctx, "className", "name " + field)
			fieldNameElems.Set(ctx, "body", field)
			fieldNameDiv.Set(ctx, "div", fieldNameElems)
			fieldElems = append(fieldElems, fieldNameDiv)

			fieldValDiv := ctx.CreateConfig()
			fieldValElems := ctx.CreateConfig()
			fieldValElems.Set(ctx, "className", "value " + field)
			fieldValElems.Set(ctx, "body", fmt.Sprintf("javascript#@#ctx.data.%s#@#", field))
			fieldValDiv.Set(ctx, "div", fieldValElems)
			fieldElems = append(fieldElems, fieldValDiv)

			fieldDiv.Set(ctx, "children", fieldElems)
			fieldDiv.Set(ctx, "className", "field " + field)

			fieldConf.Set(ctx, "div", fieldDiv)
			//fldChildren := ctx.CreateConfig()
			fields = append(fields, fieldConf)
			//fieldToBeAdded := ctx.CreateConfig()
			//fieldConf, _ := fields.GetSubConfig(ctx, field)
		}
	}
	log.Error(ctx, "fields conf ", "entity", entity.object, "fields", fields)*/
	//blkDiv.Set(ctx, "children", fields)
	blkDiv.Set(ctx, "body", "javascript###Window.displayDefaultEntity(ctx, desc, uikit)###")
	blkDiv.Set(ctx, "className", "entity default "+entity.object)
	defaultBlk.Set(ctx, "div", blkDiv)
	blocks.Set(ctx, entity.object+"_default", defaultBlk)

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

	/*

	  layoutFields = (fldToDisp, flds, className) => {
	    let fieldsArr = new Array()
	    let comp = this
	    fldToDisp.forEach(function(k) {
	      let fd = flds[k]
	      let cl = className? className + " m10": "m10"
	      fieldsArr.push(  <Field key={fd.name} name={fd.name} formValue={comp.state.formValue} {...comp.parentFormProps} time={comp.state.time} className={cl}/>      )
	    })
	    return fieldsArr
	  }

	  fields = () => {
	    let desc = this.props.description
	    console.log("desc of form ", desc)
	    let comp = this
	    if(desc && desc.fields) {
	      let flds = desc.fields
	      if(flds) {
	        if(desc.info && desc.info.tabs) {
	          let tabs = new Array()
	          let tabsToDisp = desc.info && desc.info.tabs? desc.info.layout: Object.keys(desc.info.tabs)
	          tabsToDisp.forEach(function(k) {
	            let tabFlds = desc.info.tabs[k];
	            if(tabFlds) {
	              let tabArr = comp.layoutFields(tabFlds, flds, "tabfield formfield")
	              tabs.push(
	                <comp.uikit.Tab label={k} time={comp.state.time} value={k}>
	                  {tabArr}
	                </comp.uikit.Tab>
	              )
	            }
	          })
	          let vertical = desc.info.verticaltabs? true: false
	          return (
	            <this.uikit.Tabset vertical={vertical} time={comp.state.time}>
	              {tabs}
	            </this.uikit.Tabset>
	          )
	        } else {
	          let fldToDisp = desc.info && desc.info.layout? desc.info.layout: Object.keys(flds)
	          let className=comp.props.inline?"inline formfield":"formfield"
	          return this.layoutFields(fldToDisp, flds, className)
	        }
	      }
	    }
	    return null
	  }*/

	return blocks
}
