package entity

type Message struct {
	Word       string `json:"word"`
	Occurrence int    `json:"occurrence"`
}

type MessageParam struct {
	Message string `json:"message" query:"message"`
}
