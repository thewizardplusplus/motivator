package systemutils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

func FindBackgroundProcess(
	executableName string,
	foregroundProcessPID int,
) (ps.Process, error) {
	processes, err := ps.Processes()
	if err != nil {
		return nil, fmt.Errorf("unable to get the process list: %w", err)
	}

	var backgroundProcess ps.Process
	for _, process := range processes {
		if process.Executable() == executableName &&
			process.Pid() != foregroundProcessPID {
			backgroundProcess = process
			break
		}
	}

	return backgroundProcess, nil
}

func StartBackgroundProcess(
	executableName string,
	arguments []string,
	stderrWriter io.Writer,
	stderrEndMark string,
) error {
	backgroundProcess := exec.Command(executableName, arguments...)

	stderr, err := backgroundProcess.StderrPipe()
	if err != nil {
		const message = "unable to get the stderr of the background process: %w"
		return fmt.Errorf(message, err)
	}
	defer stderr.Close()

	if err := backgroundProcess.Start(); err != nil {
		return fmt.Errorf("unable to start the background process: %w", err)
	}

	stderrScanner := bufio.NewScanner(stderr)
	for stderrScanner.Scan() {
		stderrLine := stderrScanner.Text()
		if strings.HasSuffix(stderrLine, stderrEndMark) {
			break
		}

		if _, err := fmt.Fprintln(stderrWriter, stderrLine); err != nil {
			const message = "unable to write the stderr of the background process: %w"
			return fmt.Errorf(message, err)
		}
	}
	if err := stderrScanner.Err(); err != nil {
		const message = "unable to read the stderr of the background process: %w"
		return fmt.Errorf(message, err)
	}

	return nil
}

func KillBackgroundProcess(
	executableName string,
	foregroundProcessPID int,
) error {
	backgroundProcess, err :=
		FindBackgroundProcess(executableName, foregroundProcessPID)
	if err != nil {
		return fmt.Errorf("unable to find the background process by name: %w", err)
	}
	if backgroundProcess == nil {
		return nil
	}

	osBackgroundProcess, err := os.FindProcess(backgroundProcess.Pid())
	if err != nil {
		return fmt.Errorf("unable to find the background process by PID: %w", err)
	}

	if err := osBackgroundProcess.Kill(); err != nil {
		return fmt.Errorf("unable to kill the background process: %w", err)
	}

	return nil
}
