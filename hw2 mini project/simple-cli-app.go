package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	Name            string
	Address         string
	Number          string
	DateModified    string
	NumberOfWorkers int
	Id              int
	Region          int
}

var regionCode = 0

var data []Server

func main() {

	command := flag.String("c", "", "command")
	region := flag.String("r", "", "region")
	flag.Parse()
	regionCode = regionCheck(region)
	dataInitialization()
	for {
		runCommand(*command)
		fmt.Println("please enter another command :")
		fmt.Scanln(command)

	}

}
func regionCheck(region *string) int {
	var regionToCode = map[string]int{
		"tehran":  0,
		"isfahan": 1,
		"fars":    2,
	}
	regionCode, ok := regionToCode[*region]
	if ok != true {
		fmt.Println("invalid region !")
		os.Exit(0)
	}
	return regionCode
}
func runCommand(command string) {
	switch command {
	case "list":
		{
			list()
		}
	case "exit":
		{
			os.Exit(0)
		}
	case "get":
		{
			get()
		}
	case "create":
		{
			create()
		}
	case "edit":
		{
			edit()
		}
	case "status":
		{
			status()
		}
	default:
		{
			fmt.Println("invalid command !", command)
		}
	}
}
func list() {
	for _, v := range data {
		// if v.Region == regionCode {
		// 	fmt.Println(v)
		// 	fmt.Println()
		// }
		fmt.Println(v)
		fmt.Println()
	}
}
func get() {

	fmt.Println("please enter server-id")
	var id int
	fmt.Scanln(&id)
	for _, v := range data {

		if (v.Id == id) && (v.Region == regionCode) {
			fmt.Println(v)
			fmt.Println()
			return
		}
	}
	fmt.Println("ID not found !")
}
func create() {
	var srv Server
	fmt.Println("name:")
	fmt.Scanln(&srv.Name)
	fmt.Println("address:")
	fmt.Scanln(&srv.Address)
	fmt.Println("number:")
	fmt.Scanln(&srv.Number)
	srv.DateModified = time.Now().String()
	fmt.Println("numberOfWorkers:")
	fmt.Scanln(&srv.NumberOfWorkers)
	srv.Id = len(data)
	srv.Region = regionCode

	addToList(srv)
}
func addToList(s Server) {
	data = append(data, s)
	addTofile(s)
}
func addTofile(s Server) {
	fileName := "data.txt"

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err := f.WriteString(s.string()); err != nil {

		log.Fatal(err)
	}
}
func (s Server) string() string {
	out, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return "\n" + string(out)
}

func dataInitialization() {
	var dataUnit Server
	bytes, _ := os.ReadFile("data.txt")
	str := string(bytes)
	strArr := strings.Split(str, "\n")
	for _, str := range strArr {
		str = strings.ReplaceAll(str, "{", "")
		str = strings.ReplaceAll(str, "}", "")
		for _, s := range strings.Split(str, ",") {
			dataUnit.setDatavalue(strings.Split(s, ":"))
		}
		data = append(data, dataUnit)
	}
}

func (server *Server) setDatavalue(str []string) {
	switch str[0] {
	case "\"Name\"":
		server.Name = strings.ReplaceAll(str[1], "\"", "")
	case "\"Address\"":
		server.Address = strings.ReplaceAll(str[1], "\"", "")
	case "\"Number\"":
		server.Number = strings.ReplaceAll(str[1], "\"", "")
	case "\"DateModified\"":
		server.DateModified = strings.ReplaceAll(str[1], "\"", "")
	case "\"NumberOfWorkers\"":
		i, _ := strconv.Atoi(strings.ReplaceAll(str[1], "\"", ""))
		server.NumberOfWorkers = i
	case "\"Id\"":
		i, _ := strconv.Atoi(strings.ReplaceAll(str[1], "\"", ""))
		server.Id = i
	case "\"Region\"":
		i, _ := strconv.Atoi(strings.ReplaceAll(str[1], "\"", ""))
		server.Region = i
	}
}

func edit() {
	fmt.Println(`emmm, I have another work with higher priority
but this link is simply useful
https://www.socketloop.com/tutorials/golang-read-a-text-file-and-replace-certain-words`)
}
func status() {
	i, sum := 0, 0
	for _, v := range data {
		if v.Region == regionCode {
			sum += v.NumberOfWorkers
			i++
		}
	}
	fmt.Println("number of servers:", i, "\ntotal workers :", sum)
}
