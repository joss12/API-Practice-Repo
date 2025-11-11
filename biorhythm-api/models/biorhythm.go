package models

// BiorhythmCycle represents the three main biorhythm cycles
type BiorhythmCycle struct {
	Physical     float64 `json:"physical"`
	Emotional    float64 `json:"emotional"`
	Intellectual float64 `json:"intellectual"`
}

// BiorhythmResult contains biorhythm data for a specific date
type BiorhythmResult struct {
	Date           string         `json:"date"`
	DaysSinceBirth int            `json:"days_since_birth"`
	Cycles         BiorhythmCycle `json:"cycles"`
	OverallScore   float64        `json:"overall_score"`
	Phase          string         `json:"phase"`
}

// CriticalDay represents a day when biorhythm cycles cross zero
type CriticalDay struct {
	Date  string `json:"date"`
	Cycle string `json:"cycle"`
	Type  string `json:"type"`
}

// CompatibilityResult shows biorhythm compatibility between two people
type CompatibilityResult struct {
	Date               string  `json:"date"`
	PhysicalMatch      float64 `json:"physical_match"`
	EmotionalMatch     float64 `json:"emotional_match"`
	IntellectualMatch  float64 `json:"intellectual_match"`
	OverallMatch       float64 `json:"overall_match"`
	CompatibilityLevel string  `json:"compatibility_level"`
}

// BiorhythmChart represents data for chart visualization
type BiorhythmChart struct {
	StartDate string            `json:"start_date"`
	EndDate   string            `json:"end_data"`
	Data      []BiorhythmResult `json:"data"`
}
