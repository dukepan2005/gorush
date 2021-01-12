package gorush

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// DispatchFeedback sends a feedback to the configured gateway.
func DispatchFeedback(log LogPushEntry, url string, timeout int64) error {

	if url == "" {
		return errors.New("The url can't be empty")
	}

	payload, err := json.Marshal(log)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// var client = &http.Client{
	// 	Timeout: time.Duration(timeout) * time.Second,
	// 	Transport: FeedbackTransport,
	// }
	_ = timeout
	resp, err := FeedbackClient.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return err
	}

	return nil
}
