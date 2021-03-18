(global["webpackJsonp"]=global["webpackJsonp"]||[]).push([["components/thorui/tui-tabbar/tui-tabbar"],{4537:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var u={name:"tuiTabbar",props:{current:{type:Number,default:0},color:{type:String,default:"#666"},selectedColor:{type:String,default:"#5677FC"},backgroundColor:{type:String,default:"#FFFFFF"},hump:{type:Boolean,default:!1},isFixed:{type:Boolean,default:!0},tabBar:{type:Array,default:function(){return[]}},badgeColor:{type:String,default:"#fff"},badgeBgColor:{type:String,default:"#F74D54"},unlined:{type:Boolean,default:!1}},watch:{current:function(){}},data:function(){return{}},methods:{tabbarSwitch:function(t,e,n,u){this.$emit("click",{index:t,hump:e,pagePath:n,verify:u})}}};e.default=u},"4ff9":function(t,e,n){"use strict";n.r(e);var u=n("4537"),a=n.n(u);for(var r in u)"default"!==r&&function(t){n.d(e,t,(function(){return u[t]}))}(r);e["default"]=a.a},6492:function(t,e,n){"use strict";n.r(e);var u=n("a7c4"),a=n("4ff9");for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);n("fbff");var f,o=n("f0c5"),c=Object(o["a"])(a["default"],u["b"],u["c"],!1,null,null,null,!1,u["a"],f);e["default"]=c.exports},a7c4:function(t,e,n){"use strict";var u;n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return u}));var a=function(){var t=this,e=t.$createElement;t._self._c},r=[]},cb3e:function(t,e,n){},fbff:function(t,e,n){"use strict";var u=n("cb3e"),a=n.n(u);a.a}}]);
;(global["webpackJsonp"] = global["webpackJsonp"] || []).push([
    'components/thorui/tui-tabbar/tui-tabbar-create-component',
    {
        'components/thorui/tui-tabbar/tui-tabbar-create-component':(function(module, exports, __webpack_require__){
            __webpack_require__('543d')['createComponent'](__webpack_require__("6492"))
        })
    },
    [['components/thorui/tui-tabbar/tui-tabbar-create-component']]
]);
