import { request } from '../../http'

// role list
export const fetchRoleList = (params?: any) => {
  return request.Get<Service.ResponseResult<Service.PageResp<Auth.Role[]>>>('/auth/role/list', { params })
}

// role detail
export const fetchRoleDetail = (params?: any) => {
  return request.Get<Service.ResponseResult<Auth.Role>>('/auth/role/detail', {params})
}

// role add
export const addRole = (data: Auth.Role) => {
  return request.Post<Service.ResponseResult<any>>('/auth/role/add', data)
}

// role edit
export const editRole = (data: Auth.Role) => {
  return request.Post<Service.ResponseResult<any>>('/auth/role/edit', data)
}

// role delete
export const deleteRole = (data?: any) => {
  return request.Delete<Service.ResponseResult<any>>('/auth/role/delete', data )
}

// role all
export const fetchRoleAll = () => {
  return request.Get<Service.ResponseResult<Auth.RoleSimple[]>>('/auth/role/all' )}

// role change
export const changeRole = (data: any) => {
  return request.Post<Service.ResponseResult<any>>('/auth/role/change', data)
}
