<template>
	<view>
		<view v-if="isShowShare" class="rolling-box cu-flex-center">
			<u-icon name="share-fill" color="#ffffff" size="44"></u-icon>
		</view>
		<button class="cu-clear-btn-border" v-if="isShowShare" open-type="share" style="position: absolute;width: 200rpx;height: 190rpx;z-index: 9999;background-color: rgba(0,0,0,0)"></button>
		<view class="bar-bg">
			<view class="bar-box">
				<view class="bar-text">
					{{title}}
				</view>
			</view>
		</view>
		<view class="wave-box">
			<image src="/static/custom/wave.gif" mode="scaleToFill" class="wave-gif"></image>
		</view>
	</view>
</template>

<script>
	export default {
		name: "topBox",
		props: {
			title: {
				type: String
			},
			isShowShare: {
				type: Boolean
			}
		}
	}
</script>

<style>
	@keyframes roll {
		0% {
			transform: rotate(0deg) translateY(-13rpx)
		}

		25% {
			transform: rotate(90deg) translateY(-0rpx)
		}

		50% {
			transform: rotate(180deg) translateY(-13rpx)
		}

		75% {
			transform: rotate(270deg) translateY(0rpx)
		}

		100% {
			transform: rotate(360deg) translateY(-13rpx)
		}
	}

	.rolling-box {
		animation: roll 4s infinite linear;
		z-index: 989;
		width: 80rpx;
		height: 80rpx;
		border-radius: 40rpx;
		background-color: rgba(255, 255, 255, 0.3);
		position: absolute;
		top: 90rpx;
		left: 60rpx;
	}

	.bar-bg {
		background-image: url('../../static/custom/topBlue.jpg');
		background-size: cover;
		background-position: center;
		background-repeat: no-repeat;
		padding-top: 40rpx;
		padding-bottom: 40rpx;
	}

	.bar-box {
		display: flex;
		position: relative;
		align-items: center;
		min-height: 100upx;
		justify-content: space-between;
	}

	.bar-text {
		position: absolute;
		text-align: center;
		width: calc(100% - 340upx);
		left: 0;
		right: 0;
		bottom: 0;
		top: 0;
		margin: auto;
		height: 60upx;
		font-size: 32upx;
		line-height: 60upx;
		cursor: none;
		pointer-events: none;
		text-overflow: ellipsis;
		white-space: nowrap;
		overflow: hidden;
		color: #ffffff;
		font-weight: bold;
	}

	.bar-box.ios .bar-text {
		bottom: 7px;
		height: 30px;
		font-size: 32upx;
		line-height: 30px;
	}

	.wave-gif {
		display: block;
		border: none;
		mix-blend-mode: screen;
		height: 100rpx;
		margin-top: -100rpx;
		width: 100%;
	}

	.wave-box {
		position: relative;
		margin-bottom: -2px;
	}

	.wave-box::before {
		content: "";
		display: block;
		background: inherit;
		filter: blur(10upx);
		position: absolute;
		width: 100%;
		height: 100%;
		top: 10upx;
		left: 10upx;
		z-index: -1;
		opacity: 0.4;
		transform-origin: 0 0;
		border-radius: inherit;
		transform: scale(1, 1);
	}
</style>
