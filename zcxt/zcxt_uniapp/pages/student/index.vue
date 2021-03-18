<template>
	<scroll-view :scroll-y="true" :scroll-x="false" style="width: 100%;height: 100%;">
		<top-box title="首页" :isShowShare="true"></top-box>
		<u-notice-bar mode="horizontal" :list="['欢迎使用广东技术师范大学综测系统小程序！']"></u-notice-bar>
		<swiper :indicator-dots="false" :autoplay="true" :interval="2300" :duration="600" :circular="true" style="height: 170rpx;margin-top: 20rpx;">
			<swiper-item v-for="(item,index) in banner" :key="index" class="swiper-item-box">
				<image :src="item" class="swiper-item-img" @click="navToGuide(item)" />
			</swiper-item>
		</swiper>

		<view style="width: 400rpx;margin:-20rpx 175rpx 20rpx 175rpx;">
			<u-tabs :list="topTab" :is-scroll="false" :current="currentTopTab" :bold="true" @change="changeTopBar"></u-tabs>
		</view>

		<verify-list :showOp="false" ref="underway" :isShow="currentTopTab == 0"></verify-list>
		<verify-list :showOp="false" ref="finish" :isShow="currentTopTab == 1"></verify-list>

		<view style="text-align: center;margin:70rpx 0 200rpx;width:100%;line-height: 40rpx;font-size:30rpx">
			<image src="http://www.hnyouth.org.cn/public/site/gqt20171105/images/slide2_4.png" style="height:30rpx;width:30rpx"
			 mode="widthFix"></image>
			广东技术师范大学综测工作
		</view>

		<index-bar :select="0"></index-bar>
	</scroll-view>
</template>

<script>
	var that
	import topBox from '@/components/custom/top-box'
	import verifyList from '@/components/custom/verify-list'
	import indexBar from '@/components/custom/index-bar'
	export default {
		data() {
			return {
				banner: [
					"/static/custom/baike.png",
					"/static/custom/shangshou.png"
				],
				underwayverifyListData: [],
				suspendverifyListData: [],
				currentTopTab: 0,
				topTab: [{
					name: "审核中"
				}, {
					name: "已审核"
				}],
			}
		},
		components: {
			indexBar,
			verifyList,
			topBox
		},
		onShow() {
			uni.hideTabBar()
			that = this
			this.getList('underway')
		},
		methods: {
			navToGuide(src) {
				let type
				if (src === "/static/custom/baike.png") {
					type = 'intro'
				}
				if (src === "/static/custom/shangshou.png") {
					type = 'use'
				}
				uni.navigateTo({
					url: '/pages/student/guide?type=' + type
				})
			},
			async getList(ref) {
				const component = this.$refs[ref]
				component.setLoading()
				const res = await this.$req.VerifyRead({
					state:ref
				})
				if((res === false) || (res.length === 0)) {
					component.setNoMore()
				} else {
					await component.clearData()
					await component.addData(res)
				}
			},
			changeTopBar(index) {
				if (index == 0) {
					this.getList('underway')
				} else if (index == 1) {
					this.getList('finish')
				}
				this.currentTopTab = index
			},
		}
	}
</script>

<style>
	.swiper-item-box{
		width: 100%;padding: 15rpx 80rpx;box-sizing: border-box;
	}
	
	.swiper-item-img{
		width: 100%;height: 120rpx;border-radius: 60rpx;
	}
</style>
