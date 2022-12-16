package model

type (
	Person struct {
		PersonID  int    `json:"person_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address   string `json:"address"`
		City      string `json:"city"`
	}
)
