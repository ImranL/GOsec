package pmtu


import (
	"fmt"
	"net"
	"os"
	"time"
)


//
//
// Path MTU detector
// =================
//
// -Call pmtu.PmtuTestHarness() to run a self-test on the library
//
// -To test path MTU synchronously, call pmtu.DetectPmtu; to detect asynchronously, call pmtu.DetectPmtuAsync
//
//




//
// Calls to this library are made with this struct.  Hostname is required, all other variables are optional.
//
type PmtuTest struct {
	Hostname string
	ExpectedPmtu uint
	IcmpTimeoutMS uint
}




//
// Path MTU tests either result in a pmtu value or an error string for IPv4 and IPv6 in this struct
//
type PmtuResult struct {
	Pmtu4 uint
	Pmtu6 uint
	Err4 string
	Err6 string
}




//
// Library diagnostics
// ===================
//
// Diagnostic function to test PMTU library functions, prints status to stdout
//
// USAGE: call pmtu.PmtuTestHarness()
//
func PmtuTestHarness() {
	var test PmtuTest
	test.Hostname = "google.com"
	resultChan := DetectPmtuAsync(test)
	result := <- resultChan
fmt.Println(result.Pmtu4, result.Pmtu6, result.Err4, result.Err6)
}




//
// Detect path MTU and return result synchronously
// ===============================================
//
// USAGE: result := pmtu.DetectPmtu(inputTest)
//
// -inputTest : PmtuTest (see struct definition above)
// -result : PmtuResult (see struct definition above)
// 
func DetectPmtu(test PmtuTest) PmtuResult {

	// Set up results container
	var result PmtuResult

	// Look up hostname, record v4 and/or v6 records.  Fail out if the lookup isn't successful.
	ips, err := net.LookupIP(test.Hostname)
	if err != nil {
		result.Err4 = "Hostname lookup failure"
		result.Err6 = "Hostname lookup failure"
		return result
	}
	var ip4, ip6 net.IP
	for i := 0; i < len(ips); i++ {
		if ips[i].To4() != nil {
			ip4 = ips[i]
		} else {
			ip6 = ips[i]
		}
	}
fmt.Println(ip4, ip6)

	// Ping the host, check that we can get to it with minimally-sized ping.  Simultaneously ping with the expected PMTU if set.
	conn, err := net.Dial("ip4:icmp", ip4.String())
	if err != nil {
		result.Err4 = "blah"
		result.Err6 = "blah"
		return result
	}
	defer conn.Close()
	pid := os.Getpid()
	var id1 = byte(pid & 0xff00 >> 8)
	var id2 = byte(pid & 0xff)
	var msg [64]byte
	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0
	msg[4] = id1
	msg[5] = id2
	msg[6] = 0
	msg[7] = 1
	check := CheckSum(msg[0:8])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 0xff)
	conn.SetDeadline(time.Now().Add(time.Second))
	if _, err = conn.Write(msg[0:8]); err != nil {
		result.Err4 = "blah2"
		result.Err6 = "blah2"
		return result
	}
	if _, err = conn.Read(msg[0:]); err != nil {
		result.Err4 = "blah3"
		result.Err6 = "blah3"
		return result
	}
	fmt.Println(msg)

	return result
}




func CheckSum(buf []byte) uint16 {
	var sum int32
	n := len(buf)
	if len(buf) % 2 != 0 {
		n--
	}
	for i := 0; i < n; i += 2 {
		sum += int32(buf[i]) << 8 + int32(buf[i+1])
	}
	if len(buf) % 2 != 0 {
		sum += int32(buf[n])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var ans uint16 = uint16(^sum)
	return ans
}




//
// Detect path MTU and return result asynchronously
//
// USAGE:
//	resultChan := pmtu.DetectPmtuAsync(inputTest)
//	<other code>
//	result := <- resultChan
//
// -inputTest : PmtuTest (see struct definition above)
// -result : PmtuResult (see struct definition above)
//
func DetectPmtuAsync(test PmtuTest) chan PmtuResult {
	resultChan := make(chan PmtuResult)
	go detectPmtuAsync(resultChan, test)
	return resultChan
}




//
// -=PRIVATE=- Asynchronously return PMTU test results
//
func detectPmtuAsync(resultChan chan PmtuResult, test PmtuTest) {
	resultChan <- DetectPmtu(test)
}

