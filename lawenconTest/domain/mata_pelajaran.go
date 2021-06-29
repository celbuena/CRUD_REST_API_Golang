package domain

import (
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrNotMatpelNotFound = errors.New("mata_pelajaran not found")
)

type Mata_Pelajaran struct {
	ID 			  int `json:"-"`
	IdTingkat     int    `json:"id_tingkat"`
	IdParentMapel int `json:"id_parent_mapel"`
	Icon          string `json:"icon"`
	Mapel string `json:"mapel"`
}

type UpdateInfoMapel struct {
	IdTingkat     int    `json:"id_tingkat"`
	Icon string `json:"icon"`
	IdParentMapel int `json:"id_parent_mapel"`
	Mapel string `json:"mapel"`
}
//CheckMatpel ...
type CheckMatpel struct {
	Mata_Pelajaran `json:"-"`
	IDMatpel       int `json:"id_Matpel"`
}

func (ch *CheckMatpel) Compose() {
	ch.IDMatpel = ch.Mata_Pelajaran.ID
}

//FromJSON ...
func (u *Mata_Pelajaran) FromJSON(data []byte) error {
	return json.Unmarshal(data, u)
}

func (u *UpdateInfoMapel) FromJSON(data []byte) error{
	return json.Unmarshal(data, u)
}

func (u *School) SetTimeMapel() {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

