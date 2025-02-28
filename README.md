# nixkill

`nixkill` is a lightweight command-line tool that helps you quickly identify and terminate processes using a specific port or PID on your system. This is particularly useful for developers who need to free up occupied ports or terminate specific processes without manually searching for process IDs.

## ✨ Features
- **Find and kill processes by port or PID** 🛑
- **Interactive confirmation before termination** ⚠️
- **Lightweight & easy to use** ⚡
- **Works on Linux & macOS** 🐧🍏

---

## 🚀 Installation

### Install via `go install`
To install `nixkill` using Go:
```sh
go install github.com/Benozo/nixkill@latest
```
Ensure `$GOPATH/bin` is in your system's `PATH` to run `nixkill` directly:
```sh
export PATH=$PATH:$(go env GOPATH)/bin
```

### Install via Git (Manual Build)
1. Clone the repository:
   ```sh
   git clone https://github.com/Benozo/nixkill.git
   ```
2. Navigate into the project directory:
   ```sh
   cd nixkill
   ```
3. Build the binary:
   ```sh
   go build -o nixkill
   ```
4. Make it executable:
   ```sh
   chmod +x nixkill
   ```

---

## 📌 Usage

Run `nixkill` to free up a specific port:

```sh
 ./nixkill 
Enter 'port' to kill A port or Enter 'pid' to kill BY PID (port/pid): port
Enter the port to kill: 8080
Processes [39420] are using port 8080.
Do you want to kill them? (y/n): y
Killing process 39420...
✅ Successfully killed process 39420
```

Run `nixkill` to kill a specific PID:

```sh
./nixkill 
Enter 'port' to kill A port or Enter 'pid' to kill BY PID (port/pid): pid
Enter the pid to kill: 37806
Do you want to kill process 37806? (y/n): y
Killing process 37806...
✅ Successfully killed process 37806
```

If a process is found, it will prompt for confirmation before terminating it.

---

## 🛠️ Roadmap & Contributions
We’re always looking to improve! Future enhancements may include:
- Windows support 🖥️

Contributions are welcome! Feel free to fork the repository and submit pull requests.

---

## 📜 License
`nixkill` is released under the **MIT License**.

🔗 **GitHub Repository**: [github.com/Benozo/nixkill](https://github.com/Benozo/nixkill)  

---
````
