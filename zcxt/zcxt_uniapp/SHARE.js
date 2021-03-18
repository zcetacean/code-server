export default {
	data() {
		return {
			//设置默认的分享参数
			share: {
				title: 'GPNU综测系统',
				path: '/pages/other/login',
				imageUrl: './static/custom/logo.png',
			}
		}
	},
	onShareAppMessage(res) {
		return {
			title: this.share.title,
			path: this.share.path,
			imageUrl: this.share.imageUrl,
			success(res) {
				uni.showToast({
					title: '分享成功'
				})
			},
			fail(res) {
				uni.showToast({
					title: '分享失败',
					icon: 'none'
				})
			}
		}
	}
}
