import { http } from "@/plugins/axios";
import router from "@/plugins/router";
import XEUtils from "xe-utils";
const storage = useStorage();

export default () => {
  const indexProcs=(entry_id)=>{
    return http.request({
      url: `proc/index?entry_id=${entry_id}`,
      method: "GET",
    });
  }
  // 方法
  const setPass = async (data) => {
    return await http.request({
      url: `pass`,
      method: "POST",
      data:data
    });
  };

  const setUnPass = async (data) => {
    return await http.request({
      url: `unpass`,
      method: "POST",
      data: data,
    });
  };
  
  return {
    setPass,
    setUnPass,
    indexProcs
  };
};
