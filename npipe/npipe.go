package main

import (
	"bufio"
	"fmt"
	"github.com/natefinch/npipe"
	"net"
	"os"
	"time"
)


const serverString = "The Industrial Revolution and its consequences have been a disaster" +
"for the human race. They have greatly increased the life-expectancy of" +
"those of us who live in \"advanced\" countries, but they have" +
"destabilized society, have made life unfulfilling, have subjected" +
"human beings to indignities, have led to widespread psychological" +
"suffering (in the Third World to physical suffering as well) and have" +
"inflicted severe damage on the natural world. The continued" +
"development of technology will worsen the situation. It will certainly" +
"subject human beings to greater indignities and inflict greater damage" +
"on the natural world, it will probably lead to greater social" +
"disruption and psychological suffering, and it may lead to increased" +
"physical suffering even in \"advanced\" countries." 

const clientString = "The industrial-technological system may survive or it may break" +
"down. If it survives, it MAY eventually achieve a low level of" +
"physical and psychological suffering, but only after passing through a" +
"long and very painful period of adjustment and only at the cost of" +
"permanently reducing human beings and many other living organisms to" +
"engineered products and mere cogs in the social machine. Furthermore," +
"if the system survives, the consequences will be inevitable: There is" +
"no way of reforming or modifying the system so as to prevent it from" +
"depriving people of dignity and autonomy."

// Use Dial to connect to a server and read messages from it.
func doDial() {
	conn, err := npipe.DialTimeout(`\\.\pipe\mypipe`, 250 * time.Millisecond)
	if err != nil {
		fmt.Printf("Dial error: %v", err)
		return
	}
	if _, err := fmt.Fprintln(conn, clientString); err != nil {
		fmt.Printf("Send error: %v", err)
		return
	}
	r := bufio.NewReader(conn)
	_, err = r.ReadString('\n')
	if err != nil {
		fmt.Printf("Receive error: %v", err)
	}
//	fmt.Println(msg)
}

// Use Listen to start a server, and accept connections with Accept().
func doListen() {
	ln, err := npipe.Listen(`\\.\pipe\mypipe`)
	if err != nil {
		fmt.Printf("Listen error: %v", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Accept error: %v", err)
			continue
		}
		
		time.Sleep(10 * time.Millisecond)

		// handle connection like any other net.Conn
		go func(conn net.Conn) {
			r := bufio.NewReader(conn)
			_, err := r.ReadString('\n')
			if err != nil {
				fmt.Printf("Receive error: %v", err)
				return
			}
			if _, err := fmt.Fprintln(conn, serverString); err != nil {
				fmt.Printf("Send error: %v", err)
				return
			}

//			fmt.Println(msg)
		}(conn)
	}
}

func main() {
    if len(os.Args) > 1 {
		doListen()
	} else {
		for i := 0; i < 400; i++ {
			doDial()
		}
	} 
}