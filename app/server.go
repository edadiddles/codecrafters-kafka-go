package main

import (
	"fmt"
	"net"
	"os"
    "bytes"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

    l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
	    fmt.Println("Failed to bind to port 9092")
	    os.Exit(1)
	}
    conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
	    os.Exit(1)
	}
    buff := make([]byte, 1024)
    conn.Read(buff)
    // msg = 4, req_api_key = 2, req_api_vers = 2, correlation_id = 4
    request_api_version := buff[6:8]
    correlation_id := buff[8:12]
    message_size := []byte{0, 0, 0, 0}
    supported_api_versions := [][]byte{
        {0, 0}, // version 0
        {0, 1}, // version 1
        {0, 2}, // version 2
        {0, 3}, // version 3
        {0, 4}, // version 4
    }
    is_api_version_supported := false
    for i := 0; i < len(supported_api_versions); i++ {
        if bytes.Equal(request_api_version, supported_api_versions[i]) {
            is_api_version_supported = true
        }
    }
    
    // capture error
    error_code := []byte{0,0}
    if !is_api_version_supported {
        error_code = []byte{0, 35}
    }

    conn.Write(message_size)
    conn.Write(correlation_id)
    conn.Write(error_code)
}

