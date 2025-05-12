import { request } from '../../http'

// 获取所有路由信息
export function fetchAllRoutes() {
  return request.Get<Service.ResponseResult<AppRoute.RowRoute[]>>('/auth/menu/list')
}

// 获取某个路由信息
export function fetchRoute(params?: any) {
  return request.Get<Service.ResponseResult<AppRoute.RowRoute>>(`/auth/menu/detail`, { params })
}

// 新增路由信息
export function addRoute(data: AppRoute.RowRoute) {
  return request.Post<Service.ResponseResult<AppRoute.RowRoute>>(`/auth/menu/add`, data)
}

// 修改路由信息
export function updateRoute(data: AppRoute.RowRoute) {
  return request.Post<Service.ResponseResult<AppRoute.RowRoute>>(`/auth/menu/edit`, data)
}

// 删除路由信息
export function deleteRoute(data?: any) {
  return request.Post<Service.ResponseResult<void>>(`/auth/menu/delete`, data)
}
