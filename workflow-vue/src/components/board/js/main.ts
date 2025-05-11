import * as echarts from 'echarts';

var option1 = {
    // 控制提示
    tooltip: {
        // 非轴图形，使用item的意思是放到数据对应图形上触发提示
        trigger: 'item',
        // 格式化提示内容：
        // a 代表图表名称 b 代表数据名称 c 代表数据  d代表  当前数据/总数据的比例
        formatter: "{a} <br/>{b} : {c} ({d}%)"
    },
    // 控制图表
    series: [
        {
            // 图表名称
            name: '地区',
            // 图表类型
            type: 'pie',
            // 南丁格尔玫瑰图 有两个圆  内圆半径10%  外圆半径70%
            // 百分比基于  图表DOM容器的半径
            radius: ['10%', '70%'],
            // 图表中心位置 left 50%  top 50% 距离图表DOM容器
            center: ['50%', '50%'],
            // 半径模式，另外一种是 area 面积模式
            roseType: 'radius',
            // 数据集 value 数据的值 name 数据的名称
            data: [
                { value: 20, name: '云南' },
                { value: 5, name: '北京' },
                { value: 15, name: '山东' },
                { value: 25, name: '河北' },
                { value: 20, name: '江苏' },
                { value: 35, name: '浙江' },
                { value: 30, name: '四川' },
                { value: 40, name: '湖北' }
            ],
            //文字调整
            label: {
                fontSize: 10
            },
            //引导线
            labelLine: {
                length: 8,
                length2: 10
            }
        }
    ],
    color: ['#006cff', '#60cda0', '#ed8884', '#ff9f7f', '#0096ff', '#9fe6b8', '#32c5e9', '#1d9dff']
};


var option2 = {
    series: [
        {
            type: 'pie',
            radius: ['130%', '150%'],  // 放大图形
            center: ['50%', '80%'],    // 往下移动  套住75%文字
            label: {
                show: false,
            },
            startAngle: 180,
            hoverOffset: 0,  // 鼠标经过不变大
            data: [
                {
                    value: 100,
                    itemStyle: { // 颜色渐变#00c9e0->#005fc1
                        color: {
                            type: 'linear',
                            x: 0,
                            y: 0,
                            x2: 0,
                            y2: 1,
                            colorStops: [
                                { offset: 0, color: '#00c9e0' },
                                { offset: 1, color: '#005fc1' }
                            ]
                        }
                    }
                },
                { value: 100, itemStyle: { color: '#12274d' } }, // 颜色#12274d

                { value: 200, itemStyle: { color: 'transparent' } }// 透明隐藏第三块区域
            ]
        }
    ]
};

const setup = () => {
    let el1 = document.querySelector("#pie")
    var existingInstance1 = echarts.getInstanceByDom(el1);
    if (existingInstance1) {
        console.log("existingInstance1")
        echarts.dispose(existingInstance1);
    }
    var myechart1 = echarts.init(el1);
   
    myechart1.setOption(option1);
    myechart1.on('click', function (param) {

    });

    let el2 = document.querySelector("#gauge")
    var existingInstance2 = echarts.getInstanceByDom(el2);
    if (existingInstance2) {
        console.log("existingInstance2")

        echarts.dispose(existingInstance2);
    }
    var myechart2 = echarts.init(el2);

    myechart2.setOption(option2);
    myechart2.on('click', function (param) {
    });
}

export { setup }