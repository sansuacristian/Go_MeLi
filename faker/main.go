package main

import (
	"fmt"
	"log"

	"github.com/bxcodec/faker"
)

// Sample ...
type Sample struct {
	// ID       int64
	// IPV6     string `faker:"ipv6"`
	// UserName string `faker:"username"`
	Password string `faker:"password"`
}

func main() {
	a := Sample{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)

	c := Sample{
		Password: a.Password,
	}

	if a == c {
		log.Println("se comprobo el codigo")
	} else {
		log.Println("codigo incorrecto")
	}

}
