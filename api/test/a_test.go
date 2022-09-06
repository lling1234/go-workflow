package test

import (
	"act/api/flow"
	"fmt"
	"github.com/mumushuiding/util"
	"testing"
)

func TestNodeInfo(t *testing.T) {
	str := "[{\"nodeId\":\"开始\",\"type\":\"starter\",\"assignedType\":\"\",\"level\":1,\"mode\":\"or\",\"approverNames\":\"\",\"approverIds\":\"\"},{\"nodeId\":\"node_952545981591\",\"type\":\"APPROVAL\",\"assignedType\":\"ASSIGN_USER\",\"level\":2,\"mode\":\"and\",\"approverNames\":\"李四\",\"approverIds\":\"61769798\"},{\"nodeId\":\"node_981078773171\",\"type\":\"APPROVAL\",\"assignedType\":\"ASSIGN_USER\",\"level\":3,\"mode\":\"and\",\"approverNames\":\"旅人\",\"approverIds\":\"381496\"},{\"nodeId\":\"结束\",\"type\":\"\",\"assignedType\":\"\",\"level\":4,\"mode\":\"\",\"approverNames\":\"\",\"approverIds\":\"\"}]"
	var nodeInfos []*flow.NodeInfo
	util.Str2Struct(str, &nodeInfos)
<<<<<<< HEAD

	fmt.Println(nodeInfos[3])
=======
	fmt.Println(nodeInfos[1])
}

func TestRefuseNode(t *testing.T) {
	str := "{\"id\":\"root\",\"name\":\"发起人\",\"type\":\"ROOT\",\"children\":{\"id\":\"node_630130407989\",\"parentId\":\"root\",\"props\":{\"assignedType\":\"ASSIGN_USER\",\"mode\":\"AND\",\"assignedUser\":[{\"id\":61769798,\"name\":\"李四\",\"type\":\"user\",\"selected\":false}],\"refuse\":{\"type\":\"TO_BEFORE\",\"target\":\"\"}},\"type\":\"APPROVAL\",\"name\":\"审批人\",\"children\":{\"id\":\"node_630249811630\",\"parentId\":\"node_630130407989\",\"props\":{\"assignedType\":\"ASSIGN_USER\",\"mode\":\"AND\",\"assignedUser\":[{\"id\":381496,\"name\":\"旅人\",\"type\":\"user\",\"selected\":false}],\"refuse\":{\"type\":\"TO_BEFORE\",\"target\":\"\"}},\"type\":\"APPROVAL\",\"name\":\"审批人\",\"children\":{}}},\"parentId\":null}"
	var node *flow.Node
	util.Str2Struct(str, &node)
	fmt.Println(node.Children.Props.Refuse.Type)
}

func TestUsers(t *testing.T) {
	str := "{\"id\":\"root\",\"name\":\"发起人\",\"type\":\"ROOT\",\"children\":{\"id\":\"node_630130407989\",\"parentId\":\"root\",\"props\":{\"assignedType\":\"ASSIGN_USER\",\"mode\":\"AND\",\"assignedUser\":[{\"id\":6418616,\"name\":\"张三\",\"type\":\"user\",\"sex\":null,\"selected\":false},{\"id\":61769798,\"name\":\"李四\",\"type\":\"user\",\"sex\":null,\"selected\":false}],\"refuse\":{\"type\":\"TO_BEFORE\",\"target\":\"\"}},\"type\":\"APPROVAL\",\"name\":\"审批人\",\"children\":{\"id\":\"node_630249811630\",\"parentId\":\"node_630130407989\",\"props\":{\"assignedType\":\"ASSIGN_USER\",\"mode\":\"AND\",\"assignedUser\":[{\"id\":381496,\"name\":\"旅人\",\"type\":\"user\",\"selected\":false}],\"refuse\":{\"type\":\"TO_BEFORE\",\"target\":\"\"}},\"type\":\"APPROVAL\",\"name\":\"审批人\",\"children\":{}}},\"parentId\":null}"
	var node *flow.Node
	util.Str2Struct(str, &node)
	fmt.Printf("第一个审批人：%s,第二个审批人：%s", node.Children.Props.AssignedUser[0].Name, node.Children.Props.AssignedUser[1].Name)
>>>>>>> origin/yxq
}
