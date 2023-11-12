package Controller

import (
	"awesomeProject/Config"
	"awesomeProject/Model"
	"awesomeProject/Repo"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func InitialData() {
	Repo.InsertUser(Model.User{
		UserWithoutPassword: Model.UserWithoutPassword{
			Name:  "admin",
			Email: "admin",
			Role:  "admin",
		},
		Password: "123",
	}, Config.Connect())

	Repo.InsertUser(Model.User{
		UserWithoutPassword: Model.UserWithoutPassword{
			Name:  "admin1",
			Email: "admin1",
			Role:  "patient",
		},
		Password: "123",
	}, Config.Connect())

	Repo.InsertUser(Model.User{
		UserWithoutPassword: Model.UserWithoutPassword{
			Name:  "admin2",
			Email: "admin2",
			Role:  "admin",
		},
		Password: "123",
	}, Config.Connect())

	Repo.InsertSlot(Model.Slot{
		DoctorId: "1",
		Date:     "2000-11-11 12:12:12",
	}, Config.Connect())

}
func GetAllUsers(c *gin.Context) {
	BasicAuth(c)
	var User Model.User
	var arrUser []Model.UserWithoutPassword

	db := Config.Connect()
	defer db.Close()

	rows, err := Repo.GetAllUser(db)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&User.Id, &User.Name, &User.Email, &User.Password, &User.Role)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrUser = append(arrUser, User.UserWithoutPassword)
		}
	}

	c.JSON(200, gin.H{
		"Message": "Success",
		"Data":    arrUser,
	})

}

// SignUp = Insert User API
func SignUp(c *gin.Context) {

	db := Config.Connect()
	defer db.Close()

	var user Model.User

	c.Bind(user)
	_, err := Repo.InsertUser(user, db)

	if err != nil {
		log.Print(err)
		c.JSON(409, gin.H{
			"message": "Failed to insert Data",
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "Inserted data",
		"data":    user,
	})

	fmt.Print("Insert data to database\n")
}

func BasicAuth(c *gin.Context) {
	var user Model.User
	name, password, _ := c.Request.BasicAuth()

	db := Config.Connect()
	defer db.Close()

	rows, err := Repo.GetAllUser(db)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role)

		if err != nil {

			log.Fatal(err.Error())
			c.Abort()
		} else {
			fmt.Print("Comparing With ", name, " ", password, "  ")
			fmt.Printf("Checking user: %s, Password: %s\n", user.Name, user.Password)
			if password == user.Password && name == user.Name {
				fmt.Println("User found and authenticated!")
				log.WithFields(log.Fields{
					"user": user,
				}).Info("User authenticated")
			}
		}
	}

	c.Abort()
	c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
	return

}

func Login(c *gin.Context) {
	BasicAuth(c)
	user, _, _ := c.Request.BasicAuth()

	db := Config.Connect()
	defer db.Close()

	authenticatedUser, err := Repo.GetUserbyUserName(user, db)

	fmt.Println(authenticatedUser)
	if err != nil || authenticatedUser.Id == "-1" {
		c.JSON(409, gin.H{
			"Message": "Login Failed",
			"Data":    authenticatedUser.UserWithoutPassword,
		})
	}
	c.JSON(200, gin.H{
		"Message": "Logged in Successfully",
		"Data":    authenticatedUser.UserWithoutPassword,
	})
}

func GetFreeDoctorSlot(c *gin.Context) {
	BasicAuth(c)
	var slot Model.GetListPatientResponse
	var arrSlot []Model.GetListPatientResponse

	db := Config.Connect()
	defer db.Close()

	rows, err := Repo.GetFreeDoctorSlots(db)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&slot.Id, &slot.DoctorName, &slot.Date)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrSlot = append(arrSlot, slot)
		}
	}

	fmt.Println(arrSlot)
	c.JSON(200, gin.H{
		"Message": "Got Data Successfully",
		"Data":    arrSlot,
	})

}
func AddDoctorSlot(c *gin.Context) {
	BasicAuth(c)
	db := Config.Connect()
	defer db.Close()

	user, _, _ := c.Request.BasicAuth()
	var slot Model.Slot
	c.Bind(&slot)

	u, _ := Repo.GetUserbyUserName(user, db)
	slot.DoctorId = u.Id

	_, err := Repo.InsertSlot(slot, db)
	if err != nil {
		c.JSON(409, gin.H{
			"Message": "Failed to Insert Data",
		})
		return
	}

	c.JSON(200, gin.H{
		"Message": "Insert data successfully",
	})
	fmt.Print("Insert data to database\n")

}

