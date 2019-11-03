package core

import (
	"crypto/aes"
	"crypto/cipher"
	"laatoo/sdk/common/config"
	"laatoo/sdk/server/components"
	"laatoo/sdk/server/core"
	"laatoo/sdk/server/errors"
	"laatoo/sdk/server/log"
	"laatoo/sdk/utils"
	"laatoo/server/constants"
	"path"
)

type secretsManagerImpl struct {
	storer    *utils.DiskStorer
	gcm       cipher.AEAD
	nonceSize int
}

func defaultSecretsManager(ctx core.ServerContext) (components.SecretsManager, error) {
	basedir, _ := ctx.GetString(config.BASEDIR)
	keysbase := path.Join(basedir, constants.CONF_KEYS_PATH)
	exists, _, _ := utils.FileExists(keysbase)
	if !exists {
		log.Warn(ctx, "No secrets manager provided")
		return nil, nil
	}
	log.Error(ctx, "Secret manager store found")
	c, err := aes.NewCipher([]byte(seedPass))
	// if there are any errors, handle them
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		return nil, errors.WrapError(ctx, err)
	}
	nonceSize := gcm.NonceSize()

	storer := utils.NewDiskStorer(keysbase, 5*1024*1024)
	return &secretsManagerImpl{gcm: gcm, storer: storer, nonceSize: nonceSize}, nil
}

func (mgr *secretsManagerImpl) Get(ctx core.ServerContext, key string) ([]byte, bool) {
	ciphertext, err := mgr.storer.GetObject(key)
	if err != nil {
		log.Error(ctx, "Error finding key", "err", err)
		return nil, false
	}

	if len(ciphertext) < mgr.nonceSize {
		log.Error(ctx, "Error decrypting key, malformed ciphertext", "err", err)
		return nil, false
	}

	val, err := mgr.gcm.Open(nil, ciphertext[:mgr.nonceSize], ciphertext[mgr.nonceSize:], nil)
	return val, true
}

func (mgr *secretsManagerImpl) Put(ctx core.ServerContext, key string, val []byte) error {
	return nil
}
