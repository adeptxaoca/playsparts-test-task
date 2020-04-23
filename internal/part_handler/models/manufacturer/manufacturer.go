package manufacturer

import "time"

type Manufacturer struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
