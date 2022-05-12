package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestNewRobot(t *testing.T) {
	tests := []struct {
		name  string
		robot IRobot
		want  IRobot
	}{
		{
			name:  "new robot",
			robot: NewRobot(false, TableXMaxLimit, 0, TableYMaxLimit, 0, NORTH),
			want:  &robot{placed: false, maxX: TableXMaxLimit, x: 0, maxY: TableYMaxLimit, y: 0, dir: NORTH},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.robot, tt.want) {
				t.Errorf("NewRobot() = %v, want %v", tt.robot, tt.want)
			}
		})
	}
}

func TestInputHandler_userInput(t *testing.T) {

	tests := []struct {
		name      string
		mockInput io.Reader
		want      string
	}{
		{
			name:      "test input",
			mockInput: strings.NewReader("input one two three"),
			want:      "input one two three",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InputHandler{
				Reader: tt.mockInput,
			}
			if got := s.userInput(); got != tt.want {
				t.Errorf("userInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
