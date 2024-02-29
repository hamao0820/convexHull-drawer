package graham

import (
	"sort"
)

func Scan(points []*Point) (convex []*Point) {
	// 3点以下ならそのまま返す
	if len(points) < 3 {
		return points
	}

	sort.Slice(points, func(i, j int) bool {
		return Le(points[i], points[j])
	})

	// 最も左下の点を探す
	minP := points[0]
	points_ := points[1:]

	// 角度でソート
	sort.Slice(points_, func(i, j int) bool {
		return IsClockwise(points_[i], minP, points_[j])
	})

	points = append([]*Point{minP}, points_...)

	// スタックを使って凸包を求める
	st := NewStack[*Point]()
	// 最初の3点を追加
	st.Push(points[0])
	st.Push(points[1])
	st.Push(points[2])
	for i := 3; i < len(points); i++ {
		// 時計回りなら取り除く
		for st.Len() >= 2 {
			if !IsClockwise(st.SemiTop(), st.Top(), points[i]) {
				break
			}
			st.Pop()
		}
		// 反時計回りなら追加
		st.Push(points[i])

	}

	return st.Slice()
}
