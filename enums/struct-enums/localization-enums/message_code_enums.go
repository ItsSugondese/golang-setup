package localization_enums

var MessageCodeEnums = newMessageCode()

func newMessageCode() *messageCode {
	return &messageCode{
		SAVE:          "save",
		API_OPERATION: "api.operation",
	}
}

type messageCode struct {
	SAVE          string
	API_OPERATION string
}
