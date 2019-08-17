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

func (svc *UI) processXMLFieldChildItems(ctx core.ServerContext, node Node, conf config.Config) {
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

func (svc *UI) getFields(ctx core.ServerContext, node Node, conf config.Config) {
	for _, childNode := range node.Nodes {
		if childNode.XMLName.Local == "Field" {
			fieldConf := ctx.CreateConfig()
			fieldName := svc.populateFieldNodeConf(ctx, childNode, fieldConf)
			conf.Set(ctx, fieldName, fieldConf)
		} else {
			svc.getFields(ctx, childNode, conf)
		}
	}
}

func (svc *UI) populateFieldNodeConf(ctx core.ServerContext, node Node, conf config.Config) string {
	widgetConf := ctx.CreateConfig()
	propsConf := ctx.CreateConfig()
	widgetConf.Set(ctx, "props", propsConf)
	conf.Set(ctx, "widget", widgetConf)
	widget := ""
	fieldName := ""

	for _, attr := range node.Attrs {
		switch attr.Name.Local {
		case "name":
			fieldName = attr.Value
		case "entity":
			conf.Set(ctx, attr.Name.Local, attr.Value)
		case "list":
			conf.Set(ctx, attr.Name.Local, attr.Value)
		case "type":
			conf.Set(ctx, attr.Name.Local, attr.Value)
		case "ref":
			conf.Set(ctx, attr.Name.Local, attr.Value)
		case "widget":
			widget = attr.Value
			widgetConf.Set(ctx, "name", widget)
		case "module":
			widgetConf.Set(ctx, attr.Name.Local, attr.Value)
		default:
			propsConf.Set(ctx, attr.Name.Local, attr.Value)
		}
	}
	if widget == "Select" {
		svc.processXMLFieldChildItems(ctx, node, widgetConf)
	}
	return fieldName
}

func (svc *UI) populateFieldWidgetConf(ctx core.ServerContext, node Node, conf config.Config) {
	for _, attr := range node.Attrs {
		set := false
		switch attr.Name.Local {
		case "widget":
			set = true
		case "module":
			set = true
		case "classname":
			set = true
		}
		if set {
			conf.Set(ctx, attr.Name.Local, attr.Value)
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

func (svc *UI) createField(ctx core.ServerContext, fieldName string, fieldType string, required bool, widget, conf config.Config, fieldMap *bytes.Buffer) error {
	ctx = ctx.SubContext("Create Field: " + fieldName)
	fieldAttrs := conf.Clone()
	fieldAttrs.Set(ctx, "name", fieldName)
	if widget == nil {
		widgetConf := ctx.CreateConfig()
		fieldAttrs.Set(ctx, "widget", widgetConf)
		switch fieldType {
		case config.OBJECTTYPE_STRING:
			widgetConf.Set(ctx, "name", "TextField")
			break
		case config.OBJECTTYPE_INT:
			widgetConf.Set(ctx, "name", "NumberField")
			break
		case config.OBJECTTYPE_BOOL:
			widgetConf.Set(ctx, "name", "Switch")
			break
		case config.OBJECTTYPE_STRINGARR:
			widgetConf.Set(ctx, "name", "ListField")
			break
		case config.OBJECTTYPE_DATETIME:
			widgetConf.Set(ctx, "name", "DatePicker")
			break
		case "entity":
			widgetConf.Set(ctx, "name", "Select")
			break
		case "image":
			widgetConf.Set(ctx, "name", "ImagePicker")
			break
		}
	} else {
		fieldAttrs.Set(ctx, "widget", widget)
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
		widget, ok := fieldConf.GetSubConfig(ctx, CONF_WIDGET)
		svc.createField(ctx, fieldName, fieldType, required, widget, fieldConf, fieldMap)
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
