package jwt

type SigningMethod interface {
	// 验证签名
	Verify(signingString, signedToken string, key string) error

	// 加密签名
	Sign(signingString string, key string) (string, error)

	// 返回签名加密方法名称
	Alg() string
}

var signingMethods map[string]func() SigningMethod

func RegisterSigningMethod(alg string, m func() SigningMethod) {
	signingMethods[alg] = m
}
