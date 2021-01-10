package users

type User struct {
	ID          int64  `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Email       string `json:"email,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
}
