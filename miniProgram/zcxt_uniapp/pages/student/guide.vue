<template>
	<view>
		<view v-if="!type" class="banner-img-box">
			<image src="../../static/custom/baike.png" class="banner-img" @click="()=>{type = 'intro'}"></image>
			<image src="../../static/custom/shangshou.png" class="banner-img" @click="()=>{type = 'use'}"></image>
		</view>
		<view v-if="type" class="ctx-box">
			<text v-if="type === 'intro'">
				小程序简介:
				广东技术师范大学综测系统小程序是一个用于学生评优的加分证明审核系统，上传加分证明图片（奖状、奖杯图片）后，经过审核就可以获得相应的学分。
				传统综测工作使用纸质进行综测统计，存在统计效率不高、灵活度不够、透明度不高等缺点。
				GPNU综测系统小程序是一个将综测审核工作电子化的Web应用，在综测工作老师与相关工作人员指导与配合下完成，能够很好地加大综测工作的效率。</text>
			<text v-if="type === 'use'">
				注册：用户进入小程序后绑定自己的个人学号与姓名即可。

				审核：在首页点击“评优项审核”按钮，接着点击立即上传，选择一张图片，接着根据建议加分输入申请加分，提交审核即可。

				流程：审核信息会依次通过多级审核人员进行审核，审核成功后即可获得加分，审核失败后会获得相应的反馈。
			</text>
			<u-button type="primary" @click="()=>{type=false}" style="margin-top: 20rpx;">查看更多使用说明</u-button>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				type: false
			}
		},
		onLoad(req) {
			if (req.type) {
				this.type = req.type
			} else {
				this.type = false
			}
		}
	}
</script>

<style>
	.ctx-box {
		width: 100%;
		padding: 45rpx;
		box-sizing: border-box;
		font-size: 32rpx;
		line-height: 49rpx;
	}

	.banner-img {
		width: 100%;
		margin: 20rpx 0rpx;
		height: 150rpx;
	}

	.banner-img-box {
		width: 100%;
		padding: 25rpx;
		box-sizing: border-box;
	}
</style>
