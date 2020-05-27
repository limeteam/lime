package search

import (
	"fmt"
	"github.com/rs/zerolog"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"lime/pkg/down"
	"os"
	"strings"
)

var (
	config   string
	keyword  string
	put      bool
	filename string
	driver   string
	loglevel uint8
	StartCmd = &cobra.Command{
		Use:     "search",
		Short:   "se",
		Example: "search",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/in-local.yaml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&keyword, "keyword", "k", "", "input keyword")
	StartCmd.PersistentFlags().Bool("put", false, "put")
	StartCmd.PersistentFlags().StringVarP(&filename, "filename", "n", "", "input convert filename")
	StartCmd.PersistentFlags().StringVarP(&driver, "driver", "d", "none", "please input driver")
}

func usage() {
	usageStr := `
 ___      ___   __   __  _______ 
|   |    |   | |  |_|  ||       |
|   |    |   | |       ||    ___|
|   |    |   | |       ||   |___ 
|   |___ |   | |       ||    ___|
|       ||   | | ||_|| ||   |___ 
|_______||___| |_|   |_||_______|
`
	fmt.Printf("%s\n", usageStr)
}

func setup() {
	//1.Set up log level
	zerolog.SetGlobalLevel(zerolog.Level(loglevel))
	//2.Set up configuration
	viper.SetConfigFile(config)
	content, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}
	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}
}

func run() error {
	down.Search(keyword, put, filename, driver)
	return nil
}
