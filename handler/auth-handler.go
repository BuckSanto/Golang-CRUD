package handler

import "net/http"

const USERNAME = "belajar-golang"
const PASSWORD = "Password123"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`Authentication Failed`))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte(`Invalid Username or Password`))
		return false
	}

	return true
}
