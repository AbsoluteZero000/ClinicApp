package Controller

import (
	"awesomeProject/Config"
	"awesomeProject/Model"
	"awesomeProject/Repo"
	"encoding/json"
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var Store = sessions.NewCookieStore(securecookie.GenerateRandomKey(10))

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
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var User Model.User
	var response Model.Response
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

	response.Status = 200
	response.Message = "Success"
	response.Data = arrUser

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// SignUp = Insert User API
func SignUp(w http.ResponseWriter, r *http.Request) {
	var response Model.Response

	db := Config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}

	var user Model.User
	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")
	user.Role = r.FormValue("role")

	_, err = Repo.InsertUser(user, db)

	if err != nil {
		log.Print(err)
		response.Status = 409
		response.Message = "Failed to insert data duplicate Email or Name"
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Status = 200
	response.Message = "Insert data successfully"
	fmt.Print("Insert data to database\n")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var response Model.Response

	db := Config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}
	var user Model.User
	session, _ := Store.Get(r, "Session")
	user.Id = session.Values["id"].(string)
	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")
	user.Role = r.FormValue("role")

	err = Repo.UpdateUser(user, db)

	if err != nil {
		log.Print(err)
	}

	response.Status = 200
	response.Message = "Update data successfully"
	fmt.Print("Update data successfully\n")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var response Model.Response
	db := Config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")

	_, err = Repo.DeleteUser(id, db)

	if err != nil {
		log.Print(err)
		return
	}

	response.Status = 200
	response.Message = "Delete data successfully\n"
	fmt.Print("Delete data successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ValidateAccount(name string, password string) Model.UserWithoutPassword {
	var user Model.User

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
		} else {
			fmt.Printf("Checking user: %s, Password: %s\n", user.Name, user.Password)
			if password == user.Password && name == user.Name {
				fmt.Println("User found and authenticated!")
				return user.UserWithoutPassword
			}
		}
	}

	fmt.Println("User not found or not authenticated.")
	return Model.UserWithoutPassword{Id: "-1", Name: "", Email: ""}

}

func Login(w http.ResponseWriter, r *http.Request) {
	user, pass, _ := r.BasicAuth()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Set-Cookie")

	var response Model.Response
	// Validate the user's credentials and retrieve the user's role.
	authenticatedUser := ValidateAccount(user, pass)

	if authenticatedUser.Id == "-1" {

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Create a session and set the role and authenticated flag.
	session, _ := Store.Get(r, "Session")
	session.Values["authenticated"] = true
	session.Values["role"] = authenticatedUser.Role
	session.Values["id"] = authenticatedUser.Id
	session.Save(r, w)

	response.Data = append(response.Data, authenticatedUser)
	response.Status = 200
	response.Message = "Success"

	json.NewEncoder(w).Encode(response)
}

func Signout(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "Session") // Change "session-name" to your session name
	session.Values["authenticated"] = false
	session.Save(r, w)

	// Redirect the user to the login page or any other appropriate page
	http.Redirect(w, r, "/login", http.StatusFound)
}

func AddSlot(w http.ResponseWriter, r *http.Request) {
	var response Model.Response

	db := Config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}

	var slot Model.Slot
	slot.DoctorId = r.FormValue("doctor")
	slot.Date = r.FormValue("date")

	_, err = Repo.InsertSlot(slot, db)

	if err != nil {
		log.Print(err)
		response.Status = 409
		response.Message = "idk smth weird happened"
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Status = 200
	response.Message = "Insert data successfully"
	fmt.Print("Insert data to database\n")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func GetDoctorSlots(w http.ResponseWriter, r *http.Request) {
	var slot Model.Slot
	var response Model.DoctorSlotResponse
	var arrSlot []Model.Slot
	session, _ := Store.Get(r, "Session")
	var id = session.Values["id"]

	db := Config.Connect()
	defer db.Close()

	rows, err := Repo.GetDoctorSlots(id.(string), db)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&slot.Id, &slot.DoctorId, &slot.Date)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrSlot = append(arrSlot, slot)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrSlot

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func AddPatientSlot(w http.ResponseWriter, r *http.Request) {
	var response Model.Response

	db := Config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}

	var slot Model.SlotWithPatient
	session, _ := Store.Get(r, "Session")
	slot.PatientId = session.Values["id"].(string)
	slot.SlotId = r.FormValue("slot")

	_, err = Repo.InsertPatientSlot(slot, db)

	if err != nil {
		log.Print(err)
		response.Status = 409
		response.Message = "idk smth weird happened"
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Status = 200
	response.Message = "Insert data successfully"
	fmt.Print("Insert data to database\n")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)

}
func GetPatientSlots(w http.ResponseWriter, r *http.Request) {
	var slot Model.GetListPatientResponse
	var response Model.PatientNamesSlotResponse
	var arrSlot []Model.GetListPatientResponse

	session, _ := Store.Get(r, "Session")
	id := session.Values["id"].(string)

	db := Config.Connect()
	defer db.Close()

	rows, err := Repo.GetPatientSlots(id, db)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&slot.DoctorName, &slot.Date)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrSlot = append(arrSlot, slot)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrSlot

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func DeleteSlot(w http.ResponseWriter, r *http.Request) {
	var response Model.Response
	db := Config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")

	session, _ := Store.Get(r, "Session")

	if session.Values["role"] == "patient" {
		_, err = Repo.DeletePatientSlot(id, db)
	} else {
		_, err = Repo.DeleteDoctorSlot(id, db)
	}

	if err != nil {
		log.Print(err)
		return
	}

	response.Status = 200
	response.Message = "Delete data successfully\n"
	fmt.Print("Delete data successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func EditDoctorSlot(w http.ResponseWriter, r *http.Request) {
	var response Model.DoctorSlotResponse

	db := Config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}
	var slot Model.Slot

	slot.Id = r.FormValue("id")
	slot.Date = r.FormValue("date")

	err = Repo.UpdateDoctorSlot(slot, db)

	if err != nil {
		log.Print(err)
	}

	response.Status = 200
	response.Message = "Update data successfully"
	fmt.Print("Update data successfully\n")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func EditPatientSlot(w http.ResponseWriter, r *http.Request) {
	var response Model.PatientSlotResponse

	db := Config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}
	var slot Model.SlotWithPatient

	session, _ := Store.Get(r, "Session")
	slot.Id = session.Values["id"].(string)
	slot.SlotId = r.FormValue("slot")

	err = Repo.UpdatePatientSlot(slot, db)

	if err != nil {
		log.Print(err)
	}

	response.Status = 200
	response.Message = "Update data successfully"
	fmt.Print("Update data successfully\n")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
