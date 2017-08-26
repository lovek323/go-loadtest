package loadtest

import "time"

type Summary struct {
	Start time.Time
	End   time.Time
}

func LoadTest(function func() error, max int) (*Summary, error) {
	start := now()
	count := 0
	for {
		if err := function(); err != nil {
			return nil, err
		}
		if count++; count > max {
			break
		}
	}
	end := now()
	summary := &Summary{
		Start: start,
		End:   end,
	}
	return summary, nil
}
