(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([[2],{"8b24":function(e,t,o){"use strict";o.r(t);var c=o("7a23");const a=Object(c["E"])("data-v-b3cac092");Object(c["v"])("data-v-b3cac092");const r={key:0,alt:"gopher",src:"https://crypto-gopher.herokuapp.com/gopher/gopher.png"};Object(c["t"])();const p=a(((e,t,o,p,n,s)=>{const h=Object(c["y"])("q-btn"),g=Object(c["y"])("q-card-section"),d=Object(c["y"])("q-card"),l=Object(c["y"])("q-page");return Object(c["s"])(),Object(c["d"])(l,{class:"flex flex-center"},{default:a((()=>[Object(c["g"])(d,{class:"my-card"},{default:a((()=>[e.cryptogopher?Object(c["e"])("",!0):(Object(c["s"])(),Object(c["d"])("img",r)),e.cryptogopher?(Object(c["s"])(),Object(c["d"])("img",{key:1,alt:"cryptogopher",src:e.gopherUrl},null,8,["src"])):Object(c["e"])("",!0),Object(c["g"])(g,{class:"q-pt-none text-center"},{default:a((()=>[Object(c["g"])(h,{class:"q-mt-xl",label:"New Cryptogopher",color:"primary",onClick:t[1]||(t[1]=t=>e.newGopher())})])),_:1})])),_:1})])),_:1})}));o("a79d"),o("e6cf");var n=o("bc3a"),s=o.n(n),h=o("b3fe"),g=Object(c["h"])({name:"PageIndex",setup(){const e=Object(h["a"])();return{showLoading(){e.loading.show()},hideLoading(){e.loading.hide()}}},data(){return{cryptogopher:!1,gopherSeed:null,gopherUrl:null,newGopher(){this.showLoading(),console.log("trying to get data"),s.a.post("https://crypto-gopher.herokuapp.com/api/v1/gopher",{name:"Test"}).then((e=>{this.gopherSeed=e.data.seed,this.gopherUrl="https://crypto-gopher.herokuapp.com/gopher/gopher_out_"+this.gopherSeed+".png",this.cryptogopher=!0})).catch((e=>console.log(e))).finally((()=>{this.hideLoading()}))}}}}),d=(o("a371"),o("9989")),l=o("f09f"),b=o("a370"),i=o("9c40"),u=o("eebe"),j=o.n(u);g.render=p,g.__scopeId="data-v-b3cac092";t["default"]=g;j()(g,"components",{QPage:d["a"],QCard:l["a"],QCardSection:b["a"],QBtn:i["a"]})},a371:function(e,t,o){"use strict";o("a895")},a895:function(e,t,o){}}]);