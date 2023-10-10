package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// extract id from url path 
	vars := mux.Vars(r)
	userID := vars["id"]

	var updatedUser User 

	err := json.NewDecoder(r.Body).Decode(&updatedUser) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

		// validate that id in request body matches the url id 
	if updatedUser.ID != userID {
		http.Error(w, "Mismatched user IDs", http.StatusBadRequest)
		return 
	}

	// find user by id and update
	userIndex := -1 
	for i, u := range users {
		if u.ID == updatedUser.ID {
			userIndex = i
			break
		}
	}
	
	// if not found return error 
	if userIndex == -1 {
		http.Error(w, "User not found", http.StatusBadRequest)
		return 
	}
	
	// update the user
	users[userIndex] = updatedUser 

	// return updated user 
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}