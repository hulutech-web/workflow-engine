// @ts-ignore
/* eslint-disable */
import request from '@/utils/request/http';

/** 部门 部门分页 GET /dept */
export async function deptIndex(options?: RequestOptions) {
  return request<API.PageResult>('/dept', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 部门 新增部门 POST /dept */
export async function deptStore(body: API.Dept, options?: RequestOptions) {
  return request<API.Response & { data?: Record<string, any> }>('/dept', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 部门 单个部门 GET /dept/${param0} */
export async function deptShow(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.DeptShowParams,
  options?: RequestOptions
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Response & { data?: Record<string, any> }>(`/dept/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** 部门 新增部门 PUT /dept/${param0} */
export async function deptUpdate(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.DeptUpdateParams,
  body: API.Dept,
  options?: RequestOptions
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Response & { data?: Record<string, any> }>(`/dept/${param0}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** 部门 新增部门 DELETE /dept/${param0} */
export async function deptDestroy(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.DeptDestroyParams,
  options?: RequestOptions
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Response & { data?: Record<string, any> }>(`/dept/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** 部门 新增部门 GET /dept/${param0}/tree */
export async function deptDisplayTree(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.DeptDisplayTreeParams,
  options?: RequestOptions
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Response & { data?: Record<string, any> }>(`/dept/${param0}/tree`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** 部门 新增部门 POST /dept/bind_director */
export async function deptBindDirector(body: API.BindDirectorReq, options?: RequestOptions) {
  return request<API.Response & { data?: Record<string, any> }>('/dept/bind_director', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 部门 新增部门 POST /dept/bind_manager */
export async function deptBindManager(body: API.BindManagerReq, options?: RequestOptions) {
  return request<API.Response & { data?: Record<string, any> }>('/dept/bind_manager', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 部门 部门列表 GET /dept/list */
export async function deptList(options?: RequestOptions) {
  return request<API.Response & { data?: Record<string, any> }>('/dept/list', {
    method: 'GET',
    ...(options || {}),
  });
}
