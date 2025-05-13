import { request } from '../http'
// 获取所有用户信息
export function fetchUserPage() {
  return request.Get<Service.ResponseResult<Entity.User[]>>('/userPage')
}

/**
 * 请求获取字典列表
 *
 * @param code - 字典编码，用于筛选特定的字典列表
 * @returns 返回的字典列表数据
 */
export function fetchDictList(code?: string) {
  const params = { code }
  return request.Get<Service.ResponseResult<Entity.Dict[]>>('/dict/list', { params })
}
