//括号

interface Expression {
  id: number;
  index: number;
  field: string;
  operator: string;
  value: string;
  extra: string;
}
import { message } from "ant-design-vue";
function isMySQLCondition(expression) {
  // 简化并修正后的正则表达式尝试，以匹配给定的示例
  const pattern =
    /^(?:[a-zA-Z0-9_]+(?:\s*(?:>|<|=|<>|!=|>=|<=)\s*(?:"[^"]*"|'[^']*'|[\u4e00-\u9fa5]+|\d+|NULL))\s*(?:AND|OR|NOT)?)+$/i;

  return pattern.test(expression.trim());
}

// Expression的数组类型
const ExplainConditionSql = (expressions: Expression[]) => {
  // expressions，中的前几个字段都必须有值
  let validFieldKey = ["id", "index", "field", "operator", "value"];
  if (expressions.length > 1) {
    let newExprs = expressions.slice(0, expressions.length - 1);
    // console.log("expressions", expressions, "newExprs", newExprs);
    let isComplete = true;
    isComplete = newExprs.some((item) => !!item.extra);
    if (!isComplete) {
      return {
        success: false,
        msg: "多条件时，需要添加连接条件",
      };
    } else {
      let isAllFieldValid = expressions.every((item) => {
        return validFieldKey.every((key) => item[key] !== "");
      });
      if (!isAllFieldValid) {
        return {
          success: false,
          msg: "字段不能为空",
        };
      }
    }
  }
  if (expressions.length == 1) {
    let isAllFieldValid = expressions.every((item) => {
      return validFieldKey.every((key) => item[key] !== "");
    });
    if (expressions[0]["extra"]) {
      return {
        success: false,
        msg: "单条件时不需要连接条件",
      };
    }
    if (!isAllFieldValid) {
      return {
        success: false,
        msg: "字段不能为空",
      };
    }
  }
  return true;
};

export { ExplainConditionSql };
