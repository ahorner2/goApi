package users

import (
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Path[len("/user/"):]

	userIndex := -1 
	for i, u := range users {
		if u.ID == userID {
			userIndex = i
			break
		}
	}

	if userIndex == -1 {
		http.Error(w, "User not found", http.StatusNotFound) 
		return
	}

	// delete users from slice 
	users = append(users[:userIndex], users[userIndex + 1:]...)
	// return success status, 'no content' for delete  
	w.WriteHeader(http.StatusNoContent)

}

