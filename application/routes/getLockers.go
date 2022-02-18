package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrepmagalhaes/redirector/application/types"
	"github.com/gin-gonic/gin"
)

func GetLockers(c *gin.Context) {

	client := &http.Client{}
	var data types.GetLockersParams

	if err := c.BindJSON(&data); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("param", data.Codigo_de_MSG)

	json_data, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req, err := http.NewRequest("POST", "http://ec2-34-206-134-100.compute-1.amazonaws.com:8008/msg/v01/lockers", bytes.NewBuffer(json_data))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", c.Request.Header.Get("Authorization"))

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result types.GetLockersResponse
	var result2 types.GetLockersResponseError
	decoder := json.NewDecoder(resp.Body)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()
	fmt.Println(newStr)

	if err := decoder.Decode(&result); err != nil {
		if err := decoder.Decode(&result2); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("result2", result2.Detail)
		c.IndentedJSON(http.StatusOK, &result2)
	}

	fmt.Println("result", result.Codigo_de_MSG)

	c.IndentedJSON(http.StatusOK, &result)
}
