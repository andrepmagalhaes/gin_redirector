package types

type GetLockersParams struct {
	Codigo_de_MSG                 string `json:"Codigo_de_MSG"`
	ID_de_Referencia              string `json:"ID_de_Referencia"`
	ID_do_Solicitante             string `json:"ID_do_Solicitante"`
	ID_Rede_Lockers               int    `json:"ID_Rede_Lockers"`
	Data_Hora_Solicitacao         string `json:"Data_Hora_Solicitacao"`
	Codigo_Pais_Locker            string `json:"Codigo_Pais_Locker"`
	Cidade_Locker                 string `json:"Cidade_Locker"`
	CEP_Locker                    string `json:"CEP_Locker"`
	Bairro_Locker                 string `json:"Bairro_Locker"`
	Endereco_Locker               string `json:"Endereco_Locker"`
	Numero_Locker                 string `json:"Numero_Locker"`
	Complemento_Locker            string `json:"Complemento_Locker"`
	Modelo_uso_Locker             string `json:"Modelo_uso_Locker"`
	Categoria_Locker              int    `json:"Categoria_Locker"`
	Tipo_Armazenamento            string `json:"Tipo_Armazenamento"`
	Codigo_Dimensao_Portas_Locker string `json:"Codigo_Dimensao_Portas_Locker"`
	Modelo_Operacao_Locker        string `json:"Modelo_Operacao_Locker"`
	ID_da_Estacao_do_Locker       string `json:"ID_da_Estacao_do_Locker"`
	LatLong                       string `json:"LatLong"`
	Versao_Mensageira             string `json:"Versao_Mensageira"`
}

type EstacaoLocker struct {
	Codigo_Pais_Locker      string `json:"Codigo_Pais_Locker"`
	Cidade_Locker           string `json:"Cidade_Locker"`
	Cep_Locker              string `json:"Cep_Locker"`
	Bairro_Locker           string `json:"Bairro_Locker"`
	Endereco_Locker         string `json:"Endereco_Locker"`
	Numero_Locker           string `json:"Numero_Locker"`
	Complemento_Locker      string `json:"Complemento_Locker"`
	Modelo_Uso_Locker       string `json:"Modelo_Uso_Locker"`
	Categoria_Locker        string `json:"Categoria_Locker"`
	Tipo_Armazenamento      string `json:"Tipo_Armazenamento"`
	Modelo_Operacao_Locker  string `json:"Modelo_Operacao_Locker"`
	Nome_Fantasia_do_locker string `json:"Nome_Fantasia_do_locker"`
	ID_da_estacao_do_locker string `json:"ID_da_estacao_do_locker"`
	LatLong                 string `json:"LatLong"`
	Segunda_Hora_Inicio     string `json:"Segunda_Hora_Inicio"`
	Segunda_Hora_Fim        string `json:"Segunda_Hora_Fim"`
	Terca_Hora_Inicio       string `json:"Terca_Hora_Inicio"`
	Terca_Hora_Fim          string `json:"Terca_Hora_Fim"`
	Quarta_Hora_Inicio      string `json:"Quarta_Hora_Inicio"`
	Quarta_Hora_Fim         string `json:"Quarta_Hora_Fim"`
	Quinta_Hora_Inicio      string `json:"Quinta_Hora_Inicio"`
	Quinta_Hora_Fim         string `json:"Quinta_Hora_Fim"`
	Sexta_Hora_Inicio       string `json:"Sexta_Hora_Inicio"`
	Sexta_Hora_Fim          string `json:"Sexta_Hora_Fim"`
	Sabado_Hora_Inicio      string `json:"Sabado_Hora_Inicio"`
	Sabado_Hora_Fim         string `json:"Sabado_Hora_Fim"`
	Domingo_Hora_Inicio     string `json:"Domingo_Hora_Inicio"`
	Domingo_Hora_Fim        string `json:"Domingo_Hora_Fim"`
	Feriados_Hora_Inicio    string `json:"Feriados_Hora_Inicio"`
	Feriados_Hora_Fim       string `json:"Feriados_Hora_Fim"`
}

type GetLockersResponse struct {
	Codigo_de_MSG      string          `json:"Codigo_de_MSG"`
	ID_de_Referencia   string          `json:"ID_de_Referencia"`
	ID_Rede_Lockers    string          `json:"ID_Rede_Lockers"`
	Versao_Mensageira  string          `json:"Versao_Mensageira"`
	Data_Hora_Resposta string          `json:"Data_Hora_Resposta"`
	Estacao_Locker     []EstacaoLocker `json:"Estacao_Locker"`
}
