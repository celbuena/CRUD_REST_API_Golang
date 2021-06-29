package adapter

import (
	"fmt"
	"lawenconTest/domain"
)

func (u *dataRepository) CheckMapel(mapel int) (domain.Materi_Header, error) {
	var result domain.Materi_Header
	err := u.db.Raw("SELECT id_mapel, FROM materi_header where id_mapel = ?", mapel).Scan(&result).Error
	return result, err
}

func (u *dataRepository)InsertMateri(materi *domain.Materi_Header) error{
	err := u.db.Exec(`INSERT INTO materi_header(id_mapel, chapter, title, label, materi, detail, status_materi) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		materi.IdMapel, materi.Chapter, materi.Title, materi.Label, materi.Materi, materi.Detail, materi.StatusMateri).Error
	return err
}

/*func (u *dataRepository)StoreMateriHeader(materi_header *domain.Materi_Header) error {
	err := u.db.Exec(`INSERT INTO materi_header (id_mapel, chapter, title, label, materi, detail, status_materi) VALUES (?, ?, ?, ?, ?, ?, ?)
		`,materi_header.IdMapel, materi_header.Chapter, materi_header.Title, materi_header.Label ,materi_header.Materi, materi_header.Detail,
		materi_header.StatusMateri).Error

	return err
}*/

func (u *dataRepository)GetIdMaHeader (materiHeader int) (domain.UpdateMateriHeader, error){
	var result domain.UpdateMateriHeader
	err := u.db.Raw("SELECT id FROM materi_header WHERE id = ?", materiHeader).Scan(&result).Error
	return result, err
}

func (u *dataRepository) QueryByIdMateri(materi_headerid int) (domain.Materi_Header, error){
	var result domain.Materi_Header
	err := u.db.Raw("SELECT id FROM materi_header WHERE id = ?", materi_headerid).Scan(&result).Error
	fmt.Println(err)
	return result, err
}

func (u *dataRepository)DeleteMateriHeader (materi_headerId int) error {
	err := u.db.Exec(`UPDATE materi_header SET deleted = 1 WHERE id = ?`, materi_headerId).Error
	if err != nil{
		return err
	}
	return err
}

func (u *dataRepository)UpdateMateriHeader(id string, materiHeader *domain.UpdateMateriHeader) error{
	err := u.db.Exec(
		`UPDATE materi_header SET chapter = ?, title = ?, label = ?, materi = ?, detail = ?, status_materi = ? WHERE id = ?`,
		materiHeader.Chapter, materiHeader.Title, materiHeader.Label, materiHeader.Materi, materiHeader.Detail,
		materiHeader.StatusMateri, id).Error
	fmt.Println(materiHeader.Id)
	fmt.Println(err)
	return nil
}

func (u *dataRepository) GetAllMateri(materiId int) ([]domain.Materi_Header, error) {
	var result []domain.Materi_Header
	err := u.db.Raw("SELECT * FROM materi_header").Scan(&result).Error

	return result, err
}