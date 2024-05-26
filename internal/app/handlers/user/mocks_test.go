package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/internal/app/pkg/logger"
	"github.com/sonikq/url-shortener/internal/app/services"
	"github.com/stretchr/testify/mock"
)

func NewTestHandler() (*gin.Engine, *MockServiceManager, Handler) {
	gin.SetMode(gin.TestMode)

	mockServiceManager := new(MockServiceManager)
	log := logger.New("info", "test_handler")

	handler := Handler{
		service: &services.Service{
			IUserService: mockServiceManager,
		},
		log: log,
	}

	r := gin.Default()
	return r, mockServiceManager, handler
}

type MockServiceManager struct {
	mock.Mock
}

func (m *MockServiceManager) ShorteningLink(ctx context.Context, request user.ShorteningLinkRequest) user.ShorteningLinkResponse {
	args := m.Called(ctx, request)
	return args.Get(0).(user.ShorteningLinkResponse)
}

func (m *MockServiceManager) GetFullLinkByID(ctx context.Context, request user.GetFullLinkByIDRequest) user.GetFullLinkByIDResponse {
	args := m.Called(ctx, request)
	return args.Get(0).(user.GetFullLinkByIDResponse)
}

func (m *MockServiceManager) ShorteningLinkJSON(ctx context.Context, request user.ShorteningLinkJSONRequest) user.ShorteningLinkJSONResponse {
	args := m.Called(ctx, request)
	return args.Get(0).(user.ShorteningLinkJSONResponse)
}

func (m *MockServiceManager) ShorteningBatchLinks(ctx context.Context, request user.ShorteningBatchLinksRequest) user.ShorteningBatchLinksResponse {
	args := m.Called(ctx, request)
	return args.Get(0).(user.ShorteningBatchLinksResponse)
}

func (m *MockServiceManager) GetBatchByUserID(ctx context.Context, request user.GetBatchByUserIDRequest) user.GetBatchByUserIDResponse {
	args := m.Called(ctx, request)
	return args.Get(0).(user.GetBatchByUserIDResponse)
}

func (m *MockServiceManager) PingDB(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}
