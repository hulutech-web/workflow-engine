import {generateService} from '@umijs/openapi';

generateService({
    schemaPath: 'http://localhost:8080/docs/doc.json',
    serversPath: './src',
    requestImportStatement: `import request from '@/utils/request/http';`,
    requestOptionsType: 'RequestOptions',
    hook: {
        afterOpenApiDataInited(openAPIData) {
            // 深度清理所有 GET 方法的 requestBody
            Object.entries(openAPIData.paths || {}).forEach(([path, pathItem]) => {
                Object.entries(pathItem || {}).forEach(([method, operation]) => {
                    const methodLower = method.toLowerCase();
                    if (methodLower === 'get' && operation.requestBody) {
                        console.warn(`[安全修复] 移除 ${method} ${path} 的非法请求体`);
                        delete operation.requestBody;
                        // 深度删除 content 相关定义
                        if (operation.parameters) {
                            operation.parameters = operation.parameters.filter(
                                (p: any) => p.in !== 'body'
                            );
                        }
                    }
                });
            });

            // 处理二进制上传参数（保留原有逻辑）
            const schemas = openAPIData.components?.schemas;
            if (schemas) {
                Object.values(schemas).forEach((schema) => {
                    if ('$ref' in schema || !schema.properties) return;
                    Object.values(schema.properties).forEach((prop) => {
                        if ('$ref' in prop) return;
                        if (prop.format === 'binary') {
                            prop.type = 'object';
                            prop.format = undefined;
                        }
                    });
                });
            }
            return openAPIData;
        },

        customFunctionName(operationObject) {
            console.log('[openapi]', operationObject)
            const { operationId, tags } = operationObject;
            if (!operationId || !tags?.[0]) return 'unnamedApi';
            return `${operationId}Service`;
        },

        customOptionsDefaultValue(data) {
            const { summary } = data;
            if (!summary) return {};
            if (summary.startsWith('创建') || summary.startsWith('新增')) return { successMsg: '创建成功' };
            if (summary.startsWith('更新')) return { successMsg: '更新成功' };
            if (summary.startsWith('删除')) return { successMsg: '删除成功' };
            return {};
        }
    }
});