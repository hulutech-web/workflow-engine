import { http } from '@/plugins/axios';
import router from '@/plugins/router'
const storage = useStorage()
export default () => {
    const index = async () => {
        return await http.request({
            url: "home",
            method: "get",
        })
    }
    return {
        index
    }
}

