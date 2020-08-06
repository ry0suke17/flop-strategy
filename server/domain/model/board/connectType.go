package board

// ConnectType はボードのコネクトタイプを表す、
type ConnectType int

const (
	// ConnectTypeConnected はコネクトしているボードを表す。
	ConnectTypeConnected ConnectType = 1
	// ConnectTypeDisconnected はコネクトしていないボードを表す。
	ConnectTypeDisconnected ConnectType = 2
)
