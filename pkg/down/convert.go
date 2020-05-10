package down

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"lime/pkg/down/output"
)

var outputOpt output.Option

func Covert(filename string, format string, outputpath string) error {
	if err := LoadLocalStore(filename); err != nil {
		return err
	}
	if format == "" {
		return nil
	}
	var ConversionFileName string
	if outputpath == "" {
		ConversionFileName = fmt.Sprintf("%s.%s", chapter.BookName, format)
	}
	log.Printf("Start Conversion: Format:%#v OutPath:%#v", format, ConversionFileName)
	return output.Output(*chapter, format, ConversionFileName, outputOpt)
}
