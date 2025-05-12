import { setup as init } from './init'
import { setup as main } from './main'
import { setup as user } from './user'
import { setup as sale } from './sale'
import { setup as china } from './china'

const modules = [init, main, user, sale,china]

export default function startup() {
  modules.map((setup) => setup())
}
