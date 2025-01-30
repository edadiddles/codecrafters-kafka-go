package utils

import (
    "math"
)

func Bytes_to_int(b []byte) int {
    val := 0
    for i:=0; i < len(b); i++ {
        val += int(b[i]) * int(math.Pow(2.0, float64(len(b)-1-i)))
    }

    return val
}

func Int_to_bytes(v int, num_bytes int) []byte {
    b := make([]byte, num_bytes)

    for i:=num_bytes-1; i >= 0; i-- {
        b[i] = byte(v%256)
        v /= 256
    }

    return b
}
