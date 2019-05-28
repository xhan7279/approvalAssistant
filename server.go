package main

import (
	"fmt"

	"github.com/xhan7279/approvalAssistant/models"
)

func main() {
	c := models.CreateAccount("user", "pass")
	fmt.Println("hi", c)

}
