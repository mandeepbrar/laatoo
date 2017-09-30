package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"regexp"
	"strings"
	"text/template"
)

type Node struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:"-"`
	Content []byte     `xml:",innerxml"`
	Nodes   []Node     `xml:",any"`
}

func (n *Node) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n.Attrs = start.Attr
	type node Node
	return d.DecodeElement((*node)(n), &start)
}

func (svc *UI) createDisplay(ctx core.ServerContext, itemType string, itemName string, node Node) error {
	val, err := svc.processNode(ctx, node)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	dispFunc := fmt.Sprintf("function(data, attrs) { if(!data){return null;} return %s}", val)
	//dispType := "EntityDisplay"
	svc.addRegItem(ctx, itemType, itemName, dispFunc)
	return nil
}

func (svc *UI) processText(ctx core.ServerContext, content string, replace bool) (string, error) {
	re := regexp.MustCompile(`\r?\n`)
	content = re.ReplaceAllString(content, ` `)

	contextVar := func(variable string) string {
		val, _ := ctx.GetString(variable)
		return val
	}

	entityVal := func(obj string) string {
		if replace {
			return "data." + obj
		}
		return "'+data." + obj + "+'"
	}

	attrVal := func(obj string) string {
		if replace {
			return "attrs." + obj
		}
		return "'+attrs." + obj + "+'"
	}

	funcMap := template.FuncMap{"var": contextVar, "entity": entityVal, "attribute": attrVal}

	temp, err := template.New("display").Funcs(funcMap).Parse(content)
	if err != nil {
		return "", errors.WrapError(ctx, err)
	}

	result := new(bytes.Buffer)
	anon := struct{}{}
	err = temp.Execute(result, anon)
	if err != nil {
		return "", errors.WrapError(ctx, err)
	}

	return result.String(), nil
}

func (svc *UI) processNode(ctx core.ServerContext, node Node) (string, error) {
	children := make([]string, 0)
	for _, n := range node.Nodes {
		childTxt, err := svc.processNode(ctx, n)
		if err != nil {
			return "", errors.WrapError(ctx, err)
		}
		children = append(children, childTxt)
	}
	/*
		display.WriteString("var c=[];")
		for _, c := range children {
			display.WriteString("c.push('%s')", c)
		}*/

	//element.WriteString("function() {")
	var err error
	attrs := make(map[string]interface{})
	for _, attr := range node.Attrs {
		attrs[attr.Name.Local], err = svc.processText(ctx, attr.Value, true)
		if err != nil {
			return "", errors.WrapError(ctx, err)
		}
	}

	attrStr, err := json.Marshal(attrs)
	if err != nil {
		return "", errors.WrapError(ctx, err)
	}

	if (len(children) == 0) && len(node.Content) > 0 {
		val, err := svc.processText(ctx, string(node.Content), false)
		if err != nil {
			return "", errors.WrapError(ctx, err)
		}
		children = append(children, fmt.Sprintf("'%s'", val))
	}

	//n.Attrs
	return fmt.Sprintf("_ce('%s', %s, [%s])", node.XMLName.Local, attrStr, strings.Join(children, ",")), nil
}
