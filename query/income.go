package query

import (

	"github.com/fajarhdev/go-mycash/initializers"
	"github.com/fajarhdev/go-mycash/models"
)

// the function for query adding income transaction to database
func AddIncome(income *models.Income) {
	initializers.DB.Omit("User").Create(&income)

}

// the function for query fetch income transaction with specific id user
func GetAllIncome(income *[]models.Income, id int) {
	initializers.DB.Where("user_id = ?", id).Joins("User").Find(&income)
}

// the function for query update income transaction with specific id user
func UpdateIncome(income *models.Income, updateIncome models.Income, id int) {
	// initializers.DB.Model(&income).Update(&updateIncome)
	initializers.DB.Model(&income).Where("id = ?", id).Updates(&updateIncome)
}

// the function for query delete income transaction with specific id user
func DeleteIncome(income *models.Income, id int) (err error){
	initializers.DB.Delete(&income, id)
	return
}