package main

import (
    "io/ioutil"
    "os"
    "testing"
)

func TestHelloWorld(t *testing.T) {
    // Redirect stdout to a buffer
    originalStdout := os.Stdout // keep backup of the real stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    // Restore original stdout when the function returns
    defer func() {
        w.Close()
        os.Stdout = originalStdout
    }()

    // Call your main function
    main()

    // Close the write end of the pipe
    w.Close()

    // Read what was written to the buffer
    out, _ := ioutil.ReadAll(r)

    // Check the output
    expected := "Hello, World!\n"
    if string(out) != expected {
        t.Errorf("Unexpected output: got %q want %q", string(out), expected)
    }
}
