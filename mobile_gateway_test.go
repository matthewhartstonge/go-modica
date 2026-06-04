package modica

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMobileGatewayService_CreateMessage_ErrMobileGatewaySendFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Authorization", expectedAuthHeader)
		testHeader(t, r, "Accept", mediaTypeV1)

		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `{"error-desc":"Could not queue message due to an unknown error","error":"send_failed"}`)
	})

	payload := &Message{
		Destination: "LOL This Isn't a destination number",
		Content:     "Hi, this is a test message to ensure you are texting correctly",
		Source:      "TEST",
		Scheduled:   "2017-05-05T10:00:00+12:00",
		Reference:   "alt-reference",
		Class:       "mt_message",
		Mask:        "",
		SMSClass:    2,
	}
	_, got := client.MobileGateway.CreateMessage(payload)
	if got == nil {
		t.Errorf("MobileGateway.CreateMessage should have returned ErrMobileGatewaySendFailed")
	}

	want := ErrMobileGatewaySendFailed
	if got != want {
		t.Errorf("MobileGateway.CreateMessage returned %+v, want %+v", got, want)
	}
}

func TestMobileGatewayService_CreateMessage_ErrMobileGatewayInvalidJSON(t *testing.T) {
	// This is more a theoretical test, as this case should
	// never happen due to Go's amazing JSON Marshaling!
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Authorization", expectedAuthHeader)
		testHeader(t, r, "Accept", mediaTypeV1)

		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `{"error-desc":"Invalid JSON data in the request body","error":"invalid_json"}`)
	})

	payload := &Message{}
	_, got := client.MobileGateway.CreateMessage(payload)
	if got == nil {
		t.Errorf("MobileGateway.CreateMessage should have returned ErrMobileGatewayInvalidJSON")
	}

	want := ErrMobileGatewayInvalidJSON
	if got != want {
		t.Errorf("MobileGateway.CreateMessage returned %+v, want %+v", got, want)
	}
}

func TestMobileGatewayService_CreateMessage_ErrMobileGatewayMissingAttribute(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Authorization", expectedAuthHeader)
		testHeader(t, r, "Accept", mediaTypeV1)

		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `{"error-desc":"Missing required destination attribute","error":"missing_attrib"}`)
	})

	payload := &Message{
		Destination: "",
		Content:     "Hi, this is a test message to ensure you are texting correctly",
		Source:      "TEST",
		Scheduled:   "2017-05-05T10:00:00+12:00",
		Reference:   "alt-reference",
		Class:       "mt_message",
		Mask:        "",
		SMSClass:    2,
	}
	_, got := client.MobileGateway.CreateMessage(payload)
	if got == nil {
		t.Errorf("MobileGateway.CreateMessage should have returned ErrMobileGatewayMissingAttribute")
	}

	want := ErrMobileGatewayMissingAttribute
	if got != want {
		t.Errorf("MobileGateway.CreateMessage returned %+v, want %+v", got, want)
	}
}

func TestMobileGatewayService_CreateMessage_ErrMobileGatewayInvalidAttribute(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Authorization", expectedAuthHeader)
		testHeader(t, r, "Accept", mediaTypeV1)

		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `{"error-desc":"Invalid attribute value","error":"invalid_attrib"}`)
	})

	payload := &Message{
		Destination: "LOL This Isn't a destination number",
		Content:     "Hi, this is a test message to ensure you are texting correctly",
		Source:      "TEST",
		Scheduled:   "2017-05-05T10:00:00+12:00",
		Reference:   "alt-reference",
		Class:       "mt_message",
		Mask:        "",
		SMSClass:    2,
	}
	_, got := client.MobileGateway.CreateMessage(payload)
	if got == nil {
		t.Errorf("MobileGateway.CreateMessage should have returned ErrMobileGatewayInvalidAttribute")
	}

	want := ErrMobileGatewayInvalidAttribute
	if got != want {
		t.Errorf("MobileGateway.CreateMessage returned %+v, want %+v", got, want)
	}
}

func TestMobileGatewayService_CreateMessage_ErrMobileGatewayInvalidTimestampFormat(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Authorization", expectedAuthHeader)
		testHeader(t, r, "Accept", mediaTypeV1)

		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `{"error-desc":"Invalid scheduled timestamp (must be RFC3339)","error":"400"}`)
	})

	payload := &Message{
		Destination: "+61234567890",
		Content:     "Hi, this is a test message to ensure you are texting correctly",
		Source:      "TEST",
		Scheduled:   "not RFC3339",
		Reference:   "alt-reference",
		Class:       "mt_message",
		Mask:        "",
		SMSClass:    2,
	}
	_, got := client.MobileGateway.CreateMessage(payload)
	if got == nil {
		t.Errorf("MobileGateway.CreateMessage should have returned ErrMobileGatewayInvalidTimestampFormat")
	}

	want := ErrMobileGatewayInvalidTimestampFormat
	if got != want {
		t.Errorf("MobileGateway.CreateMessage returned %+v, want %+v", got, want)
	}
}

