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

func GetLockers(c *gin.Context) {

	client := &http.Client{}
	var data types.GetLockersParams

	if err := c.BindJSON(&data); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("param", data)

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

	// if resp.Status == "200" {
	// 	var result types.GetLockersResponse
	// 	if err := json.Unmarshal(respData, &result); err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// 	c.IndentedJSON(http.StatusOK, &result)
	// 	return
	// }
	// elif resp.Status == "401" {
	// 	var result types.GetLockersResponseError
	// 	if err := json.Unmarshal(respData, &result); err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// 	c.IndentedJSON(http.StatusOK, &result)
	// 	return
	// }
	// else {

	// }

	// var result types.GetLockersResponse
	// var result2 types.GetLockersResponseError

	// if err := json.Unmarshal(respData, &result); err != nil {
	// 	fmt.Println(respData)
	// 	fmt.Println(result)
	// 	if err := json.Unmarshal(respData, &result2); err != nil {
	// 		fmt.Println(respData)
	// 		fmt.Println(result2)
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// 	fmt.Println(respData)
	// 	fmt.Println(result2)
	// 	c.IndentedJSON(http.StatusOK, &result2)
	// 	return
	// }

	// c.IndentedJSON(http.StatusOK, &result)

	// resp2 := new(http.Response)
	// resp2 = resp

	// var result types.GetLockersResponse
	// var result2 types.GetLockersResponseError
	// decoder := json.NewDecoder(resp.Body)
	// decoder2 := json.NewDecoder(resp2.Body)

	// json.Unmarshal(decoder, &result)

	// if err := decoder.Decode(&result); err != nil {
	// 	fmt.Println("asdfasdfasdf")

	// 	if err := decoder2.Decode(&result2); err != nil {

	// 		fmt.Println("birl")
	// 		fmt.Println("resp2", resp2.Body)
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// 	fmt.Println("result2", result2.Detail)
	// 	c.IndentedJSON(http.StatusOK, &result2)
	// 	return
	// }

	// fmt.Println("result len: ", result)
	// fmt.Println("result", result.Codigo_de_MSG)

}
