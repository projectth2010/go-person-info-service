package models

import "time"

type User struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName      string             `json:"firstName" bson:"firstName" validate:"required,min=2,max=50"`
	LastName       string             `json:"lastName" bson:"lastName" validate:"required,min=2,max=50"`
	Email          string             `json:"email" bson:"email" validate:"required,email"`
	Password       string             `json:"password,omitempty" bson:"password,omitempty" validate:"required,min=8"`
	Phone          string             `json:"phone,omitempty" bson:"phone,omitempty"`
	DateOfBirth    *time.Time         `json:"dateOfBirth,omitempty" bson:"dateOfBirth,omitempty"`
	Gender         string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Address        Address            `json:"address,omitempty" bson:"address,omitempty"`
	ProfilePicture string             `json:"profilePicture,omitempty" bson:"profilePicture,omitempty"`
	IsActive       bool               `json:"isActive" bson:"isActive"`
	LastLogin      *time.Time         `json:"lastLogin,omitempty" bson:"lastLogin,omitempty"`
	CreatedAt      time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type Address struct {
	Street     string `json:"street,omitempty" bson:"street,omitempty"`
	City       string `json:"city,omitempty" bson:"city,omitempty"`
	State      string `json:"state,omitempty" bson:"state,omitempty"`
	PostalCode string `json:"postalCode,omitempty" bson:"postalCode,omitempty"`
	Country    string `json:"country,omitempty" bson:"country,omitempty"`
}

type UpdateProfileRequest struct {
	FirstName   string    `json:"firstName,omitempty"`
	LastName    string    `json:"lastName,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	DateOfBirth time.Time `json:"dateOfBirth,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	Address     Address   `json:"address,omitempty"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"currentPassword" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	FirstName string `json:"firstName" validate:"required,min=2,max=50"`
	LastName  string `json:"lastName" validate:"required,min=2,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}
