package domain

import (
	"encoding/json"
	"errors"
	"io"
)

var (
	//ErrNotUserNotFound ...
	ErrNotUserNotFound = errors.New("user not found")
)

type User struct {
	ID    int    `json:"-"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Roles string `json:"roles"`
	ClassCode string `json:"class_code"`
	Deleted      bool   `json:"-"`
	CreatedBy int `json:"created_by"`
}

//Passw ...
type Passw struct {
	Password    string `json:"password"`
	OldPassword string `json:"old_password"`
	RepeatPass  string `json:"confirm_password"`
}

//IsMatchPassword ...
func (u *User) IsMatchPassword(password string) bool {
	if u.Password != password {
		return false
	}
	return true
}

//FromJSON ...
func (u *User) FromJSON(data []byte) error {
	return json.Unmarshal(data, u)
}

//FromJSON ...
func (p *Passw) FromJSON(input io.Reader) error {
	return json.NewDecoder(input).Decode(p)
}

//CheckUser ...
type CheckUser struct {
	User   `json:"-"`
	IDUser int    `json:"id_user"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Roles string `json:"roles"`
}