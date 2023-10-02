package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	FirstName string `json:"adi"`
	LastName  string `json:"lastName"`
	Followers []User `json:"followers,omitempty"` // omitempty response'da nil olan değerleri yazmaz
}

type Like struct {
	Date time.Time
}

func main() {
	u := User{
		FirstName: "Bugrahan",
		LastName:  "Ceker",
		Followers: []User{
			{
				FirstName: "Alex",
				LastName:  "Telles",
				Followers: nil,
			},
			{
				FirstName: "Mauro",
				LastName:  "Icardi",
				Followers: nil,
			},
		},
	}

	val, _ := json.Marshal(u)
	fmt.Println(string(val))

}
