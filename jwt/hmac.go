package jwt

import (
	"crypto"
	"crypto/hmac"
	"errors"
)

type SigningMethodHMAC struct {
	Name string
	Hash crypto.Hash
}

var (
	SigningMethodHS256  *SigningMethodHMAC
	SigningMethodHS384  *SigningMethodHMAC
	SigningMethodHS512  *SigningMethodHMAC
	ErrSignatureInvalid = errors.New("signature is invalid")
)

func init() {
	// HS256
	SigningMethodHS256 = &SigningMethodHMAC{"HS256", crypto.SHA256}
	RegisterSigningMethod(SigningMethodHS256.Alg(), func() SigningMethod {
		return SigningMethodHS256
	})

	// HS384
	SigningMethodHS384 = &SigningMethodHMAC{"HS384", crypto.SHA384}
	RegisterSigningMethod(SigningMethodHS384.Alg(), func() SigningMethod {
		return SigningMethodHS384
	})

	// HS512
	SigningMethodHS512 = &SigningMethodHMAC{"HS512", crypto.SHA512}
	RegisterSigningMethod(SigningMethodHS512.Alg(), func() SigningMethod {
		return SigningMethodHS512
	})
}

func (m *SigningMethodHMAC) Alg() string {
	return m.Name
}

func (m *SigningMethodHMAC) Verify(signingString, signature string, key string) error {
	if !m.Hash.Available() {
		return ErrHashUnavailable
	}

	hasher := hmac.New(m.Hash.New, []byte(key))
	hasher.Write([]byte(signingString))
	if !hmac.Equal([]byte(signature), hasher.Sum(nil)) {
		return ErrSignatureInvalid
	}

	return nil
}

func (m *SigningMethodHMAC) Sign(signingString, key string) (string, error) {
	if !m.Hash.Available() {
		return "", ErrHashUnavailable
	}

	hasher := hmac.New(m.Hash.New, []byte(key))
	hasher.Write([]byte(signingString))

	return string(hasher.Sum(nil)), nil
}