package response

type ApiVersionResponseV4 struct {
    ErrorCode [2]byte
    ApiVersionsArr ApiVersionsArray
    ThrottleTimeMs [4]byte
    TagBuffer [1]byte
}

func (r ApiVersionResponseV4) Size() int {
    size := 0
    size += len(r.ErrorCode)
    size += r.ApiVersionsArr.Size()
    size += len(r.ThrottleTimeMs)
    size += len(r.TagBuffer)

    return size
}

func (r ApiVersionResponseV4) Serialize() []byte {
    msg := make([]byte, 0)
    msg = append(msg, r.ErrorCode[:]...)
    msg = append(msg, r.ApiVersionsArr.Serialize()...)
    msg = append(msg, r.ThrottleTimeMs[:]...)
    msg = append(msg, r.TagBuffer[:]...)

    return msg
}

type ApiVersionsArray struct {
    MessageSize [1]byte
    ApiVersions []ApiVersion
}

func (ava ApiVersionsArray) Size() int {
    size := 0
    size += len(ava.MessageSize)
    for _, api_version := range ava.ApiVersions {
        size += api_version.Size()
    }

    return size
}

func (ava ApiVersionsArray) Serialize() []byte {
    msg := make([]byte, 0)

    msg = append(msg, ava.MessageSize[:]...)
    for _, apiVersion := range ava.ApiVersions {
        msg = append(msg, apiVersion.Serialize()...)
    }

    return msg
}

type ApiVersion struct {
    ApiKey [2]byte
    MinVersion [2]byte
    MaxVersion [2]byte
    TagBuffer [1]byte
}

func (av ApiVersion) Size() int {
    size := 0
    size += len(av.ApiKey)
    size += len(av.MinVersion)
    size += len(av.MaxVersion)
    size += len(av.TagBuffer)

    return size
}

func (av ApiVersion) Serialize() []byte {
    msg := make([]byte, 0)

    msg = append(msg, av.ApiKey[:]...)
    msg = append(msg, av.MinVersion[:]...)
    msg = append(msg, av.MaxVersion[:]...)
    msg = append(msg, av.TagBuffer[:]...)

    return msg
}
