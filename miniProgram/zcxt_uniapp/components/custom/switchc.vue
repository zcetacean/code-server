<template>
	<view>
		<view class="weui-switch" :class="{'weui-switch-on' : isChecked}" :value="value" :sid="index" @click="toggle" :style="{color:bgcolor, width: width+'rpx'}">
			<view v-if="isChecked && direction.length > 0" class="switch-checked" style="font-size: 28rpx;">
				{{direction[0]}}
			</view>
			<view v-if="!isChecked && direction.length > 0" class="switch-ischecked" style="font-size: 28rpx;">
				{{direction[1]}}
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		name: 'switchc',
		props: {
			value: {
				type: Boolean,
				default: true
			},
			//背景颜色
			// bgcolor: {
			// 	type: String,
			// 	default: "#00bfff"
			// },
			//宽度 rpx
			width: {
				type: Number,
				default: 104
			},
			text: {
				type: String,
				default: ''
			},
			sid: {
				type: Number,
				default: 0,
			}
		},
		data() {
			return {
				isChecked: this.value,
			}
		},
		computed: {
			direction() {
				if (this.text) {
					return this.text.split('|')
				} else {
					return []
				}
			}
		},
		watch: {
			value(val) {
				this.isChecked = val
			},
			isChecked(val) {
				let textTemp = '';
				if (val) {
					textTemp = this.text.split('|')[0];
				}
				if (!val) {
					textTemp = this.text.split('|')[1];
				}
				let ob = {
					sid: this.sid,
					value: val,
					text: textTemp
				}
				this.$emit('change', ob);
			}
		},
		methods: {
			toggle(e) {
				this.isChecked = !this.isChecked;
			}
		}
	}
</script>

<style>
	.weui-switch {
		position: relative;
		display: block;
		position: relative;
		width: 300rpx;
		height: 48rpx;
		outline: 0;
		border-radius: 24rpx;
		box-sizing: border-box;
		background-color: #DFDFDF;
		cursor: pointer;
	}

	.weui-switch:before {
		content: " ";
		position: absolute;
		top: 0;
		left: 0;
		width: 100rpx;
		height: 44rpx;
		border-radius: 22rpx;
		background-color: #FDFDFD;
	}

	.weui-switch:after {
		content: " ";
		position: absolute;
		top: 0;
		left: 0;
		width: 44rpx;
		height: 44rpx;
		border-radius: 22rpx;
		background-color: #FFFFFF;
		box-shadow: 0 2rpx 5rpx rgba(0, 0, 0, 0.4);
	}

	.weui-switch-on {
		border-color: #353535;
		background-color: #1AAD19;
	}

	/* 背景颜色设计 */
	.weui-switch-on:before {
		border-color: #353535;
		background-color: #09BB07;
	}

	.weui-switch-on:after {
		border-color: #fcc038;
		transform: translateX(58rpx);
	}

	.switch-checked {
		width: 100%;
		height: 100%;
		position: absolute;
		padding-left: 20rpx;
		line-height: 44rpx;
		color: #FFF;
		user-select: none;
	}

	.switch-ischecked {
		width: 100%;
		height: 100%;
		position: absolute;
		padding-left: 60rpx;
		line-height: 44rpx;
		color: #7A7A7A;
		user-select: none;
	}
</style>
