package usecase

import (
	"context"
	"neuro-most/news-service/internal/adapters/repo"
	"neuro-most/news-service/internal/entities"
	"time"
)

type (
	CreateNewsUseCase interface {
		Execute(ctx context.Context, input CreateNewsInput) error
	}

	CreateNewsInput struct {
		UserID    string
		Title     string
		Body      string
		ShortBody string
		ImageURL  string
	}

	createNewsInteractor struct {
		newsRepo entities.NewsRepo
	}
)

func NewCreateNewsInteractor(newsRepo repo.NewsRepo) CreateNewsUseCase {
	return &createNewsInteractor{
		newsRepo: newsRepo,
	}
}

func (i *createNewsInteractor) Execute(ctx context.Context, input CreateNewsInput) error {
	news := entities.NewNewsCreate(
		input.Title,
		input.Body,
		input.ShortBody,
		input.ImageURL,
		"",
		time.Now(),
	)
	err := i.newsRepo.Create(ctx, news)
	if err != nil {
		return err
	}
	return nil
}
