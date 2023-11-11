package jwt

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "請登入",
			})
			c.Abort()
			return
		}
		if VerifyToken(token) {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "token 錯誤",
			})
			c.Abort()
			return
		}

	}
}

// Create the JWT key used to create the signature
// 建立用於建立簽名的 JWT 金鑰
var jwtKey = []byte("my_secret_key")

// For simplification, we're storing the users information as an in-memory map in our code
// 為了簡化，我們將使用者資訊存儲為程式碼中的記憶體映射
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Create a struct to read the username and password from the request body
// 建立一個結構體以從請求正文中讀取使用者名稱和密碼
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
// 建立一個結構體，將被編碼為 JWT。
// 我們添加 jwt.RegisteredClaims 作為嵌入式類型，以提供像到期時間這樣的欄位
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateToken(name string) (string, error) {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	// 宣告 token 的到期時間
	// 在這裡，我們將其保留為 5 分鐘
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: name,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			// 在 JWT 中，到期時間表示為 Unix 毫秒
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	// 宣告使用於簽署的演算法和聲明的 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	// 建立 JWT 字串
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		// 如果建立 JWT 時發生錯誤，則返回內部伺服器錯誤

		return "", errors.New("error in creating token")
	}
	return tokenString, nil
}

func ParseToken(token string) (interface{}, error) {
	// Initialize a new instance of `Claims`
	claims := &Claims{}
	// Parse the JWT string and store the result in `claims`.
	// If the token is invalid (if it has expired according to the expiry time we set on sign in), the following code will return an error
	// 解析 JWT 字串並將結果存儲在 `claims` 中。
	// 如果 token 無效（如果根據我們在登入時設置的到期時間已過期），則以下程式碼將返回錯誤
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, jwt.ErrSignatureInvalid
		}
		return nil, err
	}
	if tkn.Valid == false {
		return nil, errors.New("token is invalid")
	}
	return claims, nil
}

func RefreshToken(token string) (interface{}, error) {

	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, jwt.ErrSignatureInvalid
		}
		return nil, err
	}
	if tkn.Valid == false {
		return nil, errors.New("token is invalid")
	}
	// (END) The code until this point is the same as the first part of the `ParseToken` Func
	// (END) 到此為止的程式碼與 `ParseToken` 函式的第一部分相同

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	// 我們確保在足夠的時間過去之前不會發行新的 token
	// 在這種情況下，只有在舊 token 在到期前 30 秒內時才會發行新 token。否則，返回錯誤的請求狀態

	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		return nil, errors.New("Token refresh time is too fast")
	}
	// Now, create a new token for the current use, with a renewed expiration time
	// 現在，為當前使用者創建一個新的 token，並更新到期時間
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	// Create the token using the token claims and the key
	// 使用 token 聲明和金鑰建立 token
	newtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the specified key
	// 使用指定的金鑰簽署 token
	tokenString, err := newtoken.SignedString(jwtKey)
	if err != nil {
		return nil, errors.New("error in creating token")
	}
	return tokenString, nil
}

func LayoutToken(token string) {

}

func VerifyToken(token string) bool {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return false
	}
	newclaims, ok := tkn.Claims.(*Claims)
	if !ok {
		return false
	}
	if newclaims.ExpiresAt.Unix() < time.Now().Unix() {
		fmt.Println("Token is expired")
		return false
	}
	return true
}
