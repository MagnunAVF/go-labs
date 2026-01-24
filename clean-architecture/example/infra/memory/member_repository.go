package memory

import (
	"ca-example/internal/domain"
	"errors"
	"sync"
)

type InMemMemberRepository struct {
	mu      sync.RWMutex
	members map[string]*domain.Member
}

func NewInMemMemberRepository() *InMemMemberRepository {
	return &InMemMemberRepository{
		members: make(map[string]*domain.Member),
	}
}

func (r *InMemMemberRepository) Get(id string) (*domain.Member, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	member, ok := r.members[id]
	if !ok {
		return nil, errors.New("member not found")
	}
	return member, nil
}

func (r *InMemMemberRepository) Save(m *domain.Member) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.members[m.ID] = m
	return nil
}
