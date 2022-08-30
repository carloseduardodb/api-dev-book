package models

type ChangePassword struct {
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}
