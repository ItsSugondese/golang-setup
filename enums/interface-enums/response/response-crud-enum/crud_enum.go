package response_crud_enum

import (
	"encoding/json"
	"fmt"
)

//var CrudEnum = newCrud()

//func newCrud() *crud {
//	return &crud{
//		CREATE: "CREATE",
//		GET:    "GET",
//		UPDATE: "UPDATE",
//		DELETE: "DELETE",
//		ERROR:  "ERROR",
//	}
//}

//type crud struct {
//	CREATE string
//	GET    string
//	UPDATE string
//	DELETE string
//	ERROR  string
//}

type CrudEnum interface {
	fmt.Stringer
	private()
}
type crud struct {
	Str string
}

func (c crud) String() string {
	return c.Str
}

func (c crud) private() {
}

// MarshalJSON method to handle custom marshaling to JSON
func (rs crud) MarshalJSON() ([]byte, error) {
	return json.Marshal(rs.Str)
}

func Create() CrudEnum {
	return crud{
		"CREATE",
	}
}

func Update() CrudEnum {
	return crud{
		"UPDATE",
	}
}

func Delete() CrudEnum {
	return crud{
		"DELETE",
	}
}

func Get() CrudEnum {
	return crud{
		"GET",
	}
}

func Error() CrudEnum {
	return crud{
		"ERROR",
	}
}
