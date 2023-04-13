package model

type User struct {
	Gender         string
	Title          string
	First          string
	Last           string
	Email          string
	Date           string
	Age            int
	RegisteredDate string
	RegisteredAge  int
	Phone          string
	Cell           string
	IdName         string
	IdValue        string
	Nat            string
}

type Location struct {
	Number      int
	Name        string
	City        string
	State       string
	Country     string
	Postcode    any
	Latitude    string
	Longitude   string
	Offset      string
	Description string
}

type Login struct {
	Uuid     string
	Username string
	Password string
	Salt     string
	MD5      string
	SHA1     string
	SHA256   string
}

type Picture struct {
	Large     string
	Medium    string
	Thumbnail string
}
