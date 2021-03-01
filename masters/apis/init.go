package apis

import (
	"database/sql"
	"github.com/gorilla/mux"
	"testAPI/masters/apis/controllers"
	"testAPI/masters/apis/controllers/kategori"
	"testAPI/masters/apis/repositories"
	"testAPI/masters/apis/repositories/category"
	"testAPI/masters/apis/usecases"
	"testAPI/masters/apis/usecases/cotegory"
)

func Init(r *mux.Router, db *sql.DB)  {
	//EducationRepo
	educationRepo := category.InitEducationRepoImpl(db)
	educationUsecase := cotegory.InitEducationUsecase(educationRepo)
	educationController := kategori.EducationController(educationUsecase)
	educationController.EducationAPI(r)

	//JobRepo
	jobRepo := category.InitJobRepoImpl(db)
	jobUsecase := cotegory.InitJobUsecase(jobRepo)
	jobController := kategori.JobsController(jobUsecase)
	jobController.JobAPI(r)

	//UserRepo
	userRepo := repositories.InitUserRepoImpl(db)
	userUsecase := usecases.InitUserUsecase(userRepo)
	userController := controllers.UsersController(userUsecase)
	userController.UserAPI(r)
}
