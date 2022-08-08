package api

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// The name of the user, can be compound, or short, Li, for example, also will and should be valid
	Name string
	// The surname of the user, if available
	Surname string
	// The last name(s) of the user
	Lastname string
	// Date of birth, not to be address publicly unless explicitly said so
	Birthdate string
	// The privacy of the birthdate
	IsBirthdatePublic bool
	// The main e-mail the user prefers as a point of contact
	Email string
	// How would the user prefer to be addressed as
	Pronouns string
	// A more advanced settings should be made as too choose which group to show it to,
	// and block certain users from seeing it, or rather,
	// even allowing only a selected group of users that can visualize it
	arePronounsPublic bool
	// A brief description the user chooses for themself
	Bio string
	// Should take into account multiple inscriptions,
	// wether by bans, or multiple courses throughout the years
	InscriptionDate string
	// The user that authorized this specific user, or enrolled them
	AuthorizedBy *uint
	// All of the users that this specific user authorized
	Authorized []User `gorm:"foreignkey:UserID"`
}
