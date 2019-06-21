package core

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"
)

type JWTHeader struct {
	typ string
	alg  string
}

type JWTPayload struct {
	iss string
	exp time.Time
	sub string
	aud string
	nbf time.Time
	iat time.Time
	jti string
	from string
	id string
}

var jwtHeader = JWTHeader{
	typ: "sha256",
	alg: "jwt",
}

func NewToken(secret string, jwtPayload JWTPayload) string {
	headerStr, payloadStr := generateBody(jwtHeader, jwtPayload)
	return headerStr + "." + payloadStr + "." + sign(secret, headerStr, payloadStr)
}

func Payload(secret string, token string) (*JWTPayload, error) {
	tokenArr := strings.Split(token, ".")
	if len(tokenArr) != 3 {
		return nil, errors.New("PARSE TOKEN ERROR")
	}

	header := tokenArr[0]
	payload := tokenArr[1]
	signature := tokenArr[2]

	if sign(secret, header, payload) != signature {
		return nil, errors.New("PARSE TOKEN ERROR")
	}

	var p *JWTPayload
	payloadStr, _ := base64.StdEncoding.DecodeString(payload)
	err := json.Unmarshal([]byte(payloadStr), &p)
	if err != nil {
		return nil, errors.New("UNMARSHAL PAYLOAD FAILED")
	}

	if p.exp.Before(time.Now()) {
		return nil, errors.New("TOKEN HAS EXPIRED")
	}

	return p, nil
}

func JWTHTTPMiddleware(h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func generateBody(jwtHeader JWTHeader, jwtPayload JWTPayload) (string, string) {
	headerByte, _ := json.Marshal(jwtHeader)
	headerStr := base64.StdEncoding.EncodeToString(headerByte)

	payloadByte, err := json.Marshal(jwtPayload)
	if err != nil {
		log.Panic("The error happened in 'json.Marshal':", err)
	}
	payloadStr := base64.StdEncoding.EncodeToString(payloadByte)

	return headerStr, payloadStr
}

func sign(secret string, header string, payload string) string {

	signature := sha256.Sum256([]byte(header + "." + payload + secret))

	return string(signature[0:])
}


