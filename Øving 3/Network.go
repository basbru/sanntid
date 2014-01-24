//go run oving3.go
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
        //Get full server address, check for errors
        //servadr, err := net.ResolveTCPAddr("tcp", serveradress+":"+port)
        //connect to server
        conn, err := net.Dial("tcp", "129.241.187.161:33546")
        //handle errors
        if err != nil {
		fmt.Println("Shit aint working..")
		return
        }
	fmt.Println("Got connection")
	

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	
	
	for {
		line, err := rw.ReadString(0)		
		if err != nil {
			fmt.Println("Got error:", err)
			break
		}
		fmt.Println("FROM REMOTE]", line)
		rw.WriteString("Hei på deg")
		rw.WriteByte(0)
		rw.Flush()
		time.Sleep(1*time.Second)
	}

	conn.Close()



/* else {
		reader := bufio.NewReader(conn)
                //make a buffer for recieved info
                read := make([]byte, 1024)
                //read from server
                conn.Read(read)
                s:=string(read)
                //print server message
                fmt.Println(s)
                msg := "Connect to: 78.91.9.210:33546\x00"
                fmt.Scanln(&msg)
                for msg != "Stop"{
                        //send message to server
                        conn.Write([]byte(msg+"\x00"))
                        read = make([]byte, 1024)
                        conn.Read(read)
                        s = string(read)
                        fmt.Println(s)
                        //needs fixing
			msg,crap = reader.ReadString(0)
                        //msg, crap1, crap2 := bufio.ReadLine()
                        time.Sleep(1000*time.Millisecond)

                }
        }
*/
