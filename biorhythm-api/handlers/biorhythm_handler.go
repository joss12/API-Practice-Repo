package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/biorhythm-api/models"
	"github.com/biorhythm-api/services"
	"github.com/gin-gonic/gin"
)

type BiorhythmHandler struct {
	service *services.BiorhythmService
}

func NewBiorhythmHandler() *BiorhythmHandler {
	return &BiorhythmHandler{
		service: services.NewBiorhythmService(),
	}
}

// GetBiorhythm handles GET /biorhythm/:birthdate
func (bh *BiorhythmHandler) GetBiorhythm(c *gin.Context) error {
	birthDate := strings.TrimSpace(c.Param("birthdate"))

	if birthDate == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Birth date is required",
			Error:   "Please provide birth date in format YYYY-MM-DD (e.g., 1990-05-15)",
		})
		return nil
	}

	result, err := bh.service.GetBiorhythmForToday(birthDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Failed to calculate biorhythm",
			Error:   err.Error(),
		})
		return nil
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Biorhythm calculated successfully",
		Data:    result,
	})
	return nil
}

// GetBiorhythmForDate handles GET /biorhythm/:birthdate/:date
func (bh *BiorhythmHandler) GetBiorhythmForDate(c *gin.Context) error {
	birthDate := strings.TrimSpace(c.Param("birthdate"))
	targetDate := strings.TrimSpace(c.Param("date"))

	if birthDate == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Birth date is required",
			Error:   "Please provide birth date in format YYYY-MM-DD (e.g., 1990-05-15)",
		})
		return nil
	}

	if targetDate == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Target date is required",
			Error:   "Please provide target date in format YYYY-MM-DD (e.g., 2024-12-25)",
		})
		return nil
	}

	result, err := bh.service.GetBiorhythmForDate(birthDate, targetDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Failed to calculate biorhythm",
			Error:   err.Error(),
		})
		return nil
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Biorhythm calculated successfully",
		Data:    result,
	})
	return nil
}

// GetCriticalDays handles GET /critical-days/:birthdate/:days
func (bh *BiorhythmHandler) GetCriticalDays(c *gin.Context) error {
	birthDate := strings.TrimSpace(c.Param("birthdate"))
	daysStr := strings.TrimSpace(c.Param("days"))

	if birthDate == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Birth date is required",
			Error:   "Please provide birth date in format YYYY-MM-DD (e.g., 1990-05-15)",
		})
		return nil
	}

	days, err := strconv.Atoi(daysStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Invalid days parameter",
			Error:   "Days must be a valid number between 1 and 365",
		})
		return nil
	}

	criticalDays, err := bh.service.GetCriticalDays(birthDate, days)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Failed to calculate critical days",
			Error:   err.Error(),
		})
		return nil
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Critical days calculated successfully",
		Data:    criticalDays,
	})
	return nil
}

// GetCompatibility handles GET /compatibility/:birthdate1/:birthdate2
func (bh *BiorhythmHandler) GetCompatibility(c *gin.Context) error {
	birthDate1 := strings.TrimSpace(c.Param("birthdate1"))
	birthDate2 := strings.TrimSpace(c.Param("birthdate2"))
	targetDate := c.Query("date") // Optional query parameter

	if birthDate1 == "" || birthDate2 == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Both birth dates are required",
			Error:   "Please provide both birth dates in format YYYY-MM-DD (e.g., 1990-05-15)",
		})
		return nil
	}

	if targetDate == "" {
		targetDate = time.Now().Format("2006-01-02")
	}

	result, err := bh.service.GetCompatibility(birthDate1, birthDate2, targetDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Failed to calculate compatibility",
			Error:   err.Error(),
		})
		return nil
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Compatibility calculated successfully",
		Data:    result,
	})
	return nil
}

// GetBiorhythmChart handles GET /chart/:birthdate/:start/:end
func (bh *BiorhythmHandler) GetBiorhythmChart(c *gin.Context) error {
	birthDate := strings.TrimSpace(c.Param("birthdate"))
	startDate := strings.TrimSpace(c.Param("start"))
	endDate := strings.TrimSpace(c.Param("end"))

	if birthDate == "" || startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "All dates are required",
			Error:   "Please provide birth date, start date, and end date in format YYYY-MM-DD",
		})
		return nil
	}

	chart, err := bh.service.GetBiorhythmChart(birthDate, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Message: "Failed to generate biorhythm chart",
			Error:   err.Error(),
		})
		return nil
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Biorhythm chart generated successfully",
		Data:    chart,
	})
	return nil
}
