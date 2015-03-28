package pmtu


import (
	"fmt"
	"runtime"
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
// Path MTU tests either result in a pmtu value or an error string in this struct
//
type pmtu struct {
	pmtu uint
	err string
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
	resultChan := DetectPmtuAsync("127.0.0.1")
	result := <- resultChan
	fmt.Println(result.pmtu)
	fmt.Println(result.err)
}




//
// Detect path MTU and return result synchronously
// ===============================================
//
// USAGE: result := pmtu.DetectPmtu(hostname, expectedPmtu)
//
// -hostname : string (required) is a valid hostname or a valid ipv4/ipv6 address.
// -expectedPmtu : uint (optional) defines the expected PMTU, if known.  If correct, the testing will be sped up considerably.  Default = 1500.
// 
func DetectPmtu(hostname string, expectedPmtu ...uint) pmtu {
	var result pmtu
	return result
}




//
// Detect path MTU and return result asynchronously
//
// USAGE:
//	resultChan := pmtu.DetectPmtuAsync(hostname, expectedPmtu)
//	<other code>
//	result := <- resultChan
//
// -hostname : string (required) is a valid hostname or a valid ipv4/ipv6 address.
// -expectedPmtu : uint (optional) defines the expected PMTU, if known.  If correct, the testing will be sped up considerably.  Default = 1500.
//
func DetectPmtuAsync(hostname string) chan pmtu {
	runtime.GOMAXPROCS(runtime.NumCPU())
	resultChan := make(chan pmtu)
	go detectPmtuAsync(resultChan, hostname)
	return resultChan
}




//
// -=PRIVATE=- Asynchronously return PMTU test results
//
func detectPmtuAsync(resultChan chan pmtu, hostname string) {
	resultChan <- DetectPmtu(hostname)
}

