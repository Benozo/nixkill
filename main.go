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

// main is the entry point of the application. It prompts the user to choose between killing a process by port or PID.
// Depending on the user's choice, it will either:
// 1. Prompt for a port/pid number, find the processes using that port, and ask for confirmation
func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter 'port' to kill A port or Enter 'pid' to kill BY PID (port/pid): ")
	killType, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	killType = strings.TrimSpace(strings.ToLower(killType))

	if killType != "port" && killType != "pid" {
		fmt.Println("Invalid input. Please enter 'port' or 'pid'.")
		return
	}

	fmt.Printf("Enter the %s to kill: ", killType)
	target, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	target = strings.TrimSpace(target)

	if target == "" {
		fmt.Printf("%s cannot be empty.\n", strings.Title(killType))
		return
	}

	if killType == "port" {
		// Validate port number
		if _, err := strconv.Atoi(target); err != nil {
			fmt.Println("Invalid port number.")
			return
		}

		// Get the PID(s) using the given port
		cmd := exec.Command("sh", "-c", fmt.Sprintf("lsof -t -i :%s -sTCP:LISTEN", target))
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error executing lsof command: %v\n", err)
			return
		}
		if len(output) == 0 {
			fmt.Println("Error: No process found using port", target)
			return
		}

		// Trim and split PIDs (Handles multiple PIDs)
		pids := strings.Fields(strings.TrimSpace(string(output)))

		fmt.Printf("Processes %v are using port %s.\n", pids, target)
		fmt.Print("Do you want to kill them? (y/n): ")
		confirm, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return
		}
		confirm = strings.TrimSpace(strings.ToLower(confirm))

		if confirm == confirmYes {
			for _, pid := range pids {
				killProcess(pid)
			}
		} else {
			fmt.Println("Aborted.")
		}
	} else if killType == "pid" {
		// Validate PID
		if _, err := strconv.Atoi(target); err != nil {
			fmt.Println("Invalid PID.")
			return
		}

		fmt.Printf("Do you want to kill process %s? (y/n): ", target)
		confirm, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return
		}
		confirm = strings.TrimSpace(strings.ToLower(confirm))

		if confirm == confirmYes {
			killProcess(target)
		} else {
			fmt.Println("Aborted.")
		}
	}
}

func killProcess(pid string) {
	fmt.Printf("Killing process %s...\n", pid)
	killCmd := exec.Command("sh", "-c", fmt.Sprintf("sudo kill -9 %s", pid))
	err := killCmd.Run()
	if err != nil {
		fmt.Printf("❌ Failed to kill process %s: %v\n", pid, err)
	} else {
		fmt.Printf("✅ Successfully killed process %s\n", pid)
	}
}
