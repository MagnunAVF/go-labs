package main

// The Bad: Forces the caller to check for nil before using the data.
// func (r *Repo) ListUsers() []User {
//     if r.count == 0 {
//         return nil
//     }
//     return r.users
// }

// The Good: Returns an empty slice, allowing the caller to use a range loop
// immediately.
type Repo struct {
	count int
	users []User
}
type User struct{}

func (r *Repo) ListUsers() []User {
	if r.count == 0 {
		return []User{} // or simply let it return the initialized slice
	}
	return r.users
}
