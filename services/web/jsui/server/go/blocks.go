package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"regexp"
	"strings"
)

type Node struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:"-"`
	Content []byte     `xml:",innerxml"`
	Nodes   []Node     `xml:",any"`
}

var (
	jsReplaceRegex    *regexp.Regexp
	jsVarReplaceRegex *regexp.Regexp
)

func (n *Node) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n.Attrs = start.Attr
	type node Node
	return d.DecodeElement((*node)(n), &start)
}

func (svc *UI) regItem(ctx core.ServerContext, itemType, itemName, cont string) {
	dispFunc := fmt.Sprintf("function(ctx, desc, uikit) { if(!ctx){ctx={};}if(!desc){desc={};}if(!uikit){uikit={};} console.log('ctx', ctx);return %s}", cont)
	//dispType := "EntityDisplay"
	svc.addRegItem(ctx, itemType, itemName, dispFunc)
}

func (svc *UI) createConfBlock(ctx core.ServerContext, itemType string, itemName string, conf config.Config) error {

	val, err := svc.processBlockConf(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Trace(ctx, "Processing conf", "itemName", itemName, "conf", conf, "val", val)
	svc.regItem(ctx, itemType, svc.getRegistryItemName(ctx, conf, itemName), val)
	return nil
}

func (svc *UI) createXMLBlock(ctx core.ServerContext, itemType string, itemName string, node Node) error {
	ctx = ctx.SubContext(fmt.Sprintf("%s_%s", itemType, itemName))
	val, err := svc.processXMLBlockNode(ctx, node)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	svc.regItem(ctx, itemType, itemName, val)
	return nil
}

func (svc *UI) processXMLBlockNode(ctx core.ServerContext, node Node) (string, error) {
	children := make([]string, 0)
	var mod = ""
	attrBuf := new(bytes.Buffer)
	content := node.Content

	processAttr := func(attrName, attrVal string, obj bool) error {
		if attrName == "module" {
			mod = attrVal
		} else {
			if attrBuf.Len() != 0 {
				attrBuf.WriteString(",")
			}
			val, err := utils.ProcessTemplate(ctx, []byte(attrVal), nil) //svc.processText(ctx, attr.Value, false)
			if err != nil {
				return errors.WrapError(ctx, err)
			}

			attrValStr := string(val)
			format := "%s:%s"
			if obj {
				conf, err := ctx.ReadConfigData([]byte(attrValStr), nil)
				if err != nil {
					log.Error(ctx, "error in reading config", "conf", conf, "str", attrValStr)
					return errors.WrapError(ctx, err)
				}
				attrValStr = processHierarchicalAttr(ctx, conf)
			} else {
				attrValStr = processJS(ctx, attrValStr)
			}
			attrBuf.WriteString(fmt.Sprintf(format, attrName, attrValStr))
		}
		return nil
	}

	for _, n := range node.Nodes {
		switch n.XMLName.Local {
		case "Attr":
			{
				name := ""
				obj := false
				for _, attr := range n.Attrs {
					if attr.Name.Local == "name" {
						name = n.Attrs[0].Value
					}
					if attr.Name.Local == "obj" {
						obj = true
					}
				}
				err := processAttr(name, string(n.Content), obj)
				if err != nil {
					return "", errors.WrapError(ctx, err)
				}
			}
		case "Content":
			{
				content = n.Content
			}
		default:
			{
				childTxt, err := svc.processXMLBlockNode(ctx, n)
				if err != nil {
					return "", errors.WrapError(ctx, err)
				}
				children = append(children, childTxt)
			}
		}
	}
	//attrs := make(map[string]interface{})
	for _, attr := range node.Attrs {
		val := attr.Value
		obj := strings.HasPrefix(val, "`")
		if obj {
			val = strings.TrimPrefix(val, "`")
		}
		err := processAttr(attr.Name.Local, val, obj)
		if err != nil {
			return "", err
		}
	}
	attrStr := fmt.Sprintf("{%s}", attrBuf.String())

	if (len(children) == 0) && len(content) > 0 {
		val, err := utils.ProcessTemplate(ctx, content, nil) //svc.processText(ctx, string(node.Content), true)
		if err != nil {
			return "", errors.WrapError(ctx, err)
		}
		children = append(children, fmt.Sprintf("%s", processJS(ctx, string(val))))
	}

	elem := ""
	switch mod {
	case "":
		elem = fmt.Sprintf("'%s'", node.XMLName.Local)
	case "uikit":
		elem = fmt.Sprintf("uikit.%s", node.XMLName.Local)
	default:
		elem = fmt.Sprintf("_$['%s'].%s", mod, node.XMLName.Local)
	}

	childArrStr := "null"
	if len(children) > 0 {
		childArrStr = fmt.Sprintf("[%s]", strings.Join(children, ","))
	}
	//n.Attrs
	return fmt.Sprintf("_ce(%s, %s, %s)", elem, attrStr, childArrStr), nil
}

