package main

import (
	"encoding/json"
	"fmt"
	"gameapp/repository/mysql"
	"gameapp/service/userservice"
	"io"
	"net/http"
)

const (
	JwtSignKey = "jwt_secret"
)

func main() {
	//us := userservice.New(mysql.New())
	//fmt.Println(us.Login(userservice.LoginRequest{"09121131116", "12345678"}))

	mux := http.NewServeMux()
	mux.HandleFunc("/check", checkHandler)
	mux.HandleFunc("/users/register", userRegisterHandler)
	//http.HandleFunc("/users/register", userRegisterHandler)
	mux.HandleFunc("/users/login", userLoginHandler)
	mux.HandleFunc("/users/profile", userProfileHandler)
	http.ListenAndServe(":2020", mux)

	fmt.Println("server is listening on port 2020 ...")

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
		us := userservice.New(mysql.New(), JwtSignKey)
		if _, lErr := us.Register(uReq); lErr != nil {
			res.Write([]byte(
				fmt.Sprintf(`{"error": "%s"}`, lErr),
			))

			return
		}
		res.Write([]byte(
			fmt.Sprintf(`{"message": "user register is ok"}`),
		))
	}
}
func userLoginHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		fmt.Fprintf(res, `{"error":"invalid method"}`)

		return
	}

	if data, rErr := io.ReadAll(req.Body); rErr != nil {
		res.Write([]byte(
			fmt.Sprintf(`{"error":"%s"}`, rErr),
		))
	} else {
		uReq := userservice.LoginRequest{}
		if jErr := json.Unmarshal(data, &uReq); jErr != nil {
			res.Write([]byte(
				fmt.Sprintf(`{"error": "%s"}`, jErr),
			))

			return
		}
		us := userservice.New(mysql.New(), JwtSignKey)
		if uRes, lErr := us.Login(uReq); lErr != nil {
			res.Write([]byte(
				fmt.Sprintf(`{"error": "%s"}`, lErr),
			))

			return
		} else {
			data, _ = json.Marshal(uRes)
			res.Write(data)
		}

	}
}

// curl --request GET \
// --url http://localhost:2020/check
func checkHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, `{"message":"it's okey !"}`)
}

func userProfileHandler(res http.ResponseWriter, req *http.Request) {

	// sessionID := req.Header.Get("sessionID")
	// validate sessionID by database and get userID

	//jwtToken := req.Header.Get("Authorization")
	// validate jwt token and retrieve user ID from payload

	if data, rErr := io.ReadAll(req.Body); rErr != nil {
		res.Write([]byte(
			fmt.Sprintf(`{"error":"%s"}`, rErr),
		))
	} else {
		uReq := userservice.ProfileRequest{}
		if jErr := json.Unmarshal(data, &uReq); jErr != nil {
			res.Write([]byte(
				fmt.Sprintf(`{"error": "%s"}`, jErr),
			))

			return
		}
		us := userservice.New(mysql.New(), JwtSignKey)
		if pRes, lErr := us.Profile(uReq); lErr != nil {
			res.Write([]byte(
				fmt.Sprintf(`{"error": "%s"}`, lErr),
			))

			return
		} else {
			//res.Write([]byte(
			//	fmt.Sprintf(`{"message": "user name is %v"}`, pRes.Name),
			//))
			data, jErr := json.Marshal(pRes)
			if jErr != nil {
				fmt.Println("error in json-marshal: %w", jErr)
			} else {
				res.Write(data)
			}

		}

	}
}
