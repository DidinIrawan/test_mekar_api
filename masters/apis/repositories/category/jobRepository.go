package category

import "testAPI/masters/apis/models"

type JobRepo interface {
	GetAllJobs() ([]*models.Jobs, error)
	InsertJobs(job *models.Jobs) error
	UpdateJob(id int, job *models.Jobs) error
}
