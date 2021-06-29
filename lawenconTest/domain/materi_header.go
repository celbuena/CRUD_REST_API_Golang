package domain

import (
	"encoding/json"
	"errors"
)

var (
	ErrNotMateriHeaderNotFound = errors.New("materi_header not found")
)

type Materi_Header struct {
	ID int `json:"-"`
	IdMapel int `json:"id_mapel"`
	Chapter int `json:"chapter"`
	Title string `json:"title"`
	Label string `json:"label"`
	Materi string `json:"materi"`
	Detail string `json:"detail"`
	StatusMateri int `json:"status_materi"`
	Deleted int `json:"deleted"`
}

type UpdateMateriHeader struct {
	Id int `json:"id"`
	Chapter int `json:"chapter"`
	Title string `json:"title"`
	Label string `json:"label"`
	Materi string `json:"materi"`
	Detail string `json:"detail"`
	StatusMateri string `json:"status_materi"`
}

//CheckMateriHeader ...
type CheckMateriHeader struct {
	Materi_Header   `json:"-"`
	IDSchool int `json:"id_user"`
}

//FromJSON ...
func (u *Materi_Header) FromJSON(data []byte) error {
	return json.Unmarshal(data, u)
}

//Compose ...
func (ch *CheckMateriHeader) Compose() {
	ch.IDSchool = ch.Materi_Header.ID
}


func (u *UpdateMateriHeader)FromJSON (data []byte)error{
	return json.Unmarshal(data, u)
}
