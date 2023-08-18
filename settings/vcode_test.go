package settings

import (
	"testing"
)

func TestVcode(t *testing.T) {
	vcode := VerificationCode()
	t.Log(vcode)
}
