package middleware

import (
	"encoding/json"
	"net/http"
)

func CheckError(w http.ResponseWriter, err error) {
	if err != nil {
		responseError := BuildSubsonicError(0, err.Error())
		w.Header().Set("Conten-Type", "application/json")
		w.Write(responseError)
	}
}

func BuildSubsonicError(code int, message string) []byte {
	errPayload := SubsnicResponseDto{
		Status:  "failed",
		Version: "1.16.1",
		Error: SubsonicError{
			Code:    code,
			Message: message,
		},
	}
	encoded, _ := json.Marshal(errPayload)

	return encoded
}
