(window.webpackJsonp=window.webpackJsonp||[]).push([[0],{21:function(e,t,a){e.exports=a.p+"static/media/grassFedBanana.8a35fba3.png"},22:function(e,t,a){e.exports=a.p+"static/media/friedBanana.a85345b1.jpg"},23:function(e,t,a){e.exports=a.p+"static/media/frozenBanana.12a7f8ba.jpg"},25:function(e,t,a){e.exports=a(38)},36:function(e,t,a){},38:function(e,t,a){"use strict";a.r(t);var n=a(0),c=a.n(n),l=a(18),r=a.n(l),i=a(5),s=a(6),m=a(8),o=a(7),u=a(9),d=a(41),p=a(43),E=a(42),f=a(40),b=function(){return c.a.createElement("nav",{className:"nav-wrapper"},c.a.createElement("div",{className:"container"},c.a.createElement(f.a,{to:"/",className:"brand-logo"},"Table Pay"),c.a.createElement("ul",{className:"right"})))},v=a(10),h=function(){return c.a.createElement("div",{class:"row"},c.a.createElement("form",{class:"col s12"},c.a.createElement("div",{class:"row"},c.a.createElement("div",{class:"input-field col s12"},c.a.createElement("input",{id:"card",type:"text",class:"validate"}),c.a.createElement("label",{for:"card"},"Card Number"))),c.a.createElement("div",{class:"row"},c.a.createElement("div",{class:"input-field col s6"},c.a.createElement("input",{id:"first_name",type:"text",class:"validate"}),c.a.createElement("label",{for:"first_name"},"Expiration Month")),c.a.createElement("div",{class:"input-field col s6"},c.a.createElement("input",{id:"last_name",type:"text",class:"validate"}),c.a.createElement("label",{for:"last_name"},"Expiration Year")))))},O=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,c=new Array(n),l=0;l<n;l++)c[l]=arguments[l];return(a=Object(m.a)(this,(e=Object(o.a)(t)).call.apply(e,[this].concat(c)))).handleCheckOut=function(){a.props.checkOut()},a.handleGetOrderDetails=function(){a.props.getOrderDetails()},a}return Object(u.a)(t,e),Object(s.a)(t,[{key:"render",value:function(){var e=this.props.items.length?this.props.items.map(function(e){return c.a.createElement("li",{className:"collection-item avatar",key:e.id},c.a.createElement("div",{className:"item-img"},c.a.createElement("img",{src:e.img,alt:e.img,className:""})),c.a.createElement("div",{className:"item-desc"},c.a.createElement("span",{className:"title"},e.title),c.a.createElement("p",null,e.desc),c.a.createElement("p",null,c.a.createElement("b",null,"Price: ",e.price,"$")),c.a.createElement("p",null,c.a.createElement("b",null,"Quantity: ",e.quantity))))}):c.a.createElement("p",null," You have not ordered anything yet");return c.a.createElement("div",{className:"container"},c.a.createElement("div",{className:"cart"},c.a.createElement("h3",null,"Order Summary"),c.a.createElement("ul",{className:"collection"},e)),c.a.createElement("div",{className:"collection"},c.a.createElement("li",{className:"collection-item"},c.a.createElement("b",null,"Total: ",this.props.total," $"))),c.a.createElement("h3",null,"Payment Info"),c.a.createElement(h,null),c.a.createElement("div",{className:"checkout"},c.a.createElement("button",{className:"waves-effect waves-light btn",handleClick:this.handleCheckOut},"Checkout"),c.a.createElement("button",{className:"waves-effect waves-light btn",style:{marginLeft:"15px"}},"Refresh")))}}]),t}(n.Component),g=Object(v.b)(function(e){return{items:e.items,total:e.total}},function(e){return{checkOut:function(){var t;e({type:"CHECK_OUT",id:t})},getOrderDetails:function(){var t;e({type:"GET_ORDER_DETAILS",id:t})}}})(O),y=function(e){function t(){return Object(i.a)(this,t),Object(m.a)(this,Object(o.a)(t).apply(this,arguments))}return Object(u.a)(t,e),Object(s.a)(t,[{key:"render",value:function(){return c.a.createElement("div",{className:"container"},c.a.createElement(g,null))}}]),t}(n.Component),j=function(e){function t(){return Object(i.a)(this,t),Object(m.a)(this,Object(o.a)(t).apply(this,arguments))}return Object(u.a)(t,e),Object(s.a)(t,[{key:"render",value:function(){return c.a.createElement("div",{className:"container",style:{width:"500px"}},c.a.createElement("h3",null,"Enter Your Table Number"),c.a.createElement("div",{class:"row"},c.a.createElement("form",{class:"col s12"},c.a.createElement("div",{class:"row"},c.a.createElement("div",{class:"input-field col s12"},c.a.createElement("input",{id:"table-num",type:"text",class:"validate"}),c.a.createElement("label",{for:"table-num"}))))),c.a.createElement("button",{className:"waves-effect waves-light btn",style:{marginLeft:"15px"}},"Submit"))}}]),t}(n.Component),N=Object(v.b)(function(e){return{items:e.items,total:e.total}},function(e){return{}})(j),w=function(e){function t(){return Object(i.a)(this,t),Object(m.a)(this,Object(o.a)(t).apply(this,arguments))}return Object(u.a)(t,e),Object(s.a)(t,[{key:"render",value:function(){return c.a.createElement(d.a,null,c.a.createElement("div",{className:"App"},c.a.createElement(b,null),c.a.createElement(p.a,null,c.a.createElement(E.a,{exact:!0,path:"/orderview",component:y}),c.a.createElement(E.a,{exact:!0,path:"/",component:N}))))}}]),t}(n.Component),x=(a(36),a(21)),k=a.n(x),C=a(22),_=a.n(C),T=a(23),D=a.n(T),B={items:[{id:1,title:"Grass Fed Organic Banana",desc:"Lorem ipsum dolor sit amet consectetur adipisicing elit. Minima, ex.",price:110,img:k.a},{id:2,title:"Deep Fried Banana",desc:"Lorem ipsum dolor sit amet consectetur adipisicing elit. Minima, ex.",price:80,img:_.a},{id:3,title:"Frozen Banana",desc:"Lorem ipsum dolor sit amet consectetur adipisicing elit. Minima, ex.",price:120,img:D.a}],addedItems:[],total:0},L=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:B,t=arguments.length>1?arguments[1]:void 0;if("CHECK_OUT"===t.type);else if("GET_ORDER_DETAILS"!==t.type)return e},I=a(11),R=Object(I.b)(L);r.a.render(c.a.createElement(v.a,{store:R},c.a.createElement(w,null)),document.getElementById("root"))}},[[25,2,1]]]);
//# sourceMappingURL=main.5d0672fb.chunk.js.map