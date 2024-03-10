package helper

import "time"

func TimeStrCurrentDayStart(st time.Time) string {
	ct := st.Format("2006-01-02")
	return ct + " 00:00:00"
}

func TimeStrCurrentDayEnd(et time.Time) string {
	ct := et.Format("2006-01-02")
	return ct + " 23:59:59"
}

func TimeCurrentDayStart(st time.Time) time.Time {
	ct := st.Format("2006-01-02")
	t, _ := time.Parse("2006-01-02 15:04:05", ct+" 00:00:00")
	return t
}

func TimeCurrentDayEnd(et time.Time) time.Time {
	ct := et.Format("2006-01-02")
	t, _ := time.Parse("2006-01-02 15:04:05", ct+" 23:59:59")
	return t
}
