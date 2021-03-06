package migrations

import (
	"github.com/fantasy9830/go-boilerplate/database"
	"github.com/fantasy9830/go-boilerplate/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db = database.GetDB()
}

// Run run the migrations.
func Run() {
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})
}

// Reverse reverse the migrations.
func Reverse() {
	db.DropTableIfExists(&models.User{}, &models.Role{}, &models.Permission{}, "user_roles", "user_permissions", "role_permissions")
}
