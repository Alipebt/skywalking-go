package main

import (
	"github.com/apache/skywalking-go/toolkit/trace"
	"net/http"
)

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

func testContext() {
	trace.CreateLocalSpan("testCaptureContext")
	captureSpanID := trace.GetSpanID()
	ctx := trace.CaptureContext()
	trace.StopSpan()

	trace.ContinueContext(ctx)
	continueSpanID := trace.GetSpanID()
	trace.CreateLocalSpan("testContinueContext")
	if captureSpanID == continueSpanID {
		trace.SetTag("testContinueContext", "success")
	}
	trace.StopSpan()
}

func testContextCarrierAndCorrelation() {
	request, _ := http.NewRequest("GET", "http://localhost/", http.NoBody)
	trace.CreateExitSpan("ExitSpan", request.Host, func(headerKey, headerValue string) error {
		request.Header.Add(headerKey, headerValue)
		return nil
	})
	trace.SetCorrelation("testCorrelation", "success")

	go func() {
		trace.CreateEntrySpan("EntrySpan", func(headerKey string) (string, error) {
			return request.Header.Get(headerKey), nil
		})
		trace.SetTag("testCorrelation", trace.GetCorrelation("testCorrelation"))
		trace.StopSpan()
	}()
	trace.StopSpan()
}
