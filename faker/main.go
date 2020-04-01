package main

import (
	"fmt"

	"github.com/bxcodec/faker"
)

// Sample ...
type Sample struct {
	ID       int64
	IPV6     string `faker:"ipv6"`
	UserName string `faker:"username"`
}

func main() {
	a := Sample{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)

}
