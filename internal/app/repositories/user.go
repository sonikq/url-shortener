package repositories

import (
	"github.com/sonikq/url-shortener/internal/app/models"
	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/internal/app/pkg/utils"
	"github.com/sonikq/url-shortener/pkg/cache"
	"time"
)

type UserRepo struct {
	c *cache.Cache
}

func NewUserRepo(c *cache.Cache) *UserRepo {
	return &UserRepo{
		c: c,
	}
}

func (r *UserRepo) ShorteningLink(request user.ShorteningLinkRequest) user.ShorteningLinkResponse {
	const sizeOfAlias = 6
	const httpPrefix = "http://"

	alias := utils.RandomString(sizeOfAlias)

	result := httpPrefix + request.BaseURL + alias

	r.c.Set(alias, request.ShorteningLink, 10*time.Minute)

	return user.ShorteningLinkResponse{
		Code:     201,
		Status:   success,
		Error:    nil,
		Response: &result,
	}
}

func (r *UserRepo) GetFullLinkByID(request user.GetFullLinkByIDRequest) user.GetFullLinkByIDResponse {

	fullLink, err := r.c.Get(request.ShortLinkID)
	if err != nil {
		return user.GetFullLinkByIDResponse{
			Code:   500,
			Status: fail,
			Error: &models.Err{
				Source:  "cache",
				Message: err.Error(),
			},
			Response: nil,
		}
	}

	return user.GetFullLinkByIDResponse{
		Code:     307,
		Status:   success,
		Error:    nil,
		Response: &fullLink,
	}
}
