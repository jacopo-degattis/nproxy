package lib

import (
	"encoding/json"
	"io"
	"net/http"
)

// Make `dstMap` inherit all fields from parent map `parentMap`
func mapInherit(parentMap map[string]any, dstMap map[string]any) {
	for key, value := range parentMap {
		dstMap[key] = value
	}
}

// This function should be moved inside the middlewares package
func Fetch(url string, method string, body io.Reader, header http.Header) ([]byte, error) {
	client := http.Client{}

	req, err := http.NewRequest(method, url, body)

	// token := header.Get("x-nd-authorization")

	// fmt.Printf("TOKEN = %s", token)

	// if token != "" {
	// 	req.Header.Add("x-nd-authorization", token)
	// }

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	return resBody, nil
}

func CheckError(w http.ResponseWriter, err error) {
	if err != nil {
		responseError := BuildSubsonicError(0, err.Error())
		w.Header().Set("Content-Type", "application/json")
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
