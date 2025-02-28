package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	confirmYes = "y"
	confirmNo  = "n"
)

// main is the entry point of the program. It prompts the user to enter a port number,
// finds the process IDs (PIDs) that are using the specified port, and asks the user
// for confirmation to kill those processes. If the user confirms, it attempts to
// kill the processes using the specified port.
//
// The program performs the following steps:
// 1. Prompts the user to enter a port number.
// 2. Validates that the port number is not empty.
// 3. Uses the `lsof` command to find the PIDs of processes listening on the specified port.
// 4. If no processes are found, it prints an error message and exits.
// 5. If processes are found, it prints the PIDs and asks the user for confirmation to kill them.
// 6. If the user confirms, it attempts to kill each process using the `kill` command.
// 7. Prints the result of each kill attempt (success or failure).
func main() {
	fmt.Print("Enter the port to kill: ")
	reader := bufio.NewReader(os.Stdin)
	port, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	port = strings.TrimSpace(port)

	if port == "" {
		fmt.Println("Port cannot be empty.")
		return
	}

	// Validate port number
	if _, err := strconv.Atoi(port); err != nil {
		fmt.Println("Invalid port number.")
		return
	}

	// Get the PID(s) using the given port
	cmd := exec.Command("sh", "-c", fmt.Sprintf("lsof -t -i :%s -sTCP:LISTEN", port))
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing lsof command: %v\n", err)
		return
	}
	if len(output) == 0 {
		fmt.Println("Error: No process found using port", port)
		return
	}

	// Trim and split PIDs (Handles multiple PIDs)
	pids := strings.Fields(strings.TrimSpace(string(output)))

	fmt.Printf("Processes %v are using port %s.\n", pids, port)
	fmt.Print("Do you want to kill them? (y/n): ")
	confirm, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	confirm = strings.TrimSpace(strings.ToLower(confirm))

	if confirm == confirmYes {
		for _, pid := range pids {
			fmt.Printf("Killing process %s...\n", pid)
			killCmd := exec.Command("sh", "-c", fmt.Sprintf("sudo kill -9 %s", pid))
			err = killCmd.Run()
			if err != nil {
				fmt.Printf("❌ Failed to kill process %s: %v\n", pid, err)
			} else {
				fmt.Printf("✅ Successfully killed process %s\n", pid)
			}
		}
	} else {
		fmt.Println("Aborted.")
	}
}
