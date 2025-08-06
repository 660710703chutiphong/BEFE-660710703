package main
import("fmt"
	"errors")

type Student struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Year int `json:"year"`
	Gpa float64 `json:"gpa"`
}
func (s *Student) IsHonor() bool{
	return s.Gpa >= 3.50;
}
func (s *Student) Validate() error{
	if s.Name == "" {return errors.New("name is required")}
	if s.Year < 1 || s.Year > 4{return errors.New("Year must be between 1-4")}
	if s.Gpa < 0 || s.Gpa > 4{return errors.New("gpa must be between 0-4")}
	return nil;
}
func main(){
	// var st Student = Student{ID:"1",Name:"Chutiphong",Email:"Buranakunkikan_c@su.ac.th",Year:3,Gpa:2.77}

	//st Student2 := Student((ID:"1",Name:"Chutiphong",Email:"Buranakunkikan_c@su.ac.th",Year:3,Gpa:2.77))
	students []Student{
		{ID:"1",Name:"Chutiphong",Email:"Buranakunkikan_c@su.ac.th",Year:3,Gpa:2.77},
		{ID:"2",Name:"Alice",Email:"alice@su.ac.th",Year:1,Gpa:2.00}}

	newstudent := Student{ID:"3",Name:"God",Email:"Godlnwza007x@god.email.com",Year:(9999999999-9999999995),Gpa:2.77}
	
	students = append(students, newstudent)
	
	for i, students:= range students{
		fmt.Printf("%d Honor: %v\n",i,st.IsHonor())
		fmt.Printf("%d Vaildation: %v\n",i,st.Validate())
	}
	// for _, students:= range students{
	// 	fmt.Printf("Honor: %v\n",st.IsHonor())
	// 	fmt.Printf("Vaildation: %v\n",st.Validate())
	// }
}