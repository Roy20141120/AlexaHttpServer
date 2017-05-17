package myserver

import (
	"net"
	"utils"
	"common"
	"log"
)

func StartSocketServer(host string, timeinterval int){
	netListen, err := net.Listen("tcp", host)
	utils.CheckError(err)
	defer netListen.Close()
	utils.Log("Waiting for clients")

	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		utils.Log(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn, timeinterval)
	}
}


//handle the connection
func handleConnection(conn net.Conn, timeout int ) {

	defer conn.Close()

	tmpBuffer := make([]byte, 0)
	buffer := make([]byte, 1024)
	messnager := make(chan byte)

	connOK := true

	go func(){
		for connOK {
			log.Println("Socket Server waiting for next Alexa event")
			msg := <- common.HttpToSocket
			log.Println("event:", msg)
			log.Println("Socket Server get Alexa event")

			conn.Write(utils.Encode([]byte(msg)))
			log.Println("Socket Server send Alexa event over")
		}
	}()

	for {
		n, err := conn.Read(buffer[:(len(buffer)-1)])

		if err != nil {
			utils.Log(conn.RemoteAddr().String(), " connection error: ", err)
			connOK = false
			return
		}
		tmpBuffer = utils.Depack(append(tmpBuffer, buffer[:n]...))

		utils.TaskDeliver(tmpBuffer,conn)
		//for test
		//var m common.Msg
		//json.Unmarshal(tmpBuffer,&m)
		//utils.Log("liueh socket ", <- common.HttpToSocket)
		//common.SocketToHttp <- m
		//str := "to client"
		//_, err = conn.Write([]byte(str))
		//if err != nil {
		//	log.Println(err)
		//}
		//utils.Log("byte = ", []byte(str))

		//start heartbeating
		go utils.HeartBeating(conn,messnager,timeout)
		//check if get message from client
		go utils.GravelChannel(tmpBuffer,messnager)

	}
}

