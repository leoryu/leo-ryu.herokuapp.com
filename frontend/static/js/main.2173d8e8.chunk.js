(window.webpackJsonp=window.webpackJsonp||[]).push([[0],{235:function(e,t,a){e.exports=a(480)},240:function(e,t,a){},480:function(e,t,a){"use strict";a.r(t);var n=a(0),r=a.n(n),c=a(21),o=a.n(c),i=(a(240),a(28)),l=a(29),s=a(32),m=a(30),p=a(33),u=a(40),d=a(2),g=a.n(d),h=a(48),b=a.n(h),f=a(9),E=a(62),v=a(41),y=a(59),j=a.n(y),O=a(16),w=a.n(O);var S=Object(f.withStyles)(function(e){return Object(f.createStyles)({toolbarMain:{borderBottom:"1px solid ".concat(e.palette.grey[300])},toolbarTitle:{flex:1},toolbarSecondary:{justifyContent:"space-evenly"}})})(function(e){var t=e.classes;return r.a.createElement(r.a.Fragment,null,r.a.createElement(j.a,{className:t.toolbarMain},r.a.createElement(w.a,{component:"h2",variant:"h5",color:"inherit",align:"center",noWrap:!0,className:t.toolbarTitle},"Leo's Blog")),r.a.createElement(j.a,{variant:"dense",className:t.toolbarSecondary}))}),C=a(27),k=a.n(C),x=a(22),N=a.n(x),P=a(73),T=a.n(P),D=a(74),A=a.n(D),B=a(75),F=a.n(B),I=a(61),M=a.n(I),R=a(60),q=a.n(R),W=a(164),L={id:"",title:"",abstract:"",content:"",created_at:0,edited_at:0},_="LeoBlogToken";var G=a(23),z=a.n(G),J=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,r=new Array(n),c=0;c<n;c++)r[c]=arguments[c];return(a=Object(s.a)(this,(e=Object(m.a)(t)).call.apply(e,[this].concat(r)))).state={papers:[L]},a.getPaper=function(){z.a.get("/api/papers?limit=10&page=1&").then(function(e){a.setState({papers:e.data})}).catch(function(e){console.log(e.response)})},a}return Object(p.a)(t,e),Object(l.a)(t,[{key:"componentDidMount",value:function(){this.getPaper()}},{key:"render",value:function(){var e=this.props.classes,t=this.state.papers;return r.a.createElement(r.a.Fragment,null,r.a.createElement(q.a,{component:function(e){return r.a.createElement(E.b,Object.assign({},e,{to:"/paper/"+t[0].id}))},underline:"none"},r.a.createElement(k.a,{className:e.mainFeaturedPost},r.a.createElement(N.a,{container:!0},r.a.createElement(N.a,{item:!0},r.a.createElement("div",{className:e.mainFeaturedPostContent},r.a.createElement(w.a,{component:"h1",variant:"h3",color:"inherit",gutterBottom:!0},t[0].title),r.a.createElement(M.a,{xsDown:!0},r.a.createElement(w.a,{variant:"h5",color:"inherit",paragraph:!0},t[0].abstract))))))),r.a.createElement(N.a,{container:!0,spacing:40,className:e.cardGrid},t.slice(1).map(function(t){return r.a.createElement(N.a,{item:!0,key:t.title,xs:12,md:12},r.a.createElement(q.a,{component:function(e){return r.a.createElement(E.b,Object.assign({},e,{to:"/paper/"+t.id}))},underline:"none"},r.a.createElement(W.a,null,r.a.createElement(T.a,{className:e.card},r.a.createElement("div",{className:e.cardDetails},r.a.createElement(A.a,null,r.a.createElement(w.a,{component:"h2",variant:"h5"},t.title),r.a.createElement(w.a,{variant:"subtitle1",color:"textSecondary",paragraph:!0},"Abstract: ",t.abstract),r.a.createElement(w.a,{variant:"subtitle1",color:"textSecondary"},new Date(1e3*t.created_at).toLocaleDateString()))),r.a.createElement(M.a,{xsDown:!0},r.a.createElement(F.a,{image:"",title:"Image title"}))))))})))}}]),t}(r.a.Component);J.propTypes={classes:g.a.object.isRequired};var $=Object(f.withStyles)(function(e){return Object(f.createStyles)({mainFeaturedPost:{backgroundColor:e.palette.grey[800],color:e.palette.common.white,marginBottom:4*e.spacing.unit},mainFeaturedPostContent:{padding:"".concat(6*e.spacing.unit,"px")},cardGrid:{paddingTop:3*e.spacing.unit,paddingBottom:3*e.spacing.unit},card:{display:"flex"},cardDetails:{flex:1},cardMedia:{width:160}})})(J),H=a(117),K=a.n(H),Q=a(165),U=a(76),V=a.n(U),X=a(166),Y=a.n(X),Z=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,r=new Array(n),c=0;c<n;c++)r[c]=arguments[c];return(a=Object(s.a)(this,(e=Object(m.a)(t)).call.apply(e,[this].concat(r)))).state={paper:{id:"",title:"",content:"",abstract:"",created_at:0,edited_at:0}},a.getPaper=Object(Q.a)(K.a.mark(function e(){return K.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:z.a.get("/api/papers/"+a.props.match.params.id).then(function(e){a.setState({paper:e.data})}).catch(function(e){console.log(e.response)});case 1:case"end":return e.stop()}},e)})),a}return Object(p.a)(t,e),Object(l.a)(t,[{key:"componentDidMount",value:function(){this.getPaper()}},{key:"render",value:function(){var e=this.state.paper;return r.a.createElement(N.a,{item:!0,xs:12,md:12},r.a.createElement(w.a,{variant:"h3",gutterBottom:!0},e.title),r.a.createElement(V.a,null),r.a.createElement(Y.a,{source:e.content}))}}]),t}(r.a.Component);Z.propTypes={classes:g.a.object.isRequired};var ee=Object(f.withStyles)(function(){return Object(f.createStyles)({})})(Z);var te=Object(f.withStyles)(function(e){return Object(f.createStyles)({footer:{backgroundColor:e.palette.background.paper,marginTop:8*e.spacing.unit,padding:"".concat(6*e.spacing.unit,"px 0")}})})(function(e){var t=e.classes;return r.a.createElement("footer",{className:t.footer},r.a.createElement(w.a,{variant:"h6",align:"center",gutterBottom:!0},"Contact me"),r.a.createElement(w.a,{variant:"subtitle1",align:"center",color:"textSecondary",component:"p"},"Email: leo-ryu@outlook.com"))}),ae=a(77),ne=a.n(ae),re=a(39),ce=a.n(re),oe=a(31),ie=a.n(oe),le=a(78),se=a.n(le),me=a(79),pe=a.n(me),ue=a(37),de=a.n(ue),ge=a(49),he=a.n(ge),be=a(167),fe=a.n(be),Ee=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,r=new Array(n),c=0;c<n;c++)r[c]=arguments[c];return(a=Object(s.a)(this,(e=Object(m.a)(t)).call.apply(e,[this].concat(r)))).state={username:"",password:""},a.onSubmit=function(e){e.preventDefault(),z.a.post("http://127.0.0.1:7777/api/verify",a.state).then(function(e){localStorage.setItem(_,e.data.token),console.log("success")}).catch(function(e){localStorage.removeItem("leo-blog-token"),console.log(e.response)})},a.onChange=function(e){a.setState(Object(u.a)({},e.target.name,e.target.value))},a}return Object(p.a)(t,e),Object(l.a)(t,[{key:"render",value:function(){var e=this.props.classes,t=this.state,a=t.username,n=t.password;return r.a.createElement("main",{className:e.main},r.a.createElement(b.a,null),r.a.createElement(k.a,{className:e.paper},r.a.createElement(ne.a,{className:e.avatar},r.a.createElement(fe.a,null)),r.a.createElement(w.a,{component:"h1",variant:"h5"},"Sign in"),r.a.createElement("form",{className:e.form,onSubmit:this.onSubmit},r.a.createElement(ie.a,{margin:"normal",required:!0,fullWidth:!0},r.a.createElement(he.a,{htmlFor:"email"},"Email Address"),r.a.createElement(de.a,{id:"username",name:"username",autoComplete:"email",value:a,onChange:this.onChange,autoFocus:!0})),r.a.createElement(ie.a,{margin:"normal",required:!0,fullWidth:!0},r.a.createElement(he.a,{htmlFor:"password"},"Password"),r.a.createElement(de.a,{name:"password",type:"password",id:"password",autoComplete:"current-password",value:n,onChange:this.onChange})),r.a.createElement(se.a,{control:r.a.createElement(pe.a,{value:"remember",color:"primary"}),label:"Remember me"}),r.a.createElement(ce.a,{type:"submit",fullWidth:!0,variant:"contained",color:"primary",className:e.submit},"Sign in"))))}}]),t}(r.a.Component);Ee.propTypes={classes:g.a.object.isRequired};var ve=Object(f.withStyles)(function(e){return Object(f.createStyles)({main:Object(u.a)({width:"auto",display:"block",marginLeft:3*e.spacing.unit,marginRight:3*e.spacing.unit},e.breakpoints.up(400+3*e.spacing.unit*2),{width:400,marginLeft:"auto",marginRight:"auto"}),paper:{marginTop:8*e.spacing.unit,display:"flex",flexDirection:"column",alignItems:"center",padding:"".concat(2*e.spacing.unit,"px ").concat(3*e.spacing.unit,"px ").concat(3*e.spacing.unit,"px")},avatar:{margin:e.spacing.unit,backgroundColor:e.palette.secondary.main},form:{width:"100%",marginTop:e.spacing.unit},submit:{marginTop:3*e.spacing.unit}})})(Ee),ye=a(50),je=a.n(ye),Oe=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,r=new Array(n),c=0;c<n;c++)r[c]=arguments[c];return(a=Object(s.a)(this,(e=Object(m.a)(t)).call.apply(e,[this].concat(r)))).state=L,a.onSubmit=function(e){e.preventDefault(),a.props.isModification?a.modifyPaper():a.createPaper()},a.createPaper=function(){z.a.post("/api/papers",a.state,{headers:{Authorization:localStorage.getItem(_)}}).catch(function(e){localStorage.removeItem(_),console.log(e.response)})},a.modifyPaper=function(){z.a.put("/api/papers/"+a.props.match.params.id,a.state,{headers:{Authorization:localStorage.getItem(_)}}).catch(function(e){localStorage.removeItem(_),console.log(e.response)})},a.onChange=function(e){a.setState(Object(u.a)({},e.target.name,e.target.value))},a.getPaper=function(){z.a.get("/api/papers/"+a.props.match.params.id).then(function(e){a.setState(e.data)}).catch(function(e){console.log(e.response)})},a.deletePaper=function(){z.a.delete("/api/papers/"+a.props.match.params.id).catch(function(e){console.log(e.response)})},a}return Object(p.a)(t,e),Object(l.a)(t,[{key:"componentDidMount",value:function(){this.props.isModification&&this.getPaper()}},{key:"render",value:function(){var e=this,t=this.props,a=t.classes,n=t.isModification,c=this.state,o=c.title,i=c.abstract,l=c.content;return r.a.createElement(N.a,null,r.a.createElement(ce.a,{type:"submit",variant:"contained",disabled:!n,onClick:function(){return e.deletePaper()}},"Delete"),r.a.createElement(k.a,{className:a.paper},r.a.createElement(w.a,{component:"h1",variant:"h5"},"Paper editor"),r.a.createElement("form",{className:a.form,onSubmit:this.onSubmit},r.a.createElement(ie.a,{margin:"normal",required:!0,fullWidth:!0},r.a.createElement(je.a,{label:"Title",id:"title",name:"title",value:o,onChange:this.onChange,autoFocus:!0})),r.a.createElement(ie.a,{margin:"normal",fullWidth:!0},r.a.createElement(je.a,{label:"Abstract",name:"abstract",id:"abstract",rows:"4",value:i,onChange:this.onChange,multiline:!0})),r.a.createElement(ie.a,{margin:"normal",required:!0,fullWidth:!0},r.a.createElement(je.a,{label:"Content",name:"content",id:"content",value:l,onChange:this.onChange,multiline:!0})),r.a.createElement(ce.a,{type:"submit",variant:"contained",color:"primary",className:a.submit},"Submit"))))}}]),t}(r.a.Component);Oe.propTypes={classes:g.a.object.isRequired};var we=Object(f.withStyles)(function(e){return Object(f.createStyles)({paper:{marginTop:8*e.spacing.unit,display:"flex",flexDirection:"column",alignItems:"center",padding:"".concat(2*e.spacing.unit,"px ").concat(3*e.spacing.unit,"px ").concat(3*e.spacing.unit,"px")},avatar:{margin:e.spacing.unit,backgroundColor:e.palette.secondary.main},form:{width:"100%",marginTop:e.spacing.unit},submit:{marginTop:3*e.spacing.unit}})})(Oe),Se=function(e){function t(){var e,a;Object(i.a)(this,t);for(var n=arguments.length,r=new Array(n),c=0;c<n;c++)r[c]=arguments[c];return(a=Object(s.a)(this,(e=Object(m.a)(t)).call.apply(e,[this].concat(r)))).state={papers:[{id:"",title:"",content:"",abstract:"",created_at:0,edited_at:0}]},a}return Object(p.a)(t,e),Object(l.a)(t,[{key:"render",value:function(){var e=this.props.classes;return r.a.createElement(E.a,null,r.a.createElement(b.a,null),r.a.createElement("div",{className:e.layout},r.a.createElement(S,null),r.a.createElement("main",null,r.a.createElement(v.a,{exact:!0,path:"/",component:$}),r.a.createElement(v.a,{path:"/paper/:id",component:ee}),r.a.createElement(v.a,{path:"/admin/signin",component:ve}),r.a.createElement(v.a,{path:"/admin/create",component:we}),r.a.createElement(v.a,{path:"/admin/edit/:id",render:function(e){return r.a.createElement(we,Object.assign({},e,{isModification:!0}))}}))),r.a.createElement(te,null))}}]),t}(r.a.Component);Se.propTypes={classes:g.a.object.isRequired};var Ce=Object(f.withStyles)(function(e){return Object(f.createStyles)({layout:Object(u.a)({width:"auto",marginLeft:3*e.spacing.unit,marginRight:3*e.spacing.unit},e.breakpoints.up(1100+3*e.spacing.unit*2),{width:1100,marginLeft:"auto",marginRight:"auto"}),mainGrid:{marginTop:3*e.spacing.unit}})})(Se);Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));o.a.render(r.a.createElement(Ce,null),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then(function(e){e.unregister()})}},[[235,1,2]]]);
//# sourceMappingURL=main.2173d8e8.chunk.js.map