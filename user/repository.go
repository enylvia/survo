package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(id int) (User, error)
	Update(user User) (User, error)
	CreateAttribut(attribut Attribut)
	UpdateAttribut(attribut Attribut)
	FindAll() ([]User, error)
	Delete(id int) (error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByID(id int) (User, error) {
	var user User
	//var attrib Attribut
	err := r.db.Preload("Attribut").Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) CreateAttribut(attribut Attribut) {
	err := r.db.Create(&attribut).Error
	if err != nil {
		return
	}
	return
}

func (r *repository) UpdateAttribut(attribut Attribut) {
	err := r.db.Save(&attribut).Error
	if err != nil {
		return
	}
	return
}

func (r *repository) FindAll() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error
	if err != nil {
		return users,err
	}
	return users,nil
}

func (r *repository)Delete(id int)(error){
	var user User
	err := r.db.Unscoped().Delete(&user, "id = ?",id).Error
	if err != nil {
		return err
	}
	return nil
}