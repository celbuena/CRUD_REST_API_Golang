package response

import (
	"encoding/json"
	"net/http"
)

const (
	//ErrNotFound ...
	ErrNotFound = `Data tidak ditemukan`
	//ErrCommonServer ...
	ErrCommonServer = `Terjadi kesalahan pada server. Silahkan coba beberapa saat lagi`
	//ErrLogin ...
	ErrLogin = `Invalid Login`
	//SuccessRes ...
	SuccessRes = `SUCCESS`
)

//Response is default response
type rs struct {
	Code         int         `json:"code"`
	Status       bool        `json:"status"`
	ErrorMessage string      `json:"error_message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

//Message ...
type Message interface {
	Default(w http.ResponseWriter)
	DefaultHTML(w http.ResponseWriter)
	IsError() bool
}

//Default is ...
func (r *rs) Default(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	json.NewEncoder(w).Encode(r)
}

//DefaultHTML ...
func (r *rs) DefaultHTML(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	encoder.Encode(r)
}

func (r *rs) IsError() bool {
	return !r.Status
}

//NotFound ...
func NotFound() Message {
	return &rs{
		ErrorMessage: ErrNotFound,
		Code:         http.StatusNotFound,
	}
}

//InternalError ...
func InternalError() Message {
	return &rs{
		ErrorMessage: ErrCommonServer,
		Code:         http.StatusInternalServerError,
	}
}

//Success ...
func Success(data interface{}) Message {
	return &rs{
		Code:   http.StatusOK,
		Status: true,
		Data:   data,
	}
}

//Errors ...
func Errors(code int, msg string) Message {
	return &rs{
		Code:         code,
		ErrorMessage: msg,
	}
}

//SuccessWithCode ...
func SuccessWithCode(code int, data interface{}) Message {
	return &rs{
		Code:   code,
		Status: true,
		Data:   data,
	}
}
