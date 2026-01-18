package main

import "fmt"

// The Bad: Mixes high-level business logic with low-level implementation
// (SQL and string formatting).
// func (s *OrderService) Complete(id string) error {
// 	order, _ := s.db.Get(id)
// 	order.Status = "Done"

// 	// Low-level plumbing mixed in
// 	q := "UPDATE orders SET status = 'Done' WHERE id = '" + id + "'"
// 	_, err := s.db.Exec(q)

// 	// Side effect: invisible logging logic
// 	fmt.Printf("LOG: Order %s completed at %v\n", id, time.Now())
// 	return err
// }

// The Good: Delegates technical details to a Repository and stays at a high
// level of abstraction.
type Repository interface {
	FindByID(string) (OrderService, error)
	Save(OrderService) error
}

type OrderService struct {
	repo Repository
}

func (o *OrderService) MarkAsCompleted() {
	fmt.Println("Marking as completed!")
}

func (s *OrderService) CompleteOrder(id string) error {
	order, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("finding order: %w", err)
	}

	order.MarkAsCompleted()

	if err := s.repo.Save(order); err != nil {
		return fmt.Errorf("saving completed order %s: %w", id, err)
	}

	return nil
}
