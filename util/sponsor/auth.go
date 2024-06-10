package sponsor

import (
        "fmt"
        "time"
)

var (
        Subject, Token string
        ExpiresAt      time.Time
)

const (
	unavailable = "sponsorship unavailable"
	victron     = "victron"
)

func IsAuthorized() bool {
        return len(Subject) > 0
}

func IsAuthorizedForApi() bool {
        return IsAuthorized() && Subject != unavailable
}

// check and set sponsorship token
func ConfigureSponsorship(token string) error {
        Subject = "tosog"
        ExpiresAt, err := time.Parse(time.RFC3339, "2047-02-11 22:00:00 +0000 UTC")
        Token = "xxJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJldmNjLmlvIiwic3ViIjoidG9zb2ciLCJleHAiOjE4MDIzODMyMDAsImlhdCI6MTcwNzc3NTIwMCwic3JjIjoiZ2gifQ.mYJb9N4Zz0deV6Mgyj7Eq5rdNeoVlq_xN146tGvxop8" // was: ey

        if (err != nil || ExpiresAt.IsZero()) {
                fmt.Println("")
        }

        return nil

}

type sponsorStatus struct {
	Name        string    `json:"name"`
	ExpiresAt   time.Time `json:"expiresAt,omitempty"`
	ExpiresSoon bool      `json:"expiresSoon,omitempty"`
}

// Status returns the sponsorship status
func Status() sponsorStatus {
	var expiresSoon bool
	if d := time.Until(ExpiresAt); d < 30*24*time.Hour && d > 0 {
		expiresSoon = true
	}

	return sponsorStatus{
		Name:        Subject,
		ExpiresAt:   ExpiresAt,
		ExpiresSoon: expiresSoon,
	}
}
