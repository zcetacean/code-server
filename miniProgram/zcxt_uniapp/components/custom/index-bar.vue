<template>
	<view>
		<u-tabbar v-model="select" :list="barConf" active-color="#2979ff" height="60px" icon-size="50" @change="tabbarSwitch"></u-tabbar>
	</view>
</template>

<script>
	export default {
		name: "indexBar",
		data() {
			return {
				barConf: [{
						iconPath: "home",
						selectedIconPath: "home-fill",
						text: '首页',
						customIcon: false,
						page:"/pages/student/index"
					},
					{
						iconPath: "plus",
						selectedIconPath: "plus",
						text: '新增',
						customIcon: false,
						page:"/pages/student/add"
					},
					{
						iconPath: "account",
						selectedIconPath: "account-fill",
						text: '我的',
						customIcon: false,
						page:"/pages/student/me"
					}
				],
			}
		},
		props: {
			select: {
				type: Number
			}
		},
		methods: {
			tabbarSwitch(index) {
				const that = this
				uni.switchTab({
					url: that.barConf[index].page
				})
			}
		}
	}
</script>

<style>
</style>
