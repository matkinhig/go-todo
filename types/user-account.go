package types

type UserAccount struct {
	USER_NAME string `json: "username"`
	PASSWORD  string `json: "password"`
	AGE       string `json: "phonenumber"`
	EMAIL     string `json: "email"`
}
