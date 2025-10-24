package user

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string        `bson:"firstName" json:"firstName" validate:"required,min=2"`
	LastName  string        `bson:"lastName" json:"lastName"`
	Email     string        `bson:"email" json:"email" validate:"required,email"`
	Password  string        `bson:"password" json:"password" validate:"required,min=8,max=32"`
	Role      string        `bson:"role" json:"role" validate:"oneof=admin user" default:"user"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
}

type GetProfileResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	User    *User  `json:"user"`
}
