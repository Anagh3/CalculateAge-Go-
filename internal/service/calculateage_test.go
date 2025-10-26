package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "exact birthday today",
			dob:      time.Date(now.Year()-25, now.Month(), now.Day(), 0, 0, 0, 0, time.Local),
			expected: 25,
		},
		{
			name:     "birthday yet to come this year",
			dob:      time.Date(now.Year()-25, now.Month()+1, now.Day(), 0, 0, 0, 0, time.Local),
			expected: 24,
		},
		{
			name:     "birthday already passed this year",
			dob:      time.Date(now.Year()-25, now.Month()-1, now.Day(), 0, 0, 0, 0, time.Local),
			expected: 25,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := calculateAge(tc.dob)
			if got != tc.expected {
				t.Errorf("%s: expected %d, got %d", tc.name, tc.expected, got)
			}
		})
	}
}
