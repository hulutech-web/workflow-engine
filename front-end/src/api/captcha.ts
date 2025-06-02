// @ts-ignore
/* eslint-disable */
import request from '@/utils/request/http';

/** 生成验证码 生成验证码 GET /captcha/get */
export async function captchaGet(options?: RequestOptions) {
  return request<API.Response & { data?: Record<string, any> }>('/captcha/get', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 验证码验证 验证码验证 POST /captcha/validate */
export async function captchaValidate(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.CaptchaValidateParams,
  options?: RequestOptions
) {
  return request<API.Response & { data?: Record<string, any> }>('/captcha/validate', {
    method: 'POST',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
