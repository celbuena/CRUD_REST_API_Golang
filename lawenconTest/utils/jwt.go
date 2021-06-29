package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type key int

const (
	jwtclaim = key(47)
	secret = "secret"
	//jwtform  = `^[A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$`
)

var (
	jKeys = []byte(secret)
	//jKeys = []byte(os.Getenv("JWT_SALT"))
	//ErrInvalidToken ...
	ErrInvalidToken = errors.New(`Invalid tokens`)
)

type Receiver struct {
	UserId int		`json:"user_id"`
	Email string	`json:"email"`
	Roles string `json:"roles"`
}

//JWTClaim is part of inside jwt
type JWTClaim interface {
	SetJWT (string, error)
	SetValueToContext(ctx context.Context) context.Context
}

//claim is struct for claim
type claim struct {
	Receiver
	jwt.StandardClaims
}

//SetJWT
func(c *claim)SetJWT()(string, error){
	fmt.Println(c, "ini error claim")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(jKeys)
}

//set value to context
func (c *claim)SetValueToContext(ctx context.Context) context.Context{
	ct := context.WithValue(ctx, jwtclaim, c.Receiver)
	return ct
}

//Get value to context
func GetValueFromContext (ctx context.Context) *Receiver{
	rs, ok :=ctx.Value(jwtclaim).(Receiver)
	if !ok{
		return &Receiver{}
	}

	return &rs
}

//New Claim
func NewClaim(userid int, email string, roles string ) *claim {
	return &claim{
		Receiver: Receiver{userid, email, roles},
	}
}

//ParseJWT ...
func ParseJWT(tn string) (*claim, error) {
	c := new(claim)

	//if match, _ := regexp.MatchString(jwtform, tn); !match {
	//	return nil, errors.New("not jwt format")
	//}

	tk, err := jwt.ParseWithClaims(tn, c,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jKeys, nil
		})

	//fmt.Println("ini errorku", tn, err)

	if err == jwt.ErrSignatureInvalid || !tk.Valid {
		return nil, ErrInvalidToken
	}

	if err != nil {
		return nil, err
	}
	return c, nil
}
