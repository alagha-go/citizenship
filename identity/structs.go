package identity


type Identity struct {
	Name        						string 						`json:"name" bson:"name"`       
	Status      						string 						`json:"status" bson:"status"`     
	Citizenship 						string 						`json:"citizenship" bson:"citizenship"`
	Dob         						string 						`json:"dob" bson:"dob"`        
	FirstName   						string 						`json:"first_name" bson:"first_name"` 
	Gender      						string 						`json:"gender" bson:"gender"`     
	IDNumber    						string 						`json:"id_number" bson:"id_number"`  
	OtherName   						string 						`json:"other_name" bson:"other_name"` 
	SerialNo    						string 						`json:"serial_no" bson:"serial_no"`  
	Photo       						string 						`json:"photo" bson:"photo"`      
	Surname     						string 						`json:"surname" bson:"surname"`    
	Valid       						bool   						`json:"valid" bson:"valid"`   
}