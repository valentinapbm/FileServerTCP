package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

//Cliente será responsable de tener la info name, current tcp conexion, channel

type client struct {
	conn     net.Conn
	name     string
	channel  *channel
	commands chan<- command //se envia n los comandos y luego se envian al server

}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil{
			return
		}
		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd{

		case "/name":
			c.commands <- command{
				id: CMD_NAME,
				client: c,
				args: args,
			}
		case "/subscribe":
			c.commands <- command{
				id: CMD_SUBSCRIBE,
				client: c,
				args: args,
			}
		case "/channels":
			c.commands <- command{
				id: CMD_CHANNELS,
				client: c,
				args: args,
			}
		case "/sendfile":
			c.commands <- command{
				id: CMD_SENDFILE,
				client: c,
				args: args,
			}
		case "/quit":
			c.commands <- command{
				id: CMD_QUIT,
				client: c,
				args: args,
			}
		default:
			c.err(fmt.Errorf("unknow command: %s", cmd))

		}
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("ERR: " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}
