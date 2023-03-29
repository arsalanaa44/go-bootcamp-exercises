package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"todo_cli/client/clientservices/taskclient"
	"todo_cli/client/clientservices/userclient"
	"todo_cli/delivery/deliveryparam"
)

var cRequest = deliveryparam.ClientRequest{}
var logged = false

func main() {

	command := flag.String("command", "no command", "command to run")
	flag.Parse()
	for {

		cRequest = deliveryparam.ClientRequest{}
		runCommand(*command)

		fmt.Println("please enter another command :")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()

	}
}

func runCommand(command string) {

	cRequest.Command = command
	if !logged {
		fmt.Println("you're not logged in")
		login()

		return
	}
	switch command {
	case "login":
		{

			login()
		}
	case "create-task":
		{
			CreateTask()
		}
	case "exit":
		{
			os.Exit(0)
		}
	default:
		{

			fmt.Println("invalid command !")
		}
	}
}

func login() {

	cRequest.LoginRequest = userclient.NewClientService().Login()
	if message, sErr := sendDataAndGetResponse(); message == "logged in successfully" {
		logged = true
		fmt.Println(message, sErr)
	}

}

func CreateTask() {
	if cReq, cErr := taskclient.NewClientService().CreateTask(); cErr != nil {
		fmt.Println("error in creating task :", cErr)
	} else {
		cRequest.CreateTaskRequest = cReq
		fmt.Println(sendDataAndGetResponse())
	}

}

func sendDataAndGetResponse() (string, error) {
	const (
		network = "tcp"
		address = "127.0.0.1:2022"
	)

	var connection net.Conn
	if c, dErr := net.Dial(network, address); dErr != nil {
		fmt.Println("dError", dErr)
	} else {
		connection = c
	}
	fmt.Println("client and server address :", connection.LocalAddr(), connection.RemoteAddr())

	if sData, sErr := json.Marshal(cRequest); sErr != nil {

		return "", fmt.Errorf("error in marshalization %v", sErr)
	} else {
		fmt.Println(connection.Write(sData))
	}

	// process request
	var rData = make([]byte, 1024)
	if numberOfReadData, rErr := connection.Read(rData); rErr != nil {

		return "", fmt.Errorf("error in reading data %v", rErr)
	} else {

		response := deliveryparam.ClientResponse{}
		fmt.Println(json.Unmarshal(rData[:numberOfReadData], &response))

		return response.Data, nil
	}
}
