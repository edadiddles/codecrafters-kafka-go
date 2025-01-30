package main

type ApiKey int

const (
    ApiVersions ApiKey = iota
    DescribeTopicPartitions
)

var api_keys = map[int]ApiKey {
    18: ApiVersions,
    75: DescribeTopicPartitions,
}
