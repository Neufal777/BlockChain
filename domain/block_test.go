package domain

import "testing"

func TestBlock_SetHash(t *testing.T) {
	type fields struct {
		Timestamp     int64
		Data          []byte
		PrevBlockHash []byte
		Hash          []byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Block{
				Timestamp:     tt.fields.Timestamp,
				Data:          tt.fields.Data,
				PrevBlockHash: tt.fields.PrevBlockHash,
				Hash:          tt.fields.Hash,
			}
			b.SetHash()
		})
	}
}
