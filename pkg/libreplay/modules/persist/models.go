package persist

type GenericResponse struct {
	Message interface{} `json:"message"`
	Error   *string     `json:"error,omitempty"`
}
