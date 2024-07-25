package user_type_enums

var UserType = newUserType()

func newUserType() *userType {
	return &userType{
		USER:    "BASE_USER",
		COMPANY: "COMPANY",
		DRIVER:  "DRIVER",
	}
}

type userType struct {
	USER    string
	COMPANY string
	DRIVER  string
}
