// This function handles errors returned by the datastore.Get function and viewTemplate’s Execute method.
// In both cases, it presents a simple error message to the user with the HTTP status code 500 (“Internal Server Error”).
// This looks like a manageable amount of code, but add some more HTTP handlers
// and you quickly end up with many copies of identical error handling code.

package main

import (
	"appengine"
	"appengine/datastore"
	"net/http"
)

func init() {
	http.HandleFunc("/view", viewRecord)
}

func viewRecord(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	// This is the error handling code
	// get the record from the datastore
	// If there is an error, it will return the error message
	if err := datastore.Get(c, key, record); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// This is the error handling code
	// give the record to the view template for rendering
	// If there is an error, it will return the error message
	if err := viewTemplate.Execute(w, record); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

// To reduce the repetition we can define our own HTTP appHandler type that includes an error return value
type appHandler func(http.ResponseWriter, *http.Request) error

// The appHandler type is a function that takes an http.ResponseWriter, an http.Request, and returns an error.

// Then change the viewRecord function to return an error instead of writing the error message to the http.ResponseWriter.
func viewRecord(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	if err := datastore.Get(c, key, record); err != nil {
		return err
	}
	return viewTemplate.Execute(w, record)
}

// but http package don't undertand the func that return errors
// So we need to create a wrapper function that converts our appHandler type into an http.HandlerFunc
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

// Now we can use the appHandler type to wrap our viewRecord function
func init() {
	http.Handle("/view", appHandler(viewRecord))
}

// Make more understandable with errors handling
type appErrros struct {
	Error   error
	Message string
	Code    int
}

// From line 34
type appHandler func(http.ResponseWriter, *http.Request) *appErrors

// Update the viewRecord function to return an *appErrors instead of an errors
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil {
		c := appengine.NewContext(r)
		c.Errorf("%v", e.Error)
		http.Error(w, e.Message, e.Code)
	}
}

// Update the viewRecord function to return an *appErrors instead of an errors
func viewRecord(w http.ResponseWriter, r *http.Request) *appErrors {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	if err := datastore.Get(c, key, record); err != nil {
		return &appErrors{err, "Record not found", 404}
	}
	if err := viewTemplate.Execute(w, record); err != nil {
		return &appErrors{err, "Can't display record", 500}
	}
	return nil
}
