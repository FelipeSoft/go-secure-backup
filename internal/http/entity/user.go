package entity

type User struct {
	Id       string
	Username string
	Email    string
	Password string

	// 0 -> standard; 1 -> administrator
	Role int
}

func (u *User) newUser(id string, username string, email string, password string, role int) *User {
	return &User{
		Id: id,
		Username: username,
		Email: email,
		Password: password,
		Role: role,
	}
}
