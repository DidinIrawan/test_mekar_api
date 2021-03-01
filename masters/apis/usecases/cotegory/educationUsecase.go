package cotegory

import "testAPI/masters/apis/models"

type EducationUsecase interface {
	GetAllEducations() ([]*models.Educations, error)
	InsertEducation(education *models.Educations) error
}
