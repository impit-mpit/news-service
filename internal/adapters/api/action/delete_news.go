package action

import (
	"context"
	newsv1 "neuro-most/news-service/gen/go/news/v1"
	"neuro-most/news-service/internal/usecase"
)

type DeleteNewsAction struct {
	uc usecase.DeleteNewsUseCase
}

func NewDeleteNewsAction(uc usecase.DeleteNewsUseCase) DeleteNewsAction {
	return DeleteNewsAction{uc: uc}
}

func (a DeleteNewsAction) Execute(ctx context.Context, input *newsv1.DeleteNewsRequest) error {
	return a.uc.Execute(ctx, usecase.DeleteNewsInput{Id: input.Id})
}
