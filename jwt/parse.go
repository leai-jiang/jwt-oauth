package jwt

import (
	"encoding/json"
	"reflect"
	"strings"
)

type Parse struct {}

func (p *Parse) Parse(tokenString string, key string) (token *Token, err error) {
	token = &Token{}

	parts := strings.Split(tokenString, ".")

	// 签名不合规范，非正确签名
	if len(parts) != 3 {
		err = &ValidationError{Errors:ValidationErrorMalformed}
		return
	}

	var e error
	var headerBytes, payloadBytes []byte

	if headerBytes, e = DecodeSegment(parts[0]); e != nil {
		err = &ValidationError{Inner: e, Errors:ValidationErrorMalformed}
		return
	}
	if e = json.Unmarshal(headerBytes, &token.Header); e != nil {
		err = &ValidationError{Inner: e, Errors:ValidationErrorMalformed}
		return
	}

	if payloadBytes, e = DecodeSegment(parts[1]); e != nil {
		err = &ValidationError{Inner: e, Errors:ValidationErrorMalformed}
		return
	}

	if e = json.Unmarshal(payloadBytes, &token.Payload); e != nil {
		err = &ValidationError{Inner: e, Errors:ValidationErrorMalformed}
		return
	}

	method := token.Header.alg
	if reflect.TypeOf(method).String() == "string" {
		if token.Method = GetSigningMethod(method); token.Method == nil {
			err = NewValidationError("the signing method is not exist", ValidationErrorUnverifiable)
		}
	} else {
		err = NewValidationError("the (alg) in header is invalid", ValidationErrorUnverifiable)
		return
	}

	if err = token.Payload.Valid(); err != nil {
		return
	}

	token.Signature = parts[2]

	if e = token.Method.Verify(strings.Join(parts[0:2], "."), token.Signature, key); e != nil {
		err = &ValidationError{Inner:e}
		return
	}

	token.Valid = true

	return token, nil
}

