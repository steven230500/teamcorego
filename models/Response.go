package models

type ResponseData struct {
	Titulo string `json:"titulo"`
	Dia    string `json:"dia"`
	Info   []InfoBody `json:"info"`
	APIVersion int `json:"api_version"`
}

type InfoBody struct {
	PreguntaID string `json:"pregunta_id"`
	Pregunta   string `json:"pregunta"`
}