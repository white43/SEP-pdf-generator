package generator

import (
	"github.com/white43/sep401/pkg/jobs"
	"github.com/white43/sep401/pkg/users"
)

type Service struct {
	jobService  *jobs.Service
	userService *users.Service
	factory     FactoryInterface
}

func NewService(jobService *jobs.Service, userService *users.Service, factory FactoryInterface) *Service {
	return &Service{jobService, userService, factory}
}

func (s Service) GetNextJob() (jobs.Job, error) {
	return s.jobService.GetNextJob()
}

func (s Service) Process(job jobs.Job) ([]byte, error) {
	return s.factory.FactoryMethod(job).Process(job.Payload)
}

func (s Service) MarkJobSuccessful(job jobs.Job, result string) error {
	if err := s.jobService.SaveJobResult(job, result); err != nil {
		return err
	}

	return s.jobService.UpdateJobStatus(job, "successful")
}

func (s Service) MarkJobFailed(job jobs.Job) error {
	return s.jobService.UpdateJobStatus(job, "error")
}

func (s Service) UpdateUserBalance(job jobs.Job, balance float64) error {
	return s.userService.DeductBalance(job.UserId, balance)
}
