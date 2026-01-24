package main

import (
	"ca-example/infra/memory"
	adapter "ca-example/internal/adapter/http"
	"ca-example/internal/usecase"
	"net/http"
)

func main() {
	repo := memory.NewInMemMemberRepository()

	rewardUC := &usecase.RewardPointsUseCase{Repo: repo}

	handler := &adapter.MemberHandler{UC: rewardUC}

	http.ListenAndServe(":8080", handler)
}
