// Go 1.2
// go run helloworld_go.go

package main

import (
    . "fmt" // Using '.' to avoid prefixing functions with their package names
    . "runtime" // This is probably not a good idea for large projects...
    . "time"
      "net"
)

var adr = "129.241.187.161"
var port = 33546


//addr = new InternetAddress(serverIP, serverPort) 
//sock = new Socket(tcp) // TCP, aka SOCK_STREAM
//sock.connect(addr)
// use sock.recv() and sock.send()


func main() {

raddr,err := net.ResolveTCPAddr("tcp", adr) //(*TCPAddr, error)

serv,err2 := net.DialTCP("tcp", nil, raddr) //(*TCPConn, error)

if err !=nil...



}
