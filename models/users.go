package models // define the namespace

// Users - Model for the uses table
// defines how this property will be named in the JSON object.
// defines datatype and size
type Users struct {
	UserId   int    `json:"user_id" orm:"auto"` //auto increment
	Email    string `json:"email" orm:"size(128)"`
	Password string `json:"password" orm:"size(64)"`
	UserName string `json:"user_name" orm:"size(32)"`
}