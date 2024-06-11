package cmd

import (
	"github.com/azusachino/golong/road-to-go/tour/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	calculateTime string
	duration      string
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "process time",
	Long:  "process time",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var nowTimeCmd = &cobra.Command{
	Use: "now",
	Run: func(cmd *cobra.Command, args []string) {
		now := timer.GetNowTime()
		log.Printf("output: %s, %d", now.Format("2006-01-02 15:04:05"), now.Unix())
	},
}

var calcTimeCmd = &cobra.Command{
	Use: "calc",
	Run: func(cmd *cobra.Command, args []string) {
		var current time.Time
		var layout = time.RFC3339
		if calculateTime == "" {
			current = timer.GetNowTime()
		} else {
			var err error
			if !strings.Contains(calculateTime, " ") {
				layout = "2006-01-02"
			}
			current, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				current = time.Unix(int64(t), 0)
			}
		}
		calculateTime, err := timer.CalculateTime(current, duration)
		if err != nil {
			log.Fatalf("err %v", err)
		}
		log.Printf("output %s, %d", calculateTime.Format(layout), calculateTime.Unix())
	},
}

func Time() {
	timeCmd.AddCommand(nowTimeCmd, calcTimeCmd)
	fs := calcTimeCmd.Flags()
	fs.StringVarP(&calculateTime, "calculate", "c", "", `calculate time duration`)
	fs.StringVarP(&duration, "duration", "d", "", `duration`)
}
