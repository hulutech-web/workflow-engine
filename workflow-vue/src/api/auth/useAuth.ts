import { http } from '@/plugins/axios';
import router from '@/plugins/router'
const storage = useStorage()
export default () => {
    const login = async (data) => {
        return await http.request({
            url: "auth/login",
            method: "post",
            data: data
        }).then(res => {
            const { data } = res
            console.log(data)

            storage.set(CacheEnum.TOKEN_NAME, data.token)
            storage.set(CacheEnum.USER, JSON.stringify(data.user))
            const route = router.resolve({ name: RouteEnum.ADMIN })
            location.href = "/#/admin/dashboard"
        })
    }

    //登录检测
    function isLogin(): boolean {
        return useStorage().get(CacheEnum.TOKEN_NAME)
    }
    return {
        login,
        isLogin
    }
}

