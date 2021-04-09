package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	// HOST - host ip
	HOST = "localhost"
	// PORT - port
	PORT = "3333"
	// SERVICEPORT - service
	SERVICEPORT = "4444"
	// TYPE - tcpip udp
	TYPE = "tcp"
	// DELIMITER - char which separate values in request
	DELIMITER = ';'
)

func main() {
	// go server()
	// Listen for incoming connections.
	conn, err := net.Dial(TYPE, HOST+":"+PORT)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	var (
		username, password string = "vlad", "1010101111001101111001101010101111001101000100110010010100110110"
		stage              int    = 1
		stringToRequest    string
	)
	// fmt.Println("Connected, write your login:")
	// fmt.Scanln(&username)
	// fmt.Println("password:")
	// fmt.Scanln(&password)

	// conn.Write([]byte(username))

	// connbuf := bufio.NewReader(conn)
	message := ""
	for {
		//read in input from stdin
		if stage == 1 {
			stringToRequest = string(DELIMITER) + username + string(DELIMITER) + password + string(DELIMITER) + string('\n')
			fmt.Printf("Message to auth server: 1 %v", stringToRequest)
			stringToRequest = strconv.Itoa(stage+1) + stringToRequest
			stage++
		} else if stage == 3 {
			stage++
			encryptedKey := strings.Split(message, string(DELIMITER))[1]
			fmt.Printf("Decrypting key from auth server\n")
			decryptedKey := Decrypt(encryptedKey)
			fmt.Printf("Decrypted\n")
			stringToRequest = strconv.Itoa(stage) + string(DELIMITER) + decryptedKey + string(DELIMITER) + username + string(DELIMITER) + string("\n")
			fmt.Printf("Sending key and name to tgs\n")
		} else if stage == 5 {
			// stage++
			// username := strings.Split(message, string(DELIMITER))[1]
			// time := strings.Split(message, string(DELIMITER))[2]
			// fmt.Printf("Got right to access to service server\n")
			// if err != nil {
			// 	fmt.Println("Error listening:", err.Error())
			// 	os.Exit(1)
			// }
			// fmt.Printf("Sending key time and username to service server\n")
			// stringToRequest = strconv.Itoa(stage) + string(DELIMITER) + username + string(DELIMITER) + time + string(DELIMITER) + string("\n")
			// break

		}

		conn.Write([]byte(stringToRequest))

		// reader := bufio.NewReader(os.Stdin)
		// fmt.Print("Text to send: ")
		// text, _ := reader.ReadString('\n')

		//send to socket
		// fmt.Fprint(conn, text+"\n")
		// conn.Write([]byte(text))

		//listen for reply
		message, _ = bufio.NewReader(conn).ReadString('\n')
		stage, _ = strconv.Atoi(strings.Split(message, string(DELIMITER))[0])
		fmt.Printf("Message from server: %v", message)
	}
	conn.Close()
	serviceConnection, err := net.Dial(TYPE, HOST+":"+SERVICEPORT)
	defer serviceConnection.Close()

	serviceConnection.Write([]byte(stringToRequest))
	message, _ = bufio.NewReader(serviceConnection).ReadString('\n')
	stage, _ = strconv.Atoi(strings.Split(message, string(DELIMITER))[0])
	fmt.Printf("Message from server: %v", message)
}
