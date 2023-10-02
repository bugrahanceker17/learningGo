package main

type User struct {
	FirstName string `json:"adi"`
	LastName  string `json:"lastName"`
	Followers []User `json:"followers,omitempty"` // omitempty response'da nil olan değerleri yazmaz
}

func main() {

}

func (u User) FollowerCount() int {
	return len(u.Followers)
}

func (u *User) AddFollower(user User) { // bir struct üzerinde değişiklik yapılacaksa bir pointer receiver tanımlamak gerekir
	if u.Followers == nil {
		u.Followers = []User{}
	}

	u.Followers = append(u.Followers, user)
}
