package flow

import (
	"errors"
	"github.com/mumushuiding/util"
	"log"
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
	END //虚拟的末级节点
)

// NodeTypes 节点类型
var NodeTypes = [...]string{ROOT: "ROOT", CONDITIONS: "CONDITIONS", CONDITION: "CONDITION", APPROVAL: "APPROVAL", NOTIFIER: "NOTIFIER", WEBHOOK: "WEBHOOK", CONCURRENTS: "CONCURRENTS", END: "END"}

type NodeProps struct {
	Mode         string          `json:"mode,optional"`
	Station      string          `json:"station,optional,optional"`
	AssignedUser []*AssignedUser `json:"assignedUser,optional"`
	Refuse       *Refuse         `json:"refuse,optional"`
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
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// NodeInfo 节点信息
type NodeInfo struct {
	NodeID       string          `json:"nodeId"`
	Type         string          `json:"type"`
	NextId       string          `json:"nextId,optional"`
	PrevId       string          `json:"prevId,optional"`
	Mode         string          `json:"mode,optional"`
	AssignedUser []*AssignedUser `json:"assignedUser,optional"`
	Refuse       *Refuse         `json:"refuse,optional"`
}

const (
	PENDING = iota
	DEALING
	REJECTTOROOT
	REJECT
	WITHDRAW
	NOTPASS
	HAVEPASS
	DISCARD
	VIEW //增加这个状态是为了识别抄送
)

var ResultTypes = [...]string{PENDING: "待处理", DEALING: "处理中", REJECTTOROOT: "驳回至发起人", REJECT: "驳回", WITHDRAW: "已撤回", NOTPASS: "未通过", HAVEPASS: "已通过", DISCARD: "废弃", VIEW: "抄送"}

//驳回的三种类型
const (
	TO_BEFORE string = "TO_BEFORE"
	TO_NODE   string = "TO_NODE"
	TO_END    string = "TO_END"
)

//或签/会签
const (
	MODE_OR  string = "OR"
	MODE_AND string = "AND"
)

func SliceApproverIds(au []*AssignedUser) []int64 {
	var approverIds []int64
	if au != nil && len(au) > 0 {
		for _, v := range au {
			approverIds = append(approverIds, v.ID)
		}
	}
	return approverIds
}
func SliceApproverNames(au []*AssignedUser) []string {
	var approverNames []string
	if au != nil && len(au) > 0 {
		for _, v := range au {
			approverNames = append(approverNames, v.Name)
		}
	}
	return approverNames
}

// CheckConditionNode 检查条件节点
func CheckConditionNode(nodes []*Node) error {
	for _, node := range nodes {
		if node.Props == nil {
			return errors.New("节点【" + node.NodeID + "】的Properties对象为空值！！")
		}
		if len(node.Branches) == 0 {
			return errors.New("节点【" + node.NodeID + "】的Conditions对象为空值！！")
		}
		err := IfProcessConfigIsValid(node)
		if err != nil {
			return err
		}
	}
	return nil
}

// IfProcessConifgIsValid 检查流程配置是否有效
func IfProcessConfigIsValid(node *Node) error {
	// 节点名称是否有效
	if len(node.NodeID) == 0 {
		log.Println("NodeID", node.NodeID)
		return errors.New("节点的【nodeId】不能为空！！")
	}
	log.Println(11111)
	// 检查类型是否有效
	if node.Type == "" && len(node.Type) == 0 {
		return errors.New("节点【" + node.NodeID + "】的类型【type】不能为空")
	}
	log.Println(222222)
	var flag = false
	for _, val := range NodeTypes {
		if val == node.Type {
			flag = true
			break
		}
	}
	log.Println(33333)
	if !flag {
		str, _ := util.ToJSONStr(NodeTypes)
		return errors.New("节点【" + node.NodeID + "】的类型为【" + node.Type + "】，为无效类型,有效类型为" + str)
	}
	log.Println(44444)
	// 当前节点是否设置有审批人
	if node.Type == NodeTypes[APPROVAL] || node.Type == NodeTypes[NOTIFIER] {
		if node.Props == nil || (node.Props.AssignedUser == nil && node.Props.Station == "") {
			return errors.New("节点【" + node.NodeID + "】的Properties属性不能为空，如：`\"properties\": {\"actionerRules\": [{\"type\": \"target_label\",\"labelNames\": \"人事\",\"memberCount\": 1,\"actType\": \"and\"}],}`")
		}
	}
	log.Println(55555)
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
	log.Println(66666)
	// 子节点是否存在
	if node.Children != nil && node.Children.NodeID != "" {
		return IfProcessConfigIsValid(node.Children)
	}
	return nil
}

func CheckConCurrentFlow(node *Node) bool {
	if node != nil && node.Name != "" {
		if node.Type == "CONCURRENT" {
			return true
		} else {
			if node.Branches != nil && len(node.Branches) > 1 {
				for _, v := range node.Branches {
					if v.Type == "CONCURRENT" {
						return true
					}
				}
			}
			CheckConCurrentFlow(node.Children)
		}

	}
	return false
}

//返回节点信息和抄送人信息
func ShiftNodeTree(n *Node, l []*NodeInfo, notifiers []int64) ([]*NodeInfo, []int64) {
	root := NodeTypes[ROOT]
	var assignedUsers []*AssignedUser
	mode := MODE_OR
	var refuse *Refuse
	if n.Props != nil {
		log.Println("n.Props.Refuse", n.Props.Refuse)
		if n.Props.AssignedUser != nil {
			assignedUsers = n.Props.AssignedUser
		}
		if n.Props.Mode != "" {
			mode = n.Props.Mode
		}
		if n.Props.Refuse != nil {
			refuse = n.Props.Refuse
		}
	}
	if n.NodeID == root {
		log.Println("root")
		l = append(l, &NodeInfo{
			NodeID:       root,
			Type:         root,
			Mode:         MODE_OR,
			AssignedUser: assignedUsers,
			Refuse:       refuse,
		})
	} else if n.Children == nil || n.Children.NodeID == "" {
		//补充上一个NodeInfo的nextId
		lastInfo := l[len(l)-1]
		if n.Type == NodeTypes[NOTIFIER] {
			lastInfo.NextId = NodeTypes[END]
			for _, v := range assignedUsers {
				if v.Type == "user" {
					notifiers = append(notifiers, v.ID)
				}
			}
		} else {
			lastInfo.NextId = n.NodeID
			l = append(l, &NodeInfo{
				NodeID:       n.NodeID,
				Type:         n.Type,
				Mode:         mode,
				PrevId:       lastInfo.NodeID,
				AssignedUser: assignedUsers,
				Refuse:       refuse,
				NextId:       NodeTypes[END],
			})
		}
		return l, notifiers
	} else {
		if n.Type == NodeTypes[NOTIFIER] {
			for _, v := range assignedUsers {
				if v.Type == "user" {
					notifiers = append(notifiers, v.ID)
				}
			}
		} else {
			//补充上一个NodeInfo的nextId
			lastInfo := l[len(l)-1]
			lastInfo.NextId = n.NodeID
			l = append(l, &NodeInfo{
				NodeID:       n.NodeID,
				Type:         n.Type,
				Mode:         mode,
				PrevId:       lastInfo.NodeID,
				AssignedUser: assignedUsers,
				Refuse:       refuse,
			})
		}
	}
	return ShiftNodeTree(n.Children, l, notifiers)
}

//判断并行节点是否为末级节点
//func IsEndNodes(l []*NodeInfo) bool {
//	if l != nil && len(l) > 0 {
//		for _, v := range l {
//			if v.NextId != "end" {
//				return false
//			}
//		}
//	}
//	return true
//}

//判断该节点是否为末级节点
//func IsEndNode(l *NodeInfo) int32 {
//	if l != nil && l.NextId != "end" {
//		return 0
//	}
//	return 1
//}

//串行流第一个节点
func FindFirstApprovalNode(n *Node) (*Node, error) {
	if n.Type == NodeTypes[NOTIFIER] {
		if n.Children != nil && n.Children.NodeID != "" {
			return FindFirstApprovalNode(n.Children)
		} else {
			return nil, errors.New("流程设计错误，没有执行节点")
		}
	}
	return n, nil
}

//func FindFirstApprovalNode(n *Node, prevId string) (*NodeInfo, error) {
//	if n != nil && n.NodeID != "" {
//		if n.Type != NodeTypes[NOTIFIER] {
//			nextNode := n.Children
//			var nextNodeNextId string
//			if n.Children.Children != nil && n.Children.Children.NodeID != "" {
//				nextNodeNextId = n.Children.Children.NodeID
//			} else {
//				nextNodeNextId = "end"
//			}
//			var refuse *Refuse
//			if nextNode.Props.Refuse != nil {
//				refuse = nextNode.Props.Refuse
//			}
//			return &NodeInfo{
//				NodeID:       n.NodeID,
//				Type:         n.Type,
//				NextId:       nextNodeNextId,
//				PrevId:       prevId,
//				Mode:         nextNode.Props.Mode,
//				AssignedUser: nextNode.Props.AssignedUser,
//				Refuse:       refuse,
//			}, nil
//		}
//		return FindFirstApprovalNode(n.Children, n.NodeID)
//	} else {
//		return nil, errors.New("错误解析")
//	}
//
//}
