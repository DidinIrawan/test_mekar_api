package category

import "testAPI/masters/apis/models"

type EducationRepository interface {
	GetAllEducations() ([]*models.Educations, error)
	InsertEducation(education *models.Educations) error
}