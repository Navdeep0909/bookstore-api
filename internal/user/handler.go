package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(w http.ResponseWriter, r *http.Request){
	var req SignupRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	fmt.Println(req)
	if err!=nil{
		http.Error(w, "InvalidInput", http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil{
		http.Error(w, "Server error", http.StatusInternalServerError)
	}

	user1 := User{
		Id: uuid.NewString(),
		Email: req.Email,
		Password: string(hashedPassword),
		Role: "Customer", //ByDefault
	}
	fmt.Println(user1)
	id := InsertUser(userCollection, user1)
	if id != nil{
		w.WriteHeader(http.StatusCreated)
		// json.NewEncoder(w).Encode(user)
	}
	w.WriteHeader(http.StatusInternalServerError)

}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	var filter map[string]interface{}
	
	if req.Email != ""{
		filter["email"] = req.Email
	}

	dbUser, err := GetUserByEmail(userCollection, &filter)
	if err != nil{
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(req.Password))
	if err != nil{
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	// JWT token generation
	expiresAt := time.Now().Add(15 * time.Minute)
	claims := &jwt.RegisteredClaims{
		Subject:   dbUser.Id,
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(AuthResponse{AccessToken: signedToken})
}	