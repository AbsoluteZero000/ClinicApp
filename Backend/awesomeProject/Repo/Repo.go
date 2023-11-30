package Repo

import (
	"awesomeProject/Model"
	"database/sql"
	"fmt"
)

func InsertUser(user Model.User, db *sql.DB) (sql.Result, error) {
	return db.Exec("INSERT INTO user(name, email, password, role) VALUES(?, ?, ?, ?)", user.Name, user.Email, user.Password, user.Role)
}

func GetAllUser(db *sql.DB) (*sql.Rows, error) {
	return db.Query("SELECT id, name, email, password, role FROM user")
}

func DeleteUser(id string, db *sql.DB) (sql.Result, error) {
	return db.Exec("DELETE FROM user WHERE id=?", id)
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
		_, err = db.Exec("UPDATE user SET name=? WHERE id=?", user.Name, user.Id)
	}
	if user.Password != "" {
		_, err = db.Exec("UPDATE user SET password=? WHERE id=?", user.Password, user.Id)
	}
	if user.Email != "" {
		_, err = db.Exec("UPDATE user SET email=? WHERE id=?", user.Email, user.Id)
	}
	if user.Role != "" {

		_, err = db.Exec("UPDATE user SET role=? WHERE id=?", user.Role, user.Id)
	}
	return err
}

// 9999-12-31 23:59:59
func InsertSlot(slot Model.Slot, db *sql.DB) (sql.Result, error) {
	return db.Exec("INSERT INTO slot(doctorid, date) VALUES(?,?)", slot.DoctorId, slot.Date)
}

func GetSlots(db *sql.DB) (*sql.Rows, error) {
	return db.Query("Select * From slot")

}

func GetDoctorSlots(id string, db *sql.DB) (*sql.Rows, error) {
	return db.Query("select id, date from slot where doctorid=?", id)

}
func GetFreeDoctorSlots(db *sql.DB) (*sql.Rows, error) {
	return db.Query("SELECT s.id AS slot_id, u.name AS doctor_name, s.date FROM slot s JOIN user u ON s.doctorid = u.id LEFT JOIN SlotWithPatient sp ON s.id = sp.slotid WHERE sp.id IS NULL;")
}

func GetPatientSlots(id string, db *sql.DB) (*sql.Rows, error) {
	return db.Query("SELECT SlotWithPatient.id, user.name, slot.date  FROM user  JOIN SlotWithPatient ON user.id = SlotWithPatient.patientid  JOIN slot ON SlotWithPatient.slotid = slot.id WHERE user.id = ?;", id)

}

func InsertPatientSlot(slot Model.SlotWithPatient, db *sql.DB) (sql.Result, error) {
	return db.Exec("INSERT INTO SlotWithPatient(patientid, slotid) VALUES(?,?)", slot.PatientId, slot.SlotId)
}

func DeletePatientSlot(id string, db *sql.DB) (sql.Result, error) {
	return db.Exec("DELETE FROM SlotWithPatient WHERE id=?", id)
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
