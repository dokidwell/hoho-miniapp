<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <text class="navbar-title">{{ title }}</text>
    </view>

    <!-- 协议内容 -->
    <scroll-view class="content" scroll-y>
      <view class="agreement-content">
        <!-- 用户服务协议 -->
        <view v-if="type === 'service'" class="section">
          <text class="section-title">HOHO Park 用户服务协议</text>
          
          <text class="section-subtitle">一、协议的范围</text>
          <text class="section-text">
            1.1 本协议是您与HOHO Park（以下简称"本平台"）之间关于您使用本平台服务所订立的协议。
          </text>
          <text class="section-text">
            1.2 本协议内容包括协议正文及所有本平台已经发布的或将来可能发布的各类规则。所有规则为本协议不可分割的组成部分，与协议正文具有同等法律效力。
          </text>

          <text class="section-subtitle">二、账号注册与使用</text>
          <text class="section-text">
            2.1 用户在使用本服务前需要注册一个账号。账号应当使用手机号码绑定注册，请用户使用尚未与本平台账号绑定的手机号码，以及未被本平台根据本协议封禁的手机号码注册账号。
          </text>
          <text class="section-text">
            2.2 用户注册成功后，本平台将给予每个用户一个用户账号及相应的密码，该用户账号和密码由用户负责保管；用户应当对以其用户账号进行的所有活动和事件负法律责任。
          </text>

          <text class="section-subtitle">三、数字作品服务</text>
          <text class="section-text">
            3.1 本平台提供的数字作品是基于区块链技术的数字凭证，具有唯一性、不可篡改性和可追溯性。
          </text>
          <text class="section-text">
            3.2 用户购买或兑换的数字作品仅供个人收藏、欣赏使用，不得用于任何商业目的。
          </text>
          <text class="section-text">
            3.3 数字作品的所有权归用户所有，但相关知识产权仍归原创作者所有。
          </text>

          <text class="section-subtitle">四、集换规则</text>
          <text class="section-text">
            4.1 用户可以在本平台内进行数字作品的集换（交易），集换使用平台积分进行结算。
          </text>
          <text class="section-text">
            4.2 用户挂售的作品必须是本人合法拥有的社区作品，第三方平台作品（如鲸探作品、Waveup作品）不可交易。
          </text>
          <text class="section-text">
            4.3 集换一旦完成，不可撤销，请用户谨慎操作。
          </text>

          <text class="section-subtitle">五、用户行为规范</text>
          <text class="section-text">
            5.1 用户在使用本服务时，必须遵守中华人民共和国相关法律法规。
          </text>
          <text class="section-text">
            5.2 用户不得利用本服务从事以下行为：
          </text>
          <text class="section-text">
            （1）发布、传送、传播、储存违反国家法律法规的内容；\n
            （2）发布、传送、传播、储存侵害他人知识产权、商业秘密等合法权利的内容；\n
            （3）恶意虚构事实、隐瞒真相以误导、欺骗他人；\n
            （4）发布、传送、传播广告信息及垃圾信息；\n
            （5）从事其他违法或违反本协议的行为。
          </text>

          <text class="section-subtitle">六、免责声明</text>
          <text class="section-text">
            6.1 本平台不对数字作品的市场价值做任何保证，用户应自行判断作品价值。
          </text>
          <text class="section-text">
            6.2 因不可抗力或者其他意外事件，使得本协议履行不可能、不必要或者无意义的，双方均不承担责任。
          </text>

          <text class="section-subtitle">七、协议的变更</text>
          <text class="section-text">
            7.1 本平台有权根据需要不时地制订、修改本协议或各类规则，并以网站公示的方式进行公告，不再单独通知用户。
          </text>
          <text class="section-text">
            7.2 变更后的协议和规则一经公布即生效，并成为本协议的一部分。用户如继续使用本服务，即视为用户接受本平台对本协议相关条款所做的修改。
          </text>
        </view>

        <!-- 隐私政策 -->
        <view v-else-if="type === 'privacy'" class="section">
          <text class="section-title">HOHO Park 隐私政策</text>
          
          <text class="section-subtitle">一、我们如何收集和使用您的个人信息</text>
          <text class="section-text">
            1.1 为了向您提供服务，我们会收集以下信息：\n
            - 账号信息：手机号码、昵称、头像\n
            - 交易信息：作品购买、兑换记录\n
            - 设备信息：设备型号、操作系统版本\n
            - 日志信息：IP地址、访问时间、访问页面
          </text>

          <text class="section-subtitle">二、我们如何使用Cookie和同类技术</text>
          <text class="section-text">
            2.1 为确保网站正常运转，我们会在您的设备上存储名为Cookie的小数据文件。Cookie通常包含标识符、站点名称以及一些号码和字符。
          </text>

          <text class="section-subtitle">三、我们如何共享、转让、公开披露您的个人信息</text>
          <text class="section-text">
            3.1 我们不会与任何公司、组织和个人分享您的个人信息，但以下情况除外：\n
            - 获得您的明确同意\n
            - 根据法律法规的要求\n
            - 与关联公司共享
          </text>

          <text class="section-subtitle">四、我们如何保护您的个人信息</text>
          <text class="section-text">
            4.1 我们已使用符合业界标准的安全防护措施保护您提供的个人信息，防止数据遭到未经授权访问、公开披露、使用、修改、损坏或丢失。
          </text>

          <text class="section-subtitle">五、您的权利</text>
          <text class="section-text">
            5.1 您对自己的个人信息享有以下权利：\n
            - 访问权：您有权访问您的个人信息\n
            - 更正权：您有权更正您的个人信息\n
            - 删除权：您有权要求我们删除您的个人信息\n
            - 撤回同意权：您有权撤回您的授权同意
          </text>

          <text class="section-subtitle">六、本政策如何更新</text>
          <text class="section-text">
            6.1 我们的隐私政策可能变更。未经您明确同意，我们不会削减您按照本隐私政策所应享有的权利。
          </text>
        </view>

        <!-- 数字作品服务协议 -->
        <view v-else-if="type === 'nft'" class="section">
          <text class="section-title">数字作品服务协议</text>
          
          <text class="section-subtitle">一、定义</text>
          <text class="section-text">
            1.1 数字作品：指基于区块链技术生成的、具有唯一性标识的数字凭证，用于证明特定数字内容的所有权。
          </text>
          <text class="section-text">
            1.2 社区作品：指在本平台内创作的数字作品，可以在平台内进行集换交易。
          </text>
          <text class="section-text">
            1.3 第三方作品：指从鲸探、Waveup等第三方平台同步的数字作品，仅供展示，可映射获得积分奖励，不可交易。
          </text>

          <text class="section-subtitle">二、作品创作与铸造</text>
          <text class="section-text">
            2.1 用户可以在本平台创作并铸造数字作品，铸造需要消耗一定数量的积分。
          </text>
          <text class="section-text">
            2.2 用户创作的作品必须是原创内容，不得侵犯他人知识产权。
          </text>
          <text class="section-text">
            2.3 作品一旦铸造成功，将永久记录在区块链上，不可删除或修改。
          </text>

          <text class="section-subtitle">三、作品权利</text>
          <text class="section-text">
            3.1 用户购买或兑换数字作品后，获得该作品的所有权，但不包括相关知识产权。
          </text>
          <text class="section-text">
            3.2 用户可以：\n
             - 持有、展示数字作品
            - 在平台内转赠或交易（社区作品）
            - 通过第三方作品映射获得积分（仅限第三方作品）
          </text>
          <text class="section-text">
            3.3 用户不得：\n
            - 用于商业目的\n
            - 侵犯原创作者的知识产权\n
            - 进行炒作、投机行为
          </text>

          <text class="section-subtitle">四、第三方作品说明</text>
          <text class="section-text">
            4.1 鲸探作品、Waveup作品等第三方平台作品，用户需要在相应平台拥有该作品，并通过账号绑定进行同步。
          </text>
          <text class="section-text">
            4.2 第三方作品仅用于展示，持有第三方作品可以映射获得积分奖励，但不可在本平台内进行交易。
          </text>
          <text class="section-text">
            4.3 第三方作品的权利和义务以原平台的协议为准。
          </text>

          <text class="section-subtitle">五、风险提示</text>
          <text class="section-text">
            5.1 数字作品市场存在不确定性，作品价值可能波动，用户应理性参与。
          </text>
          <text class="section-text">
            5.2 区块链技术仍在发展中，可能存在技术风险，用户应充分了解并承担相关风险。
          </text>
          <text class="section-text">
            5.3 本平台不对数字作品的投资价值做任何承诺或保证。
          </text>
        </view>
      </view>
    </scroll-view>
  </view>
