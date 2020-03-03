package main

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
		{"test1", fields{
			Timestamp:     716283,
			Data:          []byte("Prueba1"),
			PrevBlockHash: []byte("60ba6192797abd8"),
			Hash:          []byte("60b860ba6192797abd8"),
		}},
		{"test2", fields{
			Timestamp:     716282983,
			Data:          []byte("Prueba2"),
			PrevBlockHash: []byte("60ba61927975abd8"),
			Hash:          []byte("60b860b6a6192797abd8"),
		}},
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
