package pmtu


import (
	"fmt"
	"net"
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
	

	return result
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

