package password

type PasswordVerifier interface {
	Verify(password string, actualPassword string) bool
}

type PlainPasswordVerifier struct{}

func (p *PlainPasswordVerifier) Verify(password string, actualPassword string) bool {
	return password == actualPassword
}
