<template>
	<view style="padding: 40rpx;">
		<view v-if="step === 0">
			<u-steps :list="stepList" :current="stepListCurrent"></u-steps>
			<u-card>
				<view slot="body">
				<u-search placeholder="综测条目关键字" v-model="keyword"></u-search>
				<u-cell-group>
					<u-cell-item icon="checkmark-circle-fill" title="品德/三好学生/3分" :arrow="false"></u-cell-item>
					<u-cell-item icon="checkmark-circle" title="学业/省级互联网+金奖" :arrow="false"></u-cell-item>
				</u-cell-group>
				</view>
			</u-card>
			<u-field :disabled="true" v-model="verifyItem.display" label="当前选中" placeholder="请选择综测条目" type="textarea"></u-field>
			<u-button type="primary" @click="()=>{step = 1}">确认选择</u-button>
		</view>
		
		<view v-if="step === 1">
			参考加分：品德/三号学生/3分
			期望加分：
			<u-number-box v-model="value"></u-number-box>
			<u-upload max-count="3"></u-upload>
			<u-button type="primary" @click="()=>{step === 1}">确认申请</u-button>
			<u-button @click="()=>{step === 0}">返回上一步</u-button>
		</view>

		<!-- <view class="parent-box">
			<view class="children-box">
				<view style="box-shadow:1rpx 1rpx 8rpx 1rpx rgba(0, 0, 0, 0.10);overflow: hidden;border-radius: 30rpx;margin-left: 10%;width: 80%;margin-bottom: 40rpx;display: flex;justify-content: center;margin-top: 40rpx;">
					<view v-if="img" style="width: 530rpx;height: 380rpx;box-sizing: border-box;position: relative;">
						<image @click="preImg" :src="img" style="width: 100%;height: 100%;" mode="aspectFill"></image>
						<view @click="clearImg" style="opacity: 0.8;box-shadow:1rpx 1rpx 8rpx 1rpx rgba(0, 0, 0, 0.10);position: absolute;top:10rpx;right: 10rpx;height: 50rpx;width: 50rpx;border-radius: 25rpx;display: flex;justify-content: center;align-items: center;background-color: #ffffff;">
							<tui-icon style="margin-top: -2rpx;" name="shut" size="15" color="#bbbbbb"></tui-icon>
						</view>
					</view>
					<tui-icon @click="chooseImg" v-if="!img" style="margin: 50rpx 0rpx;" name="plus" size="35" color="#bbbbbb"></tui-icon>
				</view>
				<picker style="margin: 35rpx 0rpx 40rpx 10%;" class="shadowBox" @change="selectItem" :range="verifyItemArr.display">
					<input v-model="item.display" placeholder="请选择综测加分项" style="width: 100%;height: 60rpx;font-size: 30rpx;line-height: 60rpx;" />
				</picker>
				<view style="padding: 30rpx;width: 80%;border-radius: 30rpx;box-shadow:1rpx 1rpx 8rpx 1rpx rgba(0, 0, 0, 0.10);margin-left: 10%;box-sizing: border-box;">
					<view style="width: 100%;text-align: left;color: #aaaaaa;font-size:28rpx;margin-bottom: 30rpx;">申请加分：</view>
					<slider activeColor="#108ee9" @change="sliderChange" show-value />
				</view>
				<view style="width: 100%;display: flex;justify-content: center;">
					<view class="tui-mbtm" style="margin-top: 40rpx;" @click="confirmUpload">
						<tui-button type="primary">确认提交</tui-button>
					</view>
				</view>
			</view>
		</view> -->

		<index-bar :select="1"></index-bar>
	</view>
</template>

