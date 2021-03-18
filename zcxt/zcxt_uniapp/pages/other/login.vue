<template>
	<view>
		<scroll-view :scroll-y="true" :scroll-x="false">
			<view class="cu-flex-center logo-box">
				<image src="/static/custom/logo.png" style="width: 160rpx;height: 160rpx;"></image>
			</view>
			<view class="form-box">
				<view class="cu-flex-center-y cu-shadow-m form-id">
					<u-field v-model="id" label="帐号:" label-align="center" placeholder="请输入学号/工号/ID"></u-field>
				</view>
				<view class="cu-flex-center-y cu-shadow-m form-password">
					<u-field @right-icon-click="changePasswordShowStatus" :password="isShowPassword" right-icon="eye" :clearable="false"
					 v-model="password" label="密码:" label-align="center" placeholder="请输入密码"></u-field>
				</view>
			</view>
			<view @click="()=>{isShowTypeSelector = true}" class="cu-flex-center-x confirm-btn">
				<view class="cu-flex-center cu-shadow-m arrow-box">
					<image src="/static/custom/arrow.png" style="width: 60rpx;height: 50rpx;"></image>
				</view>
			</view>
			<view class="visitor-login-ctx" @click="visitorLogin">暂无帐号？点此使用游客帐号体验小程序</view>
		</scroll-view>
		<u-action-sheet :list="loginType" v-model="isShowTypeSelector" @click="login"></u-action-sheet>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				loginType: [{
					text: "学生登录"
				}, {
					text: "管理登录"
				}],
				isShowPassword: true,
				id: "",
				password: "",
				isShowTypeSelector: false,
			}
		},
		onShow() {
			try {
				const cookieKey = uni.getStorageSync("cookieKey")
				if (cookieKey !== "") {
					const type = uni.getStorageSync("type")
					if (type === 'admin') {
						this.toAdminPage()
					} else if (type === 'student') {
						this.toStudentPage()
					}
				}
			} catch (e) {
				uni.showToast({
					icon: "none",
					title: "登录异常"
				})
			}
		},
		methods: {
			changePasswordShowStatus() {
				this.isShowPassword = !this.isShowPassword
			},
			visitorLogin() {
				this.id = "TestId"
				this.password = "Password"
			},
			async login(index) {
				if ((this.id === "") || (this.password === "")) {
					uni.showToast({
						icon: "none",
						title: "帐号或密码不能为空"
					})
					return null
				}
				let type = index === 0 ? 'student' : 'teacher'
				const res = await this.$req.UserLogin({
					id: this.id,
					password: this.password,
					state: type
				})
				if (res) {
					if (res.hasOwnProperty('initTreeName')) {
						this.$req.util.saveUserInfo(res.initTreeName, res.name, "student")
						this.toStudentPage()
					} else {
						this.$req.util.saveUserInfo(res.permissionTreeName, res.name, "admin")
						this.toAdminPage()
					}
				} else {
					uni.showToast({
						icon: "none",
						title: "登录失败"
					})
				}
			},
			toStudentPage() {
				wx.switchTab({
					url: "/pages/student/index"
				})
			},
			toAdminPage() {
				wx.reLaunch({
					url: "/pages/admin/index"
				})
			}
		}
	}
</script>

<style lang="scss">
	.u-border-bottom:after {
		border-bottom-width: 0px;
	}

	.logo-box {
		width: 100%;
		height: 200rpx;
		margin-top: 200rpx;
	}

	.form-box {
		width: 570rpx;
		margin: 100rpx 90rpx;
		box-sizing: border-box;
	}

	.form-id {
		width: 100%;
		height: 90rpx;
		border-radius: 45rpx;
		background-color: #ffffff;
	}

	.form-password {
		margin-top: 44rpx;
		width: 100%;
		height: 90rpx;
		border-radius: 45rpx;
		background-color: #ffffff;
	}

	.confirm-btn {
		width: 100%;
		margin-top: 54rpx;
	}

	.arrow-box {
		height: 120rpx;
		width: 120rpx;
		border-radius: 60rpx;
		background-color: #5677fc;
	}

	.visitor-login-ctx {
		width: 100%;
		text-align: center;
		color: #cccccc;
		font-size: 24rpx;
		margin-top: 150rpx;
	}
</style>
