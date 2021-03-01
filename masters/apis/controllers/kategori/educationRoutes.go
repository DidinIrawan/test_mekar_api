package kategori

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"testAPI/masters/apis/models"
	"testAPI/masters/apis/usecases/cotegory"
	"testAPI/utils"
)

type EducationHandler struct {
	educationUsecase cotegory.EducationUsecase
}

func EducationController(educationUsecase cotegory.EducationUsecase) *EducationHandler  {
	return &EducationHandler{educationUsecase: educationUsecase}
}

func (e *EducationHandler) EducationAPI(r *mux.Router)  {
	r.HandleFunc("/educations", e.GetAllEducation).Methods(http.MethodGet)
	r.HandleFunc("/education", e.InsertEducation).Methods(http.MethodPost)
}

func (e *EducationHandler) GetAllEducation(w http.ResponseWriter, r *http.Request) {
	Educations, err := e.educationUsecase.GetAllEducations()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, Educations)
	}
}

func (e *EducationHandler) InsertEducation(w http.ResponseWriter, r *http.Request) {
	var education models.Educations
	err := utils.JsonDecoder(&education, r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = e.educationUsecase.InsertEducation(&education)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			utils.HandleResponse(w, http.StatusOK, education)
		}
	}
}