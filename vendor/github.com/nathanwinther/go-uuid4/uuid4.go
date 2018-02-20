package uuid4

// Create Version 4 UUIDs as specified in RFC 4122

import (
    "fmt"
    "crypto/rand"
)

func New() (string, error) {
    b := make([]byte, 16)

    _, err := rand.Read(b[:])
    if err != nil {
        return "", err
    }

    // Set the two most significant bits (bits 6 and 7) of the
    // clock_seq_hi_and_reserved to zero and one, respectively.
    b[8] = (b[8] | 0x40) & 0x7F

    // Set the four most significant bits (bits 12 through 15) of the
    // time_hi_and_version field to the 4-bit version number.
    b[6] = (b[6] & 0xF) | (4 << 4)

    // Return unparsed version of the generated UUID sequence.
    return fmt.Sprintf("%x-%x-%x-%x-%x",
        b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

