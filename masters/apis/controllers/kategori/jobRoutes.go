package kategori

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"testAPI/masters/apis/models"
	"testAPI/masters/apis/usecases/cotegory"
	"testAPI/utils"
)

type JobsHandler struct {
	jobUsecase cotegory.JobUsecases
}

func JobsController(jobUsecase cotegory.JobUsecases) *JobsHandler  {
	return &JobsHandler{jobUsecase: jobUsecase}
}

func (j *JobsHandler) JobAPI(r *mux.Router)  {
	r.HandleFunc("/jobs", j.GetAllJob).Methods(http.MethodGet)
	r.HandleFunc("/job", j.InsertJob).Methods(http.MethodPost)
	r.HandleFunc("/job/{id}", j.UpdateJob).Methods(http.MethodPut)
}

func (j *JobsHandler) GetAllJob(w http.ResponseWriter, r *http.Request) {
	Jobs, err := j.jobUsecase.GetAllJob()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, Jobs)
	}
}

func (j *JobsHandler) InsertJob(w http.ResponseWriter, r *http.Request) {
	var job models.Jobs
	err := utils.JsonDecoder(&job, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = j.jobUsecase.InsertJobs(&job)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			utils.HandleResponse(w, http.StatusOK, job)
		}
	}
}

func (j *JobsHandler) UpdateJob(w http.ResponseWriter, r *http.Request) {
	var job *models.Jobs
	vars := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&job)
	varss, _:= strconv.Atoi(vars["id"])
	err = j.jobUsecase.UpdateJob(varss, job)
	if err != nil {
		w.Write([]byte("Data not found !!"))
		log.Println(err)
		return
	}
	byteOfJob, _ := json.Marshal(utils.GenerateResponse(http.StatusOK, "Success Update", job))
	// w.Write([]byte("Update successful"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfJob)
}