// Cache archive related models and functions.
package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/log"
)

// Archive represents a cache archive.
type Archive struct {
	file *os.File
	tar  *tar.Writer
	gzip *gzip.Writer
}

// NewArchive creates a instance of Archive.
func NewArchive(pth string, compress bool) (*Archive, error) {
	log.Donef("Creating File")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	file, err := os.Create(pth)
	if err != nil {
		return nil, err
	}

	var tarWriter *tar.Writer
	var gzipWriter *gzip.Writer
	if compress {
		log.Donef("Compressing File")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		gzipWriter, err = gzip.NewWriterLevel(file, gzip.BestCompression)
		if err != nil {
			return nil, err
		}

		tarWriter = tar.NewWriter(gzipWriter)
	} else {
		log.Donef("Creating uncompressed TAR")
	log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		tarWriter = tar.NewWriter(file)
	}
	return &Archive{
		file: file,
		tar:  tarWriter,
		gzip: gzipWriter,
	}, nil
}

// Write writes the given files in the cache archive.
func (a *Archive) Write(pathToIndicator map[string]string) error {
	for pth := range pathToIndicator {
		if err := a.writeOne(pth); err != nil {
			return err
		}
	}

	return nil
}

func (a *Archive) writeOne(pth string) error {
	info, err := os.Lstat(pth)
	if err != nil {
		return fmt.Errorf("failed to lstat(%s), error: %s", pth, err)
	}

	var link string
	if info.Mode()&os.ModeSymlink != 0 {
		link, err = os.Readlink(pth)
		if err != nil {
			return fmt.Errorf("failed to read link(%s), error: %s", pth, err)
		}
	}

	header, err := tar.FileInfoHeader(info, link)
	if err != nil {
		return fmt.Errorf("failed to get tar file header(%s), error: %s", link, err)
	}

	header.Name = pth
	header.ModTime = info.ModTime()

	if err := a.tar.WriteHeader(header); err != nil {
		return fmt.Errorf("failed to write header(%v), error: %s", header, err)
	}

	// Calling Write on special types like TypeLink, TypeSymlink, TypeChar, TypeBlock, TypeDir, and TypeFifo returns (0, ErrWriteTooLong) regardless of what the Header.Size claims.
	if !info.Mode().IsRegular() {
		return nil
	}

	file, err := os.Open(pth)
	if err != nil {
		return fmt.Errorf("failed to open file(%s), error: %s", pth, err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Warnf("Failed to close file (%s): %s", pth, err)
		}
	}()

	// Write writes to the current file in the tar archive. Write returns the error ErrWriteTooLong if more than Header.Size bytes are written after WriteHeader.
	if _, err := io.CopyN(a.tar, file, info.Size()); err != nil && err != io.EOF {
		return fmt.Errorf("failed to copy, error: %s, file: %s, size: %d for header: %v", err, file.Name(), info.Size(), header)
	}

	return nil
}

// WriteHeader writes the cache descriptor file into the archive as a tar header.
func (a *Archive) WriteHeader(descriptor map[string]string, descriptorPth string) error {
	log.Donef("Marshalling Header JSON")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	b, err := json.MarshalIndent(descriptor, "", " ")
	if err != nil {
		return err
	}

	log.Donef("Actually writing data to the header")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	return a.writeData(b, descriptorPth)
}

// writeData writes the byte array into the archive.
func (a *Archive) writeData(data []byte, descriptorPth string) error {
	log.Donef("Creating Header")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	header := &tar.Header{
		Name:     descriptorPth,
		Size:     int64(len(data)),
		Typeflag: tar.TypeReg,
		Mode:     0600,
		ModTime:  time.Now(),
	}

	log.Donef("Writing Header")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	if err := a.tar.WriteHeader(header); err != nil {
		return err
	}

	log.Donef("Copying Data io.copy")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	if _, err := io.Copy(a.tar, bytes.NewReader(data)); err != nil && err != io.EOF {
		return err
	}
	return nil
}

// Close closes the archive.
func (a *Archive) Close() error {
	log.Donef("Tar Close")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	if err := a.tar.Close(); err != nil {
		return err
	}

	log.Donef("Gzip Close")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	if a.gzip != nil {
		if err := a.gzip.Close(); err != nil {
			return err
		}
	}

	log.Donef("File Close Return")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	return a.file.Close()
}

