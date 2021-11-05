package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/dgrijalva/jwt-go/v4/test"
	"golang.org/x/xerrors"
	"time"
)

func main() {
	mySigningKey := []byte("AllYourBal")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(15000),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
	TimeOutToken()
	TimeOutHmacToken()
	Token2()
}

func TimeOutHmacToken()  {
	// sample token string taken from the New example
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.SnPaU08qj8PXhKVaneoNgSzoUPf8ieDnY5_xozuEeSE"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")

		return []byte("AllYourBal"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["Foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
}

func TimeOutToken() {
	// Token from another example.  This token is expired
	var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	var uErr *jwt.UnverfiableTokenError
	var expErr *jwt.TokenExpiredError
	var nbfErr *jwt.TokenNotValidYetError

	// Use xerrors.Is to see what kind of error(s) occurred
	if token.Valid {
		fmt.Println("You look nice today")
	} else if xerrors.As(err, &uErr) {
		fmt.Println("That's not even a token")
	} else if xerrors.As(err, &expErr) {
		fmt.Println("Timing is everything")
	} else if xerrors.As(err, &nbfErr) {
		fmt.Println("Timing is everything")
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
}

func Token2()  {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.SnPaU08qj8PXhKVaneoNgSzoUPf8ieDnY5_xozuEeSE"

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// sample token is expired.  override time so it parses as valid
	test.At(time.Unix(0,0),func() {
		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("AllYourBal"), nil
		})

		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt.Unix())
		} else {
			fmt.Println(err)
		}
	})
}