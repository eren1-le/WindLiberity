package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type UploadForm interface {
	// FielfName returns field name for upload
	FieldName() string

	// FileName returns filename for upload
	FileName() string

	// ExtraFields returna extra field for upload
	ExtraFields() map[string]string

	// Buffer return the buffer of media
	Buffer() ([]byte, error)
}

type httpUpload struct {
	fieldname	string
	filename	string
	resourceURL	string
	extraFields map[string]string
}

func (u *httpUpload) FileName() string {
	return u.filename
}

func (u *httpUpload) FieldName() string {
	return u.fieldname
}

func (u *httpUpload) ExtraFields() map[string]string {
	return u.extraFields
}

func (u *httpUpload) Buffer() ([]byte, error) {
	if len(u.resourceURL) != 0 {
		resp, err := http.Get(u.resourceURL)

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error http code: %d", resp.StatusCode)
		}

		return io.ReadAll(resp.Body)
	}

	path, err := filepath.Abs(u.fieldname)

	if err != nil{
		return nil, err
	}

	return os.ReadFile(path)
}

// UploadOption configures how we set up the http upload form.
type UploadOption func(u *httpUpload)


// WithREsouceURl specifies http upload by resource url
func WithResourceURL(url string) UploadOption {
	return func(u *httpUpload) {
		u.resourceURL = url
	}
}

// WithExtraField specifies the extra field to http upload from.
func WithExtraField(key, value string) UploadOption {
	return func(u *httpUpload) {
		u.extraFields[key] = value
	}
}

// NewUploadForm returns new upload form
func NewUploadForm(fieldname, filename string, options ...UploadOption) UploadForm {
	form := &httpUpload{
		filename: filename,
		fieldname: fieldname,
	}

	if len(options) != 0 {
		form.extraFields = make(map[string]string)
		for _, f := range options {
			f(form)
		}
	}
	return form
}