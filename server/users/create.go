package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ahorner2/goApi/lib/utils"
)

type User struct {
	ID       			 string `json:"id"`
	Username 			 string `json:"username"`
	Email    			 string `json:"email"`
	Passphrase 		 string `json:"passphrase"`
	IV 	   			   string `json:"iv"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedOn 		 time.Time `json:"updatedOn"`
}

// mock user data 
var users []User 

func CreateNewUser(w http.ResponseWriter, r *http.Request){ 
	var newUser User 

	err := json.NewDecoder(r.Body).Decode(&newUser) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	envStr := os.Getenv("ENCRYPTION_KEY")
		encryptedPassphrase, generatedIV, err := utils.EncryptAES256([]byte(newUser.Passphrase), envStr)
		if err != nil {
			http.Error(w, "Error encrypting passphrase", http.StatusInternalServerError)
			return
	}
	newUser.Passphrase = encryptedPassphrase
	newUser.IV = generatedIV

	err = InsertUser(newUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to insert user into DB: %v", err), http.StatusInternalServerError)
		return
	}

	newUser.ID = strconv.Itoa(len(users) + 1) 
	users = append(users, newUser)

	w.Header().Set("Content-Type", "application/json") 
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}


