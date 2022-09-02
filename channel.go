package main

import (
	"errors"
	"io"
	"net"
	"os"
	"strings"
	"fmt"
)

type channel struct {
	name    string
	members map[net.Addr]*client
}

func (ch *channel) broadcast(sender *client, msg string){

	for addr, m := range ch.members{
		if addr != sender.conn.RemoteAddr(){
			m.msg(msg)
		}
	}
}

func (ch *channel) sendFileToChannel(sender *client, fileName string){
	
	for addr, m := range ch.members{
		if addr != sender.conn.RemoteAddr(){

		//file to read
		file, err := os.Open(strings.TrimSpace(fileName)) // For read access.
		if err != nil {
			sender.err(errors.New("File could not be sent"))
		}
		//
		defer file.Close() // make sure to close the file even if we panic.
		m.msg(sender.name +" send a file : "+fileName+". ")
		
		
		n, error1 := io.Copy(m.conn, file)
		if error1 != nil {
			sender.err(errors.New("There was an error"))
		}
				
		fmt.Println(n, "bytes sent")
		}
		}
}