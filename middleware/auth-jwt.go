package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"productBackend/config"
	"strings"
	"time"
)
type MyCustomClaims struct {
	UID int `json:"uid"`
	jwt.StandardClaims
}
//SignWxToken 生成token,uid用户id，expireSec过期秒数
func SignWxToken(uid uint, expireSec int) (tokenStr string, err error) {
	// 带权限创建令牌
	claims := make(jwt.MapClaims)
	claims["uid"] = uid
	claims["admin"] = false
	sec := time.Duration(expireSec)
	claims["exp"] = time.Now().Add(time.Second * sec).Unix() //自定义有效期，过期需要重新登录获取token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用自定义字符串加密 and get the complete encoded token as a string
	tokenStr, err = token.SignedString([]byte(config.App.SigninKey))
	return tokenStr, err
}

//WxAuth ...
func WxAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authString := c.Request.Header.Get("Authorization")
		kv := strings.Split(authString, " ")
		if len(kv) != 2 || kv[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"error": "authorization header is not exists1",
			})
			c.Abort()
			return
		}
		tokenString := kv[1]

		// Parse token
		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.App.SigninKey), nil
		})

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": "authorization header is not exists2",
			})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusOK, gin.H{
				"error": "authorization header is not exists3",
			})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(*MyCustomClaims)

		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"error": "authorization header is not exists4",
			})
			c.Abort()
			return
		}
		//将uid写入请求参数
		uid := claims.UID
		c.Set("uid", uid)

		c.Next()
	}
}

func AuthJWT(ctx *gin.Context) {
	// 得到 请求头 authorization
	authH := ctx.GetHeader("Authorization")
	// 校验是否存在
	if "" == authH {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "authorization header is not exists",
			"data":authH,
		})
		// 发现问题，终止请求的处理
		ctx.Abort()
		return
	}

	// 格式判定
	if !strings.HasPrefix(authH, "Bearer ") {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "authorization format error",
		})
		ctx.Abort()
		return
	}

	// 剥离 token 部分
	 tokenStr := authH[7:]
	// 校验 token 的合理性
	//token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
	//	return []byte(config.App.SigninKey), nil
	//})
	// Parse token
	token, _ := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.App.SigninKey), nil
	})
	// 判断 token 是否合法
	if !token.Valid {
		// 不合法
		ctx.JSON(http.StatusOK, gin.H{
			"error": "token is invalid",
		})
		ctx.Abort()
		return
	}
	claims, ok := token.Claims.(*MyCustomClaims)

	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "authorization header is not exists4",
		})
		ctx.Abort()
		return
	}
	//将uid写入请求参数
	uid := claims.UID
	ctx.Set("uid", uid)
	// 认证通过，继续执行
	ctx.Next()
}