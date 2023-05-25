package models

type AllUsers struct {
	Usercd      string `json:"usercd"`
	User_f_Name string `json:"user_f_name"`
	User_l_Name string `json:"user_l_name"`
	Password    string `json:"password"`
}
