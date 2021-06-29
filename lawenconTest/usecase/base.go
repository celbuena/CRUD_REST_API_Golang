package usecase

import (
	"lawenconTest/domain"
	"lawenconTest/response"
)


type Base interface {
	Signup(body []byte, role string) response.Message
	Login(body []byte) response.Message
	GetAllAdmin(adminid int) response.Message
	GetAllTransaction(transactionid int) response.Message
	UpdateInfoAdmin(id string, body []byte) response.Message
	DeleteAdmin(body []byte) response.Message

	//user
	RegisterSchool(body []byte, role string) response.Message
	UpdateInfoSchool(id string, body []byte, roles string) response.Message
	DeleteSchool(id string, body []byte, roles string) response.Message
	GetAllUser(userid int) response.Message

	//mata_pelajaran
	RegisterMapel(body []byte, role string) response.Message
	GetAllMapel(idMapel int) response.Message
	UpdateInfoMapel(id string, body []byte, roles string)response.Message
	DeleteMapel(body []byte, roles string)response.Message

	//materi_header
	CreateMateriHeader(body []byte) response.Message
	GetMateriHeader(materi_headerId int) response.Message
	DeleteMateri(body []byte, roles string) response.Message
	UpdateInfoMateriHeader(id string, body []byte) response.Message
}

type usecase struct {
	admin domain.BaseService
}

func NewServices(admin domain.BaseService) Base {
	return &usecase{admin: admin}
}