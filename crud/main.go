package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{
		Id:   1,
		Name: "John Doe",
		Age:  30,
	},
	{
		Id:   2,
		Name: "Jane Smith",
		Age:  25,
	},
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("POST /create", createHandler)
	mux.HandleFunc("GET /users", getUsersHandler)
	mux.HandleFunc("GET /users/{id}", getSingleUserHandler)
	mux.HandleFunc("PUT /users/{id}", updateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", deleteUserHandler)

	fmt.Println("Server is running port 5000")
	err := http.ListenAndServe(":5000", mux)

	if err != nil {
		fmt.Println("Server error: ", err)
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to go server")
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Health check passed")
}
func createHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	// w.WriteHeader(405)
	// 	fmt.Fprintln(w , "Methode not allowed");
	// 	return
	// }
	var err error

	var newUser User
	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(403)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "Create endpoint", newUser)
}
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	// users,_ := json.Marshal(users)
	// w.Write(users)

	json.NewEncoder(w).Encode(users)

	fmt.Fprintln(w, "Get Users endpoint")
}

func getSingleUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	//   fmt.Printf("the value of the id is %d", id)

	for _, user := range users {
		if user.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedUser User

	err = json.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil {
		w.WriteHeader(403)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for idx , user := range users {
		user.Id == id{
			updatedUser.Id == id
			updatedUser.Id = updatedUser

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
    
	http.Error(w, "User not found", http.StatusNotFound)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"})
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}