package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/andrepmagalhaes/redirector/application/types"
	"github.com/gin-gonic/gin"
)

func PostLockers(c *gin.Context) {

	client := &http.Client{}
	var data types.GetLockersParams

	if err := c.BindJSON(&data); err != nil {
		fmt.Println(err.Error())
		return
	}

	json_data, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req, err := http.NewRequest("POST", "http://137.184.137.84:8000/msg/v01/lockers", bytes.NewBuffer(json_data))

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

	respData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("status: ", resp.Status)
	fmt.Println("statusCode: ", resp.StatusCode)
	fmt.Println("data", respData)
	switch resp.StatusCode {
	case 200:
		{
			var result types.GetLockersResponse
			if err := json.Unmarshal(respData, &result); err != nil {
				fmt.Println(err.Error())
				return
			}
			c.IndentedJSON(resp.StatusCode, &result)
			return
		}
	case 401:
		{
			var result types.Unathorized
			if err := json.Unmarshal(respData, &result); err != nil {
				fmt.Println(err.Error())
				return
			}
			c.IndentedJSON(resp.StatusCode, &result)
			return
		}
	}

	c.IndentedJSON(resp.StatusCode, "Unhandled Status")

}
