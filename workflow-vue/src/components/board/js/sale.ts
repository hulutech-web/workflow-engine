import * as echarts from 'echarts';

var option = {
    //鼠标提示工具
    tooltip: {
        trigger: 'axis'
    },
    xAxis: {
        // 类目类型                                  
        type: 'category',
        // x轴刻度文字                                  
        data: ['1月', '2月', '3月', '4月', '5月', '6月', '7月', '8月', '9月', '10月', '11月', '12月'],
        axisTick: {
            show: false//去除刻度线
        },
        axisLabel: {
            color: '#4c9bfd'//文本颜色
        },
        axisLine: {
            show: false//去除轴线  
        },
        boundaryGap: false//去除轴内间距
    },
    yAxis: {
        // 数据作为刻度文字                                  
        type: 'value',
        axisTick: {
            show: false//去除刻度线
        },
        axisLabel: {
            color: '#4c9bfd'//文本颜色
        },
        axisLine: {
            show: false//去除轴线  
        },
        boundaryGap: false//去除轴内间距
    },
    //图例组件
    legend: {
        textStyle: {
            color: '#4c9bfd' // 图例文字颜色

        },
        right: '10%'//距离右边10%
    },
    // 设置网格样式
    grid: {
        show: true,// 显示边框
        top: '20%',
        left: '3%',
        right: '4%',
        bottom: '3%',
        borderColor: '#012f4a',// 边框颜色
        containLabel: true // 包含刻度文字在内
    },
    series: [{
        name: '预期销售额',
        // 数据                                  
        data: [24, 40, 101, 134, 90, 230, 210, 230, 120, 230, 210, 120],
        // 图表类型                                  
        type: 'line',
        // 圆滑连接                                  
        smooth: true,
        itemStyle: {
            color: '#00f2f1'  // 线颜色
        }
    },
    {
        name: '实际销售额',
        // 数据                                  
        data: [40, 64, 191, 324, 290, 330, 310, 213, 180, 200, 180, 79],
        // 图表类型                                  
        type: 'line',
        // 圆滑连接                                  
        smooth: true,
        itemStyle: {
            color: '#ed3f35'  // 线颜色
        }
    }]
};


const setup = () => {
    let ele = document.querySelector("#line")
    var lineInstance = echarts.getInstanceByDom(ele);
    if (lineInstance) {
        console.log("lineInstance")

        echarts.dispose(lineInstance);
    }
    var lineechart = echarts.init(ele);
    lineechart.setOption(option);
}
export  {setup}