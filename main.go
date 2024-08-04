package main

import (
	"fmt"

	domain "github.com/waterkemper/my-finance/my-finance-domain/domain/entity"
)

func main() {
	println("My finance app.")
	user := domain.UserBuilder{}.New().Email("rafaelwaterkemper@gmail.com").Build()
	fmt.Println(user)
	user2 := domain.User{}
	user2.SetId(123)
	fmt.Println(user2)
}
