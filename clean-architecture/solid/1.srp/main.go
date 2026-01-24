package main

import (
	"fmt"
	"os"
)

// 1. Single Responsibility Principle (SRP)
// "A class (or module/function) should have one, and only one, reason to change."
// this means a struct should focus on a single piece of functionality.

// BAD
// Here, the User struct handles both user data and saving that data to a file.
// If the file format changes, the User struct has to change.
// type User struct {
// 	Name  string
// 	Email string
// }
// func (u *User) SaveToFile() {
// 	content := fmt.Sprintf("Name: %s, Email: %s", u.Name, u.Email)
// 	_ = os.WriteFile("user.txt", []byte(content), 0644)
// }

// GOOD
// Separate the data from the persistence logic. Now, UserRepository handles
// saving, and User just holds data.
type User struct {
	Name  string
	Email string
}

type UserRepository struct{}

func (r *UserRepository) Save(u *User) {
	content := fmt.Sprintf("Name: %s, Email: %s", u.Name, u.Email)
	_ = os.WriteFile("user.txt", []byte(content), 0644)
}
