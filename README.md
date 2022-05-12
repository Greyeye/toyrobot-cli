# Toy Robot challenge.

## What does it do?  
This is a robot command line simulator.  
These are the requirements of the simulator.  
1. Table has a size limit of 5x5 units. (unit is not defined, please don't ask what if the unit is 1km!)  
2. Robot can move 1 unit per move command. 
3. Robot must not move behind the table limit. (so it does not fall down from the edge)
4. Robot can turn left or right at exactly 90 degrees angle.
5. Robot will not accept any other command but "place" in the start.
6. Robot can only be place within the boundary of the table. (inside 5x5)

## Robot Commands 
1. place *NUMBER*,*NUMBER*,*DIRECTION*  
  **NUMBER** must be between 0 and 5, anything larger or smaller will return invalid command.  
  **DIRECTION** must be one of 4, `north`, `west`, `south` and `east`.  Any other direction will return invalid command.  
  **Example:** *place 1,2,south*
2. move  
  This command will move the robot 1 unit of the direction it is facing.
3. left
  This command will turn the robot **-90** degrees to the left of the direction it is facing.
4. right
   This command will turn the robot **+90** degrees to the right of the direction it is facing.
5. report
  This command will output/print the coordinates of the robot to the console.
  ```console
  >report  
  1, 2, NORTH
  ```

## How to run
if you have `go` v1.18 installed  
```bash 
CGO_ENABLED=0 go run ./cmd
```
Alternatively, you can use docker to build the executable binaries.
```bash 
make docerbuild
```
This will download go docker image and save tar file under `bin` directory.

## How to run unit tests
You must have `go` v1.18 installed
```bash 
go test ./... -v
```
Alternatively, you can use make file.
```bash 
make test
```
Coverage can be checked by running
```bash 
go test -coverprofile=coverage.out ./cmd
go tool cover -html=coverage.out
```

## If you want more challenge..
1. Initialise the table size using input command, or command run flag. 
2. allow robot to face any direction (360 degrees)
3. give final coordinates and let robot reach it by itself.
4. send multiple batches of commands and perform only when "run" command is given.


## Future maintenance note...
1. To update go version, modify it on `go.mod` and `Dockerfile`'s GO_VER line.
2. Table Size and rotational angle are constant within [robot.go](./cmd/robot.go) file. 