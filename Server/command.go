package main

const (
	CMD_NAME int = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MESSAGE
	CMD_MEMBERS
	CMD_QUIT
)

type command struct {
	id     int
	client *client
	args   []string
}
