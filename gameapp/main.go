package main

import (
	"encoding/json"
	"fmt"
	"gameapp/repository/mysql"
	"gameapp/service/userservice"
	"io"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/check", checkHandler)
	mux.HandleFunc("/users/register", userRegisterHandler)
	//http.HandleFunc("/users/register", userRegisterHandler)
	http.ListenAndServe(":2020", mux)
	log.Println("server is listening on port 2020 ...")

}

// user curl or insomnia
// curl --request POST \
// --url http://localhost:2020/users/register \
// --header 'Content-Type: application/json' \
// --data '{
// "Name": "mehdad",
// "PhoneNumber": "09376231226"
// }'
func userRegisterHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		fmt.Fprintf(res, `{"error":"invalid method"}`)

		return
	}

	if data, rErr := io.ReadAll(req.Body); rErr != nil {
		res.Write([]byte(
			fmt.Sprintf(`{"error":"%s"}`, rErr),
		))
	} else {
		uReq := userservice.RegisterRequest{}
		json.Unmarshal(data, &uReq)
		us := userservice.New(mysql.New())
		res.Write([]byte(
			fmt.Sprint(us.Register(uReq)),
		))

	}
}

// curl --request GET \
// --url http://localhost:2020/check
func checkHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, `{"message":"it's okey !"}`)
}
