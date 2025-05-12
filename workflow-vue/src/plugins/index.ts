import { setup as axios } from './axios'
import { setup as dayjs } from './dayjs'
import { setup as antdvue } from './antdvue'
import { setup as tailwindcss } from './tailwindcss'
import { setup as router } from './router'
import { setup as pinia } from './pinia'
import { setup as antIcons } from './antIcons'
import { setup as loading } from './provider'
import { setup as vue3Lottie } from './vue3Lottie'
import { setup as Cropper } from './cropper'
import { setup as VXETable } from './vxe-table'
import { setup as Terminal } from './terminal'
import { setup as MenuContext } from './menucontext'
import { setup as Gocaptchavue } from './gocaptchavue'

const modules = [axios, dayjs, antdvue, tailwindcss, router,VXETable, pinia,antIcons,loading,vue3Lottie,Cropper,Terminal,MenuContext,Gocaptchavue]

export default function register(app: App) {
  modules.map((setup) => setup(app))
}
