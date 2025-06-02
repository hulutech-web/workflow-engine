// @ts-ignore
/* eslint-disable */
import request from '@/utils/request/http';

/** 附件分类添加 附件分类添加 POST /common/file/cateAdd */
export async function cateAdd(body: API.FileCateAddReq, options?: RequestOptions) {
  return request<API.Response>('/common/file/cateAdd', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 附件分类删除 附件分类删除 POST /common/file/cateDelete */
export async function cateDelete(body: API.IdListReq, options?: RequestOptions) {
  return request<API.Response>('/common/file/cateDelete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 附件分类列表 获取附件分类列表 GET /common/file/cateList */
export async function cateList(body: API.FileCateListReq, options?: RequestOptions) {
  return request<API.Response & { data?: any[] }>('/common/file/cateList', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 附件分类重命名 附件分类重命名 POST /common/file/cateRename */
export async function cateRename(body: API.FileCateRenameReq, options?: RequestOptions) {
  return request<API.Response>('/common/file/cateRename', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 附件删除 附件删除 POST /common/file/fileDelete */
export async function fileDelete(body: API.IdListReq, options?: RequestOptions) {
  return request<API.Response>('/common/file/fileDelete', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 附件列表 获取附件列表 GET /common/file/fileList */
export async function fileList(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.FileListParams,
  body: API.FileListReq,
  options?: RequestOptions
) {
  return request<
    API.Response & {
      data?: API.PageResp & { Count?: number; Lists?: API.FileResp[]; PageNo?: number; PageSize?: number };
    }
  >('/common/file/fileList', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    params: {
      ...params,
    },
    data: body,
    ...(options || {}),
  });
}

/** 附件移动 附件移动 POST /common/file/fileMove */
export async function fileMove(body: API.FileMoveReq, options?: RequestOptions) {
  return request<API.Response>('/common/file/fileMove', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 附件重命名 附件重命名 POST /common/file/fileRename */
export async function fileRename(body: API.FileRenameReq, options?: RequestOptions) {
  return request<API.Response>('/common/file/fileRename', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
