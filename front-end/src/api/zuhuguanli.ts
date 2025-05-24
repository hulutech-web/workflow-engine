// @ts-ignore
/* eslint-disable */
import request from '@/utils/request/http';

/** 添加租户 添加租户 POST /auth/tenant/add */
export async function unnamedApi(body: API.TenantAddReq, options?: RequestOptions) {
  return request<API.Response>('/auth/tenant/add', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 获取所有租户 获取所有租户 GET /auth/tenant/all */
export async function unnamedApi2(options?: RequestOptions) {
  return request<API.PageResp & { data?: API.TenantResp[] }>('/auth/tenant/all', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 删除租户 删除租户 POST /auth/tenant/delete */
export async function unnamedApi3(body: API.IdReq, options?: RequestOptions) {
  return request<API.Response>('/auth/tenant/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || { successMsg: '删除成功' }),
  });
}

/** 获取租户详情 获取租户详情 GET /auth/tenant/detail */
export async function unnamedApi4(options?: RequestOptions) {
  return request<API.TenantResp>('/auth/tenant/detail', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 编辑租户 编辑租户 POST /auth/tenant/edit */
export async function unnamedApi5(body: API.TenantEditReq, options?: RequestOptions) {
  return request<API.Response>('/auth/tenant/edit', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 获取租户列表 获取租户列表 POST /auth/tenant/list */
export async function unnamedApi6(body: API.PageReq, options?: RequestOptions) {
  return request<API.PageResp & { data?: API.TenantResp[] }>('/auth/tenant/list', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
