package explainshell

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "explainshell",
	Short: "A CLI to explain shell commands",
	Long: `A CLI to explain shell commands using OpenAi GPT-3`,
  }

  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }