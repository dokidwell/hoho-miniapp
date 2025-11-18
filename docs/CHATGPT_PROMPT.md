# ChatGPT开发提示词

你好！我需要你帮我完成一个Uni-app小程序项目的前端开发。

## 📋 项目信息

**项目名称**：HOHO数字藏品集换小程序

**技术栈**：
- Uni-app（Vue 3 + Composition API + `<script setup>` 语法）
- 白色主题，极简科技风格（参考OKX和Apple）
- 积分系统使用8位小数精度（DECIMAL(20,8)）

**GitHub仓库**：https://github.com/dokidwell/hoho-miniapp

**重要文档**：
- 开发指南：`/docs/FRONTEND_DEVELOPMENT_GUIDE.md`（必读！包含所有页面的完整代码示例）
- 设计稿参考：`/design/` 目录（仅供参考，可能需要调整）
- 设计稿说明：`/design/README.md`

---

## 🎯 任务要求

请严格按照 `/docs/FRONTEND_DEVELOPMENT_GUIDE.md` 文档完成以下工作：

### 1. 创建项目基础结构

```
frontend/
├── api/
│   ├── config.js          # API端点配置
│   └── request.js         # HTTP请求封装（含JWT认证拦截）
├── utils/
│   └── format.js          # 格式化工具函数
├── components/
│   └── TabBar/
│       └── TabBar.vue     # 底部导航栏组件
├── pages/                 # 所有页面
└── static/                # 静态资源
    ├── icons/             # 图标
    └── images/            # 图片
```

### 2. 创建所有13个页面

**核心页面（5个）**：
1. `pages/index/index.vue` - 首页（作品列表）
2. `pages/jijhuan/index.vue` - 集换中心
3. `pages/profile/index.vue` - 我的（个人中心）
4. `pages/ecology/index.vue` - 生态页面
5. `pages/transparent-ledger/index.vue` - 透明公示

**功能页面（8个）**：
6. `pages/create/index.vue` - 创作（铸造）
7. `pages/asset-detail/index.vue` - 藏品详情
8. `pages/jingtan-assets/index.vue` - 鲸探作品
9. `pages/third-party/index.vue` - 第三方关联
10. `pages/listing-create/index.vue` - 挂售页面
11. `pages/exchange/index.vue` - 兑换页面
12. `pages/login/index.vue` - 登录页面
13. `pages/register/index.vue` - 注册页面

**开发文档中每个页面都包含**：
- 完整的Vue代码示例
- API调用示例
- 数据格式说明
- 样式要点

### 3. 创建TabBar组件

完整代码在文档的"通用组件"章节，包含5个Tab：
- 作品、创作、集换、生态、我的

### 4. 配置pages.json

完整配置在文档中，包括：
- 所有页面路由
- TabBar配置（custom: true）
- 全局样式

### 5. 准备静态资源占位符

文档中列出了所有需要的图标和图片，请创建占位文件：
- 使用简单的纯色图片或文字占位
- 文件名和路径必须与文档一致
- 后续可以替换为实际设计资源

---

## ⚠️ 重要注意事项

### 样式规范
1. **严格使用文档中定义的CSS变量**（颜色、字体、间距、圆角）
2. **白色主题**，不是深色主题
3. **积分显示必须使用等宽字体**：
   ```vue
   <text class="price-value number-display">{{ formatPoints(price) }}</text>
   ```
   ```scss
   .number-display {
     font-family: 'Monaco', 'Menlo', 'Consolas', monospace;
   }
   ```
4. **积分格式化必须精确到8位小数**：
   ```javascript
   function formatPoints(points) {
     return parseFloat(points).toFixed(8)
   }
   ```

### API对接
1. **后端API已100%完成**，前端只需对接即可
2. **API_BASE_URL 先设置为**：`'http://localhost:8080'`（后续会替换为实际服务器地址）
3. **所有请求必须携带JWT Token**（request.js中已实现）
4. **401错误自动跳转登录页**（request.js中已实现）

### 代码质量
1. **代码必须完整可运行**，不要省略任何部分
2. **不要使用占位注释**（如 `// TODO: 实现功能`），必须实现完整功能
3. **使用Vue 3 Composition API** + `<script setup>` 语法
4. **所有数据请求使用async/await**
5. **错误处理必须完善**（使用uni.showToast显示错误）

### 设计稿参考
1. **设计稿在 `/design/` 目录**，仅供参考
2. **以开发文档为准**，设计稿可能需要调整
3. **如有冲突，优先按开发文档实现**
4. **布局、间距、颜色尽量还原设计稿**

---

## 🚀 开发流程

### Step 1: 创建基础结构
```bash
# 创建目录
mkdir -p api utils components/TabBar

# 创建文件
touch api/config.js api/request.js utils/format.js
```

