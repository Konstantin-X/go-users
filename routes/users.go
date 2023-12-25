package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-users/models"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type UserData struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
	Pass string `json:"pass"`
}

func CreateUser(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var data UserData

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong: " + err.Error(),
		})
	}

	// put into the DB
	fmt.Printf("User: %s %s\n", data.Name, data.Mail)

	password, cryptErr := bcrypt.GenerateFromPassword([]byte(data.Pass), 14)
	if cryptErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + cryptErr.Error(),
		})
	}

	newUser := models.User{
		Name:     data.Name,
		Email:    data.Mail,
		Password: string(password),
	}

	result, userErr := models.CreateUser(&newUser)
	if userErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Create user error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"userId": result.ID,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	fmt.Printf("id: %d\n", id)

	// delete this note from our DB
	return c.JSON(fiber.Map{
		"message": "Note successfully deleted",
	})
}

func ReadUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	fmt.Printf("id: %d\n", id)
	// SQL query for the actual note

	user := models.User{
		ID:       1,
		Name:     "Name",
		Email:    "",
		IsAdmin:  false,
		Password: "1111",
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	fmt.Printf("id: %d\n", id)

	// update in the DB the note
	return c.JSON(fiber.Map{
		"message": "Updated Successfully!",
		//"message": fmt.Printf("Updated Successfully! %d", id),
	})
}

func ListUsers(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)

	users := models.ListUsers(intPage, intLimit)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(users), "users": users})
}
