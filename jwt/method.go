package jwt

type SigningMethod interface {
	Verify(signingString, signedToken string, key interface{}) error
	Sign(signingString string, key interface{}) (string, error)
	Alg() string
}
