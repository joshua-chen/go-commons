package time

import "time"

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	t := time.Time(j)
	if t.IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + t.Format("2006-01-02 15:04:05") + `"`), nil
}
