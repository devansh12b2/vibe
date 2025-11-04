# 🎵 Vibe - A Git Wrapper with Personality

Vibe is a delightful Git wrapper written in Go that adds color, better UX, and helpful features to your Git workflow. It's designed to make your Git experience more enjoyable while maintaining full compatibility with standard Git commands.

## ✨ Features

- **Enhanced Status**: Beautiful, colorized status output with emojis
- **Better Logs**: Formatted git log with colors and graph visualization
- **Smart Commands**: Push and pull with friendly feedback
- **Vibes Check**: Get a fun overview of your repository health
- **Pass-through Support**: Any command not specifically handled by Vibe is passed directly to Git
- **Command Aliases**: Short aliases for common commands (e.g., `st` for status)

## 🚀 Installation

### Prerequisites

- Go 1.21 or higher
- Git installed on your system

### Build from Source

```bash
# Clone the repository
cd "/Users/devamish/Documents/Vibe Wrap"

# Download dependencies
go mod download

# Build the binary
go build -o vibe

# (Optional) Install to your PATH
sudo mv vibe /usr/local/bin/
```

Or use `go install`:

```bash
go install
```

## 📖 Usage

### Basic Commands

```bash
# Check version
vibe version

# Enhanced status
vibe status
# or use the alias
vibe st

# Commit changes
vibe commit -m "Your commit message"

# View enhanced logs
vibe log

# Push with style
vibe push origin main

# Pull with feedback
vibe pull

# Check your repo vibes
vibe vibes
```

### Pass-through Commands

Any Git command not specifically handled by Vibe is passed directly to Git:

```bash
vibe add .
vibe branch -a
vibe checkout -b feature/new-feature
vibe merge develop
vibe rebase main
```

## 🎨 Command Reference

| Command | Alias | Description |
|---------|-------|-------------|
| `vibe status` | `st` | Show working tree status with colors and emojis |
| `vibe commit` | `ci` | Record changes to the repository |
| `vibe log` | - | Show commit logs with enhanced formatting |
| `vibe push` | - | Push changes with feedback |
| `vibe pull` | - | Pull changes with feedback |
| `vibe vibes` | - | Check the overall health and stats of your repo |
| `vibe version` | - | Display Vibe version |

## 🎯 Examples

### Check Repository Status
```bash
vibe status
```
Output:
```
✨ Repository Status

📍 Branch: main

📝  M main.go
➕ A  new-file.go
❓ ?? untracked.txt
```

### View Commit History
```bash
vibe log -5
```

### Check Your Vibes
```bash
vibe vibes
```
Output:
```
🎵 Checking the vibes...

📊 Total commits: 42
👥 Contributors: 3
🌿 Current branch: main
✨ Status: Clean - immaculate vibes!

🎉 The vibes are strong with this one!
```

## 🛠️ Development

### Project Structure

```
vibe/
├── main.go       # Main application code
├── go.mod        # Go module definition
└── README.md     # Documentation
```

### Adding New Commands

To add a new command, create a new `cobra.Command` in `main.go`:

```go
var myCmd = &cobra.Command{
    Use:   "mycommand",
    Short: "Description",
    Run: func(cmd *cobra.Command, args []string) {
        // Your implementation
    },
}
rootCmd.AddCommand(myCmd)
```

## 🤝 Contributing

Contributions are welcome! Feel free to:
- Add new features
- Improve existing commands
- Fix bugs
- Enhance documentation

## 📝 License

MIT License - feel free to use this project however you'd like!

## 🌟 Why Vibe?

Git is powerful but can be intimidating. Vibe makes Git more approachable and fun while maintaining all the power you need. Whether you're a Git novice or expert, Vibe adds a touch of personality to your workflow.

Happy coding with good vibes! ✨

