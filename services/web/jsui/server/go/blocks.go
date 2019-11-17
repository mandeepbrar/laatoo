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

	//"laatoo/sdk/utils"
	"strings"
)

type Node struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:"-"`
	Content []byte     `xml:",innerxml"`
	Nodes   []Node     `xml:",any"`
}

func (node *Node) HasChildren() bool {
	return len(node.Nodes) > 0
}

func (node *Node) HasContent() bool {
	return len(node.Content) > 0
}

func (n *Node) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n.Attrs = start.Attr
	type node Node
	return d.DecodeElement((*node)(n), &start)
}

func (svc *UI) regItem(ctx core.ServerContext, itemType, itemName, cont string) {
	dispFunc := fmt.Sprintf("function(ctx, desc) { if(!ctx){ctx={};}if(!desc){desc={};} console.log('ctx', ctx);return %s}", cont)
	//dispType := "EntityDisplay"
	svc.addRegItem(ctx, itemType, itemName, dispFunc)
}

func (svc *UI) createConfBlock(ctx core.ServerContext, itemType string, itemName string, conf config.Config) error {

	val, err := svc.processBlockConf(ctx, conf, itemName)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Trace(ctx, "Processing conf", "itemName", itemName, "conf", conf, "val", val)
	svc.regItem(ctx, itemType, svc.getRegistryItemName(ctx, conf, itemName), val)
	return nil
}

func (svc *UI) createXMLBlock(ctx core.ServerContext, itemType string, itemName string, node Node) error {
	ctx = ctx.SubContext(fmt.Sprintf("%s_%s", itemType, itemName))
	log.Error(ctx, "creating xml block", "itemName", itemName)
	/*val, err := svc.processXMLBlockNode(ctx, node, itemName)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	svc.regItem(ctx, itemType, itemName, val)*/
	//conf := ctx.CreateConfig()
	conf, _, err := svc.createConfFromXML(ctx, node)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	log.Error(ctx, " xml block conf", "conf", conf)
	err = svc.createConfBlock(ctx, itemType, itemName, conf)
	if err != nil {
		log.Error(ctx, " error processing block conf", "conf", conf, "err", err)
		return errors.WrapError(ctx, err)
	}
	return nil
}

func (svc *UI) createConfFromXML(ctx core.ServerContext, node Node) (config.Config, string, error) {
	switch node.XMLName.Local {
	case "Attr":
		{
			nodeConf := ctx.CreateConfig()
			name := ""
			for _, attr := range node.Attrs {
				if attr.Name.Local == "name" {
					name = attr.Value
				}
			}
			if name == "" {
				return nil, "", errors.BadConf(ctx, "Attribute name not provided ")
			}
			if node.HasChildren() {
				childConf := ctx.CreateConfig()
				for _, attrChild := range node.Nodes {
					if attrChild.XMLName.Local == "Content" {
						childConf.SetString(ctx, "body", string(attrChild.Content))
					} else {
						aconf, _, err := svc.createConfFromXML(ctx, attrChild)
						if err != nil {
							return nil, "", errors.WrapError(ctx, err)
						}
						childConfNames := aconf.AllConfigurations(ctx)
						for _, confName := range childConfNames {
							c, _ := aconf.Get(ctx, confName)
							childConf.Set(ctx, confName, c)
						}
					}
				}
				nodeConf.Set(ctx, name, childConf)
			} else {
				nodeConf.SetString(ctx, name, string(node.Content))
			}
			log.Error(ctx, "Processing heirarchical attrs", "node", node, "nodeConf", nodeConf)
			return nodeConf, name, nil
		}
	default:
		{
			nodeConf := ctx.CreateConfig()
			if len(node.Nodes) == 0 && len(node.Content) > 0 {
				nodeConf.SetString(ctx, "body", string(node.Content))
			}
			for _, attr := range node.Attrs {
				nodeConf.SetString(ctx, attr.Name.Local, attr.Value)
			}
			if node.HasChildren() {
				childrenConf := make([]config.Config, 0)
				for _, n := range node.Nodes {
					isAttr := n.XMLName.Local == "Attr"
					//childConf := ctx.CreateConfig()
					/*if n.XMLName.Local == "Attr" {
						err := svc.createConfFromXML(ctx, n, nodeconf)
						if err!=nil {
							return errors.WrapError(ctx, err)
						}
					} else {*/
					//nconf := ctx.CreateConfig()
					childConf, name, err := svc.createConfFromXML(ctx, n)
					if err != nil {
						return nil, "", errors.WrapError(ctx, err)
					}
					//nodeconf.Set(ctx, n.XMLName.Local, nconf)
					//}
					if isAttr {
						aconf, _ := childConf.Get(ctx, name)
						nodeConf.Set(ctx, name, aconf)
					} else {
						childrenConf = append(childrenConf, childConf)
					}
				}
				nodeConf.Set(ctx, "children", childrenConf)
			} else if node.HasContent() {
				nodeConf.SetString(ctx, "body", string(node.Content))
			}
			nConf := ctx.CreateConfig()
			nConf.Set(ctx, node.XMLName.Local, nodeConf)
			return nConf, node.XMLName.Local, nil
		}
	}
	/*if childConfOnly {
		return conf, nil
	}
	nodeConf := ctx.CreateConfig()
	nodeConf.Set(ctx, node.XMLName.Local, conf)
	return nodeConf, nil*/
}

/*
func (svc *UI) processXMLBlockNode(ctx core.ServerContext, node Node, itemName string) (string, error) {
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
				childTxt, err := svc.processXMLBlockNode(ctx, n, itemName)
				if err != nil {
					return "", errors.WrapError(ctx, err)
				}
				children = append(children, childTxt)
			}
		}
	}
	log.Error(ctx, "Reading attributes")

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

	elem, elemname := getModString(ctx, node.XMLName.Local, mod, itemName)
	log.Error(ctx, "element returned", "mod", mod, "local", node.XMLName.Local, "children", children)
	childArrStr := "null"
	if len(children) > 0 {
		childArrStr = fmt.Sprintf("[%s]", strings.Join(children, ","))
	}
	//n.Attrs
	return fmt.Sprintf("_ce(%s, %s, %s, %s)", elem, attrStr, childArrStr, elemname), nil
}
*/
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
		attrBuf.WriteString(fmt.Sprintf("\"%s\":%s", attrName, strToWrite))
	}
	return fmt.Sprintf("{%s}", attrBuf.String())
}

func getModString(ctx core.ServerContext, elem, mod, itemname string) (string, string) {
	switch mod {
	case "":
		return fmt.Sprintf("_uikit.%s", elem), fmt.Sprintf("'%s.uikit.%s'", itemname, elem)
	case "html":
		return fmt.Sprintf("'%s'", elem), fmt.Sprintf("'%s.%s.%s'", itemname, mod, elem)
	default:
		return fmt.Sprintf("_$['%s'].%s", mod, elem), fmt.Sprintf("'%s.%s.%s'", itemname, mod, elem)
	}
}

func (svc *UI) processBlockConf(ctx core.ServerContext, conf config.Config, itemName string) (string, error) {

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
				childTxt, err := svc.processBlockConf(ctx, child, itemName)
				if err != nil {
					return "", errors.WrapError(ctx, err)
				}
				childStr = append(childStr, childTxt)
			}
		}

		mod := ""
		attrBuf := new(bytes.Buffer)
		content := ""
		visibility := ""
		allAttributes := root.AllConfigurations(ctx)
		//attrs := make(map[string]interface{})
		for _, key := range allAttributes {
			switch key {
			case "body":
				content, _ = root.GetString(ctx, key)

				break
			case "children":
				break
			case "visibility":
				str, ok := root.GetString(ctx, key)
				if ok {
					visibility = processJS(ctx, str)
				} else {
					val, ok := root.GetBool(ctx, key)
					if ok {
						visibility = fmt.Sprint(val)
					}
				}
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
						log.Error(ctx, "attrStr for ", "itemName", itemName, "key", key, "attrStr", attrStr)
					} else {
						val, _ := root.Get(ctx, key)
						strval, err := json.Marshal(val)
						if err != nil {
							return "", errors.WrapError(ctx, err)
						}
						attrStr = string(strval)
					}
				}

				attrBuf.WriteString(fmt.Sprintf("\"%s\":%s", key, attrStr))

			}
		}

		attrStr := fmt.Sprintf("{%s}", attrBuf.String())
		elem, elemName := getModString(ctx, rootelem, mod, itemName)

		output := ""
		//n.Attrs
		if content != "" {
			output = fmt.Sprintf("_ce(%s, %s, [%s], %s)", elem, attrStr, processJS(ctx, content), elemName)
		} else if len(childStr) > 0 {
			output = fmt.Sprintf("_ce(%s, %s, [%s], %s)", elem, attrStr, strings.Join(childStr, ","), elemName)
		} else {
			output = fmt.Sprintf("_ce(%s, %s, null, %s)", elem, attrStr, elemName)
		}
		log.Error(ctx, "visibility condition", "elemname", elemName, "visibility", visibility)
		if visibility != "" {

			return fmt.Sprintf(" (%s) ? (%s) : null", visibility, output), nil
		} else {
			return output, nil
		}
	}
	return "", nil
}
