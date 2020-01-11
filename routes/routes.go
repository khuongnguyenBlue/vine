package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khuongnguyenBlue/vine/controllers"
	"github.com/khuongnguyenBlue/vine/middlewares/jwt"
)

func Setup(c *controllers.Controller) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	api.POST("login", c.Login)
	api.POST("register", c.Register)
	api.Use(jwt.JWT())
	{
		api.GET("demo", c.Demo)
		api.GET("subjects", c.GetSubjects)
		api.GET("subjects/:id/exams", c.GetExams)
		api.GET("exams/:id", c.GetExam)
		api.GET("exams/:id/test", c.TakeExam)
		api.POST("exams/:id", c.SubmitExam)
	}
	return r
}
