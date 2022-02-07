/** 数据列表组件 date:2020-05-04   */
layui.define(["laytpl","laypage","form"],function(f){var g=layui.jquery;var j=layui.laytpl;var e=layui.laypage;var b=layui.form;var k="DataGrid";var a="ew-datagrid-loading";var m="ew-datagrid-item";var d="ew-loading",n="ew-more-end";var i={limit:10,layout:["prev","page","next","skip","count","limit"]};var h={first:true,curr:1,limit:10,text:"加载更多",loadingText:"加载中...",noMoreText:"没有更多数据了~",errorText:"加载失败，请重试"};var o=function(p){this.options=g.extend(true,{method:"GET",request:{pageName:"page",limitName:"limit"},useAdmin:false,showError:function(q){g(this.elem).empty()},showEmpty:function(q){g(this.elem).empty()},showLoading:function(){g(this.elem).addClass(a)},hideLoading:function(){g(this.elem).removeClass(a)}},p);if(p.page){this.options.page=g.extend({},i,p.page===true?{}:p.page)}if(p.loadMore){this.options.loadMore=g.extend({},h,p.loadMore===true?{}:p.loadMore)}if(typeof this.options.data==="string"){this.options.url=this.options.data;this.options.data=undefined}this.init();this.bindEvents()};o.prototype.init=function(){var s=this;var q=this.options;var r=this.getComponents();if("static"===r.$elem.css("position")){r.$elem.css("position","relative")}if(q.checkAllElem){var p=r.$checkAll.find('input[lay-filter="'+r.checkAllFilter+'"]');p.next(".layui-form-checkbox").remove();p.remove();r.$checkAll.append(['<input type="checkbox"',' lay-filter="',r.checkAllFilter,'"',' lay-skin="primary" class="ew-datagrid-checkbox" />'].join(""));if(!r.$checkAll.hasClass("layui-form")){r.$checkAll.addClass("layui-form")}if(!r.$checkAll.attr("lay-filter")){r.$checkAll.attr("lay-filter",q.checkAllElem.substring(1))}b.render("checkbox",r.$checkAll.attr("lay-filter"))}if(q.url){q.reqData=function(t,u){if(!q.where){q.where={}}q.where[q.request.pageName]=t.page;q.where[q.request.limitName]=t.limit;(q.useAdmin?layui.admin:g).ajax({url:q.url,data:q.contentType&&q.contentType.indexOf("application/json")===0?JSON.stringify(q.where):q.where,headers:q.headers,type:q.method,dataType:"json",contentType:q.contentType,success:function(v){u(q.parseData?q.parseData(v):v)},error:function(v){u({code:v.status,msg:v.statusText,xhr:v})}})}}else{if(q.data){q.reqData=undefined;if(q.loadMore){s.renderLoadMore();s.changeLoadMore(2);s.renderBody(q.data,0,false,true);q.done&&q.done(q.data,1,q.data.length)}else{if(q.page){q.page.count=q.data.length;q.page.jump=function(w,x){q.showLoading();var y=(w.curr-1)*q.page.limit;var u=y+q.page.limit;if(u>q.data.length){u=q.data.length}var t=[];for(var v=y;v<u;v++){t.push(q.data[v])}q.page.data=t;s.renderBody(t,(w.curr-1)*w.limit,false,true);q.hideLoading();if(q.data.length===0){q.showEmpty&&q.showEmpty({})}q.done&&q.done(t,w.curr,w.count)};s.renderPage()}else{s.renderBody(q.data,0,false,true);if(q.data.length===0){q.showEmpty&&q.showEmpty({})}q.done&&q.done(q.data,1,q.data.length)}}}}if(!q.reqData){return}if(q.loadMore){s.renderLoadMore().click(function(){if(g(this).hasClass(d)){return}if(q.loadMore.first){q.loadMore.first=false}else{q.loadMore.curr++}s.changeLoadMore(1);q.reqData({page:q.loadMore.curr,limit:q.loadMore.limit},function(t){if(t.code!=0){s.changeLoadMore(3);q.loadMore.curr--;return}s.changeLoadMore(0);s.renderBody(t.data,(q.loadMore.curr-1)*q.loadMore.limit,q.loadMore.curr!==1);q.done&&q.done(t.data,q.loadMore.curr,t.count||t.data.length);if(!t.data||t.data.length<q.loadMore.limit){s.changeLoadMore(2)}})}).trigger("click")}else{if(q.page){q.showLoading();q.reqData({page:1,limit:q.page.limit},function(t){q.hideLoading();if(typeof t==="string"||!t.data){return q.showError&&q.showError(t)}if(t.data.length===0){return q.showEmpty&&q.showEmpty(t)}q.page.count=t.count;q.page.jump=function(u,v){if(v){return}q.showLoading();q.reqData({page:u.curr,limit:u.limit},function(w){q.hideLoading();if(typeof w==="string"||!w.data){return q.showError&&q.showError(w)}if(w.data.length===0){return q.showEmpty&&q.showEmpty(w)}s.renderBody(w.data,(u.curr-1)*u.limit);q.done&&q.done(w.data,u.curr,u.count)})};s.renderPage();s.renderBody(t.data);q.done&&q.done(t.data,1,t.count)})}else{q.showLoading();q.reqData({},function(t){q.hideLoading();if(t.code!=0){return q.showError&&q.showError(t)}if(!t.data||t.data.length===0){return q.showEmpty&&q.showEmpty(t)}s.renderBody(t.data);q.done&&q.done(t.data,1,t.data.length)})}}};o.prototype.bindEvents=function(){var r=this;var q=this.getComponents();var p=function(u){var s=g(this);if(!s.hasClass(m)){var w=s.parent("."+m);s=w.length>0?w:s.parentsUntil("."+m).last().parent()}var t=s.data("index");var v={elem:s,data:r.getData(t),index:t,del:function(){r.del(t)},update:function(x,y){r.update(t,x,y)}};return g.extend(v,u)};q.$elem.off("click.item").on("click.item",">."+m,function(){layui.event.call(this,k,"item("+q.filter+")",p.call(this,{}))});q.$elem.off("dblclick.itemDouble").on("click.itemDouble",">."+m,function(){layui.event.call(this,k,"itemDouble("+q.filter+")",p.call(this,{}))});q.$elem.off("click.tool").on("click.tool","[lay-event]",function(t){layui.stope(t);var s=g(this);layui.event.call(this,k,"tool("+q.filter+")",p.call(this,{event:s.attr("lay-event")}))});b.on("radio("+q.radioFilter+")",function(s){var t=r.getData(s.value);t.LAY_CHECKED=true;layui.event.call(this,k,"checkbox("+q.filter+")",{checked:true,data:t})});b.on("checkbox("+q.checkboxFilter+")",function(t){var s=t.elem.checked;var u=r.getData(t.value);u.LAY_CHECKED=s;r.checkChooseAllCB();layui.event.call(this,k,"checkbox("+q.filter+")",{checked:s,data:u})});b.on("checkbox("+q.checkAllFilter+")",function(v){var u=v.elem.checked;var w=g(v.elem);var t=w.next(".layui-form-checkbox");if(!r.options.data||r.options.data.length<=0){w.prop("checked",false);t.removeClass("layui-form-checked");return}q.$elem.find('input[name="'+q.checkboxFilter+'"]').each(function(){var x=g(this);x.prop("checked",u);var y=x.next(".layui-form-checkbox");if(u){y.addClass("layui-form-checked")}else{y.removeClass("layui-form-checked")}});for(var s=0;s<r.options.data.length;s++){r.options.data[s].LAY_CHECKED=u}layui.event.call(this,k,"checkbox("+q.filter+")",{checked:u,type:"all"})})};o.prototype.getComponents=function(){var r=this;var p=g(r.options.elem);var q=p.attr("lay-filter");if(!q){q=r.options.elem.substring(1);p.attr("lay-filter",q)}return{$elem:p,templetHtml:g(r.options.templet).html(),$page:r.options.page&&r.options.page.elem?g("#"+r.options.page.elem):undefined,$loadMore:r.options.loadMore&&r.options.loadMore.elem?g("#"+r.options.loadMore.elem):undefined,filter:q,checkboxFilter:"ew_tb_checkbox_"+q,radioFilter:"ew_tb_radio_"+q,checkAllFilter:"ew_tb_checkbox_all_"+q,$checkAll:g(r.options.checkAllElem)}};o.prototype.renderBody=function(s,r,q,p){if(!s){s=[]}var x=this.options;var v=this.getComponents();if(!r){r=0}var u=[];for(var t=0;t<s.length;t++){var w=s[t];w.LAY_INDEX=t;w.LAY_NUMBER=t+r+1;w.LAY_CHECKBOX_ELEM=['<input type="checkbox" lay-skin="primary"',' name="',v.checkboxFilter,'"',' lay-filter="',v.checkboxFilter,'"',w.LAY_CHECKED?' checked="checked"':"",' class="ew-datagrid-checkbox"',' value="'+t+'" />'].join("");w.LAY_RADIO_ELEM=['<input type="radio"',' name="',v.radioFilter,'"',' lay-filter="',v.radioFilter,'"',w.LAY_CHECKED?' checked="checked"':"",' class="ew-datagrid-radio"',' value="',t,'" />'].join("");if(v.templetHtml===undefined){return console.error("DataGrid Error: Template ["+x.templet+"] not found")}j(v.templetHtml).render(w,function(y){u.push(y)})}if(q){if(!p){x.data=x.data.concat(s)}v.$elem.append(u.join(""))}else{if(!p){x.data=s}v.$elem.html(u.join(""))}this.initChildren(r);b.render("checkbox",v.filter);b.render("radio",v.filter);this.checkChooseAllCB()};o.prototype.initChildren=function(p){if(!p||!(this.options.page&&this.options.page.data)){p=0}this.getComponents().$elem.children().each(function(q){var r=g(this);r.attr("data-index",q);r.attr("data-number",q+p+1);r.addClass(m)})};o.prototype.renderPage=function(){var p=this.options;var q=this.getComponents();q.$elem.next(".ew-datagrid-page,.ew-datagrid-loadmore").remove();p.page.elem="ew-datagrid-page-"+p.elem.substring(1);q.$elem.after('<div class="ew-datagrid-page '+(p.page["class"]||"")+'" id="'+p.page.elem+'"></div>');e.render(p.page)};o.prototype.renderLoadMore=function(){var p=this.options;var q=this.getComponents();q.$elem.next(".ew-datagrid-page,.ew-datagrid-loadmore").remove();p.loadMore.elem="ew-datagrid-page-"+p.elem.substring(1);q.$elem.after(['<div id="',p.loadMore.elem,'" ','class="ew-datagrid-loadmore ',p.loadMore["class"]||"",'">',"   <div>",'      <span class="ew-icon-loading">','         <i class="layui-icon layui-icon-loading-1 layui-anim layui-anim-rotate layui-anim-loop"></i>',"      </span>",'      <span class="ew-loadmore-text">',p.loadMore.text,"</span>","   </div>","</div>"].join(""));return q.$elem.next()};o.prototype.changeLoadMore=function(s){var p=this.options;var r=this.getComponents();var q=r.$loadMore.find(".ew-loadmore-text");r.$loadMore.removeClass(d+" "+n);if(s===0){q.html(p.loadMore.text)}else{if(s===1){q.html(p.loadMore.loadingText);r.$loadMore.addClass(d)}else{if(s===2){q.html(p.loadMore.noMoreText);r.$loadMore.addClass(n)}else{q.html(p.loadMore.errorText)}}}};o.prototype.update=function(r,p,s){var v=this;var u=this.getComponents();var q=u.$elem.children('[data-index="'+r+'"]');var t=q.data("number");if(t-r!==1){g.extend(true,this.options.data[t-1],p)}else{g.extend(true,this.options.data[r],p)}if(2===s){return}j(u.templetHtml).render(v.getData(r),function(w){if(s===1){return q.html(g(w).html())}q.before(w).remove();v.initChildren(t-r-1)})};o.prototype.del=function(q){var s=this.getComponents();var p=s.$elem.children('[data-index="'+q+'"]');var r=p.data("number");p.remove();if(r-q!==1){this.options.data.splice(r-1,1)}else{this.options.data.splice(q,1)}this.initChildren(r-q-1)};o.prototype.getData=function(q){if(q===undefined){return this.options.data}var s=this.getComponents();var p=s.$elem.children('[data-index="'+q+'"]');var r=p.data("number");if(r-q!==1){return this.options.data[r-1]}return this.options.data[q]};o.prototype.checkStatus=function(){var t=this;var s=this.getComponents();var r=s.checkboxFilter;var v=s.radioFilter;var u=[];var p=s.$elem.find('input[name="'+v+'"]');if(p.length>0){var q=p.filter(":checked").val();if(q!==undefined){var w=t.getData(q);if(w){u.push(w)}}}else{s.$elem.find('input[name="'+r+'"]:checked').each(function(){var x=g(this).val();if(x!==undefined){var y=t.getData(x);if(y){u.push(y)}}})}return u};o.prototype.checkChooseAllCB=function(){var s=this.getComponents();var q=s.$checkAll.find('input[lay-filter="'+s.checkAllFilter+'"]');var p=this.options.data.length!==0;for(var r=0;r<this.options.data.length;r++){if(!this.options.data[r].LAY_CHECKED){p=false;break}}if(p){q.prop("checked",true);q.next(".layui-form-checkbox").addClass("layui-form-checked")}else{q.prop("checked",false);q.next(".layui-form-checkbox").removeClass("layui-form-checked")}};o.prototype.reload=function(p){if(p){if(p.page){if(this.options.page){p.page=g.extend({},this.options.page,p.page)}else{p.page=g.extend({},i,p.page)}if(this.options.loadMore){this.options.loadMore=undefined}}else{if(p.loadMore){if(this.options.loadMore){p.loadMore=g.extend({},this.options.loadMore,p.loadMore,{first:true,curr:1})}else{p.loadMore=g.extend({},h,p.loadMore)}if(this.options.page){this.options.page=undefined}}}g.extend(true,this.options,p)}this.init()};function c(p){if(!p){return}try{return new Function("return "+p)()}catch(q){console.error("element property data- configuration item has a syntax error: "+p)}}var l={render:function(p){if(p.onItemClick){l.onItemClick(p.elem,p.onItemClick)}if(p.onToolBarClick){l.onToolBarClick(p.elem,p.onToolBarClick)}return new o(p)},on:function(p,q){return layui.onevent.call(this,k,p,q)},onItemClick:function(q,p){if(q.indexOf("#")===0){q=q.substring(1)}return l.on("item("+q+")",p)},onToolBarClick:function(q,p){if(q.indexOf("#")===0){q=q.substring(1)}return l.on("tool("+q+")",p)}};l.autoRender=function(p){g(p||"[data-grid]").each(function(){try{var u=g(this);var s=u.attr("id");if(!s){s="ew-datagrid-"+(g('[id^="ew-datagrid-"]').length+1);u.attr("id",s)}var r=u.children("[data-grid-tpl]");if(r.length>0){r.attr("id",s+"-tpl");u.after(r);var q=c(u.attr("lay-data"));q.elem="#"+s;q.templet="#"+s+"-tpl";l.render(q)}}catch(t){console.error(t)}})};l.autoRender();g("head").append(['<style id="ew-css-datagrid">',".ew-datagrid-loadmore, .ew-datagrid-page {","    text-align: center;","}",".ew-datagrid-loadmore {","    color: #666;","    cursor: pointer;","}",".ew-datagrid-loadmore > div {","    padding: 12px;","}",".ew-datagrid-loadmore > div:hover {","    background-color: rgba(0, 0, 0, .03);","}",".ew-datagrid-loadmore .ew-icon-loading {","    margin-right: 6px;","    display: none;","}",".ew-datagrid-loadmore.",n," {","    pointer-events: none;","}",".ew-datagrid-loadmore.",d," .ew-icon-loading {","    display: inline;","}",".",a,":before {",'    content: "\\e63d";',"    font-family: layui-icon !important;","    font-size: 32px;","    color: #C3C3C3;","    position: absolute;","    left: 50%;","    top: 50%;","    margin-left: -16px;","    margin-top: -16px;","    z-index: 999;","    -webkit-animatione: layui-rotate 1s linear;","    animation: layui-rotate 1s linear;","    -webkit-animation-iteration-count: infinite;","    animation-iteration-count: infinite;","}",".ew-datagrid-checkbox + .layui-form-checkbox {","   padding-left: 18px;","}",".ew-datagrid-radio + .layui-form-radio {","   margin: 0;","   padding: 0;","   line-height: 22px;","}",".ew-datagrid-radio + .layui-form-radio .layui-icon {","   margin-right: 0;","}","</style>"].join(""));f("dataGrid",l)});