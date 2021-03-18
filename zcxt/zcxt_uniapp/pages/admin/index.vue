<template>
	<scroll-view :scroll-y="true" :scroll-x="false" style="width: 100%;height: 100%;">
		<top-box title="首页" :showShare="false"></top-box>

		<view style="margin-top: 50rpx;">
			<u-tabs :list="topTab" :is-scroll="false" :current="currentTopTab" :bold="true" @change="changeTopBar"></u-tabs>
		</view>

		<verify-list @operate="operate" :showOp="true" ref="waitVerify" :isShow="currentTopTab == 0"></verify-list>
		<tui-list-view class="tui-list-view" v-if="currentTopTab == 1">
			<tui-list-cell @click="toShowDetail(item)" :arrow="true" v-for="(item,index) in recordArr" :key="index">
				{{item.time}}
			</tui-list-cell>
		</tui-list-view>

		<tui-modal :show="showDetail" :custom="true">
			<view style="width: 100%;text-align: left;font-size: 34rpx;color: #444444;margin-bottom: 25rpx;">
				<text>{{detailCtx}}</text>
			</view>
			<tui-button type="primary" shape="circle" @click="()=>{showDetail = false}">确认</tui-button>
		</tui-modal>

		<tui-modal :show="showOperate" :custom="true">
			<view style="width: 100%;padding: 0rpx;box-sizing: border-box;margin-top: 35rpx;">
				{{isPass ? "确定通过本次审核？" : "确定不通过本次审核？"}}
			</view>
			<view style="padding: 0rpx;width: 100%;box-sizing: border-box;margin-top: 35rpx;">
				<view style="width: 100%;text-align: left;color: #aaaaaa;font-size:28rpx;margin-bottom: 30rpx;">批准加分：</view>
				<slider activeColor="#108ee9" @change="sliderChange" :value="realScore" show-value />
			</view>
			<view style="width: 100%;padding: 0rpx;box-sizing: border-box;margin-top: 35rpx;">
				<input v-model="remark" placeholder="点此输入审核备注信息(可选)" style="width: 80%;height: 60rpx;font-size: 30rpx;line-height: 60rpx;" />
			</view>
			<view style="width: 100%;margin-top: 45rpx;margin-bottom: 10rpx;">
				<view style="width: 45%;float: left;">
					<tui-button size="mini" type="gray" shape="square" @click="()=>{showOperate = false}">取消</tui-button>
				</view>
				<view style="width: 45%;margin-left: 10%;float: left;">
					<tui-button type="primary" size="mini" shape="square" @click="confirmOperate">确认</tui-button>
				</view>
			</view>
		</tui-modal>

		<view class="round" @click="openPersonalModal" style="bottom: 30rpx;right: 30rpx;background-color: #5677fc;">
			<tui-icon name="ellipsis" size="25" color="#ffffff"></tui-icon>
		</view>

		<tui-modal :show="isShowPersonal" :custom="true">
			<text class="welcome-text">
				帐号:{{adminName ? adminName : ""}} \n权限:{{adminTreeName ? adminTreeName : ""}}
			</text>
			<view style="width: 100%;margin-top: 45rpx;margin-bottom: 10rpx;">
				<view style="width: 45%;float: left;">
					<tui-button size="mini" type="primary" shape="square" @click="()=>{isShowPersonal = false}">返回</tui-button>
				</view>
				<view style="width: 45%;margin-left: 10%;float: left;">
					<tui-button size="mini" type="gray" shape="square" @click="logout">退出登录</tui-button>
				</view>
			</view>
		</tui-modal>

	</scroll-view>
</template>

<script>
	var that
	import tuiIcon from '@/components/thorui/icon/icon'
	import tuiButton from '@/components/thorui/button/button'
	import tuiModal from '@/components/thorui/modal/modal'
	import tuiListView from "@/components/thorui/list-view/list-view"
	import tuiListCell from "@/components/thorui/list-cell/list-cell"
	import topBox from '@/components/custom/top-box'
	import verifyList from '@/components/custom/verify-list'
	export default {
		data() {
			return {
				isShowPersonal:false,
				adminName:"",
				adminTreeName:"",
				curItem: null,
				remark: "",
				realScore:0,
				isPass: false,
				showOperate: false,
				showDetail: false,
				detailCtx: "",
				recordArr: [],
				underwayverifyListData: [],
				suspendverifyListData: [],
				currentTopTab: 0,
				topTab: [{
					name: "待审核"
				}, {
					name: "操作记录"
				}],
			}
		},
		components: {
			tuiIcon,
			verifyList,
			topBox,
			tuiListView,
			tuiListCell,
			tuiModal,
			tuiButton,
		},
		onShow() {
			uni.hideTabBar()
			that = this
			this.getList()
			this.adminName = uni.getStorageSync('name')
			this.adminTreeName = uni.getStorageSync('treeName')
		},
		methods: {
			openPersonalModal() {
				this.isShowPersonal = true
			},
			logout() {
				uni.clearStorageSync()
				uni.reLaunch({
					url: "/pages/other/login"
				})
			},
			sliderChange(e) {
				this.realScore = e.detail.value
			},
			operate(item, isPass, realScore,remark) {
				console.log("demodsv")
				this.remark = remark
				this.realScore = realScore
				this.curItem = item
				this.isPass = isPass
				this.showOperate = true
			},
			async confirmOperate() {
				const res = await this.$req.VerifyUpdate({
					type: this.isPass ? "yes" : "no",
					id: this.curItem.id,
					remark: this.remark,
					realScore: this.realScore
				})
				if (res) {
					this.showOperate = false
					uni.showToast({
						title: "操作成功"
					})
					this.getList()
				} else {
					uni.showToast({
						title: "操作失败",
						icon: "none"
					})
				}
			},
			toShowDetail(item) {
				this.detailCtx = item.time + "\n" + item.remark
				this.showDetail = true
			},
			async getRecordArr() {
				const res = await this.$req.RecordRead()
				this.recordArr = res ? res : []
			},
			async getList() {
				const component = this.$refs.waitVerify
				component.setLoading()
				let res = await this.$req.VerifyReadWait({
					name: uni.getStorageSync("treeName")
				})
				if ((res === false) || (res === null)) {
					component.setNoMore()
				} else {
					await component.clearData()
					await component.addData(res)
				}
			},
			onShareAppMessage() {
				const that = this
				return {
					title: "广东技术师范大学综测小程序",
					imageUrl: '/static/custom/logo.png',
					path: '/pages/student/index'
				}
			},
			changeTopBar(e) {
				const index = e.index
				if (index == 0) {
					this.getList('underway', 'underway')
				} else if (index == 1) {
					this.getRecordArr()
				}
				this.currentTopTab = index
			},
		}
	}
</script>

<style>
	.round {
		width: 100rpx;
		height: 100rpx;
		position: fixed;
		background-color: #ffffff;
		box-shadow: 1px 1px 10px 0px rgba(179, 179, 179, 0.25);
		border-radius: 50rpx;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	page {
		background-color: #fcfcfc;
		overflow-x: hidden;
		width: 100%;
	}

	.welcome-text {
		font-size: 28rpx;
		width: 100%;
		word-break: break-all;
		margin-top: 35rpx;
	}
</style>
