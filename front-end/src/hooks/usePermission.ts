import { useAuthStore } from '@/store'

/** 权限判断 */
export function usePermission() {
  const authStore = useAuthStore()

  function hasPermission(
    permission?: string[],
  ) {
    if (!permission)
      return true

    if (!authStore.userInfo)
      return false
    const permissions = authStore.permissions
    let has = permissions.includes('*')
    if (!has) {
      has = permission?.every(i => permissions.includes(i))
    }
    return has
  }

  return {
    hasPermission,
  }
}
