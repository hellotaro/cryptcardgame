(window.webpackJsonp=window.webpackJsonp||[]).push([[6],{250:function(t,r,o){var content=o(257);content.__esModule&&(content=content.default),"string"==typeof content&&(content=[[t.i,content,""]]),content.locals&&(t.exports=content.locals);(0,o(105).default)("3b52ba20",content,!0,{sourceMap:!1})},256:function(t,r,o){"use strict";o(250)},257:function(t,r,o){var n=o(104)((function(i){return i[1]}));n.push([t.i,"/*purgecss start ignore*/\n.card[data-v-72a76455]{\n  margin:0.25rem;\n  max-width:24rem;\n  border-radius:0.5rem;\n  --tw-bg-opacity:1;\n  background-color:rgba(255, 255, 255, var(--tw-bg-opacity));\n  --tw-shadow:0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);\n  box-shadow:var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow)\n}\n.card-button[data-v-72a76455]{\n  display:inline-block;\n  border-radius:0.25rem;\n  --tw-bg-opacity:1;\n  background-color:rgba(37, 99, 235, var(--tw-bg-opacity));\n  padding-left:1.5rem;\n  padding-right:1.5rem;\n  padding-top:0.625rem;\n  padding-bottom:0.625rem;\n  font-size:0.75rem;\n  line-height:1rem;\n  font-weight:500;\n  text-transform:uppercase;\n  line-height:1.25;\n  --tw-text-opacity:1;\n  color:rgba(255, 255, 255, var(--tw-text-opacity));\n  --tw-shadow:0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);\n  box-shadow:var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);\n  transition-property:background-color, border-color, color, fill, stroke, opacity, box-shadow, transform, filter, -webkit-backdrop-filter;\n  transition-property:background-color, border-color, color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;\n  transition-property:background-color, border-color, color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter, -webkit-backdrop-filter;\n  transition-timing-function:cubic-bezier(0.4, 0, 0.2, 1);\n  transition-duration:150ms;\n  transition-duration:150ms;\n  transition-timing-function:cubic-bezier(0.4, 0, 0.2, 1)\n}\n.card-button[data-v-72a76455]:hover{\n  --tw-bg-opacity:1;\n  background-color:rgba(29, 78, 216, var(--tw-bg-opacity))\n}\n\n/*purgecss end ignore*/",""]),n.locals={},t.exports=n},272:function(t,r,o){"use strict";o.r(r);var n={name:"GameCard",data:function(){return{data:"mycard"}},props:["title","img_src","content"],methods:{clickGameStartBtn:function(){this.data+="yours?",console.log(this.data)}}},e=(o(256),o(42)),component=Object(e.a)(n,(function(){var t=this,r=t.$createElement,o=t._self._c||r;return o("div",{staticClass:"card"},[o("a",{attrs:{href:"#!"}},[o("img",{staticClass:"rounded-t-lg",attrs:{src:t.img_src,alt:""}})]),t._v(" "),o("div",{staticClass:"p-6"},[o("h5",{staticClass:"mb-2 text-xl font-medium text-gray-900"},[t._v(t._s(t.title))]),t._v(" "),o("p",{staticClass:"mb-4 text-base text-gray-700"},[t._v(t._s(t.content))]),t._v(" "),o("button",{staticClass:"card-button",attrs:{type:"button"},on:{click:t.clickGameStartBtn}},[t._v("Start")])])])}),[],!1,null,"72a76455",null);r.default=component.exports}}]);