package services

import (
	"math"
	"time"

	"github.com/biorhythm-api/models"
	"github.com/biorhythm-api/utils"
)

type CalculationService struct{}

func NewCalculationService() *CalculationService {
	return &CalculationService{}
}

// CalculateBiorhythm calculates biorhythm cycles for a given date
func (cs *CalculationService) CalculateBiorhythm(birthDate, targetDate time.Time) models.BiorhythmCycle {
	daysSinceBirth := int(targetDate.Sub(birthDate).Hours() / 24)

	// Calculate each cycle using sine wave formula
	physical := math.Sin(2*math.Pi*float64(daysSinceBirth)/23.0) * 100
	emotional := math.Sin(2*math.Pi*float64(daysSinceBirth)/28.0) * 100
	intellectual := math.Sin(2*math.Pi*float64(daysSinceBirth)/33.0) * 100

	return models.BiorhythmCycle{
		Physical:     utils.RoundToDecimal(physical, 2),
		Emotional:    utils.RoundToDecimal(emotional, 2),
		Intellectual: utils.RoundToDecimal(intellectual, 2),
	}
}

// CalculateOverallScore computes weighted average of all cycles
func (cs *CalculationService) CalculateOverallScore(cycles models.BiorhythmCycle) float64 {
	// Equal weighting for all cycles
	score := (cycles.Physical + cycles.Emotional + cycles.Intellectual) / 3.0
	return utils.RoundToDecimal(score, 2)
}

// DeterminePhase determines if biorhythm is in high, low, or critical phase
func (cs *CalculationService) DeterminePhase(score float64) string {
	absScore := math.Abs(score)

	if absScore <= 10 {
		return "critical"
	} else if score > 10 {
		return "high"
	}
	return "low"
}

// FindCriticalDays identifies days when cycles cross zero
func (cs *CalculationService) FindCriticalDays(birthDate time.Time, startDate time.Time, days int) []models.CriticalDay {
	var criticalDays []models.CriticalDay

	for i := 0; i < days; i++ {
		currentDate := startDate.AddDate(0, 0, i)
		daysSinceBirth := int(currentDate.Sub(birthDate).Hours() / 24)

		// Check each cycle for zero crossings
		cycles := []struct {
			name   string
			period int
		}{
			{"physical", 23},
			{"emotional", 28},
			{"intellectual", 33},
		}

		for _, cycle := range cycles {
			value := math.Sin(2 * math.Pi * float64(daysSinceBirth) / float64(cycle.period))
			prevValue := math.Sin(2 * math.Pi * float64(daysSinceBirth-1) / float64(cycle.period))

			// Check for zero crossing
			if (value > 0 && prevValue <= 0) || (value < 0 && prevValue >= 0) {
				criticalType := "rising"
				if value < prevValue {
					criticalType = "falling"
				}

				criticalDays = append(criticalDays, models.CriticalDay{
					Date:  currentDate.Format("2006-01-02"),
					Cycle: cycle.name,
					Type:  criticalType,
				})
			}
		}
	}

	return criticalDays
}

// CalculateCompatibility computes biorhythm compatibility between two people
func (cs *CalculationService) CalculateCompatibility(birth1, birth2, targetDate time.Time) models.CompatibilityResult {
	cycles1 := cs.CalculateBiorhythm(birth1, targetDate)
	cycles2 := cs.CalculateBiorhythm(birth2, targetDate)

	// Calculate compatibility as inverse of difference (closer = more compatible)
	physicalMatch := 100 - math.Abs(cycles1.Physical-cycles2.Physical)
	emotionalMatch := 100 - math.Abs(cycles1.Emotional-cycles2.Emotional)
	intellectualMatch := 100 - math.Abs(cycles1.Intellectual-cycles2.Intellectual)

	overallMatch := (physicalMatch + emotionalMatch + intellectualMatch) / 3.0

	var level string
	switch {
	case overallMatch >= 80:
		level = "excellent"
	case overallMatch >= 60:
		level = "good"
	case overallMatch >= 40:
		level = "moderate"
	default:
		level = "challenging"
	}

	return models.CompatibilityResult{
		Date:               targetDate.Format("2006-01-02"),
		PhysicalMatch:      utils.RoundToDecimal(physicalMatch, 2),
		EmotionalMatch:     utils.RoundToDecimal(emotionalMatch, 2),
		IntellectualMatch:  utils.RoundToDecimal(intellectualMatch, 2),
		OverallMatch:       utils.RoundToDecimal(overallMatch, 2),
		CompatibilityLevel: level,
	}
}
