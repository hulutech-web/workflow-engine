import { setup as VXETable } from './vxe-table'
import { setup as ContextMenu } from './vxe-table'

const modules = [VXETable,ContextMenu]

export default function register(app: App) {
  modules.map((setup) => setup(app))
}
