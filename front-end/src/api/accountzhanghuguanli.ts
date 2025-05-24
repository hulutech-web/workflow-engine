// @ts-ignore
/* eslint-disable */
import request from '@/utils/request/http';

/** 用户登录 POST /account/login */
export async function loginService(body: API.AccountLoginReq, options?: RequestOptions) {
  return request<API.Response & { data?: API.AccountLoginResp }>('/account/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
