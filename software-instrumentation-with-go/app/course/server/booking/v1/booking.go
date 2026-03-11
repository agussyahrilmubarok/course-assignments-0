package v1

import (
	"app/course/booking"
	"app/course/catalog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *booking.Service
}

func NewHandler(service *booking.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(r *gin.RouterGroup) {
	router := r.Group("/booking")
	router.POST("", h.CreateBooking)
}

func (h *Handler) CreateBooking(c *gin.Context) {
	var req CreateBookingRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to validate request",
			"error":   err.Error(),
		})
		return
	}

	b, err := h.service.CreateBooking(c.Request.Context(), req.BatchCode, req.CustomerName)
	if err != nil {
		switch err {

		case booking.ErrInvalidCustomerName:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to validate customer name",
				"error":   err.Error(),
			})
			return

		case catalog.ErrBatchNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Failed to find batch course",
				"error":   err.Error(),
			})
			return

		case catalog.ErrBatchFull:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to booking batch full",
				"error":   err.Error(),
			})
			return

		case catalog.ErrBatchNotOpen:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to booking batch not open",
				"error":   err.Error(),
			})
			return

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to booking batch course",
				"error":   err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Booking batch course created",
		"data": gin.H{
			"id":            b.ID,
			"code":          b.Code,
			"status":        b.Status,
			"customer_name": b.CustomerName,
			"batch_code":    b.Batch.Code,
			"expired_at":    b.ExpiredAt,
		},
	})
}
