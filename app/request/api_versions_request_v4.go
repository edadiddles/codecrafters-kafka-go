package request

type ApiVersionsRequestV4 struct {
    ClientId Client
    ClientSoftwareVersion ClientSoftware
    TagBuffer [1]byte
}


type ClientSoftware struct {
    Length int
    Contents []byte
}

