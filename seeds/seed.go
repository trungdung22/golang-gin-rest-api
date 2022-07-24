package seeds

import (
	"fmt"
	"math/rand"

	"crud-api/models"

	"gorm.io/gorm"
)

func randomInt(min, max int) int {

	return rand.Intn(max-min) + min
}

func seedRoles(db *gorm.DB) {
	var countAdmin int64 = 0
	var countUser int64 = 0
	adminRole := models.Role{Name: "ROLE_ADMIN", Description: "Admin role"}
	userRole := models.Role{Name: "ROLE_USER", Description: "User role"}
	queryAdmin := db.Model(&models.Role{}).Where("name=?", "ROLE_ADMIN")
	queryUser := db.Model(&models.Role{}).Where("name=?", "ROLE_USER")
	queryAdmin.Count(&countAdmin)
	queryUser.Count(&countUser)

	if countAdmin == 0 {
		db.Create(&adminRole)
	} else {
		queryAdmin.First(&adminRole)
	}

	if countUser == 0 {
		db.Create(&userRole)
	}

	var adminUserCount int64 = 0
	db.Model(&models.User{}).Joins("inner join roles on roles.id = users.id").Where("roles.name=?", "ROLE_ADMIN").Count(&adminUserCount)

	fmt.Print("admint user count", adminUserCount)

	if adminUserCount == 0 {
		user := models.User{Username: "admin", Email: "dung.do@wizeline.com", Bio: "Bio description"}
		user.Role = adminRole
		db.Set("gorm:association_autoupdate", false).Create(&user)
		if db.Error != nil {
			print(db.Error)
		}
	}
}

func Seed() {
	db := models.GetDB()
	seedRoles(db)
}
