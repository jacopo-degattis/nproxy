package lib

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Make `dstMap` inherit all fields from parent map `parentMap`
func mapInherit(parentMap map[string]any, dstMap map[string]any) {
	for key, value := range parentMap {
		dstMap[key] = value
	}
}

// This function should be moved inside the middlewares package
// TODO: improve function stability and reliability
func Fetch(
	url string,
	method string,
	body io.Reader,
	header http.Header,
	params url.Values,
) ([]byte, error) {
	client := http.Client{}

	req, err := http.NewRequest(method, url, body)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36")

	// add params if present
	if params != nil {
		currentQuery := req.URL.Query()
		for k, v := range params {
			currentQuery.Add(k, v[0])
		}
		req.URL.RawQuery = currentQuery.Encode()
	}

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

// TODO: replace all if != nil in the code calling this function instead
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
