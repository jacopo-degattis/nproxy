package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"nproxy/config"
	"nproxy/server"
)

// /auth/login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fullUrl := fmt.Sprintf("%s%s", config.GetNavidromeUrl(), "/auth/login")

	defer r.Body.Close()
	res, err := server.Fetch(fullUrl, "POST", r.Body, r.Header, nil)

	if err != nil {
		error := server.BuildSubsonicError(0, err.Error())
		w.WriteHeader(500)
		w.Write(error)
		return
	}

	rawBody := make(map[string]any)
	err = json.NewDecoder(bytes.NewReader(res)).Decode(&rawBody)

	if err != nil {
		error := server.BuildSubsonicError(0, err.Error())
		w.WriteHeader(500)
		w.Write(error)
		return
	}

	var token = rawBody["token"].(string)
	ctx := context.WithValue(r.Context(), "token", token)

	r = r.WithContext(ctx)

	w.Write(res)
}
