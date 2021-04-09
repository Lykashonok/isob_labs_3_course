package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
)

// User - struct for user, which gave username and password
type User struct {
	username string
	password string
}

var (
	users []User
	keys  = make(map[string]string, 0)
)

func findUser(users []User, user User) bool {
	// encryptedPassword := Encrypt(user.password, user.password)
	for _, userToFind := range users {
		if user.password == userToFind.password {
			return true
		}
	}
	return false
}

const (
	// HOST - host ip
	HOST = "127.0.0.1"
	// PORT - port
	PORT = "3333"
	// TYPE - tcpip udp
	TYPE = "tcp"
	// DELIMITER - char which separate values in request
	DELIMITER = ';'
)

const letterBytes = "01"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	users = append(users, User{"vlad", "1001111000100110100111110101101011111010010011011011101101110000"})
	// Listen for incoming connections.
	l, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + HOST + ":" + PORT)

	conn, _ := l.Accept()
	for {
		err = handleRequest(conn)
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			os.Exit(1)
			break
		}
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) error {
	// Make a buffer to hold incoming data.
	// buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return err
	}
	// fmt.Printf("Recieved: %v", string(message))
	request := strings.Split(string(message), string(DELIMITER))
	response := "responsed"

	fmt.Printf("Recieved: %v\n", request)

	// Send a response back to person contacting us.
	conn.Write([]byte(response + "\n"))
	// Close the connection when you're done with it.
	// conn.Close()
	return nil
}