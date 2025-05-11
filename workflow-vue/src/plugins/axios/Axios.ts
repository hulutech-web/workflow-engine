import useRuleStore from '@/store/useRulesStore'
import router from '@/plugins/router'
import axios from 'axios'
import type { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios'
import { message as Message } from 'ant-design-vue'
import useLoadingStore from '@/store/useLoadingStore'
import { HttpCodeEnum } from '@/enum/HttpCodeEnum'


/**
 * 使用方式
 * import useLoading from '@/store/useLoadingStore'
 * //获取状态中的loading赋值给useLoading1
 * const loadingStore=useLoading()
 * const setTrans=()=>{
 *   loadingStore.setLoading(true)
 *  //调用状态中的设置方法
 *   setTimeout(()=>{
 *     loadingStore.setLoading(false)
 *   },5000)
 * }
 */
interface IOptions {
  message?: boolean
}
const storage = useStorage()

export default class Axios {
  private instance
  private options: IOptions = { message: true }
  private useLoading: null
  constructor(config: AxiosRequestConfig) {

    this.instance = axios.create(config)
    this.interceptors()
  }

  public async request<T>(config: AxiosRequestConfig, options?: IOptions) {
    this.options = Object.assign(this.options, options ?? {})
    this.useLoading = useLoadingStore()

    return new Promise(async (resolve, reject) => {
      try {
        this.useLoading.setLoading(true)
        const response = await this.instance.request<T>(config)
        resolve(response.data)
      } catch (error) {
        reject(error)
      } finally {
        this.useLoading.setLoading(false)
      }

    }) as Promise<T>
  }

  private interceptors() {
    this.interceptorsRequest()
    this.interceptorsResponse()
  }

  private interceptorsRequest() {
    this.instance.interceptors.request.use(
      (config: InternalAxiosRequestConfig) => {
        config.headers.Accept = 'application/json'
        // config.headers.Authorization = `Bearer ${storage.get(CacheEnum.TOKEN_NAME)}`
        config.headers.Authorization = `${storage.get(CacheEnum.TOKEN_NAME)}`
        return config
      },
      (error: any) => {
        return Promise.reject(error)
      },
    )
  }
  private interceptorsResponse() {
    this.instance.interceptors.response.use(
      (response) => {
        const message = response.data?.message ?? response.data?.message
        if (message && this.options.message) {
          Message.success(message ?? '成功')
        }

        this.options = { message: true }
        return response.data
      },
      (error) => {
        this.options = { message: true }
        const {
          response: { status, data },
        } = error
        const message = error.response.data.message
        switch (status) {
          case HttpCodeEnum.UNAUTHORIZED:
            storage.remove(CacheEnum.TOKEN_NAME)
            router.push({ name: RouteEnum.LOGIN })
            break
          case HttpCodeEnum.UNPROCESSABLE_ENTITY:
            const ruleStore = useRuleStore()
            // 配合antdvue的form表单验证
            ruleStore.setRules(error.response.data.errors ?? error.response.data.errors)
            break
          case HttpCodeEnum.FORBIDDEN:
            Message.error(message ?? '没有操作权限')
            break
          case HttpCodeEnum.NOT_FOUND:
            Message.error('请求资源不存在')
            router.push({ name: RouteEnum.HOME })
            break
          case HttpCodeEnum.INTERNAL_SERVER_ERROR:
            console.log(error)
            if(error.response.data.message){
              Message.error(error.response.data.message)
            }
            if(error.response.data.error){
              Message.error(error.response.data.error)
            }
            if(error.response.data.errors){
              Message.error(error.response.data.errors)
            }
            break;
          case HttpCodeEnum.TOO_MANY_REQUESTS:
            Message('请求过于频繁，请稍候再试')
            break
          default:
            if (message) {
              Message(message ?? '服务器错误')
            }
        }
        return Promise.reject(error)
      },
    )
  }

  public async Upload(url, data) {
    return await this.instance.request({
      method: 'POST',
      url: url,
      headers: {
        'Content-Type': 'multipart/form-data',
      },
      data
    })
  }
}
