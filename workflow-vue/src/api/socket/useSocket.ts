import { http } from '@/plugins/axios';

export default () => {
    const socketCmd = async (data) => {
        return await http.request({
            url: "send_to_system",
            method: "post",
            data: data
        })
    }

    return {
        socketCmd
    }
}

