package main

import(
  "fmt"
  "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Welcome!")
}

func IssueJWTToken(w http.ResponseWriter, r *http.Request) {
  payload := r.FormValue("payload")
  token, err := CreateToken(payload)

  if err != nil {
    panic(err)
  }
  fmt.Fprintln(w, token)
}

func ValidateJWTToken(w http.ResponseWriter, r *http.Request) {
  token := r.FormValue("token")
  status, result := VerifyToken(token)
  if status == true {
    fmt.Fprintln(w, result)
  } else {
    w.WriteHeader(401)
    fmt.Fprintln(w, result)
  }
}