package main

import (
	"reflect"
	"testing"
)

func Test_convertLinks(t *testing.T) {
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
			want:    []byte("Это заметка про язык программирования [[Golang]]"),
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
			name: "ContentWithoutLinks",
			args: args{
				content: []byte("12345"),
			},
			want:    []byte("12345"),
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertLinks(tt.args.content)
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

func Test_createLogseqNote(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid",
			args: args{
				content: []byte("Это заметка про язык программирования [[Golang|Go]]."),
			},
			wantErr: false,
		},
		{
			name: "Without last dot",
			args: args{
				content: []byte("Это заметка про язык программирования [[Golang|Go]]"),
			},
			wantErr: false,
		},
		{
			name: "Empty bytes slice",
			args: args{
				content: []byte{},
			},
			wantErr: false,
		},
		{
			name:    "Nil arg",
			args:    args{content: nil},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createLogseqNote(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("createLogseqNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createObsNote(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid",
			args: args{
				content: content,
			},
			wantErr: false,
		},
		{
			name: "EmptyContent",
			args: args{
				content: "",
			},
			wantErr: false,
		},
		{
			name: "ContentInDigits",
			args: args{
				content: "12345",
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createObsNote(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("createObsNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
