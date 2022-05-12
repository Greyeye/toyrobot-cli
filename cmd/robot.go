package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	TableXMaxLimit = 4
	TableYMaxLimit = 4
	RotationAngle  = 90
)

//RobotDirection angle representation of the robot's direction.
type RobotDirection int

func (e RobotDirection) String() string {
	switch e {
	case NORTH:
		return "NORTH"
	case SOUTH:
		return "SOUTH"
	case WEST:
		return "WEST"
	case EAST:
		return "EAST"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

// enum for the direction of the robot.
const (
	NORTH RobotDirection = 0
	SOUTH RobotDirection = 180
	EAST  RobotDirection = 90
	WEST  RobotDirection = 270
)

type robot struct {
	placed bool           // to check if the robot has been placed successfully or not.
	maxX   int            // maximum value of the robot's X movement.
	x      int            // robot's current X coordinates
	maxY   int            // maximum value of the robot's Y movement.
	y      int            // robot's current Y coordinates
	dir    RobotDirection // robot's current direction it is facing.
}

type IRobot interface {
	Controller(ih InputHandler)
	CheckRobotDirection(dir string) (RobotDirection, error)
	CheckPlacementCoordinates(x, y int) error
	Place(inputX, inputY, dir string)
	Report() string
	Move() error
	Left()
	Right()
}

// NewRobot is a constructor for robot, and return IRobot interface.
func NewRobot(placed bool, maxX int, x int, maxY int, y int, dir RobotDirection) IRobot {
	return &robot{placed: placed, maxX: maxX, x: x, maxY: maxY, y: y, dir: dir}
}

// Controller handles the input to the robot.
// accepts io.Reader as the argument to handle user input.
func (r *robot) Controller(ih InputHandler) {
	inputCommand := ih.userInput()
	splitCommand := strings.Split(inputCommand, " ")
	switch {
	case strings.ToLower(splitCommand[0]) == "move" && r.placed:
		err3 := r.Move()
		if err3 != nil {
			fmt.Println(err3)
		}
	case r.placed && strings.ToLower(splitCommand[0]) == "report":
		fmt.Println(r.Report())
	case r.placed && strings.ToLower(splitCommand[0]) == "right":
		r.Right()
	case r.placed && strings.ToLower(splitCommand[0]) == "left":
		r.Left()
	case strings.ToLower(splitCommand[0]) == "place" && len(splitCommand) > 0:
		placeCommandInput := strings.Split(splitCommand[1], ",")
		if len(placeCommandInput) != 3 {
			fmt.Println("invalid placement command parameters, pleases enter x,y,direction")
			break
		}
		r.Place(placeCommandInput[0], placeCommandInput[1], placeCommandInput[2])
	case !r.placed:
		fmt.Println("please place robot using place command first.")
	default:
		fmt.Println("invalid command", inputCommand)
	}
}

// CheckRobotDirection validates the string input and return the matching enumerator for RobotDirection value
func (r *robot) CheckRobotDirection(dir string) (RobotDirection, error) {
	inputDirection := strings.ToUpper(dir)
	switch inputDirection {
	case "NORTH":
		return NORTH, nil
	case "SOUTH":
		return SOUTH, nil
	case "EAST":
		return EAST, nil
	case "WEST":
		return WEST, nil
	default:
		return 0, errors.New("invalid input of the direction, " + inputDirection + ", please enter from one of the following directions\nnorth, south, east, west.")
	}

}

// CheckPlacementCoordinates validates the coordinates parameters are within the
func (r *robot) CheckPlacementCoordinates(x, y int) error {
	switch {
	case x > TableXMaxLimit:
		return errors.New("requested X coordinate is beyond the limit of " + strconv.Itoa(r.maxX))
	case x < 0:
		return errors.New("requested X coordinate is under the limit of the table, enter larger than 0")
	case y > TableYMaxLimit:
		return errors.New("requested Y coordinate is beyond the limit of " + strconv.Itoa(r.maxY))
	case y < 0:
		return errors.New("requested Y coordinate is under the limit of the table, enter larger than 0")
	}
	return nil
}

// Place moves the robot to the specified position on the table.
func (r *robot) Place(inputX, inputY, dir string) {
	x, err1 := strconv.Atoi(inputX)
	if err1 != nil {
		fmt.Printf("entered X location \"%v\" is not a number, please try again\n", inputX)
		return
	}
	y, err2 := strconv.Atoi(inputY)
	if err2 != nil {
		fmt.Printf("entered Y location \"%v\" is not a number, please try again\n", inputY)
		return
	}
	err := r.CheckPlacementCoordinates(x, y)
	if err != nil {
		fmt.Println(err)
		return
	}
	validDir, err2 := r.CheckRobotDirection(dir)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	r.x = x
	r.y = y
	r.dir = validDir
	r.placed = true
}

// Report return current coordinates and direction the robot is facing.
func (r *robot) Report() string {
	return fmt.Sprintf("%d, %d, %s", r.x, r.y, r.dir.String())
}

// Move will attempt to move the robot to the direction it's currently facing
func (r *robot) Move() error {
	switch r.dir {
	case NORTH:
		err := r.CheckPlacementCoordinates(r.x, r.y+1)
		if err != nil {
			return errors.New("unable to move robot to coordinates(" + strconv.Itoa(r.x) + "," + strconv.Itoa(r.y+1) + ")")
		}
		r.y += 1
	case EAST:
		err := r.CheckPlacementCoordinates(r.x+1, r.y)
		if err != nil {
			return errors.New("unable to move robot to coordinates(" + strconv.Itoa(r.x+1) + "," + strconv.Itoa(r.y) + ")")
		}
		r.x += 1
	case SOUTH:
		err := r.CheckPlacementCoordinates(r.x, r.y-1)
		if err != nil {
			return errors.New("unable to move robot to coordinates(" + strconv.Itoa(r.x) + "," + strconv.Itoa(r.y-1) + ")")
		}
		r.y -= 1
	case WEST:
		err := r.CheckPlacementCoordinates(r.x-1, r.y)
		if err != nil {
			return errors.New("unable to move robot to coordinates(" + strconv.Itoa(r.x-1) + "," + strconv.Itoa(r.y) + ")")
		}
		r.x -= 1
	}
	return nil
}

// Left will change the direction of the robot by the negative RotationAngle
func (r *robot) Left() {
	switch {
	case r.dir-RotationAngle < 0:
		r.dir = 270
	default:
		r.dir -= RotationAngle
	}
}

// Right will change the direction of the robot by the positive RotationAngle
func (r *robot) Right() {
	switch {
	case r.dir+RotationAngle >= 360:
		r.dir = 0
	default:
		r.dir += RotationAngle
	}
}
