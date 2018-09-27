package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
)

const (
	CONF_ENTITY      = "entity"
	CONF_FIELDS      = "fields"
	CONF_FORM_LAYOUT = "layout"
	CONF_WIDGET      = "widget"
	CONF_FORMINFO    = "info"
	CONF_WIDGET_MOD  = "widgetMod"
	CONF_WIDGET_CONF = "widgetConf"
	CONF_FIELDTYPE   = "type"
	CONF_FIELDREQD   = "required"
)

func (svc *UI) regForm(ctx core.ServerContext, itemType, itemName, cont string) {
	//formFunc := fmt.Sprintf("%s", cont)
	//dispType := "EntityDisplay"
	svc.addRegItem(ctx, itemType, itemName, cont)
}

func (svc *UI) createForm(ctx core.ServerContext, itemType string, itemName string, conf config.Config) error {
	/*obj := make(map[string]interface{})
	log.Error(ctx, "yaml block", "content", string(cont))
	err := yaml.Unmarshal(cont, &obj)
	if err != nil {
		log.Error(ctx, "unmarshalling err", "err", err)
		return errors.WrapError(ctx, err)
	}
	log.Error(ctx, "unmarshalled", "content", obj)*/
	/*val, err := svc.processBlockConf(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}*/
	log.Info(ctx, " Creating form", "form", itemName)
	formStr := new(bytes.Buffer)
	err := svc.buildFormSchema(ctx, itemType, itemName, conf, formStr)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	/*val, err := json.Marshal(formMap)
	if err != nil {
		return errors.WrapError(ctx, err)
	}*/
	svc.regForm(ctx, itemType, svc.getRegistryItemName(ctx, conf, itemName), formStr.String())
	return nil
}

func (svc *UI) createXMLForm(ctx core.ServerContext, itemType string, itemName string, node Node) error {
	conf := ctx.CreateConfig()
	formConfig, err := svc.buildXMLFormConfig(ctx, node)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	if len(node.Nodes) > 0 {
		layoutName := fmt.Sprintf("%s_FormLayout", itemName)
		svc.createXMLBlock(ctx, "Blocks", layoutName, node)
		formConfig.Set(ctx, "layout", layoutName)
	}
	conf.Set(ctx, "info", formConfig)
	formFields := ctx.CreateConfig()
	svc.getFields(ctx, node, formFields)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	conf.Set(ctx, "fields", formFields)
	svc.createForm(ctx, itemType, itemName, conf)
	return nil
}

func (svc *UI) processXMLFieldChildItems(ctx core.ServerContext, node Node, name, widget string, conf config.Config) {
	if widget == "Select" {
		if len(node.Nodes) > 0 {
			items := make([]config.Config, 0)
			for _, childNode := range node.Nodes {
				if childNode.XMLName.Local == "Item" {
					item := ctx.CreateConfig()
					for _, attr := range childNode.Attrs {
						item.Set(ctx, attr.Name.Local, attr.Value)
					}
					items = append(items, item)
				}
			}
			conf.Set(ctx, "items", items)
		}
	}
}

func (svc *UI) getFields(ctx core.ServerContext, node Node, conf config.Config) {
	for _, childNode := range node.Nodes {
		if childNode.XMLName.Local == "Field" {
			fieldName := ""
			widget := ""
			fieldConf := ctx.CreateConfig()
			for _, attr := range childNode.Attrs {
				if attr.Name.Local == "name" {
					fieldName = attr.Value
				} else {
					if attr.Name.Local == "widget" {
						widget = attr.Value
					}
					fieldConf.Set(ctx, attr.Name.Local, attr.Value)
				}
			}
			svc.processXMLFieldChildItems(ctx, childNode, fieldName, widget, fieldConf)
			conf.Set(ctx, fieldName, fieldConf)
		} else {
			svc.getFields(ctx, childNode, conf)
		}
	}
}

