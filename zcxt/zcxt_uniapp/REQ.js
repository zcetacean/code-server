export default {

	data: {

	},

	util: {
		saveUserInfo(treeName, name, type) {
			uni.setStorageSync("treeName", treeName)
			uni.setStorageSync("name", name)
			uni.setStorageSync("type", type)
		}
	},

	conf: {
		//baseUrl: "https://zcxt.soca.ink/proxy/beego/",
		baseUrl: "http://localhost:8080",
		isSuccess: function(type, e) {
			if (e.data.status === 200) {
				let result
				if (type === 'bool') {
					result = true
				} else if (type === 'data') {
					result = e.data.data
				} else {
					throw new Error("filter():filter is lose prototype[type]")
				}
				return result
			} else if (e.data.status === 135) {
				uni.showToast({
					icon: "none",
					title: "登录过期"
				})
				uni.clearStorageSync()
				uni.redirectTo({
					url: "/pages/other/login"
				})
				resolve(false)
			} else {
				return false
			}
		}
	},

	VerifyCreate: function(obj) {
		const that = this
		return new Promise(function(resolve, reject) {
			wx.uploadFile({
				url: that.conf.baseUrl + "/api/Verify/Create",
				filePath: obj.img,
				name: 'file',
				formData: {
					initTreeId: obj.initTreeId,
					id:obj.id,//ItemId
					type:obj.type,
					applyScore:obj.applyScore
				},
				header: {
					'Content-type': 'multipart/form-data',
					Cookie: uni.getStorageSync("cookieKey")
				},
				success: (res) => {
					console.log(res)
					resolve(true)
				}
			})
		})
	},

	VerifyRead: function(obj) {
		const that = this
		return new Promise(function(resolve, reject) {
			wx.request({
				url: that.conf.baseUrl + "/api/Verify/Read",
				method: "POST",
				header: {
					Cookie: uni.getStorageSync("cookieKey")
				},
				data: {
					state: obj.state
				},
				success: function(res) {
					resolve(that.conf.isSuccess('data', res))
				}
			})
		})
	},
	
	VerifyUpdate: function(obj) {
		const that = this
		return new Promise(function(resolve, reject) {
			wx.request({
				url: that.conf.baseUrl + "/api/Verify/Update",
				method: "POST",
				data: {
					type: obj.type,//yes no
					id:obj.id,
					remark:obj.remark,
					realScore:obj.realScore,
				},
				header: {
					Cookie: uni.getStorageSync("cookieKey")
				},
				success: function(res) {
					resolve(that.conf.isSuccess('bool', res))
				}
			})
		})
	},
	
	VerifyReadWait: function(obj) {
		const that = this
		return new Promise(function(resolve, reject) {
			wx.request({
				url: that.conf.baseUrl + "/api/Verify/ReadWait",
				method: "POST",
				header: {
					Cookie: uni.getStorageSync("cookieKey")
				},
				data: {
					name: obj.name
				},
				success: function(res) {
					resolve(that.conf.isSuccess('data', res))
				}
			})
		})
	},

	// 用户登录
	UserLogin: function(obj) {
		const that = this
		return new Promise(function(resolve, reject) {
			wx.request({
				url: that.conf.baseUrl + "/api/User/Login" + '?' + 'id=' + obj.id + '&' + 'password=' + obj.password + '&' +
					'state=' + obj.state,
				method: "GET",
				success: function(res) {
					if (res && res.header && res.header['Set-Cookie']) {
						uni.setStorageSync('cookieKey', res.header['Set-Cookie'])
					}
					resolve(that.conf.isSuccess('data', res))
				}
			})
		})
	},

	// 完善学生信息
	UserUpdate: function(obj) {
		const that = this
		return new Promise(function(resolve, reject) {
			wx.request({
				url: that.conf.baseUrl + "/api/User/Update",
				method: "POST",
				data: {
					name: obj.name,
					initTreeName: obj.initTreeName
				},
				header: {
					Cookie: uni.getStorageSync("cookieKey")
				},
				success: function(res) {
					resolve(that.conf.isSuccess('bool', res))
				}
			})
		})
	},

	// 获取操作列表
	RecordRead: function(obj) {
		const that = this
		return new Promise(function(resolve, reject) {
			wx.request({
				url: that.conf.baseUrl + "/api/Record/Read",
				method: "GET",
				header: {
					Cookie: uni.getStorageSync("cookieKey")
				},
				success: function(res) {
					resolve(that.conf.isSuccess('data', res))
				}
			})
		})
	},

	// 获取审核项列表
	UserReadItem: function(obj) {
		const that = this
		return new Promise(function(resolve, reject) {
			wx.request({
				url: that.conf.baseUrl + "/api/User/ReadItem",
				method: "GET",
				success: function(res) {
					resolve(that.conf.isSuccess('data', res))
				}
			})
		})
	},

	// 获取树列表
	UserReadOrganize: function(obj) {
		const that = this
		return new Promise(function(resolve, reject) {
			wx.request({
				url: that.conf.baseUrl + "/api/User/ReadOrganize",
				method: "GET",
				success: function(res) {
					if ((res.data.status === 200) && (res.data.data.length > 0)) {
						let arr = res.data.data.map((item) => {
							return item.name
						})
						resolve(Array.from(new Set(arr)))
					} else if (e.data.status === 135) {
						uni.showToast({
							icon: "none",
							title: "登录过期"
						})
						uni.clearStorageSync()
						uni.redirectTo({
							url: "/pages/other/login"
						})
						resolve(false)
					} else {
						resolve(false)
					}
				}
			})
		})
	},

}
