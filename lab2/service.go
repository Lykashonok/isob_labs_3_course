package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
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

func isUserExists(users []User, username string) bool {
	for _, userToFind := range users {
		if username == userToFind.username {
			return true
		}
	}
	return false
}

func findUser(users []User, user User) bool {
	// encryptedPassword := Encrypt(user.password, user.password)
	encryptedPassword := user.password
	for _, userToFind := range users {
		if encryptedPassword == userToFind.password {
			return true
		}
	}
	return false
}

const (
	// HOST - host ip
	HOST = "localhost"
	// PORT - port
	PORT = "4444"
	// TYPE - tcpip udp
	TYPE = "tcp"
	// DELIMITER - char which separate values in request
	DELIMITER = ';'
)

const letterBytes = "1234567890ABCDEF"

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
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return err
	}
	request := strings.Split(string(message), string(DELIMITER))
	timeLayout := "2006-01-02 15:04:05"
	keyTime, _ := time.Parse(request[2], timeLayout)
	expiredTime := time.Now().Add(1 * time.Hour)
	if isUserExists(users, request[1]) && keyTime.Before(expiredTime) {
		fmt.Printf("keyTime is not expired and user was found\n")
		response := "some confidential service servers info"
		conn.Write([]byte(response + "\n"))
	}
	fmt.Printf("Request from user: %v", request)
	return nil
}
