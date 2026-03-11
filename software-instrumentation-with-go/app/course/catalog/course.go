package catalog

import (
	"app/internal/database/entity"
	"time"
)

type Course struct {
	ID        string
	Code      string
	Title     string
	Status    CourseStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CourseStatus int

const (
	CourseStatusDraft CourseStatus = iota
	CourseStatusPublished
	CourseStatusArchived
)

func (c Course) IsPublished() bool {
	return c.Status == CourseStatusPublished
}

func (c *Course) Publish() error {
	if c.Status != CourseStatusDraft {
		return ErrInvalidCourseStatus
	}

	c.Status = CourseStatusPublished
	c.UpdatedAt = time.Now()

	return nil
}

func (c *Course) Archive() error {
	if c.Status == CourseStatusArchived {
		return ErrInvalidCourseStatus
	}

	c.Status = CourseStatusArchived
	c.UpdatedAt = time.Now()

	return nil
}

func ToCourse(e *entity.Course) *Course {
	if e == nil {
		return nil
	}

	return &Course{
		ID:        e.ID,
		Code:      e.Code,
		Title:     e.Title,
		Status:    CourseStatus(e.Status),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func ToCourseEntity(c *Course) *entity.Course {
	if c == nil {
		return nil
	}

	return &entity.Course{
		ID:        c.ID,
		Code:      c.Code,
		Title:     c.Title,
		Status:    int(c.Status),
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