func processJS(ctx core.ServerContext, input string) string {
	if jsReplaceRegex == nil {
		jsReplaceRegex, _ = regexp.Compile(`javascript#@#([a-zA-Z0-9\ _\,\(\)\'\.\(\)\[\]\{\}]+)#@#`)
	}
	if jsVarReplaceRegex == nil {
		jsVarReplaceRegex, _ = regexp.Compile(`javascript###([a-zA-Z0-9\ _\,\(\)\'\.\(\)\[\]\{\}]+)###`)
	}

	arr := jsReplaceRegex.FindAllStringIndex(input, -1)
	if len(arr) != 0 {
		return `"` + jsReplaceRegex.ReplaceAllString(input, `"+$1+"`) + `"`
	}

	arr = jsVarReplaceRegex.FindAllStringIndex(input, -1)
	if len(arr) == 0 {
		val, e := json.Marshal(input)
		if e != nil {
			log.Error(ctx, "Error in marshalling string", "string", input, "error", e)
		}
		return string(val)
	} else {
		return jsVarReplaceRegex.ReplaceAllString(input, `$1`)
	}
}

func processHierarchicalAttr(ctx core.ServerContext, conf config.Config) string {
	attrBuf := new(bytes.Buffer)
	attrNames := conf.AllConfigurations(ctx)
	for _, attrName := range attrNames {
		if attrBuf.Len() != 0 {
			attrBuf.WriteString(",")
		}
		strToWrite := ""
		strVal, ok := conf.GetString(ctx, attrName)
		if ok {
			strToWrite = processJS(ctx, strVal)
		}
		confVal, ok := conf.GetSubConfig(ctx, attrName)
		if ok {
			strToWrite = processHierarchicalAttr(ctx, confVal)
		}
		if strToWrite == "" {
			val, _ := conf.GetSubConfig(ctx, attrName)
			strval, err := json.Marshal(val)
			if err != nil {
				log.Error(ctx, "Error in marshalling", "err", err)
			} else {
				strToWrite = string(strval)
			}
		}
		attrBuf.WriteString(fmt.Sprintf("%s:%s", attrName, strToWrite))
	}
	return fmt.Sprintf("{%s}", attrBuf.String())
}

func (svc *UI) processBlockConf(ctx core.ServerContext, conf config.Config) (string, error) {

	//svc.processJS(ctx, conf)
	keys := conf.AllConfigurations(ctx)
	rootelem := ""
	for _, key := range keys {
		if key != "config" {
			rootelem = key
		}
	}

	if rootelem == "" {
		return "", errors.BadArg(ctx, "Json", "Reason", "Root element not provided for block")
	}
	root, ok := conf.GetSubConfig(ctx, rootelem)
	if ok {
		childStr := make([]string, 0)
		childrenArr, ok := root.GetConfigArray(ctx, "children")
		if ok {
			for _, child := range childrenArr {
				childTxt, err := svc.processBlockConf(ctx, child)
				if err != nil {
					return "", errors.WrapError(ctx, err)
				}
				childStr = append(childStr, childTxt)
			}
		}

		mod := ""
		attrBuf := new(bytes.Buffer)
		content := ""
		allAttributes := root.AllConfigurations(ctx)
		//attrs := make(map[string]interface{})
		for _, key := range allAttributes {
			switch key {
			case "body":
				content, _ = root.GetString(ctx, key)

				break
			case "children":
				break
			case "module":
				mod, _ = root.GetString(ctx, key)
				break
			default:
				if attrBuf.Len() != 0 {
					attrBuf.WriteString(",")
				}
				attrStr := ""
				str, ok := root.GetString(ctx, key)
				if ok {
					attrStr = processJS(ctx, str)
				} else {
					attrconf, ok := root.GetSubConfig(ctx, key)
					if ok {
						attrStr = processHierarchicalAttr(ctx, attrconf)
					} else {
						val, _ := root.Get(ctx, key)
						strval, err := json.Marshal(val)
						if err != nil {
							return "", errors.WrapError(ctx, err)
						}
						attrStr = string(strval)
					}
				}

				attrBuf.WriteString(fmt.Sprintf("%s:%s", key, attrStr))

			}
		}

		attrStr := fmt.Sprintf("{%s}", attrBuf.String())

		elem := ""
		if mod == "" {
			elem = fmt.Sprintf("'%s'", rootelem)
		} else {
			elem = fmt.Sprintf("_$['%s'].%s", mod, rootelem)
		}

		//n.Attrs
		if content != "" {
			return fmt.Sprintf("_ce(%s, %s, [%s])", elem, attrStr, processJS(ctx, content)), nil
		}
		return fmt.Sprintf("_ce(%s, %s, [%s])", elem, attrStr, strings.Join(childStr, ",")), nil

	}
	return "", nil
}
