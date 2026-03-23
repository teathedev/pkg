// Package cloudstorage contains cloud storage implementations
package cloudstorage

import "io"

type Credentials struct {
	AccountKey    string
	AccountSecret string
	ServiceURL    string
}

type StreamResponse struct {
	Size        int64
	ContentType string
	Stream      io.ReadCloser
}

type CloudStorage interface {
	CreateDownloadURL(
		creds Credentials,
		container string,
		objectKey string,
	) (string, error)
	CreateUploadURL(
		creds Credentials,
		container string,
		objectKey string,
	) (string, error)
	DownloadFile(
		creds Credentials,
		container string,
		objectKey string,
		path string,
	) error
	DownloadStream(
		creds Credentials,
		container string,
		objectKey string,
	) (StreamResponse, error)
}
