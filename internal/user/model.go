package user

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string        `bson:"firstName" json:"firstName"`
	LastName  string        `bson:"lastName" json:"lastName"`
	Email     string        `bson:"email" json:"email"`
	Password  string        `bson:"password" json:"password"`
	Role      string        `bson:"role" json:"role"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
}
