package main

import (
	"fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Category struct {
	ID		int
	Nama	string
}

type User struct {
	ID		int
	Nama	string
	Pekerjaan	string
	Password	string
}

func main() {
  	db, err := gorm.Open("postgres", "host=127.0.0.1 user=postgres dbname=kejarmimpi sslmode=disable password=postgresreza")
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Connection Success")		
	}
//	db.CreateTable(&User{})
//	db.DropTable(&User{})

	var users []User = []User{
	  User{ID: 1, Nama: "Reza Andriyunanto", Pekerjaan: "Web Developer", Password: "060195"},
	  User{ID: 2, Nama: "Ervian", Pekerjaan: "JS Developer", Password: "060195"},
	  User{ID: 3, Nama: "Naufal", Pekerjaan: "Android Developer", Password: "060195"},
	}

	for _, user := range users {
	  db.Create(&user)
	}
	
	defer db.Close()
}