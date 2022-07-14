package client_model

type ClientModel struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Function  string `json:"function"`
	CreatedBy string `json:"created_by"`
}
