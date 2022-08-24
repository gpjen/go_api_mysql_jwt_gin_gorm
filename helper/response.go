package helper

type responseOK struct {
	Status  bool        `json:"status"`
	Message string      `josn:"message"`
	Data    interface{} `json:"data"`
}

type responseFail struct {
	Status  bool        `json:"status"`
	Message string      `josn:"message"`
	Errors  interface{} `json:"errors"`
}

// used when data doesnt want to be a null json
type EpmptyObj struct{}

func ResponseOK(message string, data interface{}) responseOK {
	return responseOK{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func ResponseFail(message string, err interface{}) responseFail {

	return responseFail{
		Status:  false,
		Message: message,
		Errors:  err,
	}
}
