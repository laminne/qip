package types

// UserOutBoxResponseJSONLD /users/:id/outboxのレスポンス
type UserOutBoxResponseJSONLD struct {
	Context    []interface{} `json:"@context"`
	ID         string        `json:"id"`         // /users/:id/outbox?page=true
	Type       string        `json:"type"`       // OrderdCollection
	TotalItems int           `json:"totalItems"` // int
	First      string        `json:"first"`      // /users/:id/outbox?page=true&since_id=000000000000000000000000
	Last       string        `json:"last"`
}

type UserOutBoxPageJSONLD struct {
	Context      []interface{} `json:"@context"`
	ID           string        `json:"id"`
	PartOf       string        `json:"partOf"`
	Type         string        `json:"type"`
	TotalItems   int           `json:"totalItems"`
	OrderedItems []NoteJSONLD  `json:"orderedItems"`
	Prev         string        `json:"prev"`
	Next         string        `json:"next"`
}
