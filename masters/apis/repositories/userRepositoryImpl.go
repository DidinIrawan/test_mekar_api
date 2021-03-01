package repositories

import (
	"database/sql"
	"testAPI/masters/apis/models"
	"testAPI/utils"
)

type UserRepoImpl struct {
	db *sql.DB
}

func (u UserRepoImpl) UpdateUser(id int, user *models.Users) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	query := utils.GETUPDATEUSERBYID
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Nama,user.Date,user.NoKTP,user.JobID,user.EducationID, id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

func (u UserRepoImpl) InsertUser(user *models.Users) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	query := utils.GETINSERTUSER
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(user.Nama,user.Date,user.NoKTP,user.JobID,user.EducationID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (u UserRepoImpl) GetUserById(id int) (*models.Users, error) {
	user := new(models.Users)
	query := utils.GETUSERBYID
	if err := u.db.QueryRow(query, id).Scan(&user.UserID, &user.Nama, &user.Date,&user.NoKTP,&user.JobID,&user.Job,&user.EducationID,&user.Education,&user.Status); err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserRepoImpl) GetAllUsers() ([]*models.Users, error) {
	var listUsers []*models.Users
	query := utils.GETALLUSERS
	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		listUser := models.Users{}
		err := rows.Scan(&listUser.UserID, &listUser.Nama, &listUser.Date,&listUser.NoKTP,&listUser.JobID,&listUser.Job,&listUser.EducationID,&listUser.Education,&listUser.Status)

		if err != nil {
			return nil, err
		}

		listUsers = append(listUsers, &listUser)

	}
	return listUsers, nil
}

func InitUserRepoImpl(db *sql.DB) UserRepo {
	return &UserRepoImpl{db: db}
}