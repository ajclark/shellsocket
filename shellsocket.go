// Outrageously simple go app to bind bash to a tcp socket.

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

// Start bash(1) in interactive mode and pass in a command.
// Each command passed in is a new bash session.
func runABashCommand(command string) (string, error) {
	var out bytes.Buffer
	cmd := exec.Command("bash", "-i")
	cmd.Stdin = strings.NewReader(command)
	cmd.Stdout = &out
	error := cmd.Run()
	return out.String(), error
}

// Encapsulate runABashCommand function to execute it as a goroutine.
// Perform some basic error checking.
func run(conn net.Conn) {
	fmt.Fprintf(conn, "# ")
	for {
		cmd, _ := bufio.NewReader(conn).ReadString('\n')
		output, err := runABashCommand(string(cmd))
		if err != nil {
			errorString := err.Error()
			log.Printf(errorString)
		}
		fmt.Fprintf(conn, output)
		fmt.Fprintf(conn, "# ")
	}
}

func main() {
	// Set up usage
	portPtr := flag.String("port", "2222", "TCP port to listen on")
	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(2)
	}

	ln, err := net.Listen("tcp", ":"+*portPtr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		go run(conn)
	}
}
