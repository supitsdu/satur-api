package model

// Account represents an account entity
type AccountPersonalData struct {
	Username     string `json:"username"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	BirthDate    string `json:"birthdate"`
	EmailAddress string `json:"emailaddress"`
}
