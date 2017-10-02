package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/log"
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
	dispFunc := fmt.Sprintf("function(data, attrs) { if(!data){data={};}if(!attrs){attrs={};} return %s}", val)
	//dispType := "EntityDisplay"
	svc.addRegItem(ctx, itemType, itemName, dispFunc)
	return nil
}

func (svc *UI) processText(ctx core.ServerContext, content string, forceInsert bool) (string, error) {
	re := regexp.MustCompile(`\r?\n`)
	content = re.ReplaceAllString(content, ` `)

	contextVar := func(args ...string) string {
		val, _ := ctx.GetString(args[0])
		if forceInsert || (len(args) > 1 && args[1] == "insert") {
			return "'+" + val + "+'"
		}
		return val
	}

	dataVal := func(args ...string) string {
		if forceInsert || (len(args) > 1 && args[1] == "insert") {
			return "'+data." + args[0] + "+'"
		}
		return "data." + args[0]
	}

	attrVal := func(args ...string) string {
		if forceInsert || (len(args) > 1 && args[1] == "insert") {
			return "'+attrs." + args[0] + "+'"
		}
		return "attrs." + args[0]
	}

	funcMap := template.FuncMap{"var": contextVar, "data": dataVal, "attribute": attrVal}

	temp, err := template.New("display").Funcs(funcMap).Parse(content)
	if err != nil {
		return "", errors.WrapError(ctx, err)
	}
	//temp = temp.

	result := new(bytes.Buffer)
	log.Error(ctx, "**********8888", "result", result)
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
	//var err error
	var mod = ""
	attrBuf := new(bytes.Buffer)

	//attrs := make(map[string]interface{})
	for _, attr := range node.Attrs {
		if attr.Name.Local == "module" {
			mod = attr.Value
		} else {
			if attrBuf.Len() != 0 {
				attrBuf.WriteString(",")
			}
			log.Error(ctx, "**********8888", "attr value", attr.Value)
			val, err := svc.processText(ctx, attr.Value, false)
			if err != nil {
				return "", errors.WrapError(ctx, err)
			}
			attrVal, err := json.Marshal(val)
			if err != nil {
				return "", errors.WrapError(ctx, err)
			}
			attrValStr := string(attrVal)
			format := "%s:%s"
			if strings.HasPrefix(attrValStr, "\"`") {
				format = "%s:%s"
				attrValStr = strings.TrimPrefix(attrValStr, "\"`")
				attrValStr = strings.TrimSuffix(attrValStr, "\"")
			}
			attrBuf.WriteString(fmt.Sprintf(format, attr.Name.Local, attrValStr))
			//attrs[attr.Name.Local], err = svc.processText(ctx, attr.Value, false)
			/*if err != nil {
				return "", errors.WrapError(ctx, err)
			}*/
		}
	}
	attrStr := fmt.Sprintf("{%s}", attrBuf.String())
	/*attrStr, err := json.Marshal(attrs)
	if err != nil {
		return "", errors.WrapError(ctx, err)
	}*/

	if (len(children) == 0) && len(node.Content) > 0 {
		val, err := svc.processText(ctx, string(node.Content), true)
		if err != nil {
			return "", errors.WrapError(ctx, err)
		}
		children = append(children, fmt.Sprintf("'%s'", val))
	}
	elem := ""
	if mod == "" {
		elem = fmt.Sprintf("'%s'", node.XMLName.Local)
	} else {
		elem = fmt.Sprintf("_$['%s'].%s", mod, node.XMLName.Local)
	}
	//n.Attrs
	return fmt.Sprintf("_ce(%s, %s, [%s])", elem, attrStr, strings.Join(children, ",")), nil
}
