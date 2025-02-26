package cmd

import (
	"fmt"
	"strings"

	"github.com/jkulton/gg/internal/collection"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rot13Cmd)
}

var rot13Cmd = &cobra.Command{
	Use:   "rot13",
	Short: "Apply ROT13 cipher to a string",
	Run: func(cmd *cobra.Command, args []string) {
		r := collection.MapStrings(args, rot13)
		s := strings.Join(r, " ")
		fmt.Println(s)
	},
}

func rot13(s string) string {
	var result strings.Builder
	result.Grow(len(s))

	for i := range len(s) {
		c := s[i]
		switch {
		case 'a' <= c && c <= 'z':
			result.WriteByte('a' + (c-'a'+13)%26)
		case 'A' <= c && c <= 'Z':
			result.WriteByte('A' + (c-'A'+13)%26)
		default:
			result.WriteByte(c)
		}
	}

	return result.String()
}
