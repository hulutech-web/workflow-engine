import china from "./china.json"
import * as echarts from 'echarts';
// 注册地图
echarts.registerMap('china', china);
var option = {
    series: [
        {
            type: 'map',
            map: 'china',
            // 尺寸1.5倍
            zoom: 1.2,
            
            // 背景色
            itemStyle: {
                label: {
                    show: true,
                    textStyle: {
                        color: "#fff"
                    }
                },
                // 地图背景颜色
                areaColor: '#031525',
                // 地图边框颜色
                borderColor: '#3B5077',
                // 地图边框宽度
                borderWidth: 1,
                // 地图文字颜色
            },
            inRange: {
                // color: ['#3B5077', '#031525'] // 蓝黑
                // color: ['#ffc0cb', '#800080'] // 红紫
                // color: ['#3C3B3F', '#605C3C'] // 黑绿
                // color: ['#0f0c29', '#302b63', '#24243e'] // 黑紫黑
                // color: ['#23074d', '#cc5333'] // 紫红
                color: ['#00467F', '#A5CC82'] // 蓝绿
                // color: ['#1488CC', '#2B32B2'] // 浅蓝
                // color: ['#00467F', '#A5CC82'] // 蓝绿
                // color: ['#00467F', '#A5CC82'] // 蓝绿
                // color: ['#00467F', '#A5CC82'] // 蓝绿
                // color: ['#00467F', '#A5CC82'] // 蓝绿

            },
            geo: {
                show: true,
                map: "china",
                label: {
                    normal: {
                        show: false
                    },
                    emphasis: {
                        show: false,
                    }
                },
                roam: true,
                itemStyle: {
                    normal: {
                        areaColor: '#031525',
                        borderColor: '#3B5077',
                    },
                    emphasis: {
                        areaColor: '#2B91B7',
                    }
                }
            },
            label: {
                show: true,
                color: '#fff',
                fontSize: 12
            },
            emphasis: {
                label: {
                    color: '#fff',
                    fontSize: 12
                },
                itemStyle: {
                    areaColor: '#2B91B7',
                    borderWidth: 1,
                    borderColor: '#000'
                }
            },
            // 导入的china的json文件
            data: []
        }]
}
const setup = () => {
    let mapele = document.querySelector("#map")
    var mapInstance = echarts.getInstanceByDom(mapele);
    if (mapInstance) {
        console.log("mapInstance")

        echarts.dispose(mapInstance);
    }
    var mapechart = echarts.init(mapele);
    mapechart.setOption(option);
}

export { setup }