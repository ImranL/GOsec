package pmtu


import (
	"fmt"
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
}




//
// Path MTU tests either result in a pmtu value or an error string in this struct
//
type PmtuResult struct {
	Pmtu uint
	Err string
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
	test.Hostname = "127.0.0.1"
	resultChan := DetectPmtuAsync(test)
	result := <- resultChan
	fmt.Println(result.Pmtu)
	fmt.Println(result.Err)
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
	var result PmtuResult
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

