package usecase

import "ca-example/internal/domain"

type RewardPointsUseCase struct {
	Repo domain.MemberRepositiry
}

func (uc *RewardPointsUseCase) Execute(memberID string, points int) error {
	member, err := uc.Repo.Get(memberID)
	if err != nil {
		return err
	}

	member.AddPoints(points)

	return uc.Repo.Save(member)
}
