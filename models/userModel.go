package Models

// This model file is to select the table of database
// Example: if we want to select "other" table, then we have to create new model for "other" table

type User struct {
	// Write your column table here
	// ALERT! To always write your column table name with Uppercase letters
	// even the name of the column is not Uppercase letters
	Id    uint   `json:"id"`
	Uname string `json:"uname"`
	Pass  string `json:"pass"`
}

func (b *User) TableName() string {
	// Input table name here, if the table doesn't exist, it will create new one
	return "user_data_new"
}
