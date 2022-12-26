# This cmdlet sends Internet Control Message Protocol (ICMP) "echo request" packets ("pings") to one or more remote computers and returns the echo response replies. Here is an example of how you can use Test-Connection to ping a host in PowerShell:
# Replace "google.com" with the host you want to ping
$target = "google.com"

# Send a ping to the target and store the response
$response = Test-Connection -ComputerName $target -Count 1 -ErrorAction SilentlyContinue

# Check if the ping was successful
if ($response) {
    # The ping was successful
    Write-Host "Host is reachable"
} else {
    # The ping failed
    Write-Host "Host is not reachable"
}

