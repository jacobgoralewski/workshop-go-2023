package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  float64
}

func NewUser(name string, age float64, priv string) *User {
	return &User{Name: name, Age: age, priv: priv}
}

func (u User) String() string {
	return fmt.Sprintf("%v", u)
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (*User) Hello() {
	fmt.Println("Hello")
}

func main() {
	u := User{
		Name: "Jacob",
		Age:  24,
		priv: "asdadas",
	}
	//u.SetName("Andrzej")
	fmt.Println(u)
}
