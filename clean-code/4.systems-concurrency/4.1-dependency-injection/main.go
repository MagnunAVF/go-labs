package main

// The Bad: Hard-coded Dependencies The service creates its own database connection.
// You cannot test this without a real database running.
// type UserService struct {
//     db *sql.DB
// }

// func NewUserService() *UserService {
//     // Hidden dependency! Hard to mock/test.
//     db, _ := sql.Open("postgres", "user=pqtest dbname=pqtest sslmode=verify-full")
//     return &UserService{db: db}
// }

// The Good: Constructor Injection The service is "passive." It is given its
// dependencies, making it easy to swap a real DB for a "Mock" during testing.
type User struct{}
type UserStore interface {
	Save(u *User) error
}

type UserService struct {
	store UserStore
}

// Inversion of Control: pass the dependency in
func NewUserService(store UserStore) *UserService {
	return &UserService{store: store}
}
