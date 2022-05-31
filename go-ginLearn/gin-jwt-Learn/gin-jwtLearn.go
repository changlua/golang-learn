package gin_jwt_Learn

import (
	"log"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

//登录实体类
type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//token中存储相应信息的key名称
var identityKey = "id"

//测试需要认证接口
func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c) //取出claims
	user, _ := c.Get(identityKey) //获取自己提前定义好的身份信息
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*User).UserName,
		"text":     "Hello World.",
	})
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func Main() {
	//获取命令参数PORT，指定端口号
	port := os.Getenv("PORT")
	r := gin.Default()

	if port == "" {
		port = "8000"
	}

	//jwt插件
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{  //相应jwt插件的接口，我们这里去进行实例化
		Realm:       "test zone",  //身份
		Key:         []byte("secret key"),  //秘钥
		Timeout:     time.Minute,  //超时时长
		MaxRefresh:  time.Minute,  //最大的一个刷新时间
		IdentityKey: identityKey,  //身份标识key
		PayloadFunc: func(data interface{}) jwt.MapClaims {  //载荷信息，实际上这个data就是自定义执行的Authenticator方法的返回对象
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,  //存放的是对应的用户名
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {  //用于取出身份信息
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {  //login的身份认证
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &User{
					UserName:  userID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {  //权限认证
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {  //若是身份未认证成功。情况：①登录失败。②token不正确。③权限不足。
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",  //token查询方式：头部、query查询以及cookie携带
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",  //请求头的header的值

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	//1、登录接口
	//测试接口：http://localhost:8080/login
	/**
		{
			"username": "admin",
			"password": "admin"
		}
	 */
	r.POST("/login", authMiddleware.LoginHandler)  //LoginHandler：登录逻辑

	//2、无路由接口，执行之前会进行认证
	//测试接口：http://localhost:8080/test
	//对于没有路由的会进行认证校验
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	//3、需要认证接口
	auth := r.Group("/auth")
	// 接口：http://localhost:8080/auth/refresh_token，刷新token接口
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	//使用jwt认证插件
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		//接口：http://localhost:8080/auth/hello，用于查看当前的token信息
		auth.GET("/hello", helloHandler)
	}

	//指定端口运行
	//if err := http.ListenAndServe(":"+port, r); err != nil {
	//	log.Fatal(err)
	//}
	r.Run()
}