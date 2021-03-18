<template>
	<view>
		<u-navbar :is-back="false" title="　" :border-bottom="false">
		</u-navbar>
		<view class="u-flex u-p-l-30 u-p-r-20 u-p-b-30" style="background-color: #fff;">
			<view class="u-m-r-10">
				<u-avatar size="140"></u-avatar>
			</view>
			<view class="u-flex-1">
				<view class="u-font-18 u-p-b-20">{{name}}</view>
				<view class="u-font-14 u-tips-color">{{initTreeName}}</view>
			</view>
		</view>
		<view class="u-m-t-20">
			<u-cell-group>
				<u-cell-item icon="star" title="总分榜" @click="()=>{isShowRankPopup=true}"></u-cell-item>
				<u-cell-item icon="play-right" title="品德榜"></u-cell-item>
				<u-cell-item icon="play-right" title="学业榜"></u-cell-item>
				<u-cell-item icon="play-right" title="文体榜"></u-cell-item>
			</u-cell-group>
		</view>
		<view class="u-m-t-20">
			<u-cell-group>
				<u-cell-item icon="email" title="使用说明" @click="navTo('/pages/student/guide')"></u-cell-item>
			</u-cell-group>
		</view>
		<view class="u-m-t-20">
			<u-cell-group>
				<u-cell-item icon="close" title="退出登录" @click="backToLogin"></u-cell-item>
			</u-cell-group>
		</view>
		<index-bar :select="2"></index-bar>
		
		<u-popup mode="bottom" v-model="isShowRankPopup" :closeable="true" height="70%">
			<view style="width: 100%;text-align: center;font-size: 36rpx;margin-top: 20rpx;">班级综测总分公示榜</view>
			<view style="width: 100%;text-align: center;font-size: 28rpx;color: #999999;margin-bottom: 20rpx;">每周一、周四24：00自动更新</view>
			<u-table>
					<u-tr>
						<u-th>名次</u-th>
						<u-th>姓名</u-th>
						<u-th>分值</u-th>
					</u-tr>
					<u-tr>
						<u-td>1</u-td>
						<u-td>小明</u-td>
						<u-td>22</u-td>
					</u-tr>
					<u-tr>
						<u-td>2</u-td>
						<u-td>大明</u-td>
						<u-td>20</u-td>
					</u-tr>
				</u-table>
		</u-popup>
	</view>
</template>

<script>
	import indexBar from '@/components/custom/index-bar'
	export default {
		data() {
			return {
				name: "",
				initTreeName: "",
				isShowRankPopup:false,
			}
		},
		onLoad() {
			this.flashCache()
		},
		methods: {
			navTo(url) {
				uni.navigateTo({
					url: url
				})
			},
			backToLogin() {
				uni.clearStorageSync()
				uni.reLaunch({
					url:"/pages/other/login"
				})
			},
			flashCache() {
				const isWant = uni.getStorageSync('isWant')
				this.isWant = isWant ? isWant : false
				const name = uni.getStorageSync('name')
				this.name = name ? name : ""
				const treeName = uni.getStorageSync('treeName')
				this.initTreeName = treeName ? treeName : ""
			}
		},
		components: {
			indexBar
		}
	}
</script>

<style lang="scss">
	page{
		background-color: #f7f7f7;
	}
</style>
