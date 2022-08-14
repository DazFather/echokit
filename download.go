package echokit

import (
	"os"
	"path"

	"github.com/NicoNex/echotron/v3"
)

// DownloadDocument downloads a given document and store it in given directory
func DownloadDocument(api echotron.API, document echotron.Document, basePath string) (content []byte, err error) {
	return DownloadFile(api, document.FileID, path.Join(basePath, document.FileName))
}

// DownloadFile downloads a file with given ID at the given path
func DownloadFile(api echotron.API, fileID string, path string) (content []byte, err error) {
	content, err = FetchFile(api, fileID)
	if err == nil {
		err = os.WriteFile(path, content, os.ModePerm)
	}

	return
}

// FetchFile gets the content of a file with given ID
func FetchFile(api echotron.API, fileID string) (content []byte, err error) {
	var res echotron.APIResponseFile

	res, err = api.GetFile(fileID)
	if err != nil {
		return
	}

	return api.DownloadFile(res.Result.FilePath)
}
