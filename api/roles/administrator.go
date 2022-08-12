package roles

import (
	"gorm.io/gorm"
)

type Administrator struct {
	gorm.Model

	// It's original user
	fkUser User `gorm:"foreignkey:UserID"`

	// Courses administrated by this administrator
	CoursesAdministrated []Course `gorm:"foreignkey:CourseID"`
}
