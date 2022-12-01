package query

import (
	// "fmt"

	"github.com/fajarhdev/go-mycash/initializers"
	"github.com/fajarhdev/go-mycash/models"
)

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]models.User) {
	initializers.DB.Find(&user)
	return
}

//CreateUser ... Insert New data
func CreateUser(user models.User) {
	// if err = initializers.DB.Create(&user).Error; err != nil {
	// 	return err
	// }
	initializers.DB.Create(&user)
	// fmt.Println(result)
	// fmt.Printf(result.Error)
	return
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *models.User, id string) {
	initializers.DB.Where("id = ?", id).First(&user)
	return
}

//UpdateUser ... Update user
func UpdateUser(user *models.User, update models.User ) {
	initializers.DB.Model(&user).Updates(&update)
	// initializers.DB.Save(user)
	return
}

//DeleteUser ... Delete user
func DeleteUser(user *models.User, id string) (err error) {
	initializers.DB.Delete(&user, id)
	return
}