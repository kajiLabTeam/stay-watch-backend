package stat

import "testing"

func TestTimeToMinutes(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"00:00", 0},
		{"00:10", 10},
		{"00:15", 15},
		{"00:29", 29},
		{"01:15", 75},
		{"09:05", 545},
		{"23:59", 1439},
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
	for _, in := range []string{"invalid", "24:00", "-1:00", "12:60", "12:-1"} {
		if _, err := TimeToMinutes(in); err == nil {
			t.Errorf("TimeToMinutes(%q): expected error", in)
		}
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
		{1440, "00:00"},
		{1470, "00:30"},
		{89.4, "01:29"},
		{89.5, "01:30"},
		{89.6, "01:30"},
		{1439.6, "00:00"},
	}
	for _, c := range cases {
		got := MinutesToTime(c.in)
		if got != c.want {
			t.Errorf("MinutesToTime(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}
