package jwt

import "encoding/base64"

type Token struct {
	Method      SigningMethod
	Header 		map[string]interface{}
	Payload 	Payload
	Signature 	string
	Valid 		bool
}

func New() *Token {}

func (t *Token) SignedString() (string, error) {

}

func EncodeSegment(seg []byte) string {
	return base64.URLEncoding.EncodeToString(seg)
}

func DecodeSegment(encodedSeg string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(encodedSeg)
}