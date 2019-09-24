package midware

import (
	"github.com/gin-gonic/gin"
)

func JwtMidware(c *gin.Context) {
	//// sample token string taken from the New example
	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"
	//
	//// Parse takes the token string and a function for looking up the key. The latter is especially
	//// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	//// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	//// to the callback, providing flexibility.
	//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	//	// Don't forget to validate the alg is what you expect:
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	//	}
	//
	//	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	//	return hmacSampleSecret, nil
	//})
	//
	//if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	//	fmt.Println(claims["foo"], claims["nbf"])
	//} else {
	//	fmt.Println(err)
	//}
	//
	//token := c.Request.Header.Get("token")
	//if token == "" {
	//	c.JSON(http.StatusOK, gin.H{
	//		"status": -1,
	//		"msg":    "请求未携带token，无权限访问",
	//	})
	//	c.Abort()
	//	return
	//}
	//c.Next()

}
