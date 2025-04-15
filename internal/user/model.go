package user

type User struct{
	Id 			string	`json:"id"`
	Email 		string	`json:"email"`
	Password 	string	`json:"password,omitempty"`
	Role 		string	`json:"role"`
}

type SignupRequest struct{
	Email 		string 	`json:"email"`
	Password 	string 	`json:"password"`
}

type LoginRequest struct{
	Email 		string 	`json:"email"`
	Password 	string 	`json:"password"`
}

type AuthResponse struct{
	AccessToken string `json:"access_token"`
}