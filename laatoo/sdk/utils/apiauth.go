package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	//	"laatoosdk/log"
)

func EncryptWithKey(publickeypath string, message string) ([]byte, error) {
	publickey, err := loadPublicKey(publickeypath)
	if err != nil {
		return nil, err
	}
	encryptedmsg, err := rsa.EncryptOAEP(md5.New(), rand.Reader, publickey, []byte(message), []byte(""))
	if err != nil {
		return nil, err
	}
	return encryptedmsg, nil
}

func loadPublicKey(path string) (key *rsa.PublicKey, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("ssh: no key found")
	}
	switch block.Type {
	case "PUBLIC KEY":
		keyInt, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return keyInt.(*rsa.PublicKey), nil
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}
}
