package presence

import (
	"testing"
	"time"
)

func TestDecide(t *testing.T) {
	now := time.Date(2026, 11, 3, 10, 0, 0, 0, time.UTC)
	fresh := now.Add(-3 * time.Minute)
	stale := now.Add(-11 * time.Minute)
	const roomA, roomB = 1, 2

	tests := []struct {
		name string
		cur  Current
		obs  Observation
		want Action
	}{
		{"同じ部屋で観測(強くなる)", Current{roomA, -70, fresh}, Observation{roomA, true, -65}, Update},
		{"同じ部屋で観測(弱くなる)も更新", Current{roomA, -70, fresh}, Observation{roomA, true, -80}, Update},
		{"同じ部屋で1回取りこぼしても退室しない", Current{roomA, -70, fresh}, Observation{roomA, false, 0}, Keep},
		{"同じ部屋で猶予を超えて見えなければ退室", Current{roomA, -70, stale}, Observation{roomA, false, 0}, Exit},
		{"隣の部屋が少し強いだけでは移動しない", Current{roomA, -70, fresh}, Observation{roomB, true, -67}, Keep},
		{"隣の部屋が弱ければ移動しない", Current{roomA, -70, fresh}, Observation{roomB, true, -85}, Keep},
		{"隣の部屋が明確に強ければ移動", Current{roomA, -80, fresh}, Observation{roomB, true, -60}, Exit},
		{"今の部屋で見えなくなっていれば弱くても移動", Current{roomA, -70, stale}, Observation{roomB, true, -90}, Exit},
		{"別部屋の報告に含まれなければ何もしない", Current{roomA, -70, stale}, Observation{roomB, false, 0}, Keep},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decide(tt.cur, tt.obs, now); got != tt.want {
				t.Errorf("Decide() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 隣接2部屋で両方の受信機がユーザを拾い、RSSIが揺らぐ状況を再現する。
// ユーザは部屋Aに座ったまま。旧ロジック相当では行き来が起きるが、
// 新ロジックでは部屋Aに留まることを確認する。
func TestNoFlappingBetweenAdjacentRooms(t *testing.T) {
	const roomA, roomB = 1, 2
	start := time.Date(2026, 11, 3, 10, 0, 0, 0, time.UTC)

	// 5分おきの報告。Aは -68〜-75 で揺れ、たまに取りこぼす(Seen=false)。
	// Bは壁越しで -72〜-78 だが、時々Aを上回る。
	type report struct {
		obs Observation
		dt  time.Duration
	}
	reports := []report{
		{Observation{roomA, true, -70}, 0},
		{Observation{roomB, true, -73}, 2 * time.Minute},
		{Observation{roomA, true, -75}, 3 * time.Minute},
		{Observation{roomB, true, -72}, 2 * time.Minute}, // Aより強い
		{Observation{roomA, false, 0}, 3 * time.Minute},  // Aで取りこぼし
		{Observation{roomB, true, -74}, 2 * time.Minute},
		{Observation{roomA, true, -68}, 3 * time.Minute},
		{Observation{roomB, true, -71}, 2 * time.Minute},
	}

	cur := Current{RoomID: roomA, Rssi: -70, LastSeen: start}
	now := start
	for i, r := range reports[1:] {
		now = now.Add(r.dt)
		switch Decide(cur, r.obs, now) {
		case Update:
			cur.Rssi, cur.LastSeen = r.obs.Rssi, now
		case Exit:
			t.Fatalf("report %d: 部屋%dの報告で移動/退室と判定された", i+1, r.obs.RoomID)
		}
	}
	if cur.RoomID != roomA {
		t.Fatalf("room = %d, want %d", cur.RoomID, roomA)
	}
}
