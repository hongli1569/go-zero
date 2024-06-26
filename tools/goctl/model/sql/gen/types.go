package gen

import (
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/template"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
)

func genTypes(table Table, methods string, withCache bool) (string, error) {
	fields := table.Fields
	fieldsString, err := genFields(table, fields)
	if err != nil {
		return "", err
	}

	text, err := pathx.LoadTemplate(category, typesTemplateFile, template.Types)
	if err != nil {
		return "", err
	}

	output, err := util.With("types").
		Parse(text).
		Execute(map[string]any{
			"withCache":             withCache,
			"method":                methods,
			"upperStartCamelObject": table.Name.ToCamel(),
			"lowerStartCamelObject": stringx.From(table.Name.ToCamel()).Untitle(),
			"fields":                fieldsString,
			"originalFields":        table.Fields,
			"filterFields":          table.filterColumns,
			"aggregate":             genAggregate(),
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func genAggregate() []*parser.Field {
	return []*parser.Field{
		{
			NameOriginal: "limit",
			Name:         stringx.From("limit"),
			DataType:     "int",
		},
		{
			NameOriginal: "offset",
			Name:         stringx.From("offset"),
			DataType:     "int",
		},
		{
			NameOriginal: "project",
			Name:         stringx.From("project"),
			DataType:     "[]string",
		},

		{
			NameOriginal: "calculateTotal",
			Name:         stringx.From("calculateTotal"),
			DataType:     "bool",
		},
	}
}
