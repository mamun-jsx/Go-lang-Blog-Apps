package services

import (
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/repositories"
)

type ReactionService struct {
	repo *repositories.ReactionRepository
}

func NewReactionService(r *repositories.ReactionRepository) *ReactionService {
	return &ReactionService{repo: r}
}