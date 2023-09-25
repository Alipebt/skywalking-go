package main

import "github.com/apache/skywalking-go/toolkit/trace"

func testTag() {
	trace.CreateLocalSpan("testSetTag")
	trace.SetTag("SetTag", "success")
	trace.StopSpan()
}

func testLog() {
	trace.CreateLocalSpan("testSetLog")
	trace.SetLog("SetLog", "success")
	trace.StopSpan()
}

func testSetOperationName() {
	trace.CreateLocalSpan("testSetOperationName_failed")
	trace.SetOperationName("testSetOperationName_success")
	trace.StopSpan()
}

func testGetTraceID() {
	trace.CreateLocalSpan("testGetTraceID")
	trace.SetTag("traceID", trace.GetTraceID())
	trace.StopSpan()
}

func testGetSpanID() {
	trace.CreateLocalSpan("testGetSpanID")
	trace.SetTag("spanID", string(trace.GetSpanID()))
	trace.StopSpan()
}

func testGetSegmentID() {
	trace.CreateLocalSpan("testGetSegmentID")
	trace.SetTag("segmentID", trace.GetSegmentID())
	trace.StopSpan()
}
