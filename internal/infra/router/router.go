package router

import (
	"context"
	"log"
	"net"
	newsv1 "neuro-most/news-service/gen/go/news/v1"
	"neuro-most/news-service/internal/adapters/api/action"
	"neuro-most/news-service/internal/adapters/api/presenter"
	"neuro-most/news-service/internal/adapters/repo"
	"neuro-most/news-service/internal/usecase"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Router struct {
	db repo.GSQL
	newsv1.UnimplementedNewsServiceServer
}

func NewRouter(db repo.GSQL) Router {
	return Router{
		db: db,
	}
}

func (r *Router) Listen() {
	port := ":3001"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts = []grpc.ServerOption{}
	srv := grpc.NewServer(opts...)
	newsv1.RegisterNewsServiceServer(srv, r)

	log.Printf("Starting gRPC server on port %s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (r *Router) CreateNews(ctx context.Context, input *newsv1.CreateNewsRequest) (*emptypb.Empty, error) {
	var (
		uc = usecase.NewCreateNewsInteractor(
			repo.NewNewsRepo(r.db),
		)
		act = action.NewCreateNewsAction(uc)
	)

	return &emptypb.Empty{}, act.Execute(ctx, input)
}

func (r *Router) DeleteNews(context.Context, *newsv1.DeleteNewsRequest) (*emptypb.Empty, error) {
	var (
		uc = usecase.NewDeleteNewsInteractor(
			repo.NewNewsRepo(r.db),
		)
		act = action.NewDeleteNewsAction(uc)
	)

	return &emptypb.Empty{}, act.Execute(context.Background(), &newsv1.DeleteNewsRequest{})
}

func (r *Router) GetNewsById(ctx context.Context, input *newsv1.GetNewsByIdRequest) (*newsv1.News, error) {
	var (
		uc = usecase.NewGetByIdNewsInteractor(
			repo.NewNewsRepo(r.db),
			presenter.NewGetByIdNewsPresenter(),
		)
		act = action.NewGetByIDNewsAction(uc)
	)

	return act.Execute(ctx, input)

}

func (r *Router) GetNewsFeed(ctx context.Context, input *newsv1.GetNewsFeedRequest) (*newsv1.GetNewsFeedResponse, error) {
	var (
		uc = usecase.NewGetAllNewsInteractor(
			repo.NewNewsRepo(r.db),
			presenter.NewGetAllNewsPresenter(),
		)
		act = action.NewGetAllNewsAction(uc)
	)

	return act.Execute(ctx, input)
}

func (r *Router) UpdateNews(ctx context.Context, input *newsv1.UpdateNewsRequest) (*emptypb.Empty, error) {
	var (
		uc = usecase.NewUpdateNewsInteractor(
			repo.NewNewsRepo(r.db),
		)
		act = action.NewUpdateNewsAction(uc)
	)

	return &emptypb.Empty{}, act.Execute(ctx, input)
}
