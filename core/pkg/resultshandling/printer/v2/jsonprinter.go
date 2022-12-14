package printer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	logger "github.com/kubescape/go-logger"
	"github.com/kubescape/go-logger/helpers"
	"github.com/kubescape/kubescape/v2/core/cautils"
	"github.com/kubescape/kubescape/v2/core/pkg/resultshandling/printer"
)

const (
	jsonOutputFile = "report"
	jsonOutputExt  = ".json"
)

type JsonPrinter struct {
	writer *os.File
}

func NewJsonPrinter() *JsonPrinter {
	return &JsonPrinter{}
}

func (jp *JsonPrinter) SetWriter(outputFile string) {
	if strings.TrimSpace(outputFile) == "" {
		outputFile = jsonOutputFile
	}
	if filepath.Ext(strings.TrimSpace(outputFile)) != jsonOutputExt {
		outputFile = outputFile + jsonOutputExt
	}
	jp.writer = printer.GetWriter(outputFile)
}

func (jp *JsonPrinter) Score(score float32) {
	fmt.Fprintf(os.Stderr, "\nOverall risk-score (0- Excellent, 100- All failed): %d\n", cautils.Float32ToInt(score))
}

func (jp *JsonPrinter) ActionPrint(opaSessionObj *cautils.OPASessionObj) {
	r, err := json.Marshal(FinalizeResults(opaSessionObj))
	if err != nil {
		logger.L().Fatal("failed to Marshal posture report object")
	}

	if _, err := jp.writer.Write(r); err != nil {
		logger.L().Error("failed to write results", helpers.Error(err))
	} else {
		printer.LogOutputFile(jp.writer.Name())
	}
}
