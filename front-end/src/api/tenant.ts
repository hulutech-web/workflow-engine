// @ts-ignore
/* eslint-disable */
import request from '@/utils/request/http';

/** 添加租户 添加租户 POST /auth/tenant/add */
export async function tenantAdd(body: API.TenantAddReq, options?: RequestOptions) {
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
export async function tenantAll(options?: RequestOptions) {
  return request<API.PageResp & { data?: API.TenantResp[] }>('/auth/tenant/all', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 删除租户 删除租户 POST /auth/tenant/delete */
export async function tenantDelete(body: API.IdReq, options?: RequestOptions) {
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
export async function tenantDetail(body: API.IdReq, options?: RequestOptions) {
  return request<API.TenantResp>('/auth/tenant/detail', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 编辑租户 编辑租户 POST /auth/tenant/edit */
export async function tenantEdit(body: API.TenantEditReq, options?: RequestOptions) {
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
export async function tenantList(body: API.PageReq, options?: RequestOptions) {
  return request<API.PageResp & { data?: API.TenantResp[] }>('/auth/tenant/list', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
