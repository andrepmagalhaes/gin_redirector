package types

type GetAuthParams struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type GetAuthResponse struct {
	Token             string `json:"token,omitempty"`
	ID_do_Solicitante string `json:"ID_do_Solicitante,omitempty"`
}

type GetAuth400 struct {
	Status_code int    `json:"status_code,omitempty"`
	Details     string `json:"detail,omitempty"`
}
