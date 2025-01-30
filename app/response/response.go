package response

import (
    "github.com/codecrafters-io/kafka-starter-go/app/utils"
)

type IResponse interface {
    Size() int
    Serialize() []byte
}

type IResponseHeader interface {
    Size() int
    Serialize() []byte
}

type IResponseBody interface {
    Size() int
    Serialize() []byte
}

type Response struct {
    MessageSize [4]byte
    Header IResponseHeader
    Body IResponseBody
}

func (r Response) Size() int {
    size := 0
    //size += len(r.MessageSize)
    size += r.Header.Size()
    size += r.Body.Size()

    return size
}

func (r Response) Serialize() []byte {
    msg := make([]byte, 0)
 
    r.MessageSize = [4]byte(utils.Int_to_bytes(r.Size(), 4))
    msg = append(msg, r.MessageSize[:]...)
    msg = append(msg, r.Header.Serialize()...)
    msg = append(msg, r.Body.Serialize()...)

    return msg
}


