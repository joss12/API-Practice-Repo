package services

import (
	"errors"
	"time"

	"github.com/biorhythm-api/models"
	"github.com/biorhythm-api/utils"
)

type BiorhythmService struct {
	calculator *CalculationService
}

func NewBiorhythmService() *BiorhythmService {
	return &BiorhythmService{
		calculator: NewCalculationService(),
	}
}

// GetBiorhythmForToday returns biorhythm for current date
func (bs *BiorhythmService) GetBiorhythmForToday(birthDateStr string) (*models.BiorhythmResult, error) {
	return bs.GetBiorhythmForDate(birthDateStr, time.Now().Format("2006-01-02"))
}

// GetBiorhythmForDate returns biorhythm for specific date
func (bs *BiorhythmService) GetBiorhythmForDate(birthDateStr, targetDateStr string) (*models.BiorhythmResult, error) {
	birthDate, err := utils.ParseDate(birthDateStr)
	if err != nil {
		return nil, errors.New("invalid birth date format")
	}

	targetDate, err := utils.ParseDate(targetDateStr)
	if err != nil {
		return nil, errors.New("invalid target date format")
	}

	if targetDate.Before(birthDate) {
		return nil, errors.New("target date cannot be before birth date")
	}

	cycles := bs.calculator.CalculateBiorhythm(birthDate, targetDate)
	overallScore := bs.calculator.CalculateOverallScore(cycles)
	phase := bs.calculator.DeterminePhase(overallScore)
	daysSinceBirth := int(targetDate.Sub(birthDate).Hours() / 24)

	return &models.BiorhythmResult{
		Date:           targetDate.Format("2006-01-02"),
		DaysSinceBirth: daysSinceBirth,
		Cycles:         cycles,
		OverallScore:   overallScore,
		Phase:          phase,
	}, nil
}

// GetCriticalDays returns upcoming critical days
func (bs *BiorhythmService) GetCriticalDays(birthDateStr string, days int) ([]models.CriticalDay, error) {
	if days <= 0 || days > 365 {
		return nil, errors.New("days must be between 1 and 365")
	}

	birthDate, err := utils.ParseDate(birthDateStr)
	if err != nil {
		return nil, errors.New("invalid birth date format")
	}

	startDate := time.Now()
	criticalDays := bs.calculator.FindCriticalDays(birthDate, startDate, days)

	return criticalDays, nil
}

// GetCompatibility calculates compatibility between two birth dates
func (bs *BiorhythmService) GetCompatibility(birthDate1Str, birthDate2Str, targetDateStr string) (*models.CompatibilityResult, error) {
	birthDate1, err := utils.ParseDate(birthDate1Str)
	if err != nil {
		return nil, errors.New("invalid first birth date format")
	}

	birthDate2, err := utils.ParseDate(birthDate2Str)
	if err != nil {
		return nil, errors.New("invalid second birth date format")
	}

	var targetDate time.Time
	if targetDateStr == "" {
		targetDate = time.Now()
	} else {
		targetDate, err = utils.ParseDate(targetDateStr)
		if err != nil {
			return nil, errors.New("invalid target date format")
		}
	}

	compatibility := bs.calculator.CalculateCompatibility(birthDate1, birthDate2, targetDate)
	return &compatibility, nil
}

// GetBiorhythmChart returns biorhythm data for date range
func (bs *BiorhythmService) GetBiorhythmChart(birthDateStr, startDateStr, endDateStr string) (*models.BiorhythmChart, error) {
	// Validate birth date format (we only need validation, not the parsed value)
	_, err := utils.ParseDate(birthDateStr)
	if err != nil {
		return nil, errors.New("invalid birth date format")
	}

	startDate, err := utils.ParseDate(startDateStr)
	if err != nil {
		return nil, errors.New("invalid start date format")
	}

	endDate, err := utils.ParseDate(endDateStr)
	if err != nil {
		return nil, errors.New("invalid end date format")
	}

	if endDate.Before(startDate) {
		return nil, errors.New("end date cannot be before start date")
	}

	// Limit to maximum 90 days to prevent large responses
	if endDate.Sub(startDate).Hours()/24 > 90 {
		return nil, errors.New("date range cannot exceed 90 days")
	}

	var data []models.BiorhythmResult
	for date := startDate; !date.After(endDate); date = date.AddDate(0, 0, 1) {
		result, err := bs.GetBiorhythmForDate(birthDateStr, date.Format("2006-01-02"))
		if err != nil {
			continue // Skip invalid dates
		}
		data = append(data, *result)
	}

	return &models.BiorhythmChart{
		StartDate: startDate.Format("2006-01-02"),
		EndDate:   endDate.Format("2006-01-02"),
		Data:      data,
	}, nil
}
