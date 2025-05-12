import coordtransform from "coordtransform";
import { reactive } from "vue";
const { indexTrack } = useTrack()
let map = null;

const state = reactive({
    markerLayer: null,
    markers: [],
    infoWindow: null,
});




/**
 * 添加点标记
 * @param device 设备数据
 */
const addImgMarker = (device, tag) => {
    const gcjLatLng = coordtransform.wgs84togcj02(device.coordinates.coordinates[1], device.coordinates.coordinates[0]);
    const position = new TMap.LatLng(gcjLatLng[0], gcjLatLng[1]);

    let equipment = device.Device
    // 自定义图标
    // 过程图标
    const url = import.meta.env.VITE_API_URL
    const progressIconUrl = url + "/icons/charging.png";
    // 开始图标
    const startIconUrl = url + "/icons/start.png";
    // 结束图标
    const overIconUrl = url + "/icons/over.png";
    //初始过程图标
    let marker = null
    console.log(tag)
    if (tag == 'start') {
        marker = new TMap.MultiMarker({
            map: map,
            styles: { //点标注的相关样式
                "marker": new TMap.MarkerStyle({
                    "width": 35,
                    "height": 35,
                    "anchor": { x: 16, y: 32 },
                    "src": startIconUrl
                })
            },
            geometries: [{
                id: device.ID,
                position: position,
                styleId: "marker",
            }],
        });
    }
    if (tag == 'progress') {
        marker = new TMap.MultiMarker({
            map: map,
            styles: { //点标注的相关样式
                "marker": new TMap.MarkerStyle({
                    "width": 35,
                    "height": 35,
                    "anchor": { x: 16, y: 32 },
                    "src": progressIconUrl
                })
            },
            geometries: [{
                id: device.ID,
                position: position,
                styleId: "marker",
            }],
        });
    }
    if (tag == 'over') {
        marker = new TMap.MultiMarker({
            map: map,
            styles: { //点标注的相关样式
                "marker": new TMap.MarkerStyle({
                    "width": 35,
                    "height": 35,
                    "anchor": { x: 16, y: 32 },
                    "src": overIconUrl
                })
            },
            geometries: [{
                id: device.ID,
                position: position,
                styleId: "marker",
            }],
        });
    }
    //处理开始图标


    // 监听标记点击事件
    marker.on('click', () => {
        state.infoWindow.open();
        state.infoWindow.setPosition(position);
        state.infoWindow.setContent(`
      <div style="text-align:left;">
        <p><strong>设备名称:</strong> ${equipment.name}</p>
        <p><strong>设备编码:</strong> ${equipment.code}</p>
        <p><strong>电池状态:</strong> ${equipment.battery_status}</p>
        <p><strong>主板型号:</strong> ${equipment.board_model}</p>
        <p><strong>上线时间:</strong> ${equipment.online_at}</p>
        <p><strong>序列号:</strong> ${equipment.serial_number}</p>
        <p><strong>规格:</strong> ${equipment.spec}</p>
        <p><strong>基站编码:</strong> ${equipment.tower_code}</p>
        <p><strong>更新时间:</strong> ${equipment.UpdatedAt}</p>
        <p><strong>位置:</strong> ${device.address}</p>
      </div>
    `);
    });

    state.markers.push(marker);
};

/**
 * 初始化地图
 * @param centerCoords 初始坐标
 */
const mapInit = (centerCoords) => {
    const TMap = window.TMap;
    const gcjLatLng = coordtransform.wgs84togcj02(centerCoords[1], centerCoords[0]);
    const center = new TMap.LatLng(gcjLatLng[0], gcjLatLng[1]);

    map = new TMap.Map(document.getElementById("mapContainer"), {
        center: center,
        zoom: 14,
        viewMode: "3D",
        scrollWheel: false,
    });

    state.infoWindow = new TMap.InfoWindow({
        map: map,
        position: center,
        offset: { x: 0, y: -32 },
    });
    // 初始关闭信息窗关闭
    state.infoWindow.close();
};

/**
 * 加载设备数据并在地图上显示标记
 * @param devices 设备数组
 */
const loadDeviceData = (devices) => {
    if (devices && devices.length > 0) {
        const initialDevice = devices[0];
        const initialCoords = initialDevice.coordinates.coordinates;

        mapInit(initialCoords);

        devices.forEach((device, index) => {
            if (index == 0) {
                addImgMarker(device, 'start');
            }
            if (index == devices.length - 1) {
                addImgMarker(device, 'over');
            }
            if (index != 0 && index != devices.length - 1) {
                addImgMarker(device, 'progress');
            }
        });
        //初始化轨迹图
        addTrack(devices)
    } else {
        // 如果没有设备数据，可以设置一个默认坐标初始化地图
        mapInit([31.20929040462313, 107.46796970075786]);
    }
};

/**
 * 生成轨迹，从第0个开始到最后一个
 * @param devices 设备数组,MultiPolyline
 */
const addTrack = (devices) => {
    if (devices.length == 0) {
        return
    }
    const TMap = window.TMap;
    //创建 MultiPolyline
    let paths = devices.map(item => {
        const gcjLatLng = coordtransform.wgs84togcj02(item.coordinates.coordinates[1], item.coordinates.coordinates[0]);
        const map = new TMap.LatLng(gcjLatLng[0], gcjLatLng[1]);
        return map
    })


    // console.log("paths", paths)
    const geometries = [
        {//第1条线
            'id': 'pl_1',//折线唯一标识，删除时使用
            'styleId': 'style_blue',//绑定样式名
            'paths': paths
        },
    ]

    var polylineLayer = new TMap.MultiPolyline({
        id: 'polyline-layer', //图层唯一标识
        map: map,//设置折线图层显示到哪个地图实例中
        //折线样式定义
        styles: {
            'style_blue': new TMap.PolylineStyle({
                'color': '#3777FF', //线填充色
                'width': 6, //折线宽度
                'borderWidth': 5, //边线宽度
                'borderColor': '#FFF', //边线颜色
                'lineCap': 'butt' //线端头方式
            }),
            'style_red': new TMap.PolylineStyle({
                'color': '#CC0000', //线填充色
                'width': 6, //折线宽度
                'borderWidth': 5, //边线宽度
                'borderColor': '#CCC', //边线颜色
                'lineCap': 'round' //线端头方式
            })
        },
        //折线数据定义
        geometries: geometries
    });
}

export const loadDataAndRender = async (id, startTime, endTime) => {
    try {
        const { data, distance } = await indexTrack(id, startTime, endTime); // 获取数据
        if (data.length > 0) {
            loadDeviceData(data);
            return {
                len: data.length,
                distance
            }
        } else {
            return {
                len: data.length,
                distance
            }
        }

    } catch (error) {
        console.error("加载设备数据失败:", error);
    }
};

export default mapInit;
