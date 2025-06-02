swagger 生成文档指令：
```powershell
swag init -g ./boot/boot.go -o ./docs
```

##### 前端 生成命令

路由文件注释说明  
``route/role.go``文件中，函数名称
```go
// @BasePath /api
// @Summary 编辑角色
// @Description 编辑角色
// @Tags Role 角色管理
// @Id RoleEdit
// @Produce  json
// @Param token header string true "access_token"
// @Param request body req.RoleEditReq true "角色信息"
// @Success 200 {object} response.Response "成功"
// @Router /auth/role/edit [post]
```
##### 后端 函数声明
`` // @Tags Role 角色管理 // @Id RoleEdit ``
- 解释  必要字段  @Tags Role 角色管理 ，@Tags 后的第一个单词为前端api的文件名，@Id 为前端api的函数名。相同的@Tags Role 将合并为一个js文件，方便导入
- 示例

##### 前端 导入与调用
```js
import {account} from "@/api"; //导入模块
const {data} = await account.accountLogin(); //调用模块方法
```

```shell
cd front-end
pnpm openapi
```

#### 配置参见
``front-end/openapi.config.ts``