package roles

import (
	"gorm.io/gorm"
)

type Coordinator struct {
	gorm.Model

	// It's original user
	fkUser User `gorm:"foreignkey:UserID"`

	// Courses coordinated by this coordinator
	CoursesCoordinated []Course `gorm:"foreignkey:CourseID"`
}
