import coordtransform from "coordtransform";
import { reactive } from "vue";
const { indexDevice } = useDevice()
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
const addImgMarker = (device) => {
  const gcjLatLng = coordtransform.wgs84togcj02(device.coordinates.coordinates[1], device.coordinates.coordinates[0]);
  const position = new TMap.LatLng(gcjLatLng[0], gcjLatLng[1]);
  const url = import.meta.env.VITE_API_URL
  const progressIconUrl = url + "/icons/charging2.png";
  let marker = new TMap.MultiMarker({
    map: map,
    styles: { //点标注的相关样式
        "marker": new TMap.MarkerStyle({
            "width": 35,
            "height": 40,
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
  function handleButtonClick(deviceId, actionType) {
    console.log(`Button for device ${deviceId} with action "${actionType}" clicked.`);
    // 在这里根据按钮类型执行相应操作
  }
  // 监听标记点击事件
  marker.on('click', () => {
    state.infoWindow.open();
    state.infoWindow.setPosition(position);
    state.infoWindow.setContent(`
      <div style='text-align:left;' id="infoWindowContent">
        <p><strong>设备名称:</strong> ${device.name}</p>
        <p><strong>设备编码:</strong> ${device.code}</p>
        <p><strong>电池状态:</strong> ${device.battery_status}</p>
        <p><strong>主板型号:</strong> ${device.board_model}</p>
        <p><strong>上线时间:</strong> ${device.online_at}</p>
        <p><strong>序列号:</strong> ${device.serial_number}</p>
        <p><strong>规格:</strong> ${device.spec}</p>
        <p><strong>基站编码:</strong> ${device.tower_code}</p>
        <p><strong>更新时间:</strong> ${device.UpdatedAt}</p>
        <p><strong>位置:</strong> ${device.address}</p>
        <button type="primary" data-set="${device.ID}-historyTrack"
        style='color: #fff;background-color: #1677ff;box-shadow: 0 2px 0 rgba(5, 145, 255, 0.1);'
        >历史轨迹</button>
        <button type="primary" data-set="${device.ID}-detail"
        style='color: #fff;background-color: #1677ff;box-shadow: 0 2px 0 rgba(5, 145, 255, 0.1);'
        >详情</button>
        <button type="primary" data-set="${device.ID}-bms"
        style='color: #fff;background-color: #1677ff;box-shadow: 0 2px 0 rgba(5, 145, 255, 0.1);'
        >BMS</button>
        <button type="primary" data-set="${device.ID}-more"
        style='color: #fff;background-color: #1677ff;box-shadow: 0 2px 0 rgba(5, 145, 255, 0.1);'
        >更多</button>
      </div>
    `);
    setTimeout(() => {
      // 获取Infowindow的DOM容器
      let infoWindowDOM = document.querySelector('#infoWindowContent');

      let buttons = infoWindowDOM.querySelectorAll('button[data-set^="' + device.ID + '-"]');
      buttons.forEach(button => {
        button.addEventListener('click', function () {
          let actionType = button.dataset.set.split('-')[1]; // 获取data-set中的动作类型
          handleButtonClick(device.ID, actionType);
        });
      });
    }, 100); // 0毫秒的延迟通常足够让DOM更新
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

  // 初始化infoWindow
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

    devices.forEach(device => {
      addImgMarker(device);
    });
  } else {
    // 如果没有设备数据，可以设置一个默认坐标初始化地图
    mapInit([31.20929040462313, 107.46796970075786]);
  }
};

export const loadDataAndRender = async (startTime, endTime) => {
  try {
    const data = await indexDevice(startTime, endTime); // 获取数据
    loadDeviceData(data);
  } catch (error) {
    console.error("加载设备数据失败:", error);
  }
};

export default mapInit;
