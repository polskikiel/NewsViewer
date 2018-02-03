package handlers

import (
	"net/http"
	"crypto/rand"
	"fmt"
	"NewsViewer/internal"
	"log"
)

func NewServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{Addr: addr, Handler: handler}
}

func Run(server *http.Server) {
	log.Fatalf("%s", server.ListenAndServe())
}

func NewHandler(users *internal.UsersInMemory) http.Handler {
	h := http.NewServeMux()
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		auth, ok := r.URL.Query()["a"]

		if !ok || !isUserAuthorized(auth[0], users) {
			http.Redirect(w, r, "/register", 301)
		}

		w.Write([]byte("NEWS" + r.Header.Get("VAuth")))

	})
	h.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		token := createUserToken()
		internal.AddUserToMemory(users, *internal.NewUser(token))
		http.Redirect(w, r, "/?a="+token, 301)
	})
	return h
}

func isUserAuthorized(token string, users *internal.UsersInMemory) bool {
	for _, u := range users.Users {
		if token == u.Token {
			return true
		}
	}
	return false
}

func createUserToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
