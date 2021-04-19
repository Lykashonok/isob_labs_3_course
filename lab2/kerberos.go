package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
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
	HOST = "localhost"
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
	response := processStages(request)
	// fmt.Printf("Recieved: %v\n", request)

	// Send a response back to person contacting us.
	conn.Write([]byte(response + "\n"))
	// Close the connection when you're done with it.
	// conn.Close()
	return nil
}

// returns value for next stage
func processStages(request []string) string {
	stage, _ := strconv.Atoi(request[0])
	response := ""
	timeLayout := "2006-01-02 15:04:05"
	if stage == 2 {
		// stage;username;userpassword
		// request[2] must be encrypted
		fmt.Printf("Encrypting password and trying to find in db\n")
		if findUser(users, User{request[1], Encrypt(request[2])}) {
			fmt.Printf("Found\nCreating key, storing it and send back in encrypted form\n")
			key := randStringBytes(64)
			keys[request[1]] = key
			encryptedKey := Encrypt(key)

			response = strconv.Itoa(stage+1) + string(DELIMITER) + encryptedKey + string(DELIMITER)
		} else {
			fmt.Errorf("Can't authenticate user with credentials %v, %v", request[1], request[2])
		}
	} else if stage == 4 {
		// stage;key;username;
		fmt.Printf("Trying to find decrypted key on tgs\n")
		if keys[request[2]] == request[1] {
			// username + time
			fmt.Printf("Found\nSending back key with time. User can have access to service server\n")
			response = strconv.Itoa(stage+1) + string(DELIMITER) + request[2] + string(DELIMITER) + time.Now().Format(timeLayout) + string(DELIMITER)
		} else {
			response = "error"
		}
	}
	fmt.Printf("Request from user: %v", request)
	return response
}
