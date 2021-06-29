package adapter

import (
	"fmt"
	"gorm.io/gorm"
	"lawenconTest/domain"
)

const (
	deleteMessage = "DeleteAdmin"
)

type dataRepository struct {
	db *gorm.DB
}

func NewDataRepository(db *gorm.DB) domain.BaseService {
	return &dataRepository{db}
}

func (u *dataRepository) QueryByEmail(email string) (domain.Admin, error) {
	var result domain.Admin
	err := u.db.Raw("SELECT email FROM admins WHERE email = ?", email).Scan(&result).Error

	return result, err
}

func (u *dataRepository)StoreAdmin(admin *domain.Admin) error{
	err := u.db.Exec(`INSERT INTO admins(name, email, password, roles ) VALUES (?, ?, ?, ?)`,admin.Name, admin.Email, admin.Password, admin.Roles).Error
	return err
}

func (u *dataRepository) Find(email string, adminid int) (*domain.Admin, error) {
	var us domain.Admin

	res := u.db.First(&us, qCompose(email, adminid))

	if res.Error != nil {
		return nil, res.Error
	}

	return &us, nil
}

func qCompose(email string, id int) string {
	if len(email) == 0 {
		return fmt.Sprintf(`id = "%v"`, id)
	}
	return fmt.Sprintf(`email = "%v"`, email)
}

func (u *dataRepository) GetAllAdmin(adminid int) ([]domain.AdminInfo, error) {
	var result []domain.AdminInfo
	err := u.db.Raw("SELECT id, name, email, roles FROM admins  ").Scan(&result).Error
	/*WHERE deleted = '0'*/
	return result, err
}

func (u *dataRepository) GetAllTransaction(transactionid int) ([]domain.TransactionInfo, error) {
	var resultTransaction []domain.TransactionInfo
	err := u.db.Raw("SELECT payment_status, trx_code, payment_method FROM transactions ").Scan(&resultTransaction).Error

	return resultTransaction, err
}

func (u *dataRepository) UpdateAdmin(id string, admin *domain.UpdateAdminInfo) error {
	err := u.db.Exec(`UPDATE admins SET name = ?, email = ?, roles = ? WHERE id = ?`, admin.Name, admin.Email, admin.Roles, id).Error

	fmt.Println(id)
	fmt.Println(err)

	return nil
}

func (u *dataRepository)QueryByIdAdmin(idAdmin int) (domain.AdminInfo, error){
	var result domain.AdminInfo
	err := u.db.Raw("SELECT id FROM admins WHERE id = ?", idAdmin).Scan(&result).Error
	return result, err
}

func (u *dataRepository) DeleteAdmin(adminid int) error {
	err := u.db.Exec(`UPDATE admins SET deleted = 1 WHERE id = ?`, adminid).Error

	if err != nil {
		return err
	}
	return nil
}