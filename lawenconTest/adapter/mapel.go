package adapter

import (
	"fmt"
	"lawenconTest/domain"
)

func (u *dataRepository) QueryByMapel(mapel string) (domain.Mata_Pelajaran, error) {
	var result domain.Mata_Pelajaran
	err := u.db.Raw("SELECT mapel, FROM mata_pelajaran where mapel = ?", mapel).Scan(&result).Error
	return result, err
}

func (u *dataRepository)StoreMapel(mapel *domain.Mata_Pelajaran) error{
	err := u.db.Exec(`INSERT INTO mata_pelajaran(id_tingkat, id_parent_mapel, icon, mapel ) VALUES (?, ?, ?, ?)`,mapel.IdTingkat, mapel.IdParentMapel, mapel.Icon, mapel.Mapel).Error
	return err
}

func (u *dataRepository) GetAllMapel(mapel int) ([]domain.Mata_Pelajaran, error){
	var result []domain.Mata_Pelajaran
	err := u.db.Raw("SELECT * FROM mata_pelajaran").Scan(&result).Error

	return result, err
}

func (u *dataRepository)QueryByIdTingkat(id_tingkat int) (domain.UpdateInfoMapel, error){
	var result domain.UpdateInfoMapel
	err := u.db.Raw("SELECT id_tingkat FROM mata_pelajaran WHERE id_tingkat = ?", id_tingkat).Scan(&result).Error
	return result, err
}

func (u *dataRepository)UpdateMapel(id string, mapel *domain.UpdateInfoMapel) error {
	err := u.db.Exec(`UPDATE mata_pelajaran SET id_tingkat = ?, id_parent_mapel = ?, icon = ?,  mapel = ? WHERE id = ?`,mapel.IdTingkat, mapel.IdParentMapel,
		mapel.Icon, mapel.Mapel, id).Error
	fmt.Println(err)
	fmt.Println(mapel.IdTingkat)

	return nil
}

func (u *dataRepository)QueryByIdMatpel(mapel int) (domain.Mata_Pelajaran, error){
	var result domain.Mata_Pelajaran
	err := u.db.Raw("SELECT id FROM mata_pelajaran WHERE id = ?", mapel).Scan(&result).Error
	return result, err
}

func (u *dataRepository) DeleteMapel(mapel int) error {
	err := u.db.Exec(`UPDATE mata_pelajaran SET deleted = 1 WHERE id = ?`, mapel).Error
	if err != nil {
		return err
	}
	return nil
}