func TestMobileGatewayService_CreateMessage_ErrMobileGatewayInvalidTimestamp(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Authorization", expectedAuthHeader)
		testHeader(t, r, "Accept", mediaTypeV1)

		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `{"error-desc":"Invalid scheduled timestamp (must not be in the past)","error":"422"}`)
	})

	payload := &Message{
		Destination: "+61234567890",
		Content:     "Hi, this is a test message to ensure you are texting correctly",
		Source:      "TEST",
		Scheduled:   "2017-05-05T10:00:00+12:00",
		Reference:   "alt-reference",
		Class:       "mt_message",
		Mask:        "",
		SMSClass:    2,
	}
	_, got := client.MobileGateway.CreateMessage(payload)
	if got == nil {
		t.Errorf("MobileGateway.CreateMessage should have returned ErrMobileGatewayInvalidTimestamp")
	}

	want := ErrMobileGatewayInvalidTimestamp
	if got != want {
		t.Errorf("MobileGateway.CreateMessage returned %+v, want %+v", got, want)
	}
}

func TestMobileGatewayService_CreateMessage_ErrMobileGatewayMessageIDNotFound_EmptySlice(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Authorization", expectedAuthHeader)
		testHeader(t, r, "Accept", mediaTypeV1)

		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `[]`)
	})

	payload := &Message{
		Destination: "+61234567890",
		Content:     "Hi, this is a test message to ensure you are texting correctly",
		Source:      "TEST",
		Scheduled:   "2017-05-05T10:00:00+12:00",
		Reference:   "alt-reference",
		Class:       "mt_message",
		Mask:        "",
		SMSClass:    2,
	}
	_, got := client.MobileGateway.CreateMessage(payload)
	if got == nil {
		t.Errorf("MobileGateway.CreateMessage should have returned ErrMobileGatewayMessageIDNotFound")
	}

	want := ErrMobileGatewayMessageIDNotFound
	if got != want {
		t.Errorf("MobileGateway.CreateMessage returned %+v, want %+v", got, want)
	}
}

func TestMobileGatewayService_CreateMessage_ErrMobileGatewayMessageIDNotFound_EmptyString(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Authorization", expectedAuthHeader)
		testHeader(t, r, "Accept", mediaTypeV1)

		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, ``)
	})

	payload := &Message{
		Destination: "+61234567890",
		Content:     "Hi, this is a test message to ensure you are texting correctly",
		Source:      "TEST",
		Scheduled:   "2017-05-05T10:00:00+12:00",
		Reference:   "alt-reference",
		Class:       "mt_message",
		Mask:        "",
		SMSClass:    2,
	}
	_, got := client.MobileGateway.CreateMessage(payload)
	if got == nil {
		t.Errorf("MobileGateway.CreateMessage should have returned ErrMobileGatewayMessageIDNotFound")
	}

	want := ErrMobileGatewayMessageIDNotFound
	if got != want {
		t.Errorf("MobileGateway.CreateMessage returned %+v, want %+v", got, want)
	}
}

func TestMobileGatewayService_CreateMessage(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Authorization", expectedAuthHeader)
		testHeader(t, r, "Accept", mediaTypeV1)
		testBody(t, r, `{"destination":"+642123456789","content":"Hi, this is a test message to ensure you are texting correctly","source":"TEST","scheduled":"2017-05-05T10:00:00+12:00","reference":"alt-reference","class":"mt_message","sms_class":2}`+"\n")

		// Modica's REST api returns a non keyed raw array with a single int,
		// representing the message ID if successful.
		_, _ = fmt.Fprint(w, `[123]`)
	})

	payload := &Message{
		Destination: "+642123456789",
		Content:     "Hi, this is a test message to ensure you are texting correctly",
		Source:      "TEST",
		Scheduled:   "2017-05-05T10:00:00+12:00",
		Reference:   "alt-reference",
		Class:       "mt_message",
		Mask:        "",
		SMSClass:    2,
	}
	got, err := client.MobileGateway.CreateMessage(payload)
	if err != nil {
		t.Errorf("MobileGateway.CreateMessage returned error: %v", err)
	}

	want := 123
	if got != want {
		t.Errorf("MobileGateway.CreateMessage returned %+v, want %+v", got, want)
	}
}

