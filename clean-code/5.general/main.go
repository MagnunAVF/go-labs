package main

import "fmt"

// BAD:
// 1. Hard-coded dependencies (cannot test without DB/API).
// 2. Logic-heavy "Arrow" indentation.
// 3. Side effects (logging and formatting inside business logic).
// 4. Manual concurrency management that might leak.
// type OrderService struct {
// 	db *sql.DB
// }
// func (s *OrderService) Process(id string, isPriority bool) {
// 	// 1. Data Access & Error Handling mixed
// 	row := s.db.QueryRow("SELECT status, email FROM orders WHERE id = ?", id)
// 	var status, email string
// 	if err := row.Scan(&status, &email); err != nil {
// 		log.Println("Error!!", err)
// 		return
// 	}

// 	if status == "PENDING" {
// 		// 2. Business logic buried in nesting
// 		if isPriority {
// 			// 3. Low-level implementation detail (API call)
// 			resp, err := http.Post("https://shipping.api/priority", "application/json", nil)
// 			if err == nil && resp.StatusCode == 200 {
// 				s.db.Exec("UPDATE orders SET status = 'SHIPPED' WHERE id = ?", id)
// 				// 4. Concurrency used without control
// 				go func() {
// 					smtp.SendMail("smtp.host.com", nil, "bot@co.com", []string{email}, []byte("Shipped!"))
// 				}()
// 			}
// 		} else {
// 			// Duplicate logic for standard shipping...
// 		}
// 	}
// }

// GOOD
// Define the Domain (The "What")
type Order struct {
	ID     string
	Status string
	Email  string
}

// Interfaces allow us to swap implementations
type OrderRepository interface {
	FindByID(id string) (*Order, error)
	Update(o *Order) error
}

type Shipper interface {
	Ship(o *Order) error
}

type Notifier interface {
	Notify(email string, msg string)
}

// The Clean Service (The "How")
type OrderProcessor struct {
	repo     OrderRepository
	shipper  Shipper
	notifier Notifier
}

// NewOrderProcessor uses Dependency Injection
func NewOrderProcessor(r OrderRepository, s Shipper, n Notifier) *OrderProcessor {
	return &OrderProcessor{repo: r, shipper: s, notifier: n}
}

func (p *OrderProcessor) ProcessOrder(id string) error {
	// 1. Guard Clauses (Chapter 7)
	order, err := p.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("finding order %s: %w", id, err)
	}

	if order.Status != "PENDING" {
		return nil // Nothing to do, return early
	}

	// 2. Delegate to specialized objects
	if err := p.shipper.Ship(order); err != nil {
		return fmt.Errorf("shipping process: %w", err)
	}

	order.Status = "SHIPPED"
	if err := p.repo.Update(order); err != nil {
		return fmt.Errorf("updating order status: %w", err)
	}

	// 3. Decoupled Communication
	p.notifier.Notify(order.Email, "Your order has shipped!")

	return nil
}
