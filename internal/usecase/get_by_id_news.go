package usecase

import (
	"context"
	"neuro-most/news-service/internal/entities"
)

type (
	GetByIdNewsUseCase interface {
		Execute(ctx context.Context, input GetByIdNewsUseCaseInput) (GetByIdNewsOutput, error)
	}

	GetByIdNewsUseCaseInput struct {
		Id int64
	}

	GetByIdNewsOutput struct {
		Id        int64
		Title     string
		Body      string
		ImageURL  string
		CreatedBy string
	}

	GetByIdNewsUseCasePresenter interface {
		Output(news entities.News) GetByIdNewsOutput
	}

	GetByIdNewsInteractor struct {
		newsRepo  entities.NewsRepo
		presenter GetByIdNewsUseCasePresenter
	}
)

func NewGetByIdNewsInteractor(newsRepo entities.NewsRepo, presenter GetByIdNewsUseCasePresenter) GetByIdNewsUseCase {
	return &GetByIdNewsInteractor{
		newsRepo:  newsRepo,
		presenter: presenter,
	}
}

func (g *GetByIdNewsInteractor) Execute(ctx context.Context, input GetByIdNewsUseCaseInput) (GetByIdNewsOutput, error) {
	news, err := g.newsRepo.GetByID(ctx, input.Id)
	if err != nil {
		return GetByIdNewsOutput{}, err
	}
	return g.presenter.Output(news), nil
}
