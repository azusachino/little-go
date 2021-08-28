package cmd

import (
	"github.com/little-go/road-to-go/tour/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	MODE_UPPER = iota + 1
	MODE_LOWER
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE
	MODE_CAMELCASE_TO_UNDERCASE
)

var (
	str  string
	mode int8
)

var desc = strings.Join([]string{
	"support multiple format of words",
	"1: to upper case",
	"2: to lower case",
	"3: under score to upper case",
	"4: under score to lower case",
	"5: camel case to under score",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "Transfer format of word",
	//Long:  "Support multiple format of words",
	Long: desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = word.UnderscoreToUpperCamelCase(str)
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = word.UnderscoreToLowerCamelCase(str)
		case MODE_CAMELCASE_TO_UNDERCASE:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("mode not supported, please check help doc.")
		}
		log.Printf("output: %s", content)
	},
}

func init() {
	fs := wordCmd.Flags()
	fs.StringVarP(&str, "str", "s", "", "please input the word")
	fs.Int8VarP(&mode, "mode", "m", 0, "please select mode")
}
