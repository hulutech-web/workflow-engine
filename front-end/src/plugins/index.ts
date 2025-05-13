import { setup as VXETable } from './vxe-table'

const modules = [VXETable]

export default function register(app: App) {
  modules.map((setup) => setup(app))
}
