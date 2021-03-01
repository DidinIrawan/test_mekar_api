package cotegory

import (
	"testAPI/masters/apis/models"
	"testAPI/masters/apis/repositories/category"
)

type EducationUsecaseImpl struct {
	educationRepo category.EducationRepository
}

func (e EducationUsecaseImpl) InsertEducation(education *models.Educations) error {
	err := e.educationRepo.InsertEducation(education)
	if err != nil {
		return err
	}
	return nil
}

func (e EducationUsecaseImpl) GetAllEducations() ([]*models.Educations, error) {
	educations, err := e.educationRepo.GetAllEducations()
	if err != nil{
		return nil, err
	}
	return educations,nil
}

func InitEducationUsecase(education category.EducationRepository) EducationUsecase  {
	return &EducationUsecaseImpl{educationRepo: education}
}
