package fileutil

import (
	"fmt"
	"testing"

	"github.com/atselvan/go-utils/utils/errors"
	"github.com/stretchr/testify/assert"
)

var (
	testDataPath        = "testdata"
	testFilePath        = testDataPath + "/test.txt"
	testJsonFilePath    = testDataPath + "/test.json"
	validTestFilePath   = testDataPath + "/file.txt"
	invalidTestFilePath = "invalid/test.txt"
	validJsonFilePath   = testDataPath + "/valid.json"
	invalidJsonFilePath = testDataPath + "/invalid.json"
)

type Out struct {
	Foo string `json:"foo"`
}

func TestFileExists(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		assert.True(t, FileExists(validTestFilePath))
	})

	t.Run("error", func(t *testing.T) {
		assert.False(t, FileExists(invalidTestFilePath))
	})
}

func TestOpenFile(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		_, cErr := OpenFile(validTestFilePath)
		assert.Nil(t, cErr)
	})

	t.Run("error", func(t *testing.T) {
		_, cErr := OpenFile(invalidTestFilePath)
		assert.NotNil(t, cErr)
		assert.Equal(t, errors.ErrCodeFileOpenError, cErr.Code)
		assert.Contains(t, cErr.Message, fmt.Sprintf(errors.ErrMsg[errors.ErrCodeFileOpenError], invalidTestFilePath, ""))
	})
}

func TestCreateFile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		_, cErr := CreateFile(testFilePath)
		assert.Nil(t, cErr)
		assert.True(t, FileExists(testFilePath))
	})

	t.Run("error", func(t *testing.T) {
		_, cErr := CreateFile(invalidTestFilePath)
		assert.NotNil(t, cErr)
		assert.Equal(t, errors.ErrCodeFileCreateError, cErr.Code)
		assert.Contains(t, cErr.Message, fmt.Sprintf(errors.ErrMsg[errors.ErrCodeFileCreateError], invalidTestFilePath, ""))
	})
}

func TestRemoveFile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		cErr := RemoveFile(testFilePath)
		assert.Nil(t, cErr)
		assert.False(t, FileExists(testFilePath))
	})

	t.Run("error", func(t *testing.T) {
		cErr := RemoveFile(invalidTestFilePath)
		assert.NotNil(t, cErr)
		assert.Equal(t, errors.ErrCodeFileRemoveError, cErr.Code)
		assert.Contains(t, cErr.Message, fmt.Sprintf(errors.ErrMsg[errors.ErrCodeFileRemoveError], invalidTestFilePath, ""))
	})
}

func TestReadFile(t *testing.T) {
	t.Run("read valid file", func(t *testing.T) {
		data, cErr := ReadFile(validTestFilePath)
		assert.Nil(t, cErr)
		assert.Equal(t, "some very valuable test data", string(data))
	})

	t.Run("invalid", func(t *testing.T) {
		_, cErr := ReadFile(invalidTestFilePath)
		assert.NotNil(t, cErr)
		assert.Equal(t, errors.ErrCodeFileReadError, cErr.Code)
		assert.Contains(t, cErr.Message, fmt.Sprintf(errors.ErrMsg[errors.ErrCodeFileReadError], invalidTestFilePath, ""))
	})
}

func TestWriteFile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		cErr := WriteFile(testFilePath, []byte("something"))
		assert.Nil(t, cErr)
	})

	t.Run("error", func(t *testing.T) {
		cErr := WriteFile(invalidTestFilePath, []byte("something"))
		assert.NotNil(t, cErr)
		assert.Equal(t, errors.ErrCodeFileWriteError, cErr.Code)
		assert.Contains(t, cErr.Message, fmt.Sprintf(errors.ErrMsg[errors.ErrCodeFileWriteError], invalidTestFilePath, ""))
	})

	// cleanup
	cErr := RemoveFile(testFilePath)
	assert.Nil(t, cErr)
}

func TestReadJsonFile(t *testing.T) {
	t.Run("read valid json", func(t *testing.T) {
		out := new(Out)
		cErr := ReadJsonFile(validJsonFilePath, out)
		assert.Nil(t, cErr)
		assert.Equal(t, "bar", out.Foo)
	})

	t.Run("read error", func(t *testing.T) {
		out := new(Out)
		cErr := ReadJsonFile(invalidTestFilePath, out)
		assert.NotNil(t, cErr)
		assert.Equal(t, errors.ErrCodeFileReadError, cErr.Code)
		assert.Contains(t, cErr.Message,
			fmt.Sprintf(errors.ErrMsg[errors.ErrCodeFileReadError], invalidTestFilePath, ""))
	})

	t.Run("read invalid json", func(t *testing.T) {
		out := new(Out)
		cErr := ReadJsonFile(invalidJsonFilePath, out)
		assert.NotNil(t, cErr)
		assert.Equal(t, errors.ErrCodeJSONUnmarshalError, cErr.Code)
		assert.Contains(t, cErr.Message,
			fmt.Sprintf(errors.ErrMsg[errors.ErrCodeJSONUnmarshalError], ""))
	})
}

func TestWriteJSONFile(t *testing.T) {
	in := &Out{Foo: "bar"}
	t.Run("success", func(t *testing.T) {
		cErr := WriteJSONFile(testJsonFilePath, in)
		assert.Nil(t, cErr)
	})

	t.Run("json marshal error", func(t *testing.T) {
		in := make(chan int)
		cErr := WriteJSONFile(testJsonFilePath, &in)
		assert.NotNil(t, cErr)
		assert.Equal(t, errors.ErrCodeJSONMarshalError, cErr.Code)
		assert.Contains(t, cErr.Message, fmt.Sprintf(errors.ErrMsg[errors.ErrCodeJSONMarshalError], ""))
	})

	t.Run("write error", func(t *testing.T) {
		cErr := WriteJSONFile(invalidTestFilePath, in)
		assert.NotNil(t, cErr)
		assert.Equal(t, errors.ErrCodeFileWriteError, cErr.Code)
		assert.Contains(t, cErr.Message, fmt.Sprintf(errors.ErrMsg[errors.ErrCodeFileWriteError], invalidTestFilePath, ""))
	})

	cErr := RemoveFile(testJsonFilePath)
	assert.Nil(t, cErr)
}
