package constant

// PENDING: "待处理", DEALING: "处理中", REJECT: "驳回", WITHDRAW: "已撤回", NOTPASS: "未通过", HAVEPASS: "已通过", DISCARD: "废弃"
const (
	PENDING = iota + 1
	DEALING
	REJECT
	WITHDRAW
	NOTPASS
	HAVEPASS
	DISCARD
)
