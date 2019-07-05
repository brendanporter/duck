package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func saveProcessor() {

	saveTicker := time.NewTicker(time.Minute)

	for {
		select {
		case <-saveTicker.C:
			err := saveDataToDisk("pings", pingResults)
			if err != nil {
				log.Print(err)
			}

			pathResultsRequestChan <- 1
			prr := <-pathResultsResponseChan

			err = saveDataToDisk("paths", prr)
			if err != nil {
				log.Print(err)
			}

			pathsResult := PathsResult{
				Paths:       prr,
				MessageType: "paths",
			}

			pathsJSON, err := json.Marshal(pathsResult)
			if err != nil {
				elog.Print(err)
			}

			wsHub.broadcast <- string(pathsJSON)

			hostResultsRequestChan <- 1
			hrr := <-hostResultsResponseChan

			err = saveDataToDisk("hosts", hrr)
			if err != nil {
				log.Print(err)
			}

			hostsResult := HostsResult{
				Hosts:       hrr,
				MessageType: "hosts",
			}

			hostsJSON, err := json.Marshal(hostsResult)
			if err != nil {
				elog.Print(err)
			}

			wsHub.broadcast <- string(hostsJSON)

		}
	}
}

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
