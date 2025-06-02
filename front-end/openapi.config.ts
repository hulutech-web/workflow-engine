import {generateService} from '@umijs/openapi';

generateService({
    schemaPath: 'http://localhost:8080/docs/doc.json',
    serversPath: './src',
    requestImportStatement: `import request from '@/utils/request/http';`,
    requestOptionsType: 'RequestOptions',
    hook: {
        afterOpenApiDataInited(openAPIData) {
            const schemas = openAPIData.components?.schemas;
            if (schemas) {
                Object.values(schemas).forEach((schema) => {
                    if ('$ref' in schema) {
                        return;
                    }
                    if (schema.properties) {
                        Object.values(schema.properties).forEach((prop) => {
                            if ('$ref' in prop) {
                                return;
                            }
                            // 匡正文件上传的参数类型
                            if (prop.format === 'binary') {
                                prop.type = 'object';
                            }
                        });
                    }
                });
            }
            return openAPIData;
        },
        // @ts-ignore
        customFunctionName(operationObject) {
            const { operationId, tags } = operationObject;


            return operationId;
        },
        customType(schemaObject, namespace, defaultGetType) {
            const type = defaultGetType(schemaObject, namespace);
            // 提取出 data 的类型
            const regex = /API\.ResOp & { 'data'\?: (.+); }/;
            return type.replace(regex, '$1');
        },
        customFileNames(operationObject,
                        apiPath,
                        _apiMethod){
            const { tags } = operationObject;
            console.log(tags[0].split(" ")[0])
            if (tags[0] &&  tags[0].split(" ").length<1) {
                console.warn('[Warning] no operationId', apiPath);
                return;
            }
            const controllerName = tags[0].split(" ")[0];
            console.log("customFileNames",controllerName)
            return [controllerName];
        },
        customOptionsDefaultValue(data) {
            const {summary} = data;
            if (!summary) return {};
            if (summary.startsWith('创建') || summary.startsWith('新增')) return {successMsg: '创建成功'};
            if (summary.startsWith('更新')) return {successMsg: '更新成功'};
            if (summary.startsWith('删除')) return {successMsg: '删除成功'};
            return {};
        }
    }
});