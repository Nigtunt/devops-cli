package format

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// OutputFormat 输出格式
type OutputFormat string

const (
	FormatTable OutputFormat = "table"
	FormatJSON  OutputFormat = "json"
	FormatYAML  OutputFormat = "yaml"
)

// ParseFormat 解析格式字符串
func ParseFormat(s string) (OutputFormat, error) {
	switch strings.ToLower(s) {
	case "json":
		return FormatJSON, nil
	case "yaml", "yml":
		return FormatYAML, nil
	case "table", "":
		return FormatTable, nil
	default:
		return "", fmt.Errorf("不支持的格式：%s (支持：table, json, yaml)", s)
	}
}

// Output 根据格式输出数据
func Output(data interface{}, format OutputFormat) error {
	switch format {
	case FormatJSON:
		return outputJSON(data)
	case FormatYAML:
		return outputYAML(data)
	case FormatTable:
		return outputTable(data)
	default:
		return outputTable(data)
	}
}

func outputJSON(data interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func outputYAML(data interface{}) error {
	encoder := yaml.NewEncoder(os.Stdout)
	encoder.SetIndent(2)
	return encoder.Encode(data)
}

func outputTable(data interface{}) error {
	// 简单实现：用 fmt.Printf 打印
	// 实际项目中可以使用 github.com/olekukonko/tablewriter
	fmt.Printf("%+v\n", data)
	return nil
}
