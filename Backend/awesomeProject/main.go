package main

import (
	"awesomeProject/Controller"
	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//allowedRoles := map[string][]string{
	//	"patient": {"/editpatientslot", "/deleteslot", "/getpatientslots", "/addpatientslot", "/updateuser", "/getusers", "/deleteuser"},
	//	"doctor":  {"/editdoctorslot", "/deleteslot", "/addslot", "/getdoctorslots", "/updateuser", "/getusers", "/deleteuser"},
	//	"admin":   {"/editpatientslot", "/getpatientslots", "/addpatientslot", "/editdoctorslot", "/deleteslot", "/addslot", "/getdoctorslots", "/updateuser", "/getusers", "/deleteuser"},
	//}

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

	router.DELETE("/deleteslot", Controller.DeleteSlot)

	router.Run(":8080")
}
