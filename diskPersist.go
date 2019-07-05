package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func saveDataToDisk(name string, data interface{}) error {

	name = name + resultFileNameSuffix

	log.Printf("Saving %s to disk", name)

	dataPath, err := filepath.Abs(name + ".json")
	if err != nil {
		return err
	}

	rf, err := os.OpenFile(dataPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0775)
	if err != nil {
		return err
	}
	defer rf.Close()

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = rf.Write(dataJSON)
	if err != nil {
		return err
	}

	log.Printf("Saved %s to disk", name)

	return nil
}

func readDataFromDisk(name string, data interface{}) error {

	name = name + resultFileNameSuffix

	log.Printf("Reading %s from disk", name)

	dataPath, err := filepath.Abs(name + ".json")
	if err != nil {
		return err
	}

	rf, err := os.Open(dataPath)
	if err != nil {
		return err
	}
	defer rf.Close()

	dataJSON, err := ioutil.ReadAll(rf)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dataJSON, data)
	if err != nil {
		return err
	}

	log.Printf("Read %s from disk", name)

	return nil
}
