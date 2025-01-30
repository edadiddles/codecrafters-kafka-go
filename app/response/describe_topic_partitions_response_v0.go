package response

type DescribeTopicPartitionsResponseV0 struct {
    ThrottleTimeMs [4]byte
    TopicsArr TopicsArray
    NextCursor Cursor
    TagBuffer [1]byte
}

func (b DescribeTopicPartitionsResponseV0) Size() int {
    size := 0
    size += len(b.ThrottleTimeMs)
    size += b.TopicsArr.Size()
    size += b.NextCursor.Size()
    size += len(b.TagBuffer)

    return size
}

func (b DescribeTopicPartitionsResponseV0) Serialize() []byte {
    msg := make([]byte, 0)

    msg = append(msg, b.ThrottleTimeMs[:]...)
    msg = append(msg, b.TopicsArr.Serialize()...)
    msg = append(msg, b.NextCursor.Serialize()...)
    msg = append(msg, b.TagBuffer[:]...)

    return msg
}

type TopicsArray struct {
    MessageSize [1]byte
    Topics []Topic
}

func (ta TopicsArray) Size() int {
    size := 0
    size += len(ta.MessageSize)
    for _, topic := range ta.Topics {
        size += topic.Size()
    }
    return size
}

func (ta TopicsArray) Serialize() []byte {
    msg := make([]byte, 0)

    msg = append(msg, ta.MessageSize[:]...)
    for _, topic := range ta.Topics {
        msg = append(msg, topic.Serialize()...)
    }

    return msg
}

type Topic struct {
    ErrorCode [2]byte
    Name CompactString
    TopicId [16]byte
    IsInternal [1]byte
    PartitionsArr PartitionsArray
    TopicAuthorizedOperations [4]byte
    TagBuffer [1]byte
}

func (t Topic) Size() int {
    size := 0
    size += len(t.ErrorCode)
    size += t.Name.Size()
    size += len(t.TopicId)
    size += len(t.IsInternal)
    size += t.PartitionsArr.Size()
    size += len(t.TopicAuthorizedOperations)
    size += len(t.TagBuffer)
    
    return size
}

func (t Topic) Serialize() []byte {
    msg := make([]byte, 0)

    msg = append(msg, t.ErrorCode[:]...)
    msg = append(msg, t.Name.Serialize()...)
    msg = append(msg, t.TopicId[:]...)
    msg = append(msg, t.IsInternal[:]...)
    msg = append(msg, t.PartitionsArr.Serialize()...)
    msg = append(msg, t.TopicAuthorizedOperations[:]...)
    msg = append(msg, t.TagBuffer[:]...)

    return msg
}

type CompactString struct {
    Length [1]byte
    Contents []byte
}

func (s CompactString) Size() int {
    size := 0
    size += len(s.Length)
    size += len(s.Contents)

    return size
}

func (s CompactString) Serialize() []byte {
    msg := make([]byte, 0)

    msg = append(msg, s.Length[:]...)
    msg = append(msg, s.Contents...)

    return msg
}

type PartitionsArray struct {
    MessageSize [1]byte
    Partitions []Partition
}

func (pa PartitionsArray) Size() int {
    size := 0
    size += len(pa.MessageSize)
    for _, partition := range pa.Partitions {
        size += partition.Size()
    }

    return size
}

func (pa PartitionsArray) Serialize() []byte {
    msg := make([]byte, 0)

    msg = append(msg, pa.MessageSize[:]...)
    for _, partition := range pa.Partitions {
        msg = append(msg, partition.Serialize()...)
    }

    return msg
}

type Partition struct {
    ErrorCode [2]byte
    Index [4]byte
    LeaderId [4]byte
    LeaderEpoch [4]byte
    ReplicaNodes [4]byte
    IsrNodes [4]byte
    EligibleLeaderReplicas [4]byte
    LastKnownElr [4]byte
    OfflineReplicas [4]byte
    TagBuffer [1]byte
}

func (p Partition) Size() int {
    size := 0
    size += len(p.ErrorCode)
    size += len(p.Index)
    size += len(p.LeaderId)
    size += len(p.LeaderEpoch)
    size += len(p.ReplicaNodes)
    size += len(p.IsrNodes)
    size += len(p.EligibleLeaderReplicas)
    size += len(p.LastKnownElr)
    size += len(p.OfflineReplicas)
    size += len(p.TagBuffer)

    return size
}

func (p Partition) Serialize() []byte {
    msg := make([]byte, 0)

    msg = append(msg, p.ErrorCode[:]...)
    msg = append(msg, p.Index[:]...)
    msg = append(msg, p.LeaderId[:]...)
    msg = append(msg, p.LeaderEpoch[:]...)
    msg = append(msg, p.ReplicaNodes[:]...)
    msg = append(msg, p.IsrNodes[:]...)
    msg = append(msg, p.EligibleLeaderReplicas[:]...)
    msg = append(msg, p.LastKnownElr[:]...)
    msg = append(msg, p.OfflineReplicas[:]...)
    msg = append(msg, p.TagBuffer[:]...)

    return msg
}

type Cursor struct {
    TopicName CompactString
    PartitionIndex [4]byte
    TagBuffer [1]byte
}

func (c Cursor) Size() int {
    size := 0
    size += c.TopicName.Size()
    size += len(c.PartitionIndex)
    size += len(c.TagBuffer)

    return size
}

func (c Cursor) Serialize() []byte {
    msg := make([]byte, 0)

    msg = append(msg, c.TopicName.Serialize()...)
    msg = append(msg, c.PartitionIndex[:]...)
    msg = append(msg, c.TagBuffer[:]...)

    return msg
}
