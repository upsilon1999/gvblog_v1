package jwt

import (
	"github.com/dgrijalva/jwt-go/v4"
)

// JwtPayLoad jwt中payload数据
type JwtPayLoad struct {
	Username string `json:"username"` // 用户名
	NickName string `json:"nickName"` // 昵称
	Role     int    `json:"role"`      // 权限  1 管理员  2 普通用户  3 游客
	UserID   uint   `json:"userId"`   // 用户id
  }
  
  type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
  }
  
  
  


