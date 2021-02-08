package api

import (
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

// TestJWTMiddleware 测试jwt中间件
func TestJWTMiddleware(t *testing.T) {
	// 设置签名方法，默认jwt.SigningMethodHS256
	signingMethod := jwt.GetSigningMethod("HS256")
	// 实例化jwt对象
	tokenObj := jwt.New(signingMethod)
	// 设置载荷
	tokenObj.Claims = jwt.MapClaims{
		"live": "true",
	}
	// 签名
	token, _ := tokenObj.SignedString([]byte("test"))
	// 打印token
	fmt.Println(token)

	// 请求头 Authorization
	token = fmt.Sprintf("bearer %v", token)
	fmt.Println("Authorization", token)

	// Authorization bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsaXZlIjoidHJ1ZSJ9.h-RiquHM5wt13ut17MPxCasJFCFolC-irj4KbPXeSQM
}
