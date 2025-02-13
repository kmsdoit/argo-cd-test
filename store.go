package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var store = make(map[string]string)

type StoreBody struct {
	Value string `json:"value"`
}

func Put(key string, value string) error {
	store[key] = value
	return nil
}

func Get(key string) (string, error) {
	value, exists := store[key]
	if !exists {
		return "", http.ErrNoLocation // 키가 없을 경우 오류 반환
	}
	return value, nil
}

func Delete(key string) {

	if _, exists := store[key]; exists {
		delete(store, key)
	}
}

func KeyValuePutHandler(c *gin.Context) {
	var body StoreBody
	key := c.Param("key")

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(body, key)

	if err := Put(key, body.Value); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func KeyValueGetHandler(c *gin.Context) {
	key := c.Param("key")

	value, err := Get(key)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"value": value})
}

func KeyValueDeleteHandler(c *gin.Context) {
	key := c.Param("key")
	Delete(key)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
