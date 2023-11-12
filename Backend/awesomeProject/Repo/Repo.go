package Repo

import (
	"awesomeProject/Model"
	"database/sql"
	"fmt"
)

func InsertUser(user Model.User, db *sql.DB) (sql.Result, error) {
	return db.Exec("INSERT INTO User(name, email, password, role) VALUES(?, ?, ?, ?)", user.Name, user.Email, user.Password, user.Role)
}

func GetAllUser(db *sql.DB) (*sql.Rows, error) {
	return db.Query("SELECT id, name, email, password, role FROM User")
}

func DeleteUser(id string, db *sql.DB) (sql.Result, error) {
	return db.Exec("DELETE FROM User WHERE id=?", id)
}

func GetUserbyUserName(name string, db *sql.DB) (*Model.User, error) {
	var user Model.User
	err := db.QueryRow("select * from user where name =?", name).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}
func UpdateUser(user Model.User, db *sql.DB) error {
	var err error

	if user.Name != "" {
		_, err = db.Exec("UPDATE User SET name=? WHERE id=?", user.Name, user.Id)
	}
	if user.Password != "" {
		_, err = db.Exec("UPDATE User SET password=? WHERE id=?", user.Password, user.Id)
	}
	if user.Email != "" {
		_, err = db.Exec("UPDATE User SET email=? WHERE id=?", user.Email, user.Id)
	}
	if user.Role != "" {

		_, err = db.Exec("UPDATE User SET role=? WHERE id=?", user.Role, user.Id)
	}
	return err
}

// 9999-12-31 23:59:59
func InsertSlot(slot Model.Slot, db *sql.DB) (sql.Result, error) {
	return db.Exec("INSERT INTO Slot(doctorid, date) VALUES(?,?)", slot.DoctorId, slot.Date)
}

func GetSlots(db *sql.DB) (*sql.Rows, error) {
	return db.Query("Select * From Slot")

}

func GetDoctorSlots(id string, db *sql.DB) (*sql.Rows, error) {
	return db.Query("select id, date from Slot where doctorid=?", id)

}

func GetPatientSlots(id string, db *sql.DB) (*sql.Rows, error) {
	return db.Query("SELECT slotwithpatient.id, user.name, Slot.date  FROM user  JOIN SlotWithPatient ON user.id = SlotWithPatient.patientid  JOIN Slot ON SlotWithPatient.slotid = Slot.id WHERE user.id = ?;", id)

}

func InsertPatientSlot(slot Model.SlotWithPatient, db *sql.DB) (sql.Result, error) {
	return db.Exec("INSERT INTO SlotWithPatient(patientid, slotid) VALUES(?,?)", slot.PatientId, slot.SlotId)
}

func DeletePatientSlot(id string, db *sql.DB) (sql.Result, error) {
	return db.Exec("DELETE FROM slotwithpatient WHERE id=?", id)
}

func DeleteDoctorSlot(id string, db *sql.DB) (sql.Result, error) {
	return db.Exec("DELETE FROM slot WHERE id=?", id)
}

func UpdateDoctorSlot(slot Model.Slot, db *sql.DB) error {
	var err error
	if slot.Date != "" {
		_, err = db.Exec("UPDATE slot SET date=? WHERE id=?", slot.Date, slot.Id)
	}
	return err
}
func UpdatePatientSlot(slot Model.SlotWithPatient, db *sql.DB) error {
	var err error
	if slot.SlotId != "" {
		_, err = db.Exec("UPDATE slot SET slotid=? WHERE id=?", slot.SlotId, slot.Id)
	}
	return err
}
