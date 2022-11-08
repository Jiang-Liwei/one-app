package verifycode

type Store interface {

	// SetStore 保存验证码
	SetStore(key string, value string) bool

	// GetStore 获取验证码
	GetStore(key string, clear bool) string

	// VerifyStore 验证码是否存在
	VerifyStore(id, answer string, clear bool) bool
}
