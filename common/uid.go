package common

type UID struct {
	localID    uint32
	objectType int
	shardID    uint32
}

// UID là phương thức để tạo ra một id ảo cho toàn bộ hệ thống
// 32bit cho local ID
// 10bit cho object type
// 18 bit cho shard ID

func NewUID(localID uint32, objectType int, shardID uint32) UID {
	return UID{
		localID:    localID,
		objectType: objectType,
		shardID:    shardID,
	}
}
