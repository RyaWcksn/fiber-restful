package forms

type ErrorForm struct {
	HttpCode int    `json:"code"`
	Message  string `json:"message"`
	Error    string `json:"error"`
}
