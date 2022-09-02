package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	channels map[string]*channel
	commands chan command
}

func newServer() *server {
	return &server{
		channels: make(map[string]*channel),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NAME:
			s.name(cmd.client, cmd.args)
		case CMD_SUBSCRIBE:
			s.subscribe(cmd.client, cmd.args)
		case CMD_CHANNELS:
			s.listChannels(cmd.client, cmd.args)
		case CMD_SENDFILE:
			s.sendFile(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args)
		}
	}
}

func (s *server) newClient(conn net.Conn) *client{
	log.Printf("new client has connected: %s", conn.RemoteAddr().String())

	return &client{
		conn:     conn,
		name:     "anonymous",
		commands: s.commands,
	}
	
	//c.readInput() 

}

func (s *server) name(c *client, args []string) {
	c.name = args[1]
	c.msg(fmt.Sprintf("Hello %s ", c.name))

}
func (s *server) subscribe(c *client, args []string) {

	channelName := args[1]
	ch, ok := s.channels[channelName]
	if !ok {
		ch = &channel{
			name:    channelName,
			members: make(map[net.Addr]*client),
		}
		s.channels[channelName] = ch
	}
	ch.members[c.conn.RemoteAddr()] = c

	s.quitCurrentChannel(c)

	c.channel = ch

	ch.broadcast(c, fmt.Sprintf("%s has joined the room", c.name))
	c.msg(fmt.Sprintf("Welcome to %s", ch.name))
}
func (s *server) listChannels(c *client, args []string) {

	var channels []string
	for name := range s.channels {
		channels = append(channels, name)
	}
	c.msg(fmt.Sprintf("available rooms are: %s", strings.Join(channels, ", ")))

}
func (s *server) sendFile(c *client, args []string) {
	if c.channel == nil {
		c.err(errors.New("you must join a channel to send a file"))
		return
	}
	fileName := args[1]
	c.channel.sendFileToChannel(c, fileName)
	
	

}
func (s *server) quit(c *client, args []string) {

	log.Printf("Client has disconnected: %s", c.conn.RemoteAddr().String())
	s.quitCurrentChannel(c)
	c.msg("Good bye!")
	c.conn.Close()

}

func (s *server) quitCurrentChannel(c *client) {
	if c.channel != nil {
		delete(c.channel.members, c.conn.RemoteAddr())
		c.channel.broadcast(c, fmt.Sprintf("%s has left the room", c.name))
	}

}
