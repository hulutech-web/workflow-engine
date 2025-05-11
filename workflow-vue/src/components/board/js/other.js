var data = [
    { name: '可爱多', num: '9,086' },
    { name: '娃哈哈', num: '8,341' },
    { name: '喜之郎', num: '7,407' },
    { name: '八喜', num: '6,080' },
    { name: '小洋人', num: '6,724' },
    { name: '好多鱼', num: '2,170' },
]
$('.inner').on('mouseenter', '.sup li', function () {
    $(this).addClass('active').siblings().removeClass('active');
    //获取随机的值  sort方法 是给数组排序 a-b是从小到大
    //.5-随机0-1的数 可能为正可能为负 排序就会随机
    var radomData = data.sort(function (a, b) { return 0.5 - Math.random() });
    var html = '';
    radomData.forEach(function (item) {
        html += `<li><span>${item.name}</span><span>${item.num} <s class="icon-up"></s></span></li>`;
    });
    //渲染
    $('.sub').html(html);
});
$('.province .sup li').eq(0).mouseenter();
var index = 0;
var timer = setInterval(() => {
    index++;
    if (index > 5) {
        index = 0;
    }
    $('.sup li').eq(index).mouseenter();
}, 2000);