// uploadArchive uploads the archive file to a given destination.
// If the destination is a local file path (url has a file:// scheme) this function copies the cache archive file to the destination.
// Otherwise destination should point to the Bitrise cache API server, in this case the function has builtin retry logic with 3s sleep.
func uploadArchive(pth, url string, buildSlug string) error {
	log.Donef("Trim String Prefix")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	if strings.HasPrefix(url, "file://") {
		dst := strings.TrimPrefix(url, "file://")
		dir := filepath.Dir(dst)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		log.Donef("Return command.CopyFile")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")

		return command.CopyFile(pth, dst)
	}

	log.Donef("Checking if path exists for upload")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	fi, err := os.Stat(pth)
	if err != nil {
		return fmt.Errorf("failed to get file info (%s): %s", pth, err)
	}
	log.Donef("Getting File Size")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	sizeInBytes := fi.Size()
	log.Printf("Archive file size: %d bytes / %f MB", sizeInBytes, (float64(sizeInBytes) / 1024.0 / 1024.0))
	data := map[string]interface{}{
		"cache_archive_size": sizeInBytes,
		"build_slug":         buildSlug,
	}
	log.RInfof(stepID, "cache_archive_size", data, "Size of cache archive: %d Bytes", sizeInBytes)

	log.Donef("Getting Cache Upload URL")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	uploadURL, err := getCacheUploadURL(url, sizeInBytes)
	if err != nil {
		return fmt.Errorf("failed to generate upload url: %s", err)
	}

	log.Donef("Attempting Archive Upload")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	if err := tryToUploadArchive(uploadURL, pth); err != nil {
		fmt.Println()
		log.Warnf("First upload attempt failed, retrying...")
		fmt.Println()
		time.Sleep(3000 * time.Millisecond)
		return tryToUploadArchive(uploadURL, pth)
	}
	return nil
}

// getCacheUploadURL requests an upload url from the Bitrise cache API server.
func getCacheUploadURL(cacheAPIURL string, fileSizeInBytes int64) (string, error) {
	req, err := http.NewRequest(http.MethodPost, cacheAPIURL, bytes.NewReader([]byte(fmt.Sprintf(`{"file_size_in_bytes": %d}`, fileSizeInBytes))))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %s", err)
	}

	resp, err := (&http.Client{Timeout: 20 * time.Second}).Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %s", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Warnf("Failed to close response body: %s", err)
		}
	}()

	if resp.StatusCode < 200 || resp.StatusCode > 202 {
		return "", fmt.Errorf("upload url was rejected with status code: %d", resp.StatusCode)
	}

	var respModel map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&respModel); err != nil {
		return "", fmt.Errorf("failed to decode response body: %s", err)
	}

	uploadURL, ok := respModel["upload_url"]
	if !ok || uploadURL == "" {
		return "", fmt.Errorf("request sent, but upload url isn't received")
	}

	return uploadURL, nil
}

// tryToUploadArchive performs the cache upload.
// If the destination is a local file path (url has a file:// scheme) this function copies the cache archive file to the destination.
// Otherwise destination should be a remote url.
func tryToUploadArchive(uploadURL string, archiveFilePath string) error {
	log.Donef("Opening Archive File")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	archFile, err := os.Open(archiveFilePath)
	if err != nil {
		return fmt.Errorf("failed to open archive file for upload (%s): %s", archiveFilePath, err)
	}

	fileClosed := false
	defer func() {
		if fileClosed {
			return
		}
		if err := archFile.Close(); err != nil {
			log.Warnf("Failed to close archive file (%s): %s", archiveFilePath, err)
		}
	}()

	log.Donef("Getting start of archive file")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	fileInfo, err := archFile.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file stats of the archive file (%s): %s", archiveFilePath, err)
	}
	fileSize := fileInfo.Size()

	log.Donef("Creating new http request")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	req, err := http.NewRequest(http.MethodPut, uploadURL, archFile)
	if err != nil {
		return fmt.Errorf("failed to create upload request: %s", err)
	}

	log.Donef("Adding HTTP Headers")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	req.Header.Add("Content-Length", strconv.FormatInt(fileSize, 10))
	req.ContentLength = fileSize

	log.Donef("Perform actual upload")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to upload: %s", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("upload failed with status code: %d", resp.StatusCode)
	}

	fileClosed = true

	return nil
}
