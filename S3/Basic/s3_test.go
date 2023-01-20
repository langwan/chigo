package main

import (
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

var fileSystem *S3FileSystem

const (
	Bucket = "chihuo-test"
)

func init() {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		return
	}
	config := Config{}
	yaml.Unmarshal(data, &config)
	fileSystem, err = NewFileSystem(config.Key, config.Secret, config.Endpoint, "")
	if err != nil {
		fmt.Println(err)
	}
}

func TestS3FileSystem_PutFile(t *testing.T) {
	st, err := os.Stat("icon.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	body, err := os.Open("icon.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	fileSystem.PutFile(context.TODO(), Bucket, "icon.jpg", body, "images/jpeg", st.Size())
	body.Close()
	body, err = os.Open("icon.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	fileSystem.PutFile(context.TODO(), Bucket, "icon2.jpg", body, "images/jpeg", st.Size())
}

func TestS3FileSystem_List(t *testing.T) {
	list, err := fileSystem.List(context.TODO(), Bucket, "")
	if err != nil {
		t.Error(err)
		return
	}
	for _, name := range list {
		t.Log(name)
	}
}

func TestS3FileSystem_RemoveFile(t *testing.T) {
	fileSystem.RemoveFile(context.TODO(), Bucket, "icon.jpg")
	fileSystem.RemoveFile(context.TODO(), Bucket, "icon2.jpg")
}
