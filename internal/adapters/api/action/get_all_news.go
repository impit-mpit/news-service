package action

import (
	"context"
	newsv1 "neuro-most/news-service/gen/go/news/v1"
	"neuro-most/news-service/internal/usecase"
)

type GetAllNewsAction struct {
	uc usecase.GetAllNewsUseCase
}

func NewGetAllNewsAction(uc usecase.GetAllNewsUseCase) *GetAllNewsAction {
	return &GetAllNewsAction{
		uc: uc,
	}
}

func (a *GetAllNewsAction) Execute(ctx context.Context, input *newsv1.GetNewsFeedRequest) (*newsv1.GetNewsFeedResponse, error) {
	var usecaseInput usecase.GetAllNewsUseCaseInput
	usecaseInput.Page = int64(input.Page)
	usecaseInput.PageSize = int64(input.PageSize)

	news, total, err := a.uc.Execute(ctx, usecaseInput)
	if err != nil {
		return nil, err
	}
	newsFeed := &newsv1.GetNewsFeedResponse{
		Total: int32(total),
		News:  make([]*newsv1.ShortNews, 0),
	}
	for _, n := range news {
		newsFeed.News = append(newsFeed.News, &newsv1.ShortNews{
			Id:        n.Id,
			Title:     n.Title,
			ShortBody: n.ShortBody,
			ImageUrl:  n.ImageURL,
		})
	}
	return newsFeed, nil
}
