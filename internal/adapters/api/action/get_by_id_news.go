package action

import (
	"context"
	newsv1 "neuro-most/news-service/gen/go/news/v1"
	"neuro-most/news-service/internal/usecase"
)

type GetByIDNewsAction struct {
	uc usecase.GetByIdNewsUseCase
}

func NewGetByIDNewsAction(uc usecase.GetByIdNewsUseCase) *GetByIDNewsAction {
	return &GetByIDNewsAction{
		uc: uc,
	}
}

func (a *GetByIDNewsAction) Execute(ctx context.Context, input *newsv1.GetNewsByIdRequest) (*newsv1.News, error) {
	var usecaseInput usecase.GetByIdNewsUseCaseInput
	usecaseInput.Id = input.Id

	news, err := a.uc.Execute(ctx, usecaseInput)
	if err != nil {
		return nil, err
	}
	newsFeed := &newsv1.News{
		Id:       news.Id,
		Title:    news.Title,
		Body:     news.Body,
		ImageUrl: news.ImageURL,
	}
	return newsFeed, nil
}
