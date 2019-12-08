package domains

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	ImageUrl string `json:"imageUrl"`
}
