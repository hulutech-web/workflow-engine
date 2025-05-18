import { request } from "@/service/http";
import router from "@/plugins/router";
import XEUtils from "xe-utils";
const storage = useStorage();

export default () => {
  // 方法
  const loadFlowEntryConfig = async (id) => {
    return await request.Get(`flow/${id}/entry`);
  };

  const storeEntry = async (data) => {
    return await request.Post(`entry`, data,);
  };
  const showEntry = async (id) => {
    return await request.Get(`entry/${id}`);
  };
  const getEntryData = async (id) => {
    return await request.Get( `entry/${id}/entrydata`);
  };

  const resendEntry=(id)=>{
    return request.Post(`entry/resend`,{entry_id:id})
  }
  return {
    loadFlowEntryConfig,
    showEntry,
    storeEntry,
    getEntryData,
    resendEntry
  };
};
