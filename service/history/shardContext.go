package history

import (
	"sync/atomic"

	"code.uber.internal/devexp/minions/common/persistence"

	"github.com/uber-common/bark"
)

type (
	// ShardContext represents a history engine shard
	ShardContext interface {
		GetExecutionManager() persistence.ExecutionManager
		GetTransferTaskID() int64
		GetRangeID() int64
		GetTransferAckLevel() int64
		GetTimerSequenceNumber() int64
		UpdateAckLevel(ackLevel int64) error
		GetTransferSequenceNumber() int64
		GetLogger() bark.Logger
	}

	shardContextImpl struct {
		shardManager           persistence.ShardManager
		shardInfo              *persistence.ShardInfo
		executionManager       persistence.ExecutionManager
		transferSequenceNumber int64
		timerSequeceNumber     int64
		logger                 bark.Logger
	}
)

var _ ShardContext = (*shardContextImpl)(nil)

func (s *shardContextImpl) GetExecutionManager() persistence.ExecutionManager {
	return s.executionManager
}

func (s *shardContextImpl) GetTimerSequenceNumber() int64 {
	return atomic.AddInt64(&s.timerSequeceNumber, 1)
}

func (s *shardContextImpl) GetTransferTaskID() int64 {
	return atomic.AddInt64(&s.transferSequenceNumber, 1)
}

func (s *shardContextImpl) GetRangeID() int64 {
	return atomic.LoadInt64(&s.shardInfo.RangeID)
}

func (s *shardContextImpl) GetTransferAckLevel() int64 {
	return atomic.LoadInt64(&s.shardInfo.TransferAckLevel)
}

func (s *shardContextImpl) GetLogger() bark.Logger {
	return s.logger
}

func (s *shardContextImpl) UpdateAckLevel(ackLevel int64) error {
	atomic.StoreInt64(&s.shardInfo.TransferAckLevel, ackLevel)
	updatedShardInfo := copyShardInfo(s.shardInfo)
	updatedShardInfo.StolenSinceRenew = 0
	return s.shardManager.UpdateShard(&persistence.UpdateShardRequest{
		ShardInfo:       updatedShardInfo,
		PreviousRangeID: updatedShardInfo.RangeID,
	})
}

func (s *shardContextImpl) GetTransferSequenceNumber() int64 {
	return atomic.LoadInt64(&s.transferSequenceNumber)
}

func acquireShard(shardID int, shardManager persistence.ShardManager, executionMgr persistence.ExecutionManager,
	logger bark.Logger) (ShardContext, error) {
	response, err0 := shardManager.GetShard(&persistence.GetShardRequest{ShardID: shardID})
	if err0 != nil {
		return nil, err0
	}

	shardInfo := response.ShardInfo
	updatedShardInfo := copyShardInfo(shardInfo)
	updatedShardInfo.RangeID++
	updatedShardInfo.StolenSinceRenew++

	err1 := shardManager.UpdateShard(&persistence.UpdateShardRequest{
		ShardInfo:       updatedShardInfo,
		PreviousRangeID: shardInfo.RangeID})
	if err1 != nil {
		return nil, err1
	}

	context := &shardContextImpl{
		shardManager:           shardManager,
		shardInfo:              updatedShardInfo,
		executionManager:       executionMgr,
		transferSequenceNumber: updatedShardInfo.RangeID << 24,
	}
	context.logger = logger.WithFields(bark.Fields{
		tagHistoryShardID: shardID,
	})

	return context, nil
}

func copyShardInfo(shardInfo *persistence.ShardInfo) *persistence.ShardInfo {
	copy := &persistence.ShardInfo{
		ShardID:          shardInfo.ShardID,
		Owner:            shardInfo.Owner,
		RangeID:          shardInfo.RangeID,
		StolenSinceRenew: shardInfo.StolenSinceRenew,
		TransferAckLevel: atomic.LoadInt64(&shardInfo.TransferAckLevel),
	}

	return copy
}
