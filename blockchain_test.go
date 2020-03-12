package main

import "testing"

func TestBlockchain_AddBlock(t *testing.T) {
	type fields struct {
		blocks []*Block
	}
	type args struct {
		data string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &Blockchain{
				blocks: tt.fields.blocks,
			}
			bc.AddBlock(tt.args.data)
		})
	}
}
