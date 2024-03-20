package jobs

import (
	"github.com/GoWebProd/uuid7"
	"github.com/white43/sep401/pkg/dto"
	"github.com/white43/sep401/pkg/errors"
	"github.com/white43/sep401/pkg/users"
	"unicode"
)

type Service struct {
	repository *JobRepository
	uuid       *uuid7.Generator
}

func NewService(repository *JobRepository, uuid *uuid7.Generator) *Service {
	return &Service{repository, uuid}
}

func NewDummyService() *Service {
	return &Service{}
}

func (s Service) ValidateNewJobRequest(user users.User, request dto.NewJobRequest) error {
	if user.Balance == 0 {
		return errors.YourBalanceIsNotEnoughForNewJobs
	}

	if s.IsHTMLPayloadLengthValid(request.Payload) == false {
		return errors.PayloadWrongLength
	}

	if s.DoesStringContainOnlyPrintableCharacters(request.Payload) == false {
		return errors.StringShallContainOnlyPrintableCharacters
	}

	return nil
}

func (s Service) ValidateJobResultRequest(id string) error {
	if id == "" {
		return errors.EmptyJobID
	}

	return nil
}

func (s Service) IsHTMLPayloadLengthValid(str string) bool {
	length := len(str)
	return length > 0 && length <= 255*1024
}

func (s Service) DoesStringContainOnlyPrintableCharacters(str string) bool {
	for _, char := range str {
		if char > unicode.MaxASCII || !unicode.IsPrint(char) {
			return false
		}
	}

	return true
}

func (s Service) AddJob(jobType string, userId int, request dto.NewJobRequest) (dto.NewJobResponse, error) {
	job, err := s.repository.InsertJob(s.uuid.Next().String(), request.Payload, userId, jobType)
	if err != nil {
		return dto.NewJobResponse{}, err
	}

	if job.ID == "" {
		return dto.NewJobResponse{}, errors.NewServerError("something went wrong during job creation process")
	}

	result := dto.NewJobResponse{
		ID: job.ID,
	}

	return result, nil
}

func (s Service) GetJobResult(id string) (dto.JobResultResponse, error) {
	job, err := s.repository.GetByID(id)
	if err != nil {
		return dto.JobResultResponse{}, err
	}

	switch {
	case job.Status == "pending" || job.Status == "processing":
		return dto.JobResultResponse{}, errors.JobProcessingInProgress
	case job.Status == "error":
		return dto.JobResultResponse{}, errors.JobProcessingHasFailed
	}

	response := dto.JobResultResponse{
		Result: job.Result.String,
	}

	return response, nil
}

func (s Service) GetNextJob() (Job, error) {
	return s.repository.GetNextJob()
}

func (s Service) UpdateJobStatus(job Job, status string) error {
	return s.repository.UpdateJobStatus(job, status)
}

func (s Service) SaveJobResult(job Job, result string) error {
	return s.repository.UpdateJobResult(job, result)
}
