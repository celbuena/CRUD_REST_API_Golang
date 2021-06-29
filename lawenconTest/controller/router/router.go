package router

import (
	"github.com/go-chi/chi"
	"lawenconTest/controller/apihandler"
	"net/http"
	/*"github.com/valyala/fasthttp"*/
	"lawenconTest/middleware"
)



type route struct {
	api apihandler.Controllers
}

type Route interface {
	Mux() http.Handler
}

func NewRoute(api apihandler.Controllers) Route {
	return &route{api}
}


func(c *route) Mux() http.Handler{

	router := chi.NewRouter()
	//admin
	router.Group(func(r chi.Router) {
		r.Post("/admins/login",c.api.LoginHandler)
		r.Put("/admins/update",c.api.UpdateAdminHandler)
	})

	router.Group(func(r chi.Router) {
		r.Use(middleware.JWTValidation)
		r.Post("/admins/signup",c.api.SignUpHandler)
		r.Get("/admins/get",c.api.GetAllAdminHandler)
		r.Get("/admins/get-transactions",c.api.GetAllTransactionHandler)
		r.Put("/admins/delete-admin",c.api.DeleteAdminHandler)

		//user
		r.Post("/users/register-school",c.api.RegisterSchoolHandler)
		r.Put("/users/update-school",c.api.UpdateSchoolHandler)
		r.Put("/users/delete-school",c.api.DeleteSchoolHandler)
		r.Get("/users/get-users",c.api.GetAllUserHandler)

		//mata_pelajaran
		r.Post("/mapel/register-mapel",c.api.RegisterMapelHandler)
		r.Get("/mapel/get-mapel",c.api.GetAllMapelHandler)
		r.Put("/mapel/update-mapel",c.api.UpdateMapelHandler)
		r.Put("/mapel/delete-mapel",c.api.DeleteMapelHandler)

		// materi_header
		r.Post("/materi/register-materi",c.api.CreateMateriHeaderHandler)
		r.Get("/materi/get-materi_header",c.api.GetAllMateriHeaderHandler)
		r.Put("/materi/delete-materi_header",c.api.DeleteMateriHandler)
		r.Put("/materi/update-materi_header",c.api.UpdateMateriHeaderHandler)
	})
	return router
}