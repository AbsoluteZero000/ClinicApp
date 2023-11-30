package main

import (
	"awesomeProject/Config"
	"awesomeProject/Controller"
	"awesomeProject/Repo"
	"fmt"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := Config.Connect()
	fmt.Print(Repo.GetAllUser(db))

	Controller.InitialData()

	router := gin.Default()

	defaultConfig := cors.DefaultConfig()
	defaultConfig.AllowAllOrigins = true
	defaultConfig.AllowHeaders = append(defaultConfig.AllowHeaders, "Authorization")
	defaultConfig.AllowCredentials = true
	defaultConfig.AllowWebSockets = true
	router.Use(cors.New(defaultConfig))

	router.POST("/signup", Controller.SignUp)
	router.POST("/login", Controller.Login)

	router.GET("/getusers", Controller.GetAllUsers)

	router.POST("/addpatientslot", Controller.AddPatientSlot)
	router.GET("/getpatientslots", Controller.GetPatientSlots)
	router.PUT("/editpatientslot", Controller.EditPatientSlot)

	router.POST("/addslot", Controller.AddDoctorSlot)
	router.PUT("/editdoctorslot", Controller.EditDoctorSlot)
	router.GET("/getdoctorslots", Controller.GetDoctorSlots)
	router.GET("/getfreeslots", Controller.GetFreeDoctorSlot)
	router.DELETE("/deleteslot", Controller.DeleteSlot)

	router.Run(":8080")
}
