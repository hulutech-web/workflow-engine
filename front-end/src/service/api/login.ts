import { request } from '../http'

interface Ilogin {
  username: string
  password: string
  tenantId: number | null
}

export function fetchLogin(data: Ilogin) {
  const methodInstance = request.Post<Service.ResponseResult<Api.Login.Info>>('/account/login', data)
  methodInstance.meta = {
    authRole: null,
  }
  return methodInstance
}
export function fetchUpdateToken(data: any) {
  const method = request.Post<Service.ResponseResult<Api.Login.Info>>('/common/updateToken', data)
  method.meta = {
    authRole: 'refreshToken',
  }
  return method
}

export function fetchUserRoutes(params: { id: number }) {
  return request.Get<Service.ResponseResult<AppRoute.RowRoute[]>>('/auth/menu/route', { params })
}

export function fetchTenantSelect() {
  return request.Get<Service.ResponseResult<Base.SelectOption[]>>('/account/tenant')
}

// self
export function fetchUserInfo() {
  return request.Get<Service.ResponseResult<Api.Login.Permission>>('/user/self')
}
