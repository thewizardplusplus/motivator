package systemutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func UnmarshalJSONFile(filePath string, data interface{}) error {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("unable to read the JSON file: %w", err)
	}

	if err := json.Unmarshal(fileContent, data); err != nil {
		return fmt.Errorf("unable to unmarshal the JSON file: %w", err)
	}

	return nil
}
