package internal

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDownloadToFile(t *testing.T) {

	tmpFile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name()) // clean up
	err = DownloadToFile("https://www.google.it", tmpFile.Name(), 0777)
	if err != nil {
		t.Errorf("Error %v", err)
	}
}

func TestDownload(t *testing.T) {
	out, err := Download("https://www.google.it")
	if err != nil {
		t.Error(err)
	}
	if len(out) == 0 {
		t.Error("Empty response")
	}
}

func TestFileExists(t *testing.T) {
	mustTrue := FileExists("/etc/hosts")
	mustFalse := FileExists("/empty.file")
	if !mustTrue {
		t.Error("Dockerfile doesn't exists")
	}
	if mustFalse {
		t.Fail()
	}
}

func TestExistsInPath(t *testing.T) {
	_, mustTrue := ExistsInPath("bash")
	if !mustTrue {
		t.Fail()
	}
}

func TestUnzip(t *testing.T){
	t.Skipf("TODO: write resources for Unzip")
}

func TestCopyFile(t *testing.T){
	t.Skipf("TODO: write resources for CopyFile")
}

func TestMakeDir(t *testing.T){
	t.Skipf("TODO: write resources for MakeDir")
}

func TestTouchFile(t *testing.T){
	t.Skipf("TODO: write resources for TouchFile")
}

func TestLoadFromFileOrUri(t *testing.T){
	t.Skipf("TODO: write resources for LoadFromFileOrUri")
}


func TestHashString(t *testing.T){
	t.Skipf("TODO: write resources for SHA256String")
}

func TestHashFile(t *testing.T){
	t.Skipf("TODO: write resources for SHA256File")
}
