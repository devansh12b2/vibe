package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var version = "0.1.0"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "vibe",
		Short: "Vibe - A delightful Git wrapper with personality",
		Long:  `Vibe is a friendly Git wrapper that adds color, better UX, and helpful features to your Git workflow.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}
			// Pass through to git
			executeGitCommand(args)
		},
		SilenceErrors: true, // Suppress error output so we can handle it ourselves
	}

	// Version command
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Vibe",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Vibe v%s\n", version)
		},
	}

	// Status command with enhanced output
	var statusCmd = &cobra.Command{
		Use:   "status",
		Short: "Show the working tree status with style",
		Aliases: []string{"st"},
		Run: func(cmd *cobra.Command, args []string) {
			showStatus()
		},
	}

	// Commit command with enhanced UX
	var commitCmd = &cobra.Command{
		Use:   "commit",
		Short: "Record changes to the repository",
		Aliases: []string{"ci"},
		DisableFlagParsing: true, // Pass all flags through to git
		Run: func(cmd *cobra.Command, args []string) {
			executeGitCommandWithColor("commit", args...)
		},
	}

	// Log command with better formatting
	var logCmd = &cobra.Command{
		Use:   "log",
		Short: "Show commit logs with enhanced formatting",
		DisableFlagParsing: true, // Pass all flags through to git
		Run: func(cmd *cobra.Command, args []string) {
			showLog(args)
		},
	}

	// Push command
	var pushCmd = &cobra.Command{
		Use:   "push",
		Short: "Update remote refs along with associated objects",
		DisableFlagParsing: true, // Pass all flags through to git
		Run: func(cmd *cobra.Command, args []string) {
			executePush(args)
		},
	}

	// Pull command
	var pullCmd = &cobra.Command{
		Use:   "pull",
		Short: "Fetch from and integrate with another repository or branch",
		DisableFlagParsing: true, // Pass all flags through to git
		Run: func(cmd *cobra.Command, args []string) {
			executePull(args)
		},
	}

	// Vibes command - a fun status overview
	var vibesCmd = &cobra.Command{
		Use:   "vibes",
		Short: "Check the vibes of your repository",
		Run: func(cmd *cobra.Command, args []string) {
			checkVibes()
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(commitCmd)
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(pushCmd)
	rootCmd.AddCommand(pullCmd)
	rootCmd.AddCommand(vibesCmd)

	if err := rootCmd.Execute(); err != nil {
		// Check if it's an unknown command error
		errMsg := err.Error()
		if strings.Contains(errMsg, "unknown command") {
			// Extract the command and pass it through to git
			args := os.Args[1:]
			if len(args) > 0 {
				// Suppress the error message and just pass through
				executeGitCommand(args)
				os.Exit(0)
			}
		}
		fmt.Println(err)
		os.Exit(1)
	}
}

// executeGitCommand passes commands directly to git
func executeGitCommand(args []string) {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		os.Exit(1)
	}
}

// executeGitCommandWithColor executes git commands with enhanced output
func executeGitCommandWithColor(command string, args ...string) {
	fullArgs := append([]string{command}, args...)
	executeGitCommand(fullArgs)
}

// showStatus displays an enhanced git status
func showStatus() {
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)
	yellow := color.New(color.FgYellow)

	cyan.Println("✨ Repository Status")
	fmt.Println()

	// Get current branch
	branchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	branchOutput, err := branchCmd.Output()
	var branch string
	if err != nil {
		// Try to get branch name even if HEAD doesn't exist yet
		branchCmd = exec.Command("git", "branch", "--show-current")
		branchOutput, err = branchCmd.Output()
		if err != nil {
			red.Println("❌ Not a git repository")
			return
		}
		branch = strings.TrimSpace(string(branchOutput))
		if branch == "" {
			branch = "main (no commits yet)"
		}
	} else {
		branch = strings.TrimSpace(string(branchOutput))
	}
	cyan.Printf("📍 Branch: ")
	fmt.Println(branch)
	fmt.Println()

	// Get status
	statusCmd := exec.Command("git", "status", "--short")
	statusOutput, err := statusCmd.Output()
	if err != nil {
		red.Println("❌ Error getting status")
		return
	}

	if len(statusOutput) == 0 {
		green.Println("✅ Working tree clean - good vibes!")
		return
	}

	// Parse and colorize status
	lines := strings.Split(string(statusOutput), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "M ") || strings.HasPrefix(line, " M") {
			yellow.Printf("📝 %s\n", line)
		} else if strings.HasPrefix(line, "A ") {
			green.Printf("➕ %s\n", line)
		} else if strings.HasPrefix(line, "D ") {
			red.Printf("➖ %s\n", line)
		} else if strings.HasPrefix(line, "??") {
			cyan.Printf("❓ %s\n", line)
		} else {
			fmt.Println(line)
		}
	}
}

// showLog displays an enhanced git log
func showLog(args []string) {
	// Use pretty format with colors
	logArgs := []string{
		"log",
		"--pretty=format:%C(yellow)%h%C(reset) - %C(cyan)%an%C(reset) %C(green)(%ar)%C(reset)%n  %s%n",
		"--graph",
	}
	logArgs = append(logArgs, args...)
	executeGitCommand(logArgs)
}

// executePush handles push with feedback
func executePush(args []string) {
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	
	cyan.Println("🚀 Pushing changes...")
	fmt.Println()
	
	fullArgs := append([]string{"push"}, args...)
	executeGitCommand(fullArgs)
	
	fmt.Println()
	green.Println("✅ Push complete!")
}

// executePull handles pull with feedback
func executePull(args []string) {
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	
	cyan.Println("⬇️  Pulling changes...")
	fmt.Println()
	
	fullArgs := append([]string{"pull"}, args...)
	executeGitCommand(fullArgs)
	
	fmt.Println()
	green.Println("✅ Pull complete!")
}

// checkVibes shows a fun overview of the repository
func checkVibes() {
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen)
	yellow := color.New(color.FgYellow)
	magenta := color.New(color.FgMagenta)

	cyan.Println("🎵 Checking the vibes...")
	fmt.Println()

	// Get commit count
	commitCmd := exec.Command("git", "rev-list", "--count", "HEAD")
	commitOutput, err := commitCmd.Output()
	if err == nil {
		commits := strings.TrimSpace(string(commitOutput))
		magenta.Printf("📊 Total commits: %s\n", commits)
	}

	// Get contributor count
	contributorCmd := exec.Command("git", "shortlog", "-sn", "--all")
	contributorOutput, err := contributorCmd.Output()
	if err == nil {
		contributors := strings.Split(strings.TrimSpace(string(contributorOutput)), "\n")
		yellow.Printf("👥 Contributors: %d\n", len(contributors))
	}

	// Get current branch
	branchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	branchOutput, err := branchCmd.Output()
	if err == nil {
		branch := strings.TrimSpace(string(branchOutput))
		cyan.Printf("🌿 Current branch: %s\n", branch)
	}

	// Check if working tree is clean
	statusCmd := exec.Command("git", "status", "--short")
	statusOutput, err := statusCmd.Output()
	if err == nil {
		if len(statusOutput) == 0 {
			green.Println("✨ Status: Clean - immaculate vibes!")
		} else {
			yellow.Println("📝 Status: Changes detected - creative energy flowing!")
		}
	}

	fmt.Println()
	cyan.Println("🎉 The vibes are strong with this one!")
}

