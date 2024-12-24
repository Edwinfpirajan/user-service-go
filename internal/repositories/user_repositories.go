package repositories

import (
	"github.com/Edwinfpirajan/user-service-go/internal/core/domain"
	"github.com/Edwinfpirajan/user-service-go/internal/database"
)

// UserRepository repositorio de usuarios
type UserRepository struct {
	db *database.GormDB
}

// NewUserRepository crea una instancia del repositorio de usuarios
func NewUserRepository(db *database.GormDB) *UserRepository {
	return &UserRepository{db}
}

// Obtiene todos los usuarios
func (r *UserRepository) GetAll() (domain.Users, error) {
	var users domain.Users
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Obtiene un usuario por su ID
func (r *UserRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Crea un usuario
func (r *UserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

// Actualiza un usuario
func (r *UserRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

// Elimina un usuario
func (r *UserRepository) Delete(user *domain.User) error {
	return r.db.Delete(user).Error
}
