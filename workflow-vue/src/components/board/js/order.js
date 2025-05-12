var data = {
    day365: { orders: '20,301,987', amount: '99834' },
    day90: { orders: '301,987', amount: '9834' },
    day30: { orders: '1,987', amount: '3834' },
    day1: { orders: '987', amount: '834' }
}
//点击事件
$('.order').on('click', '.filter a', function () {
    //点击之后加类名
    $(this).addClass('active').siblings().removeClass('active');
    // 先获取点击a的 data-key自定义属性
    var key = $(this).attr('data-key');
    //获取自定义属性
    // data{}==>data.shuxing data['shuxing]
    key = data[key];//
    $('.order .item h4:eq(0)').text(key.orders);
    $('.order .item h4:eq(1)').text(key.amount);
});
//定时器
var index = 0;
var aclick = $('.order a');
setInterval(function () {
    index++;
    if (index > 3) {
        index = 0;
    }
    //每san秒调用点击事件
    aclick.eq(index).click();
}, 3000);