package user

import (
	"errors"
	"fmt"
	"regexp"
	"time"

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

type BankCardsBound struct {
	model.BaseModel
	UserId       int    `json:"user_id" gorm:"column:user_id;" binding:"required"`             // 银行卡号
	CardNumber   string `json:"card_number" gorm:"column:card_number;" binding:"required"`     // 银行卡号
	CardPassword string `json:"card_password" gorm:"column:card_password;" binding:"required"` // 银行卡密码
	BankName     string `json:"bank_name" gorm:"column:bank_name;" binding:"required"`         // 银行名称
	ExpiryDate   string `json:"expiry_date" gorm:"column:expiry_date;" binding:"required"`     // 到期日期，格式为"YYYY-MM-DD"
}

func (c *BankCardsBound) TableName() string {
	return "bank_cards_bound"
}

func (c *BankCardsBound) Valid() bool {
	// 检查UserID是否有效
	if c.UserId <= 0 {
		return false
	}

	// 检查CardNumber是否有效（假设为16位数字）
	cardNumberRegex := regexp.MustCompile(`^\d{16}$`)
	if !cardNumberRegex.MatchString(c.CardNumber) {
		return false
	}

	// 检查CardPassword是否有效（假设为6位数字）
	cardPasswordRegex := regexp.MustCompile(`^\d{6}$`)
	if !cardPasswordRegex.MatchString(c.CardPassword) {
		return false
	}

	// 检查BankName是否为空
	if c.BankName == "" {
		return false
	}

	// 检查ExpiryDate是否是将来的日期
	expiryDate, err := time.Parse("2006-01-02", c.ExpiryDate)
	if err != nil || expiryDate.Before(time.Now()) {
		return false
	}

	// 如果所有检查都通过了，则返回true
	return true
}

func (c *BankCardsBound) Create() error {
	return db.DB.
		Table(c.TableName()).
		Create(c).Error
}

func (c *BankCardsBound) Get() ([]string, error) {
	var cards []BankCardsBound
	var boundCards []string

	// 假设 db 是 *gorm.DB 的已经初始化的实例
	err := db.DB.Table(c.TableName()).Where("user_id = ?", c.UserId).Find(&cards).Error
	if err != nil {
		return nil, err
	}

	// 将查询结果中的卡号提取到 boundCards 切片中
	for _, card := range cards {
		boundCards = append(boundCards, card.CardNumber)
	}
	return boundCards, nil
}

func (c *BankCardsBound) JudgeBackCardBounded() error {
	var cardModel BankCardsBound
	// 修改查询条件，同时根据userId和cardNumber进行筛选
	db.DB.
		Table(c.TableName()).
		Where("user_id = ? AND card_number = ?", c.UserId, c.CardNumber).
		First(&cardModel)
	// 如果找到了记录，说明这张卡已经被绑定
	if cardModel.BaseModel.ID > 0 {
		return errors.New("this user has already bound this card")
	} else {
		return nil
	}
}
