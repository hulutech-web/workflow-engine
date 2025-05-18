import { request } from '../../http'

// tenant list
export const fetchTenantList = (params?: any) => {
  return request.Get<Service.ResponseResult<Service.PageResp<Auth.Tenant[]>>>('/auth/tenant/list', { params })
}

// tenant detail
export const fetchTenantDetail = (params?: any) => {
  return request.Get<Service.ResponseResult<Auth.Tenant>>('/auth/tenant/detail', {params})
}

// tenant add
export const addTenant = (data: Auth.TenantReq) => {
  return request.Post<Service.ResponseResult<any>>('/auth/tenant/add', data)
}

// tenant edit
export const editTenant = (data: Auth.TenantReq) => {
  return request.Post<Service.ResponseResult<any>>('/auth/tenant/edit', data)
}

// tenant delete
export const deleteTenant = (data?: any) => {
  return request.Post<Service.ResponseResult<any>>('/auth/tenant/delete', data )
}

