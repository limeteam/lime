package download

import (
	"lime/pkg/down"
	"fmt"
	"github.com/rs/zerolog"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

var (
	config   string
	url      string
	driver   string
	loglevel uint8
	StartCmd = &cobra.Command{
		Use:     "download",
		Short:   "d",
		Example: "download",
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
	StartCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "please input url")
	StartCmd.PersistentFlags().StringVarP(&driver, "driver", "d", "none", "please input driver")
}

func usage() {
	usageStr := `_                       _ 
                       | |                     | |
  ____  ____ ____ _ _ _| |____   ___ _   _ ____| |
 / ___)/ ___) _  | | | | |  _ \ / _ \ | | / _  ) |
( (___| |  ( ( | | | | | | | | | |_| \ V ( (/ /| |
 \____)_|   \_||_|\____|_|_| |_|\___/ \_/ \____)_|
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
	down.Download(url, driver)
	return nil
}
