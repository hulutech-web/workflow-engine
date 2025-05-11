import "go-captcha-vue/dist/style.css"
import GoCaptcha from "go-captcha-vue"

const setup = (app: App) => {
    app.use(GoCaptcha)
}
export { setup }
