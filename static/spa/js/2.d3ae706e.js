(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([[2],{"70f8":function(e,t,a){},"8b24":function(e,t,a){"use strict";a.r(t);var c=a("7a23"),n=a("ff16"),o=a.n(n);const r=Object(c["E"])("data-v-3dae397e");Object(c["u"])("data-v-3dae397e");const s=Object(c["f"])("img",{alt:"gopher",src:o.a},null,-1);Object(c["s"])();const p=r(((e,t,a,n,o,p)=>{const d=Object(c["x"])("q-btn"),b=Object(c["x"])("q-card-section"),f=Object(c["x"])("q-card"),l=Object(c["x"])("q-page");return Object(c["r"])(),Object(c["d"])(l,{class:"flex flex-center"},{default:r((()=>[Object(c["f"])(f,{class:"my-card"},{default:r((()=>[s,Object(c["f"])(b,{class:"q-pt-none text-center"},{default:r((()=>[Object(c["f"])(d,{class:"q-mt-xl",label:"New Gopher",color:"primary",onClick:t[1]||(t[1]=t=>e.newGopher())})])),_:1})])),_:1}),Object(c["e"])(" "+Object(c["z"])(e.gophers),1)])),_:1})}));var d=a("bc3a"),b=a.n(d),f=Object(c["g"])({name:"PageIndex",data(){return{gophers:[],newGopher(){console.log("trying to get data"),b.a.get("https://japanofile.herokuapp.com/api/v1/gopher").then((e=>this.gophers=e.data)).catch((e=>console.log(e)))}}}}),l=(a("d9fb"),a("9989")),g=a("f09f"),i=a("a370"),j=a("9c40"),u=a("eebe"),O=a.n(u);f.render=p,f.__scopeId="data-v-3dae397e";t["default"]=f;O()(f,"components",{QPage:l["a"],QCard:g["a"],QCardSection:i["a"],QBtn:j["a"]})},d9fb:function(e,t,a){"use strict";a("70f8")},ff16:function(e,t,a){e.exports=a.p+"img/gopher.534124c7.png"}}]);