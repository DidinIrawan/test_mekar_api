package models

type Users struct {
	UserID int `json:"user_id"`
	Nama string `json:"nama"`
	Date string `json:"date"`
	NoKTP string `json:"no_ktp"`
	JobID int `json:"job_id"`
	Job string `json:"job"`
	EducationID int `json:"education_id"`
	Education string `json:"education"`
	Status string `json:"status"`
}
