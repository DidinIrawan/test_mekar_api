package cotegory

import (
	"testAPI/masters/apis/models"
	"testAPI/masters/apis/repositories/category"
)

type JobUsecasesImpl struct {
	jobRepo category.JobRepo
}

func (j JobUsecasesImpl) UpdateJob(id int, job *models.Jobs) error {
	err := j.jobRepo.UpdateJob(id, job)
	if err != nil {
		return err
	}
	return nil
}

func (j JobUsecasesImpl) InsertJobs(job *models.Jobs) error {
	err := j.jobRepo.InsertJobs(job)
	if err != nil {
		return err
	}
	return nil
}

func (j JobUsecasesImpl) GetAllJob() ([]*models.Jobs, error) {
	jobs, err := j.jobRepo.GetAllJobs()
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func InitJobUsecase(jobRepo category.JobRepo) JobUsecases {
	return &JobUsecasesImpl{jobRepo}
}
