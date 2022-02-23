package types

type GetLockersSlotsParams struct {
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
	Versao_Mensageira             string `json:"Versao_Mensageira,omitempty"`
}

type portas struct {
	ID_da_Porta_do_Locker         string `json:"ID_da_Porta_do_Locker,omitempty"`
	IdLockerPortaFisica           int    `json:"idLockerPortaFisica,omitempty"`
	Categoria_Porta               string `json:"Categoria_Porta,omitempty"`
	Modelo_Uso_Porta              string `json:"Modelo_Uso_Porta,omitempty"`
	LockerPortaOperacao           string `json:"LockerPortaOperacao,omitempty"`
	Codigo_Dimensao_Portas_Locker string `json:"Codigo_Dimensao_Portas_Locker,omitempty"`
	Comprimento_Porta             string `json:"Comprimento_Porta,omitempty"`
	Largura_Porta                 int    `json:"Largura_Porta,omitempty"`
	Altura_Porta                  int    `json:"Altura_Porta,omitempty"`
	Peso_Maximo_Porta             int    `json:"Peso_Maximo_Porta,omitempty"`
}

type lockers struct {
	Id_da_Estacao_do_Locker string   `json:"Id_da_Estacao_do_Locker,omitempty"`
	Codigo_Pais_Locker      string   `json:"Codigo_Pais_Locker,omitempty"`
	CEP_Locker              string   `json:"CEP_Locker,omitempty"`
	Cidade_Locker           string   `json:"Cidade_Locker,omitempty"`
	Bairro_Locker           string   `json:"Bairro_Locker,omitempty"`
	Endereco_Locker         string   `json:"Endereco_Locker,omitempty"`
	Numero_Locker           string   `json:"Numero_Locker,omitempty"`
	Complemento_Locker      string   `json:"Complemento_Locker,omitempty"`
	Nome_Fantasia_do_locker string   `json:"Nome_Fantasia_do_locker,omitempty"`
	LockerUso               string   `json:"LockerUso,omitempty"`
	Categoria_Locker        string   `json:"Categoria_Locker,omitempty"`
	Tipo_Armazenamento      string   `json:"Tipo_Armazenamento,omitempty"`
	Modelo_Operacao_Locker  string   `json:"Modelo_Operacao_Locker,omitempty"`
	LatLong                 string   `json:"LatLong,omitempty"`
	Portas                  []portas `json:"Portas,omitempty"`
}

type GetLockersSlotsResponse struct {
	Codigo_de_MSG      string    `json:"Codigo_de_MSG,omitempty"`
	ID_de_Referencia   string    `json:"ID_de_Referencia,omitempty"`
	ID_do_Solicitante  string    `json:"ID_do_Solicitante,omitempty"`
	ID_Rede_Lockers    string    `json:"ID_Rede_Lockers,omitempty"`
	Data_Hora_Resposta string    `json:"Data_Hora_Resposta,omitempty"`
	Lockers            []lockers `json:"Lockers,omitempty"`
	Versao_Mensageria  string    `json:"Versao_Mensageria,omitempty"`
}

// type GetLockersSlotsIdInvalid struct {
// 	Status_code int `json:"Status_code,omitempty"`

// }
