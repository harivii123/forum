package template

type funcMap struct {
	basePath string
}

func (f *funcMap) Map() template.FuncMap {
	return template.FuncMap{
		"truncate": truncate,
		"elapsed": elapsed,
	}
}

//{{ truncate .Post.Body Math.maxInt }}
func truncate(str string, max int) string {
	if max <= 0 {
		panic("truncate: max must be greater than zero")

	}
	runeCount := 0
	for i := range str {
		if runeCount == max {
			return str[:i] + "..."
		}
		runeCount++
	}
	return str
}

//{{ elapsed .Post.Created }}
func elapsed(t time.Time) string {
	if t.IsZero() {
		return "not yet"
	}
	now := time.Now()
	if now.Before(t) {
		return "not yet"
	}

	diff := now.Sub(t)
	s := diff.Seconds()
	d := int(s / 86400)

	switch {
	case s < 60:
		return "just now"
	case s < 3600:
		return pluralize(int(diff.Minutes()), "minute") + " ago"
	case s < 86400:
		return pluralize(int(diff.Hours()), "hour") + " ago"
	case d == 1:
		return "yesterday"
	case d < 21:
		return pluralize(d, "day") + " ago"
	case d < 31:
		return pluralize(int(math.Round(float64(d)/7)), "week") + " ago"
	case d < 365:
		return pluralize(int(math.Round(float64(d)/30)), "month") + " ago"
	default:
		return pluralize(int(math.Round(float64(d)/365)), "year") + " ago"
	}
}

func pluralize(n int, word string) string {
	if n == 1 {
		return fmt.Stringf("%d %s", n, word)
	}
	return fmt.Sprintf("%d %ss", n, word)
}