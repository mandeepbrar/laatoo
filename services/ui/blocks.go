package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	"laatoo/sdk/utils"
	"strings"
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

func (svc *UI) regItem(ctx core.ServerContext, itemType, itemName, cont string) {
	dispFunc := fmt.Sprintf("function(data, attrs) { if(!data){data={};}if(!attrs){attrs={};} return %s}", cont)
	//dispType := "EntityDisplay"
	svc.addRegItem(ctx, itemType, itemName, dispFunc)
}

func (svc *UI) createConfBlock(ctx core.ServerContext, itemType string, itemName string, conf config.Config) error {
	/*obj := make(map[string]interface{})
	log.Error(ctx, "yaml block", "content", string(cont))
	err := yaml.Unmarshal(cont, &obj)
	if err != nil {
		log.Error(ctx, "unmarshalling err", "err", err)
		return errors.WrapError(ctx, err)
	}
	log.Error(ctx, "unmarshalled", "content", obj)*/
	val, err := svc.processConf(ctx, conf)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	svc.regItem(ctx, itemType, itemName, val)
	return nil
}

/*
func (svc *UI) createJsonBlock(ctx core.ServerContext, itemType string, itemName string, cont []byte) error {
	obj := make(map[string]interface{})
	err := json.Unmarshal(cont, &obj)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	val, err := svc.processMap(ctx, obj)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	svc.regItem(ctx, itemType, itemName, val)
	return nil
}*/

func (svc *UI) createXMLBlock(ctx core.ServerContext, itemType string, itemName string, node Node) error {
	val, err := svc.processXMLNode(ctx, node)
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	svc.regItem(ctx, itemType, itemName, val)
	return nil
}

/*
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
	anon := struct{}{}
	err = temp.Execute(result, anon)
	if err != nil {
		return "", errors.WrapError(ctx, err)
	}

	return result.String(), nil
}
*/
func (svc *UI) processXMLNode(ctx core.ServerContext, node Node) (string, error) {
	children := make([]string, 0)
	for _, n := range node.Nodes {
		childTxt, err := svc.processXMLNode(ctx, n)
		if err != nil {
			return "", errors.WrapError(ctx, err)
		}
		children = append(children, childTxt)
	}

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
			val, err := utils.ProcessTemplate(ctx, []byte(attr.Value), nil) //svc.processText(ctx, attr.Value, false)
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
		}
	}
	attrStr := fmt.Sprintf("{%s}", attrBuf.String())

	if (len(children) == 0) && len(node.Content) > 0 {
		val, err := utils.ProcessTemplate(ctx, node.Content, nil) //svc.processText(ctx, string(node.Content), true)
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

/*
func (svc *UI) processMapAttr(ctx core.ServerContext, obj map[string]interface{}) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	for k, v := range res {
		strV, strok := v.(string)
		if strok {
			strval, err := svc.processText(ctx, strV, false)
			if err != nil {
				return nil, errors.WrapError(ctx, err)
			}
			res[k] = strval
		} else {
			mapV, mapok := v.(map[string]interface{})
			if mapok {
				resmap, err := svc.processMapAttr(ctx, mapV)
				if err != nil {
					return nil, err
				}
				res[k] = resmap
			} else {
				res[k] = v
			}
		}
	}
	return res, nil
}*/

func (svc *UI) processConf(ctx core.ServerContext, conf config.Config) (string, error) {
	keys := conf.AllConfigurations(ctx)
	if len(keys) > 1 {
		return "", errors.BadArg(ctx, "Json", "Reason", "More than one roots of Json")
	}
	rootelem := keys[0]
	root, ok := conf.GetSubConfig(ctx, rootelem)
	if ok {
		childStr := make([]string, 0)
		childrenArr, ok := root.GetConfigArray(ctx, "children")
		if ok {
			for _, child := range childrenArr {
				childTxt, err := svc.processConf(ctx, child)
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
				/*strval, err := svc.processText(ctx, val, false)
				if err != nil {
					return "", errors.WrapError(ctx, err)
				}
				content = strval*/
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
				val, _ := root.Get(ctx, key)
				strval, err := json.Marshal(val)
				if err != nil {
					return "", errors.WrapError(ctx, err)
				}
				attrBuf.WriteString(fmt.Sprintf("%s:%s", key, string(strval)))
				/*				mapval, ok := val.(map[string]interface{})
								if ok {
									mapres, err := svc.processMapAttr(ctx, mapval)
									if err != nil {
										return "", errors.WrapError(ctx, err)
									}
									mapstr, err := json.Marshal(mapres)
									if err != nil {
										return "", errors.WrapError(ctx, err)
									}
									attrBuf.WriteString(fmt.Sprintf("%s:%s", key, string(mapstr)))
								} else {
									if strok {
										attrBuf.WriteString(fmt.Sprintf("%s:'%s'", key, strVal))
									}
								}*/
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
			return fmt.Sprintf("_ce(%s, %s, ['%s'])", elem, attrStr, content), nil
		}
		return fmt.Sprintf("_ce(%s, %s, [%s])", elem, attrStr, strings.Join(childStr, ",")), nil

	}
	return "", nil
}