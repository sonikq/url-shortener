package services

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sonikq/url-shortener/internal/app/models"
	"github.com/sonikq/url-shortener/internal/app/models/user"
	pb "github.com/sonikq/url-shortener/internal/app/proto"
	"github.com/sonikq/url-shortener/internal/app/repositories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

// ServiceGrpc -
type ServiceGrpc struct {
	pb.UnimplementedShortenerServer
	Repo    repositories.IUserRepo
	BaseURL string
}

func (s *ServiceGrpc) Shorten(ctx context.Context, req *pb.ShortenRequest) (*pb.ShortenResponse, error) {
	var resp pb.ShortenResponse

	result := s.Repo.ShorteningLink(ctx, user.ShorteningLinkRequest{
		UserID:         req.UserId,
		ShorteningLink: req.Url,
		BaseURL:        s.BaseURL,
	})
	if result.Error != nil {
		switch result.Code {
		case http.StatusConflict:
			return nil, status.Error(codes.AlreadyExists, models.ErrAlreadyExists.Error())
		default:
			return nil, status.Error(codes.Internal, result.Error.Message)
		}
	}
	resp.Shorten = *result.Response
	return &resp, nil
}

func (s *ServiceGrpc) Expand(ctx context.Context, req *pb.ExpandRequest) (*pb.ExpandResponse, error) {
	var resp pb.ExpandResponse

	result := s.Repo.GetFullLinkByID(ctx, user.GetFullLinkByIDRequest{ShortLinkID: req.ShortUrl})
	if result.Error != nil {
		switch result.Code {
		case http.StatusGone:
			return nil, status.Error(codes.DataLoss, models.ErrGetDeletedLink.Error())
		default:
			return nil, status.Error(codes.Internal, result.Error.Message)
		}
	}

	resp.Url = *result.Response
	return &resp, nil
}

func (s *ServiceGrpc) GetBatch(ctx context.Context, req *pb.GetBatchRequest) (*pb.GetBatchResponse, error) {
	var resp pb.GetBatchResponse

	result := s.Repo.GetBatchByUserID(ctx, user.GetBatchByUserIDRequest{
		BaseURL: s.BaseURL,
		UserID:  req.UserId,
	})
	if result.Error != nil {
		switch result.Code {
		case http.StatusNoContent:
			return nil, status.Error(codes.NotFound, "there are no urls")
		default:
			return nil, status.Error(codes.Internal, result.Error.Message)
		}
	}
	for _, v := range result.Response {
		resp.Rows = append(resp.Rows, &pb.UrlRow{
			OriginalURL: v.OriginalURL,
			ShortURL:    v.ShortURL,
		})
	}
	return &resp, nil
}

func (s *ServiceGrpc) Batch(ctx context.Context, req *pb.ShortBatchRequest) (*pb.ShortBatchResponse, error) {
	var resp pb.ShortBatchResponse

	var batchLinksReq user.ShorteningBatchLinksRequest

	for _, v := range req.Original {
		batchLinksReq.Body = append(batchLinksReq.Body, user.BatchUrlsInput{
			CorrelationID: v.CorrelationId,
			OriginalURL:   v.OriginalUrl,
		})
	}

	result := s.Repo.ShorteningBatchLinks(ctx, batchLinksReq)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Message)
	}

	for _, val := range result.Response {
		resp.Original = append(resp.Original, &pb.CorrelationShortURL{
			CorrelationId: val.CorrelationID,
			ShortUrl:      val.ShortURL,
		})
	}

	return &resp, nil
}

func (s *ServiceGrpc) Ping(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.Repo.PingDB(ctx)
}

func (s *ServiceGrpc) GetStats(ctx context.Context, req *empty.Empty) (*pb.GetStatsResponse, error) {
	var resp pb.GetStatsResponse

	result := s.Repo.GetStats(ctx)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Message)
	}

	resp.Urls = int32(result.Response.URL)
	resp.Users = int32(result.Response.Users)

	return &resp, nil
}
