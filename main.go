package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/andrepmagalhaes/redirector/application"
)

func main() {

	// err := godotenv.Load(".env")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	application.Init(port)

	// response, err := getAuth()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(response)

}

// type Response struct {
// 	ID     string `json:"id"`
// 	Joke   string `json:"joke"`
// 	Status int    `json:"status"`
// }

// func test() (*Response, error) {

// 	client := &http.Client{}

// 	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)

// 	if err != nil {
// 		return &Response{}, err
// 	}

// 	req.Header.Add("Accept", "application/json")
// 	req.Header.Add("Content-Type", "application/json")

// 	resp, err := client.Do(req)

// 	if err != nil {
// 		return &Response{}, err
// 	}

// 	defer resp.Body.Close()

// 	var response Response

// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		return &Response{}, err
// 	}

// 	return &response, nil
// }