### Step 2: 实现基础模块
1. 实现 `api/config.js`（API端点配置）
2. 实现 `api/request.js`（HTTP请求封装）
3. 实现 `utils/format.js`（格式化工具）
4. 实现 `components/TabBar/TabBar.vue`（底部导航栏）

### Step 3: 配置pages.json
按文档中的完整配置更新 `pages.json`

### Step 4: 逐个创建页面
**建议顺序**：
1. 登录/注册页面（先实现认证）
2. 首页（作品列表）
3. 集换中心
4. 我的（个人中心）
5. 生态页面
6. 透明公示
7. 创作（铸造）
8. 藏品详情
9. 鲸探作品
10. 第三方关联
11. 挂售页面
12. 兑换页面
13. 兑换成功页面

**每个页面都要包含**：
- 完整的模板结构
- 完整的脚本逻辑（数据请求、事件处理）
- 完整的样式（使用CSS变量）
- 加载状态、错误处理、空状态

### Step 5: 准备静态资源
为所有图标和图片创建占位文件（简单的纯色图或文字）

### Step 6: 测试和调试
确保所有页面可以正常编译和运行

---

## 📦 提交代码

完成后请使用以下命令提交到GitHub：

```bash
cd /path/to/hoho-miniapp
git add -A
git commit -m "feat: 完成前端所有页面开发

- 实现13个页面的完整功能
- 创建TabBar组件
- 配置pages.json
- 实现API请求封装和JWT认证
- 准备静态资源占位符
- 严格遵循设计规范和开发文档"
git push origin main
```

---

## 💡 开发建议

1. **先通读开发文档**，了解整体结构和规范
2. **参考设计稿**，理解UI布局和交互
3. **复制文档中的代码示例**，然后根据实际需求调整
4. **每完成一个页面就测试一次**，确保功能正常
5. **遇到问题先查文档**，文档中有详细的API说明和代码示例
6. **保持代码风格一致**，使用相同的命名规范和代码结构

---

## 🔍 关键代码示例

### 积分格式化（8位小数）
```javascript
// utils/format.js
export function formatPoints(points, decimals = 8) {
  if (points === null || points === undefined) return '0.00000000'
  return parseFloat(points).toFixed(decimals)
}
```

### API请求示例
```javascript
// 获取藏品列表
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'

async function fetchAssets() {
  try {
    const res = await request.get(API_ENDPOINTS.ASSET.LIST, {
      page: 1,
      page_size: 20,
      status: 'approved'
    })
    assets.value = res.data.list
  } catch (error) {
    uni.showToast({
      title: error.message || '加载失败',
      icon: 'none'
    })
  }
}
```

### 页面模板示例
```vue
<template>
  <view class="page">
    <!-- 内容区域 -->
    <view class="content">
      <!-- 加载中 -->
      <view v-if="loading" class="loading">加载中...</view>
      
      <!-- 数据列表 -->
      <view v-else-if="list.length > 0" class="list">
        <view v-for="item in list" :key="item.id" class="item">
          {{ item.name }}
        </view>
      </view>
      
      <!-- 空状态 -->
      <view v-else class="empty">暂无数据</view>
    </view>
    
    <!-- 底部导航栏 -->
    <TabBar :active="0" />
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import TabBar from '@/components/TabBar/TabBar.vue'

const loading = ref(false)
const list = ref([])

onMounted(() => {
  fetchData()
})

async function fetchData() {
  loading.value = true
  try {
    // API请求
  } catch (error) {
    uni.showToast({
      title: error.message || '加载失败',
      icon: 'none'
    })
  } finally {
    loading.value = false
  }
}
</script>

<style lang="scss" scoped>
.page {
  min-height: 100vh;
  background-color: var(--color-bg-primary);
}
</style>
```

---

## ✅ 完成检查清单

完成后请确认以下项目：

- [ ] 所有13个页面已创建并实现完整功能
- [ ] TabBar组件已创建并正常工作
- [ ] pages.json已正确配置
- [ ] API请求模块已实现（config.js + request.js）
- [ ] 工具函数已实现（format.js）
- [ ] 所有静态资源占位符已准备
- [ ] 样式符合设计规范（白色主题、CSS变量）
- [ ] 积分显示使用等宽字体且精确到8位小数
- [ ] 所有API调用已实现
- [ ] 错误处理完善（加载状态、错误提示、空状态）
- [ ] 代码已提交到GitHub

---

## 🎉 开始开发吧！

请开始开发，如有任何疑问随时问我。记住：

1. **开发文档是你的最佳参考**，里面有完整的代码示例
2. **设计稿仅供参考**，以开发文档为准
3. **代码必须完整可运行**，不要留TODO
4. **遇到问题先查文档**，文档中有详细说明

祝开发顺利！🚀
