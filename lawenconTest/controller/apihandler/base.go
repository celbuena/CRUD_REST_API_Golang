package apihandler

import ("lawenconTest/usecase"
	"net/http"
)

type ctrl struct {
	admin usecase.Base
}

type Controllers interface {
	//admin
	SignUpHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
	GetAllAdminHandler(w http.ResponseWriter, r *http.Request)
	GetAllTransactionHandler(w http.ResponseWriter, r *http.Request)
	UpdateAdminHandler(w http.ResponseWriter, r *http.Request)
	DeleteAdminHandler(w http.ResponseWriter, r *http.Request)

	//user
	RegisterSchoolHandler(w http.ResponseWriter, r *http.Request)
	UpdateSchoolHandler(w http.ResponseWriter, r *http.Request)
	DeleteSchoolHandler(w http.ResponseWriter, r *http.Request)
	GetAllUserHandler(w http.ResponseWriter, r *http.Request)

	//mata_pelajaran
	RegisterMapelHandler(w http.ResponseWriter, r *http.Request)
	GetAllMapelHandler(w http.ResponseWriter, r *http.Request)
	UpdateMapelHandler(w http.ResponseWriter, r *http.Request)
	DeleteMapelHandler(w http.ResponseWriter, r *http.Request)

	//materi_header
	CreateMateriHeaderHandler (w http.ResponseWriter, r *http.Request)
	GetAllMateriHeaderHandler(w http.ResponseWriter, r *http.Request)
	DeleteMateriHandler(w http.ResponseWriter, r *http.Request)
	UpdateMateriHeaderHandler(w http.ResponseWriter, r *http.Request)
}

func NewController(c usecase.Base)Controllers{
	return &ctrl{c}

}