package types

type GetLockersSlotsParams struct {
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

type portas struct {
	ID_da_Porta_do_Locker         string `json:"ID_da_Porta_do_Locker"`
	IdLockerPortaFisica           int    `json:"idLockerPortaFisica"`
	Categoria_Porta               string `json:"Categoria_Porta"`
	Modelo_Uso_Porta              string `json:"Modelo_Uso_Porta"`
	LockerPortaOperacao           string `json:"LockerPortaOperacao"`
	Codigo_Dimensao_Portas_Locker string `json:"Codigo_Dimensao_Portas_Locker"`
	Comprimento_Porta             string `json:"Comprimento_Porta"`
	Largura_Porta                 int    `json:"Largura_Porta"`
	Altura_Porta                  int    `json:"Altura_Porta"`
	Peso_Maximo_Porta             int    `json:"Peso_Maximo_Porta"`
}

type lockers struct {
	Id_da_Estacao_do_Locker string   `json:"Id_da_Estacao_do_Locker"`
	Codigo_Pais_Locker      string   `json:"Codigo_Pais_Locker"`
	CEP_Locker              string   `json:"CEP_Locker"`
	Cidade_Locker           string   `json:"Cidade_Locker"`
	Bairro_Locker           string   `json:"Bairro_Locker"`
	Endereco_Locker         string   `json:"Endereco_Locker"`
	Numero_Locker           string   `json:"Numero_Locker"`
	Complemento_Locker      string   `json:"Complemento_Locker"`
	Nome_Fantasia_do_locker string   `json:"Nome_Fantasia_do_locker"`
	LockerUso               string   `json:"LockerUso"`
	Categoria_Locker        string   `json:"Categoria_Locker"`
	Tipo_Armazenamento      string   `json:"Tipo_Armazenamento"`
	Modelo_Operacao_Locker  string   `json:"Modelo_Operacao_Locker"`
	LatLong                 string   `json:"LatLong"`
	Portas                  []portas `json:"Portas"`
}

type GetLockersSlotsResponse struct {
	Codigo_de_MSG      string    `json:"Codigo_de_MSG"`
	ID_de_Referencia   string    `json:"ID_de_Referencia"`
	ID_do_Solicitante  string    `json:"ID_do_Solicitante"`
	ID_Rede_Lockers    string    `json:"ID_Rede_Lockers"`
	Data_Hora_Resposta string    `json:"Data_Hora_Resposta"`
	Lockers            []lockers `json:"Lockers"`
}
