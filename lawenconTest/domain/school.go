package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

var (
	//ErrNotSchoolNotFound ...
	ErrNotSchoolNotFound = errors.New("school not found")
)

type School struct {
	ID          int       `json:"-"`
	SchoolName string `json:"school_name"`
	SchoolCode string `json:"school_code"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	Deleted     bool      `json:"-"`
	CreatedBy 	int 		`json:"created_by"`
}

type UpdateSchoolInfo struct {
	Id int `json:"id"`
	SchoolName string `json:"school_name"`
	SchoolCode string `json:"school_code"`
	Deleted bool `json:"deleted"`
}

//CheckMateriHeader ...
type CheckSchool struct {
	School   `json:"-"`
	IDSchool int `json:"id_user"`
}

//FromJSON ...
func (u *School) FromJSON(data []byte) error {
	return json.Unmarshal(data, u)
}

//Compose ...
func (ch *CheckSchool) Compose() {
	ch.IDSchool = ch.School.ID
}

//CreateToken ...
func (u *School) CreateToken(IDParents, IDUser int, IsAccept bool) string {
	return fmt.Sprintf(`%v:%v:%v:%v`, IDUser, u.ID, strconv.FormatBool(IsAccept), u.SchoolName)
}

func (u *School) SetTime() {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *UpdateSchoolInfo) FromJson(data []byte) error {
	return json.Unmarshal(data, u)
}
