package customer

type Model struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string  `json:"email"`
	Status string `json:"status"`
}