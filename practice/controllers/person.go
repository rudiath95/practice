package controllers

import (
	"net/http"
	"practice/database"
	"practice/repository"
	"practice/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPerson(c *gin.Context) {
	var (
		result gin.H
	)

	persons, err := repository.GetAllPerson(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": persons,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertPerson(c *gin.Context) {
	var person structs.Person

	err := c.ShouldBindJSON(&person)
	if err != nil {
		panic(err)
	}

	err = repository.InsertPerson(database.DbConnection, person)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Person",
	})

}

func UpdatedPerson(c *gin.Context) {
	var person structs.Person
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&person)
	if err != nil {
		panic(err)
	}

	person.ID = int64(id)

	err = repository.UpdatedPerson(database.DbConnection, person)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Person",
	})
}

func DeletedPerson(c *gin.Context) {
	var person structs.Person
	id, err := strconv.Atoi(c.Param("id"))

	person.ID = int64(id)

	err = repository.DeletedPerson(database.DbConnection, person)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Deleted Person",
	})
}
