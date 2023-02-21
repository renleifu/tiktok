package service

import "testing"

func TestUser(t *testing.T) {
	var name = "'fyt"
	Register(name, "fyt")
	Login(name, "fyt")
}
