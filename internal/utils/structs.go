package utils

type Registeruser struct {
	RegisterData string
}

type Data struct {
	UserId 	           string
	NickName           string `json:"name"`
	Email              string
	MobileNumber       string
	ContactPermission  bool
	UserType           string
}
