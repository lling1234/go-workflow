package service

import (
	"act/rpc/internal/flow"
	"act/rpc/internal/model"
	util2 "act/rpc/internal/workflow-util"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

func AddCandidate(nodeInfo *flow.NodeInfo, taskID int64, targetId int64, procInstID int64, tx *gorm.DB) error {
	ids := strings.Split(nodeInfo.ApproverIds, ",")
	names := strings.Split(nodeInfo.ApproverNames, ",")
	for i := 0; i < len(ids); i++ {
		approverId := ids[i]
		// ⬑ 这个不知道哪里来的，长度为3
		if len(approverId) > 4 {
			userId, _ := strconv.ParseInt(approverId, 10, 64)
			c := &model.Candidate{
				Identity: model.Identity{
					ID:         util2.NewSnowflakeId(),
					CreateTime: time.Now(),
				},
				UserId:     userId,
				UserName:   names[i],
				ProcInstID: procInstID,
				TaskID:     taskID,
				IsFinish:   2,
				TargetId:   targetId,
			}
			err := c.SaveTx(tx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
