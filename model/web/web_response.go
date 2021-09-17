package web

type WebResponse struct {
	Code   uint16
	Status string
	Data   interface{}
}
