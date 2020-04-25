package manufacturer

import "github.com/jackc/pgtype"

type Manufacturer struct {
	Id        uint64             `json:"id"`
	Name      string             `json:"name"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}
