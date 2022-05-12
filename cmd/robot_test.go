package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestRobotDirection_string(t *testing.T) {
	tests := []struct {
		name string
		e    RobotDirection
		want string
	}{
		{
			name: "north",
			e:    NORTH,
			want: "NORTH",
		},
		{
			name: "west",
			e:    WEST,
			want: "WEST",
		},
		{
			name: "east",
			e:    EAST,
			want: "EAST",
		},
		{
			name: "south",
			e:    SOUTH,
			want: "SOUTH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_robotLocation_CheckPlacementCoordinates(t *testing.T) {

	type args struct {
		inputX int
		inputY int
	}
	tests := []struct {
		name    string
		fields  robot
		args    args
		wantErr bool
	}{
		{
			name:    "normal ok test, x,y both 0",
			fields:  robot{placed: false, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{0, 0},
			wantErr: false,
		},
		{
			name:    "normal ok value on the maxX and maxY",
			fields:  robot{placed: false, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{TableXMaxLimit, TableXMaxLimit},
			wantErr: false,
		},
		{
			name:    "below the limit for x",
			fields:  robot{placed: false, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{-1, 4},
			wantErr: true,
		},
		{
			name:    "below the limit for y",
			fields:  robot{placed: false, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{0, -1},
			wantErr: true,
		},
		{
			name:    "over the limit for x",
			fields:  robot{placed: false, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{TableXMaxLimit + 1, -1},
			wantErr: true,
		},
		{
			name:    "over the limit for y",
			fields:  robot{placed: false, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{1, TableYMaxLimit + 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &robot{
				maxX: tt.fields.maxX,
				x:    tt.fields.x,
				maxY: tt.fields.maxY,
				y:    tt.fields.y,
				dir:  tt.fields.dir,
			}
			err := r.CheckPlacementCoordinates(tt.args.inputX, tt.args.inputY)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPlacementCoordinates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_robotLocation_CheckRobotDirection(t *testing.T) {

	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		fields  robot
		args    args
		want    RobotDirection
		wantErr bool
	}{
		{
			name:    "normal ok test lower case",
			fields:  robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{dir: "north"},
			want:    NORTH,
			wantErr: false,
		},
		{
			name:    "normal ok test upper case",
			fields:  robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{dir: "NORTH"},
			want:    NORTH,
			wantErr: false,
		},
		{
			name:    "normal ok test mixed case North",
			fields:  robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{dir: "North"},
			want:    NORTH,
			wantErr: false,
		},
		{
			name:    "normal ok test mixed case norTH",
			fields:  robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{dir: "norTH"},
			want:    NORTH,
			wantErr: false,
		},
		{
			name:    "normal ok test south",
			fields:  robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{dir: "south"},
			want:    SOUTH,
			wantErr: false,
		},
		{
			name:    "normal ok test east",
			fields:  robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{dir: "east"},
			want:    EAST,
			wantErr: false,
		},
		{
			name:    "normal ok test west",
			fields:  robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{dir: "west"},
			want:    WEST,
			wantErr: false,
		},
		{
			name:    "bad direction 1",
			fields:  robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{dir: "0"},
			want:    NORTH,
			wantErr: true,
		},
		{
			name:    "bad direction 2",
			fields:  robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{dir: "NORTHSOUTH"},
			want:    NORTH,
			wantErr: true,
		},
		{
			name:    "bad direction 3, with white space prefix",
			fields:  robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{dir: " NORTH"},
			want:    NORTH,
			wantErr: true,
		},
		{
			name:    "bad direction empty",
			fields:  robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:    args{dir: ""},
			want:    NORTH,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &robot{
				maxX: tt.fields.maxX,
				x:    tt.fields.x,
				maxY: tt.fields.maxY,
				y:    tt.fields.y,
				dir:  tt.fields.dir,
			}
			got, err := r.CheckRobotDirection(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckRobotDirection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckRobotDirection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_robotLocation_Left(t *testing.T) {
	tests := []struct {
		name             string
		initialDirection RobotDirection
		want             IRobot
	}{
		{name: "north to left",
			initialDirection: NORTH,
			want:             &robot{true, TableXMaxLimit, 0, TableYMaxLimit, 0, WEST},
		},
		{name: "west to left",
			initialDirection: WEST,
			want:             &robot{true, TableXMaxLimit, 0, TableYMaxLimit, 0, SOUTH},
		},
		{name: "south to left",
			initialDirection: SOUTH,
			want:             &robot{true, TableXMaxLimit, 0, TableYMaxLimit, 0, EAST},
		},
		{name: "east to left",
			initialDirection: EAST,
			want:             &robot{true, TableXMaxLimit, 0, TableYMaxLimit, 0, NORTH},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRobot(true, TableXMaxLimit, 0, TableYMaxLimit, 0, tt.initialDirection)
			r.Left()
			if !reflect.DeepEqual(r, tt.want) {
				t.Errorf("Left() got = %v, want %v", r, tt.want)
			}

		})
	}
}

func Test_robotLocation_Move(t *testing.T) {

	tests := []struct {
		name    string
		robot   robot
		wantX   int
		wantY   int
		wantErr bool
	}{
		{
			name:    "normal move north",
			robot:   robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			wantX:   0,
			wantY:   1,
			wantErr: false,
		},
		{
			name:    "on the maxX move north",
			robot:   robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: TableXMaxLimit, y: 0, dir: NORTH},
			wantX:   TableXMaxLimit,
			wantY:   1,
			wantErr: false,
		},
		{
			name:    "on the maxY move north",
			robot:   robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: TableXMaxLimit, y: TableYMaxLimit, dir: NORTH},
			wantX:   TableXMaxLimit,
			wantY:   TableYMaxLimit,
			wantErr: true,
		},
		{
			name:    "on the 0,0 move south",
			robot:   robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: SOUTH},
			wantX:   0,
			wantY:   0,
			wantErr: true,
		},
		{
			name:    "on the 0,4 move south",
			robot:   robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: TableYMaxLimit, dir: SOUTH},
			wantX:   0,
			wantY:   TableYMaxLimit - 1,
			wantErr: false,
		},
		{
			name:    "on the 4,0 move west",
			robot:   robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: TableXMaxLimit, y: 1, dir: WEST},
			wantX:   TableXMaxLimit - 1,
			wantY:   1,
			wantErr: false,
		},
		{
			name:    "on the maxX, maxY, move east",
			robot:   robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: TableXMaxLimit, y: TableYMaxLimit, dir: EAST},
			wantX:   TableXMaxLimit,
			wantY:   TableYMaxLimit,
			wantErr: true,
		},
		{
			name:    "on the 0, 0, move west",
			robot:   robot{maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: WEST},
			wantX:   0,
			wantY:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.robot
			if err := r.Move(); (err != nil) != tt.wantErr {
				t.Errorf("Move() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(r.x, tt.wantX) {
				t.Errorf("Move() got = %v, want %v", r.x, tt.wantX)
			}
			if !reflect.DeepEqual(r.y, tt.wantY) {
				t.Errorf("Move() got1 = %v, want %v", r.y, tt.wantY)
			}
		})
	}
}

func Test_robotLocation_Place(t *testing.T) {

	type args struct {
		x   string
		y   string
		dir string
	}
	tests := []struct {
		name         string
		wantLocation *robot
		args         args
	}{
		{
			name:         "normal ok condition",
			wantLocation: &robot{placed: true, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 1, y: 1, dir: NORTH},
			args:         args{x: "1", y: "1", dir: "north"},
		},
		{
			name:         "out of bounds placement x",
			wantLocation: &robot{placed: false, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:         args{x: strconv.Itoa(TableXMaxLimit + 1), y: "1", dir: "north"},
		},
		{
			name:         "out of bounds placement y",
			wantLocation: &robot{placed: false, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:         args{x: "1", y: strconv.Itoa(TableYMaxLimit + 1), dir: "north"},
		},
		{
			name:         "invalid character x",
			wantLocation: &robot{placed: false, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:         args{x: "a", y: "1", dir: "north"},
		},
		{
			name:         "invalid character y",
			wantLocation: &robot{placed: false, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:         args{x: "1", y: "y", dir: "north"},
		},
		{
			name:         "invalid direction placement",
			wantLocation: &robot{placed: false, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:         args{x: strconv.Itoa(TableXMaxLimit), y: strconv.Itoa(TableYMaxLimit), dir: "invalid"},
		},
		{
			name:         "invalid direction placement invalid x and y",
			wantLocation: &robot{placed: false, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:         args{x: strconv.Itoa(TableXMaxLimit + 1), y: strconv.Itoa(TableYMaxLimit + 1), dir: "invalid"},
		},
		{
			name:         "invalid direction placement invalid x and y",
			wantLocation: &robot{placed: false, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args:         args{x: strconv.Itoa(0 - 1), y: strconv.Itoa(0 - 1), dir: "invalid"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRobot(false, TableXMaxLimit, 0, TableYMaxLimit, 0, NORTH)
			r.Place(tt.args.x, tt.args.y, tt.args.dir)
			if !reflect.DeepEqual(r, tt.wantLocation) {
				t.Errorf("place() got = %v, want %v", r, tt.wantLocation)
			}
		})
	}
}

func Test_robotLocation_Report(t *testing.T) {

	tests := []struct {
		name  string
		robot robot
		want  string
	}{
		{
			name:  "normal",
			robot: robot{placed: true, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			want:  "0, 0, NORTH",
		},
		{
			name:  "coordinates on the border line",
			robot: robot{placed: false, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: TableXMaxLimit, y: TableYMaxLimit, dir: NORTH},
			want:  fmt.Sprintf("%v, %v, %v", TableXMaxLimit, TableYMaxLimit, NORTH.String()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.robot
			if !reflect.DeepEqual(r.Report(), tt.want) {
				t.Errorf("Report() got = %v, want %v", r, tt.want)
			}
		})
	}
}

func Test_robotLocation_Right(t *testing.T) {
	tests := []struct {
		name             string
		initialDirection RobotDirection
		want             IRobot
	}{
		{name: "north to right",
			initialDirection: NORTH,
			want:             &robot{true, TableXMaxLimit, 0, TableYMaxLimit, 0, EAST},
		},
		{name: "west to right",
			initialDirection: WEST,
			want:             &robot{true, TableXMaxLimit, 0, TableYMaxLimit, 0, NORTH},
		},
		{name: "south to right",
			initialDirection: SOUTH,
			want:             &robot{true, TableXMaxLimit, 0, TableYMaxLimit, 0, WEST},
		},
		{name: "east to right",
			initialDirection: EAST,
			want:             &robot{true, TableXMaxLimit, 0, TableYMaxLimit, 0, SOUTH},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRobot(true, TableXMaxLimit, 0, TableYMaxLimit, 0, tt.initialDirection)
			r.Right()
			if !reflect.DeepEqual(r, tt.want) {
				t.Errorf("Left() got = %v, want %v", r, tt.want)
			}

		})
	}
}

//Test_robotLocation_Controller tests multiple input chains.
//test against the final coordinates and direction of the robot.
func Test_robotLocation_Controller(t *testing.T) {
	inputTests := []struct {
		name              string
		wantFinalLocation *robot
		args              []InputHandler
	}{
		{
			name:              "place 1,1,NORTH and report",
			wantFinalLocation: &robot{placed: true, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 1, y: 1, dir: NORTH},
			args: []InputHandler{
				{strings.NewReader("place 1,1,NORTH")},
				{strings.NewReader("report")},
			},
		},
		{
			name:              "place 0,0,NORTH, move, move, move",
			wantFinalLocation: &robot{placed: true, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 3, dir: NORTH},
			args: []InputHandler{
				{strings.NewReader("place 0,0,NORTH")},
				{strings.NewReader("move")},
				{strings.NewReader("move")},
				{strings.NewReader("move")},
			},
		},
		{
			name:              "place 0,4,NORTH, move, move, move",
			wantFinalLocation: &robot{placed: true, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 4, dir: NORTH},
			args: []InputHandler{
				{strings.NewReader("place 0,4,NORTH")},
				{strings.NewReader("move")},
				{strings.NewReader("move")},
				{strings.NewReader("move")},
			},
		},
		{
			name:              "place 0,4,NORTH, right,move, move, move",
			wantFinalLocation: &robot{placed: true, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 3, y: 4, dir: EAST},
			args: []InputHandler{
				{strings.NewReader("place 0,4,NORTH")},
				{strings.NewReader("right")},
				{strings.NewReader("move")},
				{strings.NewReader("move")},
				{strings.NewReader("move")},
			},
		},
		{
			name:              "place 0,0,NORTH, place 4,4,SOUTH",
			wantFinalLocation: &robot{placed: true, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 4, y: 4, dir: SOUTH},
			args: []InputHandler{
				{strings.NewReader("place 0,4,NORTH")},
				{strings.NewReader("place 4,4,SOUTH")},
			},
		},
		{
			name:              "move move without place",
			wantFinalLocation: &robot{placed: false, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args: []InputHandler{
				{strings.NewReader("move")},
				{strings.NewReader("move")},
			},
		},
		{
			name:              "start from maxX, maxY, move move left move move left move move",
			wantFinalLocation: &robot{placed: true, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: TableXMaxLimit, y: TableYMaxLimit, dir: NORTH},
			args: []InputHandler{
				{strings.NewReader("place 4,4,WEST")},
				{strings.NewReader("move")},
				{strings.NewReader("move")},
				{strings.NewReader("left")},
				{strings.NewReader("move")},
				{strings.NewReader("move")},
				{strings.NewReader("left")},
				{strings.NewReader("move")},
				{strings.NewReader("move")},
				{strings.NewReader("left")},
				{strings.NewReader("move")},
				{strings.NewReader("move")},
			},
		},
		{
			name:              "bad place command",
			wantFinalLocation: &robot{placed: false, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 0, y: 0, dir: NORTH},
			args: []InputHandler{
				{strings.NewReader("place 1,2,NORTH,4")},
			},
		},
		{
			name:              "bad command in between",
			wantFinalLocation: &robot{placed: true, maxX: TableXMaxLimit, maxY: TableYMaxLimit, x: 1, y: 1, dir: SOUTH},
			args: []InputHandler{
				{strings.NewReader("place 1,2,SOUTH")},
				{strings.NewReader("invalidMove")},
				{strings.NewReader("move")},
			},
		},
	}
	for _, tt := range inputTests {
		t.Run(tt.name, func(t *testing.T) {
			testRobot := NewRobot(false, TableXMaxLimit, 0, TableYMaxLimit, 0, NORTH)
			for _, input := range tt.args {
				testRobot.Controller(input)
			}
			if !reflect.DeepEqual(testRobot, tt.wantFinalLocation) {
				t.Errorf("Contoller() got testRobot = %v , wantRobot = %v ", testRobot, tt.wantFinalLocation)
			}
		})
	}

}
