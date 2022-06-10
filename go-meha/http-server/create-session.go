package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
)

func main() {
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/logout", Logout)
	http.ListenAndServe(":2000", nil)
}
var (
	key = []byte("random-encription-key")
	sessionStore = sessions.NewCookieStore(key)
)
func Secret(w http.ResponseWriter, r *http.Request) {
	cookieData, _ := sessionStore.Get(r, "cookie-name")
	if cookieData.Values["authenticated"] == false {
		fmt.Fprintf(w, "The User is not Authenticated")
	} else {
		fmt.Fprintf(w, "The User is Authenticated")
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	cookieData, _ := sessionStore.Get(r, "cookie-name")
	cookieData.Values["data"] = true
	cookieData.Save(r, w)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookieData, _ := sessionStore.Get(r, "cookie-name")
	cookieData.Values["authenticated"] = false
	cookieData.Save(r, w)
}