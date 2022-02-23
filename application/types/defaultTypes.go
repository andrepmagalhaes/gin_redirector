package types

type Unathorized struct {
	Status_code int    `json:"status_code,omitempty"`
	Detail      string `json:"detail,omitempty"`
}
