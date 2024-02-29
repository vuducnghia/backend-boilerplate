package models

type User struct {
	BaseModelUUID
	Username    string `json:"username" binding:"required"`
	Password    string `json:"-"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	BaseModelAudit
	BaseModelSoftDelete
}
type Users []User

func (u *User) Create() error {
	return db.Create(u).Error
}

func (u *Users) GetAll() error {
	return db.Find(u).Error
}

func (u *User) GetById() error {
	return db.Model(u).First(u).Error
}

func (u *User) Update() error {
	return db.Model(u).Updates(u).Error
}
