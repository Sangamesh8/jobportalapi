package services

import (
	"context"
	"job-portal-api/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	CreateUser(ctx context.Context, nu models.NewUser) (models.User, error)
	Authenticate(ctx context.Context, email, password string) (jwt.RegisteredClaims, error)
	CreateCompany(ctx context.Context, newComp models.Company) (models.Company, error)
	ViewCompanies(ctx context.Context) ([]models.Company, error)
	AddJobs(jobs []models.Job, compId string) ([]models.Job, error)
	GetJobById(ctx context.Context, jobId string) (models.Job, error)
	FetchJobByCompanyId(ctx context.Context, companyId string) ([]models.Job, error)
	GetCompanyByID(ctx context.Context, companyId string) (models.Company, error)
	GetAllJobs(ctx context.Context) ([]models.Job, error)
}

type Store struct {
	Service
}

func NewStore(s Service) Store {
	return Store{Service: s}
}
