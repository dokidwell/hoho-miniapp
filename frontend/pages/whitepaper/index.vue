<template>
  <view class="whitepaper-page">
    <!-- 自定义导航栏 -->
    <view class="custom-navbar">
      <view class="navbar-content">
        <view class="back-btn" @click="goBack">
          <text class="back-icon">←</text>
        </view>
        <view class="navbar-title">HOHO Park 白皮书</view>
        <view class="placeholder"></view>
      </view>
    </view>

    <!-- 内容区域 -->
    <scroll-view class="content-scroll" scroll-y>
      <view class="content">
        <!-- 封面 -->
        <view class="cover-section">
          <view class="logo">🎨</view>
          <text class="title">HOHO Park</text>
          <text class="subtitle">数字藏品生态平台</text>
          <text class="version">白皮书 v1.0</text>
        </view>

        <!-- 目录 -->
        <view class="section">
          <text class="section-title">目录</text>
          <view class="toc">
            <text class="toc-item">1. 项目概述</text>
            <text class="toc-item">2. 平台愿景</text>
            <text class="toc-item">3. 核心功能</text>
            <text class="toc-item">4. 技术架构</text>
            <text class="toc-item">5. 积分体系</text>
            <text class="toc-item">6. 发展路线</text>
          </view>
        </view>

        <!-- 项目概述 -->
        <view class="section">
          <text class="section-title">1. 项目概述</text>
          <text class="paragraph">
            HOHO Park是一个创新的数字藏品生态平台，致力于为创作者和收藏者提供一个安全、便捷、有趣的数字资产交流空间。
          </text>
          <text class="paragraph">
            我们相信，每一件数字作品都值得被珍藏，每一位创作者都应该获得应有的回报。通过区块链技术和智能合约，我们让数字资产的所有权变得清晰可追溯。
          </text>
        </view>

        <!-- 平台愿景 -->
        <view class="section">
          <text class="section-title">2. 平台愿景</text>
          <text class="paragraph">
            <text class="highlight">让每个人都能轻松创作、收藏和交易数字藏品</text>
          </text>
          <view class="vision-list">
            <view class="vision-item">
              <text class="vision-icon">🎨</text>
              <view class="vision-content">
                <text class="vision-title">赋能创作者</text>
                <text class="vision-desc">为创作者提供便捷的铸造工具和收益渠道</text>
              </view>
            </view>
            <view class="vision-item">
              <text class="vision-icon">💎</text>
              <view class="vision-content">
                <text class="vision-title">保护收藏者</text>
                <text class="vision-desc">确保每件藏品的真实性和唯一性</text>
              </view>
            </view>
            <view class="vision-item">
              <text class="vision-icon">🔄</text>
              <view class="vision-content">
                <text class="vision-title">促进流通</text>
                <text class="vision-desc">打造活跃的二级市场，让价值自由流动</text>
              </view>
            </view>
          </view>
        </view>

        <!-- 核心功能 -->
        <view class="section">
          <text class="section-title">3. 核心功能</text>
          
          <view class="feature-card">
            <text class="feature-title">3.1 创作铸造</text>
            <text class="paragraph">
              支持图片、视频、音频等多种格式的数字作品铸造。创作者可以设置作品的发行量、价格等参数，平台自动生成唯一的数字证书。
            </text>
          </view>

          <view class="feature-card">
            <text class="feature-title">3.2 集换市场</text>
            <text class="paragraph">
              用户可以在二级市场自由挂售和购买藏品。平台采用积分体系，降低交易门槛，同时保障交易安全。
            </text>
          </view>

          <view class="feature-card">
            <text class="feature-title">3.3 第三方关联</text>
            <text class="paragraph">
              支持与鲸探等主流平台的资产关联，实现跨平台藏品展示和价值互通。
            </text>
          </view>

          <view class="feature-card">
            <text class="feature-title">3.4 社区生态</text>
            <text class="paragraph">
              建立活跃的创作者和收藏者社区，通过空投、活动等方式增强用户粘性。
            </text>
          </view>
        </view>

        <!-- 技术架构 -->
        <view class="section">
          <text class="section-title">4. 技术架构</text>
          <text class="paragraph">
            HOHO Park采用先进的技术架构，确保平台的安全性、稳定性和可扩展性。
          </text>
          <view class="tech-list">
            <view class="tech-item">
              <text class="tech-label">前端：</text>
              <text class="tech-value">uni-app + Vue 3</text>
            </view>
            <view class="tech-item">
              <text class="tech-label">后端：</text>
              <text class="tech-value">Go + Gin + GORM</text>
            </view>
            <view class="tech-item">
              <text class="tech-label">数据库：</text>
              <text class="tech-value">MySQL + Redis</text>
            </view>
            <view class="tech-item">
              <text class="tech-label">存储：</text>
              <text class="tech-value">腾讯云COS</text>
            </view>
            <view class="tech-item">
              <text class="tech-label">安全：</text>
              <text class="tech-value">HTTPS + JWT + 数据加密</text>
            </view>
          </view>
        </view>

        <!-- 积分体系 -->
        <view class="section">
          <text class="section-title">5. 积分体系</text>
          <text class="paragraph">
            HOHO积分是平台的核心流通媒介，用户可以通过多种方式获取积分：
          </text>
          <view class="points-list">
            <view class="points-item">
              <text class="points-icon">✓</text>
              <text class="points-text">注册奖励：新用户注册即送100积分</text>
            </view>
            <view class="points-item">
              <text class="points-icon">✓</text>
              <text class="points-text">每日签到：每天签到可获得10-50积分</text>
            </view>
            <view class="points-item">
              <text class="points-icon">✓</text>
              <text class="points-text">创作奖励：发布作品可获得积分奖励</text>
            </view>
            <view class="points-item">
              <text class="points-icon">✓</text>
              <text class="points-text">交易手续费：平台收取的手续费部分返还给用户</text>
            </view>
            <view class="points-item">
              <text class="points-icon">✓</text>
              <text class="points-text">活动空投：参与平台活动获得积分</text>
            </view>
          </view>
        </view>

        <!-- 发展路线 -->
        <view class="section">
          <text class="section-title">6. 发展路线</text>
          
          <view class="roadmap">
            <view class="roadmap-item">
              <view class="roadmap-time">2024 Q4</view>
              <view class="roadmap-content">
                <text class="roadmap-title">平台上线</text>
                <text class="roadmap-desc">完成基础功能开发，开放用户注册</text>
              </view>
            </view>
            <view class="roadmap-item">
              <text class="roadmap-time">2025 Q1</text>
              <view class="roadmap-content">
                <text class="roadmap-title">生态建设</text>
                <text class="roadmap-desc">引入优质创作者，举办首次空投活动</text>
              </view>
            </view>
            <view class="roadmap-item">
              <text class="roadmap-time">2025 Q2</text>
              <view class="roadmap-content">
                <text class="roadmap-title">功能升级</text>
                <text class="roadmap-desc">支持更多资产类型，优化交易体验</text>
              </view>
            </view>
            <view class="roadmap-item">
              <text class="roadmap-time">2025 Q3+</text>
              <view class="roadmap-content">
                <text class="roadmap-title">生态扩展</text>
                <text class="roadmap-desc">打通更多第三方平台，建立开放生态</text>
              </view>
            </view>
          </view>
        </view>

        <!-- 结语 -->
        <view class="section">
          <text class="section-title">结语</text>
          <text class="paragraph">
            HOHO Park不仅仅是一个数字藏品平台，更是一个连接创作者和收藏者的桥梁。我们相信，通过技术创新和社区共建，能够创造一个更加开放、公平、有趣的数字资产世界。
          </text>
          <text class="paragraph highlight">
            让我们一起，见证数字藏品的无限可能！
          </text>
        </view>

        <!-- 底部信息 -->
        <view class="footer">
          <text class="footer-text">© 2024 HOHO Park</text>
          <text class="footer-text">版本：v1.0 | 更新时间：2024年11月</text>
        </view>
      </view>
    </scroll-view>
  </view>
