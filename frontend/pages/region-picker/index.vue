<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <text class="navbar-title">选择地区</text>
    </view>

    <!-- Tab切换 -->
    <view class="tabs">
      <view 
        class="tab-item" 
        :class="{ 'active': currentTab === 0 }"
        @click="currentTab = 0"
      >
        <text>{{ selectedProvince || '省份' }}</text>
      </view>
      <view 
        class="tab-item" 
        :class="{ 'active': currentTab === 1 }"
        @click="currentTab = 1"
        v-if="selectedProvince"
      >
        <text>{{ selectedCity || '城市' }}</text>
      </view>
      <view 
        class="tab-item" 
        :class="{ 'active': currentTab === 2 }"
        @click="currentTab = 2"
        v-if="selectedCity"
      >
        <text>{{ selectedDistrict || '区县' }}</text>
      </view>
    </view>

    <!-- 列表内容 -->
    <scroll-view class="list" scroll-y>
      <!-- 省份列表 -->
      <view 
        v-if="currentTab === 0"
        class="list-item"
        v-for="province in provinces"
        :key="province"
        @click="selectProvince(province)"
      >
        <text>{{ province }}</text>
        <text class="check" v-if="selectedProvince === province">✓</text>
      </view>

      <!-- 城市列表 -->
      <view 
        v-if="currentTab === 1"
        class="list-item"
        v-for="city in cities"
        :key="city"
        @click="selectCity(city)"
      >
        <text>{{ city }}</text>
        <text class="check" v-if="selectedCity === city">✓</text>
      </view>

      <!-- 区县列表 -->
      <view 
        v-if="currentTab === 2"
        class="list-item"
        v-for="district in districts"
        :key="district"
        @click="selectDistrict(district)"
      >
        <text>{{ district }}</text>
        <text class="check" v-if="selectedDistrict === district">✓</text>
      </view>
    </scroll-view>
  </view>
</template>

<script setup>
import { ref, computed } from 'vue'
import { onLoad } from '@dcloudio/uni-app'

// 当前Tab
const currentTab = ref(0)

// 选中的地区
const selectedProvince = ref('')
const selectedCity = ref('')
const selectedDistrict = ref('')

// 省份数据（简化版，实际应该从API获取完整数据）
const provinces = ref([
  '北京市', '天津市', '河北省', '山西省', '内蒙古自治区',
  '辽宁省', '吉林省', '黑龙江省', '上海市', '江苏省',
  '浙江省', '安徽省', '福建省', '江西省', '山东省',
  '河南省', '湖北省', '湖南省', '广东省', '广西壮族自治区',
  '海南省', '重庆市', '四川省', '贵州省', '云南省',
  '西藏自治区', '陕西省', '甘肃省', '青海省', '宁夏回族自治区',
  '新疆维吾尔自治区', '台湾省', '香港特别行政区', '澳门特别行政区'
])

// 城市数据（简化版，实际应该根据省份动态加载）
const cityData = {
  '北京市': ['东城区', '西城区', '朝阳区', '丰台区', '石景山区', '海淀区', '门头沟区', '房山区', '通州区', '顺义区', '昌平区', '大兴区', '怀柔区', '平谷区', '密云区', '延庆区'],
  '上海市': ['黄浦区', '徐汇区', '长宁区', '静安区', '普陀区', '虹口区', '杨浦区', '闵行区', '宝山区', '嘉定区', '浦东新区', '金山区', '松江区', '青浦区', '奉贤区', '崇明区'],
  '广东省': ['广州市', '深圳市', '珠海市', '汕头市', '佛山市', '韶关市', '湛江市', '肇庆市', '江门市', '茂名市', '惠州市', '梅州市', '汕尾市', '河源市', '阳江市', '清远市', '东莞市', '中山市', '潮州市', '揭阳市', '云浮市'],
  // 其他省份的城市数据...
}

// 区县数据（简化版）
const districtData = {
  '广州市': ['荔湾区', '越秀区', '海珠区', '天河区', '白云区', '黄埔区', '番禺区', '花都区', '南沙区', '从化区', '增城区'],
  '深圳市': ['罗湖区', '福田区', '南山区', '宝安区', '龙岗区', '盐田区', '龙华区', '坪山区', '光明区', '大鹏新区'],
  // 其他城市的区县数据...
}

// 当前显示的城市列表
const cities = computed(() => {
  if (selectedProvince.value === '北京市' || selectedProvince.value === '上海市') {
    // 直辖市直接返回区县
    return cityData[selectedProvince.value] || []
  }
  return cityData[selectedProvince.value] || []
})

// 当前显示的区县列表
const districts = computed(() => {
  return districtData[selectedCity.value] || []
})

// 页面加载
onLoad((options) => {
  if (options.province) {
    selectedProvince.value = options.province
    currentTab.value = 1
  }
  if (options.city) {
    selectedCity.value = options.city
    currentTab.value = 2
  }
  if (options.district) {
    selectedDistrict.value = options.district
  }
})

// 选择省份
const selectProvince = (province) => {
  selectedProvince.value = province
  selectedCity.value = ''
  selectedDistrict.value = ''
  
  // 直辖市直接跳到区县选择
  if (province === '北京市' || province === '上海市' || province === '天津市' || province === '重庆市') {
    currentTab.value = 1
  } else {
    currentTab.value = 1
  }
}

// 选择城市
const selectCity = (city) => {
  selectedCity.value = city
  selectedDistrict.value = ''
  
  // 如果是直辖市，城市就是区县，直接确认
  if (selectedProvince.value === '北京市' || selectedProvince.value === '上海市' || selectedProvince.value === '天津市' || selectedProvince.value === '重庆市') {
    confirmSelection()
  } else {
    currentTab.value = 2
  }
}

// 选择区县
const selectDistrict = (district) => {
  selectedDistrict.value = district
  confirmSelection()
}

// 确认选择
const confirmSelection = () => {
  const result = {
    province: selectedProvince.value,
    city: selectedCity.value,
    district: selectedDistrict.value
  }
  
  // 触发事件
  uni.$emit('regionSelected', result)
  
  // 返回上一页
  uni.navigateBack()
}
</script>

<style scoped>
.page {
  height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: #fff;
}

.navbar {
  padding: 30rpx;
  text-align: center;
  border-bottom: 1rpx solid #eee;
}

.navbar-title {
  font-size: 36rpx;
  font-weight: 600;
}

.tabs {
  display: flex;
  border-bottom: 1rpx solid #eee;
}

.tab-item {
  flex: 1;
  padding: 24rpx 0;
  text-align: center;
  font-size: 28rpx;
  color: #666;
  position: relative;
}

.tab-item.active {
  color: #007AFF;
  font-weight: 600;
}

.tab-item.active::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 60rpx;
  height: 4rpx;
  background-color: #007AFF;
  border-radius: 2rpx;
}

.list {
  flex: 1;
  overflow-y: auto;
}

.list-item {
  padding: 30rpx;
  border-bottom: 1rpx solid #f0f0f0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 30rpx;
  color: #333;
}

.list-item:active {
  background-color: #f5f5f5;
}

.check {
  color: #007AFF;
  font-size: 32rpx;
  font-weight: bold;
}
</style>
