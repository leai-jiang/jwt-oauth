package jwt

import (
	"encoding/json"
	"reflect"
	"strings"
)

type Parse struct {}

func (p *Parse) Parse(tokenString string, key string) (token *Token, err error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {

	}

	token = &Token{}

	var headerBytes, payloadBytes []byte

	if headerBytes, err = DecodeSegment(parts[0]); err != nil {

	}
	if err = json.Unmarshal(headerBytes, &token.Header); err != nil {

	}

	if payloadBytes, err = DecodeSegment(parts[1]); err != nil {

	}

	if err = json.Unmarshal(payloadBytes, &token.Payload); err != nil {

	}

	method := token.Header.alg
	if reflect.TypeOf(method).String() == "string" {
		if token.Method = GetSigningMethod(method); token.Method == nil {

		}
	} else {

	}

	if err = token.Payload.Valid(); err != nil {

	}

	token.Signature = parts[2]

	if err = token.Method.Verify(strings.Join(parts[0:2], "."), token.Signature, key); err != nil {

	}
}

