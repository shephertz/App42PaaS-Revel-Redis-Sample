package controllers

import (
  "App42PaaS-Revel-Redis-Sample/app/models"
  "github.com/revel/revel"
  "fmt"
)

type App struct {
	ModelController
}

func (c App) New() revel.Result {
	greeting := "Revel Redis Sample App"
	return c.Render(greeting)
}

func (c App) Index(name string) revel.Result{
  fmt.Println("In instert data.")
  vals := []string{name}
	for _, v := range vals {
	    err	:= client.Rpush("username", []byte(v))
		  fmt.Println("Redis Error:", err)

	}
  
  
  fmt.Println("In get data.")
	usernames,_ := client.Lrange("username", 0, 10000)
	
  users := []models.User{}
  for _, v := range usernames {
	  user := models.User{}
	  user.Username = string(v)
	  users = append(users, user)
  }
	
  return c.Render(users)
}
