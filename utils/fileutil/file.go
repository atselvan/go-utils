package fileutil

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/atselvan/go-utils/utils/errors"
)

// OpenFile opens a file
// The method returns an errors if there is an issue with opening the file
func OpenFile(filePath string) (*os.File, *errors.Error) {
	f, err := os.Open(filePath)
	if err != nil {
		return f, errors.Newf(
			errors.ErrCodeFileOpenError,
			0,
			errors.ErrMsg[errors.ErrCodeFileOpenError],
			filePath,
			err.Error(),
		)
	}
	return f, nil
}

// FileExists checks if a file exists and returns an errors if the file was not found
func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return false
	}
	return true
}

// CreateFile creates a new file
// The method returns an errors if there was an issue with creating an new file
func CreateFile(filePath string) (*os.File, *errors.Error) {
	f, err := os.Create(filePath)
	if err != nil {
		return f, errors.Newf(
			errors.ErrCodeFileCreateError,
			0,
			errors.ErrMsg[errors.ErrCodeFileCreateError],
			filePath,
			err.Error(),
		)
	}
	return f, nil
}

// RemoveFile removes files from the provided valid filePath.
func RemoveFile(filePath string) *errors.Error {
	if err := os.Remove(filePath); err != nil {
		return errors.Newf(
			errors.ErrCodeFileRemoveError,
			0,
			errors.ErrMsg[errors.ErrCodeFileRemoveError],
			filePath,
			err.Error(),
		)
	}
	return nil
}

// ReadFile checks if a file exists and if it does tries to reads the contents of the
// file and returns the data back
// The method returns an errors the file does not exist or if there was an errors in reading the contents of the file
func ReadFile(filePath string) ([]byte, *errors.Error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.Newf(
			errors.ErrCodeFileReadError,
			0,
			errors.ErrMsg[errors.ErrCodeFileReadError],
			filePath,
			err.Error(),
		)
	}
	return data, nil
}

// WriteFile creates a new file if the file does not exists and writes data into the file
// The method returns an errors if there was an issue creating a new file
// or while writing data into the file
func WriteFile(filePath string, data []byte) *errors.Error {
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		fmt.Println(err)
		return errors.Newf(
			errors.ErrCodeFileWriteError,
			0,
			errors.ErrMsg[errors.ErrCodeFileWriteError],
			filePath,
			err.Error(),
		)
	}
	return nil
}

// ReadJsonFile reads a yaml file and puts the contents into the out variables
// out variable should be a pointer to a valid struct
// The method returns and errors if reading a file or the unmarshal process fails
func ReadJsonFile(filePath string, out any) *errors.Error {
	data, cErr := ReadFile(filePath)
	if cErr != nil {
		return cErr
	}
	if err := json.Unmarshal(data, out); err != nil {
		return errors.Newf(
			errors.ErrCodeJSONUnmarshalError,
			0,
			errors.ErrMsg[errors.ErrCodeJSONUnmarshalError], err.Error(),
		)
	}
	return nil
}

// WriteJSONFile encodes the data from an input interface into json format
// and writes the data into a file
// The in interface should be an address to a valid struct
// The method returns an errors if there is an errors with the json encode
// or with writing to the file
func WriteJSONFile(filePath string, in any) *errors.Error {
	data, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return errors.Newf(
			errors.ErrCodeJSONMarshalError,
			0,
			errors.ErrMsg[errors.ErrCodeJSONMarshalError], err.Error(),
		)
	}
	if cErr := WriteFile(filePath, data); cErr != nil {
		return cErr
	}
	return nil
}
