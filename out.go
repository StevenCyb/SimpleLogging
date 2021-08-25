package slog

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// Lock to lock the file to write logs and to be threat safe
var lock sync.Mutex

// init is used to disable timestamp prefix on logging
func init() {
	log.SetFlags(0)
}

// out is callen by the log function and handle the output stuff
func out(entry Entry) {
	timestamp := time.Now().Format(timeFormat)
	output := ""

	// Format output based on format
	if format == FORMAT_JSON {
		entry["timestamp"] = timestamp

		j, _ := json.Marshal(entry)

		output = string(j)
	} else if format == FORMAT_TEXT {
		output = timestamp + " [" + fmt.Sprintf("%+v", entry["log_level"]) + "]: "
		delete(entry, "log_level")

		first := true
		for key, value := range entry {
			tmp := fmt.Sprintf("%+v", value)
			if key != "" {
				tmp = key + "=" + tmp
			}

			if first {
				first = false
				output += tmp
				continue
			}

			output += ", " + tmp
		}
	}

	// Write log to file if specified
	if path := logFiles.GetPath(); path != "" {
		lock.Lock()
		defer lock.Unlock()

		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return
		}
		defer f.Close()

		f.Write([]byte(output + "\n"))
	}

	// Deside if its an panic or standard log
	if entry["log_level"] == "fatal" {
		log.Panic(output)
	} else {
		log.Println(output)
	}
}
