<template>
	<view v-if="isShow" style="position: relative;">

		<view style="width: 100%;margin-bottom: 30rpx" v-if="((arr.length == 0)&&(!noMore))" v-for="item in 4" :key="item">
			<tui-card>
				<view slot="body" class="tui-default">
					<view class="card_pic" style="background-color: #f8f8f8;border-radius: 15rpx;margin-bottom: 22rpx;height: 250rpx;"></view>
					<view style="width: 100%;background-color: #f8f8f8;border-radius: 4rpx;height: 28rpx;margin-bottom: 10rpx;" v-for="i in 4"></view>
				</view>
				<view slot="footer" class="tui-default" style="padding-top: 1rpx;display: flex;width: 100%;justify-content: space-between;box-sizing: border-box;">
					<view style="width: 180rpx;background-color: #f8f8f8;border-radius: 4rpx;height: 48rpx" v-for="i in 3"></view>
				</view>
			</tui-card>
		</view>

		<view class="cu-flex-center loadingio-box" v-if="((arr.length == 0)&&(!noMore))">
			<view class="loadingio-spinner-eclipse-ill2naumkh">
				<view class="ldio-zeyn0hj1rl">
					<view></view>
				</view>
			</view>
		</view>
		
		<view v-if="((arr.length == 0)&&(noMore))" style="margin-top: 40rpx;width: 100%;text-align: center;font-size: 28rpx;color: #aaaaaa;">
			暂无更多
		</view>

		<view style="margin-bottom: 30rpx" v-for="(item,index) in arr" :key="index" v-if="arr.length != 0">
			<tui-card>
				<view slot="body" class="tui-default">
					<view class="tui-tag-small tui-warning tui-tag-fillet-right cus-tag">{{item.state}}</view>
					<image :src="item.img" mode="widthFix" class="card_pic"></image>
					<view style="margin-top: 18rpx;">姓名：{{item.Student.name}}</view>
					<view>学号：{{item.Student.id}} </view>
					<view>班级：{{item.Student.initTreeName}} </view>
					<view>时间：{{item.time}}</view>
					<view>类型：{{item.type}}/{{item.Item.name}}</view>
					<view>备注：{{item.remark ? item.remark : "无"}}</view>
					<view>进度：<span style="color: #5C8DFF;text-decoration:underline" @click="">点此查看</span></view>
				</view>
				<view slot="footer" class="tui-default">
					<view class="tui-flex">
						<view class="tui-center tui-flex-1">
							<tui-icon name="star" size="10" color="#ed3f14"></tui-icon>e
							申请加分{{item.applyScore}}
						</view>
						<view class="tui-center tui-flex-1">
							<tui-icon name="notice" size="10" color="#19be6b"></tui-icon>
							参考加分{{item.Item.shouldScore}}
						</view>
						<view class="tui-center tui-flex-1">
							<tui-icon name="eye" size="10" color="#5677fc"></tui-icon>
							实际加分{{item.realScore}}
						</view>
					</view>
					<view v-if="showOp" style="height: 120rpx;width: 100%;box-sizing: border-box;padding: 30rpx;background-color: #ffffff;">
						<view @click="operate(item,true,item.realScore,item.remark)" style="background-color: #5677fc;float: right;width: 80rpx;height: 80rpx;border-radius: 40rpx;box-shadow: 1px 1px 10px 0px rgba(179,179,179,0.15);display: flex;justify-content: center;align-items: center;">
							<tui-icon name="check" size="20" color="#ffffff"></tui-icon>
						</view>
						<view @click="operate(item,false,item.realScore,item.remark)" style="background-color: #ffffff;margin-right: 30rpx;float: right;width: 80rpx;height: 80rpx;border-radius: 40rpx;box-shadow: 1px 1px 10px 0px rgba(179,179,179,0.15);display: flex;justify-content: center;align-items: center;">
							<tui-icon name="shut" size="20" color="#ed3f14"></tui-icon>
						</view>
					</view>
				</view>
			</tui-card>
		</view>

	</view>
</template>

<script>
	import tuiModal from '@/components/thorui/modal/modal'
	import tuiButton from '@/components/thorui/button/button'
	import tuiIcon from '@/components/thorui/icon/icon'
	import tuiCard from '@/components/thorui/card/card'
	import tuiRate from '@/components/thorui/rate/rate'
	export default {
		data() {
			return {
				arr: [],
				noMore:false,
			}
		},
		components: {
			tuiIcon,
			tuiRate,
			tuiCard,
			tuiModal,
			tuiButton
		},
		props: {
			isShow: {
				type: Boolean
			},
			showOp:{
				type:Boolean
			}
		},
		methods: {
			operate(item,isPass) {
				this.$emit('operate',item,isPass)
			},
			tsToStr(ts) {
				return new Date(parseInt(ts) * 1000).toLocaleString().replace(/:\d{1,2}$/, ' ')
			},
			setNoMore() {
				this.arr = []
				this.noMore = true
			},
			setLoading() {
				this.arr = []
				this.noMore = false
			},
			addData(data) {
				const that = this
				return new Promise(function(resolve, reject) {
					//加入列表
					for (let i = 0; i < data.length; i++) {
						let obj = {}
						obj = data[i]
						obj.time = that.tsToStr(obj.time)
						that.arr.push(obj)
					}
					resolve(true)
				})
			},
			clearData() {
				const that = this
				return new Promise(function(resolve, reject) {
					that.arr = []
					resolve(true)
				})
			},
		}
	}
</script>

<style>
	.tui-default {
		padding: 20rpx 30rpx;
		font-size: 30rpx;
		position: relative;
	}

	.card_pic {
		width: 100%;
		height: 80rpx;
	}

	.cus-tag {
		position: absolute;
		top: 20rpx;
		left: 30rpx;
		z-index: 99;
	}

	.card_end {
		width: 25%;
		position: absolute;
		top: 30rpx;
		right: 30rpx;
		z-index: 9999;
	}

	@keyframes ldio-zeyn0hj1rl {
		0% {
			transform: rotate(0deg)
		}

		50% {
			transform: rotate(180deg)
		}

		100% {
			transform: rotate(360deg)
		}
	}

	.ldio-zeyn0hj1rl view {
		position: absolute;
		animation: ldio-zeyn0hj1rl 1s linear infinite;
		width: 160rpx;
		height: 160rpx;
		top: 20rpx;
		left: 20rpx;
		border-radius: 50%;
		box-shadow: 0 4rpx 0 0 #5677fc;
		transform-origin: 80rpx 82rpx;
	}
	
	.loadingio-box{
		width: 100%;
		height: 400rpx;
		z-index: 999;
		position: absolute;
		top:0rpx;
		left:0rpx;
	}

	.loadingio-spinner-eclipse-ill2naumkh {
		width: 200rpx;
		height: 200rpx;
		overflow: hidden;
	}

	.ldio-zeyn0hj1rl {
		width: 100%;
		height: 100%;
		position: relative;
		transform: translateZ(0) scale(1);
		backface-visibility: hidden;
		transform-origin: 0 0;
	}

	.ldio-zeyn0hj1rl view {
		box-sizing: content-box;
	}

	image {
		will-change: transform
	}
</style>
