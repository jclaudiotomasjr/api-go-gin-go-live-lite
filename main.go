package main

import (
	
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jclaudiotomasjr/api-gin-go-live/controllers"
	"github.com/jclaudiotomasjr/api-gin-go-live/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.ConectaBD()
	router := gin.Default()

	public := router.Group("/api")

	public.POST("/score/", controllers.CreateScore)
	public.GET("/score/", controllers.AllScores)
	public.GET("/score/:id", controllers.GetScoreId)
	public.GET("/score/cpf/:cpf", controllers.GetScoreCpf)
	public.DELETE("/score/:id", controllers.DeleteScore)

	http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), router)

}
