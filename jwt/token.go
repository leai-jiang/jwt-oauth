package jwt

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type Header struct {
	// 签名类型 json web token
	typ string

	// 签名方法名称
	alg string
}

type Token struct {
	Method      SigningMethod
	Header 		*Header
	Payload 	Payload
	Signature 	string
	Valid 		bool
}

func New(m SigningMethod, p Payload) *Token {
	return &Token{
		Method: m,
		Header: &Header{
			typ: "jwt",
			alg: m.Alg(),
		},
		Payload: p,
	}
}

func (t *Token) SignedString(key string) (string, error) {
	var sig, str string
	var err error
	if str, err = t.SigningString(key); err != nil {
		return "", err
	}
	if sig, err = t.Method.Sign(str, key); err != nil {
		return "", err
	}
	return strings.Join([]string{str, sig}, "."), nil
}

func (t *Token) SigningString(key string) (string, error) {
	var err error
	parts := make([]string, 2)
	for i, _ := range parts {
		var jsonValue []byte
		if i == 0 {
			if jsonValue, err = json.Marshal(t.Header); err != nil {
				return "", err
			}
		} else {
			if jsonValue, err = json.Marshal(t.Payload); err != nil {
				return "", err
			}
		}
		parts[i] = EncodeSegment(jsonValue)
	}
	return strings.Join(parts, "."), nil
}

func EncodeSegment(seg []byte) string {
	return base64.URLEncoding.EncodeToString(seg)
}

func DecodeSegment(encodedSeg string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(encodedSeg)
}