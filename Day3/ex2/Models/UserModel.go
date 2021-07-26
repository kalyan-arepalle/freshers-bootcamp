//Models/UserModel.go
package Models

type Student struct {
	Id      uint   `json:"id"`
	FirstName    string `json:"first-name"`
	LastName    string `json:"last-name"`
	DOB   string `json:"dob"`
	Address string `json:"address"`
	Subject string `json:"subject"`
	Marks	int `json:"marks"`
	//SubjectMarks   []SubjectMark `json:"scores" gorm:"foreignKey:Id"`
}
/*
type SubjectMark struct {
	Id      int    `json:"id"`
	Subject string `json:"subject"`
	Marks   int    `json:"marks"`
}
*/
func (b *Student) TableName() string {
	return "student"
}
/*
func (b *SubjectMark) TableName() string {
	return "Subject Marks"
}
*/