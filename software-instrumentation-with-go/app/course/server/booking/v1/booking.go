package v1

import (
	"app/course/booking"
	"app/course/catalog"
	"app/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	ctx := c.Request.Context()
	log := logger.FromCtx(ctx)

	log.Info("create booking request received")

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Error("invalid request body",
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to validate request",
			"error":   err.Error(),
		})
		return
	}

	log.Info("create booking",
		zap.String("batch_code", req.BatchCode),
		zap.String("customer_name", req.CustomerName),
	)

	b, err := h.service.CreateBooking(ctx, req.BatchCode, req.CustomerName)
	if err != nil {
		switch err {

		case booking.ErrInvalidCustomerName:
			log.Warn("invalid customer name",
				zap.String("customer_name", req.CustomerName),
			)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to validate customer name",
				"error":   err.Error(),
			})
			return

		case catalog.ErrBatchNotFound:
			log.Warn("batch not found",
				zap.String("batch_code", req.BatchCode),
			)
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Failed to find batch course",
				"error":   err.Error(),
			})
			return

		case catalog.ErrBatchFull:
			log.Warn("batch full",
				zap.String("batch_code", req.BatchCode),
			)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to booking batch full",
				"error":   err.Error(),
			})
			return

		case catalog.ErrBatchNotOpen:
			log.Warn("batch not open",
				zap.String("batch_code", req.BatchCode),
			)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to booking batch not open",
				"error":   err.Error(),
			})
			return

		default:
			log.Error("booking failed",
				zap.String("batch_code", req.BatchCode),
				zap.String("customer_name", req.CustomerName),
				zap.Error(err),
			)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to booking batch course",
				"error":   err.Error(),
			})
			return
		}
	}

	log.Info("booking created",
		zap.String("booking_id", b.ID),
		zap.String("booking_code", b.Code),
		zap.String("batch_code", b.Batch.Code),
		zap.String("customer_name", b.CustomerName),
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Booking batch course created",
		"data": gin.H{
			"id":            b.ID,
			"code":          b.Code,
			"status":        b.Status,
			"customer_name": b.CustomerName,
			"batch_code":    b.Batch.Code,
		},
	})
}
