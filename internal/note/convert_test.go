package note

import (
	"reflect"
	"testing"
)

func TestConvertLinks(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Valid content",
			args: args{
				content: []byte("Это заметка про язык программирования [[Golang|Go]]."),
			},
			want:    []byte("Это заметка про язык программирования [[Golang]]."),
			wantErr: false,
		},
		{
			name: "Valid content without last dot",
			args: args{
				content: []byte("Это заметка про язык программирования [[Golang|Go]]"),
			},
			want:    []byte("Это заметка про язык программирования [[Golang]]"),
			wantErr: false,
		},
		{
			name: "Content without links",
			args: args{
				content: []byte("12345"),
			},
			want:    []byte("12345"),
			wantErr: false,
		},
		/*
			{
				name: "Invalid",
				args: args{
					content: []byte{},
				},
				want:    nil,
				wantErr: true,
			},
		*/
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertLinks(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertNote() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
