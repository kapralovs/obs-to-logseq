package note

import (
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		filename string
		content  []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Empty bytes slice",
			args: args{
				filename: "../../obs_note.md",
				content:  []byte{},
			},
			wantErr: false,
		},
		{
			name: "Nil content",
			args: args{
				filename: "../../obs_note.md",
				content:  nil,
			},
			wantErr: false,
		},
		{
			name: "Without last dot",
			args: args{
				filename: "../../obs_note.md",
				content:  []byte("Это заметка про язык программирования [[Golang|Go]]"),
			},
			wantErr: false,
		},
		{
			name: "Valid",
			args: args{
				filename: "../../obs_note.md",
				content:  []byte("Это заметка про язык программирования [[Golang|Go]]."),
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create(tt.args.filename, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOpen(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Valid",
			args: args{
				filename: "../../obs_note.md",
			},
			want:    []byte("Это заметка про язык программирования [[Golang|Go]]."),
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Open(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Open() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

// func TestCreate(t *testing.T) {
// 	type args struct {
// 		filename string
// 		content  []byte
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := Create(tt.args.filename, tt.args.content); (err != nil) != tt.wantErr {
// 				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
