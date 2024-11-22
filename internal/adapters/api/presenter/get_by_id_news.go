package presenter

import (
	"neuro-most/news-service/internal/entities"
	"neuro-most/news-service/internal/usecase"
)

type GetByIdNewsPresenter struct{}

func NewGetByIdNewsPresenter() GetByIdNewsPresenter {
	return GetByIdNewsPresenter{}
}

func (p GetByIdNewsPresenter) Output(news entities.News) usecase.GetByIdNewsOutput {
	return usecase.GetByIdNewsOutput{
		Id:        news.ID(),
		Title:     news.Title(),
		Body:      news.Body(),
		ImageURL:  news.ImageURL(),
		CreatedBy: news.CreatedBy(),
	}
}
