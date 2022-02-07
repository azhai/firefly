layui.define(function(b) {
	var a = {
		index: "/",
		lash: null,
		routers: {},
		init: function(d) {
			a.index = a.routerInfo(d.index).path.join("/");
			if (d.pop && typeof d.pop === "function") {
				a.pop = d.pop
			}
			if (d.notFound && typeof d.notFound === "function") {
				a.notFound = d.notFound
			}
			c();
			window.onhashchange = function() {
				c()
			};
			return this
		},
		reg: function(f, e) {
			if (f) {
				if (!e) {
					e = function() {}
				}
				if (f instanceof Array) {
					for (var d in f) {
						this.reg.apply(this, [f[d], e])
					}
				} else {
					if (typeof f === "string") {
						f = a.routerInfo(f).path.join("/");
						if (typeof e === "function") {
							a.routers[f] = e
						} else {
							if (typeof e === "string" && a[e]) {
								a.routers[f] = a.routers[e]
							}
						}
					}
				}
			}
			return this
		},
		routerInfo: function(d) {
			d || (d = location.hash);
			var e = d.replace(/^#+/g, "").replace(/\/+/g, "/");
			if (e.indexOf("/") !== 0) {
				e = "/" + e
			}
			return layui.router("#" + e)
		},
		refresh: function(d) {
			c(d, true)
		},
		go: function(d) {
			location.hash = "#" + a.routerInfo(d).href
		}
	};

	function c(e, f) {
		var d = a.routerInfo(e);
		a.lash = d.href;
		var g = d.path.join("/");
		if (!g || g === "/") {
			g = a.index;
			d = a.routerInfo(a.index)
		}
		a.pop && a.pop.call(this, d);
		if (a.routers[g]) {
			d.refresh = f;
			a.routers[g].call(this, d)
		} else {
			if (a.notFound) {
				a.notFound.call(this, d)
			}
		}
	}
	b("layRouter", a)
});