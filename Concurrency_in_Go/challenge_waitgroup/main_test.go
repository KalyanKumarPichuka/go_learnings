package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestUpdateMessage(t *testing.T) {

	wg.Add(1)
	go updateMessage("First")
	wg.Wait()

	if msg != "First" {
		t.Errorf("Expected 'First' but got %s", msg)
	}
}

func TestPrintMessage(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "World"
	printMessage()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "World") {
		t.Errorf("Unexpected output: %s", output)
	}
}
