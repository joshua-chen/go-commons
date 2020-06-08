/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2020-05-16 23:24:17
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 15:27:14
 */
package jwt

import (
	"fmt"
	"github.com/joshua-chen/go-commons/config"
	"github.com/joshua-chen/go-commons/middleware/models"
	"github.com/joshua-chen/go-commons/mvc/context/response"
	"github.com/joshua-chen/go-commons/mvc/context/response/msg"
	_ "log"
	"strings"
	"sync"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go/request"
	_ "github.com/iris-contrib/middleware/cors"
	_ "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	_ "github.com/spf13/cast"
)

type (
	// A function called whenever an error is encountered
	errorHandler func(context.Context, string)

	// TokenExtractor is a function that takes a context as input and returns
	// either a token or an error.  An error should only be returned if an attempt
	// to specify a token was found, but the information was somehow incorrectly
	// formed.  In the case where a token is simply not present, this should not
	// be treated as an error.  An empty string should be returned in that case.
	TokenExtractor func(context.Context) (string, error)

	// Middleware the middleware for JSON Web tokens authentication method
	JWT struct {
		Config Config
	}
)

var (
	instance *JWT
	lock     *sync.Mutex = &sync.Mutex{}
)

func Instance() *JWT {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &JWT{}
		}
	}
	return instance
}

// jwt中间件配置
func Configure() *JWT {

	instance := Instance()

	c := Config{
		ContextKey: DefaultContextKey,
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(config.AppConfig.Secret), nil
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//重要的是要避免此处的安全问题：https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		ErrorHandler: func(ctx context.Context, errMsg string) {
			ctx.StopExecution()
			ctx.JSON(response.NewUnauthorizedResult(errMsg))
		},
		// 指定func用于提取请求中的token
		Extractor: FromAuthHeader,
		// if the token was expired, expiration error will be returned
		Expiration:          true,
		Debug:               true,
		EnableAuthOnOptions: false,
	}
	instance.Config = c
	//return &JWT{Config: c}
	golog.Debugf("instance.Config: %s",instance.Config)
	return instance
}

func Filter(ctx context.Context) bool {
	instance := Configure()
	if err := instance.CheckJWT(ctx); err != nil {
		golog.Errorf("Check jwt error, %s", err)
		return false
	}
	return true
	// If everything ok then call next.
	//ctx.Next()
}

/*
func NewJWT() *jwtmiddleware.Middleware {
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(config.AppConfig.Secret), nil
		},
		Extractor: FromAuthHeader,
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式

		ErrorHandler: func(ctx context.Context, err error) {
			if err == nil {
				return
			}
			ctx.StopExecution()
			ctx.JSON(response.NewUnauthorizedResult(err.Error()))
		},
		Expiration: true,
		//Debug:               true,
		EnableAuthOnOptions: false,
	})
	return jwtHandler
}
*/
func FromAuthHeader(ctx context.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", fmt.Errorf("Authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}

// below 3 method is get token from url
// FromParameter returns a function that extracts the token from the specified
// query string parameter
func FromParameter(param string) TokenExtractor {
	return func(ctx context.Context) (string, error) {
		return ctx.URLParam(param), nil
	}
}

func FromFirst(extractors ...TokenExtractor) TokenExtractor {
	return func(ctx context.Context) (string, error) {
		for _, ex := range extractors {
			token, err := ex(ctx)
			if err != nil {
				return "", err
			}
			if token != "" {
				return token, nil
			}
		}
		return "", nil
	}
}

