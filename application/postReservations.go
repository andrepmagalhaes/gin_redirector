package types

type infoEncomendas struct {
	ID_Encomenda                         int    `json:"ID_Encomenda,omitempty"`
	Numero_Mobile_Shopper                string `json:"Numero_Mobile_Shopper,omitempty"`
	Endereco_de_Email_do_Shopper         string `json:"Endereco_de_Email_,omitempty"`
	CPF_CNPJ_Shopper                     string `json:"CPF_CNPJ_Shopper,omitempty"`
	Moeda_Shopper                        string `json:"Moeda_Sho,omitempty"`
	Valor_Encomenda_Shopper              string `json:"Valor_Encomenda_Shopper,omitempty"`
	Numero_Nota_Fiscal_Encomenda_Shopper string `json:"Numero_Nota_Fiscal_Encom,omitempty"`
	Codigo_País_Shopper                  string `json:"Codigo_País_Shopper,omitempty"`
	Cidade_Shopper                       string `json:"Cidade_Shopper,omitempty"`
	CEP_Shopper                          string `json:"CEP_Shopper,omitempty"`
	Bairro_Shopper                       string `json:"Bairro_Shopper,omitempty"`
	Endereco_Shopper                     string `json:"Endereco_Shopper,omitempty"`
	Numero_Shopper                       string `json:"Numero_Shopper,omitempty"`
	Complemento_Shopper                  string `json:"Complemento_Shopper,omitempty"`
	Codigo_Rastreamento_Encomenda        string `json:"Codigo_Rastreamento_Encomenda,omitempty"`
	Codigo_Barras_Conteudo_Encomenda     string `json:"Codigo_Barras_Conteudo_Encomenda,omitempty"`
	Descricao_Conteudo_Encomenda         string `json:"Descricao_Conteudo_Encomenda,omitempty"`
	Encomenda_Assegurada                 string `json:"Encomenda_Assegurada,omitempty"`
	Largura_Encomenda                    string `json:"Largura_Encomenda,omitempty"`
	Altura_Encomenda                     string `json:"Altura_Encomenda,omitempty"`
	Profundidade_Encomenda               string `json:"Profundidade_Encomenda,omitempty"`
}

type PostReservationParams struct {
	Codigo_de_MSG                                           string           `json:"Codigo_de_MSG,omitempty"`
	ID_de_Referencia                                        string           `json:"ID_de_Referencia,omitempty"`
	ID_do_Solicitante                                       string           `json:"ID_do_Solicitante,omitempty"`
	ID_Rede_Lockers                                         int              `json:"ID_Rede_Lockers,omitempty"`
	Data_Hora_Solicitacao                                   string           `json:"Data_Hora_Solicitacao,omitempty"`
	ID_da_Estacao_do_Locker                                 string           `json:"ID_da_Estacao_do_Locker,omitempty"`
	Tipo_de_Servico_Reserva                                 int              `json:"Tipo_de_Servico_Reserva,omitempty"`
	ID_Transacao_Unica                                      string           `json:"ID_Transacao_Unica,omitempty"`
	ID_PSL_Designado                                        int              `json:"ID_PSL_Designado,omitempty"`
	Autenticacao_Login_Operador_Logistico                   int              `json:"Autenticacao_Login,omitempty"`
	Categoria_Porta                                         string           `json:"Categoria_Porta,omitempty"`
	Geracao_de_QRCODE_na_Resposta_MS06                      int              `json:"Geracao_de_QRCODE_na_Resposta_MS06,omitempty"`
	Geracao_de_Codigo_de_Abertura_de_Porta_na_Resposta_MS06 int              `json:"Geracao_de_Codigo_de_Abertura_de_Porta_na_Resposta_MS06,omitempty"`
	Info_Encomendas                                         []infoEncomendas `json:"Info_Encomendas,omitempty"`
	URL_CALL_BACK                                           string           `json:"URL_CALL_BACK,omitempty"`
	Versao_Mensageria                                       string           `json:"Versao_Mensageira,omitempty"`
}

type infoEncomendasResponse struct {
	ID_Encomenda                   string `json:"ID_Encomenda,omitempty"`
	Etiqueta_Encomenda_Rede1minuto string `json:"Etiqueta_Encomenda_Rede1,omitempty"`
	Geracao_QRCODE                 string `json:"Geracao_QRC,omitempty"`
	Geracao_Codigo_Abertura_Porta  string `json:"Geracao_Codigo_Abert,omitempty"`
}

type PostReservationResponse struct {
	Codigo_de_MSG                  string                   `json:"Codigo_de_MSG,omitempty"`
	ID_de_Referencia               string                   `json:"ID_de_Referencia,omitempty"`
	ID_do_Solicitante              string                   `json:"ID_do_Solicitante,omitempty"`
	ID_Rede_Lockers                int                      `json:"ID_Rede_Lockers,omitempty"`
	ID_da_Estacao_do_Locker        string                   `json:"ID_da_Estacao_do_Locker,omitempty"`
	Codigo_Resposta_MS06           string                   `json:"Codigo_Resposta_MS06,omitempty"`
	Data_Hora_Resposta             string                   `json:"Data_Hora_Resposta,omitempty"`
	Codigo_Pais_Locker             string                   `json:"Codigo_Pais_Locker,omitempty"`
	CEP_Locker                     string                   `json:"CEP_Locker,omitempty"`
	Cidade_Locker                  string                   `json:"Cidade_Locker,omitempty"`
	Bairro_Locker                  string                   `json:"Bairro_Locker,omitempty"`
	Endereco_Locker                string                   `json:"Endereco_Locker,omitempty"`
	Numero_Locker                  string                   `json:"Numero_Locker,omitempty"`
	Complemento_Locker             string                   `json:"Complemento,omitempty"`
	ID_da_Porta_do_Locker          string                   `json:"ID_da_Porta_do_Locker,omitempty"`
	ID_do_Operador_da_Porta_Locker string                   `json:"ID_do_Operador_da_Porta_Locker,omitempty"`
	ID_Transacao_Unica             string                   `json:"ID_Transacao_Unica,omitempty"`
	Tipo_de_Servico_Reserva        string                   `json:"Tipo_de_Servico_Reserva,omitempty"`
	DataHora_Inicio_Reserva        string                   `json:"DataHora_Inicio_Reserva,omitempty"`
	DataHora_Final_Reserva         string                   `json:"DataHora_Final_Reserva,omitempty"`
	Info_Encomendas                []infoEncomendasResponse `json:"info_encomendas,omitempty"`
	Versao_Mensageria              string                   `json:"Versao_Mensageria,omitempty"`
}
