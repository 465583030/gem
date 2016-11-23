// Copyright 2016 The Gem Authors. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package gem

import (
	"bufio"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
)

type project struct {
	Name string `json:"name" xml:"name"`
}

func TestContext(t *testing.T) {

	router := NewRouter()
	respHtml := "OK"
	router.GET("/html", func(c *Context) {
		c.HTML(fasthttp.StatusOK, respHtml)
	})

	p := project{Name: "GEM"}
	respJson, err := json.Marshal(&p)
	if err != nil {
		t.Fatalf("json.Marshal error %s", err)
	}

	router.GET("/json", func(c *Context) {
		c.JSON(fasthttp.StatusOK, p)
	})

	jsonpCallback := []byte("callback")
	var respJsonp []byte
	respJsonp = append(respJsonp, jsonpCallback...)
	respJsonp = append(respJsonp, "("...)
	respJsonp = append(respJsonp, respJson...)
	respJsonp = append(respJsonp, ")"...)

	router.GET("/jsonp", func(c *Context) {
		c.JSONP(fasthttp.StatusOK, p, jsonpCallback)
	})

	router.GET("/xml", func(c *Context) {
		c.XML(fasthttp.StatusOK, p, xml.Header)
	})

	s := New("", router.Handler)

	// HTML
	rw1 := &readWriter{}
	rw1.r.WriteString("GET /html HTTP/1.1\r\n\r\n")

	ch1 := make(chan error)
	go func() {
		ch1 <- s.ServeConn(rw1)
	}()

	select {
	case err := <-ch1:
		if err != nil {
			t.Fatalf("return error %s", err)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("timeout")
	}

	br1 := bufio.NewReader(&rw1.w)
	var resp1 fasthttp.Response
	if err := resp1.Read(br1); err != nil {
		t.Fatalf("Unexpected error when reading response: %s", err)
	}
	if !(resp1.Header.StatusCode() == fasthttp.StatusOK) {
		t.Errorf("Regular routing failed with router chaining.")
		t.FailNow()
	}
	if !bytes.Equal(resp1.Header.PeekBytes(contentType), ContentTypeHTML) {
		t.Errorf("unexpected Content-Type got %q want %q", string(resp1.Header.PeekBytes(contentType)), string(ContentTypeHTML))
		t.FailNow()
	}
	if !bytes.Equal(resp1.Body(), []byte(respHtml)) {
		t.Errorf("unexpected response got %q want %q", string(resp1.Body()), respHtml)
		t.FailNow()
	}

	// JSON
	rw2 := &readWriter{}
	rw2.r.WriteString("GET /json HTTP/1.1\r\n\r\n")

	ch2 := make(chan error)
	go func() {
		ch2 <- s.ServeConn(rw2)
	}()

	select {
	case err := <-ch2:
		if err != nil {
			t.Fatalf("return error %s", err)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("timeout")
	}

	br2 := bufio.NewReader(&rw2.w)
	var resp2 fasthttp.Response
	if err := resp2.Read(br2); err != nil {
		t.Fatalf("Unexpected error when reading response: %s", err)
	}
	if !(resp2.Header.StatusCode() == fasthttp.StatusOK) {
		t.Errorf("Regular routing failed with router chaining.")
		t.FailNow()
	}
	if !bytes.Equal(resp2.Header.PeekBytes(contentType), ContentTypeJSON) {
		t.Errorf("unexpected Content-Type got %q want %q", string(resp2.Header.PeekBytes(contentType)), string(ContentTypeJSON))
		t.FailNow()
	}
	if !bytes.Equal(resp2.Body(), []byte(respJson)) {
		t.Errorf("unexpected response got %q want %q", string(resp2.Body()), respJson)
		t.FailNow()
	}

	// JSONP
	rw3 := &readWriter{}
	rw3.r.WriteString("GET /jsonp HTTP/1.1\r\n\r\n")

	ch3 := make(chan error)
	go func() {
		ch3 <- s.ServeConn(rw3)
	}()

	select {
	case err := <-ch3:
		if err != nil {
			t.Fatalf("return error %s", err)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("timeout")
	}

	br3 := bufio.NewReader(&rw3.w)
	var resp3 fasthttp.Response
	if err := resp3.Read(br3); err != nil {
		t.Fatalf("Unexpected error when reading response: %s", err)
	}
	if !(resp3.Header.StatusCode() == fasthttp.StatusOK) {
		t.Errorf("Regular routing failed with router chaining.")
		t.FailNow()
	}
	if !bytes.Equal(resp3.Header.PeekBytes(contentType), ContentTypeJSONP) {
		t.Errorf("unexpected Content-Type got %q want %q", string(resp3.Header.PeekBytes(contentType)), string(ContentTypeJSONP))
		t.FailNow()
	}
	if !bytes.Equal(resp3.Body(), []byte(respJsonp)) {
		t.Errorf("unexpected response got %q want %q", string(resp3.Body()), respJsonp)
		t.FailNow()
	}

	// XML
	rw4 := &readWriter{}
	rw4.r.WriteString("GET /xml HTTP/1.1\r\n\r\n")

	ch4 := make(chan error)
	go func() {
		ch4 <- s.ServeConn(rw4)
	}()

	select {
	case err := <-ch4:
		if err != nil {
			t.Fatalf("return error %s", err)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("timeout")
	}

	br4 := bufio.NewReader(&rw4.w)
	var resp4 fasthttp.Response
	if err := resp4.Read(br4); err != nil {
		t.Fatalf("Unexpected error when reading response: %s", err)
	}
	if !(resp4.Header.StatusCode() == fasthttp.StatusOK) {
		t.Errorf("Regular routing failed with router chaining.")
		t.FailNow()
	}
	if !bytes.Equal(resp4.Header.PeekBytes(contentType), ContentTypeXML) {
		t.Errorf("unexpected Content-Type got %q want %q", string(resp4.Header.PeekBytes(contentType)), string(ContentTypeXML))
		t.FailNow()
	}
	p2 := project{}
	if err := xml.Unmarshal(resp4.Body(), &p2); err != nil {
		t.Fatalf("xml.Unmarshal error %s", err)
		t.FailNow()
	}
	if p2.Name != p.Name {
		t.Errorf("unexpected project's name got %q want %q", p2.Name, p.Name)
	}
}

func TestContext2(t *testing.T) {
	router := NewRouter()

	router.GET("/", func(c *Context) {
		if !c.IsAjax() {
			t.Errorf("Expected c.IsAjax() = %t, got %t", true, c.IsAjax())
		}

		if !strings.EqualFold(string(c.Method()), c.MethodString()) {
			t.Errorf("Expected method %q, got %q", c.Method(), c.MethodString())
		}

		if !strings.EqualFold(string(c.Path()), c.PathString()) {
			t.Errorf("Expected path %q, got %q", c.Path(), c.PathString())
		}

		if !strings.EqualFold(string(c.Host()), c.HostString()) {
			t.Errorf("Expected host %q, got %q", c.Host(), c.HostString())
		}
	})

	s := New("", router.Handler)

	// HTML
	rw1 := &readWriter{}
	rw1.r.WriteString("GET http://127.0.0.1/ HTTP/1.1\r\nX-Requested-With: XMLHttpRequest\r\n\r\n")

	ch1 := make(chan error)
	go func() {
		ch1 <- s.ServeConn(rw1)
	}()

	select {
	case err := <-ch1:
		if err != nil {
			t.Fatalf("return error %s", err)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("timeout")
	}

	br1 := bufio.NewReader(&rw1.w)
	var resp1 fasthttp.Response
	if err := resp1.Read(br1); err != nil {
		t.Fatalf("Unexpected error when reading response: %s", err)
	}
	if !(resp1.Header.StatusCode() == fasthttp.StatusOK) {
		t.Errorf("Regular routing failed with router chaining.")
		t.FailNow()
	}
}
