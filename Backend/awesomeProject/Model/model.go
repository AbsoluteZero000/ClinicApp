package Model

type UserWithoutPassword struct {
	Id    string `form:"id" json:"id"`
	Name  string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
	Role  string `form:"role" json:"role"`
}

type User struct {
	UserWithoutPassword
	Password string `form:"password" json:"password"`
}

type Slot struct {
	Id       string `form:"id" json:"id"`
	DoctorId string `form:"doctor" json:"doctor"`
	Date     string `form:"date" json:"date"`
}

type SlotWithPatient struct {
	Id        string `form:"id" json:"id"`
	PatientId string `form:"patient" json:"patient"`
	SlotId    string `form:"slot" json:"slot"`
}

type DisplaySlot struct {
	Id   string `form:"id" json:"id"`
	Date string `form:"date" json:"date"`
}
type DisplaySlotResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []DisplaySlot
}
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []UserWithoutPassword
}

type DoctorSlotResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Slot
}
type PatientSlotResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []SlotWithPatient
}
type GetListPatientResponse struct {
	Id         string `json:"id"`
	Date       string `json:"date"`
	DoctorName string `json:"name"`
}

type PatientNamesSlotResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []GetListPatientResponse
}
