package main

import(
  "net/http"
)

type Route struct {
  Name string
  Method string
  Pattern string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
  Route{
    "Index",
    "GET",
    "/",
    Index,
  },
  Route{
    "Issue",
    "POST",
    "/issue",
    IssueJWTToken,
  },
  Route{
    "Validate",
    "POST",
    "/validate",
    ValidateJWTToken,
  },
}