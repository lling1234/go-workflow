package flow

import (
	"container/list"
	"errors"
	"github.com/mumushuiding/util"
	"log"
	"strconv"
)

// Node represents a specific logical unit of processing and routing
// in a workflow.
// 流程中的一个节点
type Node struct {
	Name      string     `json:"name,optional"`
	Type      string     `json:"type,optional"`
	NodeID    string     `json:"nodeId,optional"`
	ParentId  string     `json:"parentId,optional"`
	Children  *Node      `json:"children,optional"`
	Branches  []*Node    `json:"branches,optional"`
	Props     *NodeProps `json:"props,optional"`
	CondProps *CondProps `json:"conditionProps,optional"`
}

// NodeType 节点类型
type NodeType int

const (
	ROOT NodeType = iota
	CONDITIONS
	CONDITION
	APPROVAL
	NOTIFIER
	WEBHOOK
	CONCURRENTS
)

// NodeTypes 节点类型
var NodeTypes = [...]string{ROOT: "ROOT", CONDITIONS: "CONDITIONS", CONDITION: "CONDITION", APPROVAL: "APPROVAL", NOTIFIER: "NOTIFIER", WEBHOOK: "webhook", CONCURRENTS: "concurrents"}

type NodeProps struct {
	AssignedType string          `json:"assignedType,optional"`
	Mode         string          `json:"mode,optional"`
	Station      string          `json:"station,optional,optional"`
	AssignedUser []*AssignedUser `json:"assignedUser,optional"`
	Refuse       Refuse          `json:"refuse,optional"`
}
type Refuse struct {
	Type   string `json:"type,optional"`
	Target string `json:"target,optional"`
}
type CondProps struct {
	GroupsType string `json:"groupsType,optional"`
	Groups     string `json:"groups,optional"`
	Expression string `json:"expression,optional"`
}
type AssignedUser struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Selected bool   `json:"selected,optional"`
}

// NodeInfo 节点信息
type NodeInfo struct {
	NodeID       string `json:"nodeId"`
	Type         string `json:"type"`
	AssignedType string `json:"assignedType"`
	//MemberCount   int    `json:"memberCount"`
	Level         int    `json:"level,optional"`
	Mode          string `json:"mode"`
	ApproverNames string `json:"approverNames"`
	ApproverIds   string `json:"approverIds"`
}

const (
	PENDING = iota + 1
	DEALING
	REJECT
	WITHDRAW
	NOTPASS
	HAVEPASS
	DISCARD
)

var ResultTypes = [...]string{PENDING: "待处理", DEALING: "处理中", REJECT: "驳回", WITHDRAW: "已撤回", NOTPASS: "未通过", HAVEPASS: "已通过", DISCARD: "废弃"}

func (n *Node) add2ExecutionList(list *list.List) {
	switch n.Type {
	case NodeTypes[APPROVAL], NodeTypes[NOTIFIER]:
		list.PushBack(NodeInfo{
			NodeID:       n.NodeID,
			Type:         n.Type,
			Level:        list.Len() + 2,
			AssignedType: n.Props.AssignedType,
			//MemberCount:   n.Props.getMemberCount(),
			Mode:          n.Props.Mode,
			ApproverIds:   n.Props.getApproverIds(),
			ApproverNames: n.Props.getApproverNames(),
		})
		break
	default:
	}
}
func (n *NodeProps) getApproverIds() string {
	if n.AssignedType == "user" || n.AssignedType == "ASSIGN_USER" {
		str := ""
		for _, v := range n.AssignedUser {
			str += strconv.FormatInt(v.ID, 10) + ","
		}
		return str[0 : len(str)-1]
	}
	return ""
}
func (n *NodeProps) getApproverNames() string {
	if n.AssignedType == "user" || n.AssignedType == "ASSIGN_USER" {
		str := ""
		for _, v := range n.AssignedUser {
			str += v.Name + ","
		}
		return str[0 : len(str)-1]
	}
	return ""
}

