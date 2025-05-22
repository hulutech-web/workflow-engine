// @ts-ignore
/* eslint-disable */
import request from '@/utils/request/http';

/** 添加用户 添加用户相关接口 POST /user/add */
export async function unnamedApi(body: API.UserAddReq, options?: RequestOptions) {
  return request<API.Response>('/user/add', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 删除用户 删除用户相关接口 POST /user/delete/${param0} */
export async function unnamedApi2(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.unnamedApiParams,
  body: API.IdReq,
  options?: RequestOptions
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Response>(`/user/delete/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || { successMsg: '删除成功' }),
  });
}

/** 用户详情 用户详情相关接口 GET /user/detail/${param0} */
export async function unnamedApi3(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.unnamedApiParams,
  options?: RequestOptions
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Response & { data?: API.UserResp }>(`/user/detail/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** 禁用用户 禁用用户相关接口 POST /user/disable */
export async function unnamedApi4(body: API.IdReq, options?: RequestOptions) {
  return request<API.Response>('/user/disable', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 编辑用户 编辑用户相关接口 POST /user/edit */
export async function unnamedApi5(body: API.UserEditReq, options?: RequestOptions) {
  return request<API.Response>('/user/edit', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 用户列表 用户列表相关接口 GET /user/list */
export async function unnamedApi6(options?: RequestOptions) {
  return request<API.Response & { data?: API.PageResp }>('/user/list', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 用户权限 用户权限相关接口 GET /user/self */
export async function unnamedApi7(options?: RequestOptions) {
  return request<API.Response & { data?: API.UserSelfResp }>('/user/self', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 更新用户 更新用户相关接口 POST /user/update */
export async function unnamedApi8(body: API.UserUpdateReq, options?: RequestOptions) {
  return request<API.Response>('/user/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || { successMsg: '更新成功' }),
  });
}
