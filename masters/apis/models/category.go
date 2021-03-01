package models

type Jobs struct {
	IdJob int `json:"id_job"`
	Job string `json:"job"`
	Status string `json:"status"`
}

type Educations struct {
	IdEducation int `json:"id_education"`
	Education string `json:"education"`
	Status string `json:"status"`
}
