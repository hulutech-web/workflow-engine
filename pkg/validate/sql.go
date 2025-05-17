package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/xwb1989/sqlparser"
	"strings"
)

type SQL struct{}

func (S SQL) Tag() string {
	return "sql"
}

func (S SQL) Func() validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().String()

		// 尝试解析SQL语句，验证语法正确性
		stmt, err := sqlparser.Parse(value)
		if err != nil {
			return false
		}

		// 仅允许SELECT查询语句
		if _, ok := stmt.(*sqlparser.Select); !ok {
			return false
		}

		// 检查是否存在字符串字面量，强制参数化
		var hasStringLiteral bool
		sqlparser.Walk(func(node sqlparser.SQLNode) (kontinue bool, err error) {
			if val, ok := node.(*sqlparser.SQLVal); ok && val.Type == sqlparser.StrVal {
				hasStringLiteral = true
				return false, nil // 发现字面量后终止遍历
			}
			return true, nil
		}, stmt)

		if hasStringLiteral {
			return false
		}

		// 检查常见注入模式（如永真条件）
		if hasInjectionPattern(value) {
			return false
		}

		return true
	}
}

// 检测常见的SQL注入模式
func hasInjectionPattern(sql string) bool {
	// 定义常见的注入关键字和模式
	patterns := []string{
		"1=1",
		"' OR '1'='1",
		";--",
		"/*",
		"*/",
		"UNION",
		"DROP",
		"DELETE",
		"EXEC",
	}

	for _, pattern := range patterns {
		if strings.Contains(strings.ToUpper(sql), strings.ToUpper(pattern)) {
			return true
		}
	}
	return false
}

func (S SQL) Message() string {
	return "{0}不是有效的SQL语句"
}

func init() {
	Register(SQL{})
}
