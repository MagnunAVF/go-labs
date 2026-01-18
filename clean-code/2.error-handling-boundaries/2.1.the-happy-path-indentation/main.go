package main

import (
	"errors"
	"fmt"
)

// The Bad: The "Arrow" anti-pattern. The success logic is buried deep inside
// nested if blocks.
// func (s *Store) Buy(itemID string) error {
//     item, err := s.find(itemID)
//     if err == nil {
//         if item.InStock {
//             err := s.charge()
//             if err == nil {
//                 return s.ship()
//             } else {
//                 return err
//             }
//         } else {
//             return errors.New("out of stock")
//         }
//     } else {
//         return err
//     }
// }

// The Good: Uses Guard Clauses to return early, keeping the "Happy Path" left-aligned.
type Store struct{}

func (s *Store) find(string) (Item, error)
func (s *Store) charge() error
func (s *Store) ship() error

type Item struct {
	InStock bool
}

func (s *Store) Buy(itemID string) error {
	item, err := s.find(itemID)
	if err != nil {
		return fmt.Errorf("finding item: %w", err)
	}

	if !item.InStock {
		return errors.New("item out of stock")
	}

	if err := s.charge(); err != nil {
		return fmt.Errorf("charging customer: %w", err)
	}

	return s.ship()
}
