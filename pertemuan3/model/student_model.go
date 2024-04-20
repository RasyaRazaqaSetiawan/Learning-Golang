package model

type Student struct {
	StudentID int64  `json:"student_id"`
	Name      string `json:"name"`
	Age       int64  `json:"age"`
	AddressID string `json:"address"`
	HobbyID   string `json:"hobby"`
}
