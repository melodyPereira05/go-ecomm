package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name      *string            `json:"first_name" validate:"required,min=2,max=30"`
	Last_Name       *string            `json:"last_name"  validate:"required,min=2,max=30"`
	Password        *string            `json:"password"   validate:"required,min=6"`
	Email           *string            `json:"email"      validate:"email, required"`
	Phone           *string            `json:"phone"      validate:"required"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refresh_token"`
	Created_At      time.Time          `json:"created_at"`
	Updated_At      time.Time          `json:"updated_at"`
	User_ID         string             `json:"user_id"`
	UserCart        []ProductUser      `json:"usercart" bson:"usercart"`
	Address_Details []Address          `json:"address_details" bson:"address_details"`
	Order_Status    []Order            `json:"order_status" bson: "order_status"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name"bson:"product_name"`
	Price        *uint64            `json:"price"`
	Image        *string            `json:"image"`
	Rating       *uint64            `json:"rating"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        *uint64            `json:"price"`
	Rating       *uint64            `json:"rating"`
	Image        *string            `json:"image"`
}

type Address struct {
	Address_ID primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house"`
	Street     *string            `json:"street"`
	City       *string            `json:"city"`
	Pincode    *string            `json:"pincode"`
}

type Order struct {
	order_ID       primitive.ObjectID `bson:"_id"`
	order_Cart     []ProductUser      `json:"order_cart" bson:"order_id"`
	Ordered_At     time.Time          `json:"ordered_at"`
	Payment_Method Payment            `json:"payment_method"`
	Discount       *int               `json:"discount"`
	Price          *int               `json:"price"`
}

type Payment struct {
	Digital bool `json:"digital"`
	COD     bool `json:"cod"`
}
