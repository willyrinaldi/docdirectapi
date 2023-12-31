// Code generated by go-swagger; DO NOT EDIT.

package balance_sheet_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// GetDirectV1SahamLkBsURL generates an URL for the get direct v1 saham lk bs operation
type GetDirectV1SahamLkBsURL struct {
	Granularity string
	Periode     string
	Q           string
	SecCode     string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetDirectV1SahamLkBsURL) WithBasePath(bp string) *GetDirectV1SahamLkBsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetDirectV1SahamLkBsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetDirectV1SahamLkBsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/direct/v1/saham/lk/bs/"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	granularityQ := o.Granularity
	if granularityQ != "" {
		qs.Set("granularity", granularityQ)
	}

	periodeQ := o.Periode
	if periodeQ != "" {
		qs.Set("periode", periodeQ)
	}

	qQ := o.Q
	if qQ != "" {
		qs.Set("q", qQ)
	}

	secCodeQ := o.SecCode
	if secCodeQ != "" {
		qs.Set("secCode", secCodeQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetDirectV1SahamLkBsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetDirectV1SahamLkBsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetDirectV1SahamLkBsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetDirectV1SahamLkBsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetDirectV1SahamLkBsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetDirectV1SahamLkBsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
