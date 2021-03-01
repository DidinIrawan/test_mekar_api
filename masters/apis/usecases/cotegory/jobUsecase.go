package cotegory

import "testAPI/masters/apis/models"

type JobUsecases interface {
	GetAllJob()([]*models.Jobs, error)
	InsertJobs(job *models.Jobs) error
	UpdateJob(id int, job *models.Jobs) error
}