package models

import (
	"fmt"
)

type User struct {
	ID       uint16 `gorm:"primaryKey; not null"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex; not null"`
	IsAdmin  bool   `gorm:"default=false"`
	Password string `gorm:"not null"`
}

func CreateUser(user *User) (*User, error) {
	newUser := User{
		Name:     "User 2",
		Email:    "user@local",
		IsAdmin:  false,
		Password: "123",
	}

	fmt.Printf("::  EXT user: %v \n", user)
	fmt.Printf("::  INT user: %v \n", newUser)

	fmt.Printf("Gloabl DB var: %v\n", DB)

	result := DB.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
	//return GetUser(user.ID)
}

func GetUser(userID uint16) (*User, error) {
	var user User

	result := DB.First(&user, userID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func ListUsers(page int, limit int) []User {
	offset := (page - 1) * limit

	var users []User

	results := DB.Limit(limit).Offset(offset).Find(&users)
	if results.Error != nil {
		fmt.Printf("ListUsers ERROR: %v\n", results.Error)
	}

	return users
}

func UpdateUser(user *User) error {
	result := DB.Save(user)
	if result.Error != nil {
		return result.Error
	}

	//return getUser(result.ID)
	return nil
}

func DeleteUser(user *User) error {
	result := DB.Delete(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
