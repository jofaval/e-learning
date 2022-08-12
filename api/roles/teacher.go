package roles

import (
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model

	// It's original user
	fkUser User `gorm:"foreignkey:UserID"`

	// Courses teached by this teacher
	CoursesTeached []Course `gorm:"foreignkey:CourseID"`
}
