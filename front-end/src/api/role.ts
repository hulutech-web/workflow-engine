// @ts-ignore
/* eslint-disable */
import request from '@/utils/request/http';

/** 添加角色 添加角色 POST /auth/role/add */
export async function roleAdd(body: API.RoleAddReq, options?: RequestOptions) {
  return request<API.Response>('/auth/role/add', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 获取所有角色 获取所有角色 GET /auth/role/all */
export async function roleAll(options?: RequestOptions) {
  return request<API.PageResp & { data?: API.RoleSimpleResp[] }>('/auth/role/all', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 角色状态修改 角色状态修改 POST /auth/role/change */
export async function roleChange(body: API.IdReq, options?: RequestOptions) {
  return request<API.Response>('/auth/role/change', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 删除角色 删除角色 POST /auth/role/delete */
export async function roleDelete(body: API.IdReq, options?: RequestOptions) {
  return request<API.Response>('/auth/role/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || { successMsg: '删除成功' }),
  });
}

/** 获取角色详情 获取角色详情 GET /auth/role/detail/${param0} */
export async function roleDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.RoleDetailParams,
  body: API.IdReq,
  options?: RequestOptions
) {
  const { id: param0, ...queryParams } = params;
  return request<API.RoleResp>(`/auth/role/detail/${param0}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** 编辑角色 编辑角色 POST /auth/role/edit */
export async function roleEdit(body: API.RoleEditReq, options?: RequestOptions) {
  return request<API.Response>('/auth/role/edit', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 获取角色列表 获取角色列表 GET /auth/role/list */
export async function roleList(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.RoleListParams,
  options?: RequestOptions
) {
  return request<API.PageResp & { data?: API.RoleResp[] }>('/auth/role/list', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
