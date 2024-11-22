package action

import (
	"context"
	newsv1 "neuro-most/news-service/gen/go/news/v1"
	"neuro-most/news-service/internal/usecase"
)

type UpdateNewsAction struct {
	uc usecase.UpdateNewsUseCase
}

func NewUpdateNewsAction(uc usecase.UpdateNewsUseCase) *UpdateNewsAction {
	return &UpdateNewsAction{
		uc: uc,
	}
}

func (a *UpdateNewsAction) Execute(ctx context.Context, input *newsv1.UpdateNewsRequest) error {
	var usecaseInput usecase.UpdateNewsInput
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
