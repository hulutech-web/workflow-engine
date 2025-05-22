// @ts-ignore
/* eslint-disable */
import request from '@/utils/request/http';

/** 用户登出 GET /account/logout */
export async function logoutService(options?: RequestOptions) {
  return request<API.Response>('/account/logout', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 用户注册 POST /account/register */
export async function registerService(body: API.AccountRegisterReq, options?: RequestOptions) {
  return request<API.Response & { data?: API.AccountLoginResp }>('/account/register', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 租户列表 GET /account/tenant */
export async function tenantService(options?: RequestOptions) {
  return request<API.Response & { data?: API.SelectOption[] }>('/account/tenant', {
    method: 'GET',
    ...(options || {}),
  });
}