// 在登录成功的时候生成token
func NewToken(user *models.User) (string, error) {
	//expireTime := time.Now().Add(60 * time.Second)
	expireTime := time.Now().Add(time.Duration(config.AppConfig.JwtTimeout) * time.Second)

	claims := Claims{
		user.Id,
		user.Username,
		//user.Password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "iris-casbins-jwt",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//tokenClaims := instance.NewWithClaims(instance.SigningMethodHS256, instance.MapClaims{
	//	"nick_name": "iris",
	//	"email":     "go-iris@qq.com",token不存在或header设置不正确
	//	"id":        "1",
	//	"iss":       "Iris",
	//	"iat":       time.Now().Unix(),
	//	"jti":       "9527",
	//	"exp":       time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(),
	//})

	token, err := tokenClaims.SignedString([]byte(config.AppConfig.Secret))
	return token, err
}

/*
func ParseToken(tokenString string, key string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}*/
func ParseToken(ctx context.Context) (*models.User, bool) {
	//token := GetToken(ctx)
	mapClaims := (Instance().Get(ctx).Claims).(jwt.MapClaims)

	id, ok1 := mapClaims["id"].(float64)
	username, ok2 := mapClaims["username"].(string)

	//golog.Infof("*** MapClaims=%v, [id=%f, ok1=%t]; [username=%s, ok2=%t]", mapClaims, id, ok1, username, ok2)
	if !ok1 || !ok2 {
		response.Error(ctx, iris.StatusInternalServerError, msg.TokenParseFailed)
		return nil, false
	}

	user := models.User{
		Id:       int64(id),
		Username: username,
	}
	return &user, true
}

func GetToken(ctx context.Context) string {
	token := ctx.GetHeader("Authorization")
	if token != "" && len(token) > 7 {
		token = token[7:]
	}
	return token
}

/*
func GetUserID(token string) int {
	var userId = 0
	if token != "" && token != "undefined" && len(token) > 7 {
		v, _ := ParseToken(token, JwtKey)
		if v != "" {
			userId = cast.ToInt(v.(jwt.MapClaims)["id"])
		}
	}
	return userId
}
*/

// Get returns the user (&token) information for this client/request
func (m *JWT) Get(ctx context.Context) *jwt.Token {
	golog.Debugf("ContextKey: %s",m.Config.ContextKey)
	golog.Debugf("ctx.Values(): %s",ctx.Values())
	golog.Debugf("m.Config: %s",m.Config)
	return ctx.Values().Get(m.Config.ContextKey).(*jwt.Token)
}

// CheckJWT the main functionality, checks for token
func (m *JWT) CheckJWT(ctx context.Context) error {
	if !m.Config.EnableAuthOnOptions {
		if ctx.Method() == iris.MethodOptions {
			return nil
		}
	}

	// Use the specified token extractor to extract a token from the request
	token, err := m.Config.Extractor(ctx)
	// If an error occurs, call the error handler and return an error
	if err != nil {
		golog.Debug("Error extracting JWT: %v", err)
		m.Config.ErrorHandler(ctx, msg.TokenExactFailur)
		return fmt.Errorf("Error extracting token: %v", err)
	}

	// If the token is empty...
	if token == "" {
		// Check if it was required
		if m.Config.CredentialsOptional {
			golog.Debug("No credentials found (CredentialsOptional=true)")
			// No error, just no token (and that is ok given that CredentialsOptional is true)
			return nil
		}

		golog.Debug("Error: No credentials found (CredentialsOptional=false)")
		// If we get here, the required token is missing
		m.Config.ErrorHandler(ctx, msg.TokenParseFailedAndEmpty)
		return fmt.Errorf(msg.TokenParseFailedAndEmpty)
	}

	// Now parse the token

	parsedToken, err := jwt.Parse(token, m.Config.ValidationKeyGetter)
	// Check if there was an error in parsing...
	if err != nil {
		golog.Errorf("Error parsing token1: %v", err)
		m.Config.ErrorHandler(ctx, msg.TokenExpired)
		return fmt.Errorf("Error parsing token2: %v", err)
	}

	if m.Config.SigningMethod != nil && m.Config.SigningMethod.Alg() != parsedToken.Header["alg"] {
		message := fmt.Sprintf("Expected %s signing method but token specified %s",
			m.Config.SigningMethod.Alg(),
			parsedToken.Header["alg"])
		golog.Errorf("Error validating token algorithm: %s", message)
		m.Config.ErrorHandler(ctx, msg.TokenParseFailed) // 算法错误
		return fmt.Errorf("Error validating token algorithm: %s", message)
	}

	// Check if the parsed token is valid...
	if !parsedToken.Valid {
		golog.Errorf(msg.TokenParseFailedAndInvalid)
		m.Config.ErrorHandler(ctx, msg.TokenParseFailedAndInvalid)
		return fmt.Errorf(msg.TokenParseFailedAndInvalid)
	}

	if m.Config.Expiration {
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
			if expired := claims.VerifyExpiresAt(time.Now().Unix(), true); !expired {
				return fmt.Errorf(msg.TokenExpired)
			}
		}
	}

	//m.logf("JWT: %v", parsedToken)

	// If we get here, everything worked and we can set the
	// user property in context.
	ctx.Values().Set(m.Config.ContextKey, parsedToken)

	return nil
}
