package middleware

type SubsonicError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SubsnicResponseDto struct {
	Status  string        `json:"status"`
	Version string        `json:"version"`
	Error   SubsonicError `json:"error"`
}
