package entities

type Member struct {
	ID                string `json:"id"`
	UserName          string `json:"username"`
	FirstName         string `json:"firstname"`
	LastName          string `json:"lastname"`
	ProfilePictureURL string `json:"profilepictureurl"`
}
