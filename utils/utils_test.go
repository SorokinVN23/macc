package utils

import (
	"reflect"
	"testing"
)

func TestConvertToFloat64(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "From int64",
			args:    args{int64(10)},
			want:    float64(10),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToFloat64(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToInt64(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "From float64",
			args:    args{float64(10)},
			want:    int64(10),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToInt64(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		v1 interface{}
		v2 interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "TestSum from float64",
			args:    args{v1: float64(10), v2: float64(20)},
			want:    float64(30),
			wantErr: false,
		},
		{
			name:    "TestSum from int64",
			args:    args{v1: int64(10), v2: int64(20)},
			want:    int64(30),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sum(tt.args.v1, tt.args.v2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
