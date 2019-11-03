package main

import (
	"fmt"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"

	vault "github.com/hashicorp/vault/api"
)

const (
	CONF_VAULTSVC_SERVICENAME = "vaultsecretsservice"
	CONF_VS_HOST              = "host"
	CONF_VS_PATH              = "path"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_VAULTSVC_SERVICENAME, Object: VaultSecretsSvc{}}}
}

type VaultSecretsSvc struct {
	core.Service
	host   string
	path   string
	client *vault.Client
}

func (svc *VaultSecretsSvc) Initialize(ctx core.ServerContext, conf config.Config) error {
	svc.host, _ = svc.GetStringConfiguration(ctx, CONF_VS_HOST)
	svc.path, _ = svc.GetStringConfiguration(ctx, CONF_VS_PATH)
	//svc.path, _ = svc.GetStringConfiguration(ctx, CONF_VS_PATH)
	cfg := vault.DefaultConfig()
	cfg.Address = svc.host
	vc, err := vault.NewClient(cfg)
	if err != nil {
		return err
	}
	svc.client = vc
	//svc.client.SetToken(token)
	return nil
}

func (svc *VaultSecretsSvc) Get(ctx core.ServerContext, key string) ([]byte, bool) {

	sec, err := svc.client.Logical().Read(fmt.Sprintf("%s/%s", svc.path, key))
	if err != nil {
		log.Error(ctx, "Couldnt get the secret key ", "key", key, "err", err)
		return nil, false
	}
	data, ok := sec.Data["data"]
	if ok {
		byts, ok := data.([]byte)
		if ok {
			return byts, true
		}
	}
	return nil, false
}

func (svc *VaultSecretsSvc) Put(ctx core.ServerContext, key string, val []byte) error {

	_, err := svc.client.Logical().Write(fmt.Sprintf("%s/%s", svc.path, key), map[string]interface{}{"Data": val})
	if err != nil {
		return errors.WrapError(ctx, err)
	}
	return nil

}
