package irsend

import (
	"fmt"
	"os/exec"
)

// SendOnce a single command to the LIRC
func SendOnce(remote string, keys interface{}) ([]byte, error) {
	// generate the cmd to send to LIRC command
	cmdPath := "irsend"
	flags := []string{"SEND_ONCE", remote}
	switch v := keys.(type) {
	case string:
		flags = append(flags, v)
	case []string:
		for _, key := range v {
			flags = append(flags, key)
		}
	default:
		return nil, fmt.Errorf("Unknown type for remote")
	}

	// execute and return output
	cmd := exec.Command(cmdPath, flags...)
	return cmd.CombinedOutput()
}
