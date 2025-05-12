/** 不同请求服务的环境配置 */
export const serviceConfig: Record<ServiceEnvType, Record<string, string>> = {
  dev: {
    url: 'http://127.0.0.1:8080/api',
  },
  prod: {
    url: 'http://127.0.0.1:8080/api',
  },
  test: {
    url: 'http://127.0.0.1:8080/api',
  }
}
