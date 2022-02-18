package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrepmagalhaes/redirector/application/types"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {

	var data types.GetAuthParams

	if err := c.BindJSON(&data); err != nil {
		fmt.Println(err.Error())
		return
	}

	json_data, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req, err := http.Post("http://ec2-34-206-134-100.compute-1.amazonaws.com:8008/login", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer req.Body.Close()

	var result types.GetAuthResponse

	if err := json.NewDecoder(req.Body).Decode(&result); err != nil {
		fmt.Println(err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, &result)
}
