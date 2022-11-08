package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/dummy_adress", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/dummy_adress", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("form shows not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	form := New(url.Values{})

	if form.Has("a") {
		t.Error("form shows has field when fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "abc")
	form = New(postedData)

	if !form.Has("a") {
		t.Error("form shows not have field when it does")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("a", "ab")
	form := New(postedData)

	if form.MinLength("d", 1) {
		t.Error("form shows acceptable length for non-existing field")
	}

	if form.MinLength("a", 3) {
		t.Error("form shows acceptable length of the field when it is too short")
	}

	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedData.Add("b", "abc")
	postedData.Add("c", "abcd")
	form = New(postedData)

	if !form.MinLength("b", 3) {
		t.Error("form shows too short input value when it is ok")
	}

	isError = form.Errors.Get("b")
	if isError != "" {
		t.Error("should not have an error, but get one")
	}

	if !form.MinLength("c", 3) {
		t.Error("form shows too short input value when it is ok")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("a", "m@w.com")
	postedData.Add("b", "@awa.pl")
	form := New(postedData)

	form.IsEmail("a")
	if !form.Valid() {
		t.Error("form shows wrong email format when it is ok")
	}

	form.IsEmail("b")
	if form.Valid() {
		t.Error("form shows correct email format when it is wrong")
	}

	form.Errors = errors{}

	form.IsEmail("c")
	if form.Valid() {
		t.Error("form shows correct email for non-existing field")
	}
}
