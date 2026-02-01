package plansummary

//TODO: formatter系は実装途中
type ViewOption interface {
	Colored() bool
	// 後で拡張
	//MaxWidth() int
}
