package user

import (
	"errors"
	"fmt"

	db "backend/database"

	"backend/model"
)

type CommonUserModel struct {
	model.BaseModel
	Phone    string `json:"phone" gorm:"column:phone;" binding:"required"`
	Password string `json:"password" gorm:"column:password;" binding:"required"`
}

func (u *CommonUserModel) TableName() string {
	return "users"
}

func (u *CommonUserModel) Create() error {
	return db.DB.
		Table(u.TableName()).
		Create(u).Error
}

func (u *CommonUserModel) Save() (err error) {
	return db.DB.
		Table(u.TableName()).
		Save(u).Error
}

func GetCommonUserByPhoneAndPassword(phone string, password string) (*CommonUserModel, error) {
	u := &CommonUserModel{}
	d := db.DB.
		Table(u.TableName()).
		Where("phone = ? AND password = ?", phone, password).First(u)
	return u, d.Error
}

func UpdateUser(phone string, password string, newpassword string) (*CommonUserModel, error) {
	u := &CommonUserModel{}

	d := db.DB.
		Table(u.TableName()).
		Where("phone = ? AND password = ?", phone, password).
		Update("password", newpassword)
	return u, d.Error
}

func GetCommonUserById(id int) (*CommonUserModel, error) {
	u := &CommonUserModel{}
	d := db.DB.
		Table(u.TableName()).
		Where("id = ?  ?", id).First(u)
	return u, d.Error
}

// JudgeCommonUserPhoneExist 通过邮箱来判断手机号是非被注册
func JudgeCommonUserPhoneExist(phone string) error {
	var userModel CommonUserModel
	db.DB.
		Table("users").
		Where("phone = ?", phone).First(&userModel)
	if userModel.BaseModel.ID > 0 {
		fmt.Println("this emial has been registered")
		return errors.New("this phone has been registered")
	} else {
		return nil
	}
}
