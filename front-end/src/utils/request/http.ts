import { AxiosRequestConfig, AxiosResponse } from 'axios';
import createHttp from '@/utils/request/axiosHttp';
import { isResponse } from '@/types';
import NProgress from 'nprogress';
import { notification } from 'ant-design-vue';

const http = createHttp({
  timeout: 10000,
  baseURL: 'api',
  withCredentials: true,
  xsrfCookieName: 'Authorization',
  xsrfHeaderName: 'Authorization',
    headers: {
      'Content-Type': 'application/json;charset=UTF-8',
    },
});

const isAxiosResponse = (obj: any): obj is AxiosResponse => {
  return typeof obj === 'object' && obj.status && obj.statusText && obj.headers && obj.config;
};

// progress 进度条 -- 开启
http.interceptors.request.use((req: AxiosRequestConfig) => {
  if (!NProgress.isStarted()) {
    NProgress.start();
  }
  return req;
});

// 解析响应结果
http.interceptors.response.use(
  (rep: AxiosResponse<String>) => {
    const { data } = rep;
    if (isResponse(data)) {
        if (data.code === 200) {
            return Promise.resolve(data.data);
        } else if (data.code === 310) {
            notification.error({
                message: data.msg,
                description: data.data,
            });
            return Promise.reject(data);
        } else if (data.code === 500) {
            notification.error({
                message: data.msg,
            });
            console.log(data.data)
            return Promise.reject(data);
        }
    }
    return Promise.reject({ msg: rep.statusText, code: rep.status, data });
  },
  (error) => {
    if (error.response && isAxiosResponse(error.response)) {
      return Promise.reject({
        message: error.response.statusText,
        code: error.response.status,
        data: error.response.data,
      });
    }
    return Promise.reject(error);
  }
);

// progress 进度条 -- 关闭
http.interceptors.response.use(
  (rep) => {
    if (NProgress.isStarted()) {
      NProgress.done();
    }
    return rep;
  },
  (error) => {
    if (NProgress.isStarted()) {
      NProgress.done();
    }
    return error;
  }
);

export default http;

// 泛型方式定义请求函数类型
export const get = <T = any>(
    url: string,
    params?: Record<string, unknown>,
    config?: AxiosRequestConfig
): Promise<AxiosResponse<T>> => {
    return http.request<T>(url, 'GET', params, config);
};

// POST请求使用data负载
export const post = <T = any>(
    url: string,
    data?: unknown,
    config?: AxiosRequestConfig
): Promise<AxiosResponse<T>> => {
    return http.request<T>(url, 'post_json', data, {
        ...config,
        headers: {
            'Content-Type': 'application/json', // 👈 关键配置
            ...config?.headers,
        }
    });
};
