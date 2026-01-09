package output

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/projectdiscovery/utils/structs"
)

// formatJSON formats the output for json based formatting
func (w *StandardWriter) formatJSON(output *Result) ([]byte, error) {
	finalOrdMap, err := structs.FilterStructToMap(*output, nil, w.excludeOutputFields)
	if err != nil {
		return nil, err
	}

	return jsoniter.Marshal(finalOrdMap)
}
