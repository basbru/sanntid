//Adr: 129.241.187.161
package main
import(
	"net"
	"fmt"
	"bufio"
    "time"
)

var serveradress = "129.241.187.161"
var port = "33546"

func main() {
	servadr, error1 := net.ResolveTCPAddr("tcp", serveradress+":"+port)
	conn, error2 := net.DialTCP("tcp", nil, servadr)
	if error1 != nil || error2 != nil{
		fmt.Println("Shit aint working..")
	} else {
		read := make([]byte, 1024)
		conn.Read(read)
		s:=string(read)
		fmt.Println(s)
		msg := "Connect to: 78.91.9.210:33546\x00"
		fmt.Scanln(&msg)
		for msg != "Stop"{
			conn.Write([]byte(msg+"\x00"))
			read = make([]byte, 1024)
			conn.Read(read)
			s = string(read)
			fmt.Println(s)
			msg, crap1, crap2 := bufio.ReadLine()
			time.Sleep(1000*time.Millisecond)

		}
	}
}