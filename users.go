package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id       *int64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u user) create() (int64, error) {
	stmt, preparErr := db.Prepare("insert into users(name,email,password) values(:1,:2,:3)")

	if preparErr != nil {
		return 0, fmt.Errorf("prepare statement: %v", preparErr)
	}
	defer stmt.Close()
	result, execErr := stmt.Exec(u.Name, u.Email, u.Password)
	// result, err := db.Exec("insert into users(name,email,password) values(?,?,?);", u.Name, u.Email, u.Password)
	if execErr != nil {
		return 0, fmt.Errorf("Exec: %v", execErr)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("create User: %v", err)
	}
	return id, nil
}

func getUsers(c *gin.Context) {
	users, queryError := findUsers()
	if queryError != nil {
		log.Fatal("queryErr:", queryError)

	}
	fmt.Println(users)
	c.IndentedJSON(http.StatusOK, users)
}
func postUser(c *gin.Context) {
	var user user
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"JSONError": err.Error(),
		})
		return
	}
	id, insertErr := user.create()
	if insertErr != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"insertErr": insertErr.Error(),
		})
		return
	}
	user.Id = &id
	c.IndentedJSON(http.StatusOK, user)
}

func findUsers() ([]user, error) {
	var users []user
	rows, err := db.Query("Select * from users")
	if err != nil {
		return nil, fmt.Errorf("getusers : %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user user
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, fmt.Errorf("getusers : %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getusers : %v", err)
	}
	return users, nil
}
