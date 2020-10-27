package main

import (
	"reflect"
	"testing"
)

func TestBoard_GetOccupiedFields(t *testing.T) {
	tests := []struct {
		board Board
		name  string
		args  int
		want  []int
	}{
		{
			board: *NewBoard(8),
			name:  "Test",
			args:  0,
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 16, 18, 24, 27, 32, 36, 40, 45, 48, 54, 56, 63},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.board.GetOccupiedFields(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.GetOccupiedFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_getLineNumbers(t *testing.T) {
	tests := []struct {
		board Board
		name  string
		args  int
		want  []int
	}{
		{
			board: *NewBoard(8),
			name:  "Test Line 0",
			args:  0,
			want:  []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			board: *NewBoard(8),
			name:  "Test Line 1",
			args:  1,
			want:  []int{9, 10, 11, 12, 13, 14, 15},
		},
		{
			board: *NewBoard(8),
			name:  "Test Line 7",
			args:  7,
			want:  []int{57, 58, 59, 60, 61, 62, 63},
		},
		{
			board: *NewBoard(15),
			name:  "Test Line 0",
			args:  0,
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.board.getLineNumbers(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.getLineNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_getColumnNumbers(t *testing.T) {
	tests := []struct {
		board Board
		name  string
		args  int
		want  []int
	}{
		{
			board: *NewBoard(8),
			name:  "Test Column 0",
			args:  0,
			want:  []int{8, 16, 24, 32, 40, 48, 56},
		},
		{
			board: *NewBoard(8),
			name:  "Test Column 1",
			args:  1,
			want:  []int{9, 17, 25, 33, 41, 49, 57},
		},
		{
			board: *NewBoard(8),
			name:  "Test Column 7",
			args:  7,
			want:  []int{15, 23, 31, 39, 47, 55, 63},
		},
		{
			board: *NewBoard(15),
			name:  "Test Column 0",
			args:  0,
			want:  []int{15, 30, 45, 60, 75, 90, 105, 120, 135, 150, 165, 180, 195, 210},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.board.getColumnNumbers(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.getColumnNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_getDiagonals(t *testing.T) {
	type args struct {
		line, col int
	}
	tests := []struct {
		board Board
		name  string
		args  args
		want  []int
	}{
		{
			board: *NewBoard(8),
			name:  "Test Diagonals Top-Left",
			args:  args{0, 0},
			want:  []int{9, 18, 27, 36, 45, 54, 63},
		},
		{
			board: *NewBoard(8),
			name:  "Test Diagonals Top-Right",
			args:  args{0, 7},
			want:  []int{14, 21, 28, 35, 42, 49, 56},
		},
		{
			board: *NewBoard(5),
			name:  "Test Center",
			args:  args{2, 2},
			want:  []int{0, 4, 6, 8, 16, 18, 20, 24},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.board.getDiagonals(tt.args.line, tt.args.col); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.getDiagonals() = %v, want %v", got, tt.want)
			}
		})
	}
}
