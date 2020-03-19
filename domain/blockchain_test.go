package domain

import "testing"

func TestBlockchain_AddBlock(t *testing.T) {
	type fields struct {
		Blocks []*Block
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
				Blocks: tt.fields.Blocks,
			}
			bc.AddBlock(tt.args.data)
		})
	}
}
