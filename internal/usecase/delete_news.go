package usecase

import (
	"context"
	"neuro-most/news-service/internal/adapters/repo"
)

type (
	DeleteNewsUseCase interface {
		Execute(ctx context.Context, input DeleteNewsInput) error
	}

	DeleteNewsInput struct {
		Id int64
	}

	deleteNewsInteractor struct {
		repo repo.NewsRepo
	}
)

func NewDeleteNewsInteractor(repo repo.NewsRepo) DeleteNewsUseCase {
	return &deleteNewsInteractor{repo: repo}
}

func (d *deleteNewsInteractor) Execute(ctx context.Context, input DeleteNewsInput) error {
	if err := d.repo.Delete(ctx, input.Id); err != nil {
		return err
	}
	return nil
}
