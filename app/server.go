package main

import (
	"fmt"
	"net"
	"os"
    "time"

    "github.com/codecrafters-io/kafka-starter-go/app/request"
    "github.com/codecrafters-io/kafka-starter-go/app/response"
    "github.com/codecrafters-io/kafka-starter-go/app/utils"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

    l, err := net.Listen("tcp", "0.0.0.0:9092")
    if err != nil {
        fmt.Println("Failed to bind to port 9092")
        os.Exit(1)
    }
    defer l.Close()
   

    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting connection: ", err.Error())
            os.Exit(1)
        }
        defer conn.Close()
        
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    err := conn.SetReadDeadline(time.Now().Add(5*time.Second))
    if err != nil {
        fmt.Println("Error: ", err.Error())
        os.Exit(1)
    }
    run := true
    for run {
        fmt.Println("Accepting Input")
        buff := make([]byte, 1024)
        conn.Read(buff)


        // msg = 4
        // header req_api_key = 2, req_api_vers = 2, correlation_id = 4 
        //message_size := utils.bytes_to_int(buff[:4])
        req_header := request.RequestHeaderV2{}
        buff = req_header.Deserialize(buff[4:])
 
        supported_api_versions := []int{}

        if api_keys[req_header.ApiKey] == ApiVersions {
            supported_api_versions = []int{ 0,1,2,3,4 }
        } else if api_keys[req_header.ApiKey] == DescribeTopicPartitions {
            supported_api_versions = []int{ 0 }
        }
        
        is_api_version_supported := false
        for i := 0; i < len(supported_api_versions); i++ {
            if req_header.ApiVersion == supported_api_versions[i] {
                is_api_version_supported = true
            }
        }

        // capture error
        error_code := [2]byte{0,0}
        if !is_api_version_supported {
            error_code = [2]byte{0, 35}
        }
        
        resp := response.Response{}
        if api_keys[req_header.ApiKey] == ApiVersions {
            header := response.ResponseHeaderV0{}
            header.CorrelationId = req_header.CorrelationId

            //req_body := request.ApiVersionsRequestV4{}

            body := response.ApiVersionResponseV4{}
            body.ErrorCode = error_code
            body.ThrottleTimeMs = [4]byte{2,2,3,4}

            api_versions := make([]response.ApiVersion, 0) 
            api_versions = append(api_versions, response.ApiVersion{
                ApiKey: [2]byte{0,1},
                MinVersion: [2]byte{0,0},
                MaxVersion: [2]byte{0,11},
            })
            api_versions = append(api_versions, response.ApiVersion{
                ApiKey: [2]byte{0,18},
                MinVersion: [2]byte{0,0},
                MaxVersion: [2]byte{0,4},
            })
            api_versions = append(api_versions, response.ApiVersion{
                ApiKey: [2]byte{0,75},
                MinVersion: [2]byte{0,0},
                MaxVersion: [2]byte{0,0},
            })
            body.ApiVersionsArr = response.ApiVersionsArray{
                MessageSize: [1]byte{byte(len(api_versions)+1)},
                ApiVersions: api_versions,
            }

            resp.Header = header
            resp.Body = body
        } else if api_keys[req_header.ApiKey] == DescribeTopicPartitions { 
            header := response.ResponseHeaderV1{}
            body := response.DescribeTopicPartitionsResponseV0{}

            header.CorrelationId = req_header.CorrelationId

            req_body := request.DescribeTopicPartitionsRequestV0{} 
            buff = req_body.Deserialize(buff)
            body.ThrottleTimeMs = [4]byte{ 0,0,0,0 }
            req_num_topic := req_body.TopicsArr.Length
            body.TopicsArr.MessageSize = [1]byte(utils.Int_to_bytes(req_num_topic, 1))
            for i:=0; i < req_num_topic-1; i++ {
                req_topic := req_body.TopicsArr.Topics[i]
                topic := response.Topic{}
                topic.Name.Contents = req_topic.Name
                topic.Name.Length = [1]byte(utils.Int_to_bytes(req_topic.Length, 1))
                // need to check for topic existence
                error_code = [2]byte(utils.Int_to_bytes(3, 2))
                topic.ErrorCode = error_code


                body.TopicsArr.Topics = append(body.TopicsArr.Topics, topic)
            }

            
            body.NextCursor.TopicName.Length = [1]byte{255}
            resp.Header = header
            resp.Body = body
        }
 
        _, err := conn.Write(resp.Serialize())
        if err != nil {
            run = false
        }
    }
}
