package user

import "gorm.io/gorm"

type Repository interface {
	Create(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(id int) (User, error)
	Update(user User) (User, error)
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
	//TODO implement me
	panic("implement me")
}

func (r *repository) FindByEmail(email string) (User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) FindByID(id int) (User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Update(user User) (User, error) {
	//TODO implement me
	panic("implement me")
}
