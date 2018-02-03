package internal

type Article struct {
	Likes int
	Author, Company, Title, Description, Url, UrlToImage, PublishedAt string
}

type Author struct {
	Likes, ArticlesNr int
	Name, Company string
	Articles []Article
}

type Company struct {
	Likes int
	Company, Name, Language, Category, ImgUrl string
}

type Companies struct {
	Companies []Company
}


type User struct {
	Token string
}

func NewUser(token string) *User {
	return &User{Token:token}
}

type UsersInMemory struct {
	Users []User
}

func NewUsersInMemory() *UsersInMemory {
	return &UsersInMemory{}
}

func AddUserToMemory(users *UsersInMemory, user User) {
	users.Users = append(users.Users, user)
}