package handlers

import "loosidAPI/generated"

// ServerWrapper - struct for implementing ServerInterface
type ServerWrapper struct{}

var err500 = generated.Error{0, "Internal Server Error"}
