// @ts-ignore
/* eslint-disable */
import request from '@/utils/request/http';

/** 上传音频 上传音频 POST /upload/audio */
export async function unnamedApi(
  body: {
    /** 分类ID */
    cid: number;
  },
  file?: File,
  options?: RequestOptions
) {
  const formData = new FormData();

  if (file) {
    formData.append('file', file);
  }

  Object.keys(body).forEach((ele) => {
    const item = (body as any)[ele];

    if (item !== undefined && item !== null) {
      if (typeof item === 'object' && !(item instanceof File)) {
        if (item instanceof Array) {
          item.forEach((f) => formData.append(ele, f || ''));
        } else {
          formData.append(ele, JSON.stringify(item));
        }
      } else {
        formData.append(ele, item);
      }
    }
  });

  return request<API.Response & { data?: API.FileResp }>('/upload/audio', {
    method: 'POST',
    data: formData,
    requestType: 'form',
    ...(options || {}),
  });
}

/** 上传文件 上传文件 POST /upload/file */
export async function unnamedApi2(
  body: {
    /** 分类ID */
    cid: number;
  },
  file?: File,
  options?: RequestOptions
) {
  const formData = new FormData();

  if (file) {
    formData.append('file', file);
  }

  Object.keys(body).forEach((ele) => {
    const item = (body as any)[ele];

    if (item !== undefined && item !== null) {
      if (typeof item === 'object' && !(item instanceof File)) {
        if (item instanceof Array) {
          item.forEach((f) => formData.append(ele, f || ''));
        } else {
          formData.append(ele, JSON.stringify(item));
        }
      } else {
        formData.append(ele, item);
      }
    }
  });

  return request<API.Response & { data?: API.FileResp }>('/upload/file', {
    method: 'POST',
    data: formData,
    requestType: 'form',
    ...(options || {}),
  });
}

/** 上传图片 上传图片 POST /upload/image */
export async function unnamedApi3(
  body: {
    /** 分类ID */
    cid: number;
  },
  file?: File,
  options?: RequestOptions
) {
  const formData = new FormData();

  if (file) {
    formData.append('file', file);
  }

  Object.keys(body).forEach((ele) => {
    const item = (body as any)[ele];

    if (item !== undefined && item !== null) {
      if (typeof item === 'object' && !(item instanceof File)) {
        if (item instanceof Array) {
          item.forEach((f) => formData.append(ele, f || ''));
        } else {
          formData.append(ele, JSON.stringify(item));
        }
      } else {
        formData.append(ele, item);
      }
    }
  });

  return request<API.Response & { data?: API.FileResp }>('/upload/image', {
    method: 'POST',
    data: formData,
    requestType: 'form',
    ...(options || {}),
  });
}

/** 上传视频 上传视频 POST /upload/video */
export async function unnamedApi4(
  body: {
    /** 分类ID */
    cid: number;
  },
  file?: File,
  options?: RequestOptions
) {
  const formData = new FormData();

  if (file) {
    formData.append('file', file);
  }

  Object.keys(body).forEach((ele) => {
    const item = (body as any)[ele];

    if (item !== undefined && item !== null) {
      if (typeof item === 'object' && !(item instanceof File)) {
        if (item instanceof Array) {
          item.forEach((f) => formData.append(ele, f || ''));
        } else {
          formData.append(ele, JSON.stringify(item));
        }
      } else {
        formData.append(ele, item);
      }
    }
  });

  return request<API.Response & { data?: API.FileResp }>('/upload/video', {
    method: 'POST',
    data: formData,
    requestType: 'form',
    ...(options || {}),
  });
}
