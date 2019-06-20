package targzhelpers

import (
	"fmt"
	"testing"
	"os"
	"path/filepath"
)

func setUpTarIt()(string, error) {
	tempDir := os.TempDir()
	if err := os.Mkdir(tempDir, 0777); err != nil {
		return "", err
	}

	return tempDir, nil
}

func tearDownTarIt(dirPath string) error {
	err := os.RemoveAll(dirPath)
	return err
}

func testMakeTarWithWithValidPath(t *testing.T, testDirPath string) {
	testFilePath := fmt.Sprintf("%s/file.tar", testDirPath)
	filePath := filepath.Join(testDirPath, "file")
	os.Create(filePath)

	err := TarIt(testFilePath, filePath)
	if err != nil {
		t.Errorf("Failed")
	}
}

func TestTarIt(t *testing.T) {
	tempDir, err := setUpTarIt()

	testMakeTarWithWithValidPath(t, tempDir)
	if err != nil {
		t.Errorf("Failed")
	}
}
