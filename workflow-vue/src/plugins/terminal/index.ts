import Terminal from 'vue-web-terminal'
import 'vue-web-terminal/lib/theme/dark.css'
const setup = (app: App) => {
    app.use(Terminal)
}
export { setup }
