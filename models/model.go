package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Student Registration
type StudentForm struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone     string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Course    *Course            `json:"course,omitempty" bson:"course,omitempty"`
}

type Course struct {
	CourseID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CourseName string             `json:"coursename,omitempty" bson:"coursename,omitempty"`
}
