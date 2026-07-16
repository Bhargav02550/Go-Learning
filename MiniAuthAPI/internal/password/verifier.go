package password

type PasswordVerifier interface {
	VerifyPassword(password string, actualPassword string) bool
}

type PlainPasswordVerifier struct{}

func (p *PlainPasswordVerifier) VerifyPassword(password string, actualPassword string) bool {
	return password == actualPassword
}
