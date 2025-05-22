// @ts-ignore
/* eslint-disable */
import request from '@/utils/request/http';

/** 添加菜单 添加菜单 POST /auth/menu/add */
export async function unnamedApi(body: API.MenuAddReq, options?: RequestOptions) {
  return request<API.Response>('/auth/menu/add', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 删除菜单 删除菜单 POST /auth/menu/delete */
export async function unnamedApi2(body: API.IdReq, options?: RequestOptions) {
  return request<API.Response>('/auth/menu/delete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || { successMsg: '删除成功' }),
  });
}

/** 获取菜单详情 获取菜单详情 GET /auth/menu/detail/${param0} */
export async function unnamedApi3(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.unnamedApiParams,
  options?: RequestOptions
) {
  const { id: param0, ...queryParams } = params;
  return request<API.MenuResp>(`/auth/menu/detail/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** 编辑菜单 编辑菜单 POST /auth/menu/edit */
export async function unnamedApi4(body: API.MenuEditReq, options?: RequestOptions) {
  return request<API.Response>('/auth/menu/edit', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 获取菜单列表 获取菜单列表 GET /auth/menu/list */
export async function unnamedApi5(options?: RequestOptions) {
  return request<API.MenuResp[]>('/auth/menu/list', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 获取当前用户的菜单权限 获取当前用户的菜单权限 GET /auth/menu/route */
export async function unnamedApi6(options?: RequestOptions) {
  return request<API.MenuResp>('/auth/menu/route', {
    method: 'GET',
    ...(options || {}),
  });
}
