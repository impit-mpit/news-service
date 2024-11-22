package usecase

import (
	"context"
	"neuro-most/news-service/internal/entities"
)

type (
	GetAllNewsUseCase interface {
		Execute(ctx context.Context, input GetAllNewsUseCaseInput) ([]GetAllNewsOutput, int64, error)
	}

	GetAllNewsUseCaseInput struct {
		Page     int64
		PageSize int64
	}

	GetAllNewsOutput struct {
		Id        int64
		Title     string
		ShortBody string
		ImageURL  string
	}

	GetAllNewsUseCasePresenter interface {
		Output(news []entities.News) []GetAllNewsOutput
	}

	getAllNewsInteractor struct {
		newsRepo  entities.NewsRepo
		presenter GetAllNewsUseCasePresenter
	}
)

func NewGetAllNewsInteractor(newsRepo entities.NewsRepo, presenter GetAllNewsUseCasePresenter) GetAllNewsUseCase {
	return &getAllNewsInteractor{
		newsRepo:  newsRepo,
		presenter: presenter,
	}
}

func (g *getAllNewsInteractor) Execute(ctx context.Context, input GetAllNewsUseCaseInput) ([]GetAllNewsOutput, int64, error) {
	news, total, err := g.newsRepo.Fetch(ctx, input.Page, input.PageSize)
	if err != nil {
		return []GetAllNewsOutput{}, 0, err
	}
	return g.presenter.Output(news), total, nil
}
