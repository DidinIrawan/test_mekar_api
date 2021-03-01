package category

import (
	"database/sql"
	"testAPI/masters/apis/models"
	"testAPI/utils"
)

type JobRepoImpl struct {
	db *sql.DB
}

func (j *JobRepoImpl) UpdateJob(id int, job *models.Jobs) error {
	tx, err := j.db.Begin()
	if err != nil {
		return err
	}
	query := utils.GETUPDATEJOB
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(job.Job, id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

func (j *JobRepoImpl) InsertJobs(job *models.Jobs) error {
	tx, err := j.db.Begin()
	if err != nil {
		return err
	}
	query := utils.GETINSERTJOB
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(job.Job); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (j *JobRepoImpl) GetAllJobs() ([]*models.Jobs, error) {
	var listJobs []*models.Jobs
	query := utils.GETALLJOB
	rows, err := j.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		listjob := models.Jobs{}
		err := rows.Scan(&listjob.IdJob, &listjob.Job, &listjob.Status)

		if err != nil {
			return nil, err
		}

		listJobs = append(listJobs, &listjob)

	}
	return listJobs, nil
}

func InitJobRepoImpl(db *sql.DB) JobRepo {
	return &JobRepoImpl{db: db}
}