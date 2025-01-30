package request

import (
    "github.com/codecrafters-io/kafka-starter-go/app/utils"
)

type DescribeTopicPartitionsRequestV0 struct {
    TopicsArr TopicsArray
    ResponsePartitionLimit [4]byte
    Cursor [1]byte
    TagBuffer [1]byte
}

func (r *DescribeTopicPartitionsRequestV0) Deserialize(req []byte) []byte {
    req = r.TopicsArr.Deserialize(req)
    r.ResponsePartitionLimit = [4]byte(req[0:4])
    r.Cursor = [1]byte(req[4:5])

    return req[6:]
}

type TopicsArray struct {
    Length int
    Topics []Topic
}

func (ta *TopicsArray) Deserialize(req []byte) []byte {
    ta.Length = utils.Bytes_to_int(req[0:1])

    req = req[1:]
    for i:=0; i < ta.Length-1; i++ {
        t := Topic{}
        req = t.Deserialize(req)
        ta.Topics = append(ta.Topics, t)
    }

    return req
}

type Topic struct {
    Length int
    Name []byte
    TagBuffer [1]byte
}

func (t *Topic) Deserialize(req []byte) []byte {
    t.Length = utils.Bytes_to_int(req[0:1])
    t.Name = req[1:t.Length]

    return req[t.Length+1:]
}
