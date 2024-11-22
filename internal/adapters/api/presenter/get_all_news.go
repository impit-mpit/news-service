package presenter

import (
	"neuro-most/news-service/internal/entities"
	"neuro-most/news-service/internal/usecase"
)

type GetAllNewsPresenter struct{}

func NewGetAllNewsPresenter() GetAllNewsPresenter {
	return GetAllNewsPresenter{}
}

func (p GetAllNewsPresenter) Output(news []entities.News) []usecase.GetAllNewsOutput {
	newsFeed := make([]usecase.GetAllNewsOutput, len(news))
	for i, n := range news {
		newsFeed[i] = usecase.GetAllNewsOutput{
			Id:        n.ID(),
			Title:     n.Title(),
			ShortBody: n.ShortBody(),
			ImageURL:  n.ImageURL(),
		}
	}
	return newsFeed
}
