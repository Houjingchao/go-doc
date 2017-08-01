package Concurrent

import (
	"os/exec"
	"fmt"
	"bytes"
	"io"
	"bufio"
)

func main() {
	useBufferedIO := false
	cmd0 := exec.Command("echo", "-n", "my first go command~")

	stdout0, err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Println("Error: can not obtain the stdout :%s/n", err)
	}
	if err = cmd0.Start(); err != nil {
		fmt.Println("Error: cmd0 exec wrong:%s/n", err)
	}

	if !useBufferedIO {
		var outputBuf0 bytes.Buffer

		for {
			tempOutput := make([]byte, 5)
			n, err := stdout0.Read(tempOutput)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Printf("Error: Couldn't read data from the pipe: %s\n", err)
					return
				}
			}
			if n > 0 {
				outputBuf0.Write(tempOutput[:n])
			}
		}

	} else {
		outputBuf0 := bufio.NewReader(stdout0)
		output0, _, err := outputBuf0.ReadLine()
		if err != nil {
			fmt.Printf("Error: Couldn't read data from the pipe: %s\n", err)
			return
		}
		fmt.Printf("%s\n", string(output0))
	}
}
