package domain

type BaseService interface {
	// admin
	QueryByEmail(email string) (Admin, error)
	StoreAdmin(admin *Admin) error
	Find(email string, adminid int) (*Admin, error)
	GetAllAdmin(adminid int) ([]AdminInfo, error)
	GetAllTransaction(transactionid int) ([]TransactionInfo, error)
	UpdateAdmin(id string, admin *UpdateAdminInfo) error
	QueryByIdAdmin(idAdmin int) (AdminInfo, error)
	DeleteAdmin(adminid int) error

	// user
	InsertSchool(school *School) error
	QueryByKodeSekolah(kode_sekolah string) (School, error)
	QueryByIdSchool(idSchool int) (UpdateSchoolInfo, error)
	UpdateSchool(id string, school *UpdateSchoolInfo) error
	DeleteSchool(school int) error
	GetAllUser(userid int) ([]User, error)
	/*
	ChangePassword(admin *domain.Admin) error
	*/

	// mata_pelajaran
	QueryByMapel(mapel string) (Mata_Pelajaran, error)
	StoreMapel(mapel *Mata_Pelajaran) error
	GetAllMapel(mapel int) ([]Mata_Pelajaran, error)
	QueryByIdTingkat(id_tingkat int) (UpdateInfoMapel, error)
	UpdateMapel(id string, mapel *UpdateInfoMapel) error
	QueryByIdMatpel(mapel int) (Mata_Pelajaran, error)
	DeleteMapel(mapel int) error

	// materi_header
	CheckMapel(mapel int) (Materi_Header, error)
	InsertMateri(materi *Materi_Header) error
	GetIdMaHeader (materiHeader int) (UpdateMateriHeader, error)
	QueryByIdMateri(materi_headerid int) (Materi_Header, error)
	DeleteMateriHeader (materi_headerId int) error
	UpdateMateriHeader(id string, materiHeader *UpdateMateriHeader) error
	GetAllMateri(materiId int) ([]Materi_Header, error)

}

type Profile interface {
	SetAdminID(id int)
}

type Roles interface {
	Model() Profile
	FromJSON(data []byte)

}
