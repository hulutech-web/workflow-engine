import * as echarts from 'echarts';
const setup = ()=>{
    var setFont = function () {
        // 因为要定义变量可能和别的变量相互冲突，污染，所有用自调用函数
        var html = document.documentElement;// 获取html
        // 获取宽度
        var width = html.clientWidth;
    
        // 判断
        if (width < 1024) width = 1024
        if (width > 1920) width = 1920
        // 设置html的基准值
        var fontSize = width / 80 + 'px';
        // 设置给html
        html.style.fontSize = fontSize;
    }
    setFont();
    // 2、页面改变的时候也需要设置
    // 尺寸改变事件
    window.onresize = function () {
        setFont();
    }
}


export  {
    setup
};
/**
 * 
 *  $('.monitor').on('click', ' a', function () {
    //点击当前的a 加类名 active  他的兄弟删除类名
    $(this).addClass('active').siblings().removeClass('active');
    //获取一一对应的下标 
    var index = $(this).index();
    //选取content 然后狗日对应下标的 显示   当前的兄弟.content隐藏
    $('.content').eq(index).show().siblings('.content').hide();
});
//滚动
//原理：把marquee下面的子盒子都复制一遍 加入到marquee中
//      然后动画向上滚动，滚动到一半重新开始滚动
//因为选取的是两个marquee  所以要遍历
$('.monitor .marquee').each(function (index, dom) {
    //将每个 的所有子级都复制一遍
    var rows = $(dom).children().clone();
    //再将新的到的加入原来的
    $(dom).append(rows);
});
 * 
 */
 //事件委托