</template>

<script setup>
import { ref } from 'vue'

const goBack = () => {
  uni.navigateBack()
}
</script>

<style scoped>
.whitepaper-page {
  width: 100%;
  height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

/* 自定义导航栏 */
.custom-navbar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  z-index: 1000;
  padding-top: env(safe-area-inset-top);
}

.navbar-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 88rpx;
  padding: 0 30rpx;
}

.back-btn {
  width: 60rpx;
  height: 60rpx;
  display: flex;
  align-items: center;
  justify-content: center;
}

.back-icon {
  font-size: 40rpx;
  color: #333;
}

.navbar-title {
  font-size: 32rpx;
  font-weight: bold;
  color: #333;
}

.placeholder {
  width: 60rpx;
}

/* 内容滚动区域 */
.content-scroll {
  height: 100vh;
  padding-top: calc(88rpx + env(safe-area-inset-top));
}

.content {
  padding: 40rpx 30rpx 60rpx;
}

/* 封面 */
.cover-section {
  background: white;
  border-radius: 20rpx;
  padding: 80rpx 40rpx;
  text-align: center;
  margin-bottom: 40rpx;
  box-shadow: 0 10rpx 40rpx rgba(0, 0, 0, 0.1);
}

.logo {
  font-size: 120rpx;
  margin-bottom: 30rpx;
}

.title {
  display: block;
  font-size: 56rpx;
  font-weight: bold;
  color: #333;
  margin-bottom: 20rpx;
}

