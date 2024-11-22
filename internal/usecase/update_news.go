package usecase

import (
	"context"
	"neuro-most/news-service/internal/adapters/repo"
	"neuro-most/news-service/internal/entities"
)

type (
	UpdateNewsUseCase interface {
		Execute(ctx context.Context, input UpdateNewsInput) error
	}

	UpdateNewsInput struct {
		Id        int64
		UserID    string
		Title     *string
		Body      *string
		ShortBody *string
		ImageURL  *string
	}

	UpdateNewsInteractor struct {
		newsRepo entities.NewsRepo
	}
)

func NewUpdateNewsInteractor(newsRepo repo.NewsRepo) UpdateNewsUseCase {
	return &UpdateNewsInteractor{
		newsRepo: newsRepo,
	}
}

func (i *UpdateNewsInteractor) Execute(ctx context.Context, input UpdateNewsInput) error {
	news, err := i.newsRepo.GetByID(ctx, input.Id)
	if err != nil {
		return err
	}
	if input.Title != nil {
		news.SetTitle(*input.Title)
	}
	if input.Body != nil {
		news.SetBody(*input.Body)
	}
	if input.ShortBody != nil {
		news.SetShortBody(*input.ShortBody)
	}
	if input.ImageURL != nil {
		news.SetImageURL(*input.ImageURL)
	}
	err = i.newsRepo.Update(ctx, news)
	if err != nil {
		return err
	}
	return nil
}
