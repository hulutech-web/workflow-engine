import { http } from "@/plugins/axios";
import router from "@/plugins/router";
import XEUtils from "xe-utils";
const storage = useStorage();

export default () => {
  // 方法
  
  const updateFlowlink = async (data) => {
    return await http.request({
      url: `flowlink`,
      method: "POST",
      data: data,
    });
  };

  return {
    updateFlowlink
  };
};
