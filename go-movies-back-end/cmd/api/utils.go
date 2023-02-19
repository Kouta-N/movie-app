
type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"error"`
	Data    interface{} `json:"error"`
}

func (app *application) witeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {

}