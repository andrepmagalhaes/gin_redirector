package types

type GetLockersParams struct {
	Codigo_de_MSG                 string `json:"Codigo_de_MSG,omitempty"`
	ID_de_Referencia              string `json:"ID_de_Referencia,omitempty"`
	ID_do_Solicitante             string `json:"ID_do_Solicitante,omitempty"`
	ID_Rede_Lockers               int    `json:"ID_Rede_Lockers,omitempty"`
	Data_Hora_Solicitacao         string `json:"Data_Hora_Solicitacao,omitempty"`
	Codigo_Pais_Locker            string `json:"Codigo_Pais_Locker,omitempty"`
	Cidade_Locker                 string `json:"Cidade_Locker,omitempty"`
	CEP_Locker                    string `json:"CEP_Locker,omitempty"`
	Bairro_Locker                 string `json:"Bairro_Locker,omitempty"`
	Endereco_Locker               string `json:"Endereco_Locker,omitempty"`
	Numero_Locker                 string `json:"Numero_Locker,omitempty"`
	Complemento_Locker            string `json:"Complemento_Locker,omitempty"`
	Modelo_uso_Locker             string `json:"Modelo_uso_Locker,omitempty"`
	Categoria_Locker              int    `json:"Categoria_Locker,omitempty"`
	Tipo_Armazenamento            string `json:"Tipo_Armazenamento,omitempty"`
	Codigo_Dimensao_Portas_Locker string `json:"Codigo_Dimensao_Portas_Locker,omitempty"`
	Modelo_Operacao_Locker        string `json:"Modelo_Operacao_Locker,omitempty"`
	ID_da_Estacao_do_Locker       string `json:"ID_da_Estacao_do_Locker,omitempty"`
	LatLong                       string `json:"LatLong,omitempty"`
	Versao_Mensageira             string `json:"Versao_Mensageria,omitempty"`
}

type estacaoLocker struct {
	Codigo_Pais_Locker      string `json:"Codigo_Pais_Locker,omitempty"`
	Cidade_Locker           string `json:"Cidade_Locker,omitempty"`
	Cep_Locker              string `json:"Cep_Locker,omitempty"`
	Bairro_Locker           string `json:"Bairro_Locker,omitempty"`
	Endereco_Locker         string `json:"Endereco_Locker,omitempty"`
	Numero_Locker           string `json:"Numero_Locker,omitempty"`
	Complemento_Locker      string `json:"Complemento_Locker,omitempty"`
	Modelo_Uso_Locker       string `json:"Modelo_Uso_Locker,omitempty"`
	Categoria_Locker        string `json:"Categoria_Locker,omitempty"`
	Tipo_Armazenamento      string `json:"Tipo_Armazenamento,omitempty"`
	Modelo_Operacao_Locker  string `json:"Modelo_Operacao_Locker,omitempty"`
	Nome_Fantasia_do_locker string `json:"Nome_Fantasia_do_locker,omitempty"`
	ID_da_estacao_do_locker string `json:"ID_da_estacao_do_locker,omitempty"`
	LatLong                 string `json:"LatLong,omitempty"`
	Segunda_Hora_Inicio     string `json:"Segunda_Hora_Inicio,omitempty"`
	Segunda_Hora_Fim        string `json:"Segunda_Hora_Fim,omitempty"`
	Terca_Hora_Inicio       string `json:"Terca_Hora_Inicio,omitempty"`
	Terca_Hora_Fim          string `json:"Terca_Hora_Fim,omitempty"`
	Quarta_Hora_Inicio      string `json:"Quarta_Hora_Inicio,omitempty"`
	Quarta_Hora_Fim         string `json:"Quarta_Hora_Fim,omitempty"`
	Quinta_Hora_Inicio      string `json:"Quinta_Hora_Inicio,omitempty"`
	Quinta_Hora_Fim         string `json:"Quinta_Hora_Fim,omitempty"`
	Sexta_Hora_Inicio       string `json:"Sexta_Hora_Inicio,omitempty"`
	Sexta_Hora_Fim          string `json:"Sexta_Hora_Fim,omitempty"`
	Sabado_Hora_Inicio      string `json:"Sabado_Hora_Inicio,omitempty"`
	Sabado_Hora_Fim         string `json:"Sabado_Hora_Fim,omitempty"`
	Domingo_Hora_Inicio     string `json:"Domingo_Hora_Inicio,omitempty"`
	Domingo_Hora_Fim        string `json:"Domingo_Hora_Fim,omitempty"`
	Feriados_Hora_Inicio    string `json:"Feriados_Hora_Inicio,omitempty"`
	Feriados_Hora_Fim       string `json:"Feriados_Hora_Fim,omitempty"`
}

type GetLockersResponse struct {
	Codigo_de_MSG      string          `json:"Codigo_de_MSG,omitempty"`
	ID_de_Referencia   string          `json:"ID_de_Referencia,omitempty"`
	ID_Rede_Lockers    int             `json:"ID_Rede_Lockers,omitempty"`
	Versao_Mensageira  string          `json:"Versao_Mensageira,omitempty"`
	Data_Hora_Resposta string          `json:"Data_Hora_Resposta,omitempty"`
	Estacao_Locker     []estacaoLocker `json:"Estacao_Locker,omitempty"`
}

type GetLockersResponse422 struct {
	Status_code int    `json:"status_code,omitempty"`
	Detail      string `json:"detail,omitempty"`
}