</template>

<script setup>
import { ref, onLoad } from '@dcloudio/uni-app'

const type = ref('service') // service, privacy, nft
const title = ref('用户服务协议')

onLoad((options) => {
  if (options.type) {
    type.value = options.type
    
    const titleMap = {
      'service': '用户服务协议',
      'privacy': '隐私政策',
      'nft': '数字作品服务协议'
    }
    title.value = titleMap[options.type] || '用户服务协议'
  }
})
</script>

<style lang="scss" scoped>
.page {
  min-height: 100vh;
  background-color: #FFFFFF;
  display: flex;
  flex-direction: column;
}

.navbar {
  height: 88rpx;
  background-color: #FFFFFF;
  display: flex;
  align-items: center;
  justify-content: center;
  border-bottom: 1rpx solid #F0F0F0;
  
  .navbar-title {
    font-size: 32rpx;
    font-weight: 600;
    color: #000000;
  }
}

.content {
  flex: 1;
  padding: 32rpx;
}

.agreement-content {
  .section {
    .section-title {
      font-size: 36rpx;
      font-weight: 600;
      color: #000000;
      display: block;
      text-align: center;
      margin-bottom: 48rpx;
    }
    
    .section-subtitle {
      font-size: 30rpx;
      font-weight: 600;
      color: #000000;
      display: block;
      margin-top: 32rpx;
      margin-bottom: 16rpx;
    }
    
    .section-text {
      font-size: 26rpx;
      color: #333333;
      line-height: 1.8;
      display: block;
      margin-bottom: 16rpx;
      text-align: justify;
    }
  }
}
</style>
