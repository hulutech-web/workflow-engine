import VxeUITable from 'vxe-table'
import 'vxe-table/lib/style.css'
import './mytheme.css'
import VxeUIAll from 'vxe-pc-ui'
import 'vxe-pc-ui/lib/style.css'
import XEUtils from 'xe-utils'
import zhCN from 'vxe-table/es/locale/lang/zh-CN'
import 'vxe-table/styles/cssvar.scss'
import { useMessage } from 'naive-ui'
import CopyText from '@/components/custom/CopyText.vue'
const Message = useMessage()
// VXETable.config({
//     theme: 'dark'
// })

VxeUITable.formats.add('formatNumber', {
    cellFormatMethod({ cellValue }, digits = 5) {
        return XEUtils.commafy(Number(cellValue), { digits: 5 })
    }
})
VxeUITable.setup({
    i18n: (key, args) => XEUtils.toFormatString(XEUtils.get(zhCN, key), args)
})

const indicator = h(CopyText, {
    style: {
        fontSize: '24px',
    },
    spin: true,
});


// 创建自定义的单元格渲染器
VxeUITable.renderer.add("myClipboard", {
    // 默认显示模板
    renderDefault: function (renderOpts: any, params: any) {
        let { row, column } = params
        let { events } = renderOpts
        return h("div", {
            // 设置样式
            style: {
                color: "#1990FF",
                cursor: "pointer",
            },
            props: {
                size: "small",
                title: row[column.property],
            },
            async onClick(event: any) {
                let clipboard = navigator.clipboard || {
                    writeText: (text) => {
                        let copyInput = document.createElement('input');
                        copyInput.value = text;
                        document.body.appendChild(copyInput);
                        copyInput.select();
                        document.execCommand('copy');
                        document.body.removeChild(copyInput);
                    }
                }
                if (clipboard) {
                    await clipboard.writeText(event.target.innerText);
                    Message.success('复制成功')
                }
            }
        }, h("a-icon", {
            props: {
                type: "copy",
            }
        }
            , row[column.property]))
    },

})
const setup = (app: App) => {
    app.use(VxeUITable)
  app.use(VxeUIAll)
}
export { setup }
