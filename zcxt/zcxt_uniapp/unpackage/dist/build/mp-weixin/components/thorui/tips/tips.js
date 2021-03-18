(global["webpackJsonp"]=global["webpackJsonp"]||[]).push([["components/thorui/tips/tips"],{"37a9":function(t,n,e){"use strict";var i;e.d(n,"b",(function(){return u})),e.d(n,"c",(function(){return r})),e.d(n,"a",(function(){return i}));var u=function(){var t=this,n=t.$createElement;t._self._c},r=[]},"3d6e":function(t,n,e){"use strict";e.r(n);var i=e("6219"),u=e.n(i);for(var r in i)"default"!==r&&function(t){e.d(n,t,(function(){return i[t]}))}(r);n["default"]=u.a},6219:function(t,n,e){"use strict";Object.defineProperty(n,"__esModule",{value:!0}),n.default=void 0;var i={name:"tuiTips",props:{position:{type:String,default:"top"}},data:function(){return{timer:null,show:!1,msg:"无法连接到服务器~",type:"translucent"}},methods:{showTips:function(t){var n=this,e=t.type,i=void 0===e?"translucent":e,u=t.duration,r=void 0===u?2e3:u;clearTimeout(this.timer),this.show=!0,this.type=i,this.msg=t.msg,this.timer=setTimeout((function(){n.show=!1,clearTimeout(n.timer),n.timer=null}),r)}}};n.default=i},"8aa8":function(t,n,e){},"9d67":function(t,n,e){"use strict";e.r(n);var i=e("37a9"),u=e("3d6e");for(var r in u)"default"!==r&&function(t){e.d(n,t,(function(){return u[t]}))}(r);e("dbf5");var o,a=e("f0c5"),s=Object(a["a"])(u["default"],i["b"],i["c"],!1,null,null,null,!1,i["a"],o);n["default"]=s.exports},dbf5:function(t,n,e){"use strict";var i=e("8aa8"),u=e.n(i);u.a}}]);
;(global["webpackJsonp"] = global["webpackJsonp"] || []).push([
    'components/thorui/tips/tips-create-component',
    {
        'components/thorui/tips/tips-create-component':(function(module, exports, __webpack_require__){
            __webpack_require__('543d')['createComponent'](__webpack_require__("9d67"))
        })
    },
    [['components/thorui/tips/tips-create-component']]
]);
