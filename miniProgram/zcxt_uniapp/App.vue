<script>
	export default {
		onLaunch: function() {},
		onShow: async function() {},
		onHide: function() {}
	}
</script>
<style lang="scss">
	page {
		background-color: #ffffff;
	}
	
	@import "uview-ui/index.scss";
	/* css统一注册接口 */
	@import './static/thorui/thorui.css';
	@import './static/thorui/icon.css';
	
	.cu-flex-center{
		display: flex;
		align-items: center;
		justify-content: center;
	}
	
	.cu-flex-center-y{
		display: flex;
		align-items: center;
	}
	
	.cu-flex-center-x{
		display: flex;
		justify-content: center;
	}
	
	.cu-shadow-m{
		box-shadow: 1rpx 1rpx 10rpx 1rpx rgba(0, 0, 0, 0.10);
	}
	
	.cu-height-100{
		height: 100%;
	}
	
	.cu-margin-bottom-l{
		margin-bottom:80rpx;
	}
	
	.cu-clear-btn-border::after{
		border:none
	}
</style>
