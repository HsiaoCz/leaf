package jwt

import "testing"

// TestJwt 测试Token的生成和验证解析Token
func TestJwt(t *testing.T) {
	var userID int64 = 12222333
	token, err := GenToken(userID)
	if err != nil {
		t.Fatal(err)
	}
	claims, err := ParseToken(token)
	if err != nil {
		t.Fatal(err)
	}
	if userID == claims.UserID {
		t.Log("ok")
	}
}
