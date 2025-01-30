package request


type IRequestHeader interface {
    Deserialize([]byte) []byte
}

type IRequestBody interface {
    Deserialize([]byte) []byte
}
