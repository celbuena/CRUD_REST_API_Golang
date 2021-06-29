package adapter

import (
	"fmt"
	"lawenconTest/domain"

)

func (u *dataRepository)InsertSchool(school *domain.School) error{
	err := u.db.Exec(`INSERT INTO schools(school_name, school_code, created_by) VALUES (?, ?, ?)`, school.SchoolName, school.SchoolCode, school.CreatedBy).Error
	return err
}

func (u *dataRepository) QueryByKodeSekolah(kode_sekolah string) (domain.School, error) {
	var result domain.School
	err := u.db.Raw(`select id, school_name, school_code where school_code = ?`, kode_sekolah).Scan(&result).Error
	return result, err
}

func (u *dataRepository)QueryByIdSchool(idSchool int) (domain.UpdateSchoolInfo, error){
	var result domain.UpdateSchoolInfo
	err := u.db.Raw(`SELECT id from schools WHERE id = ?`, idSchool).Scan(&result).Error
	return result, err
}

func (u *dataRepository) UpdateSchool(id string, school *domain.UpdateSchoolInfo) error {
	err := u.db.Exec(`UPDATE schools SET school_name = ?, school_code = ? WHERE id = ?`, school.SchoolName, school.SchoolCode, id).Error
	fmt.Println(id)
	fmt.Println(err)
	return nil
}

func (u *dataRepository) DeleteSchool(school int) error {
	err := u.db.Exec(`UPDATE schools SET deleted = 1 WHERE id = ?`, school).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *dataRepository) GetAllUser(userid int) ([]domain.User, error) {
	var result []domain.User
	err := u.db.Raw("SELECT id, name, email, password, roles, deleted = '0' FROM users").Scan(&result).Error

	return result, err
}





