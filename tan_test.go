// Copyright 2015 The Tango Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tango

import (
	"bytes"
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestTan1(t *testing.T) {
	buff := bytes.NewBufferString("")
	recorder := httptest.NewRecorder()
	recorder.Body = buff

	l := NewLogger(os.Stdout)
	o := Classic(l)
	o.Get("/", func() string {
		return Version()
	})
	o.Logger().Debug("it's ok")

	req, err := http.NewRequest("GET", "http://localhost:8000/", nil)
	if err != nil {
		t.Error(err)
	}

	o.ServeHTTP(recorder, req)
	expect(t, recorder.Code, http.StatusOK)
	refute(t, len(buff.String()), 0)
	expect(t, buff.String(), Version())
}

func TestTan2(t *testing.T) {
	o := Classic()
	o.Get("/", func() string {
		return Version()
	})
	defer o.Shutdown(context.Background())
	go o.Run()

	time.Sleep(100 * time.Millisecond)

	resp, err := http.Get("http://localhost:8000/")
	if err != nil {
		t.Error(err)
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	expect(t, resp.StatusCode, http.StatusOK)
	expect(t, string(bs), Version())
}

func TestTan3(t *testing.T) {
	o := Classic()
	o.Get("/", func() string {
		return Version()
	})
	defer o.Shutdown(context.Background())
	go o.Run(":4040")

	time.Sleep(100 * time.Millisecond)

	resp, err := http.Get("http://localhost:4040/")
	if err != nil {
		t.Error(err)
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	expect(t, resp.StatusCode, http.StatusOK)
	expect(t, string(bs), Version())
}

func TestMinTLS(t *testing.T) {
	o := Classic()
	o.Get("/", func() string {
		return Version()
	})
	o.SetMinTLSVersion(tls.VersionTLS12)
	defer o.Shutdown(context.Background())
	go o.RunTLS("./public/cert.pem", "./public/key.pem", ":5050")

	time.Sleep(100 * time.Millisecond)

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	req, err := http.NewRequest("GET", "https://localhost:5050/", nil)
	if err != nil {
		t.Error(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	expect(t, resp.StatusCode, http.StatusOK)
	expect(t, string(bs), Version())
}

func TestMinTLSFail(t *testing.T) {
	o := Classic()
	o.Get("/", func() string {
		return Version()
	})
	o.SetMinTLSVersion(tls.VersionTLS12)
	defer o.Shutdown(context.Background())
	go o.RunTLS("./public/cert.pem", "./public/key.pem", ":5050")

	time.Sleep(100 * time.Millisecond)

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion:         tls.VersionTLS11,
				MaxVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
		},
	}
	req, err := http.NewRequest("GET", "https://localhost:5050/", nil)
	if err != nil {
		t.Error(err)
	}
	_, err = client.Do(req)
	if err == nil {
		t.Error(err)
	}
}

/* Test Helpers */
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func refute(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
