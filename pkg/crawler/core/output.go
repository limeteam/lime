package core

import (
	log "github.com/sirupsen/logrus"
	"lime/pkg/api/dao"
	"lime/pkg/common"
	"lime/pkg/crawler/novels"
)

// Output output a row data
func (ctx *Context) Output(row map[int]interface{}, namespace ...string) error {
	var outputFields []string
	var ns string
	switch len(namespace) {
	case 0:
		outputFields = ctx.task.OutputFields
		ns = ctx.task.Namespace
	case 1:
		if !ctx.task.OutputToMultipleNamespace {
			return ErrOutputToMultipleTableDisabled
		}
		multConf, ok := ctx.task.MultipleNamespaceConf[namespace[0]]
		if !ok {
			return ErrMultConfNamespaceNotFound
		}
		outputFields = multConf.OutputFields
		ns = namespace[0]
	default:
		return ErrTooManyOutputNamespace
	}

	if err := ctx.checkOutput(row, outputFields); err != nil {
		log.Errorf("checkOutput failed! err:%+v, fields:%#v, row:%+v", err, outputFields, row)
		return err
	}
	log.Debugf("output row:%+v", row)
	log.Println("type:", ctx.task.OutputConfig.Type)
	log.Println("namsespace:", ns)
	switch ctx.task.OutputConfig.Type {
	case common.OutputDbBook:
		if err := novels.SaveBook(row); err != nil {
			log.Println(err)
			return err
		}
	case common.OutputMongodb:
		daospider := dao.SpiderDao{}
		if err := daospider.Add(row); err != nil {
			log.Error(err)
			return err
		}
	case common.OutputYaml:

	default:
		return ErrOutputTypeNotSupported
	}
	return nil
}

func (ctx *Context) checkOutput(row map[int]interface{}, outputFields []string) error {
	if len(outputFields) != len(row) {
		return ErrOutputFieldsNotMatchOutputRow
	}
	for i := 0; i < len(outputFields); i++ {
		if _, ok := row[i]; !ok {
			return ErrOutputFieldsNotMatchOutputRow
		}
	}
	return nil
}