func TestMobileGatewayService_GetMessage(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/messages/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", expectedAuthHeader)
		testHeader(t, r, "Accept", mediaTypeV1)

		_, _ = fmt.Fprint(w, `{"id":123,"destination":"+642123456789","content":"Hi, this is a test message to ensure you are texting correctly","source":"TEST","reference":"alt-reference","operator":"2degrees","reply_to":"123"}`+"\n")
	})

	got, err := client.MobileGateway.GetMessage(123)
	if err != nil {
		t.Errorf("MobileGateway.CreateMessage returned error: %v", err)
	}

	want := &Message{
		ID:          123,
		Destination: "+642123456789",
		Content:     "Hi, this is a test message to ensure you are texting correctly",
		Source:      "TEST",
		Reference:   "alt-reference",
		ReplyTo:     "123",
		Operator:    "2degrees",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MobileGateway.GetMessage returned %+v, want %+v", got, want)
	}
}

func TestMobileGatewayService_GetMessage_NotFound(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/messages/321", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", expectedAuthHeader)
		testHeader(t, r, "Accept", mediaTypeV1)

		w.WriteHeader(404)
	})

	_, err := client.MobileGateway.GetMessage(321)
	if err == nil {
		t.Error("MobileGateway.GetMessage didn't return an error on not found")
	}

	if err != ErrNotFound {
		t.Errorf("MobileGateway.GetMessage returned the wrong error: got: %+v, want %+v", err, ErrNotFound)
	}
}

func TestMobileGatewayService_CreateBroadcastMessage_ErrMobileGatewayBroadcastLimit(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/messages/broadcast", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Authorization", expectedAuthHeader)
		testHeader(t, r, "Accept", mediaTypeV1)
		testBody(t, r, `{"destination":["+61234567890","+60987654321"],"content":"Hi, this is a test message to ensure you are texting correctly","source":"TEST","scheduled":"2017-05-05T10:00:00+12:00","reference":"alt-reference","class":"mt_message","sms_class":2}`+"\n")

		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `{"error-desc":"Broadcast limit has been exceeded, please consult the error description for more detail.","error":"broadcast_limit"}`)
	})

	payload := &BroadcastMessage{
		Destinations: []string{"+61234567890", "+60987654321"},
		Message: Message{
			Content:   "Hi, this is a test message to ensure you are texting correctly",
			Source:    "TEST",
			Scheduled: "2017-05-05T10:00:00+12:00",
			Reference: "alt-reference",
			Class:     "mt_message",
			Mask:      "",
			SMSClass:  2,
		},
	}
	_, got := client.MobileGateway.CreateBroadcastMessage(payload)
	if got == nil {
		t.Errorf("MobileGateway.CreateBroadcastMessage should have returned ErrMobileGatewayBroadcastLimit")
	}

	want := ErrMobileGatewayBroadcastLimit
	if got != want {
		t.Errorf("MobileGateway.CreateBroadcastMessage returned %+v, want %+v", got, want)
	}
}

func TestMobileGatewayService_CreateBroadcastMessage(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/messages/broadcast", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Authorization", expectedAuthHeader)
		testHeader(t, r, "Accept", mediaTypeV1)
		testBody(t, r, `{"destination":["+61234567890","X","0123456789"],"content":"Hi, this is a test message to ensure you are texting correctly","source":"TEST","scheduled":"2017-05-05T10:00:00+12:00","reference":"alt-reference","class":"mt_message","sms_class":2}`+"\n")

		_, _ = fmt.Fprint(w, `[{"status":"success","message":null,"destination":"+61234567890","id":123},{"status":"failure","message":"Invalid destination (X)","destination":"X","id":null},{"status":"failure","message":"That's not a real phone number is it","destination":"0123456789","id":null}]`)
	})

	payload := &BroadcastMessage{
		Destinations: []string{"+61234567890", "X", "0123456789"},
		Message: Message{
			Content:   "Hi, this is a test message to ensure you are texting correctly",
			Source:    "TEST",
			Scheduled: "2017-05-05T10:00:00+12:00",
			Reference: "alt-reference",
			Class:     "mt_message",
			Mask:      "",
			SMSClass:  2,
		},
	}
	got, err := client.MobileGateway.CreateBroadcastMessage(payload)
	if err != nil {
		t.Errorf("MobileGateway.CreateBroadcastMessage returned error: %v", err)
	}

	want := []BroadcastResponse{
		{
			Status:      "success",
			Message:     "",
			Destination: "+61234567890",
			ID:          123,
		},
		{
			Status:      "failure",
			Message:     "Invalid destination (X)",
			Destination: "X",
			ID:          0,
		},
		{
			Status:      "failure",
			Message:     "That's not a real phone number is it",
			Destination: "0123456789",
			ID:          0,
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MobileGateway.CreateBroadcastMessage returned %+v, want %+v", got, want)
	}
}
