package domain

type Member struct {
	ID     string
	Points int
}

func (m *Member) AddPoints(amount int) {
	m.Points += amount
}

type MemberRepositiry interface {
	Get(id string) (*Member, error)
	Save(m *Member) error
}
