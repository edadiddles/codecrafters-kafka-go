package response

type ResponseHeaderV1 struct {
    CorrelationId [4]byte
    TagBuffer [1]byte
}

func (h ResponseHeaderV1) Size() int {
    size := 0
    size += len(h.CorrelationId)
    size += len(h.TagBuffer)

    return size 
}

func (h ResponseHeaderV1) Serialize() []byte {
    return append(h.CorrelationId[:], h.TagBuffer[:]...)
}
