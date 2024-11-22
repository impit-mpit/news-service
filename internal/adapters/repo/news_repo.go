package repo

import (
	"context"
	"neuro-most/news-service/internal/entities"
	"time"
)

type newsGORM struct {
	ID          int64     `gorm:"primaryKey"`
	Title       string    `gorm:"not null"`
	Body        string    `gorm:"not null"`
	ShortBody   string    `gorm:"not null"`
	ImageURL    string    `gorm:"not null"`
	CreatedDate time.Time `gorm:"not null"`
	CreatedBy   string    `gorm:"not null"`
	UpdatedDate *time.Time
	UpdatedBy   *string
}

type NewsRepo struct {
	db GSQL
}

func NewNewsRepo(db GSQL) NewsRepo {
	db.AutoMigrate(&newsGORM{})
	return NewsRepo{db: db}
}

func (n NewsRepo) Create(ctx context.Context, news entities.News) error {
	newsGORM := newsGORM{
		Title:       news.Title(),
		Body:        news.Body(),
		ShortBody:   news.ShortBody(),
		ImageURL:    news.ImageURL(),
		CreatedDate: news.CreatedDate(),
		CreatedBy:   news.CreatedBy(),
	}
	if err := n.db.Create(ctx, &newsGORM); err != nil {
		return entities.ErrorNewsCreate
	}
	return nil
}

func (n NewsRepo) Update(ctx context.Context, news entities.News) error {
	updates := map[string]interface{}{
		"title":        news.Title(),
		"body":         news.Body(),
		"short_body":   news.ShortBody(),
		"image_url":    news.ImageURL(),
		"updated_date": time.Now(),
		"updated_by":   news.UpdatedBy(),
	}
	if err := n.db.UpdateOne(ctx, updates, &newsGORM{ID: news.ID()}); err != nil {
		return entities.ErrorNewsUpdate
	}
	return nil
}

func (n NewsRepo) Delete(ctx context.Context, id int64) error {
	if err := n.db.Delete(ctx, &newsGORM{}, &newsGORM{ID: id}); err != nil {
		return entities.ErrorNewsDelete
	}
	return nil
}

func (n NewsRepo) GetByID(ctx context.Context, id int64) (entities.News, error) {
	var news newsGORM
	if err := n.db.BeginFind(ctx, &news).Where(&newsGORM{ID: id}).First(&news); err != nil {
		return entities.News{}, entities.ErrNewsNotFound
	}
	return n.convertNews(news), nil
}

func (n NewsRepo) Fetch(ctx context.Context, cursor int64, num int64) ([]entities.News, int64, error) {
	var news []newsGORM
	query := n.db.BeginFind(ctx, &news)
	var total int64
	query.Count(total)
	query = query.Page(int(cursor), int(num)).OrderBy("id desc")
	err := query.Find(&news)
	if err != nil {
		return nil, 0, entities.ErrorNewsFetch
	}
	var result []entities.News
	for _, news := range news {
		result = append(result, n.convertNews(news))
	}
	return result, total, nil
}

func (n NewsRepo) convertNews(news newsGORM) entities.News {
	return entities.NewNews(
		news.ID,
		news.Title,
		news.Body,
		news.ShortBody,
		news.ImageURL,
		news.CreatedDate,
		news.CreatedBy,
		news.UpdatedDate,
		news.UpdatedBy,
	)
}
