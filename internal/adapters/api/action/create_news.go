package action

import (
	"context"
	newsv1 "neuro-most/news-service/gen/go/news/v1"
	"neuro-most/news-service/internal/usecase"
)

type CreateNewsAction struct {
	uc usecase.CreateNewsUseCase
}

func NewCreateNewsAction(uc usecase.CreateNewsUseCase) *CreateNewsAction {
	return &CreateNewsAction{
		uc: uc,
	}
}

func (a *CreateNewsAction) Execute(ctx context.Context, input *newsv1.CreateNewsRequest) error {
	var usecaseInput usecase.CreateNewsInput
	usecaseInput.Title = input.Title
	usecaseInput.Body = input.Body
	usecaseInput.ShortBody = input.ShortBody
	usecaseInput.ImageURL = input.ImageUrl

	err := a.uc.Execute(ctx, usecaseInput)
	if err != nil {
		return err
	}
	return nil
}