.subtitle {
  display: block;
  font-size: 32rpx;
  color: #666;
  margin-bottom: 40rpx;
}

.version {
  display: inline-block;
  font-size: 24rpx;
  color: #999;
  padding: 10rpx 30rpx;
  background: #f5f5f5;
  border-radius: 30rpx;
}

/* 章节 */
.section {
  background: white;
  border-radius: 20rpx;
  padding: 40rpx;
  margin-bottom: 30rpx;
  box-shadow: 0 10rpx 40rpx rgba(0, 0, 0, 0.1);
}

.section-title {
  display: block;
  font-size: 40rpx;
  font-weight: bold;
  color: #333;
  margin-bottom: 30rpx;
  padding-bottom: 20rpx;
  border-bottom: 4rpx solid #667eea;
}

.paragraph {
  display: block;
  font-size: 28rpx;
  line-height: 48rpx;
  color: #666;
  margin-bottom: 20rpx;
  text-align: justify;
}

.highlight {
  color: #667eea;
  font-weight: bold;
}

/* 目录 */
.toc {
  display: flex;
  flex-direction: column;
  gap: 20rpx;
}

.toc-item {
  font-size: 28rpx;
  color: #666;
  padding: 20rpx;
  background: #f8f9fa;
  border-radius: 10rpx;
}

/* 愿景列表 */
.vision-list {
  display: flex;
  flex-direction: column;
  gap: 30rpx;
  margin-top: 30rpx;
}

.vision-item {
  display: flex;
  align-items: flex-start;
  gap: 20rpx;
}

.vision-icon {
  font-size: 48rpx;
  flex-shrink: 0;
}

.vision-content {
  flex: 1;
}

.vision-title {
  display: block;
  font-size: 32rpx;
  font-weight: bold;
  color: #333;
  margin-bottom: 10rpx;
}

.vision-desc {
  display: block;
  font-size: 26rpx;
  color: #999;
  line-height: 40rpx;
}

/* 功能卡片 */
.feature-card {
  margin-bottom: 30rpx;
  padding: 30rpx;
  background: #f8f9fa;
  border-radius: 15rpx;
}

.feature-title {
  display: block;
  font-size: 32rpx;
  font-weight: bold;
  color: #333;
  margin-bottom: 20rpx;
}

/* 技术列表 */
.tech-list {
  display: flex;
  flex-direction: column;
  gap: 20rpx;
  margin-top: 20rpx;
}

.tech-item {
  display: flex;
  align-items: center;
  padding: 20rpx;
  background: #f8f9fa;
  border-radius: 10rpx;
}

.tech-label {
  font-size: 28rpx;
  color: #999;
  width: 150rpx;
}

.tech-value {
  font-size: 28rpx;
  color: #333;
  font-weight: 500;
}

/* 积分列表 */
.points-list {
  display: flex;
  flex-direction: column;
  gap: 20rpx;
  margin-top: 20rpx;
}

.points-item {
  display: flex;
  align-items: center;
  gap: 15rpx;
}

.points-icon {
  font-size: 28rpx;
  color: #667eea;
  font-weight: bold;
}

.points-text {
  font-size: 28rpx;
  color: #666;
  line-height: 44rpx;
}

/* 路线图 */
.roadmap {
  display: flex;
  flex-direction: column;
  gap: 30rpx;
  margin-top: 30rpx;
  position: relative;
  padding-left: 40rpx;
}

.roadmap::before {
  content: '';
  position: absolute;
  left: 15rpx;
  top: 20rpx;
  bottom: 20rpx;
  width: 4rpx;
  background: linear-gradient(to bottom, #667eea, #764ba2);
}

.roadmap-item {
  position: relative;
  padding-left: 40rpx;
}

.roadmap-item::before {
  content: '';
  position: absolute;
  left: -29rpx;
  top: 10rpx;
  width: 20rpx;
  height: 20rpx;
  background: #667eea;
  border-radius: 50%;
  border: 4rpx solid white;
  box-shadow: 0 0 0 4rpx #667eea;
}

.roadmap-time {
  font-size: 24rpx;
  color: #999;
  margin-bottom: 10rpx;
}

.roadmap-content {
  background: #f8f9fa;
  padding: 20rpx;
  border-radius: 10rpx;
}

.roadmap-title {
  display: block;
  font-size: 30rpx;
  font-weight: bold;
  color: #333;
  margin-bottom: 10rpx;
}

.roadmap-desc {
  display: block;
  font-size: 26rpx;
  color: #666;
  line-height: 40rpx;
}

/* 底部 */
.footer {
  text-align: center;
  padding: 40rpx 0;
  margin-top: 40rpx;
}

.footer-text {
  display: block;
  font-size: 24rpx;
  color: rgba(255, 255, 255, 0.8);
  line-height: 40rpx;
}
</style>
