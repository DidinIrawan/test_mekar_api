package category

import (
	"database/sql"
	"testAPI/masters/apis/models"
	"testAPI/utils"
)

type EducatinRepoImpl struct {
	db *sql.DB
}

func (e EducatinRepoImpl) GetAllEducations() ([]*models.Educations, error) {
	var educations []*models.Educations

	query := utils.GETALLEDUCATIONS
	rows, err := e.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		education := models.Educations{}
		err := rows.Scan(&education.IdEducation, &education.Education,&education.Status)

		if err != nil {
			return nil, err
		}

		educations = append(educations, &education)

	}

	return educations, nil
}

func (e EducatinRepoImpl) InsertEducation(education *models.Educations) error {
	tx, err := e.db.Begin()
	if err != nil {
		return err
	}
	query := utils.GETINSERTEDUCATION
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(education.Education); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func InitEducationRepoImpl(db *sql.DB) EducationRepository  {
	return &EducatinRepoImpl{db}
}
