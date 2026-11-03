// Package presence は受信機からの電波報告をもとに、在室者を
// 「どの部屋に置くか」を決める純粋な判定ロジックをまとめたもの。
//
// 隣接する部屋に受信機を置くと、両方の受信機が同じビーコンを拾う。
// 単純に「RSSIが強い方」「報告に含まれていない＝退室」で判定すると、
// RSSIの揺らぎやスキャンの取りこぼしで2部屋を行き来する判定（フラッピング）
// が起きるため、ここではヒステリシスと猶予時間を入れて判定する。
package presence

import (
	"os"
	"strconv"
	"time"
)

const (
	// DefaultRoomSwitchRssiMargin 別の部屋へ移ったと判定するのに必要なRSSI差[dB]。
	// 別部屋のRSSIが「今いる部屋の直近RSSI + この値」以上のときだけ移動とみなす。
	DefaultRoomSwitchRssiMargin = 5

	// DefaultAbsenceGrace 今いる部屋の受信機から見えなくなってから
	// 退室（または他部屋への移動）を認めるまでの猶予時間。
	// 受信機のPOST周期(5分)の2回分を目安にしている。
	DefaultAbsenceGrace = 10 * time.Minute
)

// RoomSwitchRssiMargin 判定に使うRSSI差[dB]。環境変数 PRESENCE_SWITCH_MARGIN_DB で上書きできる。
// .env はserviceパッケージのinitで読み込まれるため、パッケージinitではなく判定時に読む。
func RoomSwitchRssiMargin() int64 {
	if v, err := strconv.ParseInt(os.Getenv("PRESENCE_SWITCH_MARGIN_DB"), 10, 64); err == nil {
		return v
	}
	return DefaultRoomSwitchRssiMargin
}

// AbsenceGrace 判定に使う退室猶予。環境変数 PRESENCE_ABSENCE_GRACE_SEC[秒] で上書きできる。
// ローカル検証で猶予を短くする用途を想定。
func AbsenceGrace() time.Duration {
	if v, err := strconv.Atoi(os.Getenv("PRESENCE_ABSENCE_GRACE_SEC")); err == nil && v > 0 {
		return time.Duration(v) * time.Second
	}
	return DefaultAbsenceGrace
}

// Action はStayerテーブルの既存レコードに対して取る処理。
type Action int

const (
	// Keep 何もしない。
	Keep Action = iota
	// Update 同じ部屋のままRSSIと最終観測時刻を更新する。
	Update
	// Exit 現在の部屋から退室させる（他部屋の報告であれば、その後その部屋に入室する）。
	Exit
)

func (a Action) String() string {
	switch a {
	case Keep:
		return "Keep"
	case Update:
		return "Update"
	case Exit:
		return "Exit"
	}
	return "Unknown"
}

// Current は現在Stayerテーブルに登録されている在室状態。
type Current struct {
	RoomID   int64
	Rssi     int64
	LastSeen time.Time // 今いる部屋の受信機が最後にこのユーザを観測した時刻
}

// Observation は今回のリクエスト（ある部屋の受信機）での観測結果。
type Observation struct {
	RoomID int64
	Seen   bool
	Rssi   int64
}

// Decide は既存の在室状態と今回の観測から取るべき処理を返す。
func Decide(cur Current, obs Observation, now time.Time) Action {
	stale := now.Sub(cur.LastSeen) > AbsenceGrace()

	if obs.RoomID == cur.RoomID {
		if obs.Seen {
			// 弱くなった場合も更新する（直近の値で他部屋と比較するため）
			return Update
		}
		// 1回の取りこぼしでは退室にしない
		if stale {
			return Exit
		}
		return Keep
	}

	// 別の部屋の受信機からの報告
	if !obs.Seen {
		return Keep
	}
	// 今の部屋で最近見えていない、または別部屋の方が明確に強い場合のみ移動
	if stale || obs.Rssi >= cur.Rssi+RoomSwitchRssiMargin() {
		return Exit
	}
	return Keep
}
