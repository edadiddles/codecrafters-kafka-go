package request

import (
    "github.com/codecrafters-io/kafka-starter-go/app/utils"
)

type RequestHeaderV2 struct {
    ApiKey int
    ApiVersion int
    CorrelationId [4]byte
    ClientId Client
    TagBuffer [1]byte
}

func (h *RequestHeaderV2) Deserialize(req []byte) []byte {
    h.ApiKey = utils.Bytes_to_int(req[0:2])
    h.ApiVersion = utils.Bytes_to_int(req[2:4])
    h.CorrelationId = [4]byte(req[4:8])
    req = h.ClientId.Deserialize(req[8:])
    //TagBuffer
    return req[1:]
}

type Client struct {
    Length int
    Contents []byte
}

func (c *Client) Deserialize(req []byte) []byte {
    c.Length = utils.Bytes_to_int(req[0:2])
    c.Contents = []byte(req[2:c.Length+2])
    return req[c.Length+2:]
}
