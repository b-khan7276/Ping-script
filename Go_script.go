//  Here is a simple ping script in Go that sends an ICMP "echo request" message to a specified host and waits for an "echo reply." This can be used to check the reachability of a host and measure the round-trip time for a message to be sent to and received back from the host.

// To use this script, you will need to import the "net" and "fmt" packages.

package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    // Replace "google.com" with the host you want to ping
    target := "google.com"

    // Resolve the target's IP address
    ip, err := net.ResolveIPAddr("ip", target)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Create a connection to the target
    conn, err := net.DialIP("ip4:icmp", nil, ip)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()

    // Set a deadline for the connection
    err = conn.SetDeadline(time.Now().Add(time.Second * 5))
    if err != nil {
        fmt.Println(err)
        return
    }

    // Send the ICMP echo request
    _, err = conn.Write([]byte{0x08, 0x00, 0x7d, 0x4b, 0x00, 0x00, 0x00, 0x00, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69})
    if err != nil {
        fmt.Println(err)
        return
    }

    // Wait for the echo reply
    reply := make([]byte, 1024)
    n, _, err := conn.ReadFrom(reply)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print the response
    fmt.Printf("% x\n", reply[:n])
}
// This script sends an ICMP echo request message to the specified host and waits for an echo reply. If a reply is received within the specified deadline, the response is printed in hexadecimal format. If no reply is received, or if there is an error, an appropriate error message is printed.
