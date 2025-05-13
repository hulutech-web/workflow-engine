import { request } from '../http'

import router from "@/plugins/router";
import XEUtils from "xe-utils";
const storage = useStorage();

export default () => {
  // 方法

  const updateFlowlink = async (data) => {
    return await request.Post(`flowlink`,data);
  };

  return {
    updateFlowlink
  };
};
