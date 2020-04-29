package manufacturer

import "github.com/jackc/pgtype"

// Representation of the "manufacturers" table object as a structure
type Manufacturer struct {
	Id        uint64             `json:"id"`
	Name      string             `json:"name"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}
