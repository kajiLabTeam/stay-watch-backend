package stat

import "testing"

func TestTimeToMinutes(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"00:00", 0},
		{"00:10", 0},
		{"00:14", 0},
		{"00:15", 30}, // math.Round rounds half away from zero (diverges from Python's banker's rounding, which would give 0)
		{"00:16", 30},
		{"00:29", 30},
		{"00:44", 30},
		{"00:45", 60},
		{"01:15", 90},
		{"23:59", 24 * 60},
	}
	for _, c := range cases {
		got, err := TimeToMinutes(c.in)
		if err != nil {
			t.Fatalf("TimeToMinutes(%q) returned error: %v", c.in, err)
		}
		if got != c.want {
			t.Errorf("TimeToMinutes(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}

func TestTimeToMinutesInvalid(t *testing.T) {
	if _, err := TimeToMinutes("invalid"); err == nil {
		t.Error("expected error for invalid input")
	}
}

func TestMinutesToTime(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{0, "00:00"},
		{30, "00:30"},
		{90, "01:30"},
		{1439, "23:59"},
	}
	for _, c := range cases {
		got := MinutesToTime(c.in)
		if got != c.want {
			t.Errorf("MinutesToTime(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}
