package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ahorner2/goApi/server/db"
)

func FetchAllUsers(w http.ResponseWriter, r *http.Request) {
    var userList []User

    connection := db.GetDBConnection()
    defer connection.Close()

    rows, err := connection.Query(`select username, email, passphrase, iv from users`)
    if err != nil {
        http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var user User
        if err := rows.Scan(&user.Username, &user.Email, &user.Passphrase, &user.IV); err != nil {
            http.Error(w, "Failed to scan user", http.StatusInternalServerError)
            return
        }
        userList = append(userList, user)
    }

    if err := rows.Err(); err != nil {
        http.Error(w, "Rows error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html")
    json.NewEncoder(w).Encode(userList)
		
		if len(userList) == 0 {
				w.WriteHeader(http.StatusNoContent)
				return
		}

    // marshal the user slice w/ indents for readability 
    //// just for demo purposes
    jsonData, err := json.MarshalIndent(userList, "", "    ")
    if err != nil {
        http.Error(w, "Failed to encode userList data", http.StatusInternalServerError)
        return 
    }

    // wrap json data within a <pre> tag for "pretty" display
    htmlData := fmt.Sprintf("<pre>%s</pre>", jsonData)

    // write the formatted HTML to the response writer
    w.Write([]byte(htmlData))
}



func InsertUser(user User) error {
	connection := db.GetDBConnection()
	defer connection.Close()

	sql := `
		INSERT INTO users (username, email, passphrase, iv, createdAt, updatedOn)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
	`
	_, err := connection.Exec(
		sql, 
		user.Username, 
		user.Email, 
		user.Passphrase, 
		user.IV,
	)
	if err != nil {
		return err
	}

	return nil 
}