func GetDoctorSlots(c *gin.Context) {
	BasicAuth(c)
	var slot Model.DisplaySlot
	var arrSlot []Model.DisplaySlot

	user, _, _ := c.Request.BasicAuth()

	db := Config.Connect()
	defer db.Close()

	authenticatedUser, err := Repo.GetUserbyUserName(user, db)
	id := authenticatedUser.Id

	if err != nil {
		log.Print(err)
	}

	rows, err := Repo.GetDoctorSlots(id, db)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&slot.Id, &slot.Date)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrSlot = append(arrSlot, slot)
		}
	}

	fmt.Println(arrSlot)
	c.JSON(200, gin.H{
		"Message": "Got Data Successfully",
		"Data":    arrSlot,
	})
}

func AddPatientSlot(c *gin.Context) {
	BasicAuth(c)
	db := Config.Connect()
	defer db.Close()

	user, _, _ := c.Request.BasicAuth()
	u, _ := Repo.GetUserbyUserName(user, db)

	var slot Model.SlotWithPatient

	if err := c.Bind(&slot); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Invalid request body",
		})
		return
	}
	slot.PatientId = u.Id

	fmt.Println(slot)

	if _, err := Repo.InsertPatientSlot(slot, db); err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Failed to insert data",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Inserted Data Successfully",
	})
}

func GetPatientSlots(c *gin.Context) {
	var slot Model.GetListPatientResponse
	var arrSlot []Model.GetListPatientResponse
	BasicAuth(c)
	user, _, _ := c.Request.BasicAuth()

	db := Config.Connect()
	defer db.Close()

	authenticatedUser, err := Repo.GetUserbyUserName(user, db)

	if err != nil {
		log.Print(err)
	}
	id := authenticatedUser.Id
	rows, err := Repo.GetPatientSlots(id, db)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&slot.Id, &slot.DoctorName, &slot.Date)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrSlot = append(arrSlot, slot)
		}
	}
	fmt.Println(arrSlot)
	c.JSON(200, gin.H{
		"Message": "Got Data Successfully",
		"Data":    arrSlot,
	})
}

func DeleteSlot(c *gin.Context) {
	user, _, _ := c.Request.BasicAuth()
	BasicAuth(c)

	db := Config.Connect()
	defer db.Close()

	authenticatedUser, err := Repo.GetUserbyUserName(user, db)

	if err != nil {
		log.Print(err)
		return
	}
	id := c.Query("id")

	if authenticatedUser.Role == "patient" {
		_, err = Repo.DeletePatientSlot(id, db)
	} else {
		_, err = Repo.DeleteDoctorSlot(id, db)
	}

	if err != nil {
		log.Print(err)
		return
	}

	c.JSON(200, gin.H{
		"Message": "Deleted Data Successfully",
	})
}

func EditDoctorSlot(c *gin.Context) {
	BasicAuth(c)

	db := Config.Connect()
	defer db.Close()

	var slot Model.Slot

	err := c.Bind(&slot)
	if err != nil {
		log.Print(err)
		return
	}

	fmt.Println(slot)

	err = Repo.UpdateDoctorSlot(slot, db)

	if err != nil {
		log.Print(err)
	}

	c.JSON(200, gin.H{
		"Message": "Update data successfully",
	})

	fmt.Print("Update data successfully\n")

}

func EditPatientSlot(c *gin.Context) {
	BasicAuth(c)

	db := Config.Connect()
	defer db.Close()

	var slot Model.SlotWithPatient

	err := c.Bind(&slot)
	if err != nil {
		log.Print(err)
		return
	}

	fmt.Println(slot)

	err = Repo.UpdatePatientSlot(slot, db)

	if err != nil {
		log.Print(err)
	}

	c.JSON(200, gin.H{
		"Message": "Update data successfully",
	})

	fmt.Print("Update data successfully\n")
}
