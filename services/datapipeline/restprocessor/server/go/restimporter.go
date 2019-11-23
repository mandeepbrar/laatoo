package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"laatoo/sdk/common/config"
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"net/http"
	"reflect"
	"time"
)

type RestImporter struct {
	core.Service
	restEndpoint string
	method       string
	objFac       core.ObjectFactory
	list         bool
}

func (imp *RestImporter) Initialize(ctx core.ServerContext, conf config.Config) error {
	svcEndpoint, ok := conf.GetString(ctx, CONF_INP_REST_ENDPOINT)
	if !ok {
		return errors.MissingConf(ctx, CONF_INP_REST_ENDPOINT)
	}
	imp.restEndpoint = svcEndpoint

	method, ok := conf.GetString(ctx, CONF_INP_REST_METHOD)
	if !ok {
		imp.method = "GET"
	} else {
		imp.method = method
	}

	obj, ok := conf.GetString(ctx, CONF_OBJECT_TO_IMPORT)
	if ok {
		imp.objFac, ok = ctx.(core.ServerContext).GetObjectFactory(obj)
		if !ok {
			return errors.BadConf(ctx, CONF_OBJECT_TO_IMPORT)
		}
	} else {
		return errors.MissingConf(ctx, CONF_OBJECT_TO_IMPORT)
	}

	imp.list, _ = conf.GetBool(ctx, CONF_INP_LIST)

	return nil
}

func (imp *RestImporter) GetRecords(ctx core.RequestContext, initData map[string]interface{}, dataChan datapipeline.DataChan) error {

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	var resp *http.Response
	var err error
	var rdr io.Reader

	if imp.method == "GET" {
		resp, err = netClient.Get(imp.restEndpoint)
	} else {
		resp, err = netClient.Post(imp.restEndpoint, "application/json", rdr)
	}
	if err != nil {
		return errors.WrapError(ctx, err)
	}

	if imp.list {
		respObj := imp.objFac.CreateObjectCollection(ctx, 0)
		respData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(respData, &respObj)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		arr := reflect.ValueOf(respObj)
		for i := 0; i < arr.Len(); i++ {
			dataChan <- datapipeline.NewPipelineRecord(arr.Index(i).Interface(), nil)
		}
	} else {
		respObj := imp.objFac.CreateObject(ctx)
		respData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		err = json.Unmarshal(respData, &respObj)
		if err != nil {
			return errors.WrapError(ctx, err)
		}
		dataChan <- datapipeline.NewPipelineRecord(respData, nil)
	}
	return nil
}

const (
	CONF_OBJECT_TO_IMPORT  = "importobject"
	CONF_INP_REST_ENDPOINT = "importrestendpoint"
	CONF_INP_LIST          = "importlist"
	CONF_INP_REST_METHOD   = "importmethod"
)