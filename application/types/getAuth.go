package types

type GetAuthParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetAuthResponse struct {
	Token             string `json:"token"`
	ID_do_Solicitante string `json:"ID_do_Solicitante"`
}
