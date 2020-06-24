package main

import (
	//"./pack"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func main() {
	fmt.Println("Hello")
	var A interface{}
	A = 1
	fmt.Println(reflect.ValueOf(A), reflect.TypeOf(A))
	A = "hello"
	reflectDemo()
	fmt.Println(reflect.TypeOf(A), reflect.ValueOf(A))
	switch A.(type) {
	case string:
		fmt.Println("A is string")
	}
	var lover Love
	lover = Person{"ww", 28, false}
	ExtractReflect(lover)
	//fmt.Println(pack.To(), "to is called")
	router := gin.Default()

	router.GET("/", hello)
	group := router.Group("/group")
	group.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"root": "love ww",
		})
	})
	group.POST("/post", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"post": "SUCC",
		})
	})
	err := router.Run()
	if err != nil {
		fmt.Println("err occurred", err)
	}
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, gin")
}

type Person struct {
	Name string
	Age  int
	Sex  bool
}

type Love interface {
	Llove()
}

func (person Person) Llove() {
	fmt.Printf("%s love me\n", person.Name)
}

func reflectDemo() {
	var f float64 = 1.23
	value := reflect.ValueOf(f)
	convert := value.Interface().(float64)
	fmt.Println(convert)

}

func ExtractReflect(input interface{}) {
	fmt.Println(reflect.TypeOf(input), reflect.TypeOf(input).Kind())
	//inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)
	fmt.Println(inputValue)
	act:=inputValue

	for i := 0; i < act.NumField(); i++ {
		fmt.Println(act.Field(i).CanSet(), act.Field(i).Kind(), act.Field(i).Type())
	}

	fmt.Println(inputValue)
	for i := 0; i < inputValue.NumMethod(); i++ {
		fmt.Println(inputValue.Method(i).Kind(), inputValue.Method(i).Type())
		inputValue.Method(i).Call(make([]reflect.Value,0))
	}

}
