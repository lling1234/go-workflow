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
	fmt.Println(nodeInfos[3])
}
