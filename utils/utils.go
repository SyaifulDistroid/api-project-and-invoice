package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Response(c *gin.Context, code int, status interface{}, message interface{}, result interface{}) {
	c.JSON(code, gin.H{"Status": status, "Message": message, "Result": result})
}

func JoinVarAndValStruct(v reflect.Value) string {
	var NewStruct []string
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		NewStruct = append(NewStruct, fmt.Sprintf("%v:%v", typeOfS.Field(i).Name, v.Field(i).Interface()))
	}
	strJoin := strings.Join(NewStruct, ",")
	return strJoin
}

func StringToInt(str string) int {
	id, err := strconv.Atoi(str)
	if err != nil {
		logrus.Println(err)
		return 0
	}
	return id
}