<script>
	import tuiButton from '@/components/thorui/button/button'
	import tuiIcon from '@/components/thorui/icon/icon'
	import indexBar from '@/components/custom/index-bar'
	export default {
		data() {
			return {
				step:0,
				stepList: [{
					name: '选择类型'
				}, {
					name: '填写表单'
				}],
				stepListCurrent: 0,
				verifyItemArr: {
					display: [],
					source: []
				},
				verifyItem: {
					display: "品德/三好学生",
					source: null
				},
				img: "",
				remark: "",
				wantAddScore: 0,
				hackReset: true,
				tips_list: [{
					"icon": "shield",
					"text": "更快速的审核"
				}, {
					"icon": "signin",
					"text": "更透明的管理"
				}, {
					"icon": "refresh",
					"text": "随时获取结果"
				}],
			}
		},
		onLoad() {
			this.getItem()
		},
		components: {
			indexBar,
			tuiIcon,
			tuiButton
		},
		methods: {
			async confirmUpload() {
				uni.showLoading()
				const res = await this.$req.VerifyCreate({
					initTreeId: this.verifyItem.source.init_tree_id,
					img: this.img,
					id: this.verifyItem.source.id,
					type: this.verifyItem.source.type,
					applyScore: this.wantAddScore
				})
				uni.hideLoading()
				if (res) {
					uni.showToast({
						title: "申请成功"
					})
				} else {
					uni.showToast({
						icon: "none",
						title: "申请失败"
					})
				}
			},
			selectItem(e) {
				this.verifyItem.display = this.verifyItemArr.display[e.target.value]
				this.verifyItem.source = this.verifyItemArr.source[e.target.value]
			},
			async getItem() {
				const res = await this.$req.UserReadItem()
				this.verifyItemArr.source = res ? res : []
				if (res) {
					this.verifyItemArr.display = res.map(item => {
						return item.type + "/" + item.name + "/参考加分：" + item.should_score + "分"
					})
				} else {
					this.verifyItemArr.display = []
				}
			},
			preImg() {
				const that = this
				uni.previewImage({
					urls: [that.img]
				})
			},
			clearImg() {
				this.img = ""
			},
			chooseImg() {
				const that = this
				uni.chooseImage({
					success: (chooseImageRes) => {
						const tempFilePath = chooseImageRes.tempFilePaths[0]
						that.img = tempFilePath
					}
				});
			},
			sliderChange(e) {
				this.wantAddScore = e.detail.value
			},
			onShareAppMessage() {
				const that = this
				return {
					title: "广东技术师范大学综测系统小程序",
					imageUrl: '/static/custom/logo.png',
					path: '/pages/student/add'
				}
			}
		}
	}
</script>

<style lang="scss">
	.parent-box {
		width: 100%;
		height: 78%;
		display: flex;
		align-items: center;
		justify-content: center;
		position: absolute;
		top: 2%;
		left: 0rpx
	}

	.shadowBox {
		box-shadow: 1rpx 1rpx 8rpx 1rpx rgba(0, 0, 0, 0.10);
		width: 80%;
		height: 90rpx;
		border-radius: 45rpx;
		background-color: #ffffff;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.children-box {
		box-shadow: 0 0px 12px 0 rgba(0, 0, 0, .1);
		background-color: #fff;
		width: 90%;
		height: 100%;
		padding: 0rpx 0rpx 100rpx 0rpx;
		border-radius: 20rpx;
		text-align: center;
		font-size: 33rpx;
		color: #333;
		box-sizing: border-box;
		overflow: auto;
	}

	.logo {
		width: 24%;
		margin: 60rpx 38%;
	}

	.download-label {
		width: 100%;
		margin-top: 15rpx;
		font-size: 24rpx;
		color: #cccccc;
		text-align: center;
	}

	.btn-box {
		width: 100%;
		display: flex;
		justify-content: center;
	}

	.intro-list {
		color: #80848f;
		margin-top: 30rpx;
	}

	.intro-list text {
		margin-left: 18rpx;
	}

	.tui-mbtm {
		width: 50% !important;
		margin-top: 80rpx;
	}

	.tui-btn-verify {
		width: 75% !important;
		margin-top: 80rpx;
	}
</style>
