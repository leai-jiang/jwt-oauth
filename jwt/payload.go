package jwt

import (
	"fmt"
	"time"
)

type Payload interface {
	// 验证签名是否已经生效
	VerifyNotBefore(date int64) bool

	// 验证签名是否已经过期
	VerifyExpiresAt(date int64) bool

	// 验证是否满足上述两个条件
	Valid() error
}

type StandardPayload struct {
	Id        string `json:"jti,omitempty"`

	// 主题
	Subject   string `json:"sub,omitempty"`

	// 签发人
	Issuer    string `json:"iss,omitempty"`

	// 签发时间
	IssuedAt  int64  `json:"iat,omitempty"`

	// 生效时间
	NotBefore int64  `json:"nbf,omitempty"`

	// 接收人
	Audience  string `json:"aud,omitempty"`

	// 过期时间
	ExpiresAt int64  `json:"exp,omitempty"`
}

func (sp *StandardPayload) VerifyNotBefore(date int64) bool {
	return sp.NotBefore >= date
}

func (sp *StandardPayload) VerifyExpiresAt(date int64) bool {
	return sp.ExpiresAt <= date
}

func (sp *StandardPayload) Valid() error {
	vErr := new(ValidationError)
	now := time.Now().Unix()

	if sp.VerifyExpiresAt(now) == false {
		delta := time.Unix(now, 0).Sub(time.Unix(sp.ExpiresAt, 0))
		vErr.Inner = fmt.Errorf("token is expired by %v", delta)
		vErr.Errors |= ValidationErrorExpired
	}

	if sp.VerifyNotBefore(now) == false {
		vErr.Inner = fmt.Errorf("token is not valid yet")
		vErr.Errors |= ValidationErrorNotValidYet
	}

	if vErr.valid() {
		return nil
	}

	return vErr
}
