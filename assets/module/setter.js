/** EasyWeb spa v3.1.8 date:2020-03-11 */
layui.define(["table"], function(a) {
	var c = {
		baseServer: "./json/",
		pageTabs: false,
		cacheTab: true,
		defaultTheme: "",
		openTabCtxMenu: true,
		maxTabNum: 20,
		viewPath: "components",
		viewSuffix: ".html",
		reqPutToPost: true,
		apiNoCache: true,
		tableName: "firefly",
		getToken: function() {
			var d = layui.data(c.tableName);
			if (d) {
				return d.token
			}
		},
		removeToken: function() {
			layui.data(c.tableName, {
				key: "token",
				remove: true
			})
		},
		putToken: function(d) {
			layui.data(c.tableName, {
				key: "token",
				value: d
			})
		},
		getUser: function() {
			var d = layui.data(c.tableName);
			if (d) {
				return d.loginUser
			}
		},
		putUser: function(d) {
			layui.data(c.tableName, {
				key: "loginUser",
				value: d
			})
		},
		getUserAuths: function() {
			var e = [],
				d = c.getUser();
			var g = d ? d.authorities : [];
			for (var f = 0; f < g.length; f++) {
				e.push(g[f].authority)
			}
			return e
		},
		getAjaxHeaders: function(d) {
			var f = [];
			var e = c.getToken();
			if (e) {
				f.push({
					name: "Authorization",
					value: "Bearer " + e.access_token
				})
			}
			return f
		},
		ajaxSuccessBefore: function(e, d, f) {
			if (e.code === 401) {
				c.removeToken();
				layui.layer.msg("登录过期", {
					icon: 2,
					anim: 6,
					time: 1500
				}, function() {
					location.replace("components/template/login/login.html")
				});
				return false
			}
			return true
		},
		routerNotFound: function(d) {
			layui.layer.alert('路由<span class="text-danger">' + d.path.join("/") + "</span>不存在", {
				title: "提示",
				offset: "30px",
				skin: "layui-layer-admin",
				btn: [],
				anim: 6,
				shadeClose: true
			})
		}
	};
	var b = c.getToken();
	if (b && b.access_token) {
		layui.table.set({
			headers: {
				"Authorization": "Bearer " + b.access_token
			}
		})
	}
	c.base_server = c.baseServer;
	a("setter", c)
});
