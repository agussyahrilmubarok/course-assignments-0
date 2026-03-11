package v1

import (
	"app/course/catalog"
	"app/internal/database/entity"
	"app/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	service *catalog.Service
}

func NewHandler(service *catalog.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(r *gin.RouterGroup) {
	router := r.Group("/catalog")
	router.GET("/seed", h.SeedData)
	router.GET("/clean", h.CleanData)
}

func (h *Handler) SeedData(c *gin.Context) {
	ctx := c.Request.Context()

	now := util.GetJakartaTimeNow()

	courses := []entity.Course{
		{
			ID:     "COURSE-001",
			Code:   "COURSE-GIT",
			Title:  "Version Control with Git and GitHub",
			Status: int(catalog.CourseStatusPublished),
		},
		{
			ID:     "COURSE-002",
			Code:   "COURSE-API",
			Title:  "Designing RESTful APIs for Modern Applications",
			Status: int(catalog.CourseStatusPublished),
		},
		{
			ID:     "COURSE-003",
			Code:   "COURSE-DB",
			Title:  "Database Design and SQL for Developers",
			Status: int(catalog.CourseStatusPublished),
		},
		{
			ID:     "COURSE-004",
			Code:   "COURSE-ARCH",
			Title:  "Software Architecture and Clean Code Principles",
			Status: int(catalog.CourseStatusPublished),
		},
		{
			ID:     "COURSE-005",
			Code:   "COURSE-DEVOPS",
			Title:  "Introduction to DevOps and CI/CD Pipelines",
			Status: int(catalog.CourseStatusPublished),
		},
	}

	for i := range courses {
		course := courses[i]

		course.CreatedAt = now
		course.UpdatedAt = now

		err := h.service.Store.DB.
			WithContext(ctx).
			Where("code = ?", course.Code).
			FirstOrCreate(&course).
			Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to seed courses",
			})
			return
		}
	}

	startOfDay := util.GetJakartaStartOfDay(now)

	batches := []entity.Batch{
		{
			ID:             "BATCH-001",
			CourseID:       "COURSE-001",
			Code:           "GIT-BATCH-01",
			Name:           "Git Fundamentals Batch 1",
			PriceAmount:    500000,
			PriceCurrency:  "IDR",
			MaxSeats:       30,
			AvailableSeats: 30,
			Status:         1,
			StartDate:      startOfDay.AddDate(0, 0, 1).Unix(),
			EndDate:        startOfDay.AddDate(0, 0, 7).Unix(),
			CreatedAt:      now,
			UpdatedAt:      now,
		},
		{
			ID:             "BATCH-002",
			CourseID:       "COURSE-002",
			Code:           "API-BATCH-01",
			Name:           "REST API Development Batch 1",
			PriceAmount:    750000,
			PriceCurrency:  "IDR",
			MaxSeats:       25,
			AvailableSeats: 25,
			Status:         1,
			StartDate:      startOfDay.AddDate(0, 0, 1).Unix(),
			EndDate:        startOfDay.AddDate(0, 0, 7).Unix(),
			CreatedAt:      now,
			UpdatedAt:      now,
		},
		{
			ID:             "BATCH-003",
			CourseID:       "COURSE-004",
			Code:           "ARCH-BATCH-01",
			Name:           "Software Architecture Bootcamp",
			PriceAmount:    900000,
			PriceCurrency:  "IDR",
			MaxSeats:       20,
			AvailableSeats: 20,
			Status:         1,
			StartDate:      startOfDay.AddDate(0, 0, 1).Unix(),
			EndDate:        startOfDay.AddDate(0, 0, 7).Unix(),
			CreatedAt:      now,
			UpdatedAt:      now,
		},
	}

	for i := range batches {
		batch := batches[i]

		batch.CreatedAt = now
		batch.UpdatedAt = now

		err := h.service.Store.DB.
			WithContext(ctx).
			Where("code = ?", batch.Code).
			FirstOrCreate(&batch).
			Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to seed batches",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Seed courses and batches success",
	})
}

func (h *Handler) CleanData(c *gin.Context) {
	ctx := c.Request.Context()

	db := h.service.Store.DB.WithContext(ctx)

	if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&entity.Batch{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed delete batches",
			"error":   err.Error(),
		})
		return
	}

	if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&entity.Course{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed delete courses",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "clean success",
	})
}