func (svc *UI) buildXMLFormConfig(ctx core.ServerContext, node Node) (config.Config, error) {
	conf := ctx.CreateConfig()
	for _, attr := range node.Attrs {
		conf.Set(ctx, attr.Name.Local, attr.Value)
	}
	return conf, nil
}

func (svc *UI) buildEntitySchema(ctx core.ServerContext, entityName string, formconf config.Config, fieldMap *bytes.Buffer) error {
	return nil
}

func (svc *UI) createField(ctx core.ServerContext, fieldName string, fieldType string, required bool, widget, widgetMod string, conf config.Config, fieldMap *bytes.Buffer) error {
	ctx = ctx.SubContext("Create Field: " + fieldName)
	fieldAttrs := conf.Clone()
	fieldAttrs.Set(ctx, "name", fieldName)
	if widget == "" {
		switch fieldType {
		case config.OBJECTTYPE_STRING:
			fieldAttrs.Set(ctx, "widget", "TextField")
			break
		case config.OBJECTTYPE_INT:
			fieldAttrs.Set(ctx, "widget", "NumberField")
			widgetMod = ""
			break
		case config.OBJECTTYPE_BOOL:
			fieldAttrs.Set(ctx, "widget", "Switch")
			widgetMod = ""
			break
		case config.OBJECTTYPE_STRINGARR:
			fieldAttrs.Set(ctx, "widget", "ListField")
			widgetMod = ""
			break
		case config.OBJECTTYPE_DATETIME:
			fieldAttrs.Set(ctx, "widget", "DatePicker")
			break
		case "entity":
			fieldAttrs.Set(ctx, "widget", "Select")
			break
		case "image":
			fieldAttrs.Set(ctx, "widget", "ImagePicker")
			break
		}
	}
	fieldsStr, _ := json.Marshal(fieldAttrs)
	fieldMap.WriteString(fmt.Sprintf("%s:%s", fieldName, fieldsStr))
	return nil
}

func (svc *UI) buildFormSchema(ctx core.ServerContext, itemType string, itemName string, conf config.Config, formStr *bytes.Buffer) error {
	fieldMap := new(bytes.Buffer)
	/*entity, ok := conf.GetString(ctx, CONF_ENTITY)
	if ok {
		err := svc.buildEntitySchema(ctx, entity, conf, fieldMap)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
	} else {*/
	fields, ok := conf.GetSubConfig(ctx, CONF_FIELDS)
	if !ok {
		return errors.MissingConf(ctx, CONF_FIELDS)
	}
	fieldNames := fields.AllConfigurations(ctx)
	for _, fieldName := range fieldNames {
		if fieldMap.Len() > 0 {
			fieldMap.WriteString(",")
		}
		fieldConf, _ := fields.GetSubConfig(ctx, fieldName)
		fieldType, ok := fieldConf.GetString(ctx, CONF_FIELDTYPE)
		if !ok {
			fieldType = config.OBJECTTYPE_STRING
		}
		required, ok := fieldConf.GetBool(ctx, CONF_FIELDREQD)
		widget, ok := fieldConf.GetString(ctx, CONF_WIDGET)
		widgetMod, ok := fieldConf.GetString(ctx, CONF_WIDGET_MOD)
		svc.createField(ctx, fieldName, fieldType, required, widget, widgetMod, fieldConf, fieldMap)
	}
	//}
	/*layout, ok := conf.GetString(ctx, CONF_FORM_LAYOUT)
	layoutStr := ""
	if ok {
		layoutStr = fmt.Sprintf(",template: templateLayout('%s')", layout)
		//optionsMap.Set(ctx, "template", fmt.Sprintf("<Panel id=\"%s\" />", layout))
	}*/
	formInfoStr := ""
	formInfo, ok := conf.GetSubConfig(ctx, CONF_FORMINFO)
	if ok {
		info, err := json.Marshal(formInfo)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		formInfoStr = fmt.Sprintf(",info: %s", string(info))
	}
	formStr.WriteString(fmt.Sprintf("{fields:{%s} %s}", fieldMap.String(), formInfoStr))

	return nil
}