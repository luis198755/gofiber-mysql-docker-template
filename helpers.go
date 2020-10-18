package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func fetchInt(key string, c *fiber.Ctx) (value int, err error) {
	valueStr := c.Params(key)

	log.Println(c)

	value, err = strconv.Atoi(valueStr)
	if err != nil {
		return 0, Error("missing or invalid parameter " + key)
	}
	return value, nil
}

func fetchNonEmptyString(key string, c *fiber.Ctx) (value string, err error) {
	value = c.Params(key)
	if value == "" {
		return "", Error("missing or empty parameter " + key)
	}
	return value, nil
}
