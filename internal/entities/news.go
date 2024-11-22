package entities

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNewsNotFound = status.New(codes.NotFound, "news not found").Err()
	ErrorNewsCreate = status.New(codes.Internal, "error create news").Err()
	ErrorNewsUpdate = status.New(codes.Internal, "error update news").Err()
	ErrorNewsDelete = status.New(codes.Internal, "error delete news").Err()
	ErrorNewsFetch  = status.New(codes.Internal, "error fetch news").Err()
)

type (
	NewsRepo interface {
		Create(ctx context.Context, news News) error
		Update(ctx context.Context, news News) error
		Delete(ctx context.Context, id int64) error
		GetByID(ctx context.Context, id int64) (News, error)
		Fetch(ctx context.Context, cursor int64, num int64) ([]News, int64, error)
	}

	News struct {
		id          int64
		title       string
		body        string
		shortBody   string
		imageURL    string
		createdDate time.Time
		createdBy   string
		updatedDate *time.Time
		updatedBy   *string
	}
)

func NewNews(
	id int64,
	title string,
	body string,
	shortBody string,
	imageURL string,
	createDate time.Time,
	createdBy string,
	updatedDate *time.Time,
	updatedBy *string,
) News {
	return News{
		id:          id,
		title:       title,
		body:        body,
		shortBody:   shortBody,
		imageURL:    imageURL,
		createdDate: createDate,
		createdBy:   createdBy,
		updatedDate: updatedDate,
		updatedBy:   updatedBy,
	}
}

func NewNewsCreate(
	title string,
	body string,
	shortBody string,
	imageURL string,
	createdBy string,
	createdDate time.Time,
) News {
	return News{
		title:       title,
		body:        body,
		shortBody:   shortBody,
		imageURL:    imageURL,
		createdBy:   createdBy,
		createdDate: createdDate,
	}
}

func (n News) ID() int64 {
	return n.id
}

func (n News) Title() string {
	return n.title
}

func (n News) Body() string {
	return n.body
}

func (n News) ShortBody() string {
	return n.shortBody
}

func (n News) ImageURL() string {
	return n.imageURL
}

func (n News) CreatedDate() time.Time {
	return n.createdDate
}

func (n News) CreatedBy() string {
	return n.createdBy
}

func (n News) UpdatedDate() *time.Time {
	return n.updatedDate
}

func (n News) UpdatedBy() *string {
	return n.updatedBy
}

func (n *News) SetID(id int64) {
	n.id = id
}

func (n *News) SetTitle(title string) {
	n.title = title
}

func (n *News) SetBody(body string) {
	n.body = body
}

func (n *News) SetShortBody(shortBody string) {
	n.shortBody = shortBody
}

func (n *News) SetImageURL(imageURL string) {
	n.imageURL = imageURL
}

func (n *News) SetCreatedDate(createdDate time.Time) {
	n.createdDate = createdDate
}

func (n *News) SetCreatedBy(createdBy string) {
	n.createdBy = createdBy
}

func (n *News) SetUpdatedDate(updatedDate *time.Time) {
	n.updatedDate = updatedDate
}

func (n *News) SetUpdatedBy(updatedBy *string) {
	n.updatedBy = updatedBy
}
