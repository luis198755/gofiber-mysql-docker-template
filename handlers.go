package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	// ErrIDNotFound - The passed ID could not be found in the database.
	ErrIDNotFound = Error("ID not found")
)

// Error allows to define error messages as constants at compile time
type Error string

// ErrorMessage is a wrapper that can be used to wrap error messages
// as JSON resonses.
type ErrorMessage struct {
	Error string `json:"error,omitempty"`
}

// Error is the function that satisfies the error interface.
func (err Error) Error() string {
	return string(err)
}

// GetApple by ID
func GetApple(c *fiber.Ctx) error {

	id, err := fetchInt("id", c)
	if err != nil {
		return c.JSON(ErrorMessage{err.Error()})
	}

	var apple Apple
	result := db.Find(&apple, id)
	err = result.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(ErrorMessage{ErrIDNotFound.Error()})
		}
		return err
	}

	return c.JSON(&apple)
}

// GetApples returns all apples
func GetApples(c *fiber.Ctx) error {
	var apples []Apple
	result := db.Find(&apples)

	err := result.Error
	if err != nil {
		return c.JSON(ErrorMessage{err.Error()})
	}

	return c.JSON(apples)
}

// AddApple creates a new apple.
func AddApple(c *fiber.Ctx) error {

	var apple Apple
	err := c.BodyParser(&apple)
	if err != nil {
		return c.JSON(ErrorMessage{err.Error()})
	}

	result := db.Create(&apple)
	err = result.Error

	if err != nil {
		return c.JSON(ErrorMessage{err.Error()})
	}

	return c.JSON(apple)
}

// UpdateApple updates the proterties of apple with ID
func UpdateApple(c *fiber.Ctx) error {
	id, err := fetchInt("id", c)
	if err != nil {
		return c.JSON(ErrorMessage{err.Error()})
	}

	var apple Apple
	result := db.Find(&apple, id)
	err = result.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(ErrorMessage{ErrIDNotFound.Error()})
		}
	}

	err = c.BodyParser(&apple)
	if err != nil {
		return c.JSON(ErrorMessage{err.Error()})
	}

	result = db.Save(&apple)
	err = result.Error
	if err != nil {
		return err
	}

	return c.JSON(&apple)
}

// DeleteApple deletes apple with given ID
func DeleteApple(c *fiber.Ctx) error {
	var apple Apple

	id, err := fetchInt("id", c)
	if err != nil {
		return c.JSON(ErrorMessage{err.Error()})
	}

	result := db.Find(&apple, id)
	err = result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(ErrorMessage{ErrIDNotFound.Error()})
		}
		return err
	}

	result = db.Delete(&apple, id)
	err = result.Error
	if err != nil {
		return err
	}

	// return the deleted apple
	return c.JSON(&apple)
}

// FlushApples deletes apple with given ID
func FlushApples(c *fiber.Ctx) error {
	var apples []Apple

	result := db.Find(&apples)

	err := result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(ErrorMessage{ErrIDNotFound.Error()})
		}
		return err
	}

	result = db.Delete(&apples)
	err = result.Error
	if err != nil {
		return err
	}

	// return the deleted apple
	return c.JSON(&apples)
}
