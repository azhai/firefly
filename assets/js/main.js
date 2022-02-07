layui.config({
	version: "318",
	base: "assets/module/"
}).extend({
	steps: "steps/steps",
	notice: "notice/notice",
	cascader: "cascader/cascader",
	// dropdown: "dropdown/dropdown",
	fileChoose: "fileChoose/fileChoose",
	Split: "Split/Split",
	Cropper: "Cropper/Cropper",
	tagsInput: "tagsInput/tagsInput",
	citypicker: "city-picker/city-picker",
	introJs: "introJs/introJs",
	zTree: "zTree/zTree"
}).use(["layer", "setter", "index", "admin"], function() {
	var d = layui.jquery;
	var c = layui.layer;
	var e = layui.setter;
	var b = layui.index;
	var a = layui.admin;
	if (!e.getToken()) {}
	a.req("userInfo.json", function(f) {
		if (200 === f.code) {
			e.putUser(f.user);
			a.renderPerm();
			d("#huName").text(f.user.nickName)
		} else {
			c.msg("获取用户失败", {
				icon: 2,
				anim: 6
			})
		}
	});
	a.req("menus.json", function(f) {
		b.regRouter(f);
		b.renderSide(f);
		b.loadHome({
			url: "#/console/workplace",
			name: '<i class="layui-icon layui-icon-home"></i>'
		})
	})
});