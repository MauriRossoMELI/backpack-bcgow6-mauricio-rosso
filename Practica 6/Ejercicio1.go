package main

func main() {
	var user User = User{Name: "Pablo", Surname: "Diaz", Age: 28, Email: "asd@asd.com", Password: "12345"}
	println("Usuario del comienzo: ", user.Name, " ", user.Surname, " ", user.Age, " ", user.Email, " ", user.Password)
	changeName("newName", "newSurname", &user)
	changeAge(33, &user)
	changeEmail("newEmail@asd.com", &user)
	changePassword("54321", &user)
	println("Usuario del final: ", user.Name, " ", user.Surname, " ", user.Age, " ", user.Email, " ", user.Password)
}

type User struct {
	Name     string
	Surname  string
	Age      int
	Email    string
	Password string
}

func changeName(name string, surname string, user *User) {
	*&user.Name = name
	*&user.Surname = surname
}

func changeAge(age int, user *User) {
	*&user.Age = age
}

func changeEmail(email string, user *User) {
	*&user.Email = email
}

func changePassword(pass string, user *User) {
	*&user.Password = pass
}
