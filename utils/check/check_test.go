package check

import (
	"fmt"
	"testing"
)

func TestPassWordEncrypted(t *testing.T) {
	encrypted := PassWordEncrypted("716523")
	//$2a$04$OUgxddqdX/zSAi3vXmjX9.a75Yh3gnRXibP4esT4.aleQ6hU7eKeS
	fmt.Println(encrypted)
}
