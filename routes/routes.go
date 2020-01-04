package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khuongnguyenBlue/vine/controllers"
)

func Setup() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("demo", controllers.Demo)
		api.POST("login", controllers.Login)
		api.GET("subjects", controllers.GetSubjects)
		api.GET("subjects/:id/exams", controllers.GetExams)
		api.GET("exams/:id", controllers.GetExam)
	}
	return r
}
