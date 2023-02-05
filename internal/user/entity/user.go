package entity

import (
	"time"
)

type User struct {
	Id          string     `json:"id" gorm:"column:id;primary_key" bson:"_id" dynamodbav:"id" firestore:"-" avro:"id" validate:"required,max=40" match:"equal"`
	Username    string     `json:"username" gorm:"column:username" bson:"username" dynamodbav:"username" firestore:"username" avro:"username" validate:"required,username,max=100" match:"prefix"`
	Email       string     `json:"email" gorm:"column:email" bson:"email" dynamodbav:"email" firestore:"email" avro:"email" validate:"email,max=100" match:"prefix"`
	Phone       string     `json:"phone" gorm:"column:phone" bson:"phone" dynamodbav:"phone" firestore:"phone" avro:"phone" validate:"required,phone,max=18"`
	DateOfBirth *time.Time `json:"dateOfBirth" gorm:"column:date_of_birth" bson:"dateOfBirth" dynamodbav:"dateOfBirth" firestore:"dateOfBirth" avro:"dateOfBirth"`
}

// func NewUser(id string, username string, email string, phone string, dateOfBirth *time.Time) (*User, error) {
// 	u := &User{
// 		Id:          id,
// 		Username:    username,
// 		Email:       email,
// 		Phone:       phone,
// 		DateOfBirth: dateOfBirth,
// 	}
// 	err := u.Validate()
// 	if err != nil {
// 		return nil, errors.New("error invalid entity")
// 	}
// 	return u, nil
// }

// func (u *User) Validate() error {
// 	if u.Id == "" || u.Username == "" || u.Email == "" || u.Phone == "" {
// 		return errors.New("error invalid entity")
// 	}
// 	return nil
// }
