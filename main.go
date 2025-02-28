package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Print("Enter the port to kill: ")
	reader := bufio.NewReader(os.Stdin)
	port, _ := reader.ReadString('\n')
	port = strings.TrimSpace(port)

	if port == "" {
		fmt.Println("Port cannot be empty.")
		return
	}

	// Get the PID(s) using the given port
	cmd := exec.Command("sh", "-c", fmt.Sprintf("lsof -t -i :%s -sTCP:LISTEN", port))
	output, err := cmd.Output()

	if err != nil || len(output) == 0 {
		fmt.Println("Error: No process found using port", port)
		return
	}

	// Trim and split PIDs (Handles multiple PIDs)
	pids := strings.Fields(strings.TrimSpace(string(output)))

	fmt.Printf("Processes %v are using port %s.\n", pids, port)
	fmt.Print("Do you want to kill them? (y/n): ")
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(strings.ToLower(confirm))

	if confirm == "y" {
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
