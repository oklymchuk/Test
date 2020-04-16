// Task1 project main.go

package main

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_converToInt(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := converToInt(tt.args.ss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("converToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_specialSort(t *testing.T) {
	type args struct {
		si []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := specialSort(tt.args.si); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("specialSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myCompare(t *testing.T) {
	type args struct {
		first  int
		second int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myCompare(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("myCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
