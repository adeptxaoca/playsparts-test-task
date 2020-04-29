package part

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPart_ToPb(t *testing.T) {
	p := Part{Id: 1, ManufacturerId: 1, Name: "part 1", VendorCode: "VC1"}
	_ = p.CreatedAt.Set(time.Now())
	_ = p.UpdatedAt.Set(time.Now())
	_ = p.DeletedAt.Set(time.Now())
	pbPart := p.ToPb()
	assert.Equal(t, p.Id, pbPart.Id)
	assert.Equal(t, p.ManufacturerId, pbPart.ManufacturerId)
	assert.Equal(t, p.Name, pbPart.Name)
	assert.Equal(t, p.VendorCode, pbPart.VendorCode)
	assert.Equal(t, p.CreatedAt.Time.Unix(), pbPart.CreatedAt)
	assert.Equal(t, p.UpdatedAt.Time.Unix(), pbPart.UpdatedAt)
	assert.Equal(t, p.DeletedAt.Time.Unix(), pbPart.DeletedAt)
}
