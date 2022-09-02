package main

type commandID int

const (

	CMD_NAME commandID = iota
	CMD_SUBSCRIBE
	CMD_CHANNELS
	CMD_SENDFILE
	CMD_QUIT

)

type command struct {
	id commandID
	client *client
	args   []string
	
}