//func (n *NodeProps) getMemberCount() int {
//	if n.AssignedUser != nil {
//		if n.ActMode == "or" {
//			return 1
//		}
//		return len(n.AssignedUser)
//	}
//	return 0
//}
//
//// IfProcessConifgIsValid 检查流程配置是否有效
//func IfProcessConifgIsValid(node *Node) error {
//	// 节点名称是否有效
//	if len(node.NodeID) == 0 {
//		return errors.New("节点的【nodeId】不能为空！！")
//	}
//	// 检查类型是否有效
//	if len(node.Type) == 0 {
//		return errors.New("节点【" + node.NodeID + "】的类型【type】不能为空")
//	}
//	var flag = false
//	for _, val := range NodeTypes {
//		if val == node.Type {
//			flag = true
//			break
//		}
//	}
//	if !flag {
//		str, _ := util.ToJSONStr(NodeTypes)
//		return errors.New("节点【" + node.NodeID + "】的类型为【" + node.Type + "】，为无效类型,有效类型为" + str)
//	}
//	// 当前节点是否设置有审批人
//	if node.Type == NodeTypes[APPROVAL] || node.Type == NodeTypes[NOTIFIER] {
//		if node.Props == nil {
//			return errors.New("节点【" + node.NodeID + "】的Properties属性不能为空，如：`\"properties\": {\"actionerRules\": [{\"type\": \"target_label\",\"labelNames\": \"人事\",\"memberCount\": 1,\"actMode\": \"and\"}],}`")
//		}
//	}
//	// 条件节点是否存在
//	if node.Type != NodeTypes[CONDITIONS] { // 存在条件节点
//		//if len(node.ConditionNodes) == 1 {
//		//	return errors.New("节点【" + node.NodeID + "】条件节点下的节点数必须大于1")
//		//}
//		// 根据条件变量选择节点索引
//		err := CheckConditionNode(node.Branches)
//		if err != nil {
//			return err
//		}
//	}
//
//	// 子节点是否存在
//	if node.Children != nil {
//		return IfProcessConifgIsValid(node.Children)
//	}
//	return nil
//}

// CheckConditionNode 检查条件节点
func CheckConditionNode(nodes []*Node) error {
	for _, node := range nodes {
		if node.Props == nil {
			return errors.New("节点【" + node.NodeID + "】的Properties对象为空值！！")
		}
		if len(node.Branches) == 0 {
			return errors.New("节点【" + node.NodeID + "】的Conditions对象为空值！！")
		}
		err := IfProcessConifgIsValid(node)
		if err != nil {
			return err
		}
	}
	return nil
}

// ParseProcessConfig 解析流程定义json数据
func ParseProcessConfig(node *Node) (*list.List, error) {
	list := list.New()
	err := parseProcessConfig(node, list)
	return list, err
}

func parseProcessConfig(node *Node, list *list.List) (err error) {
	if node != nil {
		node.add2ExecutionList(list)
		parseProcessConfig(node.Children, list)
	}
	return nil
}

// IfProcessConifgIsValid 检查流程配置是否有效
func IfProcessConifgIsValid(node *Node) error {
	log.Println(111)
	// 节点名称是否有效
	if len(node.NodeID) == 0 {
		log.Println("NodeID", node.NodeID)
		return errors.New("节点的【nodeId】不能为空！！")
	}
	log.Println(222)
	// 检查类型是否有效
	if len(node.Type) == 0 {
		return errors.New("节点【" + node.NodeID + "】的类型【type】不能为空")
	}
	log.Println(333)
	var flag = false
	for _, val := range NodeTypes {
		if val == node.Type {
			flag = true
			break
		}
	}
	if !flag {
		str, _ := util.ToJSONStr(NodeTypes)
		return errors.New("节点【" + node.NodeID + "】的类型为【" + node.Type + "】，为无效类型,有效类型为" + str)
	}
	log.Println(4444)
	// 当前节点是否设置有审批人
	if node.Type == NodeTypes[APPROVAL] || node.Type == NodeTypes[NOTIFIER] {
		if node.Props == nil || (node.Props.AssignedUser == nil && node.Props.Station == "") {
			return errors.New("节点【" + node.NodeID + "】的Properties属性不能为空，如：`\"properties\": {\"actionerRules\": [{\"type\": \"target_label\",\"labelNames\": \"人事\",\"memberCount\": 1,\"actType\": \"and\"}],}`")
		}
	}
	log.Println(5555)
	// 条件节点是否存在
	if node.Branches != nil && node.Branches[0].CondProps != nil { // 存在条件节点
		if len(node.Branches) == 1 {
			return errors.New("节点【" + node.NodeID + "】条件节点下的节点数必须大于1")
		}
		// 根据条件变量选择节点索引
		err := CheckConditionNode(node.Branches)
		if err != nil {
			return err
		}
	}
	log.Println(666)
	// 子节点是否存在
	if node.Children != nil && node.Children.NodeID != "" {
		return IfProcessConifgIsValid(node.Children)
	}
	return nil
}
