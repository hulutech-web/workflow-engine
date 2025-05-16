import { defineStore } from "pinia";

export default defineStore("rules", {
  state: () => {
    return {
      rules: {} as Record<string, any>,
    };
  },
  getters: {
    getRule(state) {
      return (name: string) => state.rules[name];
    },
    hasRule(state) {
      return Object.keys(state.rules).length > 0;
    },
  },
  actions: {
    resetRules() {
      this.rules = {};
    },
    //修改对象obec
    setRules(rules: Record<string, string[]>) {
      let validVal = null
      Object.entries(rules).forEach(([key, value]) => {
        console.log(key, value);
        switch (Object.keys(value)[0]) {
          case "len":
            this.rules[key] = {
              type: "array",
              message: Object.values(value).join(","),
              trigger: "blur",
            };
            break;
          case "max":
            // 从value正则匹配出数字
            validVal = value["max"].match(/\d+/)[0]
            this.rules[key] = {
              required: true,
              message: Object.values(value).join(","),
              trigger: "change",
              asyncValidator: (rule, val) => {
                return new Promise((resolve, reject) => {
                  if (val < validVal) {
                     reject(Object.values(value).join(","));
                  } else {
                    resolve();
                  }
                });
              },
            };
            break;
          case "min":
            validVal = value["min"].match(/\d+/)[0]
            this.rules[key] = {
              required: true,
              message: Object.values(value).join(","),
              trigger: "change",
              asyncValidator: (rule, val) => {
                return new Promise((resolve, reject) => {
                  if (val > validVal) {
                     reject(Object.values(value).join(","));
                  } else {
                    resolve();
                  }
                });
              },
            };
            break;
          case "string":
            this.rules[key] = {
              type: "string",
              message: Object.values(value).join(","),
              trigger: "blur",
            };
            break;
          case "uint":
            validVal = value["uint"].match(/\d+/)[0]
            this.rules[key] = {
              required: true,
              message: Object.values(value).join(","),
              trigger: "change",
              asyncValidator: (rule, val) => {
                return new Promise((resolve, reject) => {
                  if (val > 0) {
                     reject(Object.values(value).join(","));
                  } else {
                    resolve();
                  }
                });
              },
            };
            break;
          case "min_len":
            validVal = value["min_len"].match(/\d+/)[0]
            this.rules[key] = {
              required: true,
              message: Object.values(value).join(","),
              trigger: "change",
              asyncValidator: (rule, val) => {
                return new Promise((resolve, reject) => {
                  if (val > 0) {
                     reject(Object.values(value).join(","));
                  } else {
                    resolve();
                  }
                });
              },
            };
            break;
          case "max_len":
            validVal = value["max_len"].match(/\d+/)[0]
            this.rules[key] = {
              required: true,
              message: Object.values(value).join(","),
              trigger: "change",
              asyncValidator: (rule, val) => {
                return new Promise((resolve, reject) => {
                  if (val > 0) {
                     reject(Object.values(value).join(","));
                  } else {
                    resolve();
                  }
                });
              },
            };
            break;
          case "ne":
            validVal = value["ne"].match(/\[\s+\]/)[0]
            this.rules[key] = {
              required: true,
              message: Object.values(value).join(","),
              trigger: "change",
              asyncValidator: (rule, val) => {
                return new Promise((resolve, reject) => {
                  if (val !== validVal) {
                     reject(Object.values(value).join(","));
                  } else {
                    resolve();
                  }
                });
              },
            };
            break;
          case "date":
            this.rules[key] = {
              required:true,
              type: "date",
              message: Object.values(value).join(","),
              trigger: "blur",
            };
            break;
          case "file":
            this.rules[key] = {
              required: true,
              message: Object.values(value).join(","),
              trigger: "blur",
            };
            break;
          case "image":
            this.rules[key] = {
              required: true,
              message: Object.values(value).join(","),
              trigger: "blur",
            };
            break;
          case "number":
            this.rules[key] = {
              type: "number",
              message: Object.values(value).join(","),
              trigger: "blur",
            };
            break;
          case "email":
            this.rules[key] = {
              type: "email",
              message: Object.values(value).join(","),
              trigger: "blur",
            };
            break;
          case "alpha_rule":
            this.rules[key] = {
              type: "string",
              message: Object.values(value).join(","),
              pattern: /^[a-zA-Z]+$/,
              trigger: "blur",
            };
            break;
          default:
            this.rules[key] = {
              required: true,
              message: Object.values(value).join(","),
              trigger: "blur",
            };
        }
        
      });
    },
    clearRule(name: string) {
      if (this.rules[name]) {
        delete this.rules[name];
      }
    },
  },
});
