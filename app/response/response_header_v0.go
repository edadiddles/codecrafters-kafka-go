package response

type ResponseHeaderV0 struct {
    CorrelationId [4]byte
}

func (h ResponseHeaderV0) Size() int {
    size := 0
    size += len(h.CorrelationId)

    return size
}

func (h ResponseHeaderV0) Serialize() []byte {
    return h.CorrelationId[:]
}
