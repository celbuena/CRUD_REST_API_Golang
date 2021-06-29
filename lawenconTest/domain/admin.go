package domain

import (
	"encoding/json"
	"errors"
	"io"
)

var (
	ErrNotAdminNotFound = errors.New("admin not found")
)

type Admin struct {
	ID       int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Roles    string `json:"roles"`
	Deleted  bool   `json:"deleted"`
}

type AdminInfo struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Roles string `json:"roles"`
}
type UpdateAdminInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Roles string `json:"roles"`
}

func (u *Admin) SetRole(roles string) {
	u.Roles = roles
}

func (u *UpdateAdminInfo) SetRole(roles string) {
	u.Roles = roles
}

func (u *Admin) SetAdminID(id int) {
	u.ID = id
}

type SuperAdmin struct {
	Admin
}

type Finance struct {
	Admin
}

type Content struct {
	Admin
}

//cek password
type Passwd struct {
	Password    string `json:"password"`
	OldPassword string `json:"old_password"`
	RepeatPass  string `json:"repeat_pass"`
}

func (u *Admin) IsMatchPassword(password string) bool {
	if u.Password != password {
		return false
	}
	return true
}

func (u *Admin) FromJSON(data []byte) error {
	return json.Unmarshal(data, u)
}

func (u *AdminInfo) FromJSON (data []byte) error{
	return json.Unmarshal(data,u)
}

func (u *UpdateAdminInfo) FromJson(data []byte) error {
	return json.Unmarshal(data, u)
}

//FromJSON ...
func (p *Passwd) FromJSON(input io.Reader) error {
	return json.NewDecoder(input).Decode(p)
}

func (s *Admin) Model() Profile {
	return s
}
