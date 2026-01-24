package main

// 5. Dependency Inversion Principle (DIP)
// "High-level modules should not depend on low-level modules. Both should
// depend on abstractions."

// BAD
// The Service is tightly coupled to a specific MySQLDatabase. You can't easily
// swap it for PostgreSQL or a Mock for testing.
// type MySQLDatabase struct{}

// type Service struct {
// 	db MySQLDatabase
// }

// func NewService() *Service {
// 	return &Service{db: MySQLDatabase{}}
// }

// GOOD
// The Service depends on an interface (Repository). You can inject any database
// that satisfies that interface.
type Repository interface {
	Save(data string)
}

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}
