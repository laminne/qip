package key

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io"
)

type RSAKey struct {
	PrivateKey []byte
	PublicKey  []byte
}

func GenRSAKey() (RSAKey, error) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return RSAKey{}, err
	}

	key := priv.PublicKey
	derRsaPrivateKey := x509.MarshalPKCS1PrivateKey(priv)
	derRsaPublicKey := x509.MarshalPKCS1PublicKey(&key)

	keys := RSAKey{}

	var i bytes.Buffer
	err = pem.Encode(&i, &pem.Block{Type: "PRIVATE KEY", Bytes: derRsaPrivateKey})
	if err != nil {
		return RSAKey{}, err
	}

	keys.PrivateKey, err = io.ReadAll(&i)
	if err != nil {
		return RSAKey{}, err
	}

	var pub bytes.Buffer
	err = pem.Encode(&pub, &pem.Block{Type: "PUBLIC KEY", Bytes: derRsaPublicKey})
	if err != nil {
		return RSAKey{}, err
	}

	keys.PublicKey, err = io.ReadAll(&pub)
	if err != nil {
		return RSAKey{}, err
	}

	return keys, nil
}
