package main

var jquery string = `/*! jQuery v3.2.1 | (c) JS Foundation and other contributors | jquery.org/license */
!function(a,b){"use strict";"object"==typeof module&&"object"==typeof module.exports?module.exports=a.document?b(a,!0):function(a){if(!a.document)throw new Error("jQuery requires a window with a document");return b(a)}:b(a)}("undefined"!=typeof window?window:this,function(a,b){"use strict";var c=[],d=a.document,e=Object.getPrototypeOf,f=c.slice,g=c.concat,h=c.push,i=c.indexOf,j={},k=j.toString,l=j.hasOwnProperty,m=l.toString,n=m.call(Object),o={};function p(a,b){b=b||d;var c=b.createElement("script");c.text=a,b.head.appendChild(c).parentNode.removeChild(c)}var q="3.2.1",r=function(a,b){return new r.fn.init(a,b)},s=/^[\s\uFEFF\xA0]+|[\s\uFEFF\xA0]+$/g,t=/^-ms-/,u=/-([a-z])/g,v=function(a,b){return b.toUpperCase()};r.fn=r.prototype={jquery:q,constructor:r,length:0,toArray:function(){return f.call(this)},get:function(a){return null==a?f.call(this):a<0?this[a+this.length]:this[a]},pushStack:function(a){var b=r.merge(this.constructor(),a);return b.prevObject=this,b},each:function(a){return r.each(this,a)},map:function(a){return this.pushStack(r.map(this,function(b,c){return a.call(b,c,b)}))},slice:function(){return this.pushStack(f.apply(this,arguments))},first:function(){return this.eq(0)},last:function(){return this.eq(-1)},eq:function(a){var b=this.length,c=+a+(a<0?b:0);return this.pushStack(c>=0&&c<b?[this[c]]:[])},end:function(){return this.prevObject||this.constructor()},push:h,sort:c.sort,splice:c.splice},r.extend=r.fn.extend=function(){var a,b,c,d,e,f,g=arguments[0]||{},h=1,i=arguments.length,j=!1;for("boolean"==typeof g&&(j=g,g=arguments[h]||{},h++),"object"==typeof g||r.isFunction(g)||(g={}),h===i&&(g=this,h--);h<i;h++)if(null!=(a=arguments[h]))for(b in a)c=g[b],d=a[b],g!==d&&(j&&d&&(r.isPlainObject(d)||(e=Array.isArray(d)))?(e?(e=!1,f=c&&Array.isArray(c)?c:[]):f=c&&r.isPlainObject(c)?c:{},g[b]=r.extend(j,f,d)):void 0!==d&&(g[b]=d));return g},r.extend({expando:"jQuery"+(q+Math.random()).replace(/\D/g,""),isReady:!0,error:function(a){throw new Error(a)},noop:function(){},isFunction:function(a){return"function"===r.type(a)},isWindow:function(a){return null!=a&&a===a.window},isNumeric:function(a){var b=r.type(a);return("number"===b||"string"===b)&&!isNaN(a-parseFloat(a))},isPlainObject:function(a){var b,c;return!(!a||"[object Object]"!==k.call(a))&&(!(b=e(a))||(c=l.call(b,"constructor")&&b.constructor,"function"==typeof c&&m.call(c)===n))},isEmptyObject:function(a){var b;for(b in a)return!1;return!0},type:function(a){return null==a?a+"":"object"==typeof a||"function"==typeof a?j[k.call(a)]||"object":typeof a},globalEval:function(a){p(a)},camelCase:function(a){return a.replace(t,"ms-").replace(u,v)},each:function(a,b){var c,d=0;if(w(a)){for(c=a.length;d<c;d++)if(b.call(a[d],d,a[d])===!1)break}else for(d in a)if(b.call(a[d],d,a[d])===!1)break;return a},trim:function(a){return null==a?"":(a+"").replace(s,"")},makeArray:function(a,b){var c=b||[];return null!=a&&(w(Object(a))?r.merge(c,"string"==typeof a?[a]:a):h.call(c,a)),c},inArray:function(a,b,c){return null==b?-1:i.call(b,a,c)},merge:function(a,b){for(var c=+b.length,d=0,e=a.length;d<c;d++)a[e++]=b[d];return a.length=e,a},grep:function(a,b,c){for(var d,e=[],f=0,g=a.length,h=!c;f<g;f++)d=!b(a[f],f),d!==h&&e.push(a[f]);return e},map:function(a,b,c){var d,e,f=0,h=[];if(w(a))for(d=a.length;f<d;f++)e=b(a[f],f,c),null!=e&&h.push(e);else for(f in a)e=b(a[f],f,c),null!=e&&h.push(e);return g.apply([],h)},guid:1,proxy:function(a,b){var c,d,e;if("string"==typeof b&&(c=a[b],b=a,a=c),r.isFunction(a))return d=f.call(arguments,2),e=function(){return a.apply(b||this,d.concat(f.call(arguments)))},e.guid=a.guid=a.guid||r.guid++,e},now:Date.now,support:o}),"function"==typeof Symbol&&(r.fn[Symbol.iterator]=c[Symbol.iterator]),r.each("Boolean Number String Function Array Date RegExp Object Error Symbol".split(" "),function(a,b){j["[object "+b+"]"]=b.toLowerCase()});function w(a){var b=!!a&&"length"in a&&a.length,c=r.type(a);return"function"!==c&&!r.isWindow(a)&&("array"===c||0===b||"number"==typeof b&&b>0&&b-1 in a)}var x=function(a){var b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u="sizzle"+1*new Date,v=a.document,w=0,x=0,y=ha(),z=ha(),A=ha(),B=function(a,b){return a===b&&(l=!0),0},C={}.hasOwnProperty,D=[],E=D.pop,F=D.push,G=D.push,H=D.slice,I=function(a,b){for(var c=0,d=a.length;c<d;c++)if(a[c]===b)return c;return-1},J="checked|selected|async|autofocus|autoplay|controls|defer|disabled|hidden|ismap|loop|multiple|open|readonly|required|scoped",K="[\\x20\\t\\r\\n\\f]",L="(?:\\\\.|[\\w-]|[^\0-\\xa0])+",M="\\["+K+"*("+L+")(?:"+K+"*([*^$|!~]?=)"+K+"*(?:'((?:\\\\.|[^\\\\'])*)'|\"((?:\\\\.|[^\\\\\"])*)\"|("+L+"))|)"+K+"*\\]",N=":("+L+")(?:\\((('((?:\\\\.|[^\\\\'])*)'|\"((?:\\\\.|[^\\\\\"])*)\")|((?:\\\\.|[^\\\\()[\\]]|"+M+")*)|.*)\\)|)",O=new RegExp(K+"+","g"),P=new RegExp("^"+K+"+|((?:^|[^\\\\])(?:\\\\.)*)"+K+"+$","g"),Q=new RegExp("^"+K+"*,"+K+"*"),R=new RegExp("^"+K+"*([>+~]|"+K+")"+K+"*"),S=new RegExp("="+K+"*([^\\]'\"]*?)"+K+"*\\]","g"),T=new RegExp(N),U=new RegExp("^"+L+"$"),V={ID:new RegExp("^#("+L+")"),CLASS:new RegExp("^\\.("+L+")"),TAG:new RegExp("^("+L+"|[*])"),ATTR:new RegExp("^"+M),PSEUDO:new RegExp("^"+N),CHILD:new RegExp("^:(only|first|last|nth|nth-last)-(child|of-type)(?:\\("+K+"*(even|odd|(([+-]|)(\\d*)n|)"+K+"*(?:([+-]|)"+K+"*(\\d+)|))"+K+"*\\)|)","i"),bool:new RegExp("^(?:"+J+")$","i"),needsContext:new RegExp("^"+K+"*[>+~]|:(even|odd|eq|gt|lt|nth|first|last)(?:\\("+K+"*((?:-\\d)?\\d*)"+K+"*\\)|)(?=[^-]|$)","i")},W=/^(?:input|select|textarea|button)$/i,X=/^h\d$/i,Y=/^[^{]+\{\s*\[native \w/,Z=/^(?:#([\w-]+)|(\w+)|\.([\w-]+))$/,$=/[+~]/,_=new RegExp("\\\\([\\da-f]{1,6}"+K+"?|("+K+")|.)","ig"),aa=function(a,b,c){var d="0x"+b-65536;return d!==d||c?b:d<0?String.fromCharCode(d+65536):String.fromCharCode(d>>10|55296,1023&d|56320)},ba=/([\0-\x1f\x7f]|^-?\d)|^-$|[^\0-\x1f\x7f-\uFFFF\w-]/g,ca=function(a,b){return b?"\0"===a?"\ufffd":a.slice(0,-1)+"\\"+a.charCodeAt(a.length-1).toString(16)+" ":"\\"+a},da=function(){m()},ea=ta(function(a){return a.disabled===!0&&("form"in a||"label"in a)},{dir:"parentNode",next:"legend"});try{G.apply(D=H.call(v.childNodes),v.childNodes),D[v.childNodes.length].nodeType}catch(fa){G={apply:D.length?function(a,b){F.apply(a,H.call(b))}:function(a,b){var c=a.length,d=0;while(a[c++]=b[d++]);a.length=c-1}}}function ga(a,b,d,e){var f,h,j,k,l,o,r,s=b&&b.ownerDocument,w=b?b.nodeType:9;if(d=d||[],"string"!=typeof a||!a||1!==w&&9!==w&&11!==w)return d;if(!e&&((b?b.ownerDocument||b:v)!==n&&m(b),b=b||n,p)){if(11!==w&&(l=Z.exec(a)))if(f=l[1]){if(9===w){if(!(j=b.getElementById(f)))return d;if(j.id===f)return d.push(j),d}else if(s&&(j=s.getElementById(f))&&t(b,j)&&j.id===f)return d.push(j),d}else{if(l[2])return G.apply(d,b.getElementsByTagName(a)),d;if((f=l[3])&&c.getElementsByClassName&&b.getElementsByClassName)return G.apply(d,b.getElementsByClassName(f)),d}if(c.qsa&&!A[a+" "]&&(!q||!q.test(a))){if(1!==w)s=b,r=a;else if("object"!==b.nodeName.toLowerCase()){(k=b.getAttribute("id"))?k=k.replace(ba,ca):b.setAttribute("id",k=u),o=g(a),h=o.length;while(h--)o[h]="#"+k+" "+sa(o[h]);r=o.join(","),s=$.test(a)&&qa(b.parentNode)||b}if(r)try{return G.apply(d,s.querySelectorAll(r)),d}catch(x){}finally{k===u&&b.removeAttribute("id")}}}return i(a.replace(P,"$1"),b,d,e)}function ha(){var a=[];function b(c,e){return a.push(c+" ")>d.cacheLength&&delete b[a.shift()],b[c+" "]=e}return b}function ia(a){return a[u]=!0,a}function ja(a){var b=n.createElement("fieldset");try{return!!a(b)}catch(c){return!1}finally{b.parentNode&&b.parentNode.removeChild(b),b=null}}function ka(a,b){var c=a.split("|"),e=c.length;while(e--)d.attrHandle[c[e]]=b}function la(a,b){var c=b&&a,d=c&&1===a.nodeType&&1===b.nodeType&&a.sourceIndex-b.sourceIndex;if(d)return d;if(c)while(c=c.nextSibling)if(c===b)return-1;return a?1:-1}function ma(a){return function(b){var c=b.nodeName.toLowerCase();return"input"===c&&b.type===a}}function na(a){return function(b){var c=b.nodeName.toLowerCase();return("input"===c||"button"===c)&&b.type===a}}function oa(a){return function(b){return"form"in b?b.parentNode&&b.disabled===!1?"label"in b?"label"in b.parentNode?b.parentNode.disabled===a:b.disabled===a:b.isDisabled===a||b.isDisabled!==!a&&ea(b)===a:b.disabled===a:"label"in b&&b.disabled===a}}function pa(a){return ia(function(b){return b=+b,ia(function(c,d){var e,f=a([],c.length,b),g=f.length;while(g--)c[e=f[g]]&&(c[e]=!(d[e]=c[e]))})})}function qa(a){return a&&"undefined"!=typeof a.getElementsByTagName&&a}c=ga.support={},f=ga.isXML=function(a){var b=a&&(a.ownerDocument||a).documentElement;return!!b&&"HTML"!==b.nodeName},m=ga.setDocument=function(a){var b,e,g=a?a.ownerDocument||a:v;return g!==n&&9===g.nodeType&&g.documentElement?(n=g,o=n.documentElement,p=!f(n),v!==n&&(e=n.defaultView)&&e.top!==e&&(e.addEventListener?e.addEventListener("unload",da,!1):e.attachEvent&&e.attachEvent("onunload",da)),c.attributes=ja(function(a){return a.className="i",!a.getAttribute("className")}),c.getElementsByTagName=ja(function(a){return a.appendChild(n.createComment("")),!a.getElementsByTagName("*").length}),c.getElementsByClassName=Y.test(n.getElementsByClassName),c.getById=ja(function(a){return o.appendChild(a).id=u,!n.getElementsByName||!n.getElementsByName(u).length}),c.getById?(d.filter.ID=function(a){var b=a.replace(_,aa);return function(a){return a.getAttribute("id")===b}},d.find.ID=function(a,b){if("undefined"!=typeof b.getElementById&&p){var c=b.getElementById(a);return c?[c]:[]}}):(d.filter.ID=function(a){var b=a.replace(_,aa);return function(a){var c="undefined"!=typeof a.getAttributeNode&&a.getAttributeNode("id");return c&&c.value===b}},d.find.ID=function(a,b){if("undefined"!=typeof b.getElementById&&p){var c,d,e,f=b.getElementById(a);if(f){if(c=f.getAttributeNode("id"),c&&c.value===a)return[f];e=b.getElementsByName(a),d=0;while(f=e[d++])if(c=f.getAttributeNode("id"),c&&c.value===a)return[f]}return[]}}),d.find.TAG=c.getElementsByTagName?function(a,b){return"undefined"!=typeof b.getElementsByTagName?b.getElementsByTagName(a):c.qsa?b.querySelectorAll(a):void 0}:function(a,b){var c,d=[],e=0,f=b.getElementsByTagName(a);if("*"===a){while(c=f[e++])1===c.nodeType&&d.push(c);return d}return f},d.find.CLASS=c.getElementsByClassName&&function(a,b){if("undefined"!=typeof b.getElementsByClassName&&p)return b.getElementsByClassName(a)},r=[],q=[],(c.qsa=Y.test(n.querySelectorAll))&&(ja(function(a){o.appendChild(a).innerHTML="<a id='"+u+"'></a><select id='"+u+"-\r\\' msallowcapture=''><option selected=''></option></select>",a.querySelectorAll("[msallowcapture^='']").length&&q.push("[*^$]="+K+"*(?:''|\"\")"),a.querySelectorAll("[selected]").length||q.push("\\["+K+"*(?:value|"+J+")"),a.querySelectorAll("[id~="+u+"-]").length||q.push("~="),a.querySelectorAll(":checked").length||q.push(":checked"),a.querySelectorAll("a#"+u+"+*").length||q.push(".#.+[+~]")}),ja(function(a){a.innerHTML="<a href='' disabled='disabled'></a><select disabled='disabled'><option/></select>";var b=n.createElement("input");b.setAttribute("type","hidden"),a.appendChild(b).setAttribute("name","D"),a.querySelectorAll("[name=d]").length&&q.push("name"+K+"*[*^$|!~]?="),2!==a.querySelectorAll(":enabled").length&&q.push(":enabled",":disabled"),o.appendChild(a).disabled=!0,2!==a.querySelectorAll(":disabled").length&&q.push(":enabled",":disabled"),a.querySelectorAll("*,:x"),q.push(",.*:")})),(c.matchesSelector=Y.test(s=o.matches||o.webkitMatchesSelector||o.mozMatchesSelector||o.oMatchesSelector||o.msMatchesSelector))&&ja(function(a){c.disconnectedMatch=s.call(a,"*"),s.call(a,"[s!='']:x"),r.push("!=",N)}),q=q.length&&new RegExp(q.join("|")),r=r.length&&new RegExp(r.join("|")),b=Y.test(o.compareDocumentPosition),t=b||Y.test(o.contains)?function(a,b){var c=9===a.nodeType?a.documentElement:a,d=b&&b.parentNode;return a===d||!(!d||1!==d.nodeType||!(c.contains?c.contains(d):a.compareDocumentPosition&&16&a.compareDocumentPosition(d)))}:function(a,b){if(b)while(b=b.parentNode)if(b===a)return!0;return!1},B=b?function(a,b){if(a===b)return l=!0,0;var d=!a.compareDocumentPosition-!b.compareDocumentPosition;return d?d:(d=(a.ownerDocument||a)===(b.ownerDocument||b)?a.compareDocumentPosition(b):1,1&d||!c.sortDetached&&b.compareDocumentPosition(a)===d?a===n||a.ownerDocument===v&&t(v,a)?-1:b===n||b.ownerDocument===v&&t(v,b)?1:k?I(k,a)-I(k,b):0:4&d?-1:1)}:function(a,b){if(a===b)return l=!0,0;var c,d=0,e=a.parentNode,f=b.parentNode,g=[a],h=[b];if(!e||!f)return a===n?-1:b===n?1:e?-1:f?1:k?I(k,a)-I(k,b):0;if(e===f)return la(a,b);c=a;while(c=c.parentNode)g.unshift(c);c=b;while(c=c.parentNode)h.unshift(c);while(g[d]===h[d])d++;return d?la(g[d],h[d]):g[d]===v?-1:h[d]===v?1:0},n):n},ga.matches=function(a,b){return ga(a,null,null,b)},ga.matchesSelector=function(a,b){if((a.ownerDocument||a)!==n&&m(a),b=b.replace(S,"='$1']"),c.matchesSelector&&p&&!A[b+" "]&&(!r||!r.test(b))&&(!q||!q.test(b)))try{var d=s.call(a,b);if(d||c.disconnectedMatch||a.document&&11!==a.document.nodeType)return d}catch(e){}return ga(b,n,null,[a]).length>0},ga.contains=function(a,b){return(a.ownerDocument||a)!==n&&m(a),t(a,b)},ga.attr=function(a,b){(a.ownerDocument||a)!==n&&m(a);var e=d.attrHandle[b.toLowerCase()],f=e&&C.call(d.attrHandle,b.toLowerCase())?e(a,b,!p):void 0;return void 0!==f?f:c.attributes||!p?a.getAttribute(b):(f=a.getAttributeNode(b))&&f.specified?f.value:null},ga.escape=function(a){return(a+"").replace(ba,ca)},ga.error=function(a){throw new Error("Syntax error, unrecognized expression: "+a)},ga.uniqueSort=function(a){var b,d=[],e=0,f=0;if(l=!c.detectDuplicates,k=!c.sortStable&&a.slice(0),a.sort(B),l){while(b=a[f++])b===a[f]&&(e=d.push(f));while(e--)a.splice(d[e],1)}return k=null,a},e=ga.getText=function(a){var b,c="",d=0,f=a.nodeType;if(f){if(1===f||9===f||11===f){if("string"==typeof a.textContent)return a.textContent;for(a=a.firstChild;a;a=a.nextSibling)c+=e(a)}else if(3===f||4===f)return a.nodeValue}else while(b=a[d++])c+=e(b);return c},d=ga.selectors={cacheLength:50,createPseudo:ia,match:V,attrHandle:{},find:{},relative:{">":{dir:"parentNode",first:!0}," ":{dir:"parentNode"},"+":{dir:"previousSibling",first:!0},"~":{dir:"previousSibling"}},preFilter:{ATTR:function(a){return a[1]=a[1].replace(_,aa),a[3]=(a[3]||a[4]||a[5]||"").replace(_,aa),"~="===a[2]&&(a[3]=" "+a[3]+" "),a.slice(0,4)},CHILD:function(a){return a[1]=a[1].toLowerCase(),"nth"===a[1].slice(0,3)?(a[3]||ga.error(a[0]),a[4]=+(a[4]?a[5]+(a[6]||1):2*("even"===a[3]||"odd"===a[3])),a[5]=+(a[7]+a[8]||"odd"===a[3])):a[3]&&ga.error(a[0]),a},PSEUDO:function(a){var b,c=!a[6]&&a[2];return V.CHILD.test(a[0])?null:(a[3]?a[2]=a[4]||a[5]||"":c&&T.test(c)&&(b=g(c,!0))&&(b=c.indexOf(")",c.length-b)-c.length)&&(a[0]=a[0].slice(0,b),a[2]=c.slice(0,b)),a.slice(0,3))}},filter:{TAG:function(a){var b=a.replace(_,aa).toLowerCase();return"*"===a?function(){return!0}:function(a){return a.nodeName&&a.nodeName.toLowerCase()===b}},CLASS:function(a){var b=y[a+" "];return b||(b=new RegExp("(^|"+K+")"+a+"("+K+"|$)"))&&y(a,function(a){return b.test("string"==typeof a.className&&a.className||"undefined"!=typeof a.getAttribute&&a.getAttribute("class")||"")})},ATTR:function(a,b,c){return function(d){var e=ga.attr(d,a);return null==e?"!="===b:!b||(e+="","="===b?e===c:"!="===b?e!==c:"^="===b?c&&0===e.indexOf(c):"*="===b?c&&e.indexOf(c)>-1:"$="===b?c&&e.slice(-c.length)===c:"~="===b?(" "+e.replace(O," ")+" ").indexOf(c)>-1:"|="===b&&(e===c||e.slice(0,c.length+1)===c+"-"))}},CHILD:function(a,b,c,d,e){var f="nth"!==a.slice(0,3),g="last"!==a.slice(-4),h="of-type"===b;return 1===d&&0===e?function(a){return!!a.parentNode}:function(b,c,i){var j,k,l,m,n,o,p=f!==g?"nextSibling":"previousSibling",q=b.parentNode,r=h&&b.nodeName.toLowerCase(),s=!i&&!h,t=!1;if(q){if(f){while(p){m=b;while(m=m[p])if(h?m.nodeName.toLowerCase()===r:1===m.nodeType)return!1;o=p="only"===a&&!o&&"nextSibling"}return!0}if(o=[g?q.firstChild:q.lastChild],g&&s){m=q,l=m[u]||(m[u]={}),k=l[m.uniqueID]||(l[m.uniqueID]={}),j=k[a]||[],n=j[0]===w&&j[1],t=n&&j[2],m=n&&q.childNodes[n];while(m=++n&&m&&m[p]||(t=n=0)||o.pop())if(1===m.nodeType&&++t&&m===b){k[a]=[w,n,t];break}}else if(s&&(m=b,l=m[u]||(m[u]={}),k=l[m.uniqueID]||(l[m.uniqueID]={}),j=k[a]||[],n=j[0]===w&&j[1],t=n),t===!1)while(m=++n&&m&&m[p]||(t=n=0)||o.pop())if((h?m.nodeName.toLowerCase()===r:1===m.nodeType)&&++t&&(s&&(l=m[u]||(m[u]={}),k=l[m.uniqueID]||(l[m.uniqueID]={}),k[a]=[w,t]),m===b))break;return t-=e,t===d||t%d===0&&t/d>=0}}},PSEUDO:function(a,b){var c,e=d.pseudos[a]||d.setFilters[a.toLowerCase()]||ga.error("unsupported pseudo: "+a);return e[u]?e(b):e.length>1?(c=[a,a,"",b],d.setFilters.hasOwnProperty(a.toLowerCase())?ia(function(a,c){var d,f=e(a,b),g=f.length;while(g--)d=I(a,f[g]),a[d]=!(c[d]=f[g])}):function(a){return e(a,0,c)}):e}},pseudos:{not:ia(function(a){var b=[],c=[],d=h(a.replace(P,"$1"));return d[u]?ia(function(a,b,c,e){var f,g=d(a,null,e,[]),h=a.length;while(h--)(f=g[h])&&(a[h]=!(b[h]=f))}):function(a,e,f){return b[0]=a,d(b,null,f,c),b[0]=null,!c.pop()}}),has:ia(function(a){return function(b){return ga(a,b).length>0}}),contains:ia(function(a){return a=a.replace(_,aa),function(b){return(b.textContent||b.innerText||e(b)).indexOf(a)>-1}}),lang:ia(function(a){return U.test(a||"")||ga.error("unsupported lang: "+a),a=a.replace(_,aa).toLowerCase(),function(b){var c;do if(c=p?b.lang:b.getAttribute("xml:lang")||b.getAttribute("lang"))return c=c.toLowerCase(),c===a||0===c.indexOf(a+"-");while((b=b.parentNode)&&1===b.nodeType);return!1}}),target:function(b){var c=a.location&&a.location.hash;return c&&c.slice(1)===b.id},root:function(a){return a===o},focus:function(a){return a===n.activeElement&&(!n.hasFocus||n.hasFocus())&&!!(a.type||a.href||~a.tabIndex)},enabled:oa(!1),disabled:oa(!0),checked:function(a){var b=a.nodeName.toLowerCase();return"input"===b&&!!a.checked||"option"===b&&!!a.selected},selected:function(a){return a.parentNode&&a.parentNode.selectedIndex,a.selected===!0},empty:function(a){for(a=a.firstChild;a;a=a.nextSibling)if(a.nodeType<6)return!1;return!0},parent:function(a){return!d.pseudos.empty(a)},header:function(a){return X.test(a.nodeName)},input:function(a){return W.test(a.nodeName)},button:function(a){var b=a.nodeName.toLowerCase();return"input"===b&&"button"===a.type||"button"===b},text:function(a){var b;return"input"===a.nodeName.toLowerCase()&&"text"===a.type&&(null==(b=a.getAttribute("type"))||"text"===b.toLowerCase())},first:pa(function(){return[0]}),last:pa(function(a,b){return[b-1]}),eq:pa(function(a,b,c){return[c<0?c+b:c]}),even:pa(function(a,b){for(var c=0;c<b;c+=2)a.push(c);return a}),odd:pa(function(a,b){for(var c=1;c<b;c+=2)a.push(c);return a}),lt:pa(function(a,b,c){for(var d=c<0?c+b:c;--d>=0;)a.push(d);return a}),gt:pa(function(a,b,c){for(var d=c<0?c+b:c;++d<b;)a.push(d);return a})}},d.pseudos.nth=d.pseudos.eq;for(b in{radio:!0,checkbox:!0,file:!0,password:!0,image:!0})d.pseudos[b]=ma(b);for(b in{submit:!0,reset:!0})d.pseudos[b]=na(b);function ra(){}ra.prototype=d.filters=d.pseudos,d.setFilters=new ra,g=ga.tokenize=function(a,b){var c,e,f,g,h,i,j,k=z[a+" "];if(k)return b?0:k.slice(0);h=a,i=[],j=d.preFilter;while(h){c&&!(e=Q.exec(h))||(e&&(h=h.slice(e[0].length)||h),i.push(f=[])),c=!1,(e=R.exec(h))&&(c=e.shift(),f.push({value:c,type:e[0].replace(P," ")}),h=h.slice(c.length));for(g in d.filter)!(e=V[g].exec(h))||j[g]&&!(e=j[g](e))||(c=e.shift(),f.push({value:c,type:g,matches:e}),h=h.slice(c.length));if(!c)break}return b?h.length:h?ga.error(a):z(a,i).slice(0)};function sa(a){for(var b=0,c=a.length,d="";b<c;b++)d+=a[b].value;return d}function ta(a,b,c){var d=b.dir,e=b.next,f=e||d,g=c&&"parentNode"===f,h=x++;return b.first?function(b,c,e){while(b=b[d])if(1===b.nodeType||g)return a(b,c,e);return!1}:function(b,c,i){var j,k,l,m=[w,h];if(i){while(b=b[d])if((1===b.nodeType||g)&&a(b,c,i))return!0}else while(b=b[d])if(1===b.nodeType||g)if(l=b[u]||(b[u]={}),k=l[b.uniqueID]||(l[b.uniqueID]={}),e&&e===b.nodeName.toLowerCase())b=b[d]||b;else{if((j=k[f])&&j[0]===w&&j[1]===h)return m[2]=j[2];if(k[f]=m,m[2]=a(b,c,i))return!0}return!1}}function ua(a){return a.length>1?function(b,c,d){var e=a.length;while(e--)if(!a[e](b,c,d))return!1;return!0}:a[0]}function va(a,b,c){for(var d=0,e=b.length;d<e;d++)ga(a,b[d],c);return c}function wa(a,b,c,d,e){for(var f,g=[],h=0,i=a.length,j=null!=b;h<i;h++)(f=a[h])&&(c&&!c(f,d,e)||(g.push(f),j&&b.push(h)));return g}function xa(a,b,c,d,e,f){return d&&!d[u]&&(d=xa(d)),e&&!e[u]&&(e=xa(e,f)),ia(function(f,g,h,i){var j,k,l,m=[],n=[],o=g.length,p=f||va(b||"*",h.nodeType?[h]:h,[]),q=!a||!f&&b?p:wa(p,m,a,h,i),r=c?e||(f?a:o||d)?[]:g:q;if(c&&c(q,r,h,i),d){j=wa(r,n),d(j,[],h,i),k=j.length;while(k--)(l=j[k])&&(r[n[k]]=!(q[n[k]]=l))}if(f){if(e||a){if(e){j=[],k=r.length;while(k--)(l=r[k])&&j.push(q[k]=l);e(null,r=[],j,i)}k=r.length;while(k--)(l=r[k])&&(j=e?I(f,l):m[k])>-1&&(f[j]=!(g[j]=l))}}else r=wa(r===g?r.splice(o,r.length):r),e?e(null,g,r,i):G.apply(g,r)})}function ya(a){for(var b,c,e,f=a.length,g=d.relative[a[0].type],h=g||d.relative[" "],i=g?1:0,k=ta(function(a){return a===b},h,!0),l=ta(function(a){return I(b,a)>-1},h,!0),m=[function(a,c,d){var e=!g&&(d||c!==j)||((b=c).nodeType?k(a,c,d):l(a,c,d));return b=null,e}];i<f;i++)if(c=d.relative[a[i].type])m=[ta(ua(m),c)];else{if(c=d.filter[a[i].type].apply(null,a[i].matches),c[u]){for(e=++i;e<f;e++)if(d.relative[a[e].type])break;return xa(i>1&&ua(m),i>1&&sa(a.slice(0,i-1).concat({value:" "===a[i-2].type?"*":""})).replace(P,"$1"),c,i<e&&ya(a.slice(i,e)),e<f&&ya(a=a.slice(e)),e<f&&sa(a))}m.push(c)}return ua(m)}function za(a,b){var c=b.length>0,e=a.length>0,f=function(f,g,h,i,k){var l,o,q,r=0,s="0",t=f&&[],u=[],v=j,x=f||e&&d.find.TAG("*",k),y=w+=null==v?1:Math.random()||.1,z=x.length;for(k&&(j=g===n||g||k);s!==z&&null!=(l=x[s]);s++){if(e&&l){o=0,g||l.ownerDocument===n||(m(l),h=!p);while(q=a[o++])if(q(l,g||n,h)){i.push(l);break}k&&(w=y)}c&&((l=!q&&l)&&r--,f&&t.push(l))}if(r+=s,c&&s!==r){o=0;while(q=b[o++])q(t,u,g,h);if(f){if(r>0)while(s--)t[s]||u[s]||(u[s]=E.call(i));u=wa(u)}G.apply(i,u),k&&!f&&u.length>0&&r+b.length>1&&ga.uniqueSort(i)}return k&&(w=y,j=v),t};return c?ia(f):f}return h=ga.compile=function(a,b){var c,d=[],e=[],f=A[a+" "];if(!f){b||(b=g(a)),c=b.length;while(c--)f=ya(b[c]),f[u]?d.push(f):e.push(f);f=A(a,za(e,d)),f.selector=a}return f},i=ga.select=function(a,b,c,e){var f,i,j,k,l,m="function"==typeof a&&a,n=!e&&g(a=m.selector||a);if(c=c||[],1===n.length){if(i=n[0]=n[0].slice(0),i.length>2&&"ID"===(j=i[0]).type&&9===b.nodeType&&p&&d.relative[i[1].type]){if(b=(d.find.ID(j.matches[0].replace(_,aa),b)||[])[0],!b)return c;m&&(b=b.parentNode),a=a.slice(i.shift().value.length)}f=V.needsContext.test(a)?0:i.length;while(f--){if(j=i[f],d.relative[k=j.type])break;if((l=d.find[k])&&(e=l(j.matches[0].replace(_,aa),$.test(i[0].type)&&qa(b.parentNode)||b))){if(i.splice(f,1),a=e.length&&sa(i),!a)return G.apply(c,e),c;break}}}return(m||h(a,n))(e,b,!p,c,!b||$.test(a)&&qa(b.parentNode)||b),c},c.sortStable=u.split("").sort(B).join("")===u,c.detectDuplicates=!!l,m(),c.sortDetached=ja(function(a){return 1&a.compareDocumentPosition(n.createElement("fieldset"))}),ja(function(a){return a.innerHTML="<a href='#'></a>","#"===a.firstChild.getAttribute("href")})||ka("type|href|height|width",function(a,b,c){if(!c)return a.getAttribute(b,"type"===b.toLowerCase()?1:2)}),c.attributes&&ja(function(a){return a.innerHTML="<input/>",a.firstChild.setAttribute("value",""),""===a.firstChild.getAttribute("value")})||ka("value",function(a,b,c){if(!c&&"input"===a.nodeName.toLowerCase())return a.defaultValue}),ja(function(a){return null==a.getAttribute("disabled")})||ka(J,function(a,b,c){var d;if(!c)return a[b]===!0?b.toLowerCase():(d=a.getAttributeNode(b))&&d.specified?d.value:null}),ga}(a);r.find=x,r.expr=x.selectors,r.expr[":"]=r.expr.pseudos,r.uniqueSort=r.unique=x.uniqueSort,r.text=x.getText,r.isXMLDoc=x.isXML,r.contains=x.contains,r.escapeSelector=x.escape;var y=function(a,b,c){var d=[],e=void 0!==c;while((a=a[b])&&9!==a.nodeType)if(1===a.nodeType){if(e&&r(a).is(c))break;d.push(a)}return d},z=function(a,b){for(var c=[];a;a=a.nextSibling)1===a.nodeType&&a!==b&&c.push(a);return c},A=r.expr.match.needsContext;function B(a,b){return a.nodeName&&a.nodeName.toLowerCase()===b.toLowerCase()}var C=/^<([a-z][^\/\0>:\x20\t\r\n\f]*)[\x20\t\r\n\f]*\/?>(?:<\/\1>|)$/i,D=/^.[^:#\[\.,]*$/;function E(a,b,c){return r.isFunction(b)?r.grep(a,function(a,d){return!!b.call(a,d,a)!==c}):b.nodeType?r.grep(a,function(a){return a===b!==c}):"string"!=typeof b?r.grep(a,function(a){return i.call(b,a)>-1!==c}):D.test(b)?r.filter(b,a,c):(b=r.filter(b,a),r.grep(a,function(a){return i.call(b,a)>-1!==c&&1===a.nodeType}))}r.filter=function(a,b,c){var d=b[0];return c&&(a=":not("+a+")"),1===b.length&&1===d.nodeType?r.find.matchesSelector(d,a)?[d]:[]:r.find.matches(a,r.grep(b,function(a){return 1===a.nodeType}))},r.fn.extend({find:function(a){var b,c,d=this.length,e=this;if("string"!=typeof a)return this.pushStack(r(a).filter(function(){for(b=0;b<d;b++)if(r.contains(e[b],this))return!0}));for(c=this.pushStack([]),b=0;b<d;b++)r.find(a,e[b],c);return d>1?r.uniqueSort(c):c},filter:function(a){return this.pushStack(E(this,a||[],!1))},not:function(a){return this.pushStack(E(this,a||[],!0))},is:function(a){return!!E(this,"string"==typeof a&&A.test(a)?r(a):a||[],!1).length}});var F,G=/^(?:\s*(<[\w\W]+>)[^>]*|#([\w-]+))$/,H=r.fn.init=function(a,b,c){var e,f;if(!a)return this;if(c=c||F,"string"==typeof a){if(e="<"===a[0]&&">"===a[a.length-1]&&a.length>=3?[null,a,null]:G.exec(a),!e||!e[1]&&b)return!b||b.jquery?(b||c).find(a):this.constructor(b).find(a);if(e[1]){if(b=b instanceof r?b[0]:b,r.merge(this,r.parseHTML(e[1],b&&b.nodeType?b.ownerDocument||b:d,!0)),C.test(e[1])&&r.isPlainObject(b))for(e in b)r.isFunction(this[e])?this[e](b[e]):this.attr(e,b[e]);return this}return f=d.getElementById(e[2]),f&&(this[0]=f,this.length=1),this}return a.nodeType?(this[0]=a,this.length=1,this):r.isFunction(a)?void 0!==c.ready?c.ready(a):a(r):r.makeArray(a,this)};H.prototype=r.fn,F=r(d);var I=/^(?:parents|prev(?:Until|All))/,J={children:!0,contents:!0,next:!0,prev:!0};r.fn.extend({has:function(a){var b=r(a,this),c=b.length;return this.filter(function(){for(var a=0;a<c;a++)if(r.contains(this,b[a]))return!0})},closest:function(a,b){var c,d=0,e=this.length,f=[],g="string"!=typeof a&&r(a);if(!A.test(a))for(;d<e;d++)for(c=this[d];c&&c!==b;c=c.parentNode)if(c.nodeType<11&&(g?g.index(c)>-1:1===c.nodeType&&r.find.matchesSelector(c,a))){f.push(c);break}return this.pushStack(f.length>1?r.uniqueSort(f):f)},index:function(a){return a?"string"==typeof a?i.call(r(a),this[0]):i.call(this,a.jquery?a[0]:a):this[0]&&this[0].parentNode?this.first().prevAll().length:-1},add:function(a,b){return this.pushStack(r.uniqueSort(r.merge(this.get(),r(a,b))))},addBack:function(a){return this.add(null==a?this.prevObject:this.prevObject.filter(a))}});function K(a,b){while((a=a[b])&&1!==a.nodeType);return a}r.each({parent:function(a){var b=a.parentNode;return b&&11!==b.nodeType?b:null},parents:function(a){return y(a,"parentNode")},parentsUntil:function(a,b,c){return y(a,"parentNode",c)},next:function(a){return K(a,"nextSibling")},prev:function(a){return K(a,"previousSibling")},nextAll:function(a){return y(a,"nextSibling")},prevAll:function(a){return y(a,"previousSibling")},nextUntil:function(a,b,c){return y(a,"nextSibling",c)},prevUntil:function(a,b,c){return y(a,"previousSibling",c)},siblings:function(a){return z((a.parentNode||{}).firstChild,a)},children:function(a){return z(a.firstChild)},contents:function(a){return B(a,"iframe")?a.contentDocument:(B(a,"template")&&(a=a.content||a),r.merge([],a.childNodes))}},function(a,b){r.fn[a]=function(c,d){var e=r.map(this,b,c);return"Until"!==a.slice(-5)&&(d=c),d&&"string"==typeof d&&(e=r.filter(d,e)),this.length>1&&(J[a]||r.uniqueSort(e),I.test(a)&&e.reverse()),this.pushStack(e)}});var L=/[^\x20\t\r\n\f]+/g;function M(a){var b={};return r.each(a.match(L)||[],function(a,c){b[c]=!0}),b}r.Callbacks=function(a){a="string"==typeof a?M(a):r.extend({},a);var b,c,d,e,f=[],g=[],h=-1,i=function(){for(e=e||a.once,d=b=!0;g.length;h=-1){c=g.shift();while(++h<f.length)f[h].apply(c[0],c[1])===!1&&a.stopOnFalse&&(h=f.length,c=!1)}a.memory||(c=!1),b=!1,e&&(f=c?[]:"")},j={add:function(){return f&&(c&&!b&&(h=f.length-1,g.push(c)),function d(b){r.each(b,function(b,c){r.isFunction(c)?a.unique&&j.has(c)||f.push(c):c&&c.length&&"string"!==r.type(c)&&d(c)})}(arguments),c&&!b&&i()),this},remove:function(){return r.each(arguments,function(a,b){var c;while((c=r.inArray(b,f,c))>-1)f.splice(c,1),c<=h&&h--}),this},has:function(a){return a?r.inArray(a,f)>-1:f.length>0},empty:function(){return f&&(f=[]),this},disable:function(){return e=g=[],f=c="",this},disabled:function(){return!f},lock:function(){return e=g=[],c||b||(f=c=""),this},locked:function(){return!!e},fireWith:function(a,c){return e||(c=c||[],c=[a,c.slice?c.slice():c],g.push(c),b||i()),this},fire:function(){return j.fireWith(this,arguments),this},fired:function(){return!!d}};return j};function N(a){return a}function O(a){throw a}function P(a,b,c,d){var e;try{a&&r.isFunction(e=a.promise)?e.call(a).done(b).fail(c):a&&r.isFunction(e=a.then)?e.call(a,b,c):b.apply(void 0,[a].slice(d))}catch(a){c.apply(void 0,[a])}}r.extend({Deferred:function(b){var c=[["notify","progress",r.Callbacks("memory"),r.Callbacks("memory"),2],["resolve","done",r.Callbacks("once memory"),r.Callbacks("once memory"),0,"resolved"],["reject","fail",r.Callbacks("once memory"),r.Callbacks("once memory"),1,"rejected"]],d="pending",e={state:function(){return d},always:function(){return f.done(arguments).fail(arguments),this},"catch":function(a){return e.then(null,a)},pipe:function(){var a=arguments;return r.Deferred(function(b){r.each(c,function(c,d){var e=r.isFunction(a[d[4]])&&a[d[4]];f[d[1]](function(){var a=e&&e.apply(this,arguments);a&&r.isFunction(a.promise)?a.promise().progress(b.notify).done(b.resolve).fail(b.reject):b[d[0]+"With"](this,e?[a]:arguments)})}),a=null}).promise()},then:function(b,d,e){var f=0;function g(b,c,d,e){return function(){var h=this,i=arguments,j=function(){var a,j;if(!(b<f)){if(a=d.apply(h,i),a===c.promise())throw new TypeError("Thenable self-resolution");j=a&&("object"==typeof a||"function"==typeof a)&&a.then,r.isFunction(j)?e?j.call(a,g(f,c,N,e),g(f,c,O,e)):(f++,j.call(a,g(f,c,N,e),g(f,c,O,e),g(f,c,N,c.notifyWith))):(d!==N&&(h=void 0,i=[a]),(e||c.resolveWith)(h,i))}},k=e?j:function(){try{j()}catch(a){r.Deferred.exceptionHook&&r.Deferred.exceptionHook(a,k.stackTrace),b+1>=f&&(d!==O&&(h=void 0,i=[a]),c.rejectWith(h,i))}};b?k():(r.Deferred.getStackHook&&(k.stackTrace=r.Deferred.getStackHook()),a.setTimeout(k))}}return r.Deferred(function(a){c[0][3].add(g(0,a,r.isFunction(e)?e:N,a.notifyWith)),c[1][3].add(g(0,a,r.isFunction(b)?b:N)),c[2][3].add(g(0,a,r.isFunction(d)?d:O))}).promise()},promise:function(a){return null!=a?r.extend(a,e):e}},f={};return r.each(c,function(a,b){var g=b[2],h=b[5];e[b[1]]=g.add,h&&g.add(function(){d=h},c[3-a][2].disable,c[0][2].lock),g.add(b[3].fire),f[b[0]]=function(){return f[b[0]+"With"](this===f?void 0:this,arguments),this},f[b[0]+"With"]=g.fireWith}),e.promise(f),b&&b.call(f,f),f},when:function(a){var b=arguments.length,c=b,d=Array(c),e=f.call(arguments),g=r.Deferred(),h=function(a){return function(c){d[a]=this,e[a]=arguments.length>1?f.call(arguments):c,--b||g.resolveWith(d,e)}};if(b<=1&&(P(a,g.done(h(c)).resolve,g.reject,!b),"pending"===g.state()||r.isFunction(e[c]&&e[c].then)))return g.then();while(c--)P(e[c],h(c),g.reject);return g.promise()}});var Q=/^(Eval|Internal|Range|Reference|Syntax|Type|URI)Error$/;r.Deferred.exceptionHook=function(b,c){a.console&&a.console.warn&&b&&Q.test(b.name)&&a.console.warn("jQuery.Deferred exception: "+b.message,b.stack,c)},r.readyException=function(b){a.setTimeout(function(){throw b})};var R=r.Deferred();r.fn.ready=function(a){return R.then(a)["catch"](function(a){r.readyException(a)}),this},r.extend({isReady:!1,readyWait:1,ready:function(a){(a===!0?--r.readyWait:r.isReady)||(r.isReady=!0,a!==!0&&--r.readyWait>0||R.resolveWith(d,[r]))}}),r.ready.then=R.then;function S(){d.removeEventListener("DOMContentLoaded",S),
a.removeEventListener("load",S),r.ready()}"complete"===d.readyState||"loading"!==d.readyState&&!d.documentElement.doScroll?a.setTimeout(r.ready):(d.addEventListener("DOMContentLoaded",S),a.addEventListener("load",S));var T=function(a,b,c,d,e,f,g){var h=0,i=a.length,j=null==c;if("object"===r.type(c)){e=!0;for(h in c)T(a,b,h,c[h],!0,f,g)}else if(void 0!==d&&(e=!0,r.isFunction(d)||(g=!0),j&&(g?(b.call(a,d),b=null):(j=b,b=function(a,b,c){return j.call(r(a),c)})),b))for(;h<i;h++)b(a[h],c,g?d:d.call(a[h],h,b(a[h],c)));return e?a:j?b.call(a):i?b(a[0],c):f},U=function(a){return 1===a.nodeType||9===a.nodeType||!+a.nodeType};function V(){this.expando=r.expando+V.uid++}V.uid=1,V.prototype={cache:function(a){var b=a[this.expando];return b||(b={},U(a)&&(a.nodeType?a[this.expando]=b:Object.defineProperty(a,this.expando,{value:b,configurable:!0}))),b},set:function(a,b,c){var d,e=this.cache(a);if("string"==typeof b)e[r.camelCase(b)]=c;else for(d in b)e[r.camelCase(d)]=b[d];return e},get:function(a,b){return void 0===b?this.cache(a):a[this.expando]&&a[this.expando][r.camelCase(b)]},access:function(a,b,c){return void 0===b||b&&"string"==typeof b&&void 0===c?this.get(a,b):(this.set(a,b,c),void 0!==c?c:b)},remove:function(a,b){var c,d=a[this.expando];if(void 0!==d){if(void 0!==b){Array.isArray(b)?b=b.map(r.camelCase):(b=r.camelCase(b),b=b in d?[b]:b.match(L)||[]),c=b.length;while(c--)delete d[b[c]]}(void 0===b||r.isEmptyObject(d))&&(a.nodeType?a[this.expando]=void 0:delete a[this.expando])}},hasData:function(a){var b=a[this.expando];return void 0!==b&&!r.isEmptyObject(b)}};var W=new V,X=new V,Y=/^(?:\{[\w\W]*\}|\[[\w\W]*\])$/,Z=/[A-Z]/g;function $(a){return"true"===a||"false"!==a&&("null"===a?null:a===+a+""?+a:Y.test(a)?JSON.parse(a):a)}function _(a,b,c){var d;if(void 0===c&&1===a.nodeType)if(d="data-"+b.replace(Z,"-$&").toLowerCase(),c=a.getAttribute(d),"string"==typeof c){try{c=$(c)}catch(e){}X.set(a,b,c)}else c=void 0;return c}r.extend({hasData:function(a){return X.hasData(a)||W.hasData(a)},data:function(a,b,c){return X.access(a,b,c)},removeData:function(a,b){X.remove(a,b)},_data:function(a,b,c){return W.access(a,b,c)},_removeData:function(a,b){W.remove(a,b)}}),r.fn.extend({data:function(a,b){var c,d,e,f=this[0],g=f&&f.attributes;if(void 0===a){if(this.length&&(e=X.get(f),1===f.nodeType&&!W.get(f,"hasDataAttrs"))){c=g.length;while(c--)g[c]&&(d=g[c].name,0===d.indexOf("data-")&&(d=r.camelCase(d.slice(5)),_(f,d,e[d])));W.set(f,"hasDataAttrs",!0)}return e}return"object"==typeof a?this.each(function(){X.set(this,a)}):T(this,function(b){var c;if(f&&void 0===b){if(c=X.get(f,a),void 0!==c)return c;if(c=_(f,a),void 0!==c)return c}else this.each(function(){X.set(this,a,b)})},null,b,arguments.length>1,null,!0)},removeData:function(a){return this.each(function(){X.remove(this,a)})}}),r.extend({queue:function(a,b,c){var d;if(a)return b=(b||"fx")+"queue",d=W.get(a,b),c&&(!d||Array.isArray(c)?d=W.access(a,b,r.makeArray(c)):d.push(c)),d||[]},dequeue:function(a,b){b=b||"fx";var c=r.queue(a,b),d=c.length,e=c.shift(),f=r._queueHooks(a,b),g=function(){r.dequeue(a,b)};"inprogress"===e&&(e=c.shift(),d--),e&&("fx"===b&&c.unshift("inprogress"),delete f.stop,e.call(a,g,f)),!d&&f&&f.empty.fire()},_queueHooks:function(a,b){var c=b+"queueHooks";return W.get(a,c)||W.access(a,c,{empty:r.Callbacks("once memory").add(function(){W.remove(a,[b+"queue",c])})})}}),r.fn.extend({queue:function(a,b){var c=2;return"string"!=typeof a&&(b=a,a="fx",c--),arguments.length<c?r.queue(this[0],a):void 0===b?this:this.each(function(){var c=r.queue(this,a,b);r._queueHooks(this,a),"fx"===a&&"inprogress"!==c[0]&&r.dequeue(this,a)})},dequeue:function(a){return this.each(function(){r.dequeue(this,a)})},clearQueue:function(a){return this.queue(a||"fx",[])},promise:function(a,b){var c,d=1,e=r.Deferred(),f=this,g=this.length,h=function(){--d||e.resolveWith(f,[f])};"string"!=typeof a&&(b=a,a=void 0),a=a||"fx";while(g--)c=W.get(f[g],a+"queueHooks"),c&&c.empty&&(d++,c.empty.add(h));return h(),e.promise(b)}});var aa=/[+-]?(?:\d*\.|)\d+(?:[eE][+-]?\d+|)/.source,ba=new RegExp("^(?:([+-])=|)("+aa+")([a-z%]*)$","i"),ca=["Top","Right","Bottom","Left"],da=function(a,b){return a=b||a,"none"===a.style.display||""===a.style.display&&r.contains(a.ownerDocument,a)&&"none"===r.css(a,"display")},ea=function(a,b,c,d){var e,f,g={};for(f in b)g[f]=a.style[f],a.style[f]=b[f];e=c.apply(a,d||[]);for(f in b)a.style[f]=g[f];return e};function fa(a,b,c,d){var e,f=1,g=20,h=d?function(){return d.cur()}:function(){return r.css(a,b,"")},i=h(),j=c&&c[3]||(r.cssNumber[b]?"":"px"),k=(r.cssNumber[b]||"px"!==j&&+i)&&ba.exec(r.css(a,b));if(k&&k[3]!==j){j=j||k[3],c=c||[],k=+i||1;do f=f||".5",k/=f,r.style(a,b,k+j);while(f!==(f=h()/i)&&1!==f&&--g)}return c&&(k=+k||+i||0,e=c[1]?k+(c[1]+1)*c[2]:+c[2],d&&(d.unit=j,d.start=k,d.end=e)),e}var ga={};function ha(a){var b,c=a.ownerDocument,d=a.nodeName,e=ga[d];return e?e:(b=c.body.appendChild(c.createElement(d)),e=r.css(b,"display"),b.parentNode.removeChild(b),"none"===e&&(e="block"),ga[d]=e,e)}function ia(a,b){for(var c,d,e=[],f=0,g=a.length;f<g;f++)d=a[f],d.style&&(c=d.style.display,b?("none"===c&&(e[f]=W.get(d,"display")||null,e[f]||(d.style.display="")),""===d.style.display&&da(d)&&(e[f]=ha(d))):"none"!==c&&(e[f]="none",W.set(d,"display",c)));for(f=0;f<g;f++)null!=e[f]&&(a[f].style.display=e[f]);return a}r.fn.extend({show:function(){return ia(this,!0)},hide:function(){return ia(this)},toggle:function(a){return"boolean"==typeof a?a?this.show():this.hide():this.each(function(){da(this)?r(this).show():r(this).hide()})}});var ja=/^(?:checkbox|radio)$/i,ka=/<([a-z][^\/\0>\x20\t\r\n\f]+)/i,la=/^$|\/(?:java|ecma)script/i,ma={option:[1,"<select multiple='multiple'>","</select>"],thead:[1,"<table>","</table>"],col:[2,"<table><colgroup>","</colgroup></table>"],tr:[2,"<table><tbody>","</tbody></table>"],td:[3,"<table><tbody><tr>","</tr></tbody></table>"],_default:[0,"",""]};ma.optgroup=ma.option,ma.tbody=ma.tfoot=ma.colgroup=ma.caption=ma.thead,ma.th=ma.td;function na(a,b){var c;return c="undefined"!=typeof a.getElementsByTagName?a.getElementsByTagName(b||"*"):"undefined"!=typeof a.querySelectorAll?a.querySelectorAll(b||"*"):[],void 0===b||b&&B(a,b)?r.merge([a],c):c}function oa(a,b){for(var c=0,d=a.length;c<d;c++)W.set(a[c],"globalEval",!b||W.get(b[c],"globalEval"))}var pa=/<|&#?\w+;/;function qa(a,b,c,d,e){for(var f,g,h,i,j,k,l=b.createDocumentFragment(),m=[],n=0,o=a.length;n<o;n++)if(f=a[n],f||0===f)if("object"===r.type(f))r.merge(m,f.nodeType?[f]:f);else if(pa.test(f)){g=g||l.appendChild(b.createElement("div")),h=(ka.exec(f)||["",""])[1].toLowerCase(),i=ma[h]||ma._default,g.innerHTML=i[1]+r.htmlPrefilter(f)+i[2],k=i[0];while(k--)g=g.lastChild;r.merge(m,g.childNodes),g=l.firstChild,g.textContent=""}else m.push(b.createTextNode(f));l.textContent="",n=0;while(f=m[n++])if(d&&r.inArray(f,d)>-1)e&&e.push(f);else if(j=r.contains(f.ownerDocument,f),g=na(l.appendChild(f),"script"),j&&oa(g),c){k=0;while(f=g[k++])la.test(f.type||"")&&c.push(f)}return l}!function(){var a=d.createDocumentFragment(),b=a.appendChild(d.createElement("div")),c=d.createElement("input");c.setAttribute("type","radio"),c.setAttribute("checked","checked"),c.setAttribute("name","t"),b.appendChild(c),o.checkClone=b.cloneNode(!0).cloneNode(!0).lastChild.checked,b.innerHTML="<textarea>x</textarea>",o.noCloneChecked=!!b.cloneNode(!0).lastChild.defaultValue}();var ra=d.documentElement,sa=/^key/,ta=/^(?:mouse|pointer|contextmenu|drag|drop)|click/,ua=/^([^.]*)(?:\.(.+)|)/;function va(){return!0}function wa(){return!1}function xa(){try{return d.activeElement}catch(a){}}function ya(a,b,c,d,e,f){var g,h;if("object"==typeof b){"string"!=typeof c&&(d=d||c,c=void 0);for(h in b)ya(a,h,c,d,b[h],f);return a}if(null==d&&null==e?(e=c,d=c=void 0):null==e&&("string"==typeof c?(e=d,d=void 0):(e=d,d=c,c=void 0)),e===!1)e=wa;else if(!e)return a;return 1===f&&(g=e,e=function(a){return r().off(a),g.apply(this,arguments)},e.guid=g.guid||(g.guid=r.guid++)),a.each(function(){r.event.add(this,b,e,d,c)})}r.event={global:{},add:function(a,b,c,d,e){var f,g,h,i,j,k,l,m,n,o,p,q=W.get(a);if(q){c.handler&&(f=c,c=f.handler,e=f.selector),e&&r.find.matchesSelector(ra,e),c.guid||(c.guid=r.guid++),(i=q.events)||(i=q.events={}),(g=q.handle)||(g=q.handle=function(b){return"undefined"!=typeof r&&r.event.triggered!==b.type?r.event.dispatch.apply(a,arguments):void 0}),b=(b||"").match(L)||[""],j=b.length;while(j--)h=ua.exec(b[j])||[],n=p=h[1],o=(h[2]||"").split(".").sort(),n&&(l=r.event.special[n]||{},n=(e?l.delegateType:l.bindType)||n,l=r.event.special[n]||{},k=r.extend({type:n,origType:p,data:d,handler:c,guid:c.guid,selector:e,needsContext:e&&r.expr.match.needsContext.test(e),namespace:o.join(".")},f),(m=i[n])||(m=i[n]=[],m.delegateCount=0,l.setup&&l.setup.call(a,d,o,g)!==!1||a.addEventListener&&a.addEventListener(n,g)),l.add&&(l.add.call(a,k),k.handler.guid||(k.handler.guid=c.guid)),e?m.splice(m.delegateCount++,0,k):m.push(k),r.event.global[n]=!0)}},remove:function(a,b,c,d,e){var f,g,h,i,j,k,l,m,n,o,p,q=W.hasData(a)&&W.get(a);if(q&&(i=q.events)){b=(b||"").match(L)||[""],j=b.length;while(j--)if(h=ua.exec(b[j])||[],n=p=h[1],o=(h[2]||"").split(".").sort(),n){l=r.event.special[n]||{},n=(d?l.delegateType:l.bindType)||n,m=i[n]||[],h=h[2]&&new RegExp("(^|\\.)"+o.join("\\.(?:.*\\.|)")+"(\\.|$)"),g=f=m.length;while(f--)k=m[f],!e&&p!==k.origType||c&&c.guid!==k.guid||h&&!h.test(k.namespace)||d&&d!==k.selector&&("**"!==d||!k.selector)||(m.splice(f,1),k.selector&&m.delegateCount--,l.remove&&l.remove.call(a,k));g&&!m.length&&(l.teardown&&l.teardown.call(a,o,q.handle)!==!1||r.removeEvent(a,n,q.handle),delete i[n])}else for(n in i)r.event.remove(a,n+b[j],c,d,!0);r.isEmptyObject(i)&&W.remove(a,"handle events")}},dispatch:function(a){var b=r.event.fix(a),c,d,e,f,g,h,i=new Array(arguments.length),j=(W.get(this,"events")||{})[b.type]||[],k=r.event.special[b.type]||{};for(i[0]=b,c=1;c<arguments.length;c++)i[c]=arguments[c];if(b.delegateTarget=this,!k.preDispatch||k.preDispatch.call(this,b)!==!1){h=r.event.handlers.call(this,b,j),c=0;while((f=h[c++])&&!b.isPropagationStopped()){b.currentTarget=f.elem,d=0;while((g=f.handlers[d++])&&!b.isImmediatePropagationStopped())b.rnamespace&&!b.rnamespace.test(g.namespace)||(b.handleObj=g,b.data=g.data,e=((r.event.special[g.origType]||{}).handle||g.handler).apply(f.elem,i),void 0!==e&&(b.result=e)===!1&&(b.preventDefault(),b.stopPropagation()))}return k.postDispatch&&k.postDispatch.call(this,b),b.result}},handlers:function(a,b){var c,d,e,f,g,h=[],i=b.delegateCount,j=a.target;if(i&&j.nodeType&&!("click"===a.type&&a.button>=1))for(;j!==this;j=j.parentNode||this)if(1===j.nodeType&&("click"!==a.type||j.disabled!==!0)){for(f=[],g={},c=0;c<i;c++)d=b[c],e=d.selector+" ",void 0===g[e]&&(g[e]=d.needsContext?r(e,this).index(j)>-1:r.find(e,this,null,[j]).length),g[e]&&f.push(d);f.length&&h.push({elem:j,handlers:f})}return j=this,i<b.length&&h.push({elem:j,handlers:b.slice(i)}),h},addProp:function(a,b){Object.defineProperty(r.Event.prototype,a,{enumerable:!0,configurable:!0,get:r.isFunction(b)?function(){if(this.originalEvent)return b(this.originalEvent)}:function(){if(this.originalEvent)return this.originalEvent[a]},set:function(b){Object.defineProperty(this,a,{enumerable:!0,configurable:!0,writable:!0,value:b})}})},fix:function(a){return a[r.expando]?a:new r.Event(a)},special:{load:{noBubble:!0},focus:{trigger:function(){if(this!==xa()&&this.focus)return this.focus(),!1},delegateType:"focusin"},blur:{trigger:function(){if(this===xa()&&this.blur)return this.blur(),!1},delegateType:"focusout"},click:{trigger:function(){if("checkbox"===this.type&&this.click&&B(this,"input"))return this.click(),!1},_default:function(a){return B(a.target,"a")}},beforeunload:{postDispatch:function(a){void 0!==a.result&&a.originalEvent&&(a.originalEvent.returnValue=a.result)}}}},r.removeEvent=function(a,b,c){a.removeEventListener&&a.removeEventListener(b,c)},r.Event=function(a,b){return this instanceof r.Event?(a&&a.type?(this.originalEvent=a,this.type=a.type,this.isDefaultPrevented=a.defaultPrevented||void 0===a.defaultPrevented&&a.returnValue===!1?va:wa,this.target=a.target&&3===a.target.nodeType?a.target.parentNode:a.target,this.currentTarget=a.currentTarget,this.relatedTarget=a.relatedTarget):this.type=a,b&&r.extend(this,b),this.timeStamp=a&&a.timeStamp||r.now(),void(this[r.expando]=!0)):new r.Event(a,b)},r.Event.prototype={constructor:r.Event,isDefaultPrevented:wa,isPropagationStopped:wa,isImmediatePropagationStopped:wa,isSimulated:!1,preventDefault:function(){var a=this.originalEvent;this.isDefaultPrevented=va,a&&!this.isSimulated&&a.preventDefault()},stopPropagation:function(){var a=this.originalEvent;this.isPropagationStopped=va,a&&!this.isSimulated&&a.stopPropagation()},stopImmediatePropagation:function(){var a=this.originalEvent;this.isImmediatePropagationStopped=va,a&&!this.isSimulated&&a.stopImmediatePropagation(),this.stopPropagation()}},r.each({altKey:!0,bubbles:!0,cancelable:!0,changedTouches:!0,ctrlKey:!0,detail:!0,eventPhase:!0,metaKey:!0,pageX:!0,pageY:!0,shiftKey:!0,view:!0,"char":!0,charCode:!0,key:!0,keyCode:!0,button:!0,buttons:!0,clientX:!0,clientY:!0,offsetX:!0,offsetY:!0,pointerId:!0,pointerType:!0,screenX:!0,screenY:!0,targetTouches:!0,toElement:!0,touches:!0,which:function(a){var b=a.button;return null==a.which&&sa.test(a.type)?null!=a.charCode?a.charCode:a.keyCode:!a.which&&void 0!==b&&ta.test(a.type)?1&b?1:2&b?3:4&b?2:0:a.which}},r.event.addProp),r.each({mouseenter:"mouseover",mouseleave:"mouseout",pointerenter:"pointerover",pointerleave:"pointerout"},function(a,b){r.event.special[a]={delegateType:b,bindType:b,handle:function(a){var c,d=this,e=a.relatedTarget,f=a.handleObj;return e&&(e===d||r.contains(d,e))||(a.type=f.origType,c=f.handler.apply(this,arguments),a.type=b),c}}}),r.fn.extend({on:function(a,b,c,d){return ya(this,a,b,c,d)},one:function(a,b,c,d){return ya(this,a,b,c,d,1)},off:function(a,b,c){var d,e;if(a&&a.preventDefault&&a.handleObj)return d=a.handleObj,r(a.delegateTarget).off(d.namespace?d.origType+"."+d.namespace:d.origType,d.selector,d.handler),this;if("object"==typeof a){for(e in a)this.off(e,b,a[e]);return this}return b!==!1&&"function"!=typeof b||(c=b,b=void 0),c===!1&&(c=wa),this.each(function(){r.event.remove(this,a,c,b)})}});var za=/<(?!area|br|col|embed|hr|img|input|link|meta|param)(([a-z][^\/\0>\x20\t\r\n\f]*)[^>]*)\/>/gi,Aa=/<script|<style|<link/i,Ba=/checked\s*(?:[^=]|=\s*.checked.)/i,Ca=/^true\/(.*)/,Da=/^\s*<!(?:\[CDATA\[|--)|(?:\]\]|--)>\s*$/g;function Ea(a,b){return B(a,"table")&&B(11!==b.nodeType?b:b.firstChild,"tr")?r(">tbody",a)[0]||a:a}function Fa(a){return a.type=(null!==a.getAttribute("type"))+"/"+a.type,a}function Ga(a){var b=Ca.exec(a.type);return b?a.type=b[1]:a.removeAttribute("type"),a}function Ha(a,b){var c,d,e,f,g,h,i,j;if(1===b.nodeType){if(W.hasData(a)&&(f=W.access(a),g=W.set(b,f),j=f.events)){delete g.handle,g.events={};for(e in j)for(c=0,d=j[e].length;c<d;c++)r.event.add(b,e,j[e][c])}X.hasData(a)&&(h=X.access(a),i=r.extend({},h),X.set(b,i))}}function Ia(a,b){var c=b.nodeName.toLowerCase();"input"===c&&ja.test(a.type)?b.checked=a.checked:"input"!==c&&"textarea"!==c||(b.defaultValue=a.defaultValue)}function Ja(a,b,c,d){b=g.apply([],b);var e,f,h,i,j,k,l=0,m=a.length,n=m-1,q=b[0],s=r.isFunction(q);if(s||m>1&&"string"==typeof q&&!o.checkClone&&Ba.test(q))return a.each(function(e){var f=a.eq(e);s&&(b[0]=q.call(this,e,f.html())),Ja(f,b,c,d)});if(m&&(e=qa(b,a[0].ownerDocument,!1,a,d),f=e.firstChild,1===e.childNodes.length&&(e=f),f||d)){for(h=r.map(na(e,"script"),Fa),i=h.length;l<m;l++)j=e,l!==n&&(j=r.clone(j,!0,!0),i&&r.merge(h,na(j,"script"))),c.call(a[l],j,l);if(i)for(k=h[h.length-1].ownerDocument,r.map(h,Ga),l=0;l<i;l++)j=h[l],la.test(j.type||"")&&!W.access(j,"globalEval")&&r.contains(k,j)&&(j.src?r._evalUrl&&r._evalUrl(j.src):p(j.textContent.replace(Da,""),k))}return a}function Ka(a,b,c){for(var d,e=b?r.filter(b,a):a,f=0;null!=(d=e[f]);f++)c||1!==d.nodeType||r.cleanData(na(d)),d.parentNode&&(c&&r.contains(d.ownerDocument,d)&&oa(na(d,"script")),d.parentNode.removeChild(d));return a}r.extend({htmlPrefilter:function(a){return a.replace(za,"<$1></$2>")},clone:function(a,b,c){var d,e,f,g,h=a.cloneNode(!0),i=r.contains(a.ownerDocument,a);if(!(o.noCloneChecked||1!==a.nodeType&&11!==a.nodeType||r.isXMLDoc(a)))for(g=na(h),f=na(a),d=0,e=f.length;d<e;d++)Ia(f[d],g[d]);if(b)if(c)for(f=f||na(a),g=g||na(h),d=0,e=f.length;d<e;d++)Ha(f[d],g[d]);else Ha(a,h);return g=na(h,"script"),g.length>0&&oa(g,!i&&na(a,"script")),h},cleanData:function(a){for(var b,c,d,e=r.event.special,f=0;void 0!==(c=a[f]);f++)if(U(c)){if(b=c[W.expando]){if(b.events)for(d in b.events)e[d]?r.event.remove(c,d):r.removeEvent(c,d,b.handle);c[W.expando]=void 0}c[X.expando]&&(c[X.expando]=void 0)}}}),r.fn.extend({detach:function(a){return Ka(this,a,!0)},remove:function(a){return Ka(this,a)},text:function(a){return T(this,function(a){return void 0===a?r.text(this):this.empty().each(function(){1!==this.nodeType&&11!==this.nodeType&&9!==this.nodeType||(this.textContent=a)})},null,a,arguments.length)},append:function(){return Ja(this,arguments,function(a){if(1===this.nodeType||11===this.nodeType||9===this.nodeType){var b=Ea(this,a);b.appendChild(a)}})},prepend:function(){return Ja(this,arguments,function(a){if(1===this.nodeType||11===this.nodeType||9===this.nodeType){var b=Ea(this,a);b.insertBefore(a,b.firstChild)}})},before:function(){return Ja(this,arguments,function(a){this.parentNode&&this.parentNode.insertBefore(a,this)})},after:function(){return Ja(this,arguments,function(a){this.parentNode&&this.parentNode.insertBefore(a,this.nextSibling)})},empty:function(){for(var a,b=0;null!=(a=this[b]);b++)1===a.nodeType&&(r.cleanData(na(a,!1)),a.textContent="");return this},clone:function(a,b){return a=null!=a&&a,b=null==b?a:b,this.map(function(){return r.clone(this,a,b)})},html:function(a){return T(this,function(a){var b=this[0]||{},c=0,d=this.length;if(void 0===a&&1===b.nodeType)return b.innerHTML;if("string"==typeof a&&!Aa.test(a)&&!ma[(ka.exec(a)||["",""])[1].toLowerCase()]){a=r.htmlPrefilter(a);try{for(;c<d;c++)b=this[c]||{},1===b.nodeType&&(r.cleanData(na(b,!1)),b.innerHTML=a);b=0}catch(e){}}b&&this.empty().append(a)},null,a,arguments.length)},replaceWith:function(){var a=[];return Ja(this,arguments,function(b){var c=this.parentNode;r.inArray(this,a)<0&&(r.cleanData(na(this)),c&&c.replaceChild(b,this))},a)}}),r.each({appendTo:"append",prependTo:"prepend",insertBefore:"before",insertAfter:"after",replaceAll:"replaceWith"},function(a,b){r.fn[a]=function(a){for(var c,d=[],e=r(a),f=e.length-1,g=0;g<=f;g++)c=g===f?this:this.clone(!0),r(e[g])[b](c),h.apply(d,c.get());return this.pushStack(d)}});var La=/^margin/,Ma=new RegExp("^("+aa+")(?!px)[a-z%]+$","i"),Na=function(b){var c=b.ownerDocument.defaultView;return c&&c.opener||(c=a),c.getComputedStyle(b)};!function(){function b(){if(i){i.style.cssText="box-sizing:border-box;position:relative;display:block;margin:auto;border:1px;padding:1px;top:1%;width:50%",i.innerHTML="",ra.appendChild(h);var b=a.getComputedStyle(i);c="1%"!==b.top,g="2px"===b.marginLeft,e="4px"===b.width,i.style.marginRight="50%",f="4px"===b.marginRight,ra.removeChild(h),i=null}}var c,e,f,g,h=d.createElement("div"),i=d.createElement("div");i.style&&(i.style.backgroundClip="content-box",i.cloneNode(!0).style.backgroundClip="",o.clearCloneStyle="content-box"===i.style.backgroundClip,h.style.cssText="border:0;width:8px;height:0;top:0;left:-9999px;padding:0;margin-top:1px;position:absolute",h.appendChild(i),r.extend(o,{pixelPosition:function(){return b(),c},boxSizingReliable:function(){return b(),e},pixelMarginRight:function(){return b(),f},reliableMarginLeft:function(){return b(),g}}))}();function Oa(a,b,c){var d,e,f,g,h=a.style;return c=c||Na(a),c&&(g=c.getPropertyValue(b)||c[b],""!==g||r.contains(a.ownerDocument,a)||(g=r.style(a,b)),!o.pixelMarginRight()&&Ma.test(g)&&La.test(b)&&(d=h.width,e=h.minWidth,f=h.maxWidth,h.minWidth=h.maxWidth=h.width=g,g=c.width,h.width=d,h.minWidth=e,h.maxWidth=f)),void 0!==g?g+"":g}function Pa(a,b){return{get:function(){return a()?void delete this.get:(this.get=b).apply(this,arguments)}}}var Qa=/^(none|table(?!-c[ea]).+)/,Ra=/^--/,Sa={position:"absolute",visibility:"hidden",display:"block"},Ta={letterSpacing:"0",fontWeight:"400"},Ua=["Webkit","Moz","ms"],Va=d.createElement("div").style;function Wa(a){if(a in Va)return a;var b=a[0].toUpperCase()+a.slice(1),c=Ua.length;while(c--)if(a=Ua[c]+b,a in Va)return a}function Xa(a){var b=r.cssProps[a];return b||(b=r.cssProps[a]=Wa(a)||a),b}function Ya(a,b,c){var d=ba.exec(b);return d?Math.max(0,d[2]-(c||0))+(d[3]||"px"):b}function Za(a,b,c,d,e){var f,g=0;for(f=c===(d?"border":"content")?4:"width"===b?1:0;f<4;f+=2)"margin"===c&&(g+=r.css(a,c+ca[f],!0,e)),d?("content"===c&&(g-=r.css(a,"padding"+ca[f],!0,e)),"margin"!==c&&(g-=r.css(a,"border"+ca[f]+"Width",!0,e))):(g+=r.css(a,"padding"+ca[f],!0,e),"padding"!==c&&(g+=r.css(a,"border"+ca[f]+"Width",!0,e)));return g}function $a(a,b,c){var d,e=Na(a),f=Oa(a,b,e),g="border-box"===r.css(a,"boxSizing",!1,e);return Ma.test(f)?f:(d=g&&(o.boxSizingReliable()||f===a.style[b]),"auto"===f&&(f=a["offset"+b[0].toUpperCase()+b.slice(1)]),f=parseFloat(f)||0,f+Za(a,b,c||(g?"border":"content"),d,e)+"px")}r.extend({cssHooks:{opacity:{get:function(a,b){if(b){var c=Oa(a,"opacity");return""===c?"1":c}}}},cssNumber:{animationIterationCount:!0,columnCount:!0,fillOpacity:!0,flexGrow:!0,flexShrink:!0,fontWeight:!0,lineHeight:!0,opacity:!0,order:!0,orphans:!0,widows:!0,zIndex:!0,zoom:!0},cssProps:{"float":"cssFloat"},style:function(a,b,c,d){if(a&&3!==a.nodeType&&8!==a.nodeType&&a.style){var e,f,g,h=r.camelCase(b),i=Ra.test(b),j=a.style;return i||(b=Xa(h)),g=r.cssHooks[b]||r.cssHooks[h],void 0===c?g&&"get"in g&&void 0!==(e=g.get(a,!1,d))?e:j[b]:(f=typeof c,"string"===f&&(e=ba.exec(c))&&e[1]&&(c=fa(a,b,e),f="number"),null!=c&&c===c&&("number"===f&&(c+=e&&e[3]||(r.cssNumber[h]?"":"px")),o.clearCloneStyle||""!==c||0!==b.indexOf("background")||(j[b]="inherit"),g&&"set"in g&&void 0===(c=g.set(a,c,d))||(i?j.setProperty(b,c):j[b]=c)),void 0)}},css:function(a,b,c,d){var e,f,g,h=r.camelCase(b),i=Ra.test(b);return i||(b=Xa(h)),g=r.cssHooks[b]||r.cssHooks[h],g&&"get"in g&&(e=g.get(a,!0,c)),void 0===e&&(e=Oa(a,b,d)),"normal"===e&&b in Ta&&(e=Ta[b]),""===c||c?(f=parseFloat(e),c===!0||isFinite(f)?f||0:e):e}}),r.each(["height","width"],function(a,b){r.cssHooks[b]={get:function(a,c,d){if(c)return!Qa.test(r.css(a,"display"))||a.getClientRects().length&&a.getBoundingClientRect().width?$a(a,b,d):ea(a,Sa,function(){return $a(a,b,d)})},set:function(a,c,d){var e,f=d&&Na(a),g=d&&Za(a,b,d,"border-box"===r.css(a,"boxSizing",!1,f),f);return g&&(e=ba.exec(c))&&"px"!==(e[3]||"px")&&(a.style[b]=c,c=r.css(a,b)),Ya(a,c,g)}}}),r.cssHooks.marginLeft=Pa(o.reliableMarginLeft,function(a,b){if(b)return(parseFloat(Oa(a,"marginLeft"))||a.getBoundingClientRect().left-ea(a,{marginLeft:0},function(){return a.getBoundingClientRect().left}))+"px"}),r.each({margin:"",padding:"",border:"Width"},function(a,b){r.cssHooks[a+b]={expand:function(c){for(var d=0,e={},f="string"==typeof c?c.split(" "):[c];d<4;d++)e[a+ca[d]+b]=f[d]||f[d-2]||f[0];return e}},La.test(a)||(r.cssHooks[a+b].set=Ya)}),r.fn.extend({css:function(a,b){return T(this,function(a,b,c){var d,e,f={},g=0;if(Array.isArray(b)){for(d=Na(a),e=b.length;g<e;g++)f[b[g]]=r.css(a,b[g],!1,d);return f}return void 0!==c?r.style(a,b,c):r.css(a,b)},a,b,arguments.length>1)}});function _a(a,b,c,d,e){return new _a.prototype.init(a,b,c,d,e)}r.Tween=_a,_a.prototype={constructor:_a,init:function(a,b,c,d,e,f){this.elem=a,this.prop=c,this.easing=e||r.easing._default,this.options=b,this.start=this.now=this.cur(),this.end=d,this.unit=f||(r.cssNumber[c]?"":"px")},cur:function(){var a=_a.propHooks[this.prop];return a&&a.get?a.get(this):_a.propHooks._default.get(this)},run:function(a){var b,c=_a.propHooks[this.prop];return this.options.duration?this.pos=b=r.easing[this.easing](a,this.options.duration*a,0,1,this.options.duration):this.pos=b=a,this.now=(this.end-this.start)*b+this.start,this.options.step&&this.options.step.call(this.elem,this.now,this),c&&c.set?c.set(this):_a.propHooks._default.set(this),this}},_a.prototype.init.prototype=_a.prototype,_a.propHooks={_default:{get:function(a){var b;return 1!==a.elem.nodeType||null!=a.elem[a.prop]&&null==a.elem.style[a.prop]?a.elem[a.prop]:(b=r.css(a.elem,a.prop,""),b&&"auto"!==b?b:0)},set:function(a){r.fx.step[a.prop]?r.fx.step[a.prop](a):1!==a.elem.nodeType||null==a.elem.style[r.cssProps[a.prop]]&&!r.cssHooks[a.prop]?a.elem[a.prop]=a.now:r.style(a.elem,a.prop,a.now+a.unit)}}},_a.propHooks.scrollTop=_a.propHooks.scrollLeft={set:function(a){a.elem.nodeType&&a.elem.parentNode&&(a.elem[a.prop]=a.now)}},r.easing={linear:function(a){return a},swing:function(a){return.5-Math.cos(a*Math.PI)/2},_default:"swing"},r.fx=_a.prototype.init,r.fx.step={};var ab,bb,cb=/^(?:toggle|show|hide)$/,db=/queueHooks$/;function eb(){bb&&(d.hidden===!1&&a.requestAnimationFrame?a.requestAnimationFrame(eb):a.setTimeout(eb,r.fx.interval),r.fx.tick())}function fb(){return a.setTimeout(function(){ab=void 0}),ab=r.now()}function gb(a,b){var c,d=0,e={height:a};for(b=b?1:0;d<4;d+=2-b)c=ca[d],e["margin"+c]=e["padding"+c]=a;return b&&(e.opacity=e.width=a),e}function hb(a,b,c){for(var d,e=(kb.tweeners[b]||[]).concat(kb.tweeners["*"]),f=0,g=e.length;f<g;f++)if(d=e[f].call(c,b,a))return d}function ib(a,b,c){var d,e,f,g,h,i,j,k,l="width"in b||"height"in b,m=this,n={},o=a.style,p=a.nodeType&&da(a),q=W.get(a,"fxshow");c.queue||(g=r._queueHooks(a,"fx"),null==g.unqueued&&(g.unqueued=0,h=g.empty.fire,g.empty.fire=function(){g.unqueued||h()}),g.unqueued++,m.always(function(){m.always(function(){g.unqueued--,r.queue(a,"fx").length||g.empty.fire()})}));for(d in b)if(e=b[d],cb.test(e)){if(delete b[d],f=f||"toggle"===e,e===(p?"hide":"show")){if("show"!==e||!q||void 0===q[d])continue;p=!0}n[d]=q&&q[d]||r.style(a,d)}if(i=!r.isEmptyObject(b),i||!r.isEmptyObject(n)){l&&1===a.nodeType&&(c.overflow=[o.overflow,o.overflowX,o.overflowY],j=q&&q.display,null==j&&(j=W.get(a,"display")),k=r.css(a,"display"),"none"===k&&(j?k=j:(ia([a],!0),j=a.style.display||j,k=r.css(a,"display"),ia([a]))),("inline"===k||"inline-block"===k&&null!=j)&&"none"===r.css(a,"float")&&(i||(m.done(function(){o.display=j}),null==j&&(k=o.display,j="none"===k?"":k)),o.display="inline-block")),c.overflow&&(o.overflow="hidden",m.always(function(){o.overflow=c.overflow[0],o.overflowX=c.overflow[1],o.overflowY=c.overflow[2]})),i=!1;for(d in n)i||(q?"hidden"in q&&(p=q.hidden):q=W.access(a,"fxshow",{display:j}),f&&(q.hidden=!p),p&&ia([a],!0),m.done(function(){p||ia([a]),W.remove(a,"fxshow");for(d in n)r.style(a,d,n[d])})),i=hb(p?q[d]:0,d,m),d in q||(q[d]=i.start,p&&(i.end=i.start,i.start=0))}}function jb(a,b){var c,d,e,f,g;for(c in a)if(d=r.camelCase(c),e=b[d],f=a[c],Array.isArray(f)&&(e=f[1],f=a[c]=f[0]),c!==d&&(a[d]=f,delete a[c]),g=r.cssHooks[d],g&&"expand"in g){f=g.expand(f),delete a[d];for(c in f)c in a||(a[c]=f[c],b[c]=e)}else b[d]=e}function kb(a,b,c){var d,e,f=0,g=kb.prefilters.length,h=r.Deferred().always(function(){delete i.elem}),i=function(){if(e)return!1;for(var b=ab||fb(),c=Math.max(0,j.startTime+j.duration-b),d=c/j.duration||0,f=1-d,g=0,i=j.tweens.length;g<i;g++)j.tweens[g].run(f);return h.notifyWith(a,[j,f,c]),f<1&&i?c:(i||h.notifyWith(a,[j,1,0]),h.resolveWith(a,[j]),!1)},j=h.promise({elem:a,props:r.extend({},b),opts:r.extend(!0,{specialEasing:{},easing:r.easing._default},c),originalProperties:b,originalOptions:c,startTime:ab||fb(),duration:c.duration,tweens:[],createTween:function(b,c){var d=r.Tween(a,j.opts,b,c,j.opts.specialEasing[b]||j.opts.easing);return j.tweens.push(d),d},stop:function(b){var c=0,d=b?j.tweens.length:0;if(e)return this;for(e=!0;c<d;c++)j.tweens[c].run(1);return b?(h.notifyWith(a,[j,1,0]),h.resolveWith(a,[j,b])):h.rejectWith(a,[j,b]),this}}),k=j.props;for(jb(k,j.opts.specialEasing);f<g;f++)if(d=kb.prefilters[f].call(j,a,k,j.opts))return r.isFunction(d.stop)&&(r._queueHooks(j.elem,j.opts.queue).stop=r.proxy(d.stop,d)),d;return r.map(k,hb,j),r.isFunction(j.opts.start)&&j.opts.start.call(a,j),j.progress(j.opts.progress).done(j.opts.done,j.opts.complete).fail(j.opts.fail).always(j.opts.always),r.fx.timer(r.extend(i,{elem:a,anim:j,queue:j.opts.queue})),j}r.Animation=r.extend(kb,{tweeners:{"*":[function(a,b){var c=this.createTween(a,b);return fa(c.elem,a,ba.exec(b),c),c}]},tweener:function(a,b){r.isFunction(a)?(b=a,a=["*"]):a=a.match(L);for(var c,d=0,e=a.length;d<e;d++)c=a[d],kb.tweeners[c]=kb.tweeners[c]||[],kb.tweeners[c].unshift(b)},prefilters:[ib],prefilter:function(a,b){b?kb.prefilters.unshift(a):kb.prefilters.push(a)}}),r.speed=function(a,b,c){var d=a&&"object"==typeof a?r.extend({},a):{complete:c||!c&&b||r.isFunction(a)&&a,duration:a,easing:c&&b||b&&!r.isFunction(b)&&b};return r.fx.off?d.duration=0:"number"!=typeof d.duration&&(d.duration in r.fx.speeds?d.duration=r.fx.speeds[d.duration]:d.duration=r.fx.speeds._default),null!=d.queue&&d.queue!==!0||(d.queue="fx"),d.old=d.complete,d.complete=function(){r.isFunction(d.old)&&d.old.call(this),d.queue&&r.dequeue(this,d.queue)},d},r.fn.extend({fadeTo:function(a,b,c,d){return this.filter(da).css("opacity",0).show().end().animate({opacity:b},a,c,d)},animate:function(a,b,c,d){var e=r.isEmptyObject(a),f=r.speed(b,c,d),g=function(){var b=kb(this,r.extend({},a),f);(e||W.get(this,"finish"))&&b.stop(!0)};return g.finish=g,e||f.queue===!1?this.each(g):this.queue(f.queue,g)},stop:function(a,b,c){var d=function(a){var b=a.stop;delete a.stop,b(c)};return"string"!=typeof a&&(c=b,b=a,a=void 0),b&&a!==!1&&this.queue(a||"fx",[]),this.each(function(){var b=!0,e=null!=a&&a+"queueHooks",f=r.timers,g=W.get(this);if(e)g[e]&&g[e].stop&&d(g[e]);else for(e in g)g[e]&&g[e].stop&&db.test(e)&&d(g[e]);for(e=f.length;e--;)f[e].elem!==this||null!=a&&f[e].queue!==a||(f[e].anim.stop(c),b=!1,f.splice(e,1));!b&&c||r.dequeue(this,a)})},finish:function(a){return a!==!1&&(a=a||"fx"),this.each(function(){var b,c=W.get(this),d=c[a+"queue"],e=c[a+"queueHooks"],f=r.timers,g=d?d.length:0;for(c.finish=!0,r.queue(this,a,[]),e&&e.stop&&e.stop.call(this,!0),b=f.length;b--;)f[b].elem===this&&f[b].queue===a&&(f[b].anim.stop(!0),f.splice(b,1));for(b=0;b<g;b++)d[b]&&d[b].finish&&d[b].finish.call(this);delete c.finish})}}),r.each(["toggle","show","hide"],function(a,b){var c=r.fn[b];r.fn[b]=function(a,d,e){return null==a||"boolean"==typeof a?c.apply(this,arguments):this.animate(gb(b,!0),a,d,e)}}),r.each({slideDown:gb("show"),slideUp:gb("hide"),slideToggle:gb("toggle"),fadeIn:{opacity:"show"},fadeOut:{opacity:"hide"},fadeToggle:{opacity:"toggle"}},function(a,b){r.fn[a]=function(a,c,d){return this.animate(b,a,c,d)}}),r.timers=[],r.fx.tick=function(){var a,b=0,c=r.timers;for(ab=r.now();b<c.length;b++)a=c[b],a()||c[b]!==a||c.splice(b--,1);c.length||r.fx.stop(),ab=void 0},r.fx.timer=function(a){r.timers.push(a),r.fx.start()},r.fx.interval=13,r.fx.start=function(){bb||(bb=!0,eb())},r.fx.stop=function(){bb=null},r.fx.speeds={slow:600,fast:200,_default:400},r.fn.delay=function(b,c){return b=r.fx?r.fx.speeds[b]||b:b,c=c||"fx",this.queue(c,function(c,d){var e=a.setTimeout(c,b);d.stop=function(){a.clearTimeout(e)}})},function(){var a=d.createElement("input"),b=d.createElement("select"),c=b.appendChild(d.createElement("option"));a.type="checkbox",o.checkOn=""!==a.value,o.optSelected=c.selected,a=d.createElement("input"),a.value="t",a.type="radio",o.radioValue="t"===a.value}();var lb,mb=r.expr.attrHandle;r.fn.extend({attr:function(a,b){return T(this,r.attr,a,b,arguments.length>1)},removeAttr:function(a){return this.each(function(){r.removeAttr(this,a)})}}),r.extend({attr:function(a,b,c){var d,e,f=a.nodeType;if(3!==f&&8!==f&&2!==f)return"undefined"==typeof a.getAttribute?r.prop(a,b,c):(1===f&&r.isXMLDoc(a)||(e=r.attrHooks[b.toLowerCase()]||(r.expr.match.bool.test(b)?lb:void 0)),void 0!==c?null===c?void r.removeAttr(a,b):e&&"set"in e&&void 0!==(d=e.set(a,c,b))?d:(a.setAttribute(b,c+""),c):e&&"get"in e&&null!==(d=e.get(a,b))?d:(d=r.find.attr(a,b),
null==d?void 0:d))},attrHooks:{type:{set:function(a,b){if(!o.radioValue&&"radio"===b&&B(a,"input")){var c=a.value;return a.setAttribute("type",b),c&&(a.value=c),b}}}},removeAttr:function(a,b){var c,d=0,e=b&&b.match(L);if(e&&1===a.nodeType)while(c=e[d++])a.removeAttribute(c)}}),lb={set:function(a,b,c){return b===!1?r.removeAttr(a,c):a.setAttribute(c,c),c}},r.each(r.expr.match.bool.source.match(/\w+/g),function(a,b){var c=mb[b]||r.find.attr;mb[b]=function(a,b,d){var e,f,g=b.toLowerCase();return d||(f=mb[g],mb[g]=e,e=null!=c(a,b,d)?g:null,mb[g]=f),e}});var nb=/^(?:input|select|textarea|button)$/i,ob=/^(?:a|area)$/i;r.fn.extend({prop:function(a,b){return T(this,r.prop,a,b,arguments.length>1)},removeProp:function(a){return this.each(function(){delete this[r.propFix[a]||a]})}}),r.extend({prop:function(a,b,c){var d,e,f=a.nodeType;if(3!==f&&8!==f&&2!==f)return 1===f&&r.isXMLDoc(a)||(b=r.propFix[b]||b,e=r.propHooks[b]),void 0!==c?e&&"set"in e&&void 0!==(d=e.set(a,c,b))?d:a[b]=c:e&&"get"in e&&null!==(d=e.get(a,b))?d:a[b]},propHooks:{tabIndex:{get:function(a){var b=r.find.attr(a,"tabindex");return b?parseInt(b,10):nb.test(a.nodeName)||ob.test(a.nodeName)&&a.href?0:-1}}},propFix:{"for":"htmlFor","class":"className"}}),o.optSelected||(r.propHooks.selected={get:function(a){var b=a.parentNode;return b&&b.parentNode&&b.parentNode.selectedIndex,null},set:function(a){var b=a.parentNode;b&&(b.selectedIndex,b.parentNode&&b.parentNode.selectedIndex)}}),r.each(["tabIndex","readOnly","maxLength","cellSpacing","cellPadding","rowSpan","colSpan","useMap","frameBorder","contentEditable"],function(){r.propFix[this.toLowerCase()]=this});function pb(a){var b=a.match(L)||[];return b.join(" ")}function qb(a){return a.getAttribute&&a.getAttribute("class")||""}r.fn.extend({addClass:function(a){var b,c,d,e,f,g,h,i=0;if(r.isFunction(a))return this.each(function(b){r(this).addClass(a.call(this,b,qb(this)))});if("string"==typeof a&&a){b=a.match(L)||[];while(c=this[i++])if(e=qb(c),d=1===c.nodeType&&" "+pb(e)+" "){g=0;while(f=b[g++])d.indexOf(" "+f+" ")<0&&(d+=f+" ");h=pb(d),e!==h&&c.setAttribute("class",h)}}return this},removeClass:function(a){var b,c,d,e,f,g,h,i=0;if(r.isFunction(a))return this.each(function(b){r(this).removeClass(a.call(this,b,qb(this)))});if(!arguments.length)return this.attr("class","");if("string"==typeof a&&a){b=a.match(L)||[];while(c=this[i++])if(e=qb(c),d=1===c.nodeType&&" "+pb(e)+" "){g=0;while(f=b[g++])while(d.indexOf(" "+f+" ")>-1)d=d.replace(" "+f+" "," ");h=pb(d),e!==h&&c.setAttribute("class",h)}}return this},toggleClass:function(a,b){var c=typeof a;return"boolean"==typeof b&&"string"===c?b?this.addClass(a):this.removeClass(a):r.isFunction(a)?this.each(function(c){r(this).toggleClass(a.call(this,c,qb(this),b),b)}):this.each(function(){var b,d,e,f;if("string"===c){d=0,e=r(this),f=a.match(L)||[];while(b=f[d++])e.hasClass(b)?e.removeClass(b):e.addClass(b)}else void 0!==a&&"boolean"!==c||(b=qb(this),b&&W.set(this,"__className__",b),this.setAttribute&&this.setAttribute("class",b||a===!1?"":W.get(this,"__className__")||""))})},hasClass:function(a){var b,c,d=0;b=" "+a+" ";while(c=this[d++])if(1===c.nodeType&&(" "+pb(qb(c))+" ").indexOf(b)>-1)return!0;return!1}});var rb=/\r/g;r.fn.extend({val:function(a){var b,c,d,e=this[0];{if(arguments.length)return d=r.isFunction(a),this.each(function(c){var e;1===this.nodeType&&(e=d?a.call(this,c,r(this).val()):a,null==e?e="":"number"==typeof e?e+="":Array.isArray(e)&&(e=r.map(e,function(a){return null==a?"":a+""})),b=r.valHooks[this.type]||r.valHooks[this.nodeName.toLowerCase()],b&&"set"in b&&void 0!==b.set(this,e,"value")||(this.value=e))});if(e)return b=r.valHooks[e.type]||r.valHooks[e.nodeName.toLowerCase()],b&&"get"in b&&void 0!==(c=b.get(e,"value"))?c:(c=e.value,"string"==typeof c?c.replace(rb,""):null==c?"":c)}}}),r.extend({valHooks:{option:{get:function(a){var b=r.find.attr(a,"value");return null!=b?b:pb(r.text(a))}},select:{get:function(a){var b,c,d,e=a.options,f=a.selectedIndex,g="select-one"===a.type,h=g?null:[],i=g?f+1:e.length;for(d=f<0?i:g?f:0;d<i;d++)if(c=e[d],(c.selected||d===f)&&!c.disabled&&(!c.parentNode.disabled||!B(c.parentNode,"optgroup"))){if(b=r(c).val(),g)return b;h.push(b)}return h},set:function(a,b){var c,d,e=a.options,f=r.makeArray(b),g=e.length;while(g--)d=e[g],(d.selected=r.inArray(r.valHooks.option.get(d),f)>-1)&&(c=!0);return c||(a.selectedIndex=-1),f}}}}),r.each(["radio","checkbox"],function(){r.valHooks[this]={set:function(a,b){if(Array.isArray(b))return a.checked=r.inArray(r(a).val(),b)>-1}},o.checkOn||(r.valHooks[this].get=function(a){return null===a.getAttribute("value")?"on":a.value})});var sb=/^(?:focusinfocus|focusoutblur)$/;r.extend(r.event,{trigger:function(b,c,e,f){var g,h,i,j,k,m,n,o=[e||d],p=l.call(b,"type")?b.type:b,q=l.call(b,"namespace")?b.namespace.split("."):[];if(h=i=e=e||d,3!==e.nodeType&&8!==e.nodeType&&!sb.test(p+r.event.triggered)&&(p.indexOf(".")>-1&&(q=p.split("."),p=q.shift(),q.sort()),k=p.indexOf(":")<0&&"on"+p,b=b[r.expando]?b:new r.Event(p,"object"==typeof b&&b),b.isTrigger=f?2:3,b.namespace=q.join("."),b.rnamespace=b.namespace?new RegExp("(^|\\.)"+q.join("\\.(?:.*\\.|)")+"(\\.|$)"):null,b.result=void 0,b.target||(b.target=e),c=null==c?[b]:r.makeArray(c,[b]),n=r.event.special[p]||{},f||!n.trigger||n.trigger.apply(e,c)!==!1)){if(!f&&!n.noBubble&&!r.isWindow(e)){for(j=n.delegateType||p,sb.test(j+p)||(h=h.parentNode);h;h=h.parentNode)o.push(h),i=h;i===(e.ownerDocument||d)&&o.push(i.defaultView||i.parentWindow||a)}g=0;while((h=o[g++])&&!b.isPropagationStopped())b.type=g>1?j:n.bindType||p,m=(W.get(h,"events")||{})[b.type]&&W.get(h,"handle"),m&&m.apply(h,c),m=k&&h[k],m&&m.apply&&U(h)&&(b.result=m.apply(h,c),b.result===!1&&b.preventDefault());return b.type=p,f||b.isDefaultPrevented()||n._default&&n._default.apply(o.pop(),c)!==!1||!U(e)||k&&r.isFunction(e[p])&&!r.isWindow(e)&&(i=e[k],i&&(e[k]=null),r.event.triggered=p,e[p](),r.event.triggered=void 0,i&&(e[k]=i)),b.result}},simulate:function(a,b,c){var d=r.extend(new r.Event,c,{type:a,isSimulated:!0});r.event.trigger(d,null,b)}}),r.fn.extend({trigger:function(a,b){return this.each(function(){r.event.trigger(a,b,this)})},triggerHandler:function(a,b){var c=this[0];if(c)return r.event.trigger(a,b,c,!0)}}),r.each("blur focus focusin focusout resize scroll click dblclick mousedown mouseup mousemove mouseover mouseout mouseenter mouseleave change select submit keydown keypress keyup contextmenu".split(" "),function(a,b){r.fn[b]=function(a,c){return arguments.length>0?this.on(b,null,a,c):this.trigger(b)}}),r.fn.extend({hover:function(a,b){return this.mouseenter(a).mouseleave(b||a)}}),o.focusin="onfocusin"in a,o.focusin||r.each({focus:"focusin",blur:"focusout"},function(a,b){var c=function(a){r.event.simulate(b,a.target,r.event.fix(a))};r.event.special[b]={setup:function(){var d=this.ownerDocument||this,e=W.access(d,b);e||d.addEventListener(a,c,!0),W.access(d,b,(e||0)+1)},teardown:function(){var d=this.ownerDocument||this,e=W.access(d,b)-1;e?W.access(d,b,e):(d.removeEventListener(a,c,!0),W.remove(d,b))}}});var tb=a.location,ub=r.now(),vb=/\?/;r.parseXML=function(b){var c;if(!b||"string"!=typeof b)return null;try{c=(new a.DOMParser).parseFromString(b,"text/xml")}catch(d){c=void 0}return c&&!c.getElementsByTagName("parsererror").length||r.error("Invalid XML: "+b),c};var wb=/\[\]$/,xb=/\r?\n/g,yb=/^(?:submit|button|image|reset|file)$/i,zb=/^(?:input|select|textarea|keygen)/i;function Ab(a,b,c,d){var e;if(Array.isArray(b))r.each(b,function(b,e){c||wb.test(a)?d(a,e):Ab(a+"["+("object"==typeof e&&null!=e?b:"")+"]",e,c,d)});else if(c||"object"!==r.type(b))d(a,b);else for(e in b)Ab(a+"["+e+"]",b[e],c,d)}r.param=function(a,b){var c,d=[],e=function(a,b){var c=r.isFunction(b)?b():b;d[d.length]=encodeURIComponent(a)+"="+encodeURIComponent(null==c?"":c)};if(Array.isArray(a)||a.jquery&&!r.isPlainObject(a))r.each(a,function(){e(this.name,this.value)});else for(c in a)Ab(c,a[c],b,e);return d.join("&")},r.fn.extend({serialize:function(){return r.param(this.serializeArray())},serializeArray:function(){return this.map(function(){var a=r.prop(this,"elements");return a?r.makeArray(a):this}).filter(function(){var a=this.type;return this.name&&!r(this).is(":disabled")&&zb.test(this.nodeName)&&!yb.test(a)&&(this.checked||!ja.test(a))}).map(function(a,b){var c=r(this).val();return null==c?null:Array.isArray(c)?r.map(c,function(a){return{name:b.name,value:a.replace(xb,"\r\n")}}):{name:b.name,value:c.replace(xb,"\r\n")}}).get()}});var Bb=/%20/g,Cb=/#.*$/,Db=/([?&])_=[^&]*/,Eb=/^(.*?):[ \t]*([^\r\n]*)$/gm,Fb=/^(?:about|app|app-storage|.+-extension|file|res|widget):$/,Gb=/^(?:GET|HEAD)$/,Hb=/^\/\//,Ib={},Jb={},Kb="*/".concat("*"),Lb=d.createElement("a");Lb.href=tb.href;function Mb(a){return function(b,c){"string"!=typeof b&&(c=b,b="*");var d,e=0,f=b.toLowerCase().match(L)||[];if(r.isFunction(c))while(d=f[e++])"+"===d[0]?(d=d.slice(1)||"*",(a[d]=a[d]||[]).unshift(c)):(a[d]=a[d]||[]).push(c)}}function Nb(a,b,c,d){var e={},f=a===Jb;function g(h){var i;return e[h]=!0,r.each(a[h]||[],function(a,h){var j=h(b,c,d);return"string"!=typeof j||f||e[j]?f?!(i=j):void 0:(b.dataTypes.unshift(j),g(j),!1)}),i}return g(b.dataTypes[0])||!e["*"]&&g("*")}function Ob(a,b){var c,d,e=r.ajaxSettings.flatOptions||{};for(c in b)void 0!==b[c]&&((e[c]?a:d||(d={}))[c]=b[c]);return d&&r.extend(!0,a,d),a}function Pb(a,b,c){var d,e,f,g,h=a.contents,i=a.dataTypes;while("*"===i[0])i.shift(),void 0===d&&(d=a.mimeType||b.getResponseHeader("Content-Type"));if(d)for(e in h)if(h[e]&&h[e].test(d)){i.unshift(e);break}if(i[0]in c)f=i[0];else{for(e in c){if(!i[0]||a.converters[e+" "+i[0]]){f=e;break}g||(g=e)}f=f||g}if(f)return f!==i[0]&&i.unshift(f),c[f]}function Qb(a,b,c,d){var e,f,g,h,i,j={},k=a.dataTypes.slice();if(k[1])for(g in a.converters)j[g.toLowerCase()]=a.converters[g];f=k.shift();while(f)if(a.responseFields[f]&&(c[a.responseFields[f]]=b),!i&&d&&a.dataFilter&&(b=a.dataFilter(b,a.dataType)),i=f,f=k.shift())if("*"===f)f=i;else if("*"!==i&&i!==f){if(g=j[i+" "+f]||j["* "+f],!g)for(e in j)if(h=e.split(" "),h[1]===f&&(g=j[i+" "+h[0]]||j["* "+h[0]])){g===!0?g=j[e]:j[e]!==!0&&(f=h[0],k.unshift(h[1]));break}if(g!==!0)if(g&&a["throws"])b=g(b);else try{b=g(b)}catch(l){return{state:"parsererror",error:g?l:"No conversion from "+i+" to "+f}}}return{state:"success",data:b}}r.extend({active:0,lastModified:{},etag:{},ajaxSettings:{url:tb.href,type:"GET",isLocal:Fb.test(tb.protocol),global:!0,processData:!0,async:!0,contentType:"application/x-www-form-urlencoded; charset=UTF-8",accepts:{"*":Kb,text:"text/plain",html:"text/html",xml:"application/xml, text/xml",json:"application/json, text/javascript"},contents:{xml:/\bxml\b/,html:/\bhtml/,json:/\bjson\b/},responseFields:{xml:"responseXML",text:"responseText",json:"responseJSON"},converters:{"* text":String,"text html":!0,"text json":JSON.parse,"text xml":r.parseXML},flatOptions:{url:!0,context:!0}},ajaxSetup:function(a,b){return b?Ob(Ob(a,r.ajaxSettings),b):Ob(r.ajaxSettings,a)},ajaxPrefilter:Mb(Ib),ajaxTransport:Mb(Jb),ajax:function(b,c){"object"==typeof b&&(c=b,b=void 0),c=c||{};var e,f,g,h,i,j,k,l,m,n,o=r.ajaxSetup({},c),p=o.context||o,q=o.context&&(p.nodeType||p.jquery)?r(p):r.event,s=r.Deferred(),t=r.Callbacks("once memory"),u=o.statusCode||{},v={},w={},x="canceled",y={readyState:0,getResponseHeader:function(a){var b;if(k){if(!h){h={};while(b=Eb.exec(g))h[b[1].toLowerCase()]=b[2]}b=h[a.toLowerCase()]}return null==b?null:b},getAllResponseHeaders:function(){return k?g:null},setRequestHeader:function(a,b){return null==k&&(a=w[a.toLowerCase()]=w[a.toLowerCase()]||a,v[a]=b),this},overrideMimeType:function(a){return null==k&&(o.mimeType=a),this},statusCode:function(a){var b;if(a)if(k)y.always(a[y.status]);else for(b in a)u[b]=[u[b],a[b]];return this},abort:function(a){var b=a||x;return e&&e.abort(b),A(0,b),this}};if(s.promise(y),o.url=((b||o.url||tb.href)+"").replace(Hb,tb.protocol+"//"),o.type=c.method||c.type||o.method||o.type,o.dataTypes=(o.dataType||"*").toLowerCase().match(L)||[""],null==o.crossDomain){j=d.createElement("a");try{j.href=o.url,j.href=j.href,o.crossDomain=Lb.protocol+"//"+Lb.host!=j.protocol+"//"+j.host}catch(z){o.crossDomain=!0}}if(o.data&&o.processData&&"string"!=typeof o.data&&(o.data=r.param(o.data,o.traditional)),Nb(Ib,o,c,y),k)return y;l=r.event&&o.global,l&&0===r.active++&&r.event.trigger("ajaxStart"),o.type=o.type.toUpperCase(),o.hasContent=!Gb.test(o.type),f=o.url.replace(Cb,""),o.hasContent?o.data&&o.processData&&0===(o.contentType||"").indexOf("application/x-www-form-urlencoded")&&(o.data=o.data.replace(Bb,"+")):(n=o.url.slice(f.length),o.data&&(f+=(vb.test(f)?"&":"?")+o.data,delete o.data),o.cache===!1&&(f=f.replace(Db,"$1"),n=(vb.test(f)?"&":"?")+"_="+ub++ +n),o.url=f+n),o.ifModified&&(r.lastModified[f]&&y.setRequestHeader("If-Modified-Since",r.lastModified[f]),r.etag[f]&&y.setRequestHeader("If-None-Match",r.etag[f])),(o.data&&o.hasContent&&o.contentType!==!1||c.contentType)&&y.setRequestHeader("Content-Type",o.contentType),y.setRequestHeader("Accept",o.dataTypes[0]&&o.accepts[o.dataTypes[0]]?o.accepts[o.dataTypes[0]]+("*"!==o.dataTypes[0]?", "+Kb+"; q=0.01":""):o.accepts["*"]);for(m in o.headers)y.setRequestHeader(m,o.headers[m]);if(o.beforeSend&&(o.beforeSend.call(p,y,o)===!1||k))return y.abort();if(x="abort",t.add(o.complete),y.done(o.success),y.fail(o.error),e=Nb(Jb,o,c,y)){if(y.readyState=1,l&&q.trigger("ajaxSend",[y,o]),k)return y;o.async&&o.timeout>0&&(i=a.setTimeout(function(){y.abort("timeout")},o.timeout));try{k=!1,e.send(v,A)}catch(z){if(k)throw z;A(-1,z)}}else A(-1,"No Transport");function A(b,c,d,h){var j,m,n,v,w,x=c;k||(k=!0,i&&a.clearTimeout(i),e=void 0,g=h||"",y.readyState=b>0?4:0,j=b>=200&&b<300||304===b,d&&(v=Pb(o,y,d)),v=Qb(o,v,y,j),j?(o.ifModified&&(w=y.getResponseHeader("Last-Modified"),w&&(r.lastModified[f]=w),w=y.getResponseHeader("etag"),w&&(r.etag[f]=w)),204===b||"HEAD"===o.type?x="nocontent":304===b?x="notmodified":(x=v.state,m=v.data,n=v.error,j=!n)):(n=x,!b&&x||(x="error",b<0&&(b=0))),y.status=b,y.statusText=(c||x)+"",j?s.resolveWith(p,[m,x,y]):s.rejectWith(p,[y,x,n]),y.statusCode(u),u=void 0,l&&q.trigger(j?"ajaxSuccess":"ajaxError",[y,o,j?m:n]),t.fireWith(p,[y,x]),l&&(q.trigger("ajaxComplete",[y,o]),--r.active||r.event.trigger("ajaxStop")))}return y},getJSON:function(a,b,c){return r.get(a,b,c,"json")},getScript:function(a,b){return r.get(a,void 0,b,"script")}}),r.each(["get","post"],function(a,b){r[b]=function(a,c,d,e){return r.isFunction(c)&&(e=e||d,d=c,c=void 0),r.ajax(r.extend({url:a,type:b,dataType:e,data:c,success:d},r.isPlainObject(a)&&a))}}),r._evalUrl=function(a){return r.ajax({url:a,type:"GET",dataType:"script",cache:!0,async:!1,global:!1,"throws":!0})},r.fn.extend({wrapAll:function(a){var b;return this[0]&&(r.isFunction(a)&&(a=a.call(this[0])),b=r(a,this[0].ownerDocument).eq(0).clone(!0),this[0].parentNode&&b.insertBefore(this[0]),b.map(function(){var a=this;while(a.firstElementChild)a=a.firstElementChild;return a}).append(this)),this},wrapInner:function(a){return r.isFunction(a)?this.each(function(b){r(this).wrapInner(a.call(this,b))}):this.each(function(){var b=r(this),c=b.contents();c.length?c.wrapAll(a):b.append(a)})},wrap:function(a){var b=r.isFunction(a);return this.each(function(c){r(this).wrapAll(b?a.call(this,c):a)})},unwrap:function(a){return this.parent(a).not("body").each(function(){r(this).replaceWith(this.childNodes)}),this}}),r.expr.pseudos.hidden=function(a){return!r.expr.pseudos.visible(a)},r.expr.pseudos.visible=function(a){return!!(a.offsetWidth||a.offsetHeight||a.getClientRects().length)},r.ajaxSettings.xhr=function(){try{return new a.XMLHttpRequest}catch(b){}};var Rb={0:200,1223:204},Sb=r.ajaxSettings.xhr();o.cors=!!Sb&&"withCredentials"in Sb,o.ajax=Sb=!!Sb,r.ajaxTransport(function(b){var c,d;if(o.cors||Sb&&!b.crossDomain)return{send:function(e,f){var g,h=b.xhr();if(h.open(b.type,b.url,b.async,b.username,b.password),b.xhrFields)for(g in b.xhrFields)h[g]=b.xhrFields[g];b.mimeType&&h.overrideMimeType&&h.overrideMimeType(b.mimeType),b.crossDomain||e["X-Requested-With"]||(e["X-Requested-With"]="XMLHttpRequest");for(g in e)h.setRequestHeader(g,e[g]);c=function(a){return function(){c&&(c=d=h.onload=h.onerror=h.onabort=h.onreadystatechange=null,"abort"===a?h.abort():"error"===a?"number"!=typeof h.status?f(0,"error"):f(h.status,h.statusText):f(Rb[h.status]||h.status,h.statusText,"text"!==(h.responseType||"text")||"string"!=typeof h.responseText?{binary:h.response}:{text:h.responseText},h.getAllResponseHeaders()))}},h.onload=c(),d=h.onerror=c("error"),void 0!==h.onabort?h.onabort=d:h.onreadystatechange=function(){4===h.readyState&&a.setTimeout(function(){c&&d()})},c=c("abort");try{h.send(b.hasContent&&b.data||null)}catch(i){if(c)throw i}},abort:function(){c&&c()}}}),r.ajaxPrefilter(function(a){a.crossDomain&&(a.contents.script=!1)}),r.ajaxSetup({accepts:{script:"text/javascript, application/javascript, application/ecmascript, application/x-ecmascript"},contents:{script:/\b(?:java|ecma)script\b/},converters:{"text script":function(a){return r.globalEval(a),a}}}),r.ajaxPrefilter("script",function(a){void 0===a.cache&&(a.cache=!1),a.crossDomain&&(a.type="GET")}),r.ajaxTransport("script",function(a){if(a.crossDomain){var b,c;return{send:function(e,f){b=r("<script>").prop({charset:a.scriptCharset,src:a.url}).on("load error",c=function(a){b.remove(),c=null,a&&f("error"===a.type?404:200,a.type)}),d.head.appendChild(b[0])},abort:function(){c&&c()}}}});var Tb=[],Ub=/(=)\?(?=&|$)|\?\?/;r.ajaxSetup({jsonp:"callback",jsonpCallback:function(){var a=Tb.pop()||r.expando+"_"+ub++;return this[a]=!0,a}}),r.ajaxPrefilter("json jsonp",function(b,c,d){var e,f,g,h=b.jsonp!==!1&&(Ub.test(b.url)?"url":"string"==typeof b.data&&0===(b.contentType||"").indexOf("application/x-www-form-urlencoded")&&Ub.test(b.data)&&"data");if(h||"jsonp"===b.dataTypes[0])return e=b.jsonpCallback=r.isFunction(b.jsonpCallback)?b.jsonpCallback():b.jsonpCallback,h?b[h]=b[h].replace(Ub,"$1"+e):b.jsonp!==!1&&(b.url+=(vb.test(b.url)?"&":"?")+b.jsonp+"="+e),b.converters["script json"]=function(){return g||r.error(e+" was not called"),g[0]},b.dataTypes[0]="json",f=a[e],a[e]=function(){g=arguments},d.always(function(){void 0===f?r(a).removeProp(e):a[e]=f,b[e]&&(b.jsonpCallback=c.jsonpCallback,Tb.push(e)),g&&r.isFunction(f)&&f(g[0]),g=f=void 0}),"script"}),o.createHTMLDocument=function(){var a=d.implementation.createHTMLDocument("").body;return a.innerHTML="<form></form><form></form>",2===a.childNodes.length}(),r.parseHTML=function(a,b,c){if("string"!=typeof a)return[];"boolean"==typeof b&&(c=b,b=!1);var e,f,g;return b||(o.createHTMLDocument?(b=d.implementation.createHTMLDocument(""),e=b.createElement("base"),e.href=d.location.href,b.head.appendChild(e)):b=d),f=C.exec(a),g=!c&&[],f?[b.createElement(f[1])]:(f=qa([a],b,g),g&&g.length&&r(g).remove(),r.merge([],f.childNodes))},r.fn.load=function(a,b,c){var d,e,f,g=this,h=a.indexOf(" ");return h>-1&&(d=pb(a.slice(h)),a=a.slice(0,h)),r.isFunction(b)?(c=b,b=void 0):b&&"object"==typeof b&&(e="POST"),g.length>0&&r.ajax({url:a,type:e||"GET",dataType:"html",data:b}).done(function(a){f=arguments,g.html(d?r("<div>").append(r.parseHTML(a)).find(d):a)}).always(c&&function(a,b){g.each(function(){c.apply(this,f||[a.responseText,b,a])})}),this},r.each(["ajaxStart","ajaxStop","ajaxComplete","ajaxError","ajaxSuccess","ajaxSend"],function(a,b){r.fn[b]=function(a){return this.on(b,a)}}),r.expr.pseudos.animated=function(a){return r.grep(r.timers,function(b){return a===b.elem}).length},r.offset={setOffset:function(a,b,c){var d,e,f,g,h,i,j,k=r.css(a,"position"),l=r(a),m={};"static"===k&&(a.style.position="relative"),h=l.offset(),f=r.css(a,"top"),i=r.css(a,"left"),j=("absolute"===k||"fixed"===k)&&(f+i).indexOf("auto")>-1,j?(d=l.position(),g=d.top,e=d.left):(g=parseFloat(f)||0,e=parseFloat(i)||0),r.isFunction(b)&&(b=b.call(a,c,r.extend({},h))),null!=b.top&&(m.top=b.top-h.top+g),null!=b.left&&(m.left=b.left-h.left+e),"using"in b?b.using.call(a,m):l.css(m)}},r.fn.extend({offset:function(a){if(arguments.length)return void 0===a?this:this.each(function(b){r.offset.setOffset(this,a,b)});var b,c,d,e,f=this[0];if(f)return f.getClientRects().length?(d=f.getBoundingClientRect(),b=f.ownerDocument,c=b.documentElement,e=b.defaultView,{top:d.top+e.pageYOffset-c.clientTop,left:d.left+e.pageXOffset-c.clientLeft}):{top:0,left:0}},position:function(){if(this[0]){var a,b,c=this[0],d={top:0,left:0};return"fixed"===r.css(c,"position")?b=c.getBoundingClientRect():(a=this.offsetParent(),b=this.offset(),B(a[0],"html")||(d=a.offset()),d={top:d.top+r.css(a[0],"borderTopWidth",!0),left:d.left+r.css(a[0],"borderLeftWidth",!0)}),{top:b.top-d.top-r.css(c,"marginTop",!0),left:b.left-d.left-r.css(c,"marginLeft",!0)}}},offsetParent:function(){return this.map(function(){var a=this.offsetParent;while(a&&"static"===r.css(a,"position"))a=a.offsetParent;return a||ra})}}),r.each({scrollLeft:"pageXOffset",scrollTop:"pageYOffset"},function(a,b){var c="pageYOffset"===b;r.fn[a]=function(d){return T(this,function(a,d,e){var f;return r.isWindow(a)?f=a:9===a.nodeType&&(f=a.defaultView),void 0===e?f?f[b]:a[d]:void(f?f.scrollTo(c?f.pageXOffset:e,c?e:f.pageYOffset):a[d]=e)},a,d,arguments.length)}}),r.each(["top","left"],function(a,b){r.cssHooks[b]=Pa(o.pixelPosition,function(a,c){if(c)return c=Oa(a,b),Ma.test(c)?r(a).position()[b]+"px":c})}),r.each({Height:"height",Width:"width"},function(a,b){r.each({padding:"inner"+a,content:b,"":"outer"+a},function(c,d){r.fn[d]=function(e,f){var g=arguments.length&&(c||"boolean"!=typeof e),h=c||(e===!0||f===!0?"margin":"border");return T(this,function(b,c,e){var f;return r.isWindow(b)?0===d.indexOf("outer")?b["inner"+a]:b.document.documentElement["client"+a]:9===b.nodeType?(f=b.documentElement,Math.max(b.body["scroll"+a],f["scroll"+a],b.body["offset"+a],f["offset"+a],f["client"+a])):void 0===e?r.css(b,c,h):r.style(b,c,e,h)},b,g?e:void 0,g)}})}),r.fn.extend({bind:function(a,b,c){return this.on(a,null,b,c)},unbind:function(a,b){return this.off(a,null,b)},delegate:function(a,b,c,d){return this.on(b,a,c,d)},undelegate:function(a,b,c){return 1===arguments.length?this.off(a,"**"):this.off(b,a||"**",c)}}),r.holdReady=function(a){a?r.readyWait++:r.ready(!0)},r.isArray=Array.isArray,r.parseJSON=JSON.parse,r.nodeName=B,"function"==typeof define&&define.amd&&define("jquery",[],function(){return r});var Vb=a.jQuery,Wb=a.$;return r.noConflict=function(b){return a.$===r&&(a.$=Wb),b&&a.jQuery===r&&(a.jQuery=Vb),r},b||(a.jQuery=a.$=r),r});
`

var bootstrapJS string = `/*!
 * Bootstrap v3.3.7 (http://getbootstrap.com)
 * Copyright 2011-2016 Twitter, Inc.
 * Licensed under the MIT license
 */
if("undefined"==typeof jQuery)throw new Error("Bootstrap's JavaScript requires jQuery");+function(a){"use strict";var b=a.fn.jquery.split(" ")[0].split(".");if(b[0]<2&&b[1]<9||1==b[0]&&9==b[1]&&b[2]<1||b[0]>3)throw new Error("Bootstrap's JavaScript requires jQuery version 1.9.1 or higher, but lower than version 4")}(jQuery),+function(a){"use strict";function b(){var a=document.createElement("bootstrap"),b={WebkitTransition:"webkitTransitionEnd",MozTransition:"transitionend",OTransition:"oTransitionEnd otransitionend",transition:"transitionend"};for(var c in b)if(void 0!==a.style[c])return{end:b[c]};return!1}a.fn.emulateTransitionEnd=function(b){var c=!1,d=this;a(this).one("bsTransitionEnd",function(){c=!0});var e=function(){c||a(d).trigger(a.support.transition.end)};return setTimeout(e,b),this},a(function(){a.support.transition=b(),a.support.transition&&(a.event.special.bsTransitionEnd={bindType:a.support.transition.end,delegateType:a.support.transition.end,handle:function(b){if(a(b.target).is(this))return b.handleObj.handler.apply(this,arguments)}})})}(jQuery),+function(a){"use strict";function b(b){return this.each(function(){var c=a(this),e=c.data("bs.alert");e||c.data("bs.alert",e=new d(this)),"string"==typeof b&&e[b].call(c)})}var c='[data-dismiss="alert"]',d=function(b){a(b).on("click",c,this.close)};d.VERSION="3.3.7",d.TRANSITION_DURATION=150,d.prototype.close=function(b){function c(){g.detach().trigger("closed.bs.alert").remove()}var e=a(this),f=e.attr("data-target");f||(f=e.attr("href"),f=f&&f.replace(/.*(?=#[^\s]*$)/,""));var g=a("#"===f?[]:f);b&&b.preventDefault(),g.length||(g=e.closest(".alert")),g.trigger(b=a.Event("close.bs.alert")),b.isDefaultPrevented()||(g.removeClass("in"),a.support.transition&&g.hasClass("fade")?g.one("bsTransitionEnd",c).emulateTransitionEnd(d.TRANSITION_DURATION):c())};var e=a.fn.alert;a.fn.alert=b,a.fn.alert.Constructor=d,a.fn.alert.noConflict=function(){return a.fn.alert=e,this},a(document).on("click.bs.alert.data-api",c,d.prototype.close)}(jQuery),+function(a){"use strict";function b(b){return this.each(function(){var d=a(this),e=d.data("bs.button"),f="object"==typeof b&&b;e||d.data("bs.button",e=new c(this,f)),"toggle"==b?e.toggle():b&&e.setState(b)})}var c=function(b,d){this.$element=a(b),this.options=a.extend({},c.DEFAULTS,d),this.isLoading=!1};c.VERSION="3.3.7",c.DEFAULTS={loadingText:"loading..."},c.prototype.setState=function(b){var c="disabled",d=this.$element,e=d.is("input")?"val":"html",f=d.data();b+="Text",null==f.resetText&&d.data("resetText",d[e]()),setTimeout(a.proxy(function(){d[e](null==f[b]?this.options[b]:f[b]),"loadingText"==b?(this.isLoading=!0,d.addClass(c).attr(c,c).prop(c,!0)):this.isLoading&&(this.isLoading=!1,d.removeClass(c).removeAttr(c).prop(c,!1))},this),0)},c.prototype.toggle=function(){var a=!0,b=this.$element.closest('[data-toggle="buttons"]');if(b.length){var c=this.$element.find("input");"radio"==c.prop("type")?(c.prop("checked")&&(a=!1),b.find(".active").removeClass("active"),this.$element.addClass("active")):"checkbox"==c.prop("type")&&(c.prop("checked")!==this.$element.hasClass("active")&&(a=!1),this.$element.toggleClass("active")),c.prop("checked",this.$element.hasClass("active")),a&&c.trigger("change")}else this.$element.attr("aria-pressed",!this.$element.hasClass("active")),this.$element.toggleClass("active")};var d=a.fn.button;a.fn.button=b,a.fn.button.Constructor=c,a.fn.button.noConflict=function(){return a.fn.button=d,this},a(document).on("click.bs.button.data-api",'[data-toggle^="button"]',function(c){var d=a(c.target).closest(".btn");b.call(d,"toggle"),a(c.target).is('input[type="radio"], input[type="checkbox"]')||(c.preventDefault(),d.is("input,button")?d.trigger("focus"):d.find("input:visible,button:visible").first().trigger("focus"))}).on("focus.bs.button.data-api blur.bs.button.data-api",'[data-toggle^="button"]',function(b){a(b.target).closest(".btn").toggleClass("focus",/^focus(in)?$/.test(b.type))})}(jQuery),+function(a){"use strict";function b(b){return this.each(function(){var d=a(this),e=d.data("bs.carousel"),f=a.extend({},c.DEFAULTS,d.data(),"object"==typeof b&&b),g="string"==typeof b?b:f.slide;e||d.data("bs.carousel",e=new c(this,f)),"number"==typeof b?e.to(b):g?e[g]():f.interval&&e.pause().cycle()})}var c=function(b,c){this.$element=a(b),this.$indicators=this.$element.find(".carousel-indicators"),this.options=c,this.paused=null,this.sliding=null,this.interval=null,this.$active=null,this.$items=null,this.options.keyboard&&this.$element.on("keydown.bs.carousel",a.proxy(this.keydown,this)),"hover"==this.options.pause&&!("ontouchstart"in document.documentElement)&&this.$element.on("mouseenter.bs.carousel",a.proxy(this.pause,this)).on("mouseleave.bs.carousel",a.proxy(this.cycle,this))};c.VERSION="3.3.7",c.TRANSITION_DURATION=600,c.DEFAULTS={interval:5e3,pause:"hover",wrap:!0,keyboard:!0},c.prototype.keydown=function(a){if(!/input|textarea/i.test(a.target.tagName)){switch(a.which){case 37:this.prev();break;case 39:this.next();break;default:return}a.preventDefault()}},c.prototype.cycle=function(b){return b||(this.paused=!1),this.interval&&clearInterval(this.interval),this.options.interval&&!this.paused&&(this.interval=setInterval(a.proxy(this.next,this),this.options.interval)),this},c.prototype.getItemIndex=function(a){return this.$items=a.parent().children(".item"),this.$items.index(a||this.$active)},c.prototype.getItemForDirection=function(a,b){var c=this.getItemIndex(b),d="prev"==a&&0===c||"next"==a&&c==this.$items.length-1;if(d&&!this.options.wrap)return b;var e="prev"==a?-1:1,f=(c+e)%this.$items.length;return this.$items.eq(f)},c.prototype.to=function(a){var b=this,c=this.getItemIndex(this.$active=this.$element.find(".item.active"));if(!(a>this.$items.length-1||a<0))return this.sliding?this.$element.one("slid.bs.carousel",function(){b.to(a)}):c==a?this.pause().cycle():this.slide(a>c?"next":"prev",this.$items.eq(a))},c.prototype.pause=function(b){return b||(this.paused=!0),this.$element.find(".next, .prev").length&&a.support.transition&&(this.$element.trigger(a.support.transition.end),this.cycle(!0)),this.interval=clearInterval(this.interval),this},c.prototype.next=function(){if(!this.sliding)return this.slide("next")},c.prototype.prev=function(){if(!this.sliding)return this.slide("prev")},c.prototype.slide=function(b,d){var e=this.$element.find(".item.active"),f=d||this.getItemForDirection(b,e),g=this.interval,h="next"==b?"left":"right",i=this;if(f.hasClass("active"))return this.sliding=!1;var j=f[0],k=a.Event("slide.bs.carousel",{relatedTarget:j,direction:h});if(this.$element.trigger(k),!k.isDefaultPrevented()){if(this.sliding=!0,g&&this.pause(),this.$indicators.length){this.$indicators.find(".active").removeClass("active");var l=a(this.$indicators.children()[this.getItemIndex(f)]);l&&l.addClass("active")}var m=a.Event("slid.bs.carousel",{relatedTarget:j,direction:h});return a.support.transition&&this.$element.hasClass("slide")?(f.addClass(b),f[0].offsetWidth,e.addClass(h),f.addClass(h),e.one("bsTransitionEnd",function(){f.removeClass([b,h].join(" ")).addClass("active"),e.removeClass(["active",h].join(" ")),i.sliding=!1,setTimeout(function(){i.$element.trigger(m)},0)}).emulateTransitionEnd(c.TRANSITION_DURATION)):(e.removeClass("active"),f.addClass("active"),this.sliding=!1,this.$element.trigger(m)),g&&this.cycle(),this}};var d=a.fn.carousel;a.fn.carousel=b,a.fn.carousel.Constructor=c,a.fn.carousel.noConflict=function(){return a.fn.carousel=d,this};var e=function(c){var d,e=a(this),f=a(e.attr("data-target")||(d=e.attr("href"))&&d.replace(/.*(?=#[^\s]+$)/,""));if(f.hasClass("carousel")){var g=a.extend({},f.data(),e.data()),h=e.attr("data-slide-to");h&&(g.interval=!1),b.call(f,g),h&&f.data("bs.carousel").to(h),c.preventDefault()}};a(document).on("click.bs.carousel.data-api","[data-slide]",e).on("click.bs.carousel.data-api","[data-slide-to]",e),a(window).on("load",function(){a('[data-ride="carousel"]').each(function(){var c=a(this);b.call(c,c.data())})})}(jQuery),+function(a){"use strict";function b(b){var c,d=b.attr("data-target")||(c=b.attr("href"))&&c.replace(/.*(?=#[^\s]+$)/,"");return a(d)}function c(b){return this.each(function(){var c=a(this),e=c.data("bs.collapse"),f=a.extend({},d.DEFAULTS,c.data(),"object"==typeof b&&b);!e&&f.toggle&&/show|hide/.test(b)&&(f.toggle=!1),e||c.data("bs.collapse",e=new d(this,f)),"string"==typeof b&&e[b]()})}var d=function(b,c){this.$element=a(b),this.options=a.extend({},d.DEFAULTS,c),this.$trigger=a('[data-toggle="collapse"][href="#'+b.id+'"],[data-toggle="collapse"][data-target="#'+b.id+'"]'),this.transitioning=null,this.options.parent?this.$parent=this.getParent():this.addAriaAndCollapsedClass(this.$element,this.$trigger),this.options.toggle&&this.toggle()};d.VERSION="3.3.7",d.TRANSITION_DURATION=350,d.DEFAULTS={toggle:!0},d.prototype.dimension=function(){var a=this.$element.hasClass("width");return a?"width":"height"},d.prototype.show=function(){if(!this.transitioning&&!this.$element.hasClass("in")){var b,e=this.$parent&&this.$parent.children(".panel").children(".in, .collapsing");if(!(e&&e.length&&(b=e.data("bs.collapse"),b&&b.transitioning))){var f=a.Event("show.bs.collapse");if(this.$element.trigger(f),!f.isDefaultPrevented()){e&&e.length&&(c.call(e,"hide"),b||e.data("bs.collapse",null));var g=this.dimension();this.$element.removeClass("collapse").addClass("collapsing")[g](0).attr("aria-expanded",!0),this.$trigger.removeClass("collapsed").attr("aria-expanded",!0),this.transitioning=1;var h=function(){this.$element.removeClass("collapsing").addClass("collapse in")[g](""),this.transitioning=0,this.$element.trigger("shown.bs.collapse")};if(!a.support.transition)return h.call(this);var i=a.camelCase(["scroll",g].join("-"));this.$element.one("bsTransitionEnd",a.proxy(h,this)).emulateTransitionEnd(d.TRANSITION_DURATION)[g](this.$element[0][i])}}}},d.prototype.hide=function(){if(!this.transitioning&&this.$element.hasClass("in")){var b=a.Event("hide.bs.collapse");if(this.$element.trigger(b),!b.isDefaultPrevented()){var c=this.dimension();this.$element[c](this.$element[c]())[0].offsetHeight,this.$element.addClass("collapsing").removeClass("collapse in").attr("aria-expanded",!1),this.$trigger.addClass("collapsed").attr("aria-expanded",!1),this.transitioning=1;var e=function(){this.transitioning=0,this.$element.removeClass("collapsing").addClass("collapse").trigger("hidden.bs.collapse")};return a.support.transition?void this.$element[c](0).one("bsTransitionEnd",a.proxy(e,this)).emulateTransitionEnd(d.TRANSITION_DURATION):e.call(this)}}},d.prototype.toggle=function(){this[this.$element.hasClass("in")?"hide":"show"]()},d.prototype.getParent=function(){return a(this.options.parent).find('[data-toggle="collapse"][data-parent="'+this.options.parent+'"]').each(a.proxy(function(c,d){var e=a(d);this.addAriaAndCollapsedClass(b(e),e)},this)).end()},d.prototype.addAriaAndCollapsedClass=function(a,b){var c=a.hasClass("in");a.attr("aria-expanded",c),b.toggleClass("collapsed",!c).attr("aria-expanded",c)};var e=a.fn.collapse;a.fn.collapse=c,a.fn.collapse.Constructor=d,a.fn.collapse.noConflict=function(){return a.fn.collapse=e,this},a(document).on("click.bs.collapse.data-api",'[data-toggle="collapse"]',function(d){var e=a(this);e.attr("data-target")||d.preventDefault();var f=b(e),g=f.data("bs.collapse"),h=g?"toggle":e.data();c.call(f,h)})}(jQuery),+function(a){"use strict";function b(b){var c=b.attr("data-target");c||(c=b.attr("href"),c=c&&/#[A-Za-z]/.test(c)&&c.replace(/.*(?=#[^\s]*$)/,""));var d=c&&a(c);return d&&d.length?d:b.parent()}function c(c){c&&3===c.which||(a(e).remove(),a(f).each(function(){var d=a(this),e=b(d),f={relatedTarget:this};e.hasClass("open")&&(c&&"click"==c.type&&/input|textarea/i.test(c.target.tagName)&&a.contains(e[0],c.target)||(e.trigger(c=a.Event("hide.bs.dropdown",f)),c.isDefaultPrevented()||(d.attr("aria-expanded","false"),e.removeClass("open").trigger(a.Event("hidden.bs.dropdown",f)))))}))}function d(b){return this.each(function(){var c=a(this),d=c.data("bs.dropdown");d||c.data("bs.dropdown",d=new g(this)),"string"==typeof b&&d[b].call(c)})}var e=".dropdown-backdrop",f='[data-toggle="dropdown"]',g=function(b){a(b).on("click.bs.dropdown",this.toggle)};g.VERSION="3.3.7",g.prototype.toggle=function(d){var e=a(this);if(!e.is(".disabled, :disabled")){var f=b(e),g=f.hasClass("open");if(c(),!g){"ontouchstart"in document.documentElement&&!f.closest(".navbar-nav").length&&a(document.createElement("div")).addClass("dropdown-backdrop").insertAfter(a(this)).on("click",c);var h={relatedTarget:this};if(f.trigger(d=a.Event("show.bs.dropdown",h)),d.isDefaultPrevented())return;e.trigger("focus").attr("aria-expanded","true"),f.toggleClass("open").trigger(a.Event("shown.bs.dropdown",h))}return!1}},g.prototype.keydown=function(c){if(/(38|40|27|32)/.test(c.which)&&!/input|textarea/i.test(c.target.tagName)){var d=a(this);if(c.preventDefault(),c.stopPropagation(),!d.is(".disabled, :disabled")){var e=b(d),g=e.hasClass("open");if(!g&&27!=c.which||g&&27==c.which)return 27==c.which&&e.find(f).trigger("focus"),d.trigger("click");var h=" li:not(.disabled):visible a",i=e.find(".dropdown-menu"+h);if(i.length){var j=i.index(c.target);38==c.which&&j>0&&j--,40==c.which&&j<i.length-1&&j++,~j||(j=0),i.eq(j).trigger("focus")}}}};var h=a.fn.dropdown;a.fn.dropdown=d,a.fn.dropdown.Constructor=g,a.fn.dropdown.noConflict=function(){return a.fn.dropdown=h,this},a(document).on("click.bs.dropdown.data-api",c).on("click.bs.dropdown.data-api",".dropdown form",function(a){a.stopPropagation()}).on("click.bs.dropdown.data-api",f,g.prototype.toggle).on("keydown.bs.dropdown.data-api",f,g.prototype.keydown).on("keydown.bs.dropdown.data-api",".dropdown-menu",g.prototype.keydown)}(jQuery),+function(a){"use strict";function b(b,d){return this.each(function(){var e=a(this),f=e.data("bs.modal"),g=a.extend({},c.DEFAULTS,e.data(),"object"==typeof b&&b);f||e.data("bs.modal",f=new c(this,g)),"string"==typeof b?f[b](d):g.show&&f.show(d)})}var c=function(b,c){this.options=c,this.$body=a(document.body),this.$element=a(b),this.$dialog=this.$element.find(".modal-dialog"),this.$backdrop=null,this.isShown=null,this.originalBodyPad=null,this.scrollbarWidth=0,this.ignoreBackdropClick=!1,this.options.remote&&this.$element.find(".modal-content").load(this.options.remote,a.proxy(function(){this.$element.trigger("loaded.bs.modal")},this))};c.VERSION="3.3.7",c.TRANSITION_DURATION=300,c.BACKDROP_TRANSITION_DURATION=150,c.DEFAULTS={backdrop:!0,keyboard:!0,show:!0},c.prototype.toggle=function(a){return this.isShown?this.hide():this.show(a)},c.prototype.show=function(b){var d=this,e=a.Event("show.bs.modal",{relatedTarget:b});this.$element.trigger(e),this.isShown||e.isDefaultPrevented()||(this.isShown=!0,this.checkScrollbar(),this.setScrollbar(),this.$body.addClass("modal-open"),this.escape(),this.resize(),this.$element.on("click.dismiss.bs.modal",'[data-dismiss="modal"]',a.proxy(this.hide,this)),this.$dialog.on("mousedown.dismiss.bs.modal",function(){d.$element.one("mouseup.dismiss.bs.modal",function(b){a(b.target).is(d.$element)&&(d.ignoreBackdropClick=!0)})}),this.backdrop(function(){var e=a.support.transition&&d.$element.hasClass("fade");d.$element.parent().length||d.$element.appendTo(d.$body),d.$element.show().scrollTop(0),d.adjustDialog(),e&&d.$element[0].offsetWidth,d.$element.addClass("in"),d.enforceFocus();var f=a.Event("shown.bs.modal",{relatedTarget:b});e?d.$dialog.one("bsTransitionEnd",function(){d.$element.trigger("focus").trigger(f)}).emulateTransitionEnd(c.TRANSITION_DURATION):d.$element.trigger("focus").trigger(f)}))},c.prototype.hide=function(b){b&&b.preventDefault(),b=a.Event("hide.bs.modal"),this.$element.trigger(b),this.isShown&&!b.isDefaultPrevented()&&(this.isShown=!1,this.escape(),this.resize(),a(document).off("focusin.bs.modal"),this.$element.removeClass("in").off("click.dismiss.bs.modal").off("mouseup.dismiss.bs.modal"),this.$dialog.off("mousedown.dismiss.bs.modal"),a.support.transition&&this.$element.hasClass("fade")?this.$element.one("bsTransitionEnd",a.proxy(this.hideModal,this)).emulateTransitionEnd(c.TRANSITION_DURATION):this.hideModal())},c.prototype.enforceFocus=function(){a(document).off("focusin.bs.modal").on("focusin.bs.modal",a.proxy(function(a){document===a.target||this.$element[0]===a.target||this.$element.has(a.target).length||this.$element.trigger("focus")},this))},c.prototype.escape=function(){this.isShown&&this.options.keyboard?this.$element.on("keydown.dismiss.bs.modal",a.proxy(function(a){27==a.which&&this.hide()},this)):this.isShown||this.$element.off("keydown.dismiss.bs.modal")},c.prototype.resize=function(){this.isShown?a(window).on("resize.bs.modal",a.proxy(this.handleUpdate,this)):a(window).off("resize.bs.modal")},c.prototype.hideModal=function(){var a=this;this.$element.hide(),this.backdrop(function(){a.$body.removeClass("modal-open"),a.resetAdjustments(),a.resetScrollbar(),a.$element.trigger("hidden.bs.modal")})},c.prototype.removeBackdrop=function(){this.$backdrop&&this.$backdrop.remove(),this.$backdrop=null},c.prototype.backdrop=function(b){var d=this,e=this.$element.hasClass("fade")?"fade":"";if(this.isShown&&this.options.backdrop){var f=a.support.transition&&e;if(this.$backdrop=a(document.createElement("div")).addClass("modal-backdrop "+e).appendTo(this.$body),this.$element.on("click.dismiss.bs.modal",a.proxy(function(a){return this.ignoreBackdropClick?void(this.ignoreBackdropClick=!1):void(a.target===a.currentTarget&&("static"==this.options.backdrop?this.$element[0].focus():this.hide()))},this)),f&&this.$backdrop[0].offsetWidth,this.$backdrop.addClass("in"),!b)return;f?this.$backdrop.one("bsTransitionEnd",b).emulateTransitionEnd(c.BACKDROP_TRANSITION_DURATION):b()}else if(!this.isShown&&this.$backdrop){this.$backdrop.removeClass("in");var g=function(){d.removeBackdrop(),b&&b()};a.support.transition&&this.$element.hasClass("fade")?this.$backdrop.one("bsTransitionEnd",g).emulateTransitionEnd(c.BACKDROP_TRANSITION_DURATION):g()}else b&&b()},c.prototype.handleUpdate=function(){this.adjustDialog()},c.prototype.adjustDialog=function(){var a=this.$element[0].scrollHeight>document.documentElement.clientHeight;this.$element.css({paddingLeft:!this.bodyIsOverflowing&&a?this.scrollbarWidth:"",paddingRight:this.bodyIsOverflowing&&!a?this.scrollbarWidth:""})},c.prototype.resetAdjustments=function(){this.$element.css({paddingLeft:"",paddingRight:""})},c.prototype.checkScrollbar=function(){var a=window.innerWidth;if(!a){var b=document.documentElement.getBoundingClientRect();a=b.right-Math.abs(b.left)}this.bodyIsOverflowing=document.body.clientWidth<a,this.scrollbarWidth=this.measureScrollbar()},c.prototype.setScrollbar=function(){var a=parseInt(this.$body.css("padding-right")||0,10);this.originalBodyPad=document.body.style.paddingRight||"",this.bodyIsOverflowing&&this.$body.css("padding-right",a+this.scrollbarWidth)},c.prototype.resetScrollbar=function(){this.$body.css("padding-right",this.originalBodyPad)},c.prototype.measureScrollbar=function(){var a=document.createElement("div");a.className="modal-scrollbar-measure",this.$body.append(a);var b=a.offsetWidth-a.clientWidth;return this.$body[0].removeChild(a),b};var d=a.fn.modal;a.fn.modal=b,a.fn.modal.Constructor=c,a.fn.modal.noConflict=function(){return a.fn.modal=d,this},a(document).on("click.bs.modal.data-api",'[data-toggle="modal"]',function(c){var d=a(this),e=d.attr("href"),f=a(d.attr("data-target")||e&&e.replace(/.*(?=#[^\s]+$)/,"")),g=f.data("bs.modal")?"toggle":a.extend({remote:!/#/.test(e)&&e},f.data(),d.data());d.is("a")&&c.preventDefault(),f.one("show.bs.modal",function(a){a.isDefaultPrevented()||f.one("hidden.bs.modal",function(){d.is(":visible")&&d.trigger("focus")})}),b.call(f,g,this)})}(jQuery),+function(a){"use strict";function b(b){return this.each(function(){var d=a(this),e=d.data("bs.tooltip"),f="object"==typeof b&&b;!e&&/destroy|hide/.test(b)||(e||d.data("bs.tooltip",e=new c(this,f)),"string"==typeof b&&e[b]())})}var c=function(a,b){this.type=null,this.options=null,this.enabled=null,this.timeout=null,this.hoverState=null,this.$element=null,this.inState=null,this.init("tooltip",a,b)};c.VERSION="3.3.7",c.TRANSITION_DURATION=150,c.DEFAULTS={animation:!0,placement:"top",selector:!1,template:'<div class="tooltip" role="tooltip"><div class="tooltip-arrow"></div><div class="tooltip-inner"></div></div>',trigger:"hover focus",title:"",delay:0,html:!1,container:!1,viewport:{selector:"body",padding:0}},c.prototype.init=function(b,c,d){if(this.enabled=!0,this.type=b,this.$element=a(c),this.options=this.getOptions(d),this.$viewport=this.options.viewport&&a(a.isFunction(this.options.viewport)?this.options.viewport.call(this,this.$element):this.options.viewport.selector||this.options.viewport),this.inState={click:!1,hover:!1,focus:!1},this.$element[0]instanceof document.constructor&&!this.options.selector)throw new Error("\"selector\" option must be specified when initializing "+this.type+" on the window.document object!");for(var e=this.options.trigger.split(" "),f=e.length;f--;){var g=e[f];if("click"==g)this.$element.on("click."+this.type,this.options.selector,a.proxy(this.toggle,this));else if("manual"!=g){var h="hover"==g?"mouseenter":"focusin",i="hover"==g?"mouseleave":"focusout";this.$element.on(h+"."+this.type,this.options.selector,a.proxy(this.enter,this)),this.$element.on(i+"."+this.type,this.options.selector,a.proxy(this.leave,this))}}this.options.selector?this._options=a.extend({},this.options,{trigger:"manual",selector:""}):this.fixTitle()},c.prototype.getDefaults=function(){return c.DEFAULTS},c.prototype.getOptions=function(b){return b=a.extend({},this.getDefaults(),this.$element.data(),b),b.delay&&"number"==typeof b.delay&&(b.delay={show:b.delay,hide:b.delay}),b},c.prototype.getDelegateOptions=function(){var b={},c=this.getDefaults();return this._options&&a.each(this._options,function(a,d){c[a]!=d&&(b[a]=d)}),b},c.prototype.enter=function(b){var c=b instanceof this.constructor?b:a(b.currentTarget).data("bs."+this.type);return c||(c=new this.constructor(b.currentTarget,this.getDelegateOptions()),a(b.currentTarget).data("bs."+this.type,c)),b instanceof a.Event&&(c.inState["focusin"==b.type?"focus":"hover"]=!0),c.tip().hasClass("in")||"in"==c.hoverState?void(c.hoverState="in"):(clearTimeout(c.timeout),c.hoverState="in",c.options.delay&&c.options.delay.show?void(c.timeout=setTimeout(function(){"in"==c.hoverState&&c.show()},c.options.delay.show)):c.show())},c.prototype.isInStateTrue=function(){for(var a in this.inState)if(this.inState[a])return!0;return!1},c.prototype.leave=function(b){var c=b instanceof this.constructor?b:a(b.currentTarget).data("bs."+this.type);if(c||(c=new this.constructor(b.currentTarget,this.getDelegateOptions()),a(b.currentTarget).data("bs."+this.type,c)),b instanceof a.Event&&(c.inState["focusout"==b.type?"focus":"hover"]=!1),!c.isInStateTrue())return clearTimeout(c.timeout),c.hoverState="out",c.options.delay&&c.options.delay.hide?void(c.timeout=setTimeout(function(){"out"==c.hoverState&&c.hide()},c.options.delay.hide)):c.hide()},c.prototype.show=function(){var b=a.Event("show.bs."+this.type);if(this.hasContent()&&this.enabled){this.$element.trigger(b);var d=a.contains(this.$element[0].ownerDocument.documentElement,this.$element[0]);if(b.isDefaultPrevented()||!d)return;var e=this,f=this.tip(),g=this.getUID(this.type);this.setContent(),f.attr("id",g),this.$element.attr("aria-describedby",g),this.options.animation&&f.addClass("fade");var h="function"==typeof this.options.placement?this.options.placement.call(this,f[0],this.$element[0]):this.options.placement,i=/\s?auto?\s?/i,j=i.test(h);j&&(h=h.replace(i,"")||"top"),f.detach().css({top:0,left:0,display:"block"}).addClass(h).data("bs."+this.type,this),this.options.container?f.appendTo(this.options.container):f.insertAfter(this.$element),this.$element.trigger("inserted.bs."+this.type);var k=this.getPosition(),l=f[0].offsetWidth,m=f[0].offsetHeight;if(j){var n=h,o=this.getPosition(this.$viewport);h="bottom"==h&&k.bottom+m>o.bottom?"top":"top"==h&&k.top-m<o.top?"bottom":"right"==h&&k.right+l>o.width?"left":"left"==h&&k.left-l<o.left?"right":h,f.removeClass(n).addClass(h)}var p=this.getCalculatedOffset(h,k,l,m);this.applyPlacement(p,h);var q=function(){var a=e.hoverState;e.$element.trigger("shown.bs."+e.type),e.hoverState=null,"out"==a&&e.leave(e)};a.support.transition&&this.$tip.hasClass("fade")?f.one("bsTransitionEnd",q).emulateTransitionEnd(c.TRANSITION_DURATION):q()}},c.prototype.applyPlacement=function(b,c){var d=this.tip(),e=d[0].offsetWidth,f=d[0].offsetHeight,g=parseInt(d.css("margin-top"),10),h=parseInt(d.css("margin-left"),10);isNaN(g)&&(g=0),isNaN(h)&&(h=0),b.top+=g,b.left+=h,a.offset.setOffset(d[0],a.extend({using:function(a){d.css({top:Math.round(a.top),left:Math.round(a.left)})}},b),0),d.addClass("in");var i=d[0].offsetWidth,j=d[0].offsetHeight;"top"==c&&j!=f&&(b.top=b.top+f-j);var k=this.getViewportAdjustedDelta(c,b,i,j);k.left?b.left+=k.left:b.top+=k.top;var l=/top|bottom/.test(c),m=l?2*k.left-e+i:2*k.top-f+j,n=l?"offsetWidth":"offsetHeight";d.offset(b),this.replaceArrow(m,d[0][n],l)},c.prototype.replaceArrow=function(a,b,c){this.arrow().css(c?"left":"top",50*(1-a/b)+"%").css(c?"top":"left","")},c.prototype.setContent=function(){var a=this.tip(),b=this.getTitle();a.find(".tooltip-inner")[this.options.html?"html":"text"](b),a.removeClass("fade in top bottom left right")},c.prototype.hide=function(b){function d(){"in"!=e.hoverState&&f.detach(),e.$element&&e.$element.removeAttr("aria-describedby").trigger("hidden.bs."+e.type),b&&b()}var e=this,f=a(this.$tip),g=a.Event("hide.bs."+this.type);if(this.$element.trigger(g),!g.isDefaultPrevented())return f.removeClass("in"),a.support.transition&&f.hasClass("fade")?f.one("bsTransitionEnd",d).emulateTransitionEnd(c.TRANSITION_DURATION):d(),this.hoverState=null,this},c.prototype.fixTitle=function(){var a=this.$element;(a.attr("title")||"string"!=typeof a.attr("data-original-title"))&&a.attr("data-original-title",a.attr("title")||"").attr("title","")},c.prototype.hasContent=function(){return this.getTitle()},c.prototype.getPosition=function(b){b=b||this.$element;var c=b[0],d="BODY"==c.tagName,e=c.getBoundingClientRect();null==e.width&&(e=a.extend({},e,{width:e.right-e.left,height:e.bottom-e.top}));var f=window.SVGElement&&c instanceof window.SVGElement,g=d?{top:0,left:0}:f?null:b.offset(),h={scroll:d?document.documentElement.scrollTop||document.body.scrollTop:b.scrollTop()},i=d?{width:a(window).width(),height:a(window).height()}:null;return a.extend({},e,h,i,g)},c.prototype.getCalculatedOffset=function(a,b,c,d){return"bottom"==a?{top:b.top+b.height,left:b.left+b.width/2-c/2}:"top"==a?{top:b.top-d,left:b.left+b.width/2-c/2}:"left"==a?{top:b.top+b.height/2-d/2,left:b.left-c}:{top:b.top+b.height/2-d/2,left:b.left+b.width}},c.prototype.getViewportAdjustedDelta=function(a,b,c,d){var e={top:0,left:0};if(!this.$viewport)return e;var f=this.options.viewport&&this.options.viewport.padding||0,g=this.getPosition(this.$viewport);if(/right|left/.test(a)){var h=b.top-f-g.scroll,i=b.top+f-g.scroll+d;h<g.top?e.top=g.top-h:i>g.top+g.height&&(e.top=g.top+g.height-i)}else{var j=b.left-f,k=b.left+f+c;j<g.left?e.left=g.left-j:k>g.right&&(e.left=g.left+g.width-k)}return e},c.prototype.getTitle=function(){var a,b=this.$element,c=this.options;return a=b.attr("data-original-title")||("function"==typeof c.title?c.title.call(b[0]):c.title)},c.prototype.getUID=function(a){do a+=~~(1e6*Math.random());while(document.getElementById(a));return a},c.prototype.tip=function(){if(!this.$tip&&(this.$tip=a(this.options.template),1!=this.$tip.length))throw new Error(this.type+" \"template\" option must consist of exactly 1 top-level element!");return this.$tip},c.prototype.arrow=function(){return this.$arrow=this.$arrow||this.tip().find(".tooltip-arrow")},c.prototype.enable=function(){this.enabled=!0},c.prototype.disable=function(){this.enabled=!1},c.prototype.toggleEnabled=function(){this.enabled=!this.enabled},c.prototype.toggle=function(b){var c=this;b&&(c=a(b.currentTarget).data("bs."+this.type),c||(c=new this.constructor(b.currentTarget,this.getDelegateOptions()),a(b.currentTarget).data("bs."+this.type,c))),b?(c.inState.click=!c.inState.click,c.isInStateTrue()?c.enter(c):c.leave(c)):c.tip().hasClass("in")?c.leave(c):c.enter(c)},c.prototype.destroy=function(){var a=this;clearTimeout(this.timeout),this.hide(function(){a.$element.off("."+a.type).removeData("bs."+a.type),a.$tip&&a.$tip.detach(),a.$tip=null,a.$arrow=null,a.$viewport=null,a.$element=null})};var d=a.fn.tooltip;a.fn.tooltip=b,a.fn.tooltip.Constructor=c,a.fn.tooltip.noConflict=function(){return a.fn.tooltip=d,this}}(jQuery),+function(a){"use strict";function b(b){return this.each(function(){var d=a(this),e=d.data("bs.popover"),f="object"==typeof b&&b;!e&&/destroy|hide/.test(b)||(e||d.data("bs.popover",e=new c(this,f)),"string"==typeof b&&e[b]())})}var c=function(a,b){this.init("popover",a,b)};if(!a.fn.tooltip)throw new Error("Popover requires tooltip.js");c.VERSION="3.3.7",c.DEFAULTS=a.extend({},a.fn.tooltip.Constructor.DEFAULTS,{placement:"right",trigger:"click",content:"",template:'<div class="popover" role="tooltip"><div class="arrow"></div><h3 class="popover-title"></h3><div class="popover-content"></div></div>'}),c.prototype=a.extend({},a.fn.tooltip.Constructor.prototype),c.prototype.constructor=c,c.prototype.getDefaults=function(){return c.DEFAULTS},c.prototype.setContent=function(){var a=this.tip(),b=this.getTitle(),c=this.getContent();a.find(".popover-title")[this.options.html?"html":"text"](b),a.find(".popover-content").children().detach().end()[this.options.html?"string"==typeof c?"html":"append":"text"](c),a.removeClass("fade top bottom left right in"),a.find(".popover-title").html()||a.find(".popover-title").hide()},c.prototype.hasContent=function(){return this.getTitle()||this.getContent()},c.prototype.getContent=function(){var a=this.$element,b=this.options;return a.attr("data-content")||("function"==typeof b.content?b.content.call(a[0]):b.content)},c.prototype.arrow=function(){return this.$arrow=this.$arrow||this.tip().find(".arrow")};var d=a.fn.popover;a.fn.popover=b,a.fn.popover.Constructor=c,a.fn.popover.noConflict=function(){return a.fn.popover=d,this}}(jQuery),+function(a){"use strict";function b(c,d){this.$body=a(document.body),this.$scrollElement=a(a(c).is(document.body)?window:c),this.options=a.extend({},b.DEFAULTS,d),this.selector=(this.options.target||"")+" .nav li > a",this.offsets=[],this.targets=[],this.activeTarget=null,this.scrollHeight=0,this.$scrollElement.on("scroll.bs.scrollspy",a.proxy(this.process,this)),this.refresh(),this.process()}function c(c){return this.each(function(){var d=a(this),e=d.data("bs.scrollspy"),f="object"==typeof c&&c;e||d.data("bs.scrollspy",e=new b(this,f)),"string"==typeof c&&e[c]()})}b.VERSION="3.3.7",b.DEFAULTS={offset:10},b.prototype.getScrollHeight=function(){return this.$scrollElement[0].scrollHeight||Math.max(this.$body[0].scrollHeight,document.documentElement.scrollHeight)},b.prototype.refresh=function(){var b=this,c="offset",d=0;this.offsets=[],this.targets=[],this.scrollHeight=this.getScrollHeight(),a.isWindow(this.$scrollElement[0])||(c="position",d=this.$scrollElement.scrollTop()),this.$body.find(this.selector).map(function(){var b=a(this),e=b.data("target")||b.attr("href"),f=/^#./.test(e)&&a(e);return f&&f.length&&f.is(":visible")&&[[f[c]().top+d,e]]||null}).sort(function(a,b){return a[0]-b[0]}).each(function(){b.offsets.push(this[0]),b.targets.push(this[1])})},b.prototype.process=function(){var a,b=this.$scrollElement.scrollTop()+this.options.offset,c=this.getScrollHeight(),d=this.options.offset+c-this.$scrollElement.height(),e=this.offsets,f=this.targets,g=this.activeTarget;if(this.scrollHeight!=c&&this.refresh(),b>=d)return g!=(a=f[f.length-1])&&this.activate(a);if(g&&b<e[0])return this.activeTarget=null,this.clear();for(a=e.length;a--;)g!=f[a]&&b>=e[a]&&(void 0===e[a+1]||b<e[a+1])&&this.activate(f[a])},b.prototype.activate=function(b){
this.activeTarget=b,this.clear();var c=this.selector+'[data-target="'+b+'"],'+this.selector+'[href="'+b+'"]',d=a(c).parents("li").addClass("active");d.parent(".dropdown-menu").length&&(d=d.closest("li.dropdown").addClass("active")),d.trigger("activate.bs.scrollspy")},b.prototype.clear=function(){a(this.selector).parentsUntil(this.options.target,".active").removeClass("active")};var d=a.fn.scrollspy;a.fn.scrollspy=c,a.fn.scrollspy.Constructor=b,a.fn.scrollspy.noConflict=function(){return a.fn.scrollspy=d,this},a(window).on("load.bs.scrollspy.data-api",function(){a('[data-spy="scroll"]').each(function(){var b=a(this);c.call(b,b.data())})})}(jQuery),+function(a){"use strict";function b(b){return this.each(function(){var d=a(this),e=d.data("bs.tab");e||d.data("bs.tab",e=new c(this)),"string"==typeof b&&e[b]()})}var c=function(b){this.element=a(b)};c.VERSION="3.3.7",c.TRANSITION_DURATION=150,c.prototype.show=function(){var b=this.element,c=b.closest("ul:not(.dropdown-menu)"),d=b.data("target");if(d||(d=b.attr("href"),d=d&&d.replace(/.*(?=#[^\s]*$)/,"")),!b.parent("li").hasClass("active")){var e=c.find(".active:last a"),f=a.Event("hide.bs.tab",{relatedTarget:b[0]}),g=a.Event("show.bs.tab",{relatedTarget:e[0]});if(e.trigger(f),b.trigger(g),!g.isDefaultPrevented()&&!f.isDefaultPrevented()){var h=a(d);this.activate(b.closest("li"),c),this.activate(h,h.parent(),function(){e.trigger({type:"hidden.bs.tab",relatedTarget:b[0]}),b.trigger({type:"shown.bs.tab",relatedTarget:e[0]})})}}},c.prototype.activate=function(b,d,e){function f(){g.removeClass("active").find("> .dropdown-menu > .active").removeClass("active").end().find('[data-toggle="tab"]').attr("aria-expanded",!1),b.addClass("active").find('[data-toggle="tab"]').attr("aria-expanded",!0),h?(b[0].offsetWidth,b.addClass("in")):b.removeClass("fade"),b.parent(".dropdown-menu").length&&b.closest("li.dropdown").addClass("active").end().find('[data-toggle="tab"]').attr("aria-expanded",!0),e&&e()}var g=d.find("> .active"),h=e&&a.support.transition&&(g.length&&g.hasClass("fade")||!!d.find("> .fade").length);g.length&&h?g.one("bsTransitionEnd",f).emulateTransitionEnd(c.TRANSITION_DURATION):f(),g.removeClass("in")};var d=a.fn.tab;a.fn.tab=b,a.fn.tab.Constructor=c,a.fn.tab.noConflict=function(){return a.fn.tab=d,this};var e=function(c){c.preventDefault(),b.call(a(this),"show")};a(document).on("click.bs.tab.data-api",'[data-toggle="tab"]',e).on("click.bs.tab.data-api",'[data-toggle="pill"]',e)}(jQuery),+function(a){"use strict";function b(b){return this.each(function(){var d=a(this),e=d.data("bs.affix"),f="object"==typeof b&&b;e||d.data("bs.affix",e=new c(this,f)),"string"==typeof b&&e[b]()})}var c=function(b,d){this.options=a.extend({},c.DEFAULTS,d),this.$target=a(this.options.target).on("scroll.bs.affix.data-api",a.proxy(this.checkPosition,this)).on("click.bs.affix.data-api",a.proxy(this.checkPositionWithEventLoop,this)),this.$element=a(b),this.affixed=null,this.unpin=null,this.pinnedOffset=null,this.checkPosition()};c.VERSION="3.3.7",c.RESET="affix affix-top affix-bottom",c.DEFAULTS={offset:0,target:window},c.prototype.getState=function(a,b,c,d){var e=this.$target.scrollTop(),f=this.$element.offset(),g=this.$target.height();if(null!=c&&"top"==this.affixed)return e<c&&"top";if("bottom"==this.affixed)return null!=c?!(e+this.unpin<=f.top)&&"bottom":!(e+g<=a-d)&&"bottom";var h=null==this.affixed,i=h?e:f.top,j=h?g:b;return null!=c&&e<=c?"top":null!=d&&i+j>=a-d&&"bottom"},c.prototype.getPinnedOffset=function(){if(this.pinnedOffset)return this.pinnedOffset;this.$element.removeClass(c.RESET).addClass("affix");var a=this.$target.scrollTop(),b=this.$element.offset();return this.pinnedOffset=b.top-a},c.prototype.checkPositionWithEventLoop=function(){setTimeout(a.proxy(this.checkPosition,this),1)},c.prototype.checkPosition=function(){if(this.$element.is(":visible")){var b=this.$element.height(),d=this.options.offset,e=d.top,f=d.bottom,g=Math.max(a(document).height(),a(document.body).height());"object"!=typeof d&&(f=e=d),"function"==typeof e&&(e=d.top(this.$element)),"function"==typeof f&&(f=d.bottom(this.$element));var h=this.getState(g,b,e,f);if(this.affixed!=h){null!=this.unpin&&this.$element.css("top","");var i="affix"+(h?"-"+h:""),j=a.Event(i+".bs.affix");if(this.$element.trigger(j),j.isDefaultPrevented())return;this.affixed=h,this.unpin="bottom"==h?this.getPinnedOffset():null,this.$element.removeClass(c.RESET).addClass(i).trigger(i.replace("affix","affixed")+".bs.affix")}"bottom"==h&&this.$element.offset({top:g-b-f})}};var d=a.fn.affix;a.fn.affix=b,a.fn.affix.Constructor=c,a.fn.affix.noConflict=function(){return a.fn.affix=d,this},a(window).on("load",function(){a('[data-spy="affix"]').each(function(){var c=a(this),d=c.data();d.offset=d.offset||{},null!=d.offsetBottom&&(d.offset.bottom=d.offsetBottom),null!=d.offsetTop&&(d.offset.top=d.offsetTop),b.call(c,d)})})}(jQuery);`

var boostrapCSS string = `/*!
 * Bootstrap v3.3.7 (http://getbootstrap.com)
 * Copyright 2011-2016 Twitter, Inc.
 * Licensed under MIT (https://github.com/twbs/bootstrap/blob/master/LICENSE)
 *//*! normalize.css v3.0.3 | MIT License | github.com/necolas/normalize.css */html{font-family:sans-serif;-webkit-text-size-adjust:100%;-ms-text-size-adjust:100%}body{margin:0}article,aside,details,figcaption,figure,footer,header,hgroup,main,menu,nav,section,summary{display:block}audio,canvas,progress,video{display:inline-block;vertical-align:baseline}audio:not([controls]){display:none;height:0}[hidden],template{display:none}a{background-color:transparent}a:active,a:hover{outline:0}abbr[title]{border-bottom:1px dotted}b,strong{font-weight:700}dfn{font-style:italic}h1{margin:.67em 0;font-size:2em}mark{color:#000;background:#ff0}small{font-size:80%}sub,sup{position:relative;font-size:75%;line-height:0;vertical-align:baseline}sup{top:-.5em}sub{bottom:-.25em}img{border:0}svg:not(:root){overflow:hidden}figure{margin:1em 40px}hr{height:0;-webkit-box-sizing:content-box;-moz-box-sizing:content-box;box-sizing:content-box}pre{overflow:auto}code,kbd,pre,samp{font-family:monospace,monospace;font-size:1em}button,input,optgroup,select,textarea{margin:0;font:inherit;color:inherit}button{overflow:visible}button,select{text-transform:none}button,html input[type=button],input[type=reset],input[type=submit]{-webkit-appearance:button;cursor:pointer}button[disabled],html input[disabled]{cursor:default}button::-moz-focus-inner,input::-moz-focus-inner{padding:0;border:0}input{line-height:normal}input[type=checkbox],input[type=radio]{-webkit-box-sizing:border-box;-moz-box-sizing:border-box;box-sizing:border-box;padding:0}input[type=number]::-webkit-inner-spin-button,input[type=number]::-webkit-outer-spin-button{height:auto}input[type=search]{-webkit-box-sizing:content-box;-moz-box-sizing:content-box;box-sizing:content-box;-webkit-appearance:textfield}input[type=search]::-webkit-search-cancel-button,input[type=search]::-webkit-search-decoration{-webkit-appearance:none}fieldset{padding:.35em .625em .75em;margin:0 2px;border:1px solid silver}legend{padding:0;border:0}textarea{overflow:auto}optgroup{font-weight:700}table{border-spacing:0;border-collapse:collapse}td,th{padding:0}/*! Source: https://github.com/h5bp/html5-boilerplate/blob/master/src/css/main.css */@media print{*,:after,:before{color:#000!important;text-shadow:none!important;background:0 0!important;-webkit-box-shadow:none!important;box-shadow:none!important}a,a:visited{text-decoration:underline}a[href]:after{content:" (" attr(href) ")"}abbr[title]:after{content:" (" attr(title) ")"}a[href^="javascript:"]:after,a[href^="#"]:after{content:""}blockquote,pre{border:1px solid #999;page-break-inside:avoid}thead{display:table-header-group}img,tr{page-break-inside:avoid}img{max-width:100%!important}h2,h3,p{orphans:3;widows:3}h2,h3{page-break-after:avoid}.navbar{display:none}.btn>.caret,.dropup>.btn>.caret{border-top-color:#000!important}.label{border:1px solid #000}.table{border-collapse:collapse!important}.table td,.table th{background-color:#fff!important}.table-bordered td,.table-bordered th{border:1px solid #ddd!important}}@font-face{font-family:'Glyphicons Halflings';src:url(../fonts/glyphicons-halflings-regular.eot);src:url(../fonts/glyphicons-halflings-regular.eot?#iefix) format('embedded-opentype'),url(../fonts/glyphicons-halflings-regular.woff2) format('woff2'),url(../fonts/glyphicons-halflings-regular.woff) format('woff'),url(../fonts/glyphicons-halflings-regular.ttf) format('truetype'),url(../fonts/glyphicons-halflings-regular.svg#glyphicons_halflingsregular) format('svg')}.glyphicon{position:relative;top:1px;display:inline-block;font-family:'Glyphicons Halflings';font-style:normal;font-weight:400;line-height:1;-webkit-font-smoothing:antialiased;-moz-osx-font-smoothing:grayscale}.glyphicon-asterisk:before{content:"\002a"}.glyphicon-plus:before{content:"\002b"}.glyphicon-eur:before,.glyphicon-euro:before{content:"\20ac"}.glyphicon-minus:before{content:"\2212"}.glyphicon-cloud:before{content:"\2601"}.glyphicon-envelope:before{content:"\2709"}.glyphicon-pencil:before{content:"\270f"}.glyphicon-glass:before{content:"\e001"}.glyphicon-music:before{content:"\e002"}.glyphicon-search:before{content:"\e003"}.glyphicon-heart:before{content:"\e005"}.glyphicon-star:before{content:"\e006"}.glyphicon-star-empty:before{content:"\e007"}.glyphicon-user:before{content:"\e008"}.glyphicon-film:before{content:"\e009"}.glyphicon-th-large:before{content:"\e010"}.glyphicon-th:before{content:"\e011"}.glyphicon-th-list:before{content:"\e012"}.glyphicon-ok:before{content:"\e013"}.glyphicon-remove:before{content:"\e014"}.glyphicon-zoom-in:before{content:"\e015"}.glyphicon-zoom-out:before{content:"\e016"}.glyphicon-off:before{content:"\e017"}.glyphicon-signal:before{content:"\e018"}.glyphicon-cog:before{content:"\e019"}.glyphicon-trash:before{content:"\e020"}.glyphicon-home:before{content:"\e021"}.glyphicon-file:before{content:"\e022"}.glyphicon-time:before{content:"\e023"}.glyphicon-road:before{content:"\e024"}.glyphicon-download-alt:before{content:"\e025"}.glyphicon-download:before{content:"\e026"}.glyphicon-upload:before{content:"\e027"}.glyphicon-inbox:before{content:"\e028"}.glyphicon-play-circle:before{content:"\e029"}.glyphicon-repeat:before{content:"\e030"}.glyphicon-refresh:before{content:"\e031"}.glyphicon-list-alt:before{content:"\e032"}.glyphicon-lock:before{content:"\e033"}.glyphicon-flag:before{content:"\e034"}.glyphicon-headphones:before{content:"\e035"}.glyphicon-volume-off:before{content:"\e036"}.glyphicon-volume-down:before{content:"\e037"}.glyphicon-volume-up:before{content:"\e038"}.glyphicon-qrcode:before{content:"\e039"}.glyphicon-barcode:before{content:"\e040"}.glyphicon-tag:before{content:"\e041"}.glyphicon-tags:before{content:"\e042"}.glyphicon-book:before{content:"\e043"}.glyphicon-bookmark:before{content:"\e044"}.glyphicon-print:before{content:"\e045"}.glyphicon-camera:before{content:"\e046"}.glyphicon-font:before{content:"\e047"}.glyphicon-bold:before{content:"\e048"}.glyphicon-italic:before{content:"\e049"}.glyphicon-text-height:before{content:"\e050"}.glyphicon-text-width:before{content:"\e051"}.glyphicon-align-left:before{content:"\e052"}.glyphicon-align-center:before{content:"\e053"}.glyphicon-align-right:before{content:"\e054"}.glyphicon-align-justify:before{content:"\e055"}.glyphicon-list:before{content:"\e056"}.glyphicon-indent-left:before{content:"\e057"}.glyphicon-indent-right:before{content:"\e058"}.glyphicon-facetime-video:before{content:"\e059"}.glyphicon-picture:before{content:"\e060"}.glyphicon-map-marker:before{content:"\e062"}.glyphicon-adjust:before{content:"\e063"}.glyphicon-tint:before{content:"\e064"}.glyphicon-edit:before{content:"\e065"}.glyphicon-share:before{content:"\e066"}.glyphicon-check:before{content:"\e067"}.glyphicon-move:before{content:"\e068"}.glyphicon-step-backward:before{content:"\e069"}.glyphicon-fast-backward:before{content:"\e070"}.glyphicon-backward:before{content:"\e071"}.glyphicon-play:before{content:"\e072"}.glyphicon-pause:before{content:"\e073"}.glyphicon-stop:before{content:"\e074"}.glyphicon-forward:before{content:"\e075"}.glyphicon-fast-forward:before{content:"\e076"}.glyphicon-step-forward:before{content:"\e077"}.glyphicon-eject:before{content:"\e078"}.glyphicon-chevron-left:before{content:"\e079"}.glyphicon-chevron-right:before{content:"\e080"}.glyphicon-plus-sign:before{content:"\e081"}.glyphicon-minus-sign:before{content:"\e082"}.glyphicon-remove-sign:before{content:"\e083"}.glyphicon-ok-sign:before{content:"\e084"}.glyphicon-question-sign:before{content:"\e085"}.glyphicon-info-sign:before{content:"\e086"}.glyphicon-screenshot:before{content:"\e087"}.glyphicon-remove-circle:before{content:"\e088"}.glyphicon-ok-circle:before{content:"\e089"}.glyphicon-ban-circle:before{content:"\e090"}.glyphicon-arrow-left:before{content:"\e091"}.glyphicon-arrow-right:before{content:"\e092"}.glyphicon-arrow-up:before{content:"\e093"}.glyphicon-arrow-down:before{content:"\e094"}.glyphicon-share-alt:before{content:"\e095"}.glyphicon-resize-full:before{content:"\e096"}.glyphicon-resize-small:before{content:"\e097"}.glyphicon-exclamation-sign:before{content:"\e101"}.glyphicon-gift:before{content:"\e102"}.glyphicon-leaf:before{content:"\e103"}.glyphicon-fire:before{content:"\e104"}.glyphicon-eye-open:before{content:"\e105"}.glyphicon-eye-close:before{content:"\e106"}.glyphicon-warning-sign:before{content:"\e107"}.glyphicon-plane:before{content:"\e108"}.glyphicon-calendar:before{content:"\e109"}.glyphicon-random:before{content:"\e110"}.glyphicon-comment:before{content:"\e111"}.glyphicon-magnet:before{content:"\e112"}.glyphicon-chevron-up:before{content:"\e113"}.glyphicon-chevron-down:before{content:"\e114"}.glyphicon-retweet:before{content:"\e115"}.glyphicon-shopping-cart:before{content:"\e116"}.glyphicon-folder-close:before{content:"\e117"}.glyphicon-folder-open:before{content:"\e118"}.glyphicon-resize-vertical:before{content:"\e119"}.glyphicon-resize-horizontal:before{content:"\e120"}.glyphicon-hdd:before{content:"\e121"}.glyphicon-bullhorn:before{content:"\e122"}.glyphicon-bell:before{content:"\e123"}.glyphicon-certificate:before{content:"\e124"}.glyphicon-thumbs-up:before{content:"\e125"}.glyphicon-thumbs-down:before{content:"\e126"}.glyphicon-hand-right:before{content:"\e127"}.glyphicon-hand-left:before{content:"\e128"}.glyphicon-hand-up:before{content:"\e129"}.glyphicon-hand-down:before{content:"\e130"}.glyphicon-circle-arrow-right:before{content:"\e131"}.glyphicon-circle-arrow-left:before{content:"\e132"}.glyphicon-circle-arrow-up:before{content:"\e133"}.glyphicon-circle-arrow-down:before{content:"\e134"}.glyphicon-globe:before{content:"\e135"}.glyphicon-wrench:before{content:"\e136"}.glyphicon-tasks:before{content:"\e137"}.glyphicon-filter:before{content:"\e138"}.glyphicon-briefcase:before{content:"\e139"}.glyphicon-fullscreen:before{content:"\e140"}.glyphicon-dashboard:before{content:"\e141"}.glyphicon-paperclip:before{content:"\e142"}.glyphicon-heart-empty:before{content:"\e143"}.glyphicon-link:before{content:"\e144"}.glyphicon-phone:before{content:"\e145"}.glyphicon-pushpin:before{content:"\e146"}.glyphicon-usd:before{content:"\e148"}.glyphicon-gbp:before{content:"\e149"}.glyphicon-sort:before{content:"\e150"}.glyphicon-sort-by-alphabet:before{content:"\e151"}.glyphicon-sort-by-alphabet-alt:before{content:"\e152"}.glyphicon-sort-by-order:before{content:"\e153"}.glyphicon-sort-by-order-alt:before{content:"\e154"}.glyphicon-sort-by-attributes:before{content:"\e155"}.glyphicon-sort-by-attributes-alt:before{content:"\e156"}.glyphicon-unchecked:before{content:"\e157"}.glyphicon-expand:before{content:"\e158"}.glyphicon-collapse-down:before{content:"\e159"}.glyphicon-collapse-up:before{content:"\e160"}.glyphicon-log-in:before{content:"\e161"}.glyphicon-flash:before{content:"\e162"}.glyphicon-log-out:before{content:"\e163"}.glyphicon-new-window:before{content:"\e164"}.glyphicon-record:before{content:"\e165"}.glyphicon-save:before{content:"\e166"}.glyphicon-open:before{content:"\e167"}.glyphicon-saved:before{content:"\e168"}.glyphicon-import:before{content:"\e169"}.glyphicon-export:before{content:"\e170"}.glyphicon-send:before{content:"\e171"}.glyphicon-floppy-disk:before{content:"\e172"}.glyphicon-floppy-saved:before{content:"\e173"}.glyphicon-floppy-remove:before{content:"\e174"}.glyphicon-floppy-save:before{content:"\e175"}.glyphicon-floppy-open:before{content:"\e176"}.glyphicon-credit-card:before{content:"\e177"}.glyphicon-transfer:before{content:"\e178"}.glyphicon-cutlery:before{content:"\e179"}.glyphicon-header:before{content:"\e180"}.glyphicon-compressed:before{content:"\e181"}.glyphicon-earphone:before{content:"\e182"}.glyphicon-phone-alt:before{content:"\e183"}.glyphicon-tower:before{content:"\e184"}.glyphicon-stats:before{content:"\e185"}.glyphicon-sd-video:before{content:"\e186"}.glyphicon-hd-video:before{content:"\e187"}.glyphicon-subtitles:before{content:"\e188"}.glyphicon-sound-stereo:before{content:"\e189"}.glyphicon-sound-dolby:before{content:"\e190"}.glyphicon-sound-5-1:before{content:"\e191"}.glyphicon-sound-6-1:before{content:"\e192"}.glyphicon-sound-7-1:before{content:"\e193"}.glyphicon-copyright-mark:before{content:"\e194"}.glyphicon-registration-mark:before{content:"\e195"}.glyphicon-cloud-download:before{content:"\e197"}.glyphicon-cloud-upload:before{content:"\e198"}.glyphicon-tree-conifer:before{content:"\e199"}.glyphicon-tree-deciduous:before{content:"\e200"}.glyphicon-cd:before{content:"\e201"}.glyphicon-save-file:before{content:"\e202"}.glyphicon-open-file:before{content:"\e203"}.glyphicon-level-up:before{content:"\e204"}.glyphicon-copy:before{content:"\e205"}.glyphicon-paste:before{content:"\e206"}.glyphicon-alert:before{content:"\e209"}.glyphicon-equalizer:before{content:"\e210"}.glyphicon-king:before{content:"\e211"}.glyphicon-queen:before{content:"\e212"}.glyphicon-pawn:before{content:"\e213"}.glyphicon-bishop:before{content:"\e214"}.glyphicon-knight:before{content:"\e215"}.glyphicon-baby-formula:before{content:"\e216"}.glyphicon-tent:before{content:"\26fa"}.glyphicon-blackboard:before{content:"\e218"}.glyphicon-bed:before{content:"\e219"}.glyphicon-apple:before{content:"\f8ff"}.glyphicon-erase:before{content:"\e221"}.glyphicon-hourglass:before{content:"\231b"}.glyphicon-lamp:before{content:"\e223"}.glyphicon-duplicate:before{content:"\e224"}.glyphicon-piggy-bank:before{content:"\e225"}.glyphicon-scissors:before{content:"\e226"}.glyphicon-bitcoin:before{content:"\e227"}.glyphicon-btc:before{content:"\e227"}.glyphicon-xbt:before{content:"\e227"}.glyphicon-yen:before{content:"\00a5"}.glyphicon-jpy:before{content:"\00a5"}.glyphicon-ruble:before{content:"\20bd"}.glyphicon-rub:before{content:"\20bd"}.glyphicon-scale:before{content:"\e230"}.glyphicon-ice-lolly:before{content:"\e231"}.glyphicon-ice-lolly-tasted:before{content:"\e232"}.glyphicon-education:before{content:"\e233"}.glyphicon-option-horizontal:before{content:"\e234"}.glyphicon-option-vertical:before{content:"\e235"}.glyphicon-menu-hamburger:before{content:"\e236"}.glyphicon-modal-window:before{content:"\e237"}.glyphicon-oil:before{content:"\e238"}.glyphicon-grain:before{content:"\e239"}.glyphicon-sunglasses:before{content:"\e240"}.glyphicon-text-size:before{content:"\e241"}.glyphicon-text-color:before{content:"\e242"}.glyphicon-text-background:before{content:"\e243"}.glyphicon-object-align-top:before{content:"\e244"}.glyphicon-object-align-bottom:before{content:"\e245"}.glyphicon-object-align-horizontal:before{content:"\e246"}.glyphicon-object-align-left:before{content:"\e247"}.glyphicon-object-align-vertical:before{content:"\e248"}.glyphicon-object-align-right:before{content:"\e249"}.glyphicon-triangle-right:before{content:"\e250"}.glyphicon-triangle-left:before{content:"\e251"}.glyphicon-triangle-bottom:before{content:"\e252"}.glyphicon-triangle-top:before{content:"\e253"}.glyphicon-console:before{content:"\e254"}.glyphicon-superscript:before{content:"\e255"}.glyphicon-subscript:before{content:"\e256"}.glyphicon-menu-left:before{content:"\e257"}.glyphicon-menu-right:before{content:"\e258"}.glyphicon-menu-down:before{content:"\e259"}.glyphicon-menu-up:before{content:"\e260"}*{-webkit-box-sizing:border-box;-moz-box-sizing:border-box;box-sizing:border-box}:after,:before{-webkit-box-sizing:border-box;-moz-box-sizing:border-box;box-sizing:border-box}html{font-size:10px;-webkit-tap-highlight-color:rgba(0,0,0,0)}body{font-family:"Helvetica Neue",Helvetica,Arial,sans-serif;font-size:14px;line-height:1.42857143;color:#333;background-color:#fff}button,input,select,textarea{font-family:inherit;font-size:inherit;line-height:inherit}a{color:#337ab7;text-decoration:none}a:focus,a:hover{color:#23527c;text-decoration:underline}a:focus{outline:5px auto -webkit-focus-ring-color;outline-offset:-2px}figure{margin:0}img{vertical-align:middle}.carousel-inner>.item>a>img,.carousel-inner>.item>img,.img-responsive,.thumbnail a>img,.thumbnail>img{display:block;max-width:100%;height:auto}.img-rounded{border-radius:6px}.img-thumbnail{display:inline-block;max-width:100%;height:auto;padding:4px;line-height:1.42857143;background-color:#fff;border:1px solid #ddd;border-radius:4px;-webkit-transition:all .2s ease-in-out;-o-transition:all .2s ease-in-out;transition:all .2s ease-in-out}.img-circle{border-radius:50%}hr{margin-top:20px;margin-bottom:20px;border:0;border-top:1px solid #eee}.sr-only{position:absolute;width:1px;height:1px;padding:0;margin:-1px;overflow:hidden;clip:rect(0,0,0,0);border:0}.sr-only-focusable:active,.sr-only-focusable:focus{position:static;width:auto;height:auto;margin:0;overflow:visible;clip:auto}[role=button]{cursor:pointer}.h1,.h2,.h3,.h4,.h5,.h6,h1,h2,h3,h4,h5,h6{font-family:inherit;font-weight:500;line-height:1.1;color:inherit}.h1 .small,.h1 small,.h2 .small,.h2 small,.h3 .small,.h3 small,.h4 .small,.h4 small,.h5 .small,.h5 small,.h6 .small,.h6 small,h1 .small,h1 small,h2 .small,h2 small,h3 .small,h3 small,h4 .small,h4 small,h5 .small,h5 small,h6 .small,h6 small{font-weight:400;line-height:1;color:#777}.h1,.h2,.h3,h1,h2,h3{margin-top:20px;margin-bottom:10px}.h1 .small,.h1 small,.h2 .small,.h2 small,.h3 .small,.h3 small,h1 .small,h1 small,h2 .small,h2 small,h3 .small,h3 small{font-size:65%}.h4,.h5,.h6,h4,h5,h6{margin-top:10px;margin-bottom:10px}.h4 .small,.h4 small,.h5 .small,.h5 small,.h6 .small,.h6 small,h4 .small,h4 small,h5 .small,h5 small,h6 .small,h6 small{font-size:75%}.h1,h1{font-size:36px}.h2,h2{font-size:30px}.h3,h3{font-size:24px}.h4,h4{font-size:18px}.h5,h5{font-size:14px}.h6,h6{font-size:12px}p{margin:0 0 10px}.lead{margin-bottom:20px;font-size:16px;font-weight:300;line-height:1.4}@media (min-width:768px){.lead{font-size:21px}}.small,small{font-size:85%}.mark,mark{padding:.2em;background-color:#fcf8e3}.text-left{text-align:left}.text-right{text-align:right}.text-center{text-align:center}.text-justify{text-align:justify}.text-nowrap{white-space:nowrap}.text-lowercase{text-transform:lowercase}.text-uppercase{text-transform:uppercase}.text-capitalize{text-transform:capitalize}.text-muted{color:#777}.text-primary{color:#337ab7}a.text-primary:focus,a.text-primary:hover{color:#286090}.text-success{color:#3c763d}a.text-success:focus,a.text-success:hover{color:#2b542c}.text-info{color:#31708f}a.text-info:focus,a.text-info:hover{color:#245269}.text-warning{color:#8a6d3b}a.text-warning:focus,a.text-warning:hover{color:#66512c}.text-danger{color:#a94442}a.text-danger:focus,a.text-danger:hover{color:#843534}.bg-primary{color:#fff;background-color:#337ab7}a.bg-primary:focus,a.bg-primary:hover{background-color:#286090}.bg-success{background-color:#dff0d8}a.bg-success:focus,a.bg-success:hover{background-color:#c1e2b3}.bg-info{background-color:#d9edf7}a.bg-info:focus,a.bg-info:hover{background-color:#afd9ee}.bg-warning{background-color:#fcf8e3}a.bg-warning:focus,a.bg-warning:hover{background-color:#f7ecb5}.bg-danger{background-color:#f2dede}a.bg-danger:focus,a.bg-danger:hover{background-color:#e4b9b9}.page-header{padding-bottom:9px;margin:40px 0 20px;border-bottom:1px solid #eee}ol,ul{margin-top:0;margin-bottom:10px}ol ol,ol ul,ul ol,ul ul{margin-bottom:0}.list-unstyled{padding-left:0;list-style:none}.list-inline{padding-left:0;margin-left:-5px;list-style:none}.list-inline>li{display:inline-block;padding-right:5px;padding-left:5px}dl{margin-top:0;margin-bottom:20px}dd,dt{line-height:1.42857143}dt{font-weight:700}dd{margin-left:0}@media (min-width:768px){.dl-horizontal dt{float:left;width:160px;overflow:hidden;clear:left;text-align:right;text-overflow:ellipsis;white-space:nowrap}.dl-horizontal dd{margin-left:180px}}abbr[data-original-title],abbr[title]{cursor:help;border-bottom:1px dotted #777}.initialism{font-size:90%;text-transform:uppercase}blockquote{padding:10px 20px;margin:0 0 20px;font-size:17.5px;border-left:5px solid #eee}blockquote ol:last-child,blockquote p:last-child,blockquote ul:last-child{margin-bottom:0}blockquote .small,blockquote footer,blockquote small{display:block;font-size:80%;line-height:1.42857143;color:#777}blockquote .small:before,blockquote footer:before,blockquote small:before{content:'\2014 \00A0'}.blockquote-reverse,blockquote.pull-right{padding-right:15px;padding-left:0;text-align:right;border-right:5px solid #eee;border-left:0}.blockquote-reverse .small:before,.blockquote-reverse footer:before,.blockquote-reverse small:before,blockquote.pull-right .small:before,blockquote.pull-right footer:before,blockquote.pull-right small:before{content:''}.blockquote-reverse .small:after,.blockquote-reverse footer:after,.blockquote-reverse small:after,blockquote.pull-right .small:after,blockquote.pull-right footer:after,blockquote.pull-right small:after{content:'\00A0 \2014'}address{margin-bottom:20px;font-style:normal;line-height:1.42857143}code,kbd,pre,samp{font-family:Menlo,Monaco,Consolas,"Courier New",monospace}code{padding:2px 4px;font-size:90%;color:#c7254e;background-color:#f9f2f4;border-radius:4px}kbd{padding:2px 4px;font-size:90%;color:#fff;background-color:#333;border-radius:3px;-webkit-box-shadow:inset 0 -1px 0 rgba(0,0,0,.25);box-shadow:inset 0 -1px 0 rgba(0,0,0,.25)}kbd kbd{padding:0;font-size:100%;font-weight:700;-webkit-box-shadow:none;box-shadow:none}pre{display:block;padding:9.5px;margin:0 0 10px;font-size:13px;line-height:1.42857143;color:#333;word-break:break-all;word-wrap:break-word;background-color:#f5f5f5;border:1px solid #ccc;border-radius:4px}pre code{padding:0;font-size:inherit;color:inherit;white-space:pre-wrap;background-color:transparent;border-radius:0}.pre-scrollable{max-height:340px;overflow-y:scroll}.container{padding-right:15px;padding-left:15px;margin-right:auto;margin-left:auto}@media (min-width:768px){.container{width:750px}}@media (min-width:992px){.container{width:970px}}@media (min-width:1200px){.container{width:1170px}}.container-fluid{padding-right:15px;padding-left:15px;margin-right:auto;margin-left:auto}.row{margin-right:-15px;margin-left:-15px}.col-lg-1,.col-lg-10,.col-lg-11,.col-lg-12,.col-lg-2,.col-lg-3,.col-lg-4,.col-lg-5,.col-lg-6,.col-lg-7,.col-lg-8,.col-lg-9,.col-md-1,.col-md-10,.col-md-11,.col-md-12,.col-md-2,.col-md-3,.col-md-4,.col-md-5,.col-md-6,.col-md-7,.col-md-8,.col-md-9,.col-sm-1,.col-sm-10,.col-sm-11,.col-sm-12,.col-sm-2,.col-sm-3,.col-sm-4,.col-sm-5,.col-sm-6,.col-sm-7,.col-sm-8,.col-sm-9,.col-xs-1,.col-xs-10,.col-xs-11,.col-xs-12,.col-xs-2,.col-xs-3,.col-xs-4,.col-xs-5,.col-xs-6,.col-xs-7,.col-xs-8,.col-xs-9{position:relative;min-height:1px;padding-right:15px;padding-left:15px}.col-xs-1,.col-xs-10,.col-xs-11,.col-xs-12,.col-xs-2,.col-xs-3,.col-xs-4,.col-xs-5,.col-xs-6,.col-xs-7,.col-xs-8,.col-xs-9{float:left}.col-xs-12{width:100%}.col-xs-11{width:91.66666667%}.col-xs-10{width:83.33333333%}.col-xs-9{width:75%}.col-xs-8{width:66.66666667%}.col-xs-7{width:58.33333333%}.col-xs-6{width:50%}.col-xs-5{width:41.66666667%}.col-xs-4{width:33.33333333%}.col-xs-3{width:25%}.col-xs-2{width:16.66666667%}.col-xs-1{width:8.33333333%}.col-xs-pull-12{right:100%}.col-xs-pull-11{right:91.66666667%}.col-xs-pull-10{right:83.33333333%}.col-xs-pull-9{right:75%}.col-xs-pull-8{right:66.66666667%}.col-xs-pull-7{right:58.33333333%}.col-xs-pull-6{right:50%}.col-xs-pull-5{right:41.66666667%}.col-xs-pull-4{right:33.33333333%}.col-xs-pull-3{right:25%}.col-xs-pull-2{right:16.66666667%}.col-xs-pull-1{right:8.33333333%}.col-xs-pull-0{right:auto}.col-xs-push-12{left:100%}.col-xs-push-11{left:91.66666667%}.col-xs-push-10{left:83.33333333%}.col-xs-push-9{left:75%}.col-xs-push-8{left:66.66666667%}.col-xs-push-7{left:58.33333333%}.col-xs-push-6{left:50%}.col-xs-push-5{left:41.66666667%}.col-xs-push-4{left:33.33333333%}.col-xs-push-3{left:25%}.col-xs-push-2{left:16.66666667%}.col-xs-push-1{left:8.33333333%}.col-xs-push-0{left:auto}.col-xs-offset-12{margin-left:100%}.col-xs-offset-11{margin-left:91.66666667%}.col-xs-offset-10{margin-left:83.33333333%}.col-xs-offset-9{margin-left:75%}.col-xs-offset-8{margin-left:66.66666667%}.col-xs-offset-7{margin-left:58.33333333%}.col-xs-offset-6{margin-left:50%}.col-xs-offset-5{margin-left:41.66666667%}.col-xs-offset-4{margin-left:33.33333333%}.col-xs-offset-3{margin-left:25%}.col-xs-offset-2{margin-left:16.66666667%}.col-xs-offset-1{margin-left:8.33333333%}.col-xs-offset-0{margin-left:0}@media (min-width:768px){.col-sm-1,.col-sm-10,.col-sm-11,.col-sm-12,.col-sm-2,.col-sm-3,.col-sm-4,.col-sm-5,.col-sm-6,.col-sm-7,.col-sm-8,.col-sm-9{float:left}.col-sm-12{width:100%}.col-sm-11{width:91.66666667%}.col-sm-10{width:83.33333333%}.col-sm-9{width:75%}.col-sm-8{width:66.66666667%}.col-sm-7{width:58.33333333%}.col-sm-6{width:50%}.col-sm-5{width:41.66666667%}.col-sm-4{width:33.33333333%}.col-sm-3{width:25%}.col-sm-2{width:16.66666667%}.col-sm-1{width:8.33333333%}.col-sm-pull-12{right:100%}.col-sm-pull-11{right:91.66666667%}.col-sm-pull-10{right:83.33333333%}.col-sm-pull-9{right:75%}.col-sm-pull-8{right:66.66666667%}.col-sm-pull-7{right:58.33333333%}.col-sm-pull-6{right:50%}.col-sm-pull-5{right:41.66666667%}.col-sm-pull-4{right:33.33333333%}.col-sm-pull-3{right:25%}.col-sm-pull-2{right:16.66666667%}.col-sm-pull-1{right:8.33333333%}.col-sm-pull-0{right:auto}.col-sm-push-12{left:100%}.col-sm-push-11{left:91.66666667%}.col-sm-push-10{left:83.33333333%}.col-sm-push-9{left:75%}.col-sm-push-8{left:66.66666667%}.col-sm-push-7{left:58.33333333%}.col-sm-push-6{left:50%}.col-sm-push-5{left:41.66666667%}.col-sm-push-4{left:33.33333333%}.col-sm-push-3{left:25%}.col-sm-push-2{left:16.66666667%}.col-sm-push-1{left:8.33333333%}.col-sm-push-0{left:auto}.col-sm-offset-12{margin-left:100%}.col-sm-offset-11{margin-left:91.66666667%}.col-sm-offset-10{margin-left:83.33333333%}.col-sm-offset-9{margin-left:75%}.col-sm-offset-8{margin-left:66.66666667%}.col-sm-offset-7{margin-left:58.33333333%}.col-sm-offset-6{margin-left:50%}.col-sm-offset-5{margin-left:41.66666667%}.col-sm-offset-4{margin-left:33.33333333%}.col-sm-offset-3{margin-left:25%}.col-sm-offset-2{margin-left:16.66666667%}.col-sm-offset-1{margin-left:8.33333333%}.col-sm-offset-0{margin-left:0}}@media (min-width:992px){.col-md-1,.col-md-10,.col-md-11,.col-md-12,.col-md-2,.col-md-3,.col-md-4,.col-md-5,.col-md-6,.col-md-7,.col-md-8,.col-md-9{float:left}.col-md-12{width:100%}.col-md-11{width:91.66666667%}.col-md-10{width:83.33333333%}.col-md-9{width:75%}.col-md-8{width:66.66666667%}.col-md-7{width:58.33333333%}.col-md-6{width:50%}.col-md-5{width:41.66666667%}.col-md-4{width:33.33333333%}.col-md-3{width:25%}.col-md-2{width:16.66666667%}.col-md-1{width:8.33333333%}.col-md-pull-12{right:100%}.col-md-pull-11{right:91.66666667%}.col-md-pull-10{right:83.33333333%}.col-md-pull-9{right:75%}.col-md-pull-8{right:66.66666667%}.col-md-pull-7{right:58.33333333%}.col-md-pull-6{right:50%}.col-md-pull-5{right:41.66666667%}.col-md-pull-4{right:33.33333333%}.col-md-pull-3{right:25%}.col-md-pull-2{right:16.66666667%}.col-md-pull-1{right:8.33333333%}.col-md-pull-0{right:auto}.col-md-push-12{left:100%}.col-md-push-11{left:91.66666667%}.col-md-push-10{left:83.33333333%}.col-md-push-9{left:75%}.col-md-push-8{left:66.66666667%}.col-md-push-7{left:58.33333333%}.col-md-push-6{left:50%}.col-md-push-5{left:41.66666667%}.col-md-push-4{left:33.33333333%}.col-md-push-3{left:25%}.col-md-push-2{left:16.66666667%}.col-md-push-1{left:8.33333333%}.col-md-push-0{left:auto}.col-md-offset-12{margin-left:100%}.col-md-offset-11{margin-left:91.66666667%}.col-md-offset-10{margin-left:83.33333333%}.col-md-offset-9{margin-left:75%}.col-md-offset-8{margin-left:66.66666667%}.col-md-offset-7{margin-left:58.33333333%}.col-md-offset-6{margin-left:50%}.col-md-offset-5{margin-left:41.66666667%}.col-md-offset-4{margin-left:33.33333333%}.col-md-offset-3{margin-left:25%}.col-md-offset-2{margin-left:16.66666667%}.col-md-offset-1{margin-left:8.33333333%}.col-md-offset-0{margin-left:0}}@media (min-width:1200px){.col-lg-1,.col-lg-10,.col-lg-11,.col-lg-12,.col-lg-2,.col-lg-3,.col-lg-4,.col-lg-5,.col-lg-6,.col-lg-7,.col-lg-8,.col-lg-9{float:left}.col-lg-12{width:100%}.col-lg-11{width:91.66666667%}.col-lg-10{width:83.33333333%}.col-lg-9{width:75%}.col-lg-8{width:66.66666667%}.col-lg-7{width:58.33333333%}.col-lg-6{width:50%}.col-lg-5{width:41.66666667%}.col-lg-4{width:33.33333333%}.col-lg-3{width:25%}.col-lg-2{width:16.66666667%}.col-lg-1{width:8.33333333%}.col-lg-pull-12{right:100%}.col-lg-pull-11{right:91.66666667%}.col-lg-pull-10{right:83.33333333%}.col-lg-pull-9{right:75%}.col-lg-pull-8{right:66.66666667%}.col-lg-pull-7{right:58.33333333%}.col-lg-pull-6{right:50%}.col-lg-pull-5{right:41.66666667%}.col-lg-pull-4{right:33.33333333%}.col-lg-pull-3{right:25%}.col-lg-pull-2{right:16.66666667%}.col-lg-pull-1{right:8.33333333%}.col-lg-pull-0{right:auto}.col-lg-push-12{left:100%}.col-lg-push-11{left:91.66666667%}.col-lg-push-10{left:83.33333333%}.col-lg-push-9{left:75%}.col-lg-push-8{left:66.66666667%}.col-lg-push-7{left:58.33333333%}.col-lg-push-6{left:50%}.col-lg-push-5{left:41.66666667%}.col-lg-push-4{left:33.33333333%}.col-lg-push-3{left:25%}.col-lg-push-2{left:16.66666667%}.col-lg-push-1{left:8.33333333%}.col-lg-push-0{left:auto}.col-lg-offset-12{margin-left:100%}.col-lg-offset-11{margin-left:91.66666667%}.col-lg-offset-10{margin-left:83.33333333%}.col-lg-offset-9{margin-left:75%}.col-lg-offset-8{margin-left:66.66666667%}.col-lg-offset-7{margin-left:58.33333333%}.col-lg-offset-6{margin-left:50%}.col-lg-offset-5{margin-left:41.66666667%}.col-lg-offset-4{margin-left:33.33333333%}.col-lg-offset-3{margin-left:25%}.col-lg-offset-2{margin-left:16.66666667%}.col-lg-offset-1{margin-left:8.33333333%}.col-lg-offset-0{margin-left:0}}table{background-color:transparent}caption{padding-top:8px;padding-bottom:8px;color:#777;text-align:left}th{text-align:left}.table{width:100%;max-width:100%;margin-bottom:20px}.table>tbody>tr>td,.table>tbody>tr>th,.table>tfoot>tr>td,.table>tfoot>tr>th,.table>thead>tr>td,.table>thead>tr>th{padding:8px;line-height:1.42857143;vertical-align:top;border-top:1px solid #ddd}.table>thead>tr>th{vertical-align:bottom;border-bottom:2px solid #ddd}.table>caption+thead>tr:first-child>td,.table>caption+thead>tr:first-child>th,.table>colgroup+thead>tr:first-child>td,.table>colgroup+thead>tr:first-child>th,.table>thead:first-child>tr:first-child>td,.table>thead:first-child>tr:first-child>th{border-top:0}.table>tbody+tbody{border-top:2px solid #ddd}.table .table{background-color:#fff}.table-condensed>tbody>tr>td,.table-condensed>tbody>tr>th,.table-condensed>tfoot>tr>td,.table-condensed>tfoot>tr>th,.table-condensed>thead>tr>td,.table-condensed>thead>tr>th{padding:5px}.table-bordered{border:1px solid #ddd}.table-bordered>tbody>tr>td,.table-bordered>tbody>tr>th,.table-bordered>tfoot>tr>td,.table-bordered>tfoot>tr>th,.table-bordered>thead>tr>td,.table-bordered>thead>tr>th{border:1px solid #ddd}.table-bordered>thead>tr>td,.table-bordered>thead>tr>th{border-bottom-width:2px}.table-striped>tbody>tr:nth-of-type(odd){background-color:#f9f9f9}.table-hover>tbody>tr:hover{background-color:#f5f5f5}table col[class*=col-]{position:static;display:table-column;float:none}table td[class*=col-],table th[class*=col-]{position:static;display:table-cell;float:none}.table>tbody>tr.active>td,.table>tbody>tr.active>th,.table>tbody>tr>td.active,.table>tbody>tr>th.active,.table>tfoot>tr.active>td,.table>tfoot>tr.active>th,.table>tfoot>tr>td.active,.table>tfoot>tr>th.active,.table>thead>tr.active>td,.table>thead>tr.active>th,.table>thead>tr>td.active,.table>thead>tr>th.active{background-color:#f5f5f5}.table-hover>tbody>tr.active:hover>td,.table-hover>tbody>tr.active:hover>th,.table-hover>tbody>tr:hover>.active,.table-hover>tbody>tr>td.active:hover,.table-hover>tbody>tr>th.active:hover{background-color:#e8e8e8}.table>tbody>tr.success>td,.table>tbody>tr.success>th,.table>tbody>tr>td.success,.table>tbody>tr>th.success,.table>tfoot>tr.success>td,.table>tfoot>tr.success>th,.table>tfoot>tr>td.success,.table>tfoot>tr>th.success,.table>thead>tr.success>td,.table>thead>tr.success>th,.table>thead>tr>td.success,.table>thead>tr>th.success{background-color:#dff0d8}.table-hover>tbody>tr.success:hover>td,.table-hover>tbody>tr.success:hover>th,.table-hover>tbody>tr:hover>.success,.table-hover>tbody>tr>td.success:hover,.table-hover>tbody>tr>th.success:hover{background-color:#d0e9c6}.table>tbody>tr.info>td,.table>tbody>tr.info>th,.table>tbody>tr>td.info,.table>tbody>tr>th.info,.table>tfoot>tr.info>td,.table>tfoot>tr.info>th,.table>tfoot>tr>td.info,.table>tfoot>tr>th.info,.table>thead>tr.info>td,.table>thead>tr.info>th,.table>thead>tr>td.info,.table>thead>tr>th.info{background-color:#d9edf7}.table-hover>tbody>tr.info:hover>td,.table-hover>tbody>tr.info:hover>th,.table-hover>tbody>tr:hover>.info,.table-hover>tbody>tr>td.info:hover,.table-hover>tbody>tr>th.info:hover{background-color:#c4e3f3}.table>tbody>tr.warning>td,.table>tbody>tr.warning>th,.table>tbody>tr>td.warning,.table>tbody>tr>th.warning,.table>tfoot>tr.warning>td,.table>tfoot>tr.warning>th,.table>tfoot>tr>td.warning,.table>tfoot>tr>th.warning,.table>thead>tr.warning>td,.table>thead>tr.warning>th,.table>thead>tr>td.warning,.table>thead>tr>th.warning{background-color:#fcf8e3}.table-hover>tbody>tr.warning:hover>td,.table-hover>tbody>tr.warning:hover>th,.table-hover>tbody>tr:hover>.warning,.table-hover>tbody>tr>td.warning:hover,.table-hover>tbody>tr>th.warning:hover{background-color:#faf2cc}.table>tbody>tr.danger>td,.table>tbody>tr.danger>th,.table>tbody>tr>td.danger,.table>tbody>tr>th.danger,.table>tfoot>tr.danger>td,.table>tfoot>tr.danger>th,.table>tfoot>tr>td.danger,.table>tfoot>tr>th.danger,.table>thead>tr.danger>td,.table>thead>tr.danger>th,.table>thead>tr>td.danger,.table>thead>tr>th.danger{background-color:#f2dede}.table-hover>tbody>tr.danger:hover>td,.table-hover>tbody>tr.danger:hover>th,.table-hover>tbody>tr:hover>.danger,.table-hover>tbody>tr>td.danger:hover,.table-hover>tbody>tr>th.danger:hover{background-color:#ebcccc}.table-responsive{min-height:.01%;overflow-x:auto}@media screen and (max-width:767px){.table-responsive{width:100%;margin-bottom:15px;overflow-y:hidden;-ms-overflow-style:-ms-autohiding-scrollbar;border:1px solid #ddd}.table-responsive>.table{margin-bottom:0}.table-responsive>.table>tbody>tr>td,.table-responsive>.table>tbody>tr>th,.table-responsive>.table>tfoot>tr>td,.table-responsive>.table>tfoot>tr>th,.table-responsive>.table>thead>tr>td,.table-responsive>.table>thead>tr>th{white-space:nowrap}.table-responsive>.table-bordered{border:0}.table-responsive>.table-bordered>tbody>tr>td:first-child,.table-responsive>.table-bordered>tbody>tr>th:first-child,.table-responsive>.table-bordered>tfoot>tr>td:first-child,.table-responsive>.table-bordered>tfoot>tr>th:first-child,.table-responsive>.table-bordered>thead>tr>td:first-child,.table-responsive>.table-bordered>thead>tr>th:first-child{border-left:0}.table-responsive>.table-bordered>tbody>tr>td:last-child,.table-responsive>.table-bordered>tbody>tr>th:last-child,.table-responsive>.table-bordered>tfoot>tr>td:last-child,.table-responsive>.table-bordered>tfoot>tr>th:last-child,.table-responsive>.table-bordered>thead>tr>td:last-child,.table-responsive>.table-bordered>thead>tr>th:last-child{border-right:0}.table-responsive>.table-bordered>tbody>tr:last-child>td,.table-responsive>.table-bordered>tbody>tr:last-child>th,.table-responsive>.table-bordered>tfoot>tr:last-child>td,.table-responsive>.table-bordered>tfoot>tr:last-child>th{border-bottom:0}}fieldset{min-width:0;padding:0;margin:0;border:0}legend{display:block;width:100%;padding:0;margin-bottom:20px;font-size:21px;line-height:inherit;color:#333;border:0;border-bottom:1px solid #e5e5e5}label{display:inline-block;max-width:100%;margin-bottom:5px;font-weight:700}input[type=search]{-webkit-box-sizing:border-box;-moz-box-sizing:border-box;box-sizing:border-box}input[type=checkbox],input[type=radio]{margin:4px 0 0;margin-top:1px\9;line-height:normal}input[type=file]{display:block}input[type=range]{display:block;width:100%}select[multiple],select[size]{height:auto}input[type=file]:focus,input[type=checkbox]:focus,input[type=radio]:focus{outline:5px auto -webkit-focus-ring-color;outline-offset:-2px}output{display:block;padding-top:7px;font-size:14px;line-height:1.42857143;color:#555}.form-control{display:block;width:100%;height:34px;padding:6px 12px;font-size:14px;line-height:1.42857143;color:#555;background-color:#fff;background-image:none;border:1px solid #ccc;border-radius:4px;-webkit-box-shadow:inset 0 1px 1px rgba(0,0,0,.075);box-shadow:inset 0 1px 1px rgba(0,0,0,.075);-webkit-transition:border-color ease-in-out .15s,-webkit-box-shadow ease-in-out .15s;-o-transition:border-color ease-in-out .15s,box-shadow ease-in-out .15s;transition:border-color ease-in-out .15s,box-shadow ease-in-out .15s}.form-control:focus{border-color:#66afe9;outline:0;-webkit-box-shadow:inset 0 1px 1px rgba(0,0,0,.075),0 0 8px rgba(102,175,233,.6);box-shadow:inset 0 1px 1px rgba(0,0,0,.075),0 0 8px rgba(102,175,233,.6)}.form-control::-moz-placeholder{color:#999;opacity:1}.form-control:-ms-input-placeholder{color:#999}.form-control::-webkit-input-placeholder{color:#999}.form-control::-ms-expand{background-color:transparent;border:0}.form-control[disabled],.form-control[readonly],fieldset[disabled] .form-control{background-color:#eee;opacity:1}.form-control[disabled],fieldset[disabled] .form-control{cursor:not-allowed}textarea.form-control{height:auto}input[type=search]{-webkit-appearance:none}@media screen and (-webkit-min-device-pixel-ratio:0){input[type=date].form-control,input[type=time].form-control,input[type=datetime-local].form-control,input[type=month].form-control{line-height:34px}.input-group-sm input[type=date],.input-group-sm input[type=time],.input-group-sm input[type=datetime-local],.input-group-sm input[type=month],input[type=date].input-sm,input[type=time].input-sm,input[type=datetime-local].input-sm,input[type=month].input-sm{line-height:30px}.input-group-lg input[type=date],.input-group-lg input[type=time],.input-group-lg input[type=datetime-local],.input-group-lg input[type=month],input[type=date].input-lg,input[type=time].input-lg,input[type=datetime-local].input-lg,input[type=month].input-lg{line-height:46px}}.form-group{margin-bottom:15px}.checkbox,.radio{position:relative;display:block;margin-top:10px;margin-bottom:10px}.checkbox label,.radio label{min-height:20px;padding-left:20px;margin-bottom:0;font-weight:400;cursor:pointer}.checkbox input[type=checkbox],.checkbox-inline input[type=checkbox],.radio input[type=radio],.radio-inline input[type=radio]{position:absolute;margin-top:4px\9;margin-left:-20px}.checkbox+.checkbox,.radio+.radio{margin-top:-5px}.checkbox-inline,.radio-inline{position:relative;display:inline-block;padding-left:20px;margin-bottom:0;font-weight:400;vertical-align:middle;cursor:pointer}.checkbox-inline+.checkbox-inline,.radio-inline+.radio-inline{margin-top:0;margin-left:10px}fieldset[disabled] input[type=checkbox],fieldset[disabled] input[type=radio],input[type=checkbox].disabled,input[type=checkbox][disabled],input[type=radio].disabled,input[type=radio][disabled]{cursor:not-allowed}.checkbox-inline.disabled,.radio-inline.disabled,fieldset[disabled] .checkbox-inline,fieldset[disabled] .radio-inline{cursor:not-allowed}.checkbox.disabled label,.radio.disabled label,fieldset[disabled] .checkbox label,fieldset[disabled] .radio label{cursor:not-allowed}.form-control-static{min-height:34px;padding-top:7px;padding-bottom:7px;margin-bottom:0}.form-control-static.input-lg,.form-control-static.input-sm{padding-right:0;padding-left:0}.input-sm{height:30px;padding:5px 10px;font-size:12px;line-height:1.5;border-radius:3px}select.input-sm{height:30px;line-height:30px}select[multiple].input-sm,textarea.input-sm{height:auto}.form-group-sm .form-control{height:30px;padding:5px 10px;font-size:12px;line-height:1.5;border-radius:3px}.form-group-sm select.form-control{height:30px;line-height:30px}.form-group-sm select[multiple].form-control,.form-group-sm textarea.form-control{height:auto}.form-group-sm .form-control-static{height:30px;min-height:32px;padding:6px 10px;font-size:12px;line-height:1.5}.input-lg{height:46px;padding:10px 16px;font-size:18px;line-height:1.3333333;border-radius:6px}select.input-lg{height:46px;line-height:46px}select[multiple].input-lg,textarea.input-lg{height:auto}.form-group-lg .form-control{height:46px;padding:10px 16px;font-size:18px;line-height:1.3333333;border-radius:6px}.form-group-lg select.form-control{height:46px;line-height:46px}.form-group-lg select[multiple].form-control,.form-group-lg textarea.form-control{height:auto}.form-group-lg .form-control-static{height:46px;min-height:38px;padding:11px 16px;font-size:18px;line-height:1.3333333}.has-feedback{position:relative}.has-feedback .form-control{padding-right:42.5px}.form-control-feedback{position:absolute;top:0;right:0;z-index:2;display:block;width:34px;height:34px;line-height:34px;text-align:center;pointer-events:none}.form-group-lg .form-control+.form-control-feedback,.input-group-lg+.form-control-feedback,.input-lg+.form-control-feedback{width:46px;height:46px;line-height:46px}.form-group-sm .form-control+.form-control-feedback,.input-group-sm+.form-control-feedback,.input-sm+.form-control-feedback{width:30px;height:30px;line-height:30px}.has-success .checkbox,.has-success .checkbox-inline,.has-success .control-label,.has-success .help-block,.has-success .radio,.has-success .radio-inline,.has-success.checkbox label,.has-success.checkbox-inline label,.has-success.radio label,.has-success.radio-inline label{color:#3c763d}.has-success .form-control{border-color:#3c763d;-webkit-box-shadow:inset 0 1px 1px rgba(0,0,0,.075);box-shadow:inset 0 1px 1px rgba(0,0,0,.075)}.has-success .form-control:focus{border-color:#2b542c;-webkit-box-shadow:inset 0 1px 1px rgba(0,0,0,.075),0 0 6px #67b168;box-shadow:inset 0 1px 1px rgba(0,0,0,.075),0 0 6px #67b168}.has-success .input-group-addon{color:#3c763d;background-color:#dff0d8;border-color:#3c763d}.has-success .form-control-feedback{color:#3c763d}.has-warning .checkbox,.has-warning .checkbox-inline,.has-warning .control-label,.has-warning .help-block,.has-warning .radio,.has-warning .radio-inline,.has-warning.checkbox label,.has-warning.checkbox-inline label,.has-warning.radio label,.has-warning.radio-inline label{color:#8a6d3b}.has-warning .form-control{border-color:#8a6d3b;-webkit-box-shadow:inset 0 1px 1px rgba(0,0,0,.075);box-shadow:inset 0 1px 1px rgba(0,0,0,.075)}.has-warning .form-control:focus{border-color:#66512c;-webkit-box-shadow:inset 0 1px 1px rgba(0,0,0,.075),0 0 6px #c0a16b;box-shadow:inset 0 1px 1px rgba(0,0,0,.075),0 0 6px #c0a16b}.has-warning .input-group-addon{color:#8a6d3b;background-color:#fcf8e3;border-color:#8a6d3b}.has-warning .form-control-feedback{color:#8a6d3b}.has-error .checkbox,.has-error .checkbox-inline,.has-error .control-label,.has-error .help-block,.has-error .radio,.has-error .radio-inline,.has-error.checkbox label,.has-error.checkbox-inline label,.has-error.radio label,.has-error.radio-inline label{color:#a94442}.has-error .form-control{border-color:#a94442;-webkit-box-shadow:inset 0 1px 1px rgba(0,0,0,.075);box-shadow:inset 0 1px 1px rgba(0,0,0,.075)}.has-error .form-control:focus{border-color:#843534;-webkit-box-shadow:inset 0 1px 1px rgba(0,0,0,.075),0 0 6px #ce8483;box-shadow:inset 0 1px 1px rgba(0,0,0,.075),0 0 6px #ce8483}.has-error .input-group-addon{color:#a94442;background-color:#f2dede;border-color:#a94442}.has-error .form-control-feedback{color:#a94442}.has-feedback label~.form-control-feedback{top:25px}.has-feedback label.sr-only~.form-control-feedback{top:0}.help-block{display:block;margin-top:5px;margin-bottom:10px;color:#737373}@media (min-width:768px){.form-inline .form-group{display:inline-block;margin-bottom:0;vertical-align:middle}.form-inline .form-control{display:inline-block;width:auto;vertical-align:middle}.form-inline .form-control-static{display:inline-block}.form-inline .input-group{display:inline-table;vertical-align:middle}.form-inline .input-group .form-control,.form-inline .input-group .input-group-addon,.form-inline .input-group .input-group-btn{width:auto}.form-inline .input-group>.form-control{width:100%}.form-inline .control-label{margin-bottom:0;vertical-align:middle}.form-inline .checkbox,.form-inline .radio{display:inline-block;margin-top:0;margin-bottom:0;vertical-align:middle}.form-inline .checkbox label,.form-inline .radio label{padding-left:0}.form-inline .checkbox input[type=checkbox],.form-inline .radio input[type=radio]{position:relative;margin-left:0}.form-inline .has-feedback .form-control-feedback{top:0}}.form-horizontal .checkbox,.form-horizontal .checkbox-inline,.form-horizontal .radio,.form-horizontal .radio-inline{padding-top:7px;margin-top:0;margin-bottom:0}.form-horizontal .checkbox,.form-horizontal .radio{min-height:27px}.form-horizontal .form-group{margin-right:-15px;margin-left:-15px}@media (min-width:768px){.form-horizontal .control-label{padding-top:7px;margin-bottom:0;text-align:right}}.form-horizontal .has-feedback .form-control-feedback{right:15px}@media (min-width:768px){.form-horizontal .form-group-lg .control-label{padding-top:11px;font-size:18px}}@media (min-width:768px){.form-horizontal .form-group-sm .control-label{padding-top:6px;font-size:12px}}.btn{display:inline-block;padding:6px 12px;margin-bottom:0;font-size:14px;font-weight:400;line-height:1.42857143;text-align:center;white-space:nowrap;vertical-align:middle;-ms-touch-action:manipulation;touch-action:manipulation;cursor:pointer;-webkit-user-select:none;-moz-user-select:none;-ms-user-select:none;user-select:none;background-image:none;border:1px solid transparent;border-radius:4px}.btn.active.focus,.btn.active:focus,.btn.focus,.btn:active.focus,.btn:active:focus,.btn:focus{outline:5px auto -webkit-focus-ring-color;outline-offset:-2px}.btn.focus,.btn:focus,.btn:hover{color:#333;text-decoration:none}.btn.active,.btn:active{background-image:none;outline:0;-webkit-box-shadow:inset 0 3px 5px rgba(0,0,0,.125);box-shadow:inset 0 3px 5px rgba(0,0,0,.125)}.btn.disabled,.btn[disabled],fieldset[disabled] .btn{cursor:not-allowed;filter:alpha(opacity=65);-webkit-box-shadow:none;box-shadow:none;opacity:.65}a.btn.disabled,fieldset[disabled] a.btn{pointer-events:none}.btn-default{color:#333;background-color:#fff;border-color:#ccc}.btn-default.focus,.btn-default:focus{color:#333;background-color:#e6e6e6;border-color:#8c8c8c}.btn-default:hover{color:#333;background-color:#e6e6e6;border-color:#adadad}.btn-default.active,.btn-default:active,.open>.dropdown-toggle.btn-default{color:#333;background-color:#e6e6e6;border-color:#adadad}.btn-default.active.focus,.btn-default.active:focus,.btn-default.active:hover,.btn-default:active.focus,.btn-default:active:focus,.btn-default:active:hover,.open>.dropdown-toggle.btn-default.focus,.open>.dropdown-toggle.btn-default:focus,.open>.dropdown-toggle.btn-default:hover{color:#333;background-color:#d4d4d4;border-color:#8c8c8c}.btn-default.active,.btn-default:active,.open>.dropdown-toggle.btn-default{background-image:none}.btn-default.disabled.focus,.btn-default.disabled:focus,.btn-default.disabled:hover,.btn-default[disabled].focus,.btn-default[disabled]:focus,.btn-default[disabled]:hover,fieldset[disabled] .btn-default.focus,fieldset[disabled] .btn-default:focus,fieldset[disabled] .btn-default:hover{background-color:#fff;border-color:#ccc}.btn-default .badge{color:#fff;background-color:#333}.btn-primary{color:#fff;background-color:#337ab7;border-color:#2e6da4}.btn-primary.focus,.btn-primary:focus{color:#fff;background-color:#286090;border-color:#122b40}.btn-primary:hover{color:#fff;background-color:#286090;border-color:#204d74}.btn-primary.active,.btn-primary:active,.open>.dropdown-toggle.btn-primary{color:#fff;background-color:#286090;border-color:#204d74}.btn-primary.active.focus,.btn-primary.active:focus,.btn-primary.active:hover,.btn-primary:active.focus,.btn-primary:active:focus,.btn-primary:active:hover,.open>.dropdown-toggle.btn-primary.focus,.open>.dropdown-toggle.btn-primary:focus,.open>.dropdown-toggle.btn-primary:hover{color:#fff;background-color:#204d74;border-color:#122b40}.btn-primary.active,.btn-primary:active,.open>.dropdown-toggle.btn-primary{background-image:none}.btn-primary.disabled.focus,.btn-primary.disabled:focus,.btn-primary.disabled:hover,.btn-primary[disabled].focus,.btn-primary[disabled]:focus,.btn-primary[disabled]:hover,fieldset[disabled] .btn-primary.focus,fieldset[disabled] .btn-primary:focus,fieldset[disabled] .btn-primary:hover{background-color:#337ab7;border-color:#2e6da4}.btn-primary .badge{color:#337ab7;background-color:#fff}.btn-success{color:#fff;background-color:#5cb85c;border-color:#4cae4c}.btn-success.focus,.btn-success:focus{color:#fff;background-color:#449d44;border-color:#255625}.btn-success:hover{color:#fff;background-color:#449d44;border-color:#398439}.btn-success.active,.btn-success:active,.open>.dropdown-toggle.btn-success{color:#fff;background-color:#449d44;border-color:#398439}.btn-success.active.focus,.btn-success.active:focus,.btn-success.active:hover,.btn-success:active.focus,.btn-success:active:focus,.btn-success:active:hover,.open>.dropdown-toggle.btn-success.focus,.open>.dropdown-toggle.btn-success:focus,.open>.dropdown-toggle.btn-success:hover{color:#fff;background-color:#398439;border-color:#255625}.btn-success.active,.btn-success:active,.open>.dropdown-toggle.btn-success{background-image:none}.btn-success.disabled.focus,.btn-success.disabled:focus,.btn-success.disabled:hover,.btn-success[disabled].focus,.btn-success[disabled]:focus,.btn-success[disabled]:hover,fieldset[disabled] .btn-success.focus,fieldset[disabled] .btn-success:focus,fieldset[disabled] .btn-success:hover{background-color:#5cb85c;border-color:#4cae4c}.btn-success .badge{color:#5cb85c;background-color:#fff}.btn-info{color:#fff;background-color:#5bc0de;border-color:#46b8da}.btn-info.focus,.btn-info:focus{color:#fff;background-color:#31b0d5;border-color:#1b6d85}.btn-info:hover{color:#fff;background-color:#31b0d5;border-color:#269abc}.btn-info.active,.btn-info:active,.open>.dropdown-toggle.btn-info{color:#fff;background-color:#31b0d5;border-color:#269abc}.btn-info.active.focus,.btn-info.active:focus,.btn-info.active:hover,.btn-info:active.focus,.btn-info:active:focus,.btn-info:active:hover,.open>.dropdown-toggle.btn-info.focus,.open>.dropdown-toggle.btn-info:focus,.open>.dropdown-toggle.btn-info:hover{color:#fff;background-color:#269abc;border-color:#1b6d85}.btn-info.active,.btn-info:active,.open>.dropdown-toggle.btn-info{background-image:none}.btn-info.disabled.focus,.btn-info.disabled:focus,.btn-info.disabled:hover,.btn-info[disabled].focus,.btn-info[disabled]:focus,.btn-info[disabled]:hover,fieldset[disabled] .btn-info.focus,fieldset[disabled] .btn-info:focus,fieldset[disabled] .btn-info:hover{background-color:#5bc0de;border-color:#46b8da}.btn-info .badge{color:#5bc0de;background-color:#fff}.btn-warning{color:#fff;background-color:#f0ad4e;border-color:#eea236}.btn-warning.focus,.btn-warning:focus{color:#fff;background-color:#ec971f;border-color:#985f0d}.btn-warning:hover{color:#fff;background-color:#ec971f;border-color:#d58512}.btn-warning.active,.btn-warning:active,.open>.dropdown-toggle.btn-warning{color:#fff;background-color:#ec971f;border-color:#d58512}.btn-warning.active.focus,.btn-warning.active:focus,.btn-warning.active:hover,.btn-warning:active.focus,.btn-warning:active:focus,.btn-warning:active:hover,.open>.dropdown-toggle.btn-warning.focus,.open>.dropdown-toggle.btn-warning:focus,.open>.dropdown-toggle.btn-warning:hover{color:#fff;background-color:#d58512;border-color:#985f0d}.btn-warning.active,.btn-warning:active,.open>.dropdown-toggle.btn-warning{background-image:none}.btn-warning.disabled.focus,.btn-warning.disabled:focus,.btn-warning.disabled:hover,.btn-warning[disabled].focus,.btn-warning[disabled]:focus,.btn-warning[disabled]:hover,fieldset[disabled] .btn-warning.focus,fieldset[disabled] .btn-warning:focus,fieldset[disabled] .btn-warning:hover{background-color:#f0ad4e;border-color:#eea236}.btn-warning .badge{color:#f0ad4e;background-color:#fff}.btn-danger{color:#fff;background-color:#d9534f;border-color:#d43f3a}.btn-danger.focus,.btn-danger:focus{color:#fff;background-color:#c9302c;border-color:#761c19}.btn-danger:hover{color:#fff;background-color:#c9302c;border-color:#ac2925}.btn-danger.active,.btn-danger:active,.open>.dropdown-toggle.btn-danger{color:#fff;background-color:#c9302c;border-color:#ac2925}.btn-danger.active.focus,.btn-danger.active:focus,.btn-danger.active:hover,.btn-danger:active.focus,.btn-danger:active:focus,.btn-danger:active:hover,.open>.dropdown-toggle.btn-danger.focus,.open>.dropdown-toggle.btn-danger:focus,.open>.dropdown-toggle.btn-danger:hover{color:#fff;background-color:#ac2925;border-color:#761c19}.btn-danger.active,.btn-danger:active,.open>.dropdown-toggle.btn-danger{background-image:none}.btn-danger.disabled.focus,.btn-danger.disabled:focus,.btn-danger.disabled:hover,.btn-danger[disabled].focus,.btn-danger[disabled]:focus,.btn-danger[disabled]:hover,fieldset[disabled] .btn-danger.focus,fieldset[disabled] .btn-danger:focus,fieldset[disabled] .btn-danger:hover{background-color:#d9534f;border-color:#d43f3a}.btn-danger .badge{color:#d9534f;background-color:#fff}.btn-link{font-weight:400;color:#337ab7;border-radius:0}.btn-link,.btn-link.active,.btn-link:active,.btn-link[disabled],fieldset[disabled] .btn-link{background-color:transparent;-webkit-box-shadow:none;box-shadow:none}.btn-link,.btn-link:active,.btn-link:focus,.btn-link:hover{border-color:transparent}.btn-link:focus,.btn-link:hover{color:#23527c;text-decoration:underline;background-color:transparent}.btn-link[disabled]:focus,.btn-link[disabled]:hover,fieldset[disabled] .btn-link:focus,fieldset[disabled] .btn-link:hover{color:#777;text-decoration:none}.btn-group-lg>.btn,.btn-lg{padding:10px 16px;font-size:18px;line-height:1.3333333;border-radius:6px}.btn-group-sm>.btn,.btn-sm{padding:5px 10px;font-size:12px;line-height:1.5;border-radius:3px}.btn-group-xs>.btn,.btn-xs{padding:1px 5px;font-size:12px;line-height:1.5;border-radius:3px}.btn-block{display:block;width:100%}.btn-block+.btn-block{margin-top:5px}input[type=button].btn-block,input[type=reset].btn-block,input[type=submit].btn-block{width:100%}.fade{opacity:0;-webkit-transition:opacity .15s linear;-o-transition:opacity .15s linear;transition:opacity .15s linear}.fade.in{opacity:1}.collapse{display:none}.collapse.in{display:block}tr.collapse.in{display:table-row}tbody.collapse.in{display:table-row-group}.collapsing{position:relative;height:0;overflow:hidden;-webkit-transition-timing-function:ease;-o-transition-timing-function:ease;transition-timing-function:ease;-webkit-transition-duration:.35s;-o-transition-duration:.35s;transition-duration:.35s;-webkit-transition-property:height,visibility;-o-transition-property:height,visibility;transition-property:height,visibility}.caret{display:inline-block;width:0;height:0;margin-left:2px;vertical-align:middle;border-top:4px dashed;border-top:4px solid\9;border-right:4px solid transparent;border-left:4px solid transparent}.dropdown,.dropup{position:relative}.dropdown-toggle:focus{outline:0}.dropdown-menu{position:absolute;top:100%;left:0;z-index:1000;display:none;float:left;min-width:160px;padding:5px 0;margin:2px 0 0;font-size:14px;text-align:left;list-style:none;background-color:#fff;-webkit-background-clip:padding-box;background-clip:padding-box;border:1px solid #ccc;border:1px solid rgba(0,0,0,.15);border-radius:4px;-webkit-box-shadow:0 6px 12px rgba(0,0,0,.175);box-shadow:0 6px 12px rgba(0,0,0,.175)}.dropdown-menu.pull-right{right:0;left:auto}.dropdown-menu .divider{height:1px;margin:9px 0;overflow:hidden;background-color:#e5e5e5}.dropdown-menu>li>a{display:block;padding:3px 20px;clear:both;font-weight:400;line-height:1.42857143;color:#333;white-space:nowrap}.dropdown-menu>li>a:focus,.dropdown-menu>li>a:hover{color:#262626;text-decoration:none;background-color:#f5f5f5}.dropdown-menu>.active>a,.dropdown-menu>.active>a:focus,.dropdown-menu>.active>a:hover{color:#fff;text-decoration:none;background-color:#337ab7;outline:0}.dropdown-menu>.disabled>a,.dropdown-menu>.disabled>a:focus,.dropdown-menu>.disabled>a:hover{color:#777}.dropdown-menu>.disabled>a:focus,.dropdown-menu>.disabled>a:hover{text-decoration:none;cursor:not-allowed;background-color:transparent;background-image:none;filter:progid:DXImageTransform.Microsoft.gradient(enabled=false)}.open>.dropdown-menu{display:block}.open>a{outline:0}.dropdown-menu-right{right:0;left:auto}.dropdown-menu-left{right:auto;left:0}.dropdown-header{display:block;padding:3px 20px;font-size:12px;line-height:1.42857143;color:#777;white-space:nowrap}.dropdown-backdrop{position:fixed;top:0;right:0;bottom:0;left:0;z-index:990}.pull-right>.dropdown-menu{right:0;left:auto}.dropup .caret,.navbar-fixed-bottom .dropdown .caret{content:"";border-top:0;border-bottom:4px dashed;border-bottom:4px solid\9}.dropup .dropdown-menu,.navbar-fixed-bottom .dropdown .dropdown-menu{top:auto;bottom:100%;margin-bottom:2px}@media (min-width:768px){.navbar-right .dropdown-menu{right:0;left:auto}.navbar-right .dropdown-menu-left{right:auto;left:0}}.btn-group,.btn-group-vertical{position:relative;display:inline-block;vertical-align:middle}.btn-group-vertical>.btn,.btn-group>.btn{position:relative;float:left}.btn-group-vertical>.btn.active,.btn-group-vertical>.btn:active,.btn-group-vertical>.btn:focus,.btn-group-vertical>.btn:hover,.btn-group>.btn.active,.btn-group>.btn:active,.btn-group>.btn:focus,.btn-group>.btn:hover{z-index:2}.btn-group .btn+.btn,.btn-group .btn+.btn-group,.btn-group .btn-group+.btn,.btn-group .btn-group+.btn-group{margin-left:-1px}.btn-toolbar{margin-left:-5px}.btn-toolbar .btn,.btn-toolbar .btn-group,.btn-toolbar .input-group{float:left}.btn-toolbar>.btn,.btn-toolbar>.btn-group,.btn-toolbar>.input-group{margin-left:5px}.btn-group>.btn:not(:first-child):not(:last-child):not(.dropdown-toggle){border-radius:0}.btn-group>.btn:first-child{margin-left:0}.btn-group>.btn:first-child:not(:last-child):not(.dropdown-toggle){border-top-right-radius:0;border-bottom-right-radius:0}.btn-group>.btn:last-child:not(:first-child),.btn-group>.dropdown-toggle:not(:first-child){border-top-left-radius:0;border-bottom-left-radius:0}.btn-group>.btn-group{float:left}.btn-group>.btn-group:not(:first-child):not(:last-child)>.btn{border-radius:0}.btn-group>.btn-group:first-child:not(:last-child)>.btn:last-child,.btn-group>.btn-group:first-child:not(:last-child)>.dropdown-toggle{border-top-right-radius:0;border-bottom-right-radius:0}.btn-group>.btn-group:last-child:not(:first-child)>.btn:first-child{border-top-left-radius:0;border-bottom-left-radius:0}.btn-group .dropdown-toggle:active,.btn-group.open .dropdown-toggle{outline:0}.btn-group>.btn+.dropdown-toggle{padding-right:8px;padding-left:8px}.btn-group>.btn-lg+.dropdown-toggle{padding-right:12px;padding-left:12px}.btn-group.open .dropdown-toggle{-webkit-box-shadow:inset 0 3px 5px rgba(0,0,0,.125);box-shadow:inset 0 3px 5px rgba(0,0,0,.125)}.btn-group.open .dropdown-toggle.btn-link{-webkit-box-shadow:none;box-shadow:none}.btn .caret{margin-left:0}.btn-lg .caret{border-width:5px 5px 0;border-bottom-width:0}.dropup .btn-lg .caret{border-width:0 5px 5px}.btn-group-vertical>.btn,.btn-group-vertical>.btn-group,.btn-group-vertical>.btn-group>.btn{display:block;float:none;width:100%;max-width:100%}.btn-group-vertical>.btn-group>.btn{float:none}.btn-group-vertical>.btn+.btn,.btn-group-vertical>.btn+.btn-group,.btn-group-vertical>.btn-group+.btn,.btn-group-vertical>.btn-group+.btn-group{margin-top:-1px;margin-left:0}.btn-group-vertical>.btn:not(:first-child):not(:last-child){border-radius:0}.btn-group-vertical>.btn:first-child:not(:last-child){border-top-left-radius:4px;border-top-right-radius:4px;border-bottom-right-radius:0;border-bottom-left-radius:0}.btn-group-vertical>.btn:last-child:not(:first-child){border-top-left-radius:0;border-top-right-radius:0;border-bottom-right-radius:4px;border-bottom-left-radius:4px}.btn-group-vertical>.btn-group:not(:first-child):not(:last-child)>.btn{border-radius:0}.btn-group-vertical>.btn-group:first-child:not(:last-child)>.btn:last-child,.btn-group-vertical>.btn-group:first-child:not(:last-child)>.dropdown-toggle{border-bottom-right-radius:0;border-bottom-left-radius:0}.btn-group-vertical>.btn-group:last-child:not(:first-child)>.btn:first-child{border-top-left-radius:0;border-top-right-radius:0}.btn-group-justified{display:table;width:100%;table-layout:fixed;border-collapse:separate}.btn-group-justified>.btn,.btn-group-justified>.btn-group{display:table-cell;float:none;width:1%}.btn-group-justified>.btn-group .btn{width:100%}.btn-group-justified>.btn-group .dropdown-menu{left:auto}[data-toggle=buttons]>.btn input[type=checkbox],[data-toggle=buttons]>.btn input[type=radio],[data-toggle=buttons]>.btn-group>.btn input[type=checkbox],[data-toggle=buttons]>.btn-group>.btn input[type=radio]{position:absolute;clip:rect(0,0,0,0);pointer-events:none}.input-group{position:relative;display:table;border-collapse:separate}.input-group[class*=col-]{float:none;padding-right:0;padding-left:0}.input-group .form-control{position:relative;z-index:2;float:left;width:100%;margin-bottom:0}.input-group .form-control:focus{z-index:3}.input-group-lg>.form-control,.input-group-lg>.input-group-addon,.input-group-lg>.input-group-btn>.btn{height:46px;padding:10px 16px;font-size:18px;line-height:1.3333333;border-radius:6px}select.input-group-lg>.form-control,select.input-group-lg>.input-group-addon,select.input-group-lg>.input-group-btn>.btn{height:46px;line-height:46px}select[multiple].input-group-lg>.form-control,select[multiple].input-group-lg>.input-group-addon,select[multiple].input-group-lg>.input-group-btn>.btn,textarea.input-group-lg>.form-control,textarea.input-group-lg>.input-group-addon,textarea.input-group-lg>.input-group-btn>.btn{height:auto}.input-group-sm>.form-control,.input-group-sm>.input-group-addon,.input-group-sm>.input-group-btn>.btn{height:30px;padding:5px 10px;font-size:12px;line-height:1.5;border-radius:3px}select.input-group-sm>.form-control,select.input-group-sm>.input-group-addon,select.input-group-sm>.input-group-btn>.btn{height:30px;line-height:30px}select[multiple].input-group-sm>.form-control,select[multiple].input-group-sm>.input-group-addon,select[multiple].input-group-sm>.input-group-btn>.btn,textarea.input-group-sm>.form-control,textarea.input-group-sm>.input-group-addon,textarea.input-group-sm>.input-group-btn>.btn{height:auto}.input-group .form-control,.input-group-addon,.input-group-btn{display:table-cell}.input-group .form-control:not(:first-child):not(:last-child),.input-group-addon:not(:first-child):not(:last-child),.input-group-btn:not(:first-child):not(:last-child){border-radius:0}.input-group-addon,.input-group-btn{width:1%;white-space:nowrap;vertical-align:middle}.input-group-addon{padding:6px 12px;font-size:14px;font-weight:400;line-height:1;color:#555;text-align:center;background-color:#eee;border:1px solid #ccc;border-radius:4px}.input-group-addon.input-sm{padding:5px 10px;font-size:12px;border-radius:3px}.input-group-addon.input-lg{padding:10px 16px;font-size:18px;border-radius:6px}.input-group-addon input[type=checkbox],.input-group-addon input[type=radio]{margin-top:0}.input-group .form-control:first-child,.input-group-addon:first-child,.input-group-btn:first-child>.btn,.input-group-btn:first-child>.btn-group>.btn,.input-group-btn:first-child>.dropdown-toggle,.input-group-btn:last-child>.btn-group:not(:last-child)>.btn,.input-group-btn:last-child>.btn:not(:last-child):not(.dropdown-toggle){border-top-right-radius:0;border-bottom-right-radius:0}.input-group-addon:first-child{border-right:0}.input-group .form-control:last-child,.input-group-addon:last-child,.input-group-btn:first-child>.btn-group:not(:first-child)>.btn,.input-group-btn:first-child>.btn:not(:first-child),.input-group-btn:last-child>.btn,.input-group-btn:last-child>.btn-group>.btn,.input-group-btn:last-child>.dropdown-toggle{border-top-left-radius:0;border-bottom-left-radius:0}.input-group-addon:last-child{border-left:0}.input-group-btn{position:relative;font-size:0;white-space:nowrap}.input-group-btn>.btn{position:relative}.input-group-btn>.btn+.btn{margin-left:-1px}.input-group-btn>.btn:active,.input-group-btn>.btn:focus,.input-group-btn>.btn:hover{z-index:2}.input-group-btn:first-child>.btn,.input-group-btn:first-child>.btn-group{margin-right:-1px}.input-group-btn:last-child>.btn,.input-group-btn:last-child>.btn-group{z-index:2;margin-left:-1px}.nav{padding-left:0;margin-bottom:0;list-style:none}.nav>li{position:relative;display:block}.nav>li>a{position:relative;display:block;padding:10px 15px}.nav>li>a:focus,.nav>li>a:hover{text-decoration:none;background-color:#eee}.nav>li.disabled>a{color:#777}.nav>li.disabled>a:focus,.nav>li.disabled>a:hover{color:#777;text-decoration:none;cursor:not-allowed;background-color:transparent}.nav .open>a,.nav .open>a:focus,.nav .open>a:hover{background-color:#eee;border-color:#337ab7}.nav .nav-divider{height:1px;margin:9px 0;overflow:hidden;background-color:#e5e5e5}.nav>li>a>img{max-width:none}.nav-tabs{border-bottom:1px solid #ddd}.nav-tabs>li{float:left;margin-bottom:-1px}.nav-tabs>li>a{margin-right:2px;line-height:1.42857143;border:1px solid transparent;border-radius:4px 4px 0 0}.nav-tabs>li>a:hover{border-color:#eee #eee #ddd}.nav-tabs>li.active>a,.nav-tabs>li.active>a:focus,.nav-tabs>li.active>a:hover{color:#555;cursor:default;background-color:#fff;border:1px solid #ddd;border-bottom-color:transparent}.nav-tabs.nav-justified{width:100%;border-bottom:0}.nav-tabs.nav-justified>li{float:none}.nav-tabs.nav-justified>li>a{margin-bottom:5px;text-align:center}.nav-tabs.nav-justified>.dropdown .dropdown-menu{top:auto;left:auto}@media (min-width:768px){.nav-tabs.nav-justified>li{display:table-cell;width:1%}.nav-tabs.nav-justified>li>a{margin-bottom:0}}.nav-tabs.nav-justified>li>a{margin-right:0;border-radius:4px}.nav-tabs.nav-justified>.active>a,.nav-tabs.nav-justified>.active>a:focus,.nav-tabs.nav-justified>.active>a:hover{border:1px solid #ddd}@media (min-width:768px){.nav-tabs.nav-justified>li>a{border-bottom:1px solid #ddd;border-radius:4px 4px 0 0}.nav-tabs.nav-justified>.active>a,.nav-tabs.nav-justified>.active>a:focus,.nav-tabs.nav-justified>.active>a:hover{border-bottom-color:#fff}}.nav-pills>li{float:left}.nav-pills>li>a{border-radius:4px}.nav-pills>li+li{margin-left:2px}.nav-pills>li.active>a,.nav-pills>li.active>a:focus,.nav-pills>li.active>a:hover{color:#fff;background-color:#337ab7}.nav-stacked>li{float:none}.nav-stacked>li+li{margin-top:2px;margin-left:0}.nav-justified{width:100%}.nav-justified>li{float:none}.nav-justified>li>a{margin-bottom:5px;text-align:center}.nav-justified>.dropdown .dropdown-menu{top:auto;left:auto}@media (min-width:768px){.nav-justified>li{display:table-cell;width:1%}.nav-justified>li>a{margin-bottom:0}}.nav-tabs-justified{border-bottom:0}.nav-tabs-justified>li>a{margin-right:0;border-radius:4px}.nav-tabs-justified>.active>a,.nav-tabs-justified>.active>a:focus,.nav-tabs-justified>.active>a:hover{border:1px solid #ddd}@media (min-width:768px){.nav-tabs-justified>li>a{border-bottom:1px solid #ddd;border-radius:4px 4px 0 0}.nav-tabs-justified>.active>a,.nav-tabs-justified>.active>a:focus,.nav-tabs-justified>.active>a:hover{border-bottom-color:#fff}}.tab-content>.tab-pane{display:none}.tab-content>.active{display:block}.nav-tabs .dropdown-menu{margin-top:-1px;border-top-left-radius:0;border-top-right-radius:0}.navbar{position:relative;min-height:50px;margin-bottom:20px;border:1px solid transparent}@media (min-width:768px){.navbar{border-radius:4px}}@media (min-width:768px){.navbar-header{float:left}}.navbar-collapse{padding-right:15px;padding-left:15px;overflow-x:visible;-webkit-overflow-scrolling:touch;border-top:1px solid transparent;-webkit-box-shadow:inset 0 1px 0 rgba(255,255,255,.1);box-shadow:inset 0 1px 0 rgba(255,255,255,.1)}.navbar-collapse.in{overflow-y:auto}@media (min-width:768px){.navbar-collapse{width:auto;border-top:0;-webkit-box-shadow:none;box-shadow:none}.navbar-collapse.collapse{display:block!important;height:auto!important;padding-bottom:0;overflow:visible!important}.navbar-collapse.in{overflow-y:visible}.navbar-fixed-bottom .navbar-collapse,.navbar-fixed-top .navbar-collapse,.navbar-static-top .navbar-collapse{padding-right:0;padding-left:0}}.navbar-fixed-bottom .navbar-collapse,.navbar-fixed-top .navbar-collapse{max-height:340px}@media (max-device-width:480px) and (orientation:landscape){.navbar-fixed-bottom .navbar-collapse,.navbar-fixed-top .navbar-collapse{max-height:200px}}.container-fluid>.navbar-collapse,.container-fluid>.navbar-header,.container>.navbar-collapse,.container>.navbar-header{margin-right:-15px;margin-left:-15px}@media (min-width:768px){.container-fluid>.navbar-collapse,.container-fluid>.navbar-header,.container>.navbar-collapse,.container>.navbar-header{margin-right:0;margin-left:0}}.navbar-static-top{z-index:1000;border-width:0 0 1px}@media (min-width:768px){.navbar-static-top{border-radius:0}}.navbar-fixed-bottom,.navbar-fixed-top{position:fixed;right:0;left:0;z-index:1030}@media (min-width:768px){.navbar-fixed-bottom,.navbar-fixed-top{border-radius:0}}.navbar-fixed-top{top:0;border-width:0 0 1px}.navbar-fixed-bottom{bottom:0;margin-bottom:0;border-width:1px 0 0}.navbar-brand{float:left;height:50px;padding:15px 15px;font-size:18px;line-height:20px}.navbar-brand:focus,.navbar-brand:hover{text-decoration:none}.navbar-brand>img{display:block}@media (min-width:768px){.navbar>.container .navbar-brand,.navbar>.container-fluid .navbar-brand{margin-left:-15px}}.navbar-toggle{position:relative;float:right;padding:9px 10px;margin-top:8px;margin-right:15px;margin-bottom:8px;background-color:transparent;background-image:none;border:1px solid transparent;border-radius:4px}.navbar-toggle:focus{outline:0}.navbar-toggle .icon-bar{display:block;width:22px;height:2px;border-radius:1px}.navbar-toggle .icon-bar+.icon-bar{margin-top:4px}@media (min-width:768px){.navbar-toggle{display:none}}.navbar-nav{margin:7.5px -15px}.navbar-nav>li>a{padding-top:10px;padding-bottom:10px;line-height:20px}@media (max-width:767px){.navbar-nav .open .dropdown-menu{position:static;float:none;width:auto;margin-top:0;background-color:transparent;border:0;-webkit-box-shadow:none;box-shadow:none}.navbar-nav .open .dropdown-menu .dropdown-header,.navbar-nav .open .dropdown-menu>li>a{padding:5px 15px 5px 25px}.navbar-nav .open .dropdown-menu>li>a{line-height:20px}.navbar-nav .open .dropdown-menu>li>a:focus,.navbar-nav .open .dropdown-menu>li>a:hover{background-image:none}}@media (min-width:768px){.navbar-nav{float:left;margin:0}.navbar-nav>li{float:left}.navbar-nav>li>a{padding-top:15px;padding-bottom:15px}}.navbar-form{padding:10px 15px;margin-top:8px;margin-right:-15px;margin-bottom:8px;margin-left:-15px;border-top:1px solid transparent;border-bottom:1px solid transparent;-webkit-box-shadow:inset 0 1px 0 rgba(255,255,255,.1),0 1px 0 rgba(255,255,255,.1);box-shadow:inset 0 1px 0 rgba(255,255,255,.1),0 1px 0 rgba(255,255,255,.1)}@media (min-width:768px){.navbar-form .form-group{display:inline-block;margin-bottom:0;vertical-align:middle}.navbar-form .form-control{display:inline-block;width:auto;vertical-align:middle}.navbar-form .form-control-static{display:inline-block}.navbar-form .input-group{display:inline-table;vertical-align:middle}.navbar-form .input-group .form-control,.navbar-form .input-group .input-group-addon,.navbar-form .input-group .input-group-btn{width:auto}.navbar-form .input-group>.form-control{width:100%}.navbar-form .control-label{margin-bottom:0;vertical-align:middle}.navbar-form .checkbox,.navbar-form .radio{display:inline-block;margin-top:0;margin-bottom:0;vertical-align:middle}.navbar-form .checkbox label,.navbar-form .radio label{padding-left:0}.navbar-form .checkbox input[type=checkbox],.navbar-form .radio input[type=radio]{position:relative;margin-left:0}.navbar-form .has-feedback .form-control-feedback{top:0}}@media (max-width:767px){.navbar-form .form-group{margin-bottom:5px}.navbar-form .form-group:last-child{margin-bottom:0}}@media (min-width:768px){.navbar-form{width:auto;padding-top:0;padding-bottom:0;margin-right:0;margin-left:0;border:0;-webkit-box-shadow:none;box-shadow:none}}.navbar-nav>li>.dropdown-menu{margin-top:0;border-top-left-radius:0;border-top-right-radius:0}.navbar-fixed-bottom .navbar-nav>li>.dropdown-menu{margin-bottom:0;border-top-left-radius:4px;border-top-right-radius:4px;border-bottom-right-radius:0;border-bottom-left-radius:0}.navbar-btn{margin-top:8px;margin-bottom:8px}.navbar-btn.btn-sm{margin-top:10px;margin-bottom:10px}.navbar-btn.btn-xs{margin-top:14px;margin-bottom:14px}.navbar-text{margin-top:15px;margin-bottom:15px}@media (min-width:768px){.navbar-text{float:left;margin-right:15px;margin-left:15px}}@media (min-width:768px){.navbar-left{float:left!important}.navbar-right{float:right!important;margin-right:-15px}.navbar-right~.navbar-right{margin-right:0}}.navbar-default{background-color:#f8f8f8;border-color:#e7e7e7}.navbar-default .navbar-brand{color:#777}.navbar-default .navbar-brand:focus,.navbar-default .navbar-brand:hover{color:#5e5e5e;background-color:transparent}.navbar-default .navbar-text{color:#777}.navbar-default .navbar-nav>li>a{color:#777}.navbar-default .navbar-nav>li>a:focus,.navbar-default .navbar-nav>li>a:hover{color:#333;background-color:transparent}.navbar-default .navbar-nav>.active>a,.navbar-default .navbar-nav>.active>a:focus,.navbar-default .navbar-nav>.active>a:hover{color:#555;background-color:#e7e7e7}.navbar-default .navbar-nav>.disabled>a,.navbar-default .navbar-nav>.disabled>a:focus,.navbar-default .navbar-nav>.disabled>a:hover{color:#ccc;background-color:transparent}.navbar-default .navbar-toggle{border-color:#ddd}.navbar-default .navbar-toggle:focus,.navbar-default .navbar-toggle:hover{background-color:#ddd}.navbar-default .navbar-toggle .icon-bar{background-color:#888}.navbar-default .navbar-collapse,.navbar-default .navbar-form{border-color:#e7e7e7}.navbar-default .navbar-nav>.open>a,.navbar-default .navbar-nav>.open>a:focus,.navbar-default .navbar-nav>.open>a:hover{color:#555;background-color:#e7e7e7}@media (max-width:767px){.navbar-default .navbar-nav .open .dropdown-menu>li>a{color:#777}.navbar-default .navbar-nav .open .dropdown-menu>li>a:focus,.navbar-default .navbar-nav .open .dropdown-menu>li>a:hover{color:#333;background-color:transparent}.navbar-default .navbar-nav .open .dropdown-menu>.active>a,.navbar-default .navbar-nav .open .dropdown-menu>.active>a:focus,.navbar-default .navbar-nav .open .dropdown-menu>.active>a:hover{color:#555;background-color:#e7e7e7}.navbar-default .navbar-nav .open .dropdown-menu>.disabled>a,.navbar-default .navbar-nav .open .dropdown-menu>.disabled>a:focus,.navbar-default .navbar-nav .open .dropdown-menu>.disabled>a:hover{color:#ccc;background-color:transparent}}.navbar-default .navbar-link{color:#777}.navbar-default .navbar-link:hover{color:#333}.navbar-default .btn-link{color:#777}.navbar-default .btn-link:focus,.navbar-default .btn-link:hover{color:#333}.navbar-default .btn-link[disabled]:focus,.navbar-default .btn-link[disabled]:hover,fieldset[disabled] .navbar-default .btn-link:focus,fieldset[disabled] .navbar-default .btn-link:hover{color:#ccc}.navbar-inverse{background-color:#222;border-color:#080808}.navbar-inverse .navbar-brand{color:#9d9d9d}.navbar-inverse .navbar-brand:focus,.navbar-inverse .navbar-brand:hover{color:#fff;background-color:transparent}.navbar-inverse .navbar-text{color:#9d9d9d}.navbar-inverse .navbar-nav>li>a{color:#9d9d9d}.navbar-inverse .navbar-nav>li>a:focus,.navbar-inverse .navbar-nav>li>a:hover{color:#fff;background-color:transparent}.navbar-inverse .navbar-nav>.active>a,.navbar-inverse .navbar-nav>.active>a:focus,.navbar-inverse .navbar-nav>.active>a:hover{color:#fff;background-color:#080808}.navbar-inverse .navbar-nav>.disabled>a,.navbar-inverse .navbar-nav>.disabled>a:focus,.navbar-inverse .navbar-nav>.disabled>a:hover{color:#444;background-color:transparent}.navbar-inverse .navbar-toggle{border-color:#333}.navbar-inverse .navbar-toggle:focus,.navbar-inverse .navbar-toggle:hover{background-color:#333}.navbar-inverse .navbar-toggle .icon-bar{background-color:#fff}.navbar-inverse .navbar-collapse,.navbar-inverse .navbar-form{border-color:#101010}.navbar-inverse .navbar-nav>.open>a,.navbar-inverse .navbar-nav>.open>a:focus,.navbar-inverse .navbar-nav>.open>a:hover{color:#fff;background-color:#080808}@media (max-width:767px){.navbar-inverse .navbar-nav .open .dropdown-menu>.dropdown-header{border-color:#080808}.navbar-inverse .navbar-nav .open .dropdown-menu .divider{background-color:#080808}.navbar-inverse .navbar-nav .open .dropdown-menu>li>a{color:#9d9d9d}.navbar-inverse .navbar-nav .open .dropdown-menu>li>a:focus,.navbar-inverse .navbar-nav .open .dropdown-menu>li>a:hover{color:#fff;background-color:transparent}.navbar-inverse .navbar-nav .open .dropdown-menu>.active>a,.navbar-inverse .navbar-nav .open .dropdown-menu>.active>a:focus,.navbar-inverse .navbar-nav .open .dropdown-menu>.active>a:hover{color:#fff;background-color:#080808}.navbar-inverse .navbar-nav .open .dropdown-menu>.disabled>a,.navbar-inverse .navbar-nav .open .dropdown-menu>.disabled>a:focus,.navbar-inverse .navbar-nav .open .dropdown-menu>.disabled>a:hover{color:#444;background-color:transparent}}.navbar-inverse .navbar-link{color:#9d9d9d}.navbar-inverse .navbar-link:hover{color:#fff}.navbar-inverse .btn-link{color:#9d9d9d}.navbar-inverse .btn-link:focus,.navbar-inverse .btn-link:hover{color:#fff}.navbar-inverse .btn-link[disabled]:focus,.navbar-inverse .btn-link[disabled]:hover,fieldset[disabled] .navbar-inverse .btn-link:focus,fieldset[disabled] .navbar-inverse .btn-link:hover{color:#444}.breadcrumb{padding:8px 15px;margin-bottom:20px;list-style:none;background-color:#f5f5f5;border-radius:4px}.breadcrumb>li{display:inline-block}.breadcrumb>li+li:before{padding:0 5px;color:#ccc;content:"/\00a0"}.breadcrumb>.active{color:#777}.pagination{display:inline-block;padding-left:0;margin:20px 0;border-radius:4px}.pagination>li{display:inline}.pagination>li>a,.pagination>li>span{position:relative;float:left;padding:6px 12px;margin-left:-1px;line-height:1.42857143;color:#337ab7;text-decoration:none;background-color:#fff;border:1px solid #ddd}.pagination>li:first-child>a,.pagination>li:first-child>span{margin-left:0;border-top-left-radius:4px;border-bottom-left-radius:4px}.pagination>li:last-child>a,.pagination>li:last-child>span{border-top-right-radius:4px;border-bottom-right-radius:4px}.pagination>li>a:focus,.pagination>li>a:hover,.pagination>li>span:focus,.pagination>li>span:hover{z-index:2;color:#23527c;background-color:#eee;border-color:#ddd}.pagination>.active>a,.pagination>.active>a:focus,.pagination>.active>a:hover,.pagination>.active>span,.pagination>.active>span:focus,.pagination>.active>span:hover{z-index:3;color:#fff;cursor:default;background-color:#337ab7;border-color:#337ab7}.pagination>.disabled>a,.pagination>.disabled>a:focus,.pagination>.disabled>a:hover,.pagination>.disabled>span,.pagination>.disabled>span:focus,.pagination>.disabled>span:hover{color:#777;cursor:not-allowed;background-color:#fff;border-color:#ddd}.pagination-lg>li>a,.pagination-lg>li>span{padding:10px 16px;font-size:18px;line-height:1.3333333}.pagination-lg>li:first-child>a,.pagination-lg>li:first-child>span{border-top-left-radius:6px;border-bottom-left-radius:6px}.pagination-lg>li:last-child>a,.pagination-lg>li:last-child>span{border-top-right-radius:6px;border-bottom-right-radius:6px}.pagination-sm>li>a,.pagination-sm>li>span{padding:5px 10px;font-size:12px;line-height:1.5}.pagination-sm>li:first-child>a,.pagination-sm>li:first-child>span{border-top-left-radius:3px;border-bottom-left-radius:3px}.pagination-sm>li:last-child>a,.pagination-sm>li:last-child>span{border-top-right-radius:3px;border-bottom-right-radius:3px}.pager{padding-left:0;margin:20px 0;text-align:center;list-style:none}.pager li{display:inline}.pager li>a,.pager li>span{display:inline-block;padding:5px 14px;background-color:#fff;border:1px solid #ddd;border-radius:15px}.pager li>a:focus,.pager li>a:hover{text-decoration:none;background-color:#eee}.pager .next>a,.pager .next>span{float:right}.pager .previous>a,.pager .previous>span{float:left}.pager .disabled>a,.pager .disabled>a:focus,.pager .disabled>a:hover,.pager .disabled>span{color:#777;cursor:not-allowed;background-color:#fff}.label{display:inline;padding:.2em .6em .3em;font-size:75%;font-weight:700;line-height:1;color:#fff;text-align:center;white-space:nowrap;vertical-align:baseline;border-radius:.25em}a.label:focus,a.label:hover{color:#fff;text-decoration:none;cursor:pointer}.label:empty{display:none}.btn .label{position:relative;top:-1px}.label-default{background-color:#777}.label-default[href]:focus,.label-default[href]:hover{background-color:#5e5e5e}.label-primary{background-color:#337ab7}.label-primary[href]:focus,.label-primary[href]:hover{background-color:#286090}.label-success{background-color:#5cb85c}.label-success[href]:focus,.label-success[href]:hover{background-color:#449d44}.label-info{background-color:#5bc0de}.label-info[href]:focus,.label-info[href]:hover{background-color:#31b0d5}.label-warning{background-color:#f0ad4e}.label-warning[href]:focus,.label-warning[href]:hover{background-color:#ec971f}.label-danger{background-color:#d9534f}.label-danger[href]:focus,.label-danger[href]:hover{background-color:#c9302c}.badge{display:inline-block;min-width:10px;padding:3px 7px;font-size:12px;font-weight:700;line-height:1;color:#fff;text-align:center;white-space:nowrap;vertical-align:middle;background-color:#777;border-radius:10px}.badge:empty{display:none}.btn .badge{position:relative;top:-1px}.btn-group-xs>.btn .badge,.btn-xs .badge{top:0;padding:1px 5px}a.badge:focus,a.badge:hover{color:#fff;text-decoration:none;cursor:pointer}.list-group-item.active>.badge,.nav-pills>.active>a>.badge{color:#337ab7;background-color:#fff}.list-group-item>.badge{float:right}.list-group-item>.badge+.badge{margin-right:5px}.nav-pills>li>a>.badge{margin-left:3px}.jumbotron{padding-top:30px;padding-bottom:30px;margin-bottom:30px;color:inherit;background-color:#eee}.jumbotron .h1,.jumbotron h1{color:inherit}.jumbotron p{margin-bottom:15px;font-size:21px;font-weight:200}.jumbotron>hr{border-top-color:#d5d5d5}.container .jumbotron,.container-fluid .jumbotron{padding-right:15px;padding-left:15px;border-radius:6px}.jumbotron .container{max-width:100%}@media screen and (min-width:768px){.jumbotron{padding-top:48px;padding-bottom:48px}.container .jumbotron,.container-fluid .jumbotron{padding-right:60px;padding-left:60px}.jumbotron .h1,.jumbotron h1{font-size:63px}}.thumbnail{display:block;padding:4px;margin-bottom:20px;line-height:1.42857143;background-color:#fff;border:1px solid #ddd;border-radius:4px;-webkit-transition:border .2s ease-in-out;-o-transition:border .2s ease-in-out;transition:border .2s ease-in-out}.thumbnail a>img,.thumbnail>img{margin-right:auto;margin-left:auto}a.thumbnail.active,a.thumbnail:focus,a.thumbnail:hover{border-color:#337ab7}.thumbnail .caption{padding:9px;color:#333}.alert{padding:15px;margin-bottom:20px;border:1px solid transparent;border-radius:4px}.alert h4{margin-top:0;color:inherit}.alert .alert-link{font-weight:700}.alert>p,.alert>ul{margin-bottom:0}.alert>p+p{margin-top:5px}.alert-dismissable,.alert-dismissible{padding-right:35px}.alert-dismissable .close,.alert-dismissible .close{position:relative;top:-2px;right:-21px;color:inherit}.alert-success{color:#3c763d;background-color:#dff0d8;border-color:#d6e9c6}.alert-success hr{border-top-color:#c9e2b3}.alert-success .alert-link{color:#2b542c}.alert-info{color:#31708f;background-color:#d9edf7;border-color:#bce8f1}.alert-info hr{border-top-color:#a6e1ec}.alert-info .alert-link{color:#245269}.alert-warning{color:#8a6d3b;background-color:#fcf8e3;border-color:#faebcc}.alert-warning hr{border-top-color:#f7e1b5}.alert-warning .alert-link{color:#66512c}.alert-danger{color:#a94442;background-color:#f2dede;border-color:#ebccd1}.alert-danger hr{border-top-color:#e4b9c0}.alert-danger .alert-link{color:#843534}@-webkit-keyframes progress-bar-stripes{from{background-position:40px 0}to{background-position:0 0}}@-o-keyframes progress-bar-stripes{from{background-position:40px 0}to{background-position:0 0}}@keyframes progress-bar-stripes{from{background-position:40px 0}to{background-position:0 0}}.progress{height:20px;margin-bottom:20px;overflow:hidden;background-color:#f5f5f5;border-radius:4px;-webkit-box-shadow:inset 0 1px 2px rgba(0,0,0,.1);box-shadow:inset 0 1px 2px rgba(0,0,0,.1)}.progress-bar{float:left;width:0;height:100%;font-size:12px;line-height:20px;color:#fff;text-align:center;background-color:#337ab7;-webkit-box-shadow:inset 0 -1px 0 rgba(0,0,0,.15);box-shadow:inset 0 -1px 0 rgba(0,0,0,.15);-webkit-transition:width .6s ease;-o-transition:width .6s ease;transition:width .6s ease}.progress-bar-striped,.progress-striped .progress-bar{background-image:-webkit-linear-gradient(45deg,rgba(255,255,255,.15) 25%,transparent 25%,transparent 50%,rgba(255,255,255,.15) 50%,rgba(255,255,255,.15) 75%,transparent 75%,transparent);background-image:-o-linear-gradient(45deg,rgba(255,255,255,.15) 25%,transparent 25%,transparent 50%,rgba(255,255,255,.15) 50%,rgba(255,255,255,.15) 75%,transparent 75%,transparent);background-image:linear-gradient(45deg,rgba(255,255,255,.15) 25%,transparent 25%,transparent 50%,rgba(255,255,255,.15) 50%,rgba(255,255,255,.15) 75%,transparent 75%,transparent);-webkit-background-size:40px 40px;background-size:40px 40px}.progress-bar.active,.progress.active .progress-bar{-webkit-animation:progress-bar-stripes 2s linear infinite;-o-animation:progress-bar-stripes 2s linear infinite;animation:progress-bar-stripes 2s linear infinite}.progress-bar-success{background-color:#5cb85c}.progress-striped .progress-bar-success{background-image:-webkit-linear-gradient(45deg,rgba(255,255,255,.15) 25%,transparent 25%,transparent 50%,rgba(255,255,255,.15) 50%,rgba(255,255,255,.15) 75%,transparent 75%,transparent);background-image:-o-linear-gradient(45deg,rgba(255,255,255,.15) 25%,transparent 25%,transparent 50%,rgba(255,255,255,.15) 50%,rgba(255,255,255,.15) 75%,transparent 75%,transparent);background-image:linear-gradient(45deg,rgba(255,255,255,.15) 25%,transparent 25%,transparent 50%,rgba(255,255,255,.15) 50%,rgba(255,255,255,.15) 75%,transparent 75%,transparent)}.progress-bar-info{background-color:#5bc0de}.progress-striped .progress-bar-info{background-image:-webkit-linear-gradient(45deg,rgba(255,255,255,.15) 25%,transparent 25%,transparent 50%,rgba(255,255,255,.15) 50%,rgba(255,255,255,.15) 75%,transparent 75%,transparent);background-image:-o-linear-gradient(45deg,rgba(255,255,255,.15) 25%,transparent 25%,transparent 50%,rgba(255,255,255,.15) 50%,rgba(255,255,255,.15) 75%,transparent 75%,transparent);background-image:linear-gradient(45deg,rgba(255,255,255,.15) 25%,transparent 25%,transparent 50%,rgba(255,255,255,.15) 50%,rgba(255,255,255,.15) 75%,transparent 75%,transparent)}.progress-bar-warning{background-color:#f0ad4e}.progress-striped .progress-bar-warning{background-image:-webkit-linear-gradient(45deg,rgba(255,255,255,.15) 25%,transparent 25%,transparent 50%,rgba(255,255,255,.15) 50%,rgba(255,255,255,.15) 75%,transparent 75%,transparent);background-image:-o-linear-gradient(45deg,rgba(255,255,255,.15) 25%,transparent 25%,transparent 50%,rgba(255,255,255,.15) 50%,rgba(255,255,255,.15) 75%,transparent 75%,transparent);background-image:linear-gradient(45deg,rgba(255,255,255,.15) 25%,transparent 25%,transparent 50%,rgba(255,255,255,.15) 50%,rgba(255,255,255,.15) 75%,transparent 75%,transparent)}.progress-bar-danger{background-color:#d9534f}.progress-striped .progress-bar-danger{background-image:-webkit-linear-gradient(45deg,rgba(255,255,255,.15) 25%,transparent 25%,transparent 50%,rgba(255,255,255,.15) 50%,rgba(255,255,255,.15) 75%,transparent 75%,transparent);background-image:-o-linear-gradient(45deg,rgba(255,255,255,.15) 25%,transparent 25%,transparent 50%,rgba(255,255,255,.15) 50%,rgba(255,255,255,.15) 75%,transparent 75%,transparent);background-image:linear-gradient(45deg,rgba(255,255,255,.15) 25%,transparent 25%,transparent 50%,rgba(255,255,255,.15) 50%,rgba(255,255,255,.15) 75%,transparent 75%,transparent)}.media{margin-top:15px}.media:first-child{margin-top:0}.media,.media-body{overflow:hidden;zoom:1}.media-body{width:10000px}.media-object{display:block}.media-object.img-thumbnail{max-width:none}.media-right,.media>.pull-right{padding-left:10px}.media-left,.media>.pull-left{padding-right:10px}.media-body,.media-left,.media-right{display:table-cell;vertical-align:top}.media-middle{vertical-align:middle}.media-bottom{vertical-align:bottom}.media-heading{margin-top:0;margin-bottom:5px}.media-list{padding-left:0;list-style:none}.list-group{padding-left:0;margin-bottom:20px}.list-group-item{position:relative;display:block;padding:10px 15px;margin-bottom:-1px;background-color:#fff;border:1px solid #ddd}.list-group-item:first-child{border-top-left-radius:4px;border-top-right-radius:4px}.list-group-item:last-child{margin-bottom:0;border-bottom-right-radius:4px;border-bottom-left-radius:4px}a.list-group-item,button.list-group-item{color:#555}a.list-group-item .list-group-item-heading,button.list-group-item .list-group-item-heading{color:#333}a.list-group-item:focus,a.list-group-item:hover,button.list-group-item:focus,button.list-group-item:hover{color:#555;text-decoration:none;background-color:#f5f5f5}button.list-group-item{width:100%;text-align:left}.list-group-item.disabled,.list-group-item.disabled:focus,.list-group-item.disabled:hover{color:#777;cursor:not-allowed;background-color:#eee}.list-group-item.disabled .list-group-item-heading,.list-group-item.disabled:focus .list-group-item-heading,.list-group-item.disabled:hover .list-group-item-heading{color:inherit}.list-group-item.disabled .list-group-item-text,.list-group-item.disabled:focus .list-group-item-text,.list-group-item.disabled:hover .list-group-item-text{color:#777}.list-group-item.active,.list-group-item.active:focus,.list-group-item.active:hover{z-index:2;color:#fff;background-color:#337ab7;border-color:#337ab7}.list-group-item.active .list-group-item-heading,.list-group-item.active .list-group-item-heading>.small,.list-group-item.active .list-group-item-heading>small,.list-group-item.active:focus .list-group-item-heading,.list-group-item.active:focus .list-group-item-heading>.small,.list-group-item.active:focus .list-group-item-heading>small,.list-group-item.active:hover .list-group-item-heading,.list-group-item.active:hover .list-group-item-heading>.small,.list-group-item.active:hover .list-group-item-heading>small{color:inherit}.list-group-item.active .list-group-item-text,.list-group-item.active:focus .list-group-item-text,.list-group-item.active:hover .list-group-item-text{color:#c7ddef}.list-group-item-success{color:#3c763d;background-color:#dff0d8}a.list-group-item-success,button.list-group-item-success{color:#3c763d}a.list-group-item-success .list-group-item-heading,button.list-group-item-success .list-group-item-heading{color:inherit}a.list-group-item-success:focus,a.list-group-item-success:hover,button.list-group-item-success:focus,button.list-group-item-success:hover{color:#3c763d;background-color:#d0e9c6}a.list-group-item-success.active,a.list-group-item-success.active:focus,a.list-group-item-success.active:hover,button.list-group-item-success.active,button.list-group-item-success.active:focus,button.list-group-item-success.active:hover{color:#fff;background-color:#3c763d;border-color:#3c763d}.list-group-item-info{color:#31708f;background-color:#d9edf7}a.list-group-item-info,button.list-group-item-info{color:#31708f}a.list-group-item-info .list-group-item-heading,button.list-group-item-info .list-group-item-heading{color:inherit}a.list-group-item-info:focus,a.list-group-item-info:hover,button.list-group-item-info:focus,button.list-group-item-info:hover{color:#31708f;background-color:#c4e3f3}a.list-group-item-info.active,a.list-group-item-info.active:focus,a.list-group-item-info.active:hover,button.list-group-item-info.active,button.list-group-item-info.active:focus,button.list-group-item-info.active:hover{color:#fff;background-color:#31708f;border-color:#31708f}.list-group-item-warning{color:#8a6d3b;background-color:#fcf8e3}a.list-group-item-warning,button.list-group-item-warning{color:#8a6d3b}a.list-group-item-warning .list-group-item-heading,button.list-group-item-warning .list-group-item-heading{color:inherit}a.list-group-item-warning:focus,a.list-group-item-warning:hover,button.list-group-item-warning:focus,button.list-group-item-warning:hover{color:#8a6d3b;background-color:#faf2cc}a.list-group-item-warning.active,a.list-group-item-warning.active:focus,a.list-group-item-warning.active:hover,button.list-group-item-warning.active,button.list-group-item-warning.active:focus,button.list-group-item-warning.active:hover{color:#fff;background-color:#8a6d3b;border-color:#8a6d3b}.list-group-item-danger{color:#a94442;background-color:#f2dede}a.list-group-item-danger,button.list-group-item-danger{color:#a94442}a.list-group-item-danger .list-group-item-heading,button.list-group-item-danger .list-group-item-heading{color:inherit}a.list-group-item-danger:focus,a.list-group-item-danger:hover,button.list-group-item-danger:focus,button.list-group-item-danger:hover{color:#a94442;background-color:#ebcccc}a.list-group-item-danger.active,a.list-group-item-danger.active:focus,a.list-group-item-danger.active:hover,button.list-group-item-danger.active,button.list-group-item-danger.active:focus,button.list-group-item-danger.active:hover{color:#fff;background-color:#a94442;border-color:#a94442}.list-group-item-heading{margin-top:0;margin-bottom:5px}.list-group-item-text{margin-bottom:0;line-height:1.3}.panel{margin-bottom:20px;background-color:#fff;border:1px solid transparent;border-radius:4px;-webkit-box-shadow:0 1px 1px rgba(0,0,0,.05);box-shadow:0 1px 1px rgba(0,0,0,.05)}.panel-body{padding:15px}.panel-heading{padding:10px 15px;border-bottom:1px solid transparent;border-top-left-radius:3px;border-top-right-radius:3px}.panel-heading>.dropdown .dropdown-toggle{color:inherit}.panel-title{margin-top:0;margin-bottom:0;font-size:16px;color:inherit}.panel-title>.small,.panel-title>.small>a,.panel-title>a,.panel-title>small,.panel-title>small>a{color:inherit}.panel-footer{padding:10px 15px;background-color:#f5f5f5;border-top:1px solid #ddd;border-bottom-right-radius:3px;border-bottom-left-radius:3px}.panel>.list-group,.panel>.panel-collapse>.list-group{margin-bottom:0}.panel>.list-group .list-group-item,.panel>.panel-collapse>.list-group .list-group-item{border-width:1px 0;border-radius:0}.panel>.list-group:first-child .list-group-item:first-child,.panel>.panel-collapse>.list-group:first-child .list-group-item:first-child{border-top:0;border-top-left-radius:3px;border-top-right-radius:3px}.panel>.list-group:last-child .list-group-item:last-child,.panel>.panel-collapse>.list-group:last-child .list-group-item:last-child{border-bottom:0;border-bottom-right-radius:3px;border-bottom-left-radius:3px}.panel>.panel-heading+.panel-collapse>.list-group .list-group-item:first-child{border-top-left-radius:0;border-top-right-radius:0}.panel-heading+.list-group .list-group-item:first-child{border-top-width:0}.list-group+.panel-footer{border-top-width:0}.panel>.panel-collapse>.table,.panel>.table,.panel>.table-responsive>.table{margin-bottom:0}.panel>.panel-collapse>.table caption,.panel>.table caption,.panel>.table-responsive>.table caption{padding-right:15px;padding-left:15px}.panel>.table-responsive:first-child>.table:first-child,.panel>.table:first-child{border-top-left-radius:3px;border-top-right-radius:3px}.panel>.table-responsive:first-child>.table:first-child>tbody:first-child>tr:first-child,.panel>.table-responsive:first-child>.table:first-child>thead:first-child>tr:first-child,.panel>.table:first-child>tbody:first-child>tr:first-child,.panel>.table:first-child>thead:first-child>tr:first-child{border-top-left-radius:3px;border-top-right-radius:3px}.panel>.table-responsive:first-child>.table:first-child>tbody:first-child>tr:first-child td:first-child,.panel>.table-responsive:first-child>.table:first-child>tbody:first-child>tr:first-child th:first-child,.panel>.table-responsive:first-child>.table:first-child>thead:first-child>tr:first-child td:first-child,.panel>.table-responsive:first-child>.table:first-child>thead:first-child>tr:first-child th:first-child,.panel>.table:first-child>tbody:first-child>tr:first-child td:first-child,.panel>.table:first-child>tbody:first-child>tr:first-child th:first-child,.panel>.table:first-child>thead:first-child>tr:first-child td:first-child,.panel>.table:first-child>thead:first-child>tr:first-child th:first-child{border-top-left-radius:3px}.panel>.table-responsive:first-child>.table:first-child>tbody:first-child>tr:first-child td:last-child,.panel>.table-responsive:first-child>.table:first-child>tbody:first-child>tr:first-child th:last-child,.panel>.table-responsive:first-child>.table:first-child>thead:first-child>tr:first-child td:last-child,.panel>.table-responsive:first-child>.table:first-child>thead:first-child>tr:first-child th:last-child,.panel>.table:first-child>tbody:first-child>tr:first-child td:last-child,.panel>.table:first-child>tbody:first-child>tr:first-child th:last-child,.panel>.table:first-child>thead:first-child>tr:first-child td:last-child,.panel>.table:first-child>thead:first-child>tr:first-child th:last-child{border-top-right-radius:3px}.panel>.table-responsive:last-child>.table:last-child,.panel>.table:last-child{border-bottom-right-radius:3px;border-bottom-left-radius:3px}.panel>.table-responsive:last-child>.table:last-child>tbody:last-child>tr:last-child,.panel>.table-responsive:last-child>.table:last-child>tfoot:last-child>tr:last-child,.panel>.table:last-child>tbody:last-child>tr:last-child,.panel>.table:last-child>tfoot:last-child>tr:last-child{border-bottom-right-radius:3px;border-bottom-left-radius:3px}.panel>.table-responsive:last-child>.table:last-child>tbody:last-child>tr:last-child td:first-child,.panel>.table-responsive:last-child>.table:last-child>tbody:last-child>tr:last-child th:first-child,.panel>.table-responsive:last-child>.table:last-child>tfoot:last-child>tr:last-child td:first-child,.panel>.table-responsive:last-child>.table:last-child>tfoot:last-child>tr:last-child th:first-child,.panel>.table:last-child>tbody:last-child>tr:last-child td:first-child,.panel>.table:last-child>tbody:last-child>tr:last-child th:first-child,.panel>.table:last-child>tfoot:last-child>tr:last-child td:first-child,.panel>.table:last-child>tfoot:last-child>tr:last-child th:first-child{border-bottom-left-radius:3px}.panel>.table-responsive:last-child>.table:last-child>tbody:last-child>tr:last-child td:last-child,.panel>.table-responsive:last-child>.table:last-child>tbody:last-child>tr:last-child th:last-child,.panel>.table-responsive:last-child>.table:last-child>tfoot:last-child>tr:last-child td:last-child,.panel>.table-responsive:last-child>.table:last-child>tfoot:last-child>tr:last-child th:last-child,.panel>.table:last-child>tbody:last-child>tr:last-child td:last-child,.panel>.table:last-child>tbody:last-child>tr:last-child th:last-child,.panel>.table:last-child>tfoot:last-child>tr:last-child td:last-child,.panel>.table:last-child>tfoot:last-child>tr:last-child th:last-child{border-bottom-right-radius:3px}.panel>.panel-body+.table,.panel>.panel-body+.table-responsive,.panel>.table+.panel-body,.panel>.table-responsive+.panel-body{border-top:1px solid #ddd}.panel>.table>tbody:first-child>tr:first-child td,.panel>.table>tbody:first-child>tr:first-child th{border-top:0}.panel>.table-bordered,.panel>.table-responsive>.table-bordered{border:0}.panel>.table-bordered>tbody>tr>td:first-child,.panel>.table-bordered>tbody>tr>th:first-child,.panel>.table-bordered>tfoot>tr>td:first-child,.panel>.table-bordered>tfoot>tr>th:first-child,.panel>.table-bordered>thead>tr>td:first-child,.panel>.table-bordered>thead>tr>th:first-child,.panel>.table-responsive>.table-bordered>tbody>tr>td:first-child,.panel>.table-responsive>.table-bordered>tbody>tr>th:first-child,.panel>.table-responsive>.table-bordered>tfoot>tr>td:first-child,.panel>.table-responsive>.table-bordered>tfoot>tr>th:first-child,.panel>.table-responsive>.table-bordered>thead>tr>td:first-child,.panel>.table-responsive>.table-bordered>thead>tr>th:first-child{border-left:0}.panel>.table-bordered>tbody>tr>td:last-child,.panel>.table-bordered>tbody>tr>th:last-child,.panel>.table-bordered>tfoot>tr>td:last-child,.panel>.table-bordered>tfoot>tr>th:last-child,.panel>.table-bordered>thead>tr>td:last-child,.panel>.table-bordered>thead>tr>th:last-child,.panel>.table-responsive>.table-bordered>tbody>tr>td:last-child,.panel>.table-responsive>.table-bordered>tbody>tr>th:last-child,.panel>.table-responsive>.table-bordered>tfoot>tr>td:last-child,.panel>.table-responsive>.table-bordered>tfoot>tr>th:last-child,.panel>.table-responsive>.table-bordered>thead>tr>td:last-child,.panel>.table-responsive>.table-bordered>thead>tr>th:last-child{border-right:0}.panel>.table-bordered>tbody>tr:first-child>td,.panel>.table-bordered>tbody>tr:first-child>th,.panel>.table-bordered>thead>tr:first-child>td,.panel>.table-bordered>thead>tr:first-child>th,.panel>.table-responsive>.table-bordered>tbody>tr:first-child>td,.panel>.table-responsive>.table-bordered>tbody>tr:first-child>th,.panel>.table-responsive>.table-bordered>thead>tr:first-child>td,.panel>.table-responsive>.table-bordered>thead>tr:first-child>th{border-bottom:0}.panel>.table-bordered>tbody>tr:last-child>td,.panel>.table-bordered>tbody>tr:last-child>th,.panel>.table-bordered>tfoot>tr:last-child>td,.panel>.table-bordered>tfoot>tr:last-child>th,.panel>.table-responsive>.table-bordered>tbody>tr:last-child>td,.panel>.table-responsive>.table-bordered>tbody>tr:last-child>th,.panel>.table-responsive>.table-bordered>tfoot>tr:last-child>td,.panel>.table-responsive>.table-bordered>tfoot>tr:last-child>th{border-bottom:0}.panel>.table-responsive{margin-bottom:0;border:0}.panel-group{margin-bottom:20px}.panel-group .panel{margin-bottom:0;border-radius:4px}.panel-group .panel+.panel{margin-top:5px}.panel-group .panel-heading{border-bottom:0}.panel-group .panel-heading+.panel-collapse>.list-group,.panel-group .panel-heading+.panel-collapse>.panel-body{border-top:1px solid #ddd}.panel-group .panel-footer{border-top:0}.panel-group .panel-footer+.panel-collapse .panel-body{border-bottom:1px solid #ddd}.panel-default{border-color:#ddd}.panel-default>.panel-heading{color:#333;background-color:#f5f5f5;border-color:#ddd}.panel-default>.panel-heading+.panel-collapse>.panel-body{border-top-color:#ddd}.panel-default>.panel-heading .badge{color:#f5f5f5;background-color:#333}.panel-default>.panel-footer+.panel-collapse>.panel-body{border-bottom-color:#ddd}.panel-primary{border-color:#337ab7}.panel-primary>.panel-heading{color:#fff;background-color:#337ab7;border-color:#337ab7}.panel-primary>.panel-heading+.panel-collapse>.panel-body{border-top-color:#337ab7}.panel-primary>.panel-heading .badge{color:#337ab7;background-color:#fff}.panel-primary>.panel-footer+.panel-collapse>.panel-body{border-bottom-color:#337ab7}.panel-success{border-color:#d6e9c6}.panel-success>.panel-heading{color:#3c763d;background-color:#dff0d8;border-color:#d6e9c6}.panel-success>.panel-heading+.panel-collapse>.panel-body{border-top-color:#d6e9c6}.panel-success>.panel-heading .badge{color:#dff0d8;background-color:#3c763d}.panel-success>.panel-footer+.panel-collapse>.panel-body{border-bottom-color:#d6e9c6}.panel-info{border-color:#bce8f1}.panel-info>.panel-heading{color:#31708f;background-color:#d9edf7;border-color:#bce8f1}.panel-info>.panel-heading+.panel-collapse>.panel-body{border-top-color:#bce8f1}.panel-info>.panel-heading .badge{color:#d9edf7;background-color:#31708f}.panel-info>.panel-footer+.panel-collapse>.panel-body{border-bottom-color:#bce8f1}.panel-warning{border-color:#faebcc}.panel-warning>.panel-heading{color:#8a6d3b;background-color:#fcf8e3;border-color:#faebcc}.panel-warning>.panel-heading+.panel-collapse>.panel-body{border-top-color:#faebcc}.panel-warning>.panel-heading .badge{color:#fcf8e3;background-color:#8a6d3b}.panel-warning>.panel-footer+.panel-collapse>.panel-body{border-bottom-color:#faebcc}.panel-danger{border-color:#ebccd1}.panel-danger>.panel-heading{color:#a94442;background-color:#f2dede;border-color:#ebccd1}.panel-danger>.panel-heading+.panel-collapse>.panel-body{border-top-color:#ebccd1}.panel-danger>.panel-heading .badge{color:#f2dede;background-color:#a94442}.panel-danger>.panel-footer+.panel-collapse>.panel-body{border-bottom-color:#ebccd1}.embed-responsive{position:relative;display:block;height:0;padding:0;overflow:hidden}.embed-responsive .embed-responsive-item,.embed-responsive embed,.embed-responsive iframe,.embed-responsive object,.embed-responsive video{position:absolute;top:0;bottom:0;left:0;width:100%;height:100%;border:0}.embed-responsive-16by9{padding-bottom:56.25%}.embed-responsive-4by3{padding-bottom:75%}.well{min-height:20px;padding:19px;margin-bottom:20px;background-color:#f5f5f5;border:1px solid #e3e3e3;border-radius:4px;-webkit-box-shadow:inset 0 1px 1px rgba(0,0,0,.05);box-shadow:inset 0 1px 1px rgba(0,0,0,.05)}.well blockquote{border-color:#ddd;border-color:rgba(0,0,0,.15)}.well-lg{padding:24px;border-radius:6px}.well-sm{padding:9px;border-radius:3px}.close{float:right;font-size:21px;font-weight:700;line-height:1;color:#000;text-shadow:0 1px 0 #fff;filter:alpha(opacity=20);opacity:.2}.close:focus,.close:hover{color:#000;text-decoration:none;cursor:pointer;filter:alpha(opacity=50);opacity:.5}button.close{-webkit-appearance:none;padding:0;cursor:pointer;background:0 0;border:0}.modal-open{overflow:hidden}.modal{position:fixed;top:0;right:0;bottom:0;left:0;z-index:1050;display:none;overflow:hidden;-webkit-overflow-scrolling:touch;outline:0}.modal.fade .modal-dialog{-webkit-transition:-webkit-transform .3s ease-out;-o-transition:-o-transform .3s ease-out;transition:transform .3s ease-out;-webkit-transform:translate(0,-25%);-ms-transform:translate(0,-25%);-o-transform:translate(0,-25%);transform:translate(0,-25%)}.modal.in .modal-dialog{-webkit-transform:translate(0,0);-ms-transform:translate(0,0);-o-transform:translate(0,0);transform:translate(0,0)}.modal-open .modal{overflow-x:hidden;overflow-y:auto}.modal-dialog{position:relative;width:auto;margin:10px}.modal-content{position:relative;background-color:#fff;-webkit-background-clip:padding-box;background-clip:padding-box;border:1px solid #999;border:1px solid rgba(0,0,0,.2);border-radius:6px;outline:0;-webkit-box-shadow:0 3px 9px rgba(0,0,0,.5);box-shadow:0 3px 9px rgba(0,0,0,.5)}.modal-backdrop{position:fixed;top:0;right:0;bottom:0;left:0;z-index:1040;background-color:#000}.modal-backdrop.fade{filter:alpha(opacity=0);opacity:0}.modal-backdrop.in{filter:alpha(opacity=50);opacity:.5}.modal-header{padding:15px;border-bottom:1px solid #e5e5e5}.modal-header .close{margin-top:-2px}.modal-title{margin:0;line-height:1.42857143}.modal-body{position:relative;padding:15px}.modal-footer{padding:15px;text-align:right;border-top:1px solid #e5e5e5}.modal-footer .btn+.btn{margin-bottom:0;margin-left:5px}.modal-footer .btn-group .btn+.btn{margin-left:-1px}.modal-footer .btn-block+.btn-block{margin-left:0}.modal-scrollbar-measure{position:absolute;top:-9999px;width:50px;height:50px;overflow:scroll}@media (min-width:768px){.modal-dialog{width:600px;margin:30px auto}.modal-content{-webkit-box-shadow:0 5px 15px rgba(0,0,0,.5);box-shadow:0 5px 15px rgba(0,0,0,.5)}.modal-sm{width:300px}}@media (min-width:992px){.modal-lg{width:900px}}.tooltip{position:absolute;z-index:1070;display:block;font-family:"Helvetica Neue",Helvetica,Arial,sans-serif;font-size:12px;font-style:normal;font-weight:400;line-height:1.42857143;text-align:left;text-align:start;text-decoration:none;text-shadow:none;text-transform:none;letter-spacing:normal;word-break:normal;word-spacing:normal;word-wrap:normal;white-space:normal;filter:alpha(opacity=0);opacity:0;line-break:auto}.tooltip.in{filter:alpha(opacity=90);opacity:.9}.tooltip.top{padding:5px 0;margin-top:-3px}.tooltip.right{padding:0 5px;margin-left:3px}.tooltip.bottom{padding:5px 0;margin-top:3px}.tooltip.left{padding:0 5px;margin-left:-3px}.tooltip-inner{max-width:200px;padding:3px 8px;color:#fff;text-align:center;background-color:#000;border-radius:4px}.tooltip-arrow{position:absolute;width:0;height:0;border-color:transparent;border-style:solid}.tooltip.top .tooltip-arrow{bottom:0;left:50%;margin-left:-5px;border-width:5px 5px 0;border-top-color:#000}.tooltip.top-left .tooltip-arrow{right:5px;bottom:0;margin-bottom:-5px;border-width:5px 5px 0;border-top-color:#000}.tooltip.top-right .tooltip-arrow{bottom:0;left:5px;margin-bottom:-5px;border-width:5px 5px 0;border-top-color:#000}.tooltip.right .tooltip-arrow{top:50%;left:0;margin-top:-5px;border-width:5px 5px 5px 0;border-right-color:#000}.tooltip.left .tooltip-arrow{top:50%;right:0;margin-top:-5px;border-width:5px 0 5px 5px;border-left-color:#000}.tooltip.bottom .tooltip-arrow{top:0;left:50%;margin-left:-5px;border-width:0 5px 5px;border-bottom-color:#000}.tooltip.bottom-left .tooltip-arrow{top:0;right:5px;margin-top:-5px;border-width:0 5px 5px;border-bottom-color:#000}.tooltip.bottom-right .tooltip-arrow{top:0;left:5px;margin-top:-5px;border-width:0 5px 5px;border-bottom-color:#000}.popover{position:absolute;top:0;left:0;z-index:1060;display:none;max-width:276px;padding:1px;font-family:"Helvetica Neue",Helvetica,Arial,sans-serif;font-size:14px;font-style:normal;font-weight:400;line-height:1.42857143;text-align:left;text-align:start;text-decoration:none;text-shadow:none;text-transform:none;letter-spacing:normal;word-break:normal;word-spacing:normal;word-wrap:normal;white-space:normal;background-color:#fff;-webkit-background-clip:padding-box;background-clip:padding-box;border:1px solid #ccc;border:1px solid rgba(0,0,0,.2);border-radius:6px;-webkit-box-shadow:0 5px 10px rgba(0,0,0,.2);box-shadow:0 5px 10px rgba(0,0,0,.2);line-break:auto}.popover.top{margin-top:-10px}.popover.right{margin-left:10px}.popover.bottom{margin-top:10px}.popover.left{margin-left:-10px}.popover-title{padding:8px 14px;margin:0;font-size:14px;background-color:#f7f7f7;border-bottom:1px solid #ebebeb;border-radius:5px 5px 0 0}.popover-content{padding:9px 14px}.popover>.arrow,.popover>.arrow:after{position:absolute;display:block;width:0;height:0;border-color:transparent;border-style:solid}.popover>.arrow{border-width:11px}.popover>.arrow:after{content:"";border-width:10px}.popover.top>.arrow{bottom:-11px;left:50%;margin-left:-11px;border-top-color:#999;border-top-color:rgba(0,0,0,.25);border-bottom-width:0}.popover.top>.arrow:after{bottom:1px;margin-left:-10px;content:" ";border-top-color:#fff;border-bottom-width:0}.popover.right>.arrow{top:50%;left:-11px;margin-top:-11px;border-right-color:#999;border-right-color:rgba(0,0,0,.25);border-left-width:0}.popover.right>.arrow:after{bottom:-10px;left:1px;content:" ";border-right-color:#fff;border-left-width:0}.popover.bottom>.arrow{top:-11px;left:50%;margin-left:-11px;border-top-width:0;border-bottom-color:#999;border-bottom-color:rgba(0,0,0,.25)}.popover.bottom>.arrow:after{top:1px;margin-left:-10px;content:" ";border-top-width:0;border-bottom-color:#fff}.popover.left>.arrow{top:50%;right:-11px;margin-top:-11px;border-right-width:0;border-left-color:#999;border-left-color:rgba(0,0,0,.25)}.popover.left>.arrow:after{right:1px;bottom:-10px;content:" ";border-right-width:0;border-left-color:#fff}.carousel{position:relative}.carousel-inner{position:relative;width:100%;overflow:hidden}.carousel-inner>.item{position:relative;display:none;-webkit-transition:.6s ease-in-out left;-o-transition:.6s ease-in-out left;transition:.6s ease-in-out left}.carousel-inner>.item>a>img,.carousel-inner>.item>img{line-height:1}@media all and (transform-3d),(-webkit-transform-3d){.carousel-inner>.item{-webkit-transition:-webkit-transform .6s ease-in-out;-o-transition:-o-transform .6s ease-in-out;transition:transform .6s ease-in-out;-webkit-backface-visibility:hidden;backface-visibility:hidden;-webkit-perspective:1000px;perspective:1000px}.carousel-inner>.item.active.right,.carousel-inner>.item.next{left:0;-webkit-transform:translate3d(100%,0,0);transform:translate3d(100%,0,0)}.carousel-inner>.item.active.left,.carousel-inner>.item.prev{left:0;-webkit-transform:translate3d(-100%,0,0);transform:translate3d(-100%,0,0)}.carousel-inner>.item.active,.carousel-inner>.item.next.left,.carousel-inner>.item.prev.right{left:0;-webkit-transform:translate3d(0,0,0);transform:translate3d(0,0,0)}}.carousel-inner>.active,.carousel-inner>.next,.carousel-inner>.prev{display:block}.carousel-inner>.active{left:0}.carousel-inner>.next,.carousel-inner>.prev{position:absolute;top:0;width:100%}.carousel-inner>.next{left:100%}.carousel-inner>.prev{left:-100%}.carousel-inner>.next.left,.carousel-inner>.prev.right{left:0}.carousel-inner>.active.left{left:-100%}.carousel-inner>.active.right{left:100%}.carousel-control{position:absolute;top:0;bottom:0;left:0;width:15%;font-size:20px;color:#fff;text-align:center;text-shadow:0 1px 2px rgba(0,0,0,.6);background-color:rgba(0,0,0,0);filter:alpha(opacity=50);opacity:.5}.carousel-control.left{background-image:-webkit-linear-gradient(left,rgba(0,0,0,.5) 0,rgba(0,0,0,.0001) 100%);background-image:-o-linear-gradient(left,rgba(0,0,0,.5) 0,rgba(0,0,0,.0001) 100%);background-image:-webkit-gradient(linear,left top,right top,from(rgba(0,0,0,.5)),to(rgba(0,0,0,.0001)));background-image:linear-gradient(to right,rgba(0,0,0,.5) 0,rgba(0,0,0,.0001) 100%);filter:progid:DXImageTransform.Microsoft.gradient(startColorstr='#80000000', endColorstr='#00000000', GradientType=1);background-repeat:repeat-x}.carousel-control.right{right:0;left:auto;background-image:-webkit-linear-gradient(left,rgba(0,0,0,.0001) 0,rgba(0,0,0,.5) 100%);background-image:-o-linear-gradient(left,rgba(0,0,0,.0001) 0,rgba(0,0,0,.5) 100%);background-image:-webkit-gradient(linear,left top,right top,from(rgba(0,0,0,.0001)),to(rgba(0,0,0,.5)));background-image:linear-gradient(to right,rgba(0,0,0,.0001) 0,rgba(0,0,0,.5) 100%);filter:progid:DXImageTransform.Microsoft.gradient(startColorstr='#00000000', endColorstr='#80000000', GradientType=1);background-repeat:repeat-x}.carousel-control:focus,.carousel-control:hover{color:#fff;text-decoration:none;filter:alpha(opacity=90);outline:0;opacity:.9}.carousel-control .glyphicon-chevron-left,.carousel-control .glyphicon-chevron-right,.carousel-control .icon-next,.carousel-control .icon-prev{position:absolute;top:50%;z-index:5;display:inline-block;margin-top:-10px}.carousel-control .glyphicon-chevron-left,.carousel-control .icon-prev{left:50%;margin-left:-10px}.carousel-control .glyphicon-chevron-right,.carousel-control .icon-next{right:50%;margin-right:-10px}.carousel-control .icon-next,.carousel-control .icon-prev{width:20px;height:20px;font-family:serif;line-height:1}.carousel-control .icon-prev:before{content:'\2039'}.carousel-control .icon-next:before{content:'\203a'}.carousel-indicators{position:absolute;bottom:10px;left:50%;z-index:15;width:60%;padding-left:0;margin-left:-30%;text-align:center;list-style:none}.carousel-indicators li{display:inline-block;width:10px;height:10px;margin:1px;text-indent:-999px;cursor:pointer;background-color:#000\9;background-color:rgba(0,0,0,0);border:1px solid #fff;border-radius:10px}.carousel-indicators .active{width:12px;height:12px;margin:0;background-color:#fff}.carousel-caption{position:absolute;right:15%;bottom:20px;left:15%;z-index:10;padding-top:20px;padding-bottom:20px;color:#fff;text-align:center;text-shadow:0 1px 2px rgba(0,0,0,.6)}.carousel-caption .btn{text-shadow:none}@media screen and (min-width:768px){.carousel-control .glyphicon-chevron-left,.carousel-control .glyphicon-chevron-right,.carousel-control .icon-next,.carousel-control .icon-prev{width:30px;height:30px;margin-top:-10px;font-size:30px}.carousel-control .glyphicon-chevron-left,.carousel-control .icon-prev{margin-left:-10px}.carousel-control .glyphicon-chevron-right,.carousel-control .icon-next{margin-right:-10px}.carousel-caption{right:20%;left:20%;padding-bottom:30px}.carousel-indicators{bottom:20px}}.btn-group-vertical>.btn-group:after,.btn-group-vertical>.btn-group:before,.btn-toolbar:after,.btn-toolbar:before,.clearfix:after,.clearfix:before,.container-fluid:after,.container-fluid:before,.container:after,.container:before,.dl-horizontal dd:after,.dl-horizontal dd:before,.form-horizontal .form-group:after,.form-horizontal .form-group:before,.modal-footer:after,.modal-footer:before,.modal-header:after,.modal-header:before,.nav:after,.nav:before,.navbar-collapse:after,.navbar-collapse:before,.navbar-header:after,.navbar-header:before,.navbar:after,.navbar:before,.pager:after,.pager:before,.panel-body:after,.panel-body:before,.row:after,.row:before{display:table;content:" "}.btn-group-vertical>.btn-group:after,.btn-toolbar:after,.clearfix:after,.container-fluid:after,.container:after,.dl-horizontal dd:after,.form-horizontal .form-group:after,.modal-footer:after,.modal-header:after,.nav:after,.navbar-collapse:after,.navbar-header:after,.navbar:after,.pager:after,.panel-body:after,.row:after{clear:both}.center-block{display:block;margin-right:auto;margin-left:auto}.pull-right{float:right!important}.pull-left{float:left!important}.hide{display:none!important}.show{display:block!important}.invisible{visibility:hidden}.text-hide{font:0/0 a;color:transparent;text-shadow:none;background-color:transparent;border:0}.hidden{display:none!important}.affix{position:fixed}@-ms-viewport{width:device-width}.visible-lg,.visible-md,.visible-sm,.visible-xs{display:none!important}.visible-lg-block,.visible-lg-inline,.visible-lg-inline-block,.visible-md-block,.visible-md-inline,.visible-md-inline-block,.visible-sm-block,.visible-sm-inline,.visible-sm-inline-block,.visible-xs-block,.visible-xs-inline,.visible-xs-inline-block{display:none!important}@media (max-width:767px){.visible-xs{display:block!important}table.visible-xs{display:table!important}tr.visible-xs{display:table-row!important}td.visible-xs,th.visible-xs{display:table-cell!important}}@media (max-width:767px){.visible-xs-block{display:block!important}}@media (max-width:767px){.visible-xs-inline{display:inline!important}}@media (max-width:767px){.visible-xs-inline-block{display:inline-block!important}}@media (min-width:768px) and (max-width:991px){.visible-sm{display:block!important}table.visible-sm{display:table!important}tr.visible-sm{display:table-row!important}td.visible-sm,th.visible-sm{display:table-cell!important}}@media (min-width:768px) and (max-width:991px){.visible-sm-block{display:block!important}}@media (min-width:768px) and (max-width:991px){.visible-sm-inline{display:inline!important}}@media (min-width:768px) and (max-width:991px){.visible-sm-inline-block{display:inline-block!important}}@media (min-width:992px) and (max-width:1199px){.visible-md{display:block!important}table.visible-md{display:table!important}tr.visible-md{display:table-row!important}td.visible-md,th.visible-md{display:table-cell!important}}@media (min-width:992px) and (max-width:1199px){.visible-md-block{display:block!important}}@media (min-width:992px) and (max-width:1199px){.visible-md-inline{display:inline!important}}@media (min-width:992px) and (max-width:1199px){.visible-md-inline-block{display:inline-block!important}}@media (min-width:1200px){.visible-lg{display:block!important}table.visible-lg{display:table!important}tr.visible-lg{display:table-row!important}td.visible-lg,th.visible-lg{display:table-cell!important}}@media (min-width:1200px){.visible-lg-block{display:block!important}}@media (min-width:1200px){.visible-lg-inline{display:inline!important}}@media (min-width:1200px){.visible-lg-inline-block{display:inline-block!important}}@media (max-width:767px){.hidden-xs{display:none!important}}@media (min-width:768px) and (max-width:991px){.hidden-sm{display:none!important}}@media (min-width:992px) and (max-width:1199px){.hidden-md{display:none!important}}@media (min-width:1200px){.hidden-lg{display:none!important}}.visible-print{display:none!important}@media print{.visible-print{display:block!important}table.visible-print{display:table!important}tr.visible-print{display:table-row!important}td.visible-print,th.visible-print{display:table-cell!important}}.visible-print-block{display:none!important}@media print{.visible-print-block{display:block!important}}.visible-print-inline{display:none!important}@media print{.visible-print-inline{display:inline!important}}.visible-print-inline-block{display:none!important}@media print{.visible-print-inline-block{display:inline-block!important}}@media print{.hidden-print{display:none!important}}
/*# sourceMappingURL=bootstrap.min.css.map */`

var bootstrapSlateCSS string = `/*! normalize.css v2.1.3 | MIT License | git.io/normalize */

article,
aside,
details,
figcaption,
figure,
footer,
header,
hgroup,
main,
nav,
section,
summary {
  display: block;
}

audio,
canvas,
video {
  display: inline-block;
}

audio:not([controls]) {
  display: none;
  height: 0;
}

[hidden],
template {
  display: none;
}

html {
  font-family: sans-serif;
  -webkit-text-size-adjust: 100%;
      -ms-text-size-adjust: 100%;
}

body {
  margin: 0;
}

a {
  background: transparent;
}

a:focus {
  outline: thin dotted;
}

a:active,
a:hover {
  outline: 0;
}

h1 {
  margin: 0.67em 0;
  font-size: 2em;
}

abbr[title] {
  border-bottom: 1px dotted;
}

b,
strong {
  font-weight: bold;
}

dfn {
  font-style: italic;
}

hr {
  height: 0;
  -moz-box-sizing: content-box;
       box-sizing: content-box;
}

mark {
  color: #000;
  background: #ff0;
}

code,
kbd,
pre,
samp {
  font-family: monospace, serif;
  font-size: 1em;
}

pre {
  white-space: pre-wrap;
}

q {
  quotes: "\201C" "\201D" "\2018" "\2019";
}

small {
  font-size: 80%;
}

sub,
sup {
  position: relative;
  font-size: 75%;
  line-height: 0;
  vertical-align: baseline;
}

sup {
  top: -0.5em;
}

sub {
  bottom: -0.25em;
}

img {
  border: 0;
}

svg:not(:root) {
  overflow: hidden;
}

figure {
  margin: 0;
}

fieldset {
  padding: 0.35em 0.625em 0.75em;
  margin: 0 2px;
  border: 1px solid #c0c0c0;
}

legend {
  padding: 0;
  border: 0;
}

button,
input,
select,
textarea {
  margin: 0;
  font-family: inherit;
  font-size: 100%;
}

button,
input {
  line-height: normal;
}

button,
select {
  text-transform: none;
}

button,
html input[type="button"],
input[type="reset"],
input[type="submit"] {
  cursor: pointer;
  -webkit-appearance: button;
}

button[disabled],
html input[disabled] {
  cursor: default;
}

input[type="checkbox"],
input[type="radio"] {
  padding: 0;
  box-sizing: border-box;
}

input[type="search"] {
  -webkit-box-sizing: content-box;
     -moz-box-sizing: content-box;
          box-sizing: content-box;
  -webkit-appearance: textfield;
}

input[type="search"]::-webkit-search-cancel-button,
input[type="search"]::-webkit-search-decoration {
  -webkit-appearance: none;
}

button::-moz-focus-inner,
input::-moz-focus-inner {
  padding: 0;
  border: 0;
}

textarea {
  overflow: auto;
  vertical-align: top;
}

table {
  border-collapse: collapse;
  border-spacing: 0;
}

@media print {
  * {
    color: #000 !important;
    text-shadow: none !important;
    background: transparent !important;
    box-shadow: none !important;
  }
  a,
  a:visited {
    text-decoration: underline;
  }
  a[href]:after {
    content: " (" attr(href) ")";
  }
  abbr[title]:after {
    content: " (" attr(title) ")";
  }
  a[href^="javascript:"]:after,
  a[href^="#"]:after {
    content: "";
  }
  pre,
  blockquote {
    border: 1px solid #999;
    page-break-inside: avoid;
  }
  thead {
    display: table-header-group;
  }
  tr,
  img {
    page-break-inside: avoid;
  }
  img {
    max-width: 100% !important;
  }
  @page  {
    margin: 2cm .5cm;
  }
  p,
  h2,
  h3 {
    orphans: 3;
    widows: 3;
  }
  h2,
  h3 {
    page-break-after: avoid;
  }
  select {
    background: #fff !important;
  }
  .navbar {
    display: none;
  }
  .table td,
  .table th {
    background-color: #fff !important;
  }
  .btn > .caret,
  .dropup > .btn > .caret {
    border-top-color: #000 !important;
  }
  .label {
    border: 1px solid #000;
  }
  .table {
    border-collapse: collapse !important;
  }
  .table-bordered th,
  .table-bordered td {
    border: 1px solid #ddd !important;
  }
}

*,
*:before,
*:after {
  -webkit-box-sizing: border-box;
     -moz-box-sizing: border-box;
          box-sizing: border-box;
}

html {
  font-size: 62.5%;
  -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
}

body {
  font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
  font-size: 14px;
  line-height: 1.428571429;
  color: #c8c8c8;
  background-color: #272b30;
}

input,
button,
select,
textarea {
  font-family: inherit;
  font-size: inherit;
  line-height: inherit;
}

a {
  color: #ffffff;
  text-decoration: none;
}

a:hover,
a:focus {
  color: #ffffff;
  text-decoration: underline;
}

a:focus {
  outline: thin dotted;
  outline: 5px auto -webkit-focus-ring-color;
  outline-offset: -2px;
}

img {
  vertical-align: middle;
}

.img-responsive {
  display: block;
  height: auto;
  max-width: 100%;
}

.img-rounded {
  border-radius: 6px;
}

.img-thumbnail {
  display: inline-block;
  height: auto;
  max-width: 100%;
  padding: 4px;
  line-height: 1.428571429;
  background-color: #272b30;
  border: 1px solid #dddddd;
  border-radius: 4px;
  -webkit-transition: all 0.2s ease-in-out;
          transition: all 0.2s ease-in-out;
}

.img-circle {
  border-radius: 50%;
}

hr {
  margin-top: 20px;
  margin-bottom: 20px;
  border: 0;
  border-top: 1px solid #1c1e22;
}

.sr-only {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  border: 0;
}

h1,
h2,
h3,
h4,
h5,
h6,
.h1,
.h2,
.h3,
.h4,
.h5,
.h6 {
  font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
  font-weight: 500;
  line-height: 1.1;
  color: inherit;
}

h1 small,
h2 small,
h3 small,
h4 small,
h5 small,
h6 small,
.h1 small,
.h2 small,
.h3 small,
.h4 small,
.h5 small,
.h6 small,
h1 .small,
h2 .small,
h3 .small,
h4 .small,
h5 .small,
h6 .small,
.h1 .small,
.h2 .small,
.h3 .small,
.h4 .small,
.h5 .small,
.h6 .small {
  font-weight: normal;
  line-height: 1;
  color: #7a8288;
}

h1,
h2,
h3 {
  margin-top: 20px;
  margin-bottom: 10px;
}

h1 small,
h2 small,
h3 small,
h1 .small,
h2 .small,
h3 .small {
  font-size: 65%;
}

h4,
h5,
h6 {
  margin-top: 10px;
  margin-bottom: 10px;
}

h4 small,
h5 small,
h6 small,
h4 .small,
h5 .small,
h6 .small {
  font-size: 75%;
}

h1,
.h1 {
  font-size: 36px;
}

h2,
.h2 {
  font-size: 30px;
}

h3,
.h3 {
  font-size: 24px;
}

h4,
.h4 {
  font-size: 18px;
}

h5,
.h5 {
  font-size: 14px;
}

h6,
.h6 {
  font-size: 12px;
}

p {
  margin: 0 0 10px;
}

.lead {
  margin-bottom: 20px;
  font-size: 16px;
  font-weight: 200;
  line-height: 1.4;
}

@media (min-width: 768px) {
  .lead {
    font-size: 21px;
  }
}

small,
.small {
  font-size: 85%;
}

cite {
  font-style: normal;
}

.text-muted {
  color: #7a8288;
}

.text-primary {
  color: #7a8288;
}

.text-primary:hover {
  color: #62686d;
}

.text-warning {
  color: #ffffff;
}

.text-warning:hover {
  color: #e6e6e6;
}

.text-danger {
  color: #ffffff;
}

.text-danger:hover {
  color: #e6e6e6;
}

.text-success {
  color: #ffffff;
}

.text-success:hover {
  color: #e6e6e6;
}

.text-info {
  color: #ffffff;
}

.text-info:hover {
  color: #e6e6e6;
}

.text-left {
  text-align: left;
}

.text-right {
  text-align: right;
}

.text-center {
  text-align: center;
}

.page-header {
  padding-bottom: 9px;
  margin: 40px 0 20px;
  border-bottom: 1px solid #1c1e22;
}

ul,
ol {
  margin-top: 0;
  margin-bottom: 10px;
}

ul ul,
ol ul,
ul ol,
ol ol {
  margin-bottom: 0;
}

.list-unstyled {
  padding-left: 0;
  list-style: none;
}

.list-inline {
  padding-left: 0;
  list-style: none;
}

.list-inline > li {
  display: inline-block;
  padding-right: 5px;
  padding-left: 5px;
}

.list-inline > li:first-child {
  padding-left: 0;
}

dl {
  margin-top: 0;
  margin-bottom: 20px;
}

dt,
dd {
  line-height: 1.428571429;
}

dt {
  font-weight: bold;
}

dd {
  margin-left: 0;
}

@media (min-width: 768px) {
  .dl-horizontal dt {
    float: left;
    width: 160px;
    overflow: hidden;
    clear: left;
    text-align: right;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
  .dl-horizontal dd {
    margin-left: 180px;
  }
  .dl-horizontal dd:before,
  .dl-horizontal dd:after {
    display: table;
    content: " ";
  }
  .dl-horizontal dd:after {
    clear: both;
  }
  .dl-horizontal dd:before,
  .dl-horizontal dd:after {
    display: table;
    content: " ";
  }
  .dl-horizontal dd:after {
    clear: both;
  }
  .dl-horizontal dd:before,
  .dl-horizontal dd:after {
    display: table;
    content: " ";
  }
  .dl-horizontal dd:after {
    clear: both;
  }
  .dl-horizontal dd:before,
  .dl-horizontal dd:after {
    display: table;
    content: " ";
  }
  .dl-horizontal dd:after {
    clear: both;
  }
  .dl-horizontal dd:before,
  .dl-horizontal dd:after {
    display: table;
    content: " ";
  }
  .dl-horizontal dd:after {
    clear: both;
  }
}

abbr[title],
abbr[data-original-title] {
  cursor: help;
  border-bottom: 1px dotted #7a8288;
}

.initialism {
  font-size: 90%;
  text-transform: uppercase;
}

blockquote {
  padding: 10px 20px;
  margin: 0 0 20px;
  border-left: 5px solid #7a8288;
}

blockquote p {
  font-size: 17.5px;
  font-weight: 300;
  line-height: 1.25;
}

blockquote p:last-child {
  margin-bottom: 0;
}

blockquote small,
blockquote .small {
  display: block;
  line-height: 1.428571429;
  color: #7a8288;
}

blockquote small:before,
blockquote .small:before {
  content: '\2014 \00A0';
}

blockquote.pull-right {
  padding-right: 15px;
  padding-left: 0;
  border-right: 5px solid #7a8288;
  border-left: 0;
}

blockquote.pull-right p,
blockquote.pull-right small,
blockquote.pull-right .small {
  text-align: right;
}

blockquote.pull-right small:before,
blockquote.pull-right .small:before {
  content: '';
}

blockquote.pull-right small:after,
blockquote.pull-right .small:after {
  content: '\00A0 \2014';
}

blockquote:before,
blockquote:after {
  content: "";
}

address {
  margin-bottom: 20px;
  font-style: normal;
  line-height: 1.428571429;
}

code,
kbd,
pre,
samp {
  font-family: Monaco, Menlo, Consolas, "Courier New", monospace;
}

code {
  padding: 2px 4px;
  font-size: 90%;
  color: #c7254e;
  white-space: nowrap;
  background-color: #f9f2f4;
  border-radius: 4px;
}

pre {
  display: block;
  padding: 9.5px;
  margin: 0 0 10px;
  font-size: 13px;
  line-height: 1.428571429;
  color: #3a3f44;
  word-break: break-all;
  word-wrap: break-word;
  background-color: #f5f5f5;
  border: 1px solid #cccccc;
  border-radius: 4px;
}

pre code {
  padding: 0;
  font-size: inherit;
  color: inherit;
  white-space: pre-wrap;
  background-color: transparent;
  border-radius: 0;
}

.pre-scrollable {
  max-height: 340px;
  overflow-y: scroll;
}

.container {
  padding-right: 15px;
  padding-left: 15px;
  margin-right: auto;
  margin-left: auto;
}

.container:before,
.container:after {
  display: table;
  content: " ";
}

.container:after {
  clear: both;
}

.container:before,
.container:after {
  display: table;
  content: " ";
}

.container:after {
  clear: both;
}

.container:before,
.container:after {
  display: table;
  content: " ";
}

.container:after {
  clear: both;
}

.container:before,
.container:after {
  display: table;
  content: " ";
}

.container:after {
  clear: both;
}

.container:before,
.container:after {
  display: table;
  content: " ";
}

.container:after {
  clear: both;
}

@media (min-width: 768px) {
  .container {
    width: 750px;
  }
}

@media (min-width: 992px) {
  .container {
    width: 970px;
  }
}

@media (min-width: 1200px) {
  .container {
    width: 1170px;
  }
}

.row {
  margin-right: -15px;
  margin-left: -15px;
}

.row:before,
.row:after {
  display: table;
  content: " ";
}

.row:after {
  clear: both;
}

.row:before,
.row:after {
  display: table;
  content: " ";
}

.row:after {
  clear: both;
}

.row:before,
.row:after {
  display: table;
  content: " ";
}

.row:after {
  clear: both;
}

.row:before,
.row:after {
  display: table;
  content: " ";
}

.row:after {
  clear: both;
}

.row:before,
.row:after {
  display: table;
  content: " ";
}

.row:after {
  clear: both;
}

.col-xs-1,
.col-sm-1,
.col-md-1,
.col-lg-1,
.col-xs-2,
.col-sm-2,
.col-md-2,
.col-lg-2,
.col-xs-3,
.col-sm-3,
.col-md-3,
.col-lg-3,
.col-xs-4,
.col-sm-4,
.col-md-4,
.col-lg-4,
.col-xs-5,
.col-sm-5,
.col-md-5,
.col-lg-5,
.col-xs-6,
.col-sm-6,
.col-md-6,
.col-lg-6,
.col-xs-7,
.col-sm-7,
.col-md-7,
.col-lg-7,
.col-xs-8,
.col-sm-8,
.col-md-8,
.col-lg-8,
.col-xs-9,
.col-sm-9,
.col-md-9,
.col-lg-9,
.col-xs-10,
.col-sm-10,
.col-md-10,
.col-lg-10,
.col-xs-11,
.col-sm-11,
.col-md-11,
.col-lg-11,
.col-xs-12,
.col-sm-12,
.col-md-12,
.col-lg-12 {
  position: relative;
  min-height: 1px;
  padding-right: 15px;
  padding-left: 15px;
}

.col-xs-1,
.col-xs-2,
.col-xs-3,
.col-xs-4,
.col-xs-5,
.col-xs-6,
.col-xs-7,
.col-xs-8,
.col-xs-9,
.col-xs-10,
.col-xs-11,
.col-xs-12 {
  float: left;
}

.col-xs-12 {
  width: 100%;
}

.col-xs-11 {
  width: 91.66666666666666%;
}

.col-xs-10 {
  width: 83.33333333333334%;
}

.col-xs-9 {
  width: 75%;
}

.col-xs-8 {
  width: 66.66666666666666%;
}

.col-xs-7 {
  width: 58.333333333333336%;
}

.col-xs-6 {
  width: 50%;
}

.col-xs-5 {
  width: 41.66666666666667%;
}

.col-xs-4 {
  width: 33.33333333333333%;
}

.col-xs-3 {
  width: 25%;
}

.col-xs-2 {
  width: 16.666666666666664%;
}

.col-xs-1 {
  width: 8.333333333333332%;
}

.col-xs-pull-12 {
  right: 100%;
}

.col-xs-pull-11 {
  right: 91.66666666666666%;
}

.col-xs-pull-10 {
  right: 83.33333333333334%;
}

.col-xs-pull-9 {
  right: 75%;
}

.col-xs-pull-8 {
  right: 66.66666666666666%;
}

.col-xs-pull-7 {
  right: 58.333333333333336%;
}

.col-xs-pull-6 {
  right: 50%;
}

.col-xs-pull-5 {
  right: 41.66666666666667%;
}

.col-xs-pull-4 {
  right: 33.33333333333333%;
}

.col-xs-pull-3 {
  right: 25%;
}

.col-xs-pull-2 {
  right: 16.666666666666664%;
}

.col-xs-pull-1 {
  right: 8.333333333333332%;
}

.col-xs-pull-0 {
  right: 0;
}

.col-xs-push-12 {
  left: 100%;
}

.col-xs-push-11 {
  left: 91.66666666666666%;
}

.col-xs-push-10 {
  left: 83.33333333333334%;
}

.col-xs-push-9 {
  left: 75%;
}

.col-xs-push-8 {
  left: 66.66666666666666%;
}

.col-xs-push-7 {
  left: 58.333333333333336%;
}

.col-xs-push-6 {
  left: 50%;
}

.col-xs-push-5 {
  left: 41.66666666666667%;
}

.col-xs-push-4 {
  left: 33.33333333333333%;
}

.col-xs-push-3 {
  left: 25%;
}

.col-xs-push-2 {
  left: 16.666666666666664%;
}

.col-xs-push-1 {
  left: 8.333333333333332%;
}

.col-xs-push-0 {
  left: 0;
}

.col-xs-offset-12 {
  margin-left: 100%;
}

.col-xs-offset-11 {
  margin-left: 91.66666666666666%;
}

.col-xs-offset-10 {
  margin-left: 83.33333333333334%;
}

.col-xs-offset-9 {
  margin-left: 75%;
}

.col-xs-offset-8 {
  margin-left: 66.66666666666666%;
}

.col-xs-offset-7 {
  margin-left: 58.333333333333336%;
}

.col-xs-offset-6 {
  margin-left: 50%;
}

.col-xs-offset-5 {
  margin-left: 41.66666666666667%;
}

.col-xs-offset-4 {
  margin-left: 33.33333333333333%;
}

.col-xs-offset-3 {
  margin-left: 25%;
}

.col-xs-offset-2 {
  margin-left: 16.666666666666664%;
}

.col-xs-offset-1 {
  margin-left: 8.333333333333332%;
}

.col-xs-offset-0 {
  margin-left: 0;
}

@media (min-width: 768px) {
  .col-sm-1,
  .col-sm-2,
  .col-sm-3,
  .col-sm-4,
  .col-sm-5,
  .col-sm-6,
  .col-sm-7,
  .col-sm-8,
  .col-sm-9,
  .col-sm-10,
  .col-sm-11,
  .col-sm-12 {
    float: left;
  }
  .col-sm-12 {
    width: 100%;
  }
  .col-sm-11 {
    width: 91.66666666666666%;
  }
  .col-sm-10 {
    width: 83.33333333333334%;
  }
  .col-sm-9 {
    width: 75%;
  }
  .col-sm-8 {
    width: 66.66666666666666%;
  }
  .col-sm-7 {
    width: 58.333333333333336%;
  }
  .col-sm-6 {
    width: 50%;
  }
  .col-sm-5 {
    width: 41.66666666666667%;
  }
  .col-sm-4 {
    width: 33.33333333333333%;
  }
  .col-sm-3 {
    width: 25%;
  }
  .col-sm-2 {
    width: 16.666666666666664%;
  }
  .col-sm-1 {
    width: 8.333333333333332%;
  }
  .col-sm-pull-12 {
    right: 100%;
  }
  .col-sm-pull-11 {
    right: 91.66666666666666%;
  }
  .col-sm-pull-10 {
    right: 83.33333333333334%;
  }
  .col-sm-pull-9 {
    right: 75%;
  }
  .col-sm-pull-8 {
    right: 66.66666666666666%;
  }
  .col-sm-pull-7 {
    right: 58.333333333333336%;
  }
  .col-sm-pull-6 {
    right: 50%;
  }
  .col-sm-pull-5 {
    right: 41.66666666666667%;
  }
  .col-sm-pull-4 {
    right: 33.33333333333333%;
  }
  .col-sm-pull-3 {
    right: 25%;
  }
  .col-sm-pull-2 {
    right: 16.666666666666664%;
  }
  .col-sm-pull-1 {
    right: 8.333333333333332%;
  }
  .col-sm-pull-0 {
    right: 0;
  }
  .col-sm-push-12 {
    left: 100%;
  }
  .col-sm-push-11 {
    left: 91.66666666666666%;
  }
  .col-sm-push-10 {
    left: 83.33333333333334%;
  }
  .col-sm-push-9 {
    left: 75%;
  }
  .col-sm-push-8 {
    left: 66.66666666666666%;
  }
  .col-sm-push-7 {
    left: 58.333333333333336%;
  }
  .col-sm-push-6 {
    left: 50%;
  }
  .col-sm-push-5 {
    left: 41.66666666666667%;
  }
  .col-sm-push-4 {
    left: 33.33333333333333%;
  }
  .col-sm-push-3 {
    left: 25%;
  }
  .col-sm-push-2 {
    left: 16.666666666666664%;
  }
  .col-sm-push-1 {
    left: 8.333333333333332%;
  }
  .col-sm-push-0 {
    left: 0;
  }
  .col-sm-offset-12 {
    margin-left: 100%;
  }
  .col-sm-offset-11 {
    margin-left: 91.66666666666666%;
  }
  .col-sm-offset-10 {
    margin-left: 83.33333333333334%;
  }
  .col-sm-offset-9 {
    margin-left: 75%;
  }
  .col-sm-offset-8 {
    margin-left: 66.66666666666666%;
  }
  .col-sm-offset-7 {
    margin-left: 58.333333333333336%;
  }
  .col-sm-offset-6 {
    margin-left: 50%;
  }
  .col-sm-offset-5 {
    margin-left: 41.66666666666667%;
  }
  .col-sm-offset-4 {
    margin-left: 33.33333333333333%;
  }
  .col-sm-offset-3 {
    margin-left: 25%;
  }
  .col-sm-offset-2 {
    margin-left: 16.666666666666664%;
  }
  .col-sm-offset-1 {
    margin-left: 8.333333333333332%;
  }
  .col-sm-offset-0 {
    margin-left: 0;
  }
}

@media (min-width: 992px) {
  .col-md-1,
  .col-md-2,
  .col-md-3,
  .col-md-4,
  .col-md-5,
  .col-md-6,
  .col-md-7,
  .col-md-8,
  .col-md-9,
  .col-md-10,
  .col-md-11,
  .col-md-12 {
    float: left;
  }
  .col-md-12 {
    width: 100%;
  }
  .col-md-11 {
    width: 91.66666666666666%;
  }
  .col-md-10 {
    width: 83.33333333333334%;
  }
  .col-md-9 {
    width: 75%;
  }
  .col-md-8 {
    width: 66.66666666666666%;
  }
  .col-md-7 {
    width: 58.333333333333336%;
  }
  .col-md-6 {
    width: 50%;
  }
  .col-md-5 {
    width: 41.66666666666667%;
  }
  .col-md-4 {
    width: 33.33333333333333%;
  }
  .col-md-3 {
    width: 25%;
  }
  .col-md-2 {
    width: 16.666666666666664%;
  }
  .col-md-1 {
    width: 8.333333333333332%;
  }
  .col-md-pull-12 {
    right: 100%;
  }
  .col-md-pull-11 {
    right: 91.66666666666666%;
  }
  .col-md-pull-10 {
    right: 83.33333333333334%;
  }
  .col-md-pull-9 {
    right: 75%;
  }
  .col-md-pull-8 {
    right: 66.66666666666666%;
  }
  .col-md-pull-7 {
    right: 58.333333333333336%;
  }
  .col-md-pull-6 {
    right: 50%;
  }
  .col-md-pull-5 {
    right: 41.66666666666667%;
  }
  .col-md-pull-4 {
    right: 33.33333333333333%;
  }
  .col-md-pull-3 {
    right: 25%;
  }
  .col-md-pull-2 {
    right: 16.666666666666664%;
  }
  .col-md-pull-1 {
    right: 8.333333333333332%;
  }
  .col-md-pull-0 {
    right: 0;
  }
  .col-md-push-12 {
    left: 100%;
  }
  .col-md-push-11 {
    left: 91.66666666666666%;
  }
  .col-md-push-10 {
    left: 83.33333333333334%;
  }
  .col-md-push-9 {
    left: 75%;
  }
  .col-md-push-8 {
    left: 66.66666666666666%;
  }
  .col-md-push-7 {
    left: 58.333333333333336%;
  }
  .col-md-push-6 {
    left: 50%;
  }
  .col-md-push-5 {
    left: 41.66666666666667%;
  }
  .col-md-push-4 {
    left: 33.33333333333333%;
  }
  .col-md-push-3 {
    left: 25%;
  }
  .col-md-push-2 {
    left: 16.666666666666664%;
  }
  .col-md-push-1 {
    left: 8.333333333333332%;
  }
  .col-md-push-0 {
    left: 0;
  }
  .col-md-offset-12 {
    margin-left: 100%;
  }
  .col-md-offset-11 {
    margin-left: 91.66666666666666%;
  }
  .col-md-offset-10 {
    margin-left: 83.33333333333334%;
  }
  .col-md-offset-9 {
    margin-left: 75%;
  }
  .col-md-offset-8 {
    margin-left: 66.66666666666666%;
  }
  .col-md-offset-7 {
    margin-left: 58.333333333333336%;
  }
  .col-md-offset-6 {
    margin-left: 50%;
  }
  .col-md-offset-5 {
    margin-left: 41.66666666666667%;
  }
  .col-md-offset-4 {
    margin-left: 33.33333333333333%;
  }
  .col-md-offset-3 {
    margin-left: 25%;
  }
  .col-md-offset-2 {
    margin-left: 16.666666666666664%;
  }
  .col-md-offset-1 {
    margin-left: 8.333333333333332%;
  }
  .col-md-offset-0 {
    margin-left: 0;
  }
}

@media (min-width: 1200px) {
  .col-lg-1,
  .col-lg-2,
  .col-lg-3,
  .col-lg-4,
  .col-lg-5,
  .col-lg-6,
  .col-lg-7,
  .col-lg-8,
  .col-lg-9,
  .col-lg-10,
  .col-lg-11,
  .col-lg-12 {
    float: left;
  }
  .col-lg-12 {
    width: 100%;
  }
  .col-lg-11 {
    width: 91.66666666666666%;
  }
  .col-lg-10 {
    width: 83.33333333333334%;
  }
  .col-lg-9 {
    width: 75%;
  }
  .col-lg-8 {
    width: 66.66666666666666%;
  }
  .col-lg-7 {
    width: 58.333333333333336%;
  }
  .col-lg-6 {
    width: 50%;
  }
  .col-lg-5 {
    width: 41.66666666666667%;
  }
  .col-lg-4 {
    width: 33.33333333333333%;
  }
  .col-lg-3 {
    width: 25%;
  }
  .col-lg-2 {
    width: 16.666666666666664%;
  }
  .col-lg-1 {
    width: 8.333333333333332%;
  }
  .col-lg-pull-12 {
    right: 100%;
  }
  .col-lg-pull-11 {
    right: 91.66666666666666%;
  }
  .col-lg-pull-10 {
    right: 83.33333333333334%;
  }
  .col-lg-pull-9 {
    right: 75%;
  }
  .col-lg-pull-8 {
    right: 66.66666666666666%;
  }
  .col-lg-pull-7 {
    right: 58.333333333333336%;
  }
  .col-lg-pull-6 {
    right: 50%;
  }
  .col-lg-pull-5 {
    right: 41.66666666666667%;
  }
  .col-lg-pull-4 {
    right: 33.33333333333333%;
  }
  .col-lg-pull-3 {
    right: 25%;
  }
  .col-lg-pull-2 {
    right: 16.666666666666664%;
  }
  .col-lg-pull-1 {
    right: 8.333333333333332%;
  }
  .col-lg-pull-0 {
    right: 0;
  }
  .col-lg-push-12 {
    left: 100%;
  }
  .col-lg-push-11 {
    left: 91.66666666666666%;
  }
  .col-lg-push-10 {
    left: 83.33333333333334%;
  }
  .col-lg-push-9 {
    left: 75%;
  }
  .col-lg-push-8 {
    left: 66.66666666666666%;
  }
  .col-lg-push-7 {
    left: 58.333333333333336%;
  }
  .col-lg-push-6 {
    left: 50%;
  }
  .col-lg-push-5 {
    left: 41.66666666666667%;
  }
  .col-lg-push-4 {
    left: 33.33333333333333%;
  }
  .col-lg-push-3 {
    left: 25%;
  }
  .col-lg-push-2 {
    left: 16.666666666666664%;
  }
  .col-lg-push-1 {
    left: 8.333333333333332%;
  }
  .col-lg-push-0 {
    left: 0;
  }
  .col-lg-offset-12 {
    margin-left: 100%;
  }
  .col-lg-offset-11 {
    margin-left: 91.66666666666666%;
  }
  .col-lg-offset-10 {
    margin-left: 83.33333333333334%;
  }
  .col-lg-offset-9 {
    margin-left: 75%;
  }
  .col-lg-offset-8 {
    margin-left: 66.66666666666666%;
  }
  .col-lg-offset-7 {
    margin-left: 58.333333333333336%;
  }
  .col-lg-offset-6 {
    margin-left: 50%;
  }
  .col-lg-offset-5 {
    margin-left: 41.66666666666667%;
  }
  .col-lg-offset-4 {
    margin-left: 33.33333333333333%;
  }
  .col-lg-offset-3 {
    margin-left: 25%;
  }
  .col-lg-offset-2 {
    margin-left: 16.666666666666664%;
  }
  .col-lg-offset-1 {
    margin-left: 8.333333333333332%;
  }
  .col-lg-offset-0 {
    margin-left: 0;
  }
}

table {
  max-width: 100%;
  background-color: #2e3338;
}

th {
  text-align: left;
}

.table {
  width: 100%;
  margin-bottom: 20px;
}

.table > thead > tr > th,
.table > tbody > tr > th,
.table > tfoot > tr > th,
.table > thead > tr > td,
.table > tbody > tr > td,
.table > tfoot > tr > td {
  padding: 8px;
  line-height: 1.428571429;
  vertical-align: top;
  border-top: 1px solid #1c1e22;
}

.table > thead > tr > th {
  vertical-align: bottom;
  border-bottom: 2px solid #1c1e22;
}

.table > caption + thead > tr:first-child > th,
.table > colgroup + thead > tr:first-child > th,
.table > thead:first-child > tr:first-child > th,
.table > caption + thead > tr:first-child > td,
.table > colgroup + thead > tr:first-child > td,
.table > thead:first-child > tr:first-child > td {
  border-top: 0;
}

.table > tbody + tbody {
  border-top: 2px solid #1c1e22;
}

.table .table {
  background-color: #272b30;
}

.table-condensed > thead > tr > th,
.table-condensed > tbody > tr > th,
.table-condensed > tfoot > tr > th,
.table-condensed > thead > tr > td,
.table-condensed > tbody > tr > td,
.table-condensed > tfoot > tr > td {
  padding: 5px;
}

.table-bordered {
  border: 1px solid #1c1e22;
}

.table-bordered > thead > tr > th,
.table-bordered > tbody > tr > th,
.table-bordered > tfoot > tr > th,
.table-bordered > thead > tr > td,
.table-bordered > tbody > tr > td,
.table-bordered > tfoot > tr > td {
  border: 1px solid #1c1e22;
}

.table-bordered > thead > tr > th,
.table-bordered > thead > tr > td {
  border-bottom-width: 2px;
}

.table-striped > tbody > tr:nth-child(odd) > td,
.table-striped > tbody > tr:nth-child(odd) > th {
  background-color: #353a41;
}

.table-hover > tbody > tr:hover > td,
.table-hover > tbody > tr:hover > th {
  background-color: #49515a;
}

table col[class*="col-"] {
  position: static;
  display: table-column;
  float: none;
}

table td[class*="col-"],
table th[class*="col-"] {
  display: table-cell;
  float: none;
}

.table > thead > tr > .active,
.table > tbody > tr > .active,
.table > tfoot > tr > .active,
.table > thead > .active > td,
.table > tbody > .active > td,
.table > tfoot > .active > td,
.table > thead > .active > th,
.table > tbody > .active > th,
.table > tfoot > .active > th {
  background-color: #49515a;
}

.table-hover > tbody > tr > .active:hover,
.table-hover > tbody > .active:hover > td,
.table-hover > tbody > .active:hover > th {
  background-color: #3e444c;
}

.table > thead > tr > .success,
.table > tbody > tr > .success,
.table > tfoot > tr > .success,
.table > thead > .success > td,
.table > tbody > .success > td,
.table > tfoot > .success > td,
.table > thead > .success > th,
.table > tbody > .success > th,
.table > tfoot > .success > th {
  background-color: #62c462;
}

.table-hover > tbody > tr > .success:hover,
.table-hover > tbody > .success:hover > td,
.table-hover > tbody > .success:hover > th {
  background-color: #4fbd4f;
}

.table > thead > tr > .danger,
.table > tbody > tr > .danger,
.table > tfoot > tr > .danger,
.table > thead > .danger > td,
.table > tbody > .danger > td,
.table > tfoot > .danger > td,
.table > thead > .danger > th,
.table > tbody > .danger > th,
.table > tfoot > .danger > th {
  background-color: #ee5f5b;
}

.table-hover > tbody > tr > .danger:hover,
.table-hover > tbody > .danger:hover > td,
.table-hover > tbody > .danger:hover > th {
  background-color: #ec4844;
}

.table > thead > tr > .warning,
.table > tbody > tr > .warning,
.table > tfoot > tr > .warning,
.table > thead > .warning > td,
.table > tbody > .warning > td,
.table > tfoot > .warning > td,
.table > thead > .warning > th,
.table > tbody > .warning > th,
.table > tfoot > .warning > th {
  background-color: #f89406;
}

.table-hover > tbody > tr > .warning:hover,
.table-hover > tbody > .warning:hover > td,
.table-hover > tbody > .warning:hover > th {
  background-color: #df8505;
}

@media (max-width: 767px) {
  .table-responsive {
    width: 100%;
    margin-bottom: 15px;
    overflow-x: scroll;
    overflow-y: hidden;
    border: 1px solid #1c1e22;
    -ms-overflow-style: -ms-autohiding-scrollbar;
    -webkit-overflow-scrolling: touch;
  }
  .table-responsive > .table {
    margin-bottom: 0;
  }
  .table-responsive > .table > thead > tr > th,
  .table-responsive > .table > tbody > tr > th,
  .table-responsive > .table > tfoot > tr > th,
  .table-responsive > .table > thead > tr > td,
  .table-responsive > .table > tbody > tr > td,
  .table-responsive > .table > tfoot > tr > td {
    white-space: nowrap;
  }
  .table-responsive > .table-bordered {
    border: 0;
  }
  .table-responsive > .table-bordered > thead > tr > th:first-child,
  .table-responsive > .table-bordered > tbody > tr > th:first-child,
  .table-responsive > .table-bordered > tfoot > tr > th:first-child,
  .table-responsive > .table-bordered > thead > tr > td:first-child,
  .table-responsive > .table-bordered > tbody > tr > td:first-child,
  .table-responsive > .table-bordered > tfoot > tr > td:first-child {
    border-left: 0;
  }
  .table-responsive > .table-bordered > thead > tr > th:last-child,
  .table-responsive > .table-bordered > tbody > tr > th:last-child,
  .table-responsive > .table-bordered > tfoot > tr > th:last-child,
  .table-responsive > .table-bordered > thead > tr > td:last-child,
  .table-responsive > .table-bordered > tbody > tr > td:last-child,
  .table-responsive > .table-bordered > tfoot > tr > td:last-child {
    border-right: 0;
  }
  .table-responsive > .table-bordered > tbody > tr:last-child > th,
  .table-responsive > .table-bordered > tfoot > tr:last-child > th,
  .table-responsive > .table-bordered > tbody > tr:last-child > td,
  .table-responsive > .table-bordered > tfoot > tr:last-child > td {
    border-bottom: 0;
  }
}

fieldset {
  padding: 0;
  margin: 0;
  border: 0;
}

legend {
  display: block;
  width: 100%;
  padding: 0;
  margin-bottom: 20px;
  font-size: 21px;
  line-height: inherit;
  color: #c8c8c8;
  border: 0;
  border-bottom: 1px solid #1c1e22;
}

label {
  display: inline-block;
  margin-bottom: 5px;
  font-weight: bold;
}

input[type="search"] {
  -webkit-box-sizing: border-box;
     -moz-box-sizing: border-box;
          box-sizing: border-box;
}

input[type="radio"],
input[type="checkbox"] {
  margin: 4px 0 0;
  margin-top: 1px \9;
  /* IE8-9 */

  line-height: normal;
}

input[type="file"] {
  display: block;
}

select[multiple],
select[size] {
  height: auto;
}

select optgroup {
  font-family: inherit;
  font-size: inherit;
  font-style: inherit;
}

input[type="file"]:focus,
input[type="radio"]:focus,
input[type="checkbox"]:focus {
  outline: thin dotted;
  outline: 5px auto -webkit-focus-ring-color;
  outline-offset: -2px;
}

input[type="number"]::-webkit-outer-spin-button,
input[type="number"]::-webkit-inner-spin-button {
  height: auto;
}

output {
  display: block;
  padding-top: 9px;
  font-size: 14px;
  line-height: 1.428571429;
  color: #272b30;
  vertical-align: middle;
}

.form-control {
  display: block;
  width: 100%;
  height: 38px;
  padding: 8px 12px;
  font-size: 14px;
  line-height: 1.428571429;
  color: #272b30;
  vertical-align: middle;
  background-color: #ffffff;
  background-image: none;
  border: 1px solid #cccccc;
  border-radius: 4px;
  -webkit-box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075);
          box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075);
  -webkit-transition: border-color ease-in-out 0.15s, box-shadow ease-in-out 0.15s;
          transition: border-color ease-in-out 0.15s, box-shadow ease-in-out 0.15s;
}

.form-control:focus {
  border-color: #66afe9;
  outline: 0;
  -webkit-box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075), 0 0 8px rgba(102, 175, 233, 0.6);
          box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075), 0 0 8px rgba(102, 175, 233, 0.6);
}

.form-control:-moz-placeholder {
  color: #7a8288;
}

.form-control::-moz-placeholder {
  color: #7a8288;
  opacity: 1;
}

.form-control:-ms-input-placeholder {
  color: #7a8288;
}

.form-control::-webkit-input-placeholder {
  color: #7a8288;
}

.form-control[disabled],
.form-control[readonly],
fieldset[disabled] .form-control {
  cursor: not-allowed;
  background-color: #999999;
}

textarea.form-control {
  height: auto;
}

.form-group {
  margin-bottom: 15px;
}

.radio,
.checkbox {
  display: block;
  min-height: 20px;
  padding-left: 20px;
  margin-top: 10px;
  margin-bottom: 10px;
  vertical-align: middle;
}

.radio label,
.checkbox label {
  display: inline;
  margin-bottom: 0;
  font-weight: normal;
  cursor: pointer;
}

.radio input[type="radio"],
.radio-inline input[type="radio"],
.checkbox input[type="checkbox"],
.checkbox-inline input[type="checkbox"] {
  float: left;
  margin-left: -20px;
}

.radio + .radio,
.checkbox + .checkbox {
  margin-top: -5px;
}

.radio-inline,
.checkbox-inline {
  display: inline-block;
  padding-left: 20px;
  margin-bottom: 0;
  font-weight: normal;
  vertical-align: middle;
  cursor: pointer;
}

.radio-inline + .radio-inline,
.checkbox-inline + .checkbox-inline {
  margin-top: 0;
  margin-left: 10px;
}

input[type="radio"][disabled],
input[type="checkbox"][disabled],
.radio[disabled],
.radio-inline[disabled],
.checkbox[disabled],
.checkbox-inline[disabled],
fieldset[disabled] input[type="radio"],
fieldset[disabled] input[type="checkbox"],
fieldset[disabled] .radio,
fieldset[disabled] .radio-inline,
fieldset[disabled] .checkbox,
fieldset[disabled] .checkbox-inline {
  cursor: not-allowed;
}

.input-sm {
  height: 30px;
  padding: 5px 10px;
  font-size: 12px;
  line-height: 1.5;
  border-radius: 3px;
}

select.input-sm {
  height: 30px;
  line-height: 30px;
}

textarea.input-sm {
  height: auto;
}

.input-lg {
  height: 56px;
  padding: 14px 16px;
  font-size: 18px;
  line-height: 1.33;
  border-radius: 6px;
}

select.input-lg {
  height: 56px;
  line-height: 56px;
}

textarea.input-lg {
  height: auto;
}

.has-warning .help-block,
.has-warning .control-label,
.has-warning .radio,
.has-warning .checkbox,
.has-warning .radio-inline,
.has-warning .checkbox-inline {
  color: #ffffff;
}

.has-warning .form-control {
  border-color: #ffffff;
  -webkit-box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075);
          box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075);
}

.has-warning .form-control:focus {
  border-color: #e6e6e6;
  -webkit-box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075), 0 0 6px #ffffff;
          box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075), 0 0 6px #ffffff;
}

.has-warning .input-group-addon {
  color: #ffffff;
  background-color: #f89406;
  border-color: #ffffff;
}

.has-error .help-block,
.has-error .control-label,
.has-error .radio,
.has-error .checkbox,
.has-error .radio-inline,
.has-error .checkbox-inline {
  color: #ffffff;
}

.has-error .form-control {
  border-color: #ffffff;
  -webkit-box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075);
          box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075);
}

.has-error .form-control:focus {
  border-color: #e6e6e6;
  -webkit-box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075), 0 0 6px #ffffff;
          box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075), 0 0 6px #ffffff;
}

.has-error .input-group-addon {
  color: #ffffff;
  background-color: #ee5f5b;
  border-color: #ffffff;
}

.has-success .help-block,
.has-success .control-label,
.has-success .radio,
.has-success .checkbox,
.has-success .radio-inline,
.has-success .checkbox-inline {
  color: #ffffff;
}

.has-success .form-control {
  border-color: #ffffff;
  -webkit-box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075);
          box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075);
}

.has-success .form-control:focus {
  border-color: #e6e6e6;
  -webkit-box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075), 0 0 6px #ffffff;
          box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075), 0 0 6px #ffffff;
}

.has-success .input-group-addon {
  color: #ffffff;
  background-color: #62c462;
  border-color: #ffffff;
}

.form-control-static {
  margin-bottom: 0;
}

.help-block {
  display: block;
  margin-top: 5px;
  margin-bottom: 10px;
  color: #ffffff;
}

@media (min-width: 768px) {
  .form-inline .form-group {
    display: inline-block;
    margin-bottom: 0;
    vertical-align: middle;
  }
  .form-inline .form-control {
    display: inline-block;
  }
  .form-inline select.form-control {
    width: auto;
  }
  .form-inline .radio,
  .form-inline .checkbox {
    display: inline-block;
    padding-left: 0;
    margin-top: 0;
    margin-bottom: 0;
  }
  .form-inline .radio input[type="radio"],
  .form-inline .checkbox input[type="checkbox"] {
    float: none;
    margin-left: 0;
  }
}

.form-horizontal .control-label,
.form-horizontal .radio,
.form-horizontal .checkbox,
.form-horizontal .radio-inline,
.form-horizontal .checkbox-inline {
  padding-top: 9px;
  margin-top: 0;
  margin-bottom: 0;
}

.form-horizontal .radio,
.form-horizontal .checkbox {
  min-height: 29px;
}

.form-horizontal .form-group {
  margin-right: -15px;
  margin-left: -15px;
}

.form-horizontal .form-group:before,
.form-horizontal .form-group:after {
  display: table;
  content: " ";
}

.form-horizontal .form-group:after {
  clear: both;
}

.form-horizontal .form-group:before,
.form-horizontal .form-group:after {
  display: table;
  content: " ";
}

.form-horizontal .form-group:after {
  clear: both;
}

.form-horizontal .form-group:before,
.form-horizontal .form-group:after {
  display: table;
  content: " ";
}

.form-horizontal .form-group:after {
  clear: both;
}

.form-horizontal .form-group:before,
.form-horizontal .form-group:after {
  display: table;
  content: " ";
}

.form-horizontal .form-group:after {
  clear: both;
}

.form-horizontal .form-group:before,
.form-horizontal .form-group:after {
  display: table;
  content: " ";
}

.form-horizontal .form-group:after {
  clear: both;
}

.form-horizontal .form-control-static {
  padding-top: 9px;
}

@media (min-width: 768px) {
  .form-horizontal .control-label {
    text-align: right;
  }
}

.btn {
  display: inline-block;
  padding: 8px 12px;
  margin-bottom: 0;
  font-size: 14px;
  font-weight: normal;
  line-height: 1.428571429;
  text-align: center;
  white-space: nowrap;
  vertical-align: middle;
  cursor: pointer;
  background-image: none;
  border: 1px solid transparent;
  border-radius: 4px;
  -webkit-user-select: none;
     -moz-user-select: none;
      -ms-user-select: none;
       -o-user-select: none;
          user-select: none;
}

.btn:focus {
  outline: thin dotted;
  outline: 5px auto -webkit-focus-ring-color;
  outline-offset: -2px;
}

.btn:hover,
.btn:focus {
  color: #ffffff;
  text-decoration: none;
}

.btn:active,
.btn.active {
  background-image: none;
  outline: 0;
  -webkit-box-shadow: inset 0 3px 5px rgba(0, 0, 0, 0.125);
          box-shadow: inset 0 3px 5px rgba(0, 0, 0, 0.125);
}

.btn.disabled,
.btn[disabled],
fieldset[disabled] .btn {
  pointer-events: none;
  cursor: not-allowed;
  opacity: 0.65;
  filter: alpha(opacity=65);
  -webkit-box-shadow: none;
          box-shadow: none;
}

.btn-default {
  color: #ffffff;
  background-color: #3a3f44;
  border-color: #3a3f44;
}

.btn-default:hover,
.btn-default:focus,
.btn-default:active,
.btn-default.active,
.open .dropdown-toggle.btn-default {
  color: #ffffff;
  background-color: #272b2e;
  border-color: #1e2023;
}

.btn-default:active,
.btn-default.active,
.open .dropdown-toggle.btn-default {
  background-image: none;
}

.btn-default.disabled,
.btn-default[disabled],
fieldset[disabled] .btn-default,
.btn-default.disabled:hover,
.btn-default[disabled]:hover,
fieldset[disabled] .btn-default:hover,
.btn-default.disabled:focus,
.btn-default[disabled]:focus,
fieldset[disabled] .btn-default:focus,
.btn-default.disabled:active,
.btn-default[disabled]:active,
fieldset[disabled] .btn-default:active,
.btn-default.disabled.active,
.btn-default[disabled].active,
fieldset[disabled] .btn-default.active {
  background-color: #3a3f44;
  border-color: #3a3f44;
}

.btn-default .badge {
  color: #3a3f44;
  background-color: #fff;
}

.btn-primary {
  color: #ffffff;
  background-color: #7a8288;
  border-color: #7a8288;
}

.btn-primary:hover,
.btn-primary:focus,
.btn-primary:active,
.btn-primary.active,
.open .dropdown-toggle.btn-primary {
  color: #ffffff;
  background-color: #676d73;
  border-color: #5d6368;
}

.btn-primary:active,
.btn-primary.active,
.open .dropdown-toggle.btn-primary {
  background-image: none;
}

.btn-primary.disabled,
.btn-primary[disabled],
fieldset[disabled] .btn-primary,
.btn-primary.disabled:hover,
.btn-primary[disabled]:hover,
fieldset[disabled] .btn-primary:hover,
.btn-primary.disabled:focus,
.btn-primary[disabled]:focus,
fieldset[disabled] .btn-primary:focus,
.btn-primary.disabled:active,
.btn-primary[disabled]:active,
fieldset[disabled] .btn-primary:active,
.btn-primary.disabled.active,
.btn-primary[disabled].active,
fieldset[disabled] .btn-primary.active {
  background-color: #7a8288;
  border-color: #7a8288;
}

.btn-primary .badge {
  color: #7a8288;
  background-color: #fff;
}

.btn-warning {
  color: #ffffff;
  background-color: #f89406;
  border-color: #f89406;
}

.btn-warning:hover,
.btn-warning:focus,
.btn-warning:active,
.btn-warning.active,
.open .dropdown-toggle.btn-warning {
  color: #ffffff;
  background-color: #d07c05;
  border-color: #bc7005;
}

.btn-warning:active,
.btn-warning.active,
.open .dropdown-toggle.btn-warning {
  background-image: none;
}

.btn-warning.disabled,
.btn-warning[disabled],
fieldset[disabled] .btn-warning,
.btn-warning.disabled:hover,
.btn-warning[disabled]:hover,
fieldset[disabled] .btn-warning:hover,
.btn-warning.disabled:focus,
.btn-warning[disabled]:focus,
fieldset[disabled] .btn-warning:focus,
.btn-warning.disabled:active,
.btn-warning[disabled]:active,
fieldset[disabled] .btn-warning:active,
.btn-warning.disabled.active,
.btn-warning[disabled].active,
fieldset[disabled] .btn-warning.active {
  background-color: #f89406;
  border-color: #f89406;
}

.btn-warning .badge {
  color: #f89406;
  background-color: #fff;
}

.btn-danger {
  color: #ffffff;
  background-color: #ee5f5b;
  border-color: #ee5f5b;
}

.btn-danger:hover,
.btn-danger:focus,
.btn-danger:active,
.btn-danger.active,
.open .dropdown-toggle.btn-danger {
  color: #ffffff;
  background-color: #ea3b36;
  border-color: #e82924;
}

.btn-danger:active,
.btn-danger.active,
.open .dropdown-toggle.btn-danger {
  background-image: none;
}

.btn-danger.disabled,
.btn-danger[disabled],
fieldset[disabled] .btn-danger,
.btn-danger.disabled:hover,
.btn-danger[disabled]:hover,
fieldset[disabled] .btn-danger:hover,
.btn-danger.disabled:focus,
.btn-danger[disabled]:focus,
fieldset[disabled] .btn-danger:focus,
.btn-danger.disabled:active,
.btn-danger[disabled]:active,
fieldset[disabled] .btn-danger:active,
.btn-danger.disabled.active,
.btn-danger[disabled].active,
fieldset[disabled] .btn-danger.active {
  background-color: #ee5f5b;
  border-color: #ee5f5b;
}

.btn-danger .badge {
  color: #ee5f5b;
  background-color: #fff;
}

.btn-success {
  color: #ffffff;
  background-color: #62c462;
  border-color: #62c462;
}

.btn-success:hover,
.btn-success:focus,
.btn-success:active,
.btn-success.active,
.open .dropdown-toggle.btn-success {
  color: #ffffff;
  background-color: #45b845;
  border-color: #40a940;
}

.btn-success:active,
.btn-success.active,
.open .dropdown-toggle.btn-success {
  background-image: none;
}

.btn-success.disabled,
.btn-success[disabled],
fieldset[disabled] .btn-success,
.btn-success.disabled:hover,
.btn-success[disabled]:hover,
fieldset[disabled] .btn-success:hover,
.btn-success.disabled:focus,
.btn-success[disabled]:focus,
fieldset[disabled] .btn-success:focus,
.btn-success.disabled:active,
.btn-success[disabled]:active,
fieldset[disabled] .btn-success:active,
.btn-success.disabled.active,
.btn-success[disabled].active,
fieldset[disabled] .btn-success.active {
  background-color: #62c462;
  border-color: #62c462;
}

.btn-success .badge {
  color: #62c462;
  background-color: #fff;
}

.btn-info {
  color: #ffffff;
  background-color: #5bc0de;
  border-color: #5bc0de;
}

.btn-info:hover,
.btn-info:focus,
.btn-info:active,
.btn-info.active,
.open .dropdown-toggle.btn-info {
  color: #ffffff;
  background-color: #39b3d7;
  border-color: #2aabd2;
}

.btn-info:active,
.btn-info.active,
.open .dropdown-toggle.btn-info {
  background-image: none;
}

.btn-info.disabled,
.btn-info[disabled],
fieldset[disabled] .btn-info,
.btn-info.disabled:hover,
.btn-info[disabled]:hover,
fieldset[disabled] .btn-info:hover,
.btn-info.disabled:focus,
.btn-info[disabled]:focus,
fieldset[disabled] .btn-info:focus,
.btn-info.disabled:active,
.btn-info[disabled]:active,
fieldset[disabled] .btn-info:active,
.btn-info.disabled.active,
.btn-info[disabled].active,
fieldset[disabled] .btn-info.active {
  background-color: #5bc0de;
  border-color: #5bc0de;
}

.btn-info .badge {
  color: #5bc0de;
  background-color: #fff;
}

.btn-link {
  font-weight: normal;
  color: #ffffff;
  cursor: pointer;
  border-radius: 0;
}

.btn-link,
.btn-link:active,
.btn-link[disabled],
fieldset[disabled] .btn-link {
  background-color: transparent;
  -webkit-box-shadow: none;
          box-shadow: none;
}

.btn-link,
.btn-link:hover,
.btn-link:focus,
.btn-link:active {
  border-color: transparent;
}

.btn-link:hover,
.btn-link:focus {
  color: #ffffff;
  text-decoration: underline;
  background-color: transparent;
}

.btn-link[disabled]:hover,
fieldset[disabled] .btn-link:hover,
.btn-link[disabled]:focus,
fieldset[disabled] .btn-link:focus {
  color: #7a8288;
  text-decoration: none;
}

.btn-lg {
  padding: 14px 16px;
  font-size: 18px;
  line-height: 1.33;
  border-radius: 6px;
}

.btn-sm {
  padding: 5px 10px;
  font-size: 12px;
  line-height: 1.5;
  border-radius: 3px;
}

.btn-xs {
  padding: 1px 5px;
  font-size: 12px;
  line-height: 1.5;
  border-radius: 3px;
}

.btn-block {
  display: block;
  width: 100%;
  padding-right: 0;
  padding-left: 0;
}

.btn-block + .btn-block {
  margin-top: 5px;
}

input[type="submit"].btn-block,
input[type="reset"].btn-block,
input[type="button"].btn-block {
  width: 100%;
}

.fade {
  opacity: 0;
  -webkit-transition: opacity 0.15s linear;
          transition: opacity 0.15s linear;
}

.fade.in {
  opacity: 1;
}

.collapse {
  display: none;
}

.collapse.in {
  display: block;
}

.collapsing {
  position: relative;
  height: 0;
  overflow: hidden;
  -webkit-transition: height 0.35s ease;
          transition: height 0.35s ease;
}

@font-face {
  font-family: 'Glyphicons Halflings';
  src: url('../fonts/glyphicons-halflings-regular.eot');
  src: url('../fonts/glyphicons-halflings-regular.eot?#iefix') format('embedded-opentype'), url('../fonts/glyphicons-halflings-regular.woff') format('woff'), url('../fonts/glyphicons-halflings-regular.ttf') format('truetype'), url('../fonts/glyphicons-halflings-regular.svg#glyphicons-halflingsregular') format('svg');
}

.glyphicon {
  position: relative;
  top: 1px;
  display: inline-block;
  font-family: 'Glyphicons Halflings';
  -webkit-font-smoothing: antialiased;
  font-style: normal;
  font-weight: normal;
  line-height: 1;
  -moz-osx-font-smoothing: grayscale;
}

.glyphicon:empty {
  width: 1em;
}

.glyphicon-asterisk:before {
  content: "\2a";
}

.glyphicon-plus:before {
  content: "\2b";
}

.glyphicon-euro:before {
  content: "\20ac";
}

.glyphicon-minus:before {
  content: "\2212";
}

.glyphicon-cloud:before {
  content: "\2601";
}

.glyphicon-envelope:before {
  content: "\2709";
}

.glyphicon-pencil:before {
  content: "\270f";
}

.glyphicon-glass:before {
  content: "\e001";
}

.glyphicon-music:before {
  content: "\e002";
}

.glyphicon-search:before {
  content: "\e003";
}

.glyphicon-heart:before {
  content: "\e005";
}

.glyphicon-star:before {
  content: "\e006";
}

.glyphicon-star-empty:before {
  content: "\e007";
}

.glyphicon-user:before {
  content: "\e008";
}

.glyphicon-film:before {
  content: "\e009";
}

.glyphicon-th-large:before {
  content: "\e010";
}

.glyphicon-th:before {
  content: "\e011";
}

.glyphicon-th-list:before {
  content: "\e012";
}

.glyphicon-ok:before {
  content: "\e013";
}

.glyphicon-remove:before {
  content: "\e014";
}

.glyphicon-zoom-in:before {
  content: "\e015";
}

.glyphicon-zoom-out:before {
  content: "\e016";
}

.glyphicon-off:before {
  content: "\e017";
}

.glyphicon-signal:before {
  content: "\e018";
}

.glyphicon-cog:before {
  content: "\e019";
}

.glyphicon-trash:before {
  content: "\e020";
}

.glyphicon-home:before {
  content: "\e021";
}

.glyphicon-file:before {
  content: "\e022";
}

.glyphicon-time:before {
  content: "\e023";
}

.glyphicon-road:before {
  content: "\e024";
}

.glyphicon-download-alt:before {
  content: "\e025";
}

.glyphicon-download:before {
  content: "\e026";
}

.glyphicon-upload:before {
  content: "\e027";
}

.glyphicon-inbox:before {
  content: "\e028";
}

.glyphicon-play-circle:before {
  content: "\e029";
}

.glyphicon-repeat:before {
  content: "\e030";
}

.glyphicon-refresh:before {
  content: "\e031";
}

.glyphicon-list-alt:before {
  content: "\e032";
}

.glyphicon-lock:before {
  content: "\e033";
}

.glyphicon-flag:before {
  content: "\e034";
}

.glyphicon-headphones:before {
  content: "\e035";
}

.glyphicon-volume-off:before {
  content: "\e036";
}

.glyphicon-volume-down:before {
  content: "\e037";
}

.glyphicon-volume-up:before {
  content: "\e038";
}

.glyphicon-qrcode:before {
  content: "\e039";
}

.glyphicon-barcode:before {
  content: "\e040";
}

.glyphicon-tag:before {
  content: "\e041";
}

.glyphicon-tags:before {
  content: "\e042";
}

.glyphicon-book:before {
  content: "\e043";
}

.glyphicon-bookmark:before {
  content: "\e044";
}

.glyphicon-print:before {
  content: "\e045";
}

.glyphicon-camera:before {
  content: "\e046";
}

.glyphicon-font:before {
  content: "\e047";
}

.glyphicon-bold:before {
  content: "\e048";
}

.glyphicon-italic:before {
  content: "\e049";
}

.glyphicon-text-height:before {
  content: "\e050";
}

.glyphicon-text-width:before {
  content: "\e051";
}

.glyphicon-align-left:before {
  content: "\e052";
}

.glyphicon-align-center:before {
  content: "\e053";
}

.glyphicon-align-right:before {
  content: "\e054";
}

.glyphicon-align-justify:before {
  content: "\e055";
}

.glyphicon-list:before {
  content: "\e056";
}

.glyphicon-indent-left:before {
  content: "\e057";
}

.glyphicon-indent-right:before {
  content: "\e058";
}

.glyphicon-facetime-video:before {
  content: "\e059";
}

.glyphicon-picture:before {
  content: "\e060";
}

.glyphicon-map-marker:before {
  content: "\e062";
}

.glyphicon-adjust:before {
  content: "\e063";
}

.glyphicon-tint:before {
  content: "\e064";
}

.glyphicon-edit:before {
  content: "\e065";
}

.glyphicon-share:before {
  content: "\e066";
}

.glyphicon-check:before {
  content: "\e067";
}

.glyphicon-move:before {
  content: "\e068";
}

.glyphicon-step-backward:before {
  content: "\e069";
}

.glyphicon-fast-backward:before {
  content: "\e070";
}

.glyphicon-backward:before {
  content: "\e071";
}

.glyphicon-play:before {
  content: "\e072";
}

.glyphicon-pause:before {
  content: "\e073";
}

.glyphicon-stop:before {
  content: "\e074";
}

.glyphicon-forward:before {
  content: "\e075";
}

.glyphicon-fast-forward:before {
  content: "\e076";
}

.glyphicon-step-forward:before {
  content: "\e077";
}

.glyphicon-eject:before {
  content: "\e078";
}

.glyphicon-chevron-left:before {
  content: "\e079";
}

.glyphicon-chevron-right:before {
  content: "\e080";
}

.glyphicon-plus-sign:before {
  content: "\e081";
}

.glyphicon-minus-sign:before {
  content: "\e082";
}

.glyphicon-remove-sign:before {
  content: "\e083";
}

.glyphicon-ok-sign:before {
  content: "\e084";
}

.glyphicon-question-sign:before {
  content: "\e085";
}

.glyphicon-info-sign:before {
  content: "\e086";
}

.glyphicon-screenshot:before {
  content: "\e087";
}

.glyphicon-remove-circle:before {
  content: "\e088";
}

.glyphicon-ok-circle:before {
  content: "\e089";
}

.glyphicon-ban-circle:before {
  content: "\e090";
}

.glyphicon-arrow-left:before {
  content: "\e091";
}

.glyphicon-arrow-right:before {
  content: "\e092";
}

.glyphicon-arrow-up:before {
  content: "\e093";
}

.glyphicon-arrow-down:before {
  content: "\e094";
}

.glyphicon-share-alt:before {
  content: "\e095";
}

.glyphicon-resize-full:before {
  content: "\e096";
}

.glyphicon-resize-small:before {
  content: "\e097";
}

.glyphicon-exclamation-sign:before {
  content: "\e101";
}

.glyphicon-gift:before {
  content: "\e102";
}

.glyphicon-leaf:before {
  content: "\e103";
}

.glyphicon-fire:before {
  content: "\e104";
}

.glyphicon-eye-open:before {
  content: "\e105";
}

.glyphicon-eye-close:before {
  content: "\e106";
}

.glyphicon-warning-sign:before {
  content: "\e107";
}

.glyphicon-plane:before {
  content: "\e108";
}

.glyphicon-calendar:before {
  content: "\e109";
}

.glyphicon-random:before {
  content: "\e110";
}

.glyphicon-comment:before {
  content: "\e111";
}

.glyphicon-magnet:before {
  content: "\e112";
}

.glyphicon-chevron-up:before {
  content: "\e113";
}

.glyphicon-chevron-down:before {
  content: "\e114";
}

.glyphicon-retweet:before {
  content: "\e115";
}

.glyphicon-shopping-cart:before {
  content: "\e116";
}

.glyphicon-folder-close:before {
  content: "\e117";
}

.glyphicon-folder-open:before {
  content: "\e118";
}

.glyphicon-resize-vertical:before {
  content: "\e119";
}

.glyphicon-resize-horizontal:before {
  content: "\e120";
}

.glyphicon-hdd:before {
  content: "\e121";
}

.glyphicon-bullhorn:before {
  content: "\e122";
}

.glyphicon-bell:before {
  content: "\e123";
}

.glyphicon-certificate:before {
  content: "\e124";
}

.glyphicon-thumbs-up:before {
  content: "\e125";
}

.glyphicon-thumbs-down:before {
  content: "\e126";
}

.glyphicon-hand-right:before {
  content: "\e127";
}

.glyphicon-hand-left:before {
  content: "\e128";
}

.glyphicon-hand-up:before {
  content: "\e129";
}

.glyphicon-hand-down:before {
  content: "\e130";
}

.glyphicon-circle-arrow-right:before {
  content: "\e131";
}

.glyphicon-circle-arrow-left:before {
  content: "\e132";
}

.glyphicon-circle-arrow-up:before {
  content: "\e133";
}

.glyphicon-circle-arrow-down:before {
  content: "\e134";
}

.glyphicon-globe:before {
  content: "\e135";
}

.glyphicon-wrench:before {
  content: "\e136";
}

.glyphicon-tasks:before {
  content: "\e137";
}

.glyphicon-filter:before {
  content: "\e138";
}

.glyphicon-briefcase:before {
  content: "\e139";
}

.glyphicon-fullscreen:before {
  content: "\e140";
}

.glyphicon-dashboard:before {
  content: "\e141";
}

.glyphicon-paperclip:before {
  content: "\e142";
}

.glyphicon-heart-empty:before {
  content: "\e143";
}

.glyphicon-link:before {
  content: "\e144";
}

.glyphicon-phone:before {
  content: "\e145";
}

.glyphicon-pushpin:before {
  content: "\e146";
}

.glyphicon-usd:before {
  content: "\e148";
}

.glyphicon-gbp:before {
  content: "\e149";
}

.glyphicon-sort:before {
  content: "\e150";
}

.glyphicon-sort-by-alphabet:before {
  content: "\e151";
}

.glyphicon-sort-by-alphabet-alt:before {
  content: "\e152";
}

.glyphicon-sort-by-order:before {
  content: "\e153";
}

.glyphicon-sort-by-order-alt:before {
  content: "\e154";
}

.glyphicon-sort-by-attributes:before {
  content: "\e155";
}

.glyphicon-sort-by-attributes-alt:before {
  content: "\e156";
}

.glyphicon-unchecked:before {
  content: "\e157";
}

.glyphicon-expand:before {
  content: "\e158";
}

.glyphicon-collapse-down:before {
  content: "\e159";
}

.glyphicon-collapse-up:before {
  content: "\e160";
}

.glyphicon-log-in:before {
  content: "\e161";
}

.glyphicon-flash:before {
  content: "\e162";
}

.glyphicon-log-out:before {
  content: "\e163";
}

.glyphicon-new-window:before {
  content: "\e164";
}

.glyphicon-record:before {
  content: "\e165";
}

.glyphicon-save:before {
  content: "\e166";
}

.glyphicon-open:before {
  content: "\e167";
}

.glyphicon-saved:before {
  content: "\e168";
}

.glyphicon-import:before {
  content: "\e169";
}

.glyphicon-export:before {
  content: "\e170";
}

.glyphicon-send:before {
  content: "\e171";
}

.glyphicon-floppy-disk:before {
  content: "\e172";
}

.glyphicon-floppy-saved:before {
  content: "\e173";
}

.glyphicon-floppy-remove:before {
  content: "\e174";
}

.glyphicon-floppy-save:before {
  content: "\e175";
}

.glyphicon-floppy-open:before {
  content: "\e176";
}

.glyphicon-credit-card:before {
  content: "\e177";
}

.glyphicon-transfer:before {
  content: "\e178";
}

.glyphicon-cutlery:before {
  content: "\e179";
}

.glyphicon-header:before {
  content: "\e180";
}

.glyphicon-compressed:before {
  content: "\e181";
}

.glyphicon-earphone:before {
  content: "\e182";
}

.glyphicon-phone-alt:before {
  content: "\e183";
}

.glyphicon-tower:before {
  content: "\e184";
}

.glyphicon-stats:before {
  content: "\e185";
}

.glyphicon-sd-video:before {
  content: "\e186";
}

.glyphicon-hd-video:before {
  content: "\e187";
}

.glyphicon-subtitles:before {
  content: "\e188";
}

.glyphicon-sound-stereo:before {
  content: "\e189";
}

.glyphicon-sound-dolby:before {
  content: "\e190";
}

.glyphicon-sound-5-1:before {
  content: "\e191";
}

.glyphicon-sound-6-1:before {
  content: "\e192";
}

.glyphicon-sound-7-1:before {
  content: "\e193";
}

.glyphicon-copyright-mark:before {
  content: "\e194";
}

.glyphicon-registration-mark:before {
  content: "\e195";
}

.glyphicon-cloud-download:before {
  content: "\e197";
}

.glyphicon-cloud-upload:before {
  content: "\e198";
}

.glyphicon-tree-conifer:before {
  content: "\e199";
}

.glyphicon-tree-deciduous:before {
  content: "\e200";
}

.caret {
  display: inline-block;
  width: 0;
  height: 0;
  margin-left: 2px;
  vertical-align: middle;
  border-top: 4px solid;
  border-right: 4px solid transparent;
  border-left: 4px solid transparent;
}

.dropdown {
  position: relative;
}

.dropdown-toggle:focus {
  outline: 0;
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  left: 0;
  z-index: 1000;
  display: none;
  float: left;
  min-width: 160px;
  padding: 5px 0;
  margin: 2px 0 0;
  font-size: 14px;
  list-style: none;
  background-color: #3a3f44;
  border: 1px solid #272b30;
  border: 1px solid rgba(0, 0, 0, 0.15);
  border-radius: 4px;
  -webkit-box-shadow: 0 6px 12px rgba(0, 0, 0, 0.175);
          box-shadow: 0 6px 12px rgba(0, 0, 0, 0.175);
  background-clip: padding-box;
}

.dropdown-menu.pull-right {
  right: 0;
  left: auto;
}

.dropdown-menu .divider {
  height: 1px;
  margin: 9px 0;
  overflow: hidden;
  background-color: #272b30;
}

.dropdown-menu > li > a {
  display: block;
  padding: 3px 20px;
  clear: both;
  font-weight: normal;
  line-height: 1.428571429;
  color: #c8c8c8;
  white-space: nowrap;
}

.dropdown-menu > li > a:hover,
.dropdown-menu > li > a:focus {
  color: #ffffff;
  text-decoration: none;
  background-color: #272b30;
}

.dropdown-menu > .active > a,
.dropdown-menu > .active > a:hover,
.dropdown-menu > .active > a:focus {
  color: #ffffff;
  text-decoration: none;
  background-color: #272b30;
  outline: 0;
}

.dropdown-menu > .disabled > a,
.dropdown-menu > .disabled > a:hover,
.dropdown-menu > .disabled > a:focus {
  color: #7a8288;
}

.dropdown-menu > .disabled > a:hover,
.dropdown-menu > .disabled > a:focus {
  text-decoration: none;
  cursor: not-allowed;
  background-color: transparent;
  background-image: none;
  filter: progid:DXImageTransform.Microsoft.gradient(enabled=false);
}

.open > .dropdown-menu {
  display: block;
}

.open > a {
  outline: 0;
}

.dropdown-header {
  display: block;
  padding: 3px 20px;
  font-size: 12px;
  line-height: 1.428571429;
  color: #7a8288;
}

.dropdown-backdrop {
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  z-index: 990;
}

.pull-right > .dropdown-menu {
  right: 0;
  left: auto;
}

.dropup .caret,
.navbar-fixed-bottom .dropdown .caret {
  border-top: 0;
  border-bottom: 4px solid;
  content: "";
}

.dropup .dropdown-menu,
.navbar-fixed-bottom .dropdown .dropdown-menu {
  top: auto;
  bottom: 100%;
  margin-bottom: 1px;
}

@media (min-width: 768px) {
  .navbar-right .dropdown-menu {
    right: 0;
    left: auto;
  }
}

.btn-group,
.btn-group-vertical {
  position: relative;
  display: inline-block;
  vertical-align: middle;
}

.btn-group > .btn,
.btn-group-vertical > .btn {
  position: relative;
  float: left;
}

.btn-group > .btn:hover,
.btn-group-vertical > .btn:hover,
.btn-group > .btn:focus,
.btn-group-vertical > .btn:focus,
.btn-group > .btn:active,
.btn-group-vertical > .btn:active,
.btn-group > .btn.active,
.btn-group-vertical > .btn.active {
  z-index: 2;
}

.btn-group > .btn:focus,
.btn-group-vertical > .btn:focus {
  outline: none;
}

.btn-group .btn + .btn,
.btn-group .btn + .btn-group,
.btn-group .btn-group + .btn,
.btn-group .btn-group + .btn-group {
  margin-left: -1px;
}

.btn-toolbar:before,
.btn-toolbar:after {
  display: table;
  content: " ";
}

.btn-toolbar:after {
  clear: both;
}

.btn-toolbar:before,
.btn-toolbar:after {
  display: table;
  content: " ";
}

.btn-toolbar:after {
  clear: both;
}

.btn-toolbar:before,
.btn-toolbar:after {
  display: table;
  content: " ";
}

.btn-toolbar:after {
  clear: both;
}

.btn-toolbar:before,
.btn-toolbar:after {
  display: table;
  content: " ";
}

.btn-toolbar:after {
  clear: both;
}

.btn-toolbar:before,
.btn-toolbar:after {
  display: table;
  content: " ";
}

.btn-toolbar:after {
  clear: both;
}

.btn-toolbar .btn-group {
  float: left;
}

.btn-toolbar > .btn + .btn,
.btn-toolbar > .btn-group + .btn,
.btn-toolbar > .btn + .btn-group,
.btn-toolbar > .btn-group + .btn-group {
  margin-left: 5px;
}

.btn-group > .btn:not(:first-child):not(:last-child):not(.dropdown-toggle) {
  border-radius: 0;
}

.btn-group > .btn:first-child {
  margin-left: 0;
}

.btn-group > .btn:first-child:not(:last-child):not(.dropdown-toggle) {
  border-top-right-radius: 0;
  border-bottom-right-radius: 0;
}

.btn-group > .btn:last-child:not(:first-child),
.btn-group > .dropdown-toggle:not(:first-child) {
  border-bottom-left-radius: 0;
  border-top-left-radius: 0;
}

.btn-group > .btn-group {
  float: left;
}

.btn-group > .btn-group:not(:first-child):not(:last-child) > .btn {
  border-radius: 0;
}

.btn-group > .btn-group:first-child > .btn:last-child,
.btn-group > .btn-group:first-child > .dropdown-toggle {
  border-top-right-radius: 0;
  border-bottom-right-radius: 0;
}

.btn-group > .btn-group:last-child > .btn:first-child {
  border-bottom-left-radius: 0;
  border-top-left-radius: 0;
}

.btn-group .dropdown-toggle:active,
.btn-group.open .dropdown-toggle {
  outline: 0;
}

.btn-group-xs > .btn {
  padding: 1px 5px;
  font-size: 12px;
  line-height: 1.5;
  border-radius: 3px;
}

.btn-group-sm > .btn {
  padding: 5px 10px;
  font-size: 12px;
  line-height: 1.5;
  border-radius: 3px;
}

.btn-group-lg > .btn {
  padding: 14px 16px;
  font-size: 18px;
  line-height: 1.33;
  border-radius: 6px;
}

.btn-group > .btn + .dropdown-toggle {
  padding-right: 8px;
  padding-left: 8px;
}

.btn-group > .btn-lg + .dropdown-toggle {
  padding-right: 12px;
  padding-left: 12px;
}

.btn-group.open .dropdown-toggle {
  -webkit-box-shadow: inset 0 3px 5px rgba(0, 0, 0, 0.125);
          box-shadow: inset 0 3px 5px rgba(0, 0, 0, 0.125);
}

.btn-group.open .dropdown-toggle.btn-link {
  -webkit-box-shadow: none;
          box-shadow: none;
}

.btn .caret {
  margin-left: 0;
}

.btn-lg .caret {
  border-width: 5px 5px 0;
  border-bottom-width: 0;
}

.dropup .btn-lg .caret {
  border-width: 0 5px 5px;
}

.btn-group-vertical > .btn,
.btn-group-vertical > .btn-group,
.btn-group-vertical > .btn-group > .btn {
  display: block;
  float: none;
  width: 100%;
  max-width: 100%;
}

.btn-group-vertical > .btn-group:before,
.btn-group-vertical > .btn-group:after {
  display: table;
  content: " ";
}

.btn-group-vertical > .btn-group:after {
  clear: both;
}

.btn-group-vertical > .btn-group:before,
.btn-group-vertical > .btn-group:after {
  display: table;
  content: " ";
}

.btn-group-vertical > .btn-group:after {
  clear: both;
}

.btn-group-vertical > .btn-group:before,
.btn-group-vertical > .btn-group:after {
  display: table;
  content: " ";
}

.btn-group-vertical > .btn-group:after {
  clear: both;
}

.btn-group-vertical > .btn-group:before,
.btn-group-vertical > .btn-group:after {
  display: table;
  content: " ";
}

.btn-group-vertical > .btn-group:after {
  clear: both;
}

.btn-group-vertical > .btn-group:before,
.btn-group-vertical > .btn-group:after {
  display: table;
  content: " ";
}

.btn-group-vertical > .btn-group:after {
  clear: both;
}

.btn-group-vertical > .btn-group > .btn {
  float: none;
}

.btn-group-vertical > .btn + .btn,
.btn-group-vertical > .btn + .btn-group,
.btn-group-vertical > .btn-group + .btn,
.btn-group-vertical > .btn-group + .btn-group {
  margin-top: -1px;
  margin-left: 0;
}

.btn-group-vertical > .btn:not(:first-child):not(:last-child) {
  border-radius: 0;
}

.btn-group-vertical > .btn:first-child:not(:last-child) {
  border-top-right-radius: 4px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}

.btn-group-vertical > .btn:last-child:not(:first-child) {
  border-top-right-radius: 0;
  border-bottom-left-radius: 4px;
  border-top-left-radius: 0;
}

.btn-group-vertical > .btn-group:not(:first-child):not(:last-child) > .btn {
  border-radius: 0;
}

.btn-group-vertical > .btn-group:first-child > .btn:last-child,
.btn-group-vertical > .btn-group:first-child > .dropdown-toggle {
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}

.btn-group-vertical > .btn-group:last-child > .btn:first-child {
  border-top-right-radius: 0;
  border-top-left-radius: 0;
}

.btn-group-justified {
  display: table;
  width: 100%;
  border-collapse: separate;
  table-layout: fixed;
}

.btn-group-justified > .btn,
.btn-group-justified > .btn-group {
  display: table-cell;
  float: none;
  width: 1%;
}

.btn-group-justified > .btn-group .btn {
  width: 100%;
}

[data-toggle="buttons"] > .btn > input[type="radio"],
[data-toggle="buttons"] > .btn > input[type="checkbox"] {
  display: none;
}

.input-group {
  position: relative;
  display: table;
  border-collapse: separate;
}

.input-group[class*="col-"] {
  float: none;
  padding-right: 0;
  padding-left: 0;
}

.input-group .form-control {
  width: 100%;
  margin-bottom: 0;
}

.input-group-lg > .form-control,
.input-group-lg > .input-group-addon,
.input-group-lg > .input-group-btn > .btn {
  height: 56px;
  padding: 14px 16px;
  font-size: 18px;
  line-height: 1.33;
  border-radius: 6px;
}

select.input-group-lg > .form-control,
select.input-group-lg > .input-group-addon,
select.input-group-lg > .input-group-btn > .btn {
  height: 56px;
  line-height: 56px;
}

textarea.input-group-lg > .form-control,
textarea.input-group-lg > .input-group-addon,
textarea.input-group-lg > .input-group-btn > .btn {
  height: auto;
}

.input-group-sm > .form-control,
.input-group-sm > .input-group-addon,
.input-group-sm > .input-group-btn > .btn {
  height: 30px;
  padding: 5px 10px;
  font-size: 12px;
  line-height: 1.5;
  border-radius: 3px;
}

select.input-group-sm > .form-control,
select.input-group-sm > .input-group-addon,
select.input-group-sm > .input-group-btn > .btn {
  height: 30px;
  line-height: 30px;
}

textarea.input-group-sm > .form-control,
textarea.input-group-sm > .input-group-addon,
textarea.input-group-sm > .input-group-btn > .btn {
  height: auto;
}

.input-group-addon,
.input-group-btn,
.input-group .form-control {
  display: table-cell;
}

.input-group-addon:not(:first-child):not(:last-child),
.input-group-btn:not(:first-child):not(:last-child),
.input-group .form-control:not(:first-child):not(:last-child) {
  border-radius: 0;
}

.input-group-addon,
.input-group-btn {
  width: 1%;
  white-space: nowrap;
  vertical-align: middle;
}

.input-group-addon {
  padding: 8px 12px;
  font-size: 14px;
  font-weight: normal;
  line-height: 1;
  color: #272b30;
  text-align: center;
  background-color: #999999;
  border: 1px solid #cccccc;
  border-radius: 4px;
}

.input-group-addon.input-sm {
  padding: 5px 10px;
  font-size: 12px;
  border-radius: 3px;
}

.input-group-addon.input-lg {
  padding: 14px 16px;
  font-size: 18px;
  border-radius: 6px;
}

.input-group-addon input[type="radio"],
.input-group-addon input[type="checkbox"] {
  margin-top: 0;
}

.input-group .form-control:first-child,
.input-group-addon:first-child,
.input-group-btn:first-child > .btn,
.input-group-btn:first-child > .dropdown-toggle,
.input-group-btn:last-child > .btn:not(:last-child):not(.dropdown-toggle) {
  border-top-right-radius: 0;
  border-bottom-right-radius: 0;
}

.input-group-addon:first-child {
  border-right: 0;
}

.input-group .form-control:last-child,
.input-group-addon:last-child,
.input-group-btn:last-child > .btn,
.input-group-btn:last-child > .dropdown-toggle,
.input-group-btn:first-child > .btn:not(:first-child) {
  border-bottom-left-radius: 0;
  border-top-left-radius: 0;
}

.input-group-addon:last-child {
  border-left: 0;
}

.input-group-btn {
  position: relative;
  white-space: nowrap;
}

.input-group-btn:first-child > .btn {
  margin-right: -1px;
}

.input-group-btn:last-child > .btn {
  margin-left: -1px;
}

.input-group-btn > .btn {
  position: relative;
}

.input-group-btn > .btn + .btn {
  margin-left: -4px;
}

.input-group-btn > .btn:hover,
.input-group-btn > .btn:active {
  z-index: 2;
}

.nav {
  padding-left: 0;
  margin-bottom: 0;
  list-style: none;
}

.nav:before,
.nav:after {
  display: table;
  content: " ";
}

.nav:after {
  clear: both;
}

.nav:before,
.nav:after {
  display: table;
  content: " ";
}

.nav:after {
  clear: both;
}

.nav:before,
.nav:after {
  display: table;
  content: " ";
}

.nav:after {
  clear: both;
}

.nav:before,
.nav:after {
  display: table;
  content: " ";
}

.nav:after {
  clear: both;
}

.nav:before,
.nav:after {
  display: table;
  content: " ";
}

.nav:after {
  clear: both;
}

.nav > li {
  position: relative;
  display: block;
}

.nav > li > a {
  position: relative;
  display: block;
  padding: 10px 15px;
}

.nav > li > a:hover,
.nav > li > a:focus {
  text-decoration: none;
  background-color: #3e444c;
}

.nav > li.disabled > a {
  color: #7a8288;
}

.nav > li.disabled > a:hover,
.nav > li.disabled > a:focus {
  color: #7a8288;
  text-decoration: none;
  cursor: not-allowed;
  background-color: transparent;
}

.nav .open > a,
.nav .open > a:hover,
.nav .open > a:focus {
  background-color: #3e444c;
  border-color: #ffffff;
}

.nav .nav-divider {
  height: 1px;
  margin: 9px 0;
  overflow: hidden;
  background-color: #e5e5e5;
}

.nav > li > a > img {
  max-width: none;
}

.nav-tabs {
  border-bottom: 1px solid #1c1e22;
}

.nav-tabs > li {
  float: left;
  margin-bottom: -1px;
}

.nav-tabs > li > a {
  margin-right: 2px;
  line-height: 1.428571429;
  border: 1px solid transparent;
  border-radius: 4px 4px 0 0;
}

.nav-tabs > li > a:hover {
  border-color: #1c1e22 #1c1e22 #1c1e22;
}

.nav-tabs > li.active > a,
.nav-tabs > li.active > a:hover,
.nav-tabs > li.active > a:focus {
  color: #ffffff;
  cursor: default;
  background-color: #3e444c;
  border: 1px solid #1c1e22;
  border-bottom-color: transparent;
}

.nav-tabs.nav-justified {
  width: 100%;
  border-bottom: 0;
}

.nav-tabs.nav-justified > li {
  float: none;
}

.nav-tabs.nav-justified > li > a {
  margin-bottom: 5px;
  text-align: center;
}

.nav-tabs.nav-justified > .dropdown .dropdown-menu {
  top: auto;
  left: auto;
}

@media (min-width: 768px) {
  .nav-tabs.nav-justified > li {
    display: table-cell;
    width: 1%;
  }
  .nav-tabs.nav-justified > li > a {
    margin-bottom: 0;
  }
}

.nav-tabs.nav-justified > li > a {
  margin-right: 0;
  border-radius: 4px;
}

.nav-tabs.nav-justified > .active > a,
.nav-tabs.nav-justified > .active > a:hover,
.nav-tabs.nav-justified > .active > a:focus {
  border: 1px solid #1c1e22;
}

@media (min-width: 768px) {
  .nav-tabs.nav-justified > li > a {
    border-bottom: 1px solid #1c1e22;
    border-radius: 4px 4px 0 0;
  }
  .nav-tabs.nav-justified > .active > a,
  .nav-tabs.nav-justified > .active > a:hover,
  .nav-tabs.nav-justified > .active > a:focus {
    border-bottom-color: #272b30;
  }
}

.nav-pills > li {
  float: left;
}

.nav-pills > li > a {
  border-radius: 4px;
}

.nav-pills > li + li {
  margin-left: 2px;
}

.nav-pills > li.active > a,
.nav-pills > li.active > a:hover,
.nav-pills > li.active > a:focus {
  color: #ffffff;
  background-color: transparent;
}

.nav-stacked > li {
  float: none;
}

.nav-stacked > li + li {
  margin-top: 2px;
  margin-left: 0;
}

.nav-justified {
  width: 100%;
}

.nav-justified > li {
  float: none;
}

.nav-justified > li > a {
  margin-bottom: 5px;
  text-align: center;
}

.nav-justified > .dropdown .dropdown-menu {
  top: auto;
  left: auto;
}

@media (min-width: 768px) {
  .nav-justified > li {
    display: table-cell;
    width: 1%;
  }
  .nav-justified > li > a {
    margin-bottom: 0;
  }
}

.nav-tabs-justified {
  border-bottom: 0;
}

.nav-tabs-justified > li > a {
  margin-right: 0;
  border-radius: 4px;
}

.nav-tabs-justified > .active > a,
.nav-tabs-justified > .active > a:hover,
.nav-tabs-justified > .active > a:focus {
  border: 1px solid #1c1e22;
}

@media (min-width: 768px) {
  .nav-tabs-justified > li > a {
    border-bottom: 1px solid #1c1e22;
    border-radius: 4px 4px 0 0;
  }
  .nav-tabs-justified > .active > a,
  .nav-tabs-justified > .active > a:hover,
  .nav-tabs-justified > .active > a:focus {
    border-bottom-color: #272b30;
  }
}

.tab-content > .tab-pane {
  display: none;
}

.tab-content > .active {
  display: block;
}

.nav-tabs .dropdown-menu {
  margin-top: -1px;
  border-top-right-radius: 0;
  border-top-left-radius: 0;
}

.navbar {
  position: relative;
  min-height: 50px;
  margin-bottom: 20px;
  border: 1px solid transparent;
}

.navbar:before,
.navbar:after {
  display: table;
  content: " ";
}

.navbar:after {
  clear: both;
}

.navbar:before,
.navbar:after {
  display: table;
  content: " ";
}

.navbar:after {
  clear: both;
}

.navbar:before,
.navbar:after {
  display: table;
  content: " ";
}

.navbar:after {
  clear: both;
}

.navbar:before,
.navbar:after {
  display: table;
  content: " ";
}

.navbar:after {
  clear: both;
}

.navbar:before,
.navbar:after {
  display: table;
  content: " ";
}

.navbar:after {
  clear: both;
}

@media (min-width: 768px) {
  .navbar {
    border-radius: 4px;
  }
}

.navbar-header:before,
.navbar-header:after {
  display: table;
  content: " ";
}

.navbar-header:after {
  clear: both;
}

.navbar-header:before,
.navbar-header:after {
  display: table;
  content: " ";
}

.navbar-header:after {
  clear: both;
}

.navbar-header:before,
.navbar-header:after {
  display: table;
  content: " ";
}

.navbar-header:after {
  clear: both;
}

.navbar-header:before,
.navbar-header:after {
  display: table;
  content: " ";
}

.navbar-header:after {
  clear: both;
}

.navbar-header:before,
.navbar-header:after {
  display: table;
  content: " ";
}

.navbar-header:after {
  clear: both;
}

@media (min-width: 768px) {
  .navbar-header {
    float: left;
  }
}

.navbar-collapse {
  max-height: 340px;
  padding-right: 15px;
  padding-left: 15px;
  overflow-x: visible;
  border-top: 1px solid transparent;
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.1);
  -webkit-overflow-scrolling: touch;
}

.navbar-collapse:before,
.navbar-collapse:after {
  display: table;
  content: " ";
}

.navbar-collapse:after {
  clear: both;
}

.navbar-collapse:before,
.navbar-collapse:after {
  display: table;
  content: " ";
}

.navbar-collapse:after {
  clear: both;
}

.navbar-collapse:before,
.navbar-collapse:after {
  display: table;
  content: " ";
}

.navbar-collapse:after {
  clear: both;
}

.navbar-collapse:before,
.navbar-collapse:after {
  display: table;
  content: " ";
}

.navbar-collapse:after {
  clear: both;
}

.navbar-collapse:before,
.navbar-collapse:after {
  display: table;
  content: " ";
}

.navbar-collapse:after {
  clear: both;
}

.navbar-collapse.in {
  overflow-y: auto;
}

@media (min-width: 768px) {
  .navbar-collapse {
    width: auto;
    border-top: 0;
    box-shadow: none;
  }
  .navbar-collapse.collapse {
    display: block !important;
    height: auto !important;
    padding-bottom: 0;
    overflow: visible !important;
  }
  .navbar-collapse.in {
    overflow-y: visible;
  }
  .navbar-fixed-top .navbar-collapse,
  .navbar-static-top .navbar-collapse,
  .navbar-fixed-bottom .navbar-collapse {
    padding-right: 0;
    padding-left: 0;
  }
}

.container > .navbar-header,
.container > .navbar-collapse {
  margin-right: -15px;
  margin-left: -15px;
}

@media (min-width: 768px) {
  .container > .navbar-header,
  .container > .navbar-collapse {
    margin-right: 0;
    margin-left: 0;
  }
}

.navbar-static-top {
  z-index: 1000;
  border-width: 0 0 1px;
}

@media (min-width: 768px) {
  .navbar-static-top {
    border-radius: 0;
  }
}

.navbar-fixed-top,
.navbar-fixed-bottom {
  position: fixed;
  right: 0;
  left: 0;
  z-index: 1030;
}

@media (min-width: 768px) {
  .navbar-fixed-top,
  .navbar-fixed-bottom {
    border-radius: 0;
  }
}

.navbar-fixed-top {
  top: 0;
  border-width: 0 0 1px;
}

.navbar-fixed-bottom {
  bottom: 0;
  margin-bottom: 0;
  border-width: 1px 0 0;
}

.navbar-brand {
  float: left;
  padding: 15px 15px;
  font-size: 18px;
  line-height: 20px;
}

.navbar-brand:hover,
.navbar-brand:focus {
  text-decoration: none;
}

@media (min-width: 768px) {
  .navbar > .container .navbar-brand {
    margin-left: -15px;
  }
}

.navbar-toggle {
  position: relative;
  float: right;
  padding: 9px 10px;
  margin-top: 8px;
  margin-right: 15px;
  margin-bottom: 8px;
  background-color: transparent;
  background-image: none;
  border: 1px solid transparent;
  border-radius: 4px;
}

.navbar-toggle .icon-bar {
  display: block;
  width: 22px;
  height: 2px;
  border-radius: 1px;
}

.navbar-toggle .icon-bar + .icon-bar {
  margin-top: 4px;
}

@media (min-width: 768px) {
  .navbar-toggle {
    display: none;
  }
}

.navbar-nav {
  margin: 7.5px -15px;
}

.navbar-nav > li > a {
  padding-top: 10px;
  padding-bottom: 10px;
  line-height: 20px;
}

@media (max-width: 767px) {
  .navbar-nav .open .dropdown-menu {
    position: static;
    float: none;
    width: auto;
    margin-top: 0;
    background-color: transparent;
    border: 0;
    box-shadow: none;
  }
  .navbar-nav .open .dropdown-menu > li > a,
  .navbar-nav .open .dropdown-menu .dropdown-header {
    padding: 5px 15px 5px 25px;
  }
  .navbar-nav .open .dropdown-menu > li > a {
    line-height: 20px;
  }
  .navbar-nav .open .dropdown-menu > li > a:hover,
  .navbar-nav .open .dropdown-menu > li > a:focus {
    background-image: none;
  }
}

@media (min-width: 768px) {
  .navbar-nav {
    float: left;
    margin: 0;
  }
  .navbar-nav > li {
    float: left;
  }
  .navbar-nav > li > a {
    padding-top: 15px;
    padding-bottom: 15px;
  }
  .navbar-nav.navbar-right:last-child {
    margin-right: -15px;
  }
}

@media (min-width: 768px) {
  .navbar-left {
    float: left !important;
  }
  .navbar-right {
    float: right !important;
  }
}

.navbar-form {
  padding: 10px 15px;
  margin-top: 6px;
  margin-right: -15px;
  margin-bottom: 6px;
  margin-left: -15px;
  border-top: 1px solid transparent;
  border-bottom: 1px solid transparent;
  -webkit-box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 0 rgba(255, 255, 255, 0.1);
          box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.1), 0 1px 0 rgba(255, 255, 255, 0.1);
}

@media (min-width: 768px) {
  .navbar-form .form-group {
    display: inline-block;
    margin-bottom: 0;
    vertical-align: middle;
  }
  .navbar-form .form-control {
    display: inline-block;
  }
  .navbar-form select.form-control {
    width: auto;
  }
  .navbar-form .radio,
  .navbar-form .checkbox {
    display: inline-block;
    padding-left: 0;
    margin-top: 0;
    margin-bottom: 0;
  }
  .navbar-form .radio input[type="radio"],
  .navbar-form .checkbox input[type="checkbox"] {
    float: none;
    margin-left: 0;
  }
}

@media (max-width: 767px) {
  .navbar-form .form-group {
    margin-bottom: 5px;
  }
}

@media (min-width: 768px) {
  .navbar-form {
    width: auto;
    padding-top: 0;
    padding-bottom: 0;
    margin-right: 0;
    margin-left: 0;
    border: 0;
    -webkit-box-shadow: none;
            box-shadow: none;
  }
  .navbar-form.navbar-right:last-child {
    margin-right: -15px;
  }
}

.navbar-nav > li > .dropdown-menu {
  margin-top: 0;
  border-top-right-radius: 0;
  border-top-left-radius: 0;
}

.navbar-fixed-bottom .navbar-nav > li > .dropdown-menu {
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}

.navbar-nav.pull-right > li > .dropdown-menu,
.navbar-nav > li > .dropdown-menu.pull-right {
  right: 0;
  left: auto;
}

.navbar-btn {
  margin-top: 6px;
  margin-bottom: 6px;
}

.navbar-btn.btn-sm {
  margin-top: 10px;
  margin-bottom: 10px;
}

.navbar-btn.btn-xs {
  margin-top: 14px;
  margin-bottom: 14px;
}

.navbar-text {
  margin-top: 15px;
  margin-bottom: 15px;
}

@media (min-width: 768px) {
  .navbar-text {
    float: left;
    margin-right: 15px;
    margin-left: 15px;
  }
  .navbar-text.navbar-right:last-child {
    margin-right: 0;
  }
}

.navbar-default {
  background-color: #3a3f44;
  border-color: #2b2e32;
}

.navbar-default .navbar-brand {
  color: #c8c8c8;
}

.navbar-default .navbar-brand:hover,
.navbar-default .navbar-brand:focus {
  color: #ffffff;
  background-color: none;
}

.navbar-default .navbar-text {
  color: #c8c8c8;
}

.navbar-default .navbar-nav > li > a {
  color: #c8c8c8;
}

.navbar-default .navbar-nav > li > a:hover,
.navbar-default .navbar-nav > li > a:focus {
  color: #ffffff;
  background-color: #272b2e;
}

.navbar-default .navbar-nav > .active > a,
.navbar-default .navbar-nav > .active > a:hover,
.navbar-default .navbar-nav > .active > a:focus {
  color: #ffffff;
  background-color: #272b2e;
}

.navbar-default .navbar-nav > .disabled > a,
.navbar-default .navbar-nav > .disabled > a:hover,
.navbar-default .navbar-nav > .disabled > a:focus {
  color: #cccccc;
  background-color: transparent;
}

.navbar-default .navbar-toggle {
  border-color: #272b2e;
}

.navbar-default .navbar-toggle:hover,
.navbar-default .navbar-toggle:focus {
  background-color: #272b2e;
}

.navbar-default .navbar-toggle .icon-bar {
  background-color: #c8c8c8;
}

.navbar-default .navbar-collapse,
.navbar-default .navbar-form {
  border-color: #2b2e32;
}

.navbar-default .navbar-nav > .open > a,
.navbar-default .navbar-nav > .open > a:hover,
.navbar-default .navbar-nav > .open > a:focus {
  color: #ffffff;
  background-color: #272b2e;
}

@media (max-width: 767px) {
  .navbar-default .navbar-nav .open .dropdown-menu > li > a {
    color: #c8c8c8;
  }
  .navbar-default .navbar-nav .open .dropdown-menu > li > a:hover,
  .navbar-default .navbar-nav .open .dropdown-menu > li > a:focus {
    color: #ffffff;
    background-color: #272b2e;
  }
  .navbar-default .navbar-nav .open .dropdown-menu > .active > a,
  .navbar-default .navbar-nav .open .dropdown-menu > .active > a:hover,
  .navbar-default .navbar-nav .open .dropdown-menu > .active > a:focus {
    color: #ffffff;
    background-color: #272b2e;
  }
  .navbar-default .navbar-nav .open .dropdown-menu > .disabled > a,
  .navbar-default .navbar-nav .open .dropdown-menu > .disabled > a:hover,
  .navbar-default .navbar-nav .open .dropdown-menu > .disabled > a:focus {
    color: #cccccc;
    background-color: transparent;
  }
}

.navbar-default .navbar-link {
  color: #c8c8c8;
}

.navbar-default .navbar-link:hover {
  color: #ffffff;
}

.navbar-inverse {
  background-color: #7a8288;
  border-color: #62686d;
}

.navbar-inverse .navbar-brand {
  color: #cccccc;
}

.navbar-inverse .navbar-brand:hover,
.navbar-inverse .navbar-brand:focus {
  color: #ffffff;
  background-color: none;
}

.navbar-inverse .navbar-text {
  color: #cccccc;
}

.navbar-inverse .navbar-nav > li > a {
  color: #cccccc;
}

.navbar-inverse .navbar-nav > li > a:hover,
.navbar-inverse .navbar-nav > li > a:focus {
  color: #ffffff;
  background-color: #5d6368;
}

.navbar-inverse .navbar-nav > .active > a,
.navbar-inverse .navbar-nav > .active > a:hover,
.navbar-inverse .navbar-nav > .active > a:focus {
  color: #ffffff;
  background-color: #5d6368;
}

.navbar-inverse .navbar-nav > .disabled > a,
.navbar-inverse .navbar-nav > .disabled > a:hover,
.navbar-inverse .navbar-nav > .disabled > a:focus {
  color: #cccccc;
  background-color: transparent;
}

.navbar-inverse .navbar-toggle {
  border-color: #5d6368;
}

.navbar-inverse .navbar-toggle:hover,
.navbar-inverse .navbar-toggle:focus {
  background-color: #5d6368;
}

.navbar-inverse .navbar-toggle .icon-bar {
  background-color: #ffffff;
}

.navbar-inverse .navbar-collapse,
.navbar-inverse .navbar-form {
  border-color: #697075;
}

.navbar-inverse .navbar-nav > .open > a,
.navbar-inverse .navbar-nav > .open > a:hover,
.navbar-inverse .navbar-nav > .open > a:focus {
  color: #ffffff;
  background-color: #5d6368;
}

@media (max-width: 767px) {
  .navbar-inverse .navbar-nav .open .dropdown-menu > .dropdown-header {
    border-color: #62686d;
  }
  .navbar-inverse .navbar-nav .open .dropdown-menu .divider {
    background-color: #62686d;
  }
  .navbar-inverse .navbar-nav .open .dropdown-menu > li > a {
    color: #cccccc;
  }
  .navbar-inverse .navbar-nav .open .dropdown-menu > li > a:hover,
  .navbar-inverse .navbar-nav .open .dropdown-menu > li > a:focus {
    color: #ffffff;
    background-color: #5d6368;
  }
  .navbar-inverse .navbar-nav .open .dropdown-menu > .active > a,
  .navbar-inverse .navbar-nav .open .dropdown-menu > .active > a:hover,
  .navbar-inverse .navbar-nav .open .dropdown-menu > .active > a:focus {
    color: #ffffff;
    background-color: #5d6368;
  }
  .navbar-inverse .navbar-nav .open .dropdown-menu > .disabled > a,
  .navbar-inverse .navbar-nav .open .dropdown-menu > .disabled > a:hover,
  .navbar-inverse .navbar-nav .open .dropdown-menu > .disabled > a:focus {
    color: #cccccc;
    background-color: transparent;
  }
}

.navbar-inverse .navbar-link {
  color: #cccccc;
}

.navbar-inverse .navbar-link:hover {
  color: #ffffff;
}

.breadcrumb {
  padding: 8px 15px;
  margin-bottom: 20px;
  list-style: none;
  background-color: transparent;
  border-radius: 4px;
}

.breadcrumb > li {
  display: inline-block;
}

.breadcrumb > li + li:before {
  padding: 0 5px;
  color: #cccccc;
  content: "/\00a0";
}

.breadcrumb > .active {
  color: #7a8288;
}

.pagination {
  display: inline-block;
  padding-left: 0;
  margin: 20px 0;
  border-radius: 4px;
}

.pagination > li {
  display: inline;
}

.pagination > li > a,
.pagination > li > span {
  position: relative;
  float: left;
  padding: 8px 12px;
  margin-left: -1px;
  line-height: 1.428571429;
  text-decoration: none;
  background-color: #3a3f44;
  border: 1px solid rgba(0, 0, 0, 0.6);
}

.pagination > li:first-child > a,
.pagination > li:first-child > span {
  margin-left: 0;
  border-bottom-left-radius: 4px;
  border-top-left-radius: 4px;
}

.pagination > li:last-child > a,
.pagination > li:last-child > span {
  border-top-right-radius: 4px;
  border-bottom-right-radius: 4px;
}

.pagination > li > a:hover,
.pagination > li > span:hover,
.pagination > li > a:focus,
.pagination > li > span:focus {
  background-color: transparent;
}

.pagination > .active > a,
.pagination > .active > span,
.pagination > .active > a:hover,
.pagination > .active > span:hover,
.pagination > .active > a:focus,
.pagination > .active > span:focus {
  z-index: 2;
  color: #ffffff;
  cursor: default;
  background-color: #232628;
  border-color: #232628;
}

.pagination > .disabled > span,
.pagination > .disabled > span:hover,
.pagination > .disabled > span:focus,
.pagination > .disabled > a,
.pagination > .disabled > a:hover,
.pagination > .disabled > a:focus {
  color: #7a8288;
  cursor: not-allowed;
  background-color: #3a3f44;
  border-color: rgba(0, 0, 0, 0.6);
}

.pagination-lg > li > a,
.pagination-lg > li > span {
  padding: 14px 16px;
  font-size: 18px;
}

.pagination-lg > li:first-child > a,
.pagination-lg > li:first-child > span {
  border-bottom-left-radius: 6px;
  border-top-left-radius: 6px;
}

.pagination-lg > li:last-child > a,
.pagination-lg > li:last-child > span {
  border-top-right-radius: 6px;
  border-bottom-right-radius: 6px;
}

.pagination-sm > li > a,
.pagination-sm > li > span {
  padding: 5px 10px;
  font-size: 12px;
}

.pagination-sm > li:first-child > a,
.pagination-sm > li:first-child > span {
  border-bottom-left-radius: 3px;
  border-top-left-radius: 3px;
}

.pagination-sm > li:last-child > a,
.pagination-sm > li:last-child > span {
  border-top-right-radius: 3px;
  border-bottom-right-radius: 3px;
}

.pager {
  padding-left: 0;
  margin: 20px 0;
  text-align: center;
  list-style: none;
}

.pager:before,
.pager:after {
  display: table;
  content: " ";
}

.pager:after {
  clear: both;
}

.pager:before,
.pager:after {
  display: table;
  content: " ";
}

.pager:after {
  clear: both;
}

.pager:before,
.pager:after {
  display: table;
  content: " ";
}

.pager:after {
  clear: both;
}

.pager:before,
.pager:after {
  display: table;
  content: " ";
}

.pager:after {
  clear: both;
}

.pager:before,
.pager:after {
  display: table;
  content: " ";
}

.pager:after {
  clear: both;
}

.pager li {
  display: inline;
}

.pager li > a,
.pager li > span {
  display: inline-block;
  padding: 5px 14px;
  background-color: #3a3f44;
  border: 1px solid rgba(0, 0, 0, 0.6);
  border-radius: 15px;
}

.pager li > a:hover,
.pager li > a:focus {
  text-decoration: none;
  background-color: transparent;
}

.pager .next > a,
.pager .next > span {
  float: right;
}

.pager .previous > a,
.pager .previous > span {
  float: left;
}

.pager .disabled > a,
.pager .disabled > a:hover,
.pager .disabled > a:focus,
.pager .disabled > span {
  color: #7a8288;
  cursor: not-allowed;
  background-color: #3a3f44;
}

.label {
  display: inline;
  padding: .2em .6em .3em;
  font-size: 75%;
  font-weight: bold;
  line-height: 1;
  color: #ffffff;
  text-align: center;
  white-space: nowrap;
  vertical-align: baseline;
  border-radius: .25em;
}

.label[href]:hover,
.label[href]:focus {
  color: #ffffff;
  text-decoration: none;
  cursor: pointer;
}

.label:empty {
  display: none;
}

.btn .label {
  position: relative;
  top: -1px;
}

.label-default {
  background-color: #3a3f44;
}

.label-default[href]:hover,
.label-default[href]:focus {
  background-color: #232628;
}

.label-primary {
  background-color: #7a8288;
}

.label-primary[href]:hover,
.label-primary[href]:focus {
  background-color: #62686d;
}

.label-success {
  background-color: #62c462;
}

.label-success[href]:hover,
.label-success[href]:focus {
  background-color: #42b142;
}

.label-info {
  background-color: #5bc0de;
}

.label-info[href]:hover,
.label-info[href]:focus {
  background-color: #31b0d5;
}

.label-warning {
  background-color: #f89406;
}

.label-warning[href]:hover,
.label-warning[href]:focus {
  background-color: #c67605;
}

.label-danger {
  background-color: #ee5f5b;
}

.label-danger[href]:hover,
.label-danger[href]:focus {
  background-color: #e9322d;
}

.badge {
  display: inline-block;
  min-width: 10px;
  padding: 3px 7px;
  font-size: 12px;
  font-weight: bold;
  line-height: 1;
  color: #ffffff;
  text-align: center;
  white-space: nowrap;
  vertical-align: baseline;
  background-color: #7a8288;
  border-radius: 10px;
}

.badge:empty {
  display: none;
}

.btn .badge {
  position: relative;
  top: -1px;
}

a.badge:hover,
a.badge:focus {
  color: #ffffff;
  text-decoration: none;
  cursor: pointer;
}

a.list-group-item.active > .badge,
.nav-pills > .active > a > .badge {
  color: #ffffff;
  background-color: #7a8288;
}

.nav-pills > li > a > .badge {
  margin-left: 3px;
}

.jumbotron {
  padding: 30px;
  margin-bottom: 30px;
  font-size: 21px;
  font-weight: 200;
  line-height: 2.1428571435;
  color: inherit;
  background-color: #1c1e22;
}

.jumbotron h1,
.jumbotron .h1 {
  line-height: 1;
  color: inherit;
}

.jumbotron p {
  line-height: 1.4;
}

.container .jumbotron {
  border-radius: 6px;
}

.jumbotron .container {
  max-width: 100%;
}

@media screen and (min-width: 768px) {
  .jumbotron {
    padding-top: 48px;
    padding-bottom: 48px;
  }
  .container .jumbotron {
    padding-right: 60px;
    padding-left: 60px;
  }
  .jumbotron h1,
  .jumbotron .h1 {
    font-size: 63px;
  }
}

.thumbnail {
  display: block;
  padding: 4px;
  margin-bottom: 20px;
  line-height: 1.428571429;
  background-color: #272b30;
  border: 1px solid #dddddd;
  border-radius: 4px;
  -webkit-transition: all 0.2s ease-in-out;
          transition: all 0.2s ease-in-out;
}

.thumbnail > img,
.thumbnail a > img {
  display: block;
  height: auto;
  max-width: 100%;
  margin-right: auto;
  margin-left: auto;
}

a.thumbnail:hover,
a.thumbnail:focus,
a.thumbnail.active {
  border-color: #ffffff;
}

.thumbnail .caption {
  padding: 9px;
  color: #c8c8c8;
}

.alert {
  padding: 15px;
  margin-bottom: 20px;
  border: 1px solid transparent;
  border-radius: 4px;
}

.alert h4 {
  margin-top: 0;
  color: inherit;
}

.alert .alert-link {
  font-weight: bold;
}

.alert > p,
.alert > ul {
  margin-bottom: 0;
}

.alert > p + p {
  margin-top: 5px;
}

.alert-dismissable {
  padding-right: 35px;
}

.alert-dismissable .close {
  position: relative;
  top: -2px;
  right: -21px;
  color: inherit;
}

.alert-success {
  color: #ffffff;
  background-color: #62c462;
  border-color: #62bd4f;
}

.alert-success hr {
  border-top-color: #55b142;
}

.alert-success .alert-link {
  color: #e6e6e6;
}

.alert-info {
  color: #ffffff;
  background-color: #5bc0de;
  border-color: #3dced8;
}

.alert-info hr {
  border-top-color: #2ac7d2;
}

.alert-info .alert-link {
  color: #e6e6e6;
}

.alert-warning {
  color: #ffffff;
  background-color: #f89406;
  border-color: #e96506;
}

.alert-warning hr {
  border-top-color: #d05a05;
}

.alert-warning .alert-link {
  color: #e6e6e6;
}

.alert-danger {
  color: #ffffff;
  background-color: #ee5f5b;
  border-color: #ed4d63;
}

.alert-danger hr {
  border-top-color: #ea364f;
}

.alert-danger .alert-link {
  color: #e6e6e6;
}

@-webkit-keyframes progress-bar-stripes {
  from {
    background-position: 40px 0;
  }
  to {
    background-position: 0 0;
  }
}

@keyframes progress-bar-stripes {
  from {
    background-position: 40px 0;
  }
  to {
    background-position: 0 0;
  }
}

.progress {
  height: 20px;
  margin-bottom: 20px;
  overflow: hidden;
  background-color: #1c1e22;
  border-radius: 4px;
  -webkit-box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.1);
          box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.1);
}

.progress-bar {
  float: left;
  width: 0;
  height: 100%;
  font-size: 12px;
  line-height: 20px;
  color: #ffffff;
  text-align: center;
  background-color: #7a8288;
  -webkit-box-shadow: inset 0 -1px 0 rgba(0, 0, 0, 0.15);
          box-shadow: inset 0 -1px 0 rgba(0, 0, 0, 0.15);
  -webkit-transition: width 0.6s ease;
          transition: width 0.6s ease;
}

.progress-striped .progress-bar {
  background-image: -webkit-linear-gradient(45deg, rgba(255, 255, 255, 0.15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, 0.15) 50%, rgba(255, 255, 255, 0.15) 75%, transparent 75%, transparent);
  background-image: linear-gradient(45deg, rgba(255, 255, 255, 0.15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, 0.15) 50%, rgba(255, 255, 255, 0.15) 75%, transparent 75%, transparent);
  background-size: 40px 40px;
}

.progress.active .progress-bar {
  -webkit-animation: progress-bar-stripes 2s linear infinite;
          animation: progress-bar-stripes 2s linear infinite;
}

.progress-bar-success {
  background-color: #62c462;
}

.progress-striped .progress-bar-success {
  background-image: -webkit-linear-gradient(45deg, rgba(255, 255, 255, 0.15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, 0.15) 50%, rgba(255, 255, 255, 0.15) 75%, transparent 75%, transparent);
  background-image: linear-gradient(45deg, rgba(255, 255, 255, 0.15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, 0.15) 50%, rgba(255, 255, 255, 0.15) 75%, transparent 75%, transparent);
}

.progress-bar-info {
  background-color: #5bc0de;
}

.progress-striped .progress-bar-info {
  background-image: -webkit-linear-gradient(45deg, rgba(255, 255, 255, 0.15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, 0.15) 50%, rgba(255, 255, 255, 0.15) 75%, transparent 75%, transparent);
  background-image: linear-gradient(45deg, rgba(255, 255, 255, 0.15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, 0.15) 50%, rgba(255, 255, 255, 0.15) 75%, transparent 75%, transparent);
}

.progress-bar-warning {
  background-color: #f89406;
}

.progress-striped .progress-bar-warning {
  background-image: -webkit-linear-gradient(45deg, rgba(255, 255, 255, 0.15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, 0.15) 50%, rgba(255, 255, 255, 0.15) 75%, transparent 75%, transparent);
  background-image: linear-gradient(45deg, rgba(255, 255, 255, 0.15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, 0.15) 50%, rgba(255, 255, 255, 0.15) 75%, transparent 75%, transparent);
}

.progress-bar-danger {
  background-color: #ee5f5b;
}

.progress-striped .progress-bar-danger {
  background-image: -webkit-linear-gradient(45deg, rgba(255, 255, 255, 0.15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, 0.15) 50%, rgba(255, 255, 255, 0.15) 75%, transparent 75%, transparent);
  background-image: linear-gradient(45deg, rgba(255, 255, 255, 0.15) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, 0.15) 50%, rgba(255, 255, 255, 0.15) 75%, transparent 75%, transparent);
}

.media,
.media-body {
  overflow: hidden;
  zoom: 1;
}

.media,
.media .media {
  margin-top: 15px;
}

.media:first-child {
  margin-top: 0;
}

.media-object {
  display: block;
}

.media-heading {
  margin: 0 0 5px;
}

.media > .pull-left {
  margin-right: 10px;
}

.media > .pull-right {
  margin-left: 10px;
}

.media-list {
  padding-left: 0;
  list-style: none;
}

.list-group {
  padding-left: 0;
  margin-bottom: 20px;
}

.list-group-item {
  position: relative;
  display: block;
  padding: 10px 15px;
  margin-bottom: -1px;
  background-color: transparent;
  border: 1px solid rgba(0, 0, 0, 0.6);
}

.list-group-item:first-child {
  border-top-right-radius: 4px;
  border-top-left-radius: 4px;
}

.list-group-item:last-child {
  margin-bottom: 0;
  border-bottom-right-radius: 4px;
  border-bottom-left-radius: 4px;
}

.list-group-item > .badge {
  float: right;
}

.list-group-item > .badge + .badge {
  margin-right: 5px;
}

a.list-group-item {
  color: #c8c8c8;
}

a.list-group-item .list-group-item-heading {
  color: #ffffff;
}

a.list-group-item:hover,
a.list-group-item:focus {
  text-decoration: none;
  background-color: #3e444c;
}

a.list-group-item.active,
a.list-group-item.active:hover,
a.list-group-item.active:focus {
  z-index: 2;
  color: #ffffff;
  background-color: #3e444c;
  border-color: rgba(0, 0, 0, 0.6);
}

a.list-group-item.active .list-group-item-heading,
a.list-group-item.active:hover .list-group-item-heading,
a.list-group-item.active:focus .list-group-item-heading {
  color: inherit;
}

a.list-group-item.active .list-group-item-text,
a.list-group-item.active:hover .list-group-item-text,
a.list-group-item.active:focus .list-group-item-text {
  color: #a2aab4;
}

.list-group-item-heading {
  margin-top: 0;
  margin-bottom: 5px;
}

.list-group-item-text {
  margin-bottom: 0;
  line-height: 1.3;
}

.panel {
  margin-bottom: 20px;
  background-color: #2e3338;
  border: 1px solid transparent;
  border-radius: 4px;
  -webkit-box-shadow: 0 1px 1px rgba(0, 0, 0, 0.05);
          box-shadow: 0 1px 1px rgba(0, 0, 0, 0.05);
}

.panel-body {
  padding: 15px;
}

.panel-body:before,
.panel-body:after {
  display: table;
  content: " ";
}

.panel-body:after {
  clear: both;
}

.panel-body:before,
.panel-body:after {
  display: table;
  content: " ";
}

.panel-body:after {
  clear: both;
}

.panel-body:before,
.panel-body:after {
  display: table;
  content: " ";
}

.panel-body:after {
  clear: both;
}

.panel-body:before,
.panel-body:after {
  display: table;
  content: " ";
}

.panel-body:after {
  clear: both;
}

.panel-body:before,
.panel-body:after {
  display: table;
  content: " ";
}

.panel-body:after {
  clear: both;
}

.panel > .list-group {
  margin-bottom: 0;
}

.panel > .list-group .list-group-item {
  border-width: 1px 0;
}

.panel > .list-group .list-group-item:first-child {
  border-top-right-radius: 0;
  border-top-left-radius: 0;
}

.panel > .list-group .list-group-item:last-child {
  border-bottom: 0;
}

.panel-heading + .list-group .list-group-item:first-child {
  border-top-width: 0;
}

.panel > .table,
.panel > .table-responsive > .table {
  margin-bottom: 0;
}

.panel > .panel-body + .table,
.panel > .panel-body + .table-responsive {
  border-top: 1px solid #1c1e22;
}

.panel > .table > tbody:first-child th,
.panel > .table > tbody:first-child td {
  border-top: 0;
}

.panel > .table-bordered,
.panel > .table-responsive > .table-bordered {
  border: 0;
}

.panel > .table-bordered > thead > tr > th:first-child,
.panel > .table-responsive > .table-bordered > thead > tr > th:first-child,
.panel > .table-bordered > tbody > tr > th:first-child,
.panel > .table-responsive > .table-bordered > tbody > tr > th:first-child,
.panel > .table-bordered > tfoot > tr > th:first-child,
.panel > .table-responsive > .table-bordered > tfoot > tr > th:first-child,
.panel > .table-bordered > thead > tr > td:first-child,
.panel > .table-responsive > .table-bordered > thead > tr > td:first-child,
.panel > .table-bordered > tbody > tr > td:first-child,
.panel > .table-responsive > .table-bordered > tbody > tr > td:first-child,
.panel > .table-bordered > tfoot > tr > td:first-child,
.panel > .table-responsive > .table-bordered > tfoot > tr > td:first-child {
  border-left: 0;
}

.panel > .table-bordered > thead > tr > th:last-child,
.panel > .table-responsive > .table-bordered > thead > tr > th:last-child,
.panel > .table-bordered > tbody > tr > th:last-child,
.panel > .table-responsive > .table-bordered > tbody > tr > th:last-child,
.panel > .table-bordered > tfoot > tr > th:last-child,
.panel > .table-responsive > .table-bordered > tfoot > tr > th:last-child,
.panel > .table-bordered > thead > tr > td:last-child,
.panel > .table-responsive > .table-bordered > thead > tr > td:last-child,
.panel > .table-bordered > tbody > tr > td:last-child,
.panel > .table-responsive > .table-bordered > tbody > tr > td:last-child,
.panel > .table-bordered > tfoot > tr > td:last-child,
.panel > .table-responsive > .table-bordered > tfoot > tr > td:last-child {
  border-right: 0;
}

.panel > .table-bordered > thead > tr:last-child > th,
.panel > .table-responsive > .table-bordered > thead > tr:last-child > th,
.panel > .table-bordered > tbody > tr:last-child > th,
.panel > .table-responsive > .table-bordered > tbody > tr:last-child > th,
.panel > .table-bordered > tfoot > tr:last-child > th,
.panel > .table-responsive > .table-bordered > tfoot > tr:last-child > th,
.panel > .table-bordered > thead > tr:last-child > td,
.panel > .table-responsive > .table-bordered > thead > tr:last-child > td,
.panel > .table-bordered > tbody > tr:last-child > td,
.panel > .table-responsive > .table-bordered > tbody > tr:last-child > td,
.panel > .table-bordered > tfoot > tr:last-child > td,
.panel > .table-responsive > .table-bordered > tfoot > tr:last-child > td {
  border-bottom: 0;
}

.panel > .table-responsive {
  margin-bottom: 0;
  border: 0;
}

.panel-heading {
  padding: 10px 15px;
  border-bottom: 1px solid transparent;
  border-top-right-radius: 3px;
  border-top-left-radius: 3px;
}

.panel-heading > .dropdown .dropdown-toggle {
  color: inherit;
}

.panel-title {
  margin-top: 0;
  margin-bottom: 0;
  font-size: 16px;
  color: inherit;
}

.panel-title > a {
  color: inherit;
}

.panel-footer {
  padding: 10px 15px;
  background-color: #3e444c;
  border-top: 1px solid rgba(0, 0, 0, 0.6);
  border-bottom-right-radius: 3px;
  border-bottom-left-radius: 3px;
}

.panel-group .panel {
  margin-bottom: 0;
  overflow: hidden;
  border-radius: 4px;
}

.panel-group .panel + .panel {
  margin-top: 5px;
}

.panel-group .panel-heading {
  border-bottom: 0;
}

.panel-group .panel-heading + .panel-collapse .panel-body {
  border-top: 1px solid rgba(0, 0, 0, 0.6);
}

.panel-group .panel-footer {
  border-top: 0;
}

.panel-group .panel-footer + .panel-collapse .panel-body {
  border-bottom: 1px solid rgba(0, 0, 0, 0.6);
}

.panel-default {
  border-color: rgba(0, 0, 0, 0.6);
}

.panel-default > .panel-heading {
  color: #c8c8c8;
  background-color: #3e444c;
  border-color: rgba(0, 0, 0, 0.6);
}

.panel-default > .panel-heading + .panel-collapse .panel-body {
  border-top-color: rgba(0, 0, 0, 0.6);
}

.panel-default > .panel-footer + .panel-collapse .panel-body {
  border-bottom-color: rgba(0, 0, 0, 0.6);
}

.panel-primary {
  border-color: rgba(0, 0, 0, 0.6);
}

.panel-primary > .panel-heading {
  color: #ffffff;
  background-color: #7a8288;
  border-color: rgba(0, 0, 0, 0.6);
}

.panel-primary > .panel-heading + .panel-collapse .panel-body {
  border-top-color: rgba(0, 0, 0, 0.6);
}

.panel-primary > .panel-footer + .panel-collapse .panel-body {
  border-bottom-color: rgba(0, 0, 0, 0.6);
}

.panel-success {
  border-color: rgba(0, 0, 0, 0.6);
}

.panel-success > .panel-heading {
  color: #ffffff;
  background-color: #62c462;
  border-color: rgba(0, 0, 0, 0.6);
}

.panel-success > .panel-heading + .panel-collapse .panel-body {
  border-top-color: rgba(0, 0, 0, 0.6);
}

.panel-success > .panel-footer + .panel-collapse .panel-body {
  border-bottom-color: rgba(0, 0, 0, 0.6);
}

.panel-warning {
  border-color: rgba(0, 0, 0, 0.6);
}

.panel-warning > .panel-heading {
  color: #ffffff;
  background-color: #f89406;
  border-color: rgba(0, 0, 0, 0.6);
}

.panel-warning > .panel-heading + .panel-collapse .panel-body {
  border-top-color: rgba(0, 0, 0, 0.6);
}

.panel-warning > .panel-footer + .panel-collapse .panel-body {
  border-bottom-color: rgba(0, 0, 0, 0.6);
}

.panel-danger {
  border-color: rgba(0, 0, 0, 0.6);
}

.panel-danger > .panel-heading {
  color: #ffffff;
  background-color: #ee5f5b;
  border-color: rgba(0, 0, 0, 0.6);
}

.panel-danger > .panel-heading + .panel-collapse .panel-body {
  border-top-color: rgba(0, 0, 0, 0.6);
}

.panel-danger > .panel-footer + .panel-collapse .panel-body {
  border-bottom-color: rgba(0, 0, 0, 0.6);
}

.panel-info {
  border-color: rgba(0, 0, 0, 0.6);
}

.panel-info > .panel-heading {
  color: #ffffff;
  background-color: #5bc0de;
  border-color: rgba(0, 0, 0, 0.6);
}

.panel-info > .panel-heading + .panel-collapse .panel-body {
  border-top-color: rgba(0, 0, 0, 0.6);
}

.panel-info > .panel-footer + .panel-collapse .panel-body {
  border-bottom-color: rgba(0, 0, 0, 0.6);
}

.well {
  min-height: 20px;
  padding: 19px;
  margin-bottom: 20px;
  background-color: #1c1e22;
  border: 1px solid #0c0d0e;
  border-radius: 4px;
  -webkit-box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.05);
          box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.05);
}

.well blockquote {
  border-color: #ddd;
  border-color: rgba(0, 0, 0, 0.15);
}

.well-lg {
  padding: 24px;
  border-radius: 6px;
}

.well-sm {
  padding: 9px;
  border-radius: 3px;
}

.close {
  float: right;
  font-size: 21px;
  font-weight: bold;
  line-height: 1;
  color: #000000;
  text-shadow: 0 1px 0 #ffffff;
  opacity: 0.2;
  filter: alpha(opacity=20);
}

.close:hover,
.close:focus {
  color: #000000;
  text-decoration: none;
  cursor: pointer;
  opacity: 0.5;
  filter: alpha(opacity=50);
}

button.close {
  padding: 0;
  cursor: pointer;
  background: transparent;
  border: 0;
  -webkit-appearance: none;
}

.modal-open {
  overflow: hidden;
}

.modal {
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  z-index: 1040;
  display: none;
  overflow: auto;
  overflow-y: scroll;
}

.modal.fade .modal-dialog {
  -webkit-transform: translate(0, -25%);
      -ms-transform: translate(0, -25%);
          transform: translate(0, -25%);
  -webkit-transition: -webkit-transform 0.3s ease-out;
     -moz-transition: -moz-transform 0.3s ease-out;
       -o-transition: -o-transform 0.3s ease-out;
          transition: transform 0.3s ease-out;
}

.modal.in .modal-dialog {
  -webkit-transform: translate(0, 0);
      -ms-transform: translate(0, 0);
          transform: translate(0, 0);
}

.modal-dialog {
  position: relative;
  z-index: 1050;
  width: auto;
  margin: 10px;
}

.modal-content {
  position: relative;
  background-color: #2e3338;
  border: 1px solid #999999;
  border: 1px solid rgba(0, 0, 0, 0.2);
  border-radius: 6px;
  outline: none;
  -webkit-box-shadow: 0 3px 9px rgba(0, 0, 0, 0.5);
          box-shadow: 0 3px 9px rgba(0, 0, 0, 0.5);
  background-clip: padding-box;
}

.modal-backdrop {
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  z-index: 1030;
  background-color: #000000;
}

.modal-backdrop.fade {
  opacity: 0;
  filter: alpha(opacity=0);
}

.modal-backdrop.in {
  opacity: 0.5;
  filter: alpha(opacity=50);
}

.modal-header {
  min-height: 16.428571429px;
  padding: 15px;
  border-bottom: 1px solid #1c1e22;
}

.modal-header .close {
  margin-top: -2px;
}

.modal-title {
  margin: 0;
  line-height: 1.428571429;
}

.modal-body {
  position: relative;
  padding: 20px;
}

.modal-footer {
  padding: 19px 20px 20px;
  margin-top: 15px;
  text-align: right;
  border-top: 1px solid #1c1e22;
}

.modal-footer:before,
.modal-footer:after {
  display: table;
  content: " ";
}

.modal-footer:after {
  clear: both;
}

.modal-footer:before,
.modal-footer:after {
  display: table;
  content: " ";
}

.modal-footer:after {
  clear: both;
}

.modal-footer:before,
.modal-footer:after {
  display: table;
  content: " ";
}

.modal-footer:after {
  clear: both;
}

.modal-footer:before,
.modal-footer:after {
  display: table;
  content: " ";
}

.modal-footer:after {
  clear: both;
}

.modal-footer:before,
.modal-footer:after {
  display: table;
  content: " ";
}

.modal-footer:after {
  clear: both;
}

.modal-footer .btn + .btn {
  margin-bottom: 0;
  margin-left: 5px;
}

.modal-footer .btn-group .btn + .btn {
  margin-left: -1px;
}

.modal-footer .btn-block + .btn-block {
  margin-left: 0;
}

@media screen and (min-width: 768px) {
  .modal-dialog {
    width: 600px;
    margin: 30px auto;
  }
  .modal-content {
    -webkit-box-shadow: 0 5px 15px rgba(0, 0, 0, 0.5);
            box-shadow: 0 5px 15px rgba(0, 0, 0, 0.5);
  }
}

.tooltip {
  position: absolute;
  z-index: 1030;
  display: block;
  font-size: 12px;
  line-height: 1.4;
  opacity: 0;
  filter: alpha(opacity=0);
  visibility: visible;
}

.tooltip.in {
  opacity: 0.9;
  filter: alpha(opacity=90);
}

.tooltip.top {
  padding: 5px 0;
  margin-top: -3px;
}

.tooltip.right {
  padding: 0 5px;
  margin-left: 3px;
}

.tooltip.bottom {
  padding: 5px 0;
  margin-top: 3px;
}

.tooltip.left {
  padding: 0 5px;
  margin-left: -3px;
}

.tooltip-inner {
  max-width: 200px;
  padding: 3px 8px;
  color: #ffffff;
  text-align: center;
  text-decoration: none;
  background-color: rgba(0, 0, 0, 0.9);
  border-radius: 4px;
}

.tooltip-arrow {
  position: absolute;
  width: 0;
  height: 0;
  border-color: transparent;
  border-style: solid;
}

.tooltip.top .tooltip-arrow {
  bottom: 0;
  left: 50%;
  margin-left: -5px;
  border-top-color: rgba(0, 0, 0, 0.9);
  border-width: 5px 5px 0;
}

.tooltip.top-left .tooltip-arrow {
  bottom: 0;
  left: 5px;
  border-top-color: rgba(0, 0, 0, 0.9);
  border-width: 5px 5px 0;
}

.tooltip.top-right .tooltip-arrow {
  right: 5px;
  bottom: 0;
  border-top-color: rgba(0, 0, 0, 0.9);
  border-width: 5px 5px 0;
}

.tooltip.right .tooltip-arrow {
  top: 50%;
  left: 0;
  margin-top: -5px;
  border-right-color: rgba(0, 0, 0, 0.9);
  border-width: 5px 5px 5px 0;
}

.tooltip.left .tooltip-arrow {
  top: 50%;
  right: 0;
  margin-top: -5px;
  border-left-color: rgba(0, 0, 0, 0.9);
  border-width: 5px 0 5px 5px;
}

.tooltip.bottom .tooltip-arrow {
  top: 0;
  left: 50%;
  margin-left: -5px;
  border-bottom-color: rgba(0, 0, 0, 0.9);
  border-width: 0 5px 5px;
}

.tooltip.bottom-left .tooltip-arrow {
  top: 0;
  left: 5px;
  border-bottom-color: rgba(0, 0, 0, 0.9);
  border-width: 0 5px 5px;
}

.tooltip.bottom-right .tooltip-arrow {
  top: 0;
  right: 5px;
  border-bottom-color: rgba(0, 0, 0, 0.9);
  border-width: 0 5px 5px;
}

.popover {
  position: absolute;
  top: 0;
  left: 0;
  z-index: 1010;
  display: none;
  max-width: 276px;
  padding: 1px;
  text-align: left;
  white-space: normal;
  background-color: #2e3338;
  border: 1px solid #999999;
  border: 1px solid rgba(0, 0, 0, 0.2);
  border-radius: 6px;
  -webkit-box-shadow: 0 5px 10px rgba(0, 0, 0, 0.2);
          box-shadow: 0 5px 10px rgba(0, 0, 0, 0.2);
  background-clip: padding-box;
}

.popover.top {
  margin-top: -10px;
}

.popover.right {
  margin-left: 10px;
}

.popover.bottom {
  margin-top: 10px;
}

.popover.left {
  margin-left: -10px;
}

.popover-title {
  padding: 8px 14px;
  margin: 0;
  font-size: 14px;
  font-weight: normal;
  line-height: 18px;
  background-color: #2e3338;
  border-bottom: 1px solid #22262a;
  border-radius: 5px 5px 0 0;
}

.popover-content {
  padding: 9px 14px;
}

.popover .arrow,
.popover .arrow:after {
  position: absolute;
  display: block;
  width: 0;
  height: 0;
  border-color: transparent;
  border-style: solid;
}

.popover .arrow {
  border-width: 11px;
}

.popover .arrow:after {
  border-width: 10px;
  content: "";
}

.popover.top .arrow {
  bottom: -11px;
  left: 50%;
  margin-left: -11px;
  border-top-color: #999999;
  border-top-color: rgba(0, 0, 0, 0.25);
  border-bottom-width: 0;
}

.popover.top .arrow:after {
  bottom: 1px;
  margin-left: -10px;
  border-top-color: #2e3338;
  border-bottom-width: 0;
  content: " ";
}

.popover.right .arrow {
  top: 50%;
  left: -11px;
  margin-top: -11px;
  border-right-color: #999999;
  border-right-color: rgba(0, 0, 0, 0.25);
  border-left-width: 0;
}

.popover.right .arrow:after {
  bottom: -10px;
  left: 1px;
  border-right-color: #2e3338;
  border-left-width: 0;
  content: " ";
}

.popover.bottom .arrow {
  top: -11px;
  left: 50%;
  margin-left: -11px;
  border-bottom-color: #999999;
  border-bottom-color: rgba(0, 0, 0, 0.25);
  border-top-width: 0;
}

.popover.bottom .arrow:after {
  top: 1px;
  margin-left: -10px;
  border-bottom-color: #2e3338;
  border-top-width: 0;
  content: " ";
}

.popover.left .arrow {
  top: 50%;
  right: -11px;
  margin-top: -11px;
  border-left-color: #999999;
  border-left-color: rgba(0, 0, 0, 0.25);
  border-right-width: 0;
}

.popover.left .arrow:after {
  right: 1px;
  bottom: -10px;
  border-left-color: #2e3338;
  border-right-width: 0;
  content: " ";
}

.carousel {
  position: relative;
}

.carousel-inner {
  position: relative;
  width: 100%;
  overflow: hidden;
}

.carousel-inner > .item {
  position: relative;
  display: none;
  -webkit-transition: 0.6s ease-in-out left;
          transition: 0.6s ease-in-out left;
}

.carousel-inner > .item > img,
.carousel-inner > .item > a > img {
  display: block;
  height: auto;
  max-width: 100%;
  line-height: 1;
}

.carousel-inner > .active,
.carousel-inner > .next,
.carousel-inner > .prev {
  display: block;
}

.carousel-inner > .active {
  left: 0;
}

.carousel-inner > .next,
.carousel-inner > .prev {
  position: absolute;
  top: 0;
  width: 100%;
}

.carousel-inner > .next {
  left: 100%;
}

.carousel-inner > .prev {
  left: -100%;
}

.carousel-inner > .next.left,
.carousel-inner > .prev.right {
  left: 0;
}

.carousel-inner > .active.left {
  left: -100%;
}

.carousel-inner > .active.right {
  left: 100%;
}

.carousel-control {
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  width: 15%;
  font-size: 20px;
  color: #ffffff;
  text-align: center;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.6);
  opacity: 0.5;
  filter: alpha(opacity=50);
}

.carousel-control.left {
  background-image: -webkit-linear-gradient(left, color-stop(rgba(0, 0, 0, 0.5) 0), color-stop(rgba(0, 0, 0, 0.0001) 100%));
  background-image: linear-gradient(to right, rgba(0, 0, 0, 0.5) 0, rgba(0, 0, 0, 0.0001) 100%);
  background-repeat: repeat-x;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#80000000', endColorstr='#00000000', GradientType=1);
}

.carousel-control.right {
  right: 0;
  left: auto;
  background-image: -webkit-linear-gradient(left, color-stop(rgba(0, 0, 0, 0.0001) 0), color-stop(rgba(0, 0, 0, 0.5) 100%));
  background-image: linear-gradient(to right, rgba(0, 0, 0, 0.0001) 0, rgba(0, 0, 0, 0.5) 100%);
  background-repeat: repeat-x;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#00000000', endColorstr='#80000000', GradientType=1);
}

.carousel-control:hover,
.carousel-control:focus {
  color: #ffffff;
  text-decoration: none;
  outline: none;
  opacity: 0.9;
  filter: alpha(opacity=90);
}

.carousel-control .icon-prev,
.carousel-control .icon-next,
.carousel-control .glyphicon-chevron-left,
.carousel-control .glyphicon-chevron-right {
  position: absolute;
  top: 50%;
  z-index: 5;
  display: inline-block;
}

.carousel-control .icon-prev,
.carousel-control .glyphicon-chevron-left {
  left: 50%;
}

.carousel-control .icon-next,
.carousel-control .glyphicon-chevron-right {
  right: 50%;
}

.carousel-control .icon-prev,
.carousel-control .icon-next {
  width: 20px;
  height: 20px;
  margin-top: -10px;
  margin-left: -10px;
  font-family: serif;
}

.carousel-control .icon-prev:before {
  content: '\2039';
}

.carousel-control .icon-next:before {
  content: '\203a';
}

.carousel-indicators {
  position: absolute;
  bottom: 10px;
  left: 50%;
  z-index: 15;
  width: 60%;
  padding-left: 0;
  margin-left: -30%;
  text-align: center;
  list-style: none;
}

.carousel-indicators li {
  display: inline-block;
  width: 10px;
  height: 10px;
  margin: 1px;
  text-indent: -999px;
  cursor: pointer;
  background-color: #000 \9;
  background-color: rgba(0, 0, 0, 0);
  border: 1px solid #ffffff;
  border-radius: 10px;
}

.carousel-indicators .active {
  width: 12px;
  height: 12px;
  margin: 0;
  background-color: #ffffff;
}

.carousel-caption {
  position: absolute;
  right: 15%;
  bottom: 20px;
  left: 15%;
  z-index: 10;
  padding-top: 20px;
  padding-bottom: 20px;
  color: #ffffff;
  text-align: center;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.6);
}

.carousel-caption .btn {
  text-shadow: none;
}

@media screen and (min-width: 768px) {
  .carousel-control .glyphicons-chevron-left,
  .carousel-control .glyphicons-chevron-right,
  .carousel-control .icon-prev,
  .carousel-control .icon-next {
    width: 30px;
    height: 30px;
    margin-top: -15px;
    margin-left: -15px;
    font-size: 30px;
  }
  .carousel-caption {
    right: 20%;
    left: 20%;
    padding-bottom: 30px;
  }
  .carousel-indicators {
    bottom: 20px;
  }
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: " ";
}

.clearfix:after {
  clear: both;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: " ";
}

.clearfix:after {
  clear: both;
}

.center-block {
  display: block;
  margin-right: auto;
  margin-left: auto;
}

.pull-right {
  float: right !important;
}

.pull-left {
  float: left !important;
}

.hide {
  display: none !important;
}

.show {
  display: block !important;
}

.invisible {
  visibility: hidden;
}

.text-hide {
  font: 0/0 a;
  color: transparent;
  text-shadow: none;
  background-color: transparent;
  border: 0;
}

.hidden {
  display: none !important;
  visibility: hidden !important;
}

.affix {
  position: fixed;
}

@-ms-viewport {
  width: device-width;
}

.visible-xs,
tr.visible-xs,
th.visible-xs,
td.visible-xs {
  display: none !important;
}

@media (max-width: 767px) {
  .visible-xs {
    display: block !important;
  }
  table.visible-xs {
    display: table;
  }
  tr.visible-xs {
    display: table-row !important;
  }
  th.visible-xs,
  td.visible-xs {
    display: table-cell !important;
  }
}

@media (min-width: 768px) and (max-width: 991px) {
  .visible-xs.visible-sm {
    display: block !important;
  }
  table.visible-xs.visible-sm {
    display: table;
  }
  tr.visible-xs.visible-sm {
    display: table-row !important;
  }
  th.visible-xs.visible-sm,
  td.visible-xs.visible-sm {
    display: table-cell !important;
  }
}

@media (min-width: 992px) and (max-width: 1199px) {
  .visible-xs.visible-md {
    display: block !important;
  }
  table.visible-xs.visible-md {
    display: table;
  }
  tr.visible-xs.visible-md {
    display: table-row !important;
  }
  th.visible-xs.visible-md,
  td.visible-xs.visible-md {
    display: table-cell !important;
  }
}

@media (min-width: 1200px) {
  .visible-xs.visible-lg {
    display: block !important;
  }
  table.visible-xs.visible-lg {
    display: table;
  }
  tr.visible-xs.visible-lg {
    display: table-row !important;
  }
  th.visible-xs.visible-lg,
  td.visible-xs.visible-lg {
    display: table-cell !important;
  }
}

.visible-sm,
tr.visible-sm,
th.visible-sm,
td.visible-sm {
  display: none !important;
}

@media (max-width: 767px) {
  .visible-sm.visible-xs {
    display: block !important;
  }
  table.visible-sm.visible-xs {
    display: table;
  }
  tr.visible-sm.visible-xs {
    display: table-row !important;
  }
  th.visible-sm.visible-xs,
  td.visible-sm.visible-xs {
    display: table-cell !important;
  }
}

@media (min-width: 768px) and (max-width: 991px) {
  .visible-sm {
    display: block !important;
  }
  table.visible-sm {
    display: table;
  }
  tr.visible-sm {
    display: table-row !important;
  }
  th.visible-sm,
  td.visible-sm {
    display: table-cell !important;
  }
}

@media (min-width: 992px) and (max-width: 1199px) {
  .visible-sm.visible-md {
    display: block !important;
  }
  table.visible-sm.visible-md {
    display: table;
  }
  tr.visible-sm.visible-md {
    display: table-row !important;
  }
  th.visible-sm.visible-md,
  td.visible-sm.visible-md {
    display: table-cell !important;
  }
}

@media (min-width: 1200px) {
  .visible-sm.visible-lg {
    display: block !important;
  }
  table.visible-sm.visible-lg {
    display: table;
  }
  tr.visible-sm.visible-lg {
    display: table-row !important;
  }
  th.visible-sm.visible-lg,
  td.visible-sm.visible-lg {
    display: table-cell !important;
  }
}

.visible-md,
tr.visible-md,
th.visible-md,
td.visible-md {
  display: none !important;
}

@media (max-width: 767px) {
  .visible-md.visible-xs {
    display: block !important;
  }
  table.visible-md.visible-xs {
    display: table;
  }
  tr.visible-md.visible-xs {
    display: table-row !important;
  }
  th.visible-md.visible-xs,
  td.visible-md.visible-xs {
    display: table-cell !important;
  }
}

@media (min-width: 768px) and (max-width: 991px) {
  .visible-md.visible-sm {
    display: block !important;
  }
  table.visible-md.visible-sm {
    display: table;
  }
  tr.visible-md.visible-sm {
    display: table-row !important;
  }
  th.visible-md.visible-sm,
  td.visible-md.visible-sm {
    display: table-cell !important;
  }
}

@media (min-width: 992px) and (max-width: 1199px) {
  .visible-md {
    display: block !important;
  }
  table.visible-md {
    display: table;
  }
  tr.visible-md {
    display: table-row !important;
  }
  th.visible-md,
  td.visible-md {
    display: table-cell !important;
  }
}

@media (min-width: 1200px) {
  .visible-md.visible-lg {
    display: block !important;
  }
  table.visible-md.visible-lg {
    display: table;
  }
  tr.visible-md.visible-lg {
    display: table-row !important;
  }
  th.visible-md.visible-lg,
  td.visible-md.visible-lg {
    display: table-cell !important;
  }
}

.visible-lg,
tr.visible-lg,
th.visible-lg,
td.visible-lg {
  display: none !important;
}

@media (max-width: 767px) {
  .visible-lg.visible-xs {
    display: block !important;
  }
  table.visible-lg.visible-xs {
    display: table;
  }
  tr.visible-lg.visible-xs {
    display: table-row !important;
  }
  th.visible-lg.visible-xs,
  td.visible-lg.visible-xs {
    display: table-cell !important;
  }
}

@media (min-width: 768px) and (max-width: 991px) {
  .visible-lg.visible-sm {
    display: block !important;
  }
  table.visible-lg.visible-sm {
    display: table;
  }
  tr.visible-lg.visible-sm {
    display: table-row !important;
  }
  th.visible-lg.visible-sm,
  td.visible-lg.visible-sm {
    display: table-cell !important;
  }
}

@media (min-width: 992px) and (max-width: 1199px) {
  .visible-lg.visible-md {
    display: block !important;
  }
  table.visible-lg.visible-md {
    display: table;
  }
  tr.visible-lg.visible-md {
    display: table-row !important;
  }
  th.visible-lg.visible-md,
  td.visible-lg.visible-md {
    display: table-cell !important;
  }
}

@media (min-width: 1200px) {
  .visible-lg {
    display: block !important;
  }
  table.visible-lg {
    display: table;
  }
  tr.visible-lg {
    display: table-row !important;
  }
  th.visible-lg,
  td.visible-lg {
    display: table-cell !important;
  }
}

.hidden-xs {
  display: block !important;
}

table.hidden-xs {
  display: table;
}

tr.hidden-xs {
  display: table-row !important;
}

th.hidden-xs,
td.hidden-xs {
  display: table-cell !important;
}

@media (max-width: 767px) {
  .hidden-xs,
  tr.hidden-xs,
  th.hidden-xs,
  td.hidden-xs {
    display: none !important;
  }
}

@media (min-width: 768px) and (max-width: 991px) {
  .hidden-xs.hidden-sm,
  tr.hidden-xs.hidden-sm,
  th.hidden-xs.hidden-sm,
  td.hidden-xs.hidden-sm {
    display: none !important;
  }
}

@media (min-width: 992px) and (max-width: 1199px) {
  .hidden-xs.hidden-md,
  tr.hidden-xs.hidden-md,
  th.hidden-xs.hidden-md,
  td.hidden-xs.hidden-md {
    display: none !important;
  }
}

@media (min-width: 1200px) {
  .hidden-xs.hidden-lg,
  tr.hidden-xs.hidden-lg,
  th.hidden-xs.hidden-lg,
  td.hidden-xs.hidden-lg {
    display: none !important;
  }
}

.hidden-sm {
  display: block !important;
}

table.hidden-sm {
  display: table;
}

tr.hidden-sm {
  display: table-row !important;
}

th.hidden-sm,
td.hidden-sm {
  display: table-cell !important;
}

@media (max-width: 767px) {
  .hidden-sm.hidden-xs,
  tr.hidden-sm.hidden-xs,
  th.hidden-sm.hidden-xs,
  td.hidden-sm.hidden-xs {
    display: none !important;
  }
}

@media (min-width: 768px) and (max-width: 991px) {
  .hidden-sm,
  tr.hidden-sm,
  th.hidden-sm,
  td.hidden-sm {
    display: none !important;
  }
}

@media (min-width: 992px) and (max-width: 1199px) {
  .hidden-sm.hidden-md,
  tr.hidden-sm.hidden-md,
  th.hidden-sm.hidden-md,
  td.hidden-sm.hidden-md {
    display: none !important;
  }
}

@media (min-width: 1200px) {
  .hidden-sm.hidden-lg,
  tr.hidden-sm.hidden-lg,
  th.hidden-sm.hidden-lg,
  td.hidden-sm.hidden-lg {
    display: none !important;
  }
}

.hidden-md {
  display: block !important;
}

table.hidden-md {
  display: table;
}

tr.hidden-md {
  display: table-row !important;
}

th.hidden-md,
td.hidden-md {
  display: table-cell !important;
}

@media (max-width: 767px) {
  .hidden-md.hidden-xs,
  tr.hidden-md.hidden-xs,
  th.hidden-md.hidden-xs,
  td.hidden-md.hidden-xs {
    display: none !important;
  }
}

@media (min-width: 768px) and (max-width: 991px) {
  .hidden-md.hidden-sm,
  tr.hidden-md.hidden-sm,
  th.hidden-md.hidden-sm,
  td.hidden-md.hidden-sm {
    display: none !important;
  }
}

@media (min-width: 992px) and (max-width: 1199px) {
  .hidden-md,
  tr.hidden-md,
  th.hidden-md,
  td.hidden-md {
    display: none !important;
  }
}

@media (min-width: 1200px) {
  .hidden-md.hidden-lg,
  tr.hidden-md.hidden-lg,
  th.hidden-md.hidden-lg,
  td.hidden-md.hidden-lg {
    display: none !important;
  }
}

.hidden-lg {
  display: block !important;
}

table.hidden-lg {
  display: table;
}

tr.hidden-lg {
  display: table-row !important;
}

th.hidden-lg,
td.hidden-lg {
  display: table-cell !important;
}

@media (max-width: 767px) {
  .hidden-lg.hidden-xs,
  tr.hidden-lg.hidden-xs,
  th.hidden-lg.hidden-xs,
  td.hidden-lg.hidden-xs {
    display: none !important;
  }
}

@media (min-width: 768px) and (max-width: 991px) {
  .hidden-lg.hidden-sm,
  tr.hidden-lg.hidden-sm,
  th.hidden-lg.hidden-sm,
  td.hidden-lg.hidden-sm {
    display: none !important;
  }
}

@media (min-width: 992px) and (max-width: 1199px) {
  .hidden-lg.hidden-md,
  tr.hidden-lg.hidden-md,
  th.hidden-lg.hidden-md,
  td.hidden-lg.hidden-md {
    display: none !important;
  }
}

@media (min-width: 1200px) {
  .hidden-lg,
  tr.hidden-lg,
  th.hidden-lg,
  td.hidden-lg {
    display: none !important;
  }
}

.visible-print,
tr.visible-print,
th.visible-print,
td.visible-print {
  display: none !important;
}

@media print {
  .visible-print {
    display: block !important;
  }
  table.visible-print {
    display: table;
  }
  tr.visible-print {
    display: table-row !important;
  }
  th.visible-print,
  td.visible-print {
    display: table-cell !important;
  }
  .hidden-print,
  tr.hidden-print,
  th.hidden-print,
  td.hidden-print {
    display: none !important;
  }
}

.navbar {
  text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.3);
  background-image: -webkit-linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-image: linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-repeat: no-repeat;
  border: 1px solid rgba(0, 0, 0, 0.6);
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff484e55', endColorstr='#ff313539', GradientType=0);
  filter: none;
}

.navbar-inverse {
  background-image: -webkit-linear-gradient(#8a9196, #7a8288 60%, #70787d);
  background-image: linear-gradient(#8a9196, #7a8288 60%, #70787d);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff8a9196', endColorstr='#ff70787d', GradientType=0);
  filter: none;
}

.navbar-nav > li > a {
  border-right: 1px solid rgba(0, 0, 0, 0.2);
  border-left: 1px solid rgba(255, 255, 255, 0.1);
}

.navbar-nav > li > a:hover {
  background-image: -webkit-linear-gradient(#020202, #101112 40%, #191b1d);
  background-image: linear-gradient(#020202, #101112 40%, #191b1d);
  background-repeat: no-repeat;
  border-left-color: transparent;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff020202', endColorstr='#ff191b1d', GradientType=0);
  filter: none;
}

.navbar .nav .open > a {
  border-color: transparent;
}

.navbar-nav > li.active > a {
  border-left-color: transparent;
}

.navbar-form {
  margin-right: 5px;
  margin-left: 5px;
}

.btn,
.btn:hover {
  text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.3);
  border-color: rgba(0, 0, 0, 0.6);
}

.btn-default {
  background-image: -webkit-linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-image: linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff484e55', endColorstr='#ff313539', GradientType=0);
  filter: none;
}

.btn-primary {
  background-image: -webkit-linear-gradient(#8a9196, #7a8288 60%, #70787d);
  background-image: linear-gradient(#8a9196, #7a8288 60%, #70787d);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff8a9196', endColorstr='#ff70787d', GradientType=0);
  filter: none;
}

.btn-success {
  background-image: -webkit-linear-gradient(#78cc78, #62c462 60%, #53be53);
  background-image: linear-gradient(#78cc78, #62c462 60%, #53be53);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff78cc78', endColorstr='#ff53be53', GradientType=0);
  filter: none;
}

.btn-info {
  background-image: -webkit-linear-gradient(#74cae3, #5bc0de 60%, #4ab9db);
  background-image: linear-gradient(#74cae3, #5bc0de 60%, #4ab9db);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff74cae3', endColorstr='#ff4ab9db', GradientType=0);
  filter: none;
}

.btn-warning {
  background-image: -webkit-linear-gradient(#faa123, #f89406 60%, #e48806);
  background-image: linear-gradient(#faa123, #f89406 60%, #e48806);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#fffaa123', endColorstr='#ffe48806', GradientType=0);
  filter: none;
}

.btn-danger {
  background-image: -webkit-linear-gradient(#f17a77, #ee5f5b 60%, #ec4d49);
  background-image: linear-gradient(#f17a77, #ee5f5b 60%, #ec4d49);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#fff17a77', endColorstr='#ffec4d49', GradientType=0);
  filter: none;
}

.btn-default:hover {
  background-image: -webkit-linear-gradient(#020202, #101112 40%, #191b1d);
  background-image: linear-gradient(#020202, #101112 40%, #191b1d);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff020202', endColorstr='#ff191b1d', GradientType=0);
  filter: none;
}

.btn-primary:hover {
  background-image: -webkit-linear-gradient(#404448, #4e5458 40%, #585e62);
  background-image: linear-gradient(#404448, #4e5458 40%, #585e62);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff404448', endColorstr='#ff585e62', GradientType=0);
  filter: none;
}

.btn-success:hover {
  background-image: -webkit-linear-gradient(#2f7d2f, #379337 40%, #3da23d);
  background-image: linear-gradient(#2f7d2f, #379337 40%, #3da23d);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff2f7d2f', endColorstr='#ff3da23d', GradientType=0);
  filter: none;
}

.btn-info:hover {
  background-image: -webkit-linear-gradient(#20829f, #2596b8 40%, #28a4c9);
  background-image: linear-gradient(#20829f, #2596b8 40%, #28a4c9);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff20829f', endColorstr='#ff28a4c9', GradientType=0);
  filter: none;
}

.btn-warning:hover {
  background-image: -webkit-linear-gradient(#804d03, #9e5f04 40%, #b26a04);
  background-image: linear-gradient(#804d03, #9e5f04 40%, #b26a04);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff804d03', endColorstr='#ffb26a04', GradientType=0);
  filter: none;
}

.btn-danger:hover {
  background-image: -webkit-linear-gradient(#bb1813, #d71c16 40%, #e7201a);
  background-image: linear-gradient(#bb1813, #d71c16 40%, #e7201a);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ffbb1813', endColorstr='#ffe7201a', GradientType=0);
  filter: none;
}

h1,
h2,
h3,
h4,
h5,
h6 {
  text-shadow: -1px -1px 0 rgba(0, 0, 0, 0.3);
}

.text-primary,
.text-primary:hover {
  color: #7a8288;
}

.text-success,
.text-success:hover {
  color: #62c462;
}

.text-danger,
.text-danger:hover {
  color: #ee5f5b;
}

.text-warning,
.text-warning:hover {
  color: #f89406;
}

.text-info,
.text-info:hover {
  color: #5bc0de;
}

.table tr.success,
.table tr.warning,
.table tr.danger {
  color: #fff;
}

.table-bordered tbody tr.success td,
.table-bordered tbody tr.warning td,
.table-bordered tbody tr.danger td,
.table-bordered tbody tr.success:hover td,
.table-bordered tbody tr.warning:hover td,
.table-bordered tbody tr.danger:hover td {
  border-color: #1c1e22;
}

.table-responsive > .table {
  background-color: #2e3338;
}

.has-warning .help-block,
.has-warning .control-label {
  color: #f89406;
}

.has-warning .form-control,
.has-warning .form-control:focus {
  border-color: #f89406;
}

.has-error .help-block,
.has-error .control-label {
  color: #ee5f5b;
}

.has-error .form-control,
.has-error .form-control:focus {
  border-color: #ee5f5b;
}

.has-success .help-block,
.has-success .control-label {
  color: #62c462;
}

.has-success .form-control,
.has-success .form-control:focus {
  border-color: #62c462;
}

legend {
  color: #fff;
}

.input-group-addon {
  color: #ffffff;
  text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.3);
  background-image: -webkit-linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-image: linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-repeat: no-repeat;
  border-color: rgba(0, 0, 0, 0.6);
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff484e55', endColorstr='#ff313539', GradientType=0);
  filter: none;
}

.nav .open > a,
.nav .open > a:hover,
.nav .open > a:focus {
  border-color: rgba(0, 0, 0, 0.6);
}

.nav-pills > li > a {
  text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.3);
  background-image: -webkit-linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-image: linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-repeat: no-repeat;
  border: 1px solid rgba(0, 0, 0, 0.6);
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff484e55', endColorstr='#ff313539', GradientType=0);
  filter: none;
}

.nav-pills > li > a:hover {
  background-image: -webkit-linear-gradient(#020202, #101112 40%, #191b1d);
  background-image: linear-gradient(#020202, #101112 40%, #191b1d);
  background-repeat: no-repeat;
  border: 1px solid rgba(0, 0, 0, 0.6);
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff020202', endColorstr='#ff191b1d', GradientType=0);
  filter: none;
}

.nav-pills > li.active > a,
.nav-pills > li.active > a:hover {
  background-color: none;
  background-image: -webkit-linear-gradient(#020202, #101112 40%, #191b1d);
  background-image: linear-gradient(#020202, #101112 40%, #191b1d);
  background-repeat: no-repeat;
  border: 1px solid rgba(0, 0, 0, 0.6);
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff020202', endColorstr='#ff191b1d', GradientType=0);
  filter: none;
}

.nav-pills > li.disabled > a,
.nav-pills > li.disabled > a:hover {
  background-image: -webkit-linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-image: linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff484e55', endColorstr='#ff313539', GradientType=0);
  filter: none;
}

.pagination > li > a {
  text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.3);
  background-image: -webkit-linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-image: linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff484e55', endColorstr='#ff313539', GradientType=0);
  filter: none;
}

.pagination > li > a:hover {
  background-image: -webkit-linear-gradient(#020202, #101112 40%, #191b1d);
  background-image: linear-gradient(#020202, #101112 40%, #191b1d);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff020202', endColorstr='#ff191b1d', GradientType=0);
  filter: none;
}

.pagination > li.active > a {
  background-image: -webkit-linear-gradient(#020202, #101112 40%, #191b1d);
  background-image: linear-gradient(#020202, #101112 40%, #191b1d);
  background-repeat: no-repeat;
  border-color: rgba(0, 0, 0, 0.6);
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff020202', endColorstr='#ff191b1d', GradientType=0);
  filter: none;
}

.pagination > li.disabled > a,
.pagination > li.disabled > a:hover {
  background-color: transparent;
  background-image: -webkit-linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-image: linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff484e55', endColorstr='#ff313539', GradientType=0);
  filter: none;
}

.pager > li > a {
  text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.3);
  background-image: -webkit-linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-image: linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-repeat: no-repeat;
  border: 1px solid rgba(0, 0, 0, 0.6);
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff484e55', endColorstr='#ff313539', GradientType=0);
  filter: none;
}

.pager > li > a:hover {
  background-image: -webkit-linear-gradient(#020202, #101112 40%, #191b1d);
  background-image: linear-gradient(#020202, #101112 40%, #191b1d);
  background-repeat: no-repeat;
  border: 1px solid rgba(0, 0, 0, 0.6);
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff020202', endColorstr='#ff191b1d', GradientType=0);
  filter: none;
}

.pager > li.disabled > a,
.pager > li.disabled > a:hover {
  background-color: transparent;
  background-image: -webkit-linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-image: linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-repeat: no-repeat;
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff484e55', endColorstr='#ff313539', GradientType=0);
  filter: none;
}

.breadcrumb {
  text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.3);
  background-image: -webkit-linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-image: linear-gradient(#484e55, #3a3f44 60%, #313539);
  background-repeat: no-repeat;
  border: 1px solid rgba(0, 0, 0, 0.6);
  filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#ff484e55', endColorstr='#ff313539', GradientType=0);
  filter: none;
}

.alert .alert-link,
.alert a {
  color: #fff;
  text-decoration: underline;
}

.jumbotron {
  border: 1px solid rgba(0, 0, 0, 0.6);
}

.list-group-item {
  background-color: #32383e;
}

.panel-primary .panel-heading,
.panel-success .panel-heading,
.panel-danger .panel-heading,
.panel-warning .panel-heading,
.panel-info .panel-heading {
  border-color: #000;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: " ";
}

.clearfix:after {
  clear: both;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: " ";
}

.clearfix:after {
  clear: both;
}

.center-block {
  display: block;
  margin-right: auto;
  margin-left: auto;
}

.pull-right {
  float: right !important;
}

.pull-left {
  float: left !important;
}

.hide {
  display: none !important;
}

.show {
  display: block !important;
}

.invisible {
  visibility: hidden;
}

.text-hide {
  font: 0/0 a;
  color: transparent;
  text-shadow: none;
  background-color: transparent;
  border: 0;
}

.hidden {
  display: none !important;
  visibility: hidden !important;
}

.affix {
  position: fixed;
}`

var highchartsMoreJS string = `
/*
 Highcharts JS v7.0.0 (2018-12-11)

 (c) 2009-2018 Torstein Honsi

 License: www.highcharts.com/license
*/
(function(y){"object"===typeof module&&module.exports?module.exports=y:"function"===typeof define&&define.amd?define(function(){return y}):y("undefined"!==typeof Highcharts?Highcharts:void 0)})(function(y){(function(a){function v(a,d){this.init(a,d)}var w=a.CenteredSeriesMixin,e=a.extend,m=a.merge,u=a.splat;e(v.prototype,{coll:"pane",init:function(a,d){this.chart=d;this.background=[];d.pane.push(this);this.setOptions(a)},setOptions:function(a){this.options=m(this.defaultOptions,this.chart.angular?
{background:{}}:void 0,a)},render:function(){var a=this.options,d=this.options.background,c=this.chart.renderer;this.group||(this.group=c.g("pane-group").attr({zIndex:a.zIndex||0}).add());this.updateCenter();if(d)for(d=u(d),a=Math.max(d.length,this.background.length||0),c=0;c<a;c++)d[c]&&this.axis?this.renderBackground(m(this.defaultBackgroundOptions,d[c]),c):this.background[c]&&(this.background[c]=this.background[c].destroy(),this.background.splice(c,1))},renderBackground:function(a,d){var c="animate",
n={"class":"highcharts-pane "+(a.className||"")};this.chart.styledMode||e(n,{fill:a.backgroundColor,stroke:a.borderColor,"stroke-width":a.borderWidth});this.background[d]||(this.background[d]=this.chart.renderer.path().add(this.group),c="attr");this.background[d][c]({d:this.axis.getPlotBandPath(a.from,a.to,a)}).attr(n)},defaultOptions:{center:["50%","50%"],size:"85%",startAngle:0},defaultBackgroundOptions:{shape:"circle",borderWidth:1,borderColor:"#cccccc",backgroundColor:{linearGradient:{x1:0,y1:0,
x2:0,y2:1},stops:[[0,"#ffffff"],[1,"#e6e6e6"]]},from:-Number.MAX_VALUE,innerRadius:0,to:Number.MAX_VALUE,outerRadius:"105%"},updateCenter:function(a){this.center=(a||this.axis||{}).center=w.getCenter.call(this)},update:function(a,d){m(!0,this.options,a);this.setOptions(this.options);this.render();this.chart.axes.forEach(function(c){c.pane===this&&(c.pane=null,c.update({},d))},this)}});a.Pane=v})(y);(function(a){var v=a.addEvent,w=a.Axis,e=a.extend,m=a.merge,u=a.noop,q=a.pick,d=a.pInt,c=a.Tick,n=a.wrap,
b=a.correctFloat,l,k,h=w.prototype,f=c.prototype;a.radialAxisExtended||(a.radialAxisExtended=!0,l={getOffset:u,redraw:function(){this.isDirty=!1},render:function(){this.isDirty=!1},setScale:u,setCategories:u,setTitle:u},k={defaultRadialGaugeOptions:{labels:{align:"center",x:0,y:null},minorGridLineWidth:0,minorTickInterval:"auto",minorTickLength:10,minorTickPosition:"inside",minorTickWidth:1,tickLength:10,tickPosition:"inside",tickWidth:2,title:{rotation:0},zIndex:2},defaultRadialXOptions:{gridLineWidth:1,
labels:{align:null,distance:15,x:0,y:null,style:{textOverflow:"none"}},maxPadding:0,minPadding:0,showLastLabel:!1,tickLength:0},defaultRadialYOptions:{gridLineInterpolation:"circle",labels:{align:"right",x:-3,y:-2},showLastLabel:!1,title:{x:4,text:null,rotation:90}},setOptions:function(b){b=this.options=m(this.defaultOptions,this.defaultRadialOptions,b);b.plotBands||(b.plotBands=[]);a.fireEvent(this,"afterSetOptions")},getOffset:function(){h.getOffset.call(this);this.chart.axisOffset[this.side]=0},
getLinePath:function(b,c){b=this.center;var g=this.chart,p=q(c,b[2]/2-this.offset);this.isCircular||void 0!==c?(c=this.chart.renderer.symbols.arc(this.left+b[0],this.top+b[1],p,p,{start:this.startAngleRad,end:this.endAngleRad,open:!0,innerR:0}),c.xBounds=[this.left+b[0]],c.yBounds=[this.top+b[1]-p]):(c=this.postTranslate(this.angleRad,p),c=["M",b[0]+g.plotLeft,b[1]+g.plotTop,"L",c.x,c.y]);return c},setAxisTranslation:function(){h.setAxisTranslation.call(this);this.center&&(this.transA=this.isCircular?
(this.endAngleRad-this.startAngleRad)/(this.max-this.min||1):this.center[2]/2/(this.max-this.min||1),this.minPixelPadding=this.isXAxis?this.transA*this.minPointOffset:0)},beforeSetTickPositions:function(){if(this.autoConnect=this.isCircular&&void 0===q(this.userMax,this.options.max)&&b(this.endAngleRad-this.startAngleRad)===b(2*Math.PI))this.max+=this.categories&&1||this.pointRange||this.closestPointRange||0},setAxisSize:function(){h.setAxisSize.call(this);this.isRadial&&(this.pane.updateCenter(this),
this.isCircular&&(this.sector=this.endAngleRad-this.startAngleRad),this.len=this.width=this.height=this.center[2]*q(this.sector,1)/2)},getPosition:function(b,c){return this.postTranslate(this.isCircular?this.translate(b):this.angleRad,q(this.isCircular?c:this.translate(b),this.center[2]/2)-this.offset)},postTranslate:function(b,c){var g=this.chart,p=this.center;b=this.startAngleRad+b;return{x:g.plotLeft+p[0]+Math.cos(b)*c,y:g.plotTop+p[1]+Math.sin(b)*c}},getPlotBandPath:function(b,c,g){var p=this.center,
n=this.startAngleRad,l=p[2]/2,a=[q(g.outerRadius,"100%"),g.innerRadius,q(g.thickness,10)],t=Math.min(this.offset,0),h=/%$/,k,f=this.isCircular;"polygon"===this.options.gridLineInterpolation?p=this.getPlotLinePath(b).concat(this.getPlotLinePath(c,!0)):(b=Math.max(b,this.min),c=Math.min(c,this.max),f||(a[0]=this.translate(b),a[1]=this.translate(c)),a=a.map(function(g){h.test(g)&&(g=d(g,10)*l/100);return g}),"circle"!==g.shape&&f?(b=n+this.translate(b),c=n+this.translate(c)):(b=-Math.PI/2,c=1.5*Math.PI,
k=!0),a[0]-=t,a[2]-=t,p=this.chart.renderer.symbols.arc(this.left+p[0],this.top+p[1],a[0],a[0],{start:Math.min(b,c),end:Math.max(b,c),innerR:q(a[1],a[0]-a[2]),open:k}));return p},getPlotLinePath:function(b,c){var g=this,p=g.center,n=g.chart,d=g.getPosition(b),a,l,t;g.isCircular?t=["M",p[0]+n.plotLeft,p[1]+n.plotTop,"L",d.x,d.y]:"circle"===g.options.gridLineInterpolation?(b=g.translate(b),t=g.getLinePath(0,b)):(n.xAxis.forEach(function(b){b.pane===g.pane&&(a=b)}),t=[],b=g.translate(b),p=a.tickPositions,
a.autoConnect&&(p=p.concat([p[0]])),c&&(p=[].concat(p).reverse()),p.forEach(function(g,c){l=a.getPosition(g,b);t.push(c?"L":"M",l.x,l.y)}));return t},getTitlePosition:function(){var b=this.center,c=this.chart,g=this.options.title;return{x:c.plotLeft+b[0]+(g.x||0),y:c.plotTop+b[1]-{high:.5,middle:.25,low:0}[g.align]*b[2]+(g.y||0)}}},v(w,"init",function(b){var c=this.chart,g=c.angular,p=c.polar,n=this.isXAxis,d=g&&n,a,h=c.options;b=b.userOptions.pane||0;b=this.pane=c.pane&&c.pane[b];if(g){if(e(this,
d?l:k),a=!n)this.defaultRadialOptions=this.defaultRadialGaugeOptions}else p&&(e(this,k),this.defaultRadialOptions=(a=n)?this.defaultRadialXOptions:m(this.defaultYAxisOptions,this.defaultRadialYOptions));g||p?(this.isRadial=!0,c.inverted=!1,h.chart.zoomType=null):this.isRadial=!1;b&&a&&(b.axis=this);this.isCircular=a}),v(w,"afterInit",function(){var b=this.chart,c=this.options,g=this.pane,p=g&&g.options;b.angular&&this.isXAxis||!g||!b.angular&&!b.polar||(this.angleRad=(c.angle||0)*Math.PI/180,this.startAngleRad=
(p.startAngle-90)*Math.PI/180,this.endAngleRad=(q(p.endAngle,p.startAngle+360)-90)*Math.PI/180,this.offset=c.offset||0)}),n(h,"autoLabelAlign",function(b){if(!this.isRadial)return b.apply(this,[].slice.call(arguments,1))}),v(c,"afterGetPosition",function(b){this.axis.getPosition&&e(b.pos,this.axis.getPosition(this.pos))}),v(c,"afterGetLabelPosition",function(b){var c=this.axis,g=this.label,p=c.options.labels,n=p.y,d,a=20,l=p.align,h=(c.translate(this.pos)+c.startAngleRad+Math.PI/2)/Math.PI*180%360;
c.isRadial&&(d=c.getPosition(this.pos,c.center[2]/2+q(p.distance,-25)),"auto"===p.rotation?g.attr({rotation:h}):null===n&&(n=c.chart.renderer.fontMetrics(g.styles&&g.styles.fontSize).b-g.getBBox().height/2),null===l&&(c.isCircular?(this.label.getBBox().width>c.len*c.tickInterval/(c.max-c.min)&&(a=0),l=h>a&&h<180-a?"left":h>180+a&&h<360-a?"right":"center"):l="center",g.attr({align:l})),b.pos.x=d.x+p.x,b.pos.y=d.y+n)}),n(f,"getMarkPath",function(b,c,g,p,n,d,a){var l=this.axis;l.isRadial?(b=l.getPosition(this.pos,
l.center[2]/2+p),c=["M",c,g,"L",b.x,b.y]):c=b.call(this,c,g,p,n,d,a);return c}))})(y);(function(a){var v=a.pick,w=a.extend,e=a.isArray,m=a.defined,u=a.seriesType,q=a.seriesTypes,d=a.Series.prototype,c=a.Point.prototype;u("arearange","area",{lineWidth:1,threshold:null,tooltip:{pointFormat:'\x3cspan style\x3d"color:{series.color}"\x3e\u25cf\x3c/span\x3e {series.name}: \x3cb\x3e{point.low}\x3c/b\x3e - \x3cb\x3e{point.high}\x3c/b\x3e\x3cbr/\x3e'},trackByArea:!0,dataLabels:{align:null,verticalAlign:null,
xLow:0,xHigh:0,yLow:0,yHigh:0}},{pointArrayMap:["low","high"],toYData:function(c){return[c.low,c.high]},pointValKey:"low",deferTranslatePolar:!0,highToXY:function(c){var b=this.chart,n=this.xAxis.postTranslate(c.rectPlotX,this.yAxis.len-c.plotHigh);c.plotHighX=n.x-b.plotLeft;c.plotHigh=n.y-b.plotTop;c.plotLowX=c.plotX},translate:function(){var c=this,b=c.yAxis,d=!!c.modifyValue;q.area.prototype.translate.apply(c);c.points.forEach(function(a){var n=a.low,l=a.high,k=a.plotY;null===l||null===n?(a.isNull=
!0,a.plotY=null):(a.plotLow=k,a.plotHigh=b.translate(d?c.modifyValue(l,a):l,0,1,0,1),d&&(a.yBottom=a.plotHigh))});this.chart.polar&&this.points.forEach(function(b){c.highToXY(b);b.tooltipPos=[(b.plotHighX+b.plotLowX)/2,(b.plotHigh+b.plotLow)/2]})},getGraphPath:function(c){var b=[],a=[],d,n=q.area.prototype.getGraphPath,f,r,t;t=this.options;var g=this.chart.polar&&!1!==t.connectEnds,p=t.connectNulls,x=t.step;c=c||this.points;for(d=c.length;d--;)f=c[d],f.isNull||g||p||c[d+1]&&!c[d+1].isNull||a.push({plotX:f.plotX,
plotY:f.plotY,doCurve:!1}),r={polarPlotY:f.polarPlotY,rectPlotX:f.rectPlotX,yBottom:f.yBottom,plotX:v(f.plotHighX,f.plotX),plotY:f.plotHigh,isNull:f.isNull},a.push(r),b.push(r),f.isNull||g||p||c[d-1]&&!c[d-1].isNull||a.push({plotX:f.plotX,plotY:f.plotY,doCurve:!1});c=n.call(this,c);x&&(!0===x&&(x="left"),t.step={left:"right",center:"center",right:"left"}[x]);b=n.call(this,b);a=n.call(this,a);t.step=x;t=[].concat(c,b);this.chart.polar||"M"!==a[0]||(a[0]="L");this.graphPath=t;this.areaPath=c.concat(a);
t.isArea=!0;t.xMap=c.xMap;this.areaPath.xMap=c.xMap;return t},drawDataLabels:function(){var c=this.points,b=c.length,a,k=[],h=this.options.dataLabels,f,r,t=this.chart.inverted,g,p;e(h)?1<h.length?(g=h[0],p=h[1]):(g=h[0],p={enabled:!1}):(g=w({},h),g.x=h.xHigh,g.y=h.yHigh,p=w({},h),p.x=h.xLow,p.y=h.yLow);if(g.enabled||this._hasPointLabels){for(a=b;a--;)if(f=c[a])r=g.inside?f.plotHigh<f.plotLow:f.plotHigh>f.plotLow,f.y=f.high,f._plotY=f.plotY,f.plotY=f.plotHigh,k[a]=f.dataLabel,f.dataLabel=f.dataLabelUpper,
f.below=r,t?g.align||(g.align=r?"right":"left"):g.verticalAlign||(g.verticalAlign=r?"top":"bottom");this.options.dataLabels=g;d.drawDataLabels&&d.drawDataLabels.apply(this,arguments);for(a=b;a--;)if(f=c[a])f.dataLabelUpper=f.dataLabel,f.dataLabel=k[a],delete f.dataLabels,f.y=f.low,f.plotY=f._plotY}if(p.enabled||this._hasPointLabels){for(a=b;a--;)if(f=c[a])r=p.inside?f.plotHigh<f.plotLow:f.plotHigh>f.plotLow,f.below=!r,t?p.align||(p.align=r?"left":"right"):p.verticalAlign||(p.verticalAlign=r?"bottom":
"top");this.options.dataLabels=p;d.drawDataLabels&&d.drawDataLabels.apply(this,arguments)}if(g.enabled)for(a=b;a--;)if(f=c[a])f.dataLabels=[f.dataLabel,f.dataLabelUpper].filter(function(b){return!!b});this.options.dataLabels=h},alignDataLabel:function(){q.column.prototype.alignDataLabel.apply(this,arguments)},drawPoints:function(){var c=this.points.length,b,l;d.drawPoints.apply(this,arguments);for(l=0;l<c;)b=this.points[l],b.origProps={plotY:b.plotY,plotX:b.plotX,isInside:b.isInside,negative:b.negative,
zone:b.zone,y:b.y},b.lowerGraphic=b.graphic,b.graphic=b.upperGraphic,b.plotY=b.plotHigh,m(b.plotHighX)&&(b.plotX=b.plotHighX),b.y=b.high,b.negative=b.high<(this.options.threshold||0),b.zone=this.zones.length&&b.getZone(),this.chart.polar||(b.isInside=b.isTopInside=void 0!==b.plotY&&0<=b.plotY&&b.plotY<=this.yAxis.len&&0<=b.plotX&&b.plotX<=this.xAxis.len),l++;d.drawPoints.apply(this,arguments);for(l=0;l<c;)b=this.points[l],b.upperGraphic=b.graphic,b.graphic=b.lowerGraphic,a.extend(b,b.origProps),delete b.origProps,
l++},setStackedPoints:a.noop},{setState:function(){var a=this.state,b=this.series,d=b.chart.polar;m(this.plotHigh)||(this.plotHigh=b.yAxis.toPixels(this.high,!0));m(this.plotLow)||(this.plotLow=this.plotY=b.yAxis.toPixels(this.low,!0));b.stateMarkerGraphic&&(b.lowerStateMarkerGraphic=b.stateMarkerGraphic,b.stateMarkerGraphic=b.upperStateMarkerGraphic);this.graphic=this.upperGraphic;this.plotY=this.plotHigh;d&&(this.plotX=this.plotHighX);c.setState.apply(this,arguments);this.state=a;this.plotY=this.plotLow;
this.graphic=this.lowerGraphic;d&&(this.plotX=this.plotLowX);b.stateMarkerGraphic&&(b.upperStateMarkerGraphic=b.stateMarkerGraphic,b.stateMarkerGraphic=b.lowerStateMarkerGraphic,b.lowerStateMarkerGraphic=void 0);c.setState.apply(this,arguments)},haloPath:function(){var a=this.series.chart.polar,b=[];this.plotY=this.plotLow;a&&(this.plotX=this.plotLowX);this.isInside&&(b=c.haloPath.apply(this,arguments));this.plotY=this.plotHigh;a&&(this.plotX=this.plotHighX);this.isTopInside&&(b=b.concat(c.haloPath.apply(this,
arguments)));return b},destroyElements:function(){["lowerGraphic","upperGraphic"].forEach(function(c){this[c]&&(this[c]=this[c].destroy())},this);this.graphic=null;return c.destroyElements.apply(this,arguments)}})})(y);(function(a){var v=a.seriesType;v("areasplinerange","arearange",null,{getPointSpline:a.seriesTypes.spline.prototype.getPointSpline})})(y);(function(a){var v=a.defaultPlotOptions,w=a.merge,e=a.noop,m=a.pick,u=a.seriesType,q=a.seriesTypes.column.prototype;u("columnrange","arearange",
w(v.column,v.arearange,{pointRange:null,marker:null,states:{hover:{halo:!1}}}),{translate:function(){var a=this,c=a.yAxis,n=a.xAxis,b=n.startAngleRad,l,k=a.chart,h=a.xAxis.isRadial,f=Math.max(k.chartWidth,k.chartHeight)+999,r;q.translate.apply(a);a.points.forEach(function(d){var g=d.shapeArgs,p=a.options.minPointLength,x,t;d.plotHigh=r=Math.min(Math.max(-f,c.translate(d.high,0,1,0,1)),f);d.plotLow=Math.min(Math.max(-f,d.plotY),f);t=r;x=m(d.rectPlotY,d.plotY)-r;Math.abs(x)<p?(p-=x,x+=p,t-=p/2):0>x&&
(x*=-1,t-=x);h?(l=d.barX+b,d.shapeType="path",d.shapeArgs={d:a.polarArc(t+x,t,l,l+d.pointWidth)}):(g.height=x,g.y=t,d.tooltipPos=k.inverted?[c.len+c.pos-k.plotLeft-t-x/2,n.len+n.pos-k.plotTop-g.x-g.width/2,x]:[n.left-k.plotLeft+g.x+g.width/2,c.pos-k.plotTop+t+x/2,x])})},directTouch:!0,trackerGroups:["group","dataLabelsGroup"],drawGraph:e,getSymbol:e,crispCol:function(){return q.crispCol.apply(this,arguments)},drawPoints:function(){return q.drawPoints.apply(this,arguments)},drawTracker:function(){return q.drawTracker.apply(this,
arguments)},getColumnMetrics:function(){return q.getColumnMetrics.apply(this,arguments)},pointAttribs:function(){return q.pointAttribs.apply(this,arguments)},animate:function(){return q.animate.apply(this,arguments)},polarArc:function(){return q.polarArc.apply(this,arguments)},translate3dPoints:function(){return q.translate3dPoints.apply(this,arguments)},translate3dShapes:function(){return q.translate3dShapes.apply(this,arguments)}},{setState:q.pointClass.prototype.setState})})(y);(function(a){var v=
a.pick,w=a.seriesType,e=a.seriesTypes.column.prototype;w("columnpyramid","column",{},{translate:function(){var a=this,u=a.chart,q=a.options,d=a.dense=2>a.closestPointRange*a.xAxis.transA,d=a.borderWidth=v(q.borderWidth,d?0:1),c=a.yAxis,n=q.threshold,b=a.translatedThreshold=c.getThreshold(n),l=v(q.minPointLength,5),k=a.getColumnMetrics(),h=k.width,f=a.barW=Math.max(h,1+2*d),r=a.pointXOffset=k.offset;u.inverted&&(b-=.5);q.pointPadding&&(f=Math.ceil(f));e.translate.apply(a);a.points.forEach(function(d){var g=
v(d.yBottom,b),p=999+Math.abs(g),k=Math.min(Math.max(-p,d.plotY),c.len+p),p=d.plotX+r,t=f/2,e=Math.min(k,g),g=Math.max(k,g)-e,m,z,A,w,C,D;d.barX=p;d.pointWidth=h;d.tooltipPos=u.inverted?[c.len+c.pos-u.plotLeft-k,a.xAxis.len-p-t,g]:[p+t,k+c.pos-u.plotTop,g];k=n+(d.total||d.y);"percent"===q.stacking&&(k=n+(0>d.y)?-100:100);k=c.toPixels(k,!0);m=u.plotHeight-k-(u.plotHeight-b);z=t*(e-k)/m;A=t*(e+g-k)/m;m=p-z+t;z=p+z+t;w=p+A+t;A=p-A+t;C=e-l;D=e+g;0>d.y&&(C=e,D=e+g+l);u.inverted&&(w=u.plotWidth-e,m=k-(u.plotWidth-
b),z=t*(k-w)/m,A=t*(k-(w-g))/m,m=p+t+z,z=m-2*z,w=p-A+t,A=p+A+t,C=e,D=e+g-l,0>d.y&&(D=e+g+l));d.shapeType="path";d.shapeArgs={x:m,y:C,width:z-m,height:g,d:["M",m,C,"L",z,C,w,D,A,D,"Z"]}})}})})(y);(function(a){var v=a.isNumber,w=a.merge,e=a.pick,m=a.pInt,u=a.Series,q=a.seriesType,d=a.TrackerMixin;q("gauge","line",{dataLabels:{enabled:!0,defer:!1,y:15,borderRadius:3,crop:!1,verticalAlign:"top",zIndex:2,borderWidth:1,borderColor:"#cccccc"},dial:{},pivot:{},tooltip:{headerFormat:""},showInLegend:!1},{angular:!0,
directTouch:!0,drawGraph:a.noop,fixedBox:!0,forceDL:!0,noSharedTooltip:!0,trackerGroups:["group","dataLabelsGroup"],translate:function(){var c=this.yAxis,a=this.options,b=c.center;this.generatePoints();this.points.forEach(function(d){var n=w(a.dial,d.dial),h=m(e(n.radius,80))*b[2]/200,f=m(e(n.baseLength,70))*h/100,l=m(e(n.rearLength,10))*h/100,t=n.baseWidth||3,g=n.topWidth||1,p=a.overshoot,x=c.startAngleRad+c.translate(d.y,null,null,null,!0);v(p)?(p=p/180*Math.PI,x=Math.max(c.startAngleRad-p,Math.min(c.endAngleRad+
p,x))):!1===a.wrap&&(x=Math.max(c.startAngleRad,Math.min(c.endAngleRad,x)));x=180*x/Math.PI;d.shapeType="path";d.shapeArgs={d:n.path||["M",-l,-t/2,"L",f,-t/2,h,-g/2,h,g/2,f,t/2,-l,t/2,"z"],translateX:b[0],translateY:b[1],rotation:x};d.plotX=b[0];d.plotY=b[1]})},drawPoints:function(){var c=this,a=c.chart,b=c.yAxis.center,d=c.pivot,k=c.options,h=k.pivot,f=a.renderer;c.points.forEach(function(b){var d=b.graphic,g=b.shapeArgs,p=g.d,n=w(k.dial,b.dial);d?(d.animate(g),g.d=p):(b.graphic=f[b.shapeType](g).attr({rotation:g.rotation,
zIndex:1}).addClass("highcharts-dial").add(c.group),a.styledMode||b.graphic.attr({stroke:n.borderColor||"none","stroke-width":n.borderWidth||0,fill:n.backgroundColor||"#000000"}))});d?d.animate({translateX:b[0],translateY:b[1]}):(c.pivot=f.circle(0,0,e(h.radius,5)).attr({zIndex:2}).addClass("highcharts-pivot").translate(b[0],b[1]).add(c.group),a.styledMode||c.pivot.attr({"stroke-width":h.borderWidth||0,stroke:h.borderColor||"#cccccc",fill:h.backgroundColor||"#000000"}))},animate:function(c){var a=
this;c||(a.points.forEach(function(b){var c=b.graphic;c&&(c.attr({rotation:180*a.yAxis.startAngleRad/Math.PI}),c.animate({rotation:b.shapeArgs.rotation},a.options.animation))}),a.animate=null)},render:function(){this.group=this.plotGroup("group","series",this.visible?"visible":"hidden",this.options.zIndex,this.chart.seriesGroup);u.prototype.render.call(this);this.group.clip(this.chart.clipRect)},setData:function(c,a){u.prototype.setData.call(this,c,!1);this.processData();this.generatePoints();e(a,
!0)&&this.chart.redraw()},drawTracker:d&&d.drawTrackerPoint},{setState:function(c){this.state=c}})})(y);(function(a){var v=a.noop,w=a.pick,e=a.seriesType,m=a.seriesTypes;e("boxplot","column",{threshold:null,tooltip:{pointFormat:'\x3cspan style\x3d"color:{point.color}"\x3e\u25cf\x3c/span\x3e \x3cb\x3e {series.name}\x3c/b\x3e\x3cbr/\x3eMaximum: {point.high}\x3cbr/\x3eUpper quartile: {point.q3}\x3cbr/\x3eMedian: {point.median}\x3cbr/\x3eLower quartile: {point.q1}\x3cbr/\x3eMinimum: {point.low}\x3cbr/\x3e'},
whiskerLength:"50%",fillColor:"#ffffff",lineWidth:1,medianWidth:2,whiskerWidth:2},{pointArrayMap:["low","q1","median","q3","high"],toYData:function(a){return[a.low,a.q1,a.median,a.q3,a.high]},pointValKey:"high",pointAttribs:function(){return{}},drawDataLabels:v,translate:function(){var a=this.yAxis,e=this.pointArrayMap;m.column.prototype.translate.apply(this);this.points.forEach(function(d){e.forEach(function(c){null!==d[c]&&(d[c+"Plot"]=a.translate(d[c],0,1,0,1))})})},drawPoints:function(){var a=
this,e=a.options,d=a.chart,c=d.renderer,n,b,l,k,h,f,m=0,t,g,p,x,E=!1!==a.doQuartiles,F,B=a.options.whiskerLength;a.points.forEach(function(r){var q=r.graphic,u=q?"animate":"attr",z=r.shapeArgs,v={},H={},J={},y={},G=r.color||a.color;void 0!==r.plotY&&(t=z.width,g=Math.floor(z.x),p=g+t,x=Math.round(t/2),n=Math.floor(E?r.q1Plot:r.lowPlot),b=Math.floor(E?r.q3Plot:r.lowPlot),l=Math.floor(r.highPlot),k=Math.floor(r.lowPlot),q||(r.graphic=q=c.g("point").add(a.group),r.stem=c.path().addClass("highcharts-boxplot-stem").add(q),
B&&(r.whiskers=c.path().addClass("highcharts-boxplot-whisker").add(q)),E&&(r.box=c.path(void 0).addClass("highcharts-boxplot-box").add(q)),r.medianShape=c.path(void 0).addClass("highcharts-boxplot-median").add(q)),d.styledMode||(H.stroke=r.stemColor||e.stemColor||G,H["stroke-width"]=w(r.stemWidth,e.stemWidth,e.lineWidth),H.dashstyle=r.stemDashStyle||e.stemDashStyle,r.stem.attr(H),B&&(J.stroke=r.whiskerColor||e.whiskerColor||G,J["stroke-width"]=w(r.whiskerWidth,e.whiskerWidth,e.lineWidth),r.whiskers.attr(J)),
E&&(v.fill=r.fillColor||e.fillColor||G,v.stroke=e.lineColor||G,v["stroke-width"]=e.lineWidth||0,r.box.attr(v)),y.stroke=r.medianColor||e.medianColor||G,y["stroke-width"]=w(r.medianWidth,e.medianWidth,e.lineWidth),r.medianShape.attr(y)),f=r.stem.strokeWidth()%2/2,m=g+x+f,r.stem[u]({d:["M",m,b,"L",m,l,"M",m,n,"L",m,k]}),E&&(f=r.box.strokeWidth()%2/2,n=Math.floor(n)+f,b=Math.floor(b)+f,g+=f,p+=f,r.box[u]({d:["M",g,b,"L",g,n,"L",p,n,"L",p,b,"L",g,b,"z"]})),B&&(f=r.whiskers.strokeWidth()%2/2,l+=f,k+=f,
F=/%$/.test(B)?x*parseFloat(B)/100:B/2,r.whiskers[u]({d:["M",m-F,l,"L",m+F,l,"M",m-F,k,"L",m+F,k]})),h=Math.round(r.medianPlot),f=r.medianShape.strokeWidth()%2/2,h+=f,r.medianShape[u]({d:["M",g,h,"L",p,h]}))})},setStackedPoints:v})})(y);(function(a){var v=a.noop,w=a.seriesType,e=a.seriesTypes;w("errorbar","boxplot",{color:"#000000",grouping:!1,linkedTo:":previous",tooltip:{pointFormat:'\x3cspan style\x3d"color:{point.color}"\x3e\u25cf\x3c/span\x3e {series.name}: \x3cb\x3e{point.low}\x3c/b\x3e - \x3cb\x3e{point.high}\x3c/b\x3e\x3cbr/\x3e'},
whiskerWidth:null},{type:"errorbar",pointArrayMap:["low","high"],toYData:function(a){return[a.low,a.high]},pointValKey:"high",doQuartiles:!1,drawDataLabels:e.arearange?function(){var a=this.pointValKey;e.arearange.prototype.drawDataLabels.call(this);this.data.forEach(function(e){e.y=e[a]})}:v,getColumnMetrics:function(){return this.linkedParent&&this.linkedParent.columnMetrics||e.column.prototype.getColumnMetrics.call(this)}})})(y);(function(a){var v=a.correctFloat,w=a.isNumber,e=a.pick,m=a.Point,
u=a.Series,q=a.seriesType,d=a.seriesTypes;q("waterfall","column",{dataLabels:{inside:!0},lineWidth:1,lineColor:"#333333",dashStyle:"Dot",borderColor:"#333333",states:{hover:{lineWidthPlus:0}}},{pointValKey:"y",showLine:!0,generatePoints:function(){var c=this.options.threshold,a,b,l,k;d.column.prototype.generatePoints.apply(this);l=0;for(b=this.points.length;l<b;l++)a=this.points[l],k=this.processedYData[l],a.isSum?a.y=v(k):a.isIntermediateSum&&(a.y=v(k-c),c=k)},translate:function(){var c=this.options,
a=this.yAxis,b,l,k,h,f,r,t,g,p,m,q=e(c.minPointLength,5),u=q/2,B=c.threshold,z=c.stacking,v;d.column.prototype.translate.apply(this);g=p=B;l=this.points;b=0;for(c=l.length;b<c;b++)k=l[b],t=this.processedYData[b],h=k.shapeArgs,f=z&&a.stacks[(this.negStacks&&t<B?"-":"")+this.stackKey],v=this.getStackIndicator(v,k.x,this.index),m=e(f&&f[k.x].points[v.key],[0,t]),r=Math.max(g,g+k.y)+m[0],h.y=a.translate(r,0,1,0,1),k.isSum?(h.y=a.translate(m[1],0,1,0,1),h.height=Math.min(a.translate(m[0],0,1,0,1),a.len)-
h.y):k.isIntermediateSum?(h.y=a.translate(m[1],0,1,0,1),h.height=Math.min(a.translate(p,0,1,0,1),a.len)-h.y,p=m[1]):(h.height=0<t?a.translate(g,0,1,0,1)-h.y:a.translate(g,0,1,0,1)-a.translate(g-t,0,1,0,1),g+=f&&f[k.x]?f[k.x].total:t,k.below=g<e(B,0)),0>h.height&&(h.y+=h.height,h.height*=-1),k.plotY=h.y=Math.round(h.y)-this.borderWidth%2/2,h.height=Math.max(Math.round(h.height),.001),k.yBottom=h.y+h.height,h.height<=q&&!k.isNull?(h.height=q,h.y-=u,k.plotY=h.y,k.minPointLengthOffset=0>k.y?-u:u):(k.isNull&&
(h.width=0),k.minPointLengthOffset=0),h=k.plotY+(k.negative?h.height:0),this.chart.inverted?k.tooltipPos[0]=a.len-h:k.tooltipPos[1]=h},processData:function(a){var c=this.yData,b=this.options.data,d,k=c.length,h,f,r,e,g,p;f=h=r=e=this.options.threshold||0;for(p=0;p<k;p++)g=c[p],d=b&&b[p]?b[p]:{},"sum"===g||d.isSum?c[p]=v(f):"intermediateSum"===g||d.isIntermediateSum?c[p]=v(h):(f+=g,h+=g),r=Math.min(f,r),e=Math.max(f,e);u.prototype.processData.call(this,a);this.options.stacking||(this.dataMin=r,this.dataMax=
e)},toYData:function(a){return a.isSum?0===a.x?null:"sum":a.isIntermediateSum?0===a.x?null:"intermediateSum":a.y},pointAttribs:function(a,n){var b=this.options.upColor;b&&!a.options.color&&(a.color=0<a.y?b:null);a=d.column.prototype.pointAttribs.call(this,a,n);delete a.dashstyle;return a},getGraphPath:function(){return["M",0,0]},getCrispPath:function(){var a=this.data,d=a.length,b=this.graph.strokeWidth()+this.borderWidth,b=Math.round(b)%2/2,l=this.xAxis.reversed,k=this.yAxis.reversed,h=[],f,r,e;
for(e=1;e<d;e++){r=a[e].shapeArgs;f=a[e-1].shapeArgs;r=["M",f.x+(l?0:f.width),f.y+a[e-1].minPointLengthOffset+b,"L",r.x+(l?f.width:0),f.y+a[e-1].minPointLengthOffset+b];if(0>a[e-1].y&&!k||0<a[e-1].y&&k)r[2]+=f.height,r[5]+=f.height;h=h.concat(r)}return h},drawGraph:function(){u.prototype.drawGraph.call(this);this.graph.attr({d:this.getCrispPath()})},setStackedPoints:function(){var a=this.options,d,b;u.prototype.setStackedPoints.apply(this,arguments);d=this.stackedYData?this.stackedYData.length:0;
for(b=1;b<d;b++)a.data[b].isSum||a.data[b].isIntermediateSum||(this.stackedYData[b]+=this.stackedYData[b-1])},getExtremes:function(){if(this.options.stacking)return u.prototype.getExtremes.apply(this,arguments)}},{getClassName:function(){var a=m.prototype.getClassName.call(this);this.isSum?a+=" highcharts-sum":this.isIntermediateSum&&(a+=" highcharts-intermediate-sum");return a},isValid:function(){return w(this.y,!0)||this.isSum||this.isIntermediateSum}})})(y);(function(a){var v=a.Series,w=a.seriesType,
e=a.seriesTypes;w("polygon","scatter",{marker:{enabled:!1,states:{hover:{enabled:!1}}},stickyTracking:!1,tooltip:{followPointer:!0,pointFormat:""},trackByArea:!0},{type:"polygon",getGraphPath:function(){for(var a=v.prototype.getGraphPath.call(this),e=a.length+1;e--;)(e===a.length||"M"===a[e])&&0<e&&a.splice(e,0,"z");return this.areaPath=a},drawGraph:function(){this.options.fillColor=this.color;e.area.prototype.drawGraph.call(this)},drawLegendSymbol:a.LegendSymbolMixin.drawRectangle,drawTracker:v.prototype.drawTracker,
setStackedPoints:a.noop})})(y);(function(a){var v=a.Series,w=a.Legend,e=a.Chart,m=a.addEvent,u=a.wrap,q=a.color,d=a.isNumber,c=a.numberFormat,n=a.objectEach,b=a.merge,l=a.noop,k=a.pick,h=a.stableSort,f=a.setOptions,r=a.arrayMin,t=a.arrayMax;f({legend:{bubbleLegend:{borderColor:void 0,borderWidth:2,className:void 0,color:void 0,connectorClassName:void 0,connectorColor:void 0,connectorDistance:60,connectorWidth:1,enabled:!1,labels:{className:void 0,allowOverlap:!1,format:"",formatter:void 0,align:"right",
style:{fontSize:10,color:void 0},x:0,y:0},maxSize:60,minSize:10,legendIndex:0,ranges:{value:void 0,borderColor:void 0,color:void 0,connectorColor:void 0},sizeBy:"area",sizeByAbsoluteValue:!1,zIndex:1,zThreshold:0}}});a.BubbleLegend=function(a,b){this.init(a,b)};a.BubbleLegend.prototype={init:function(a,b){this.options=a;this.visible=!0;this.chart=b.chart;this.legend=b},setState:l,addToLegend:function(a){a.splice(this.options.legendIndex,0,this)},drawLegendSymbol:function(a){var b=this.chart,c=this.options,
g=k(a.options.itemDistance,20),f,e=c.ranges;f=c.connectorDistance;this.fontMetrics=b.renderer.fontMetrics(c.labels.style.fontSize.toString()+"px");e&&e.length&&d(e[0].value)?(h(e,function(a,b){return b.value-a.value}),this.ranges=e,this.setOptions(),this.render(),b=this.getMaxLabelSize(),e=this.ranges[0].radius,a=2*e,f=f-e+b.width,f=0<f?f:0,this.maxLabel=b,this.movementX="left"===c.labels.align?f:0,this.legendItemWidth=a+f+g,this.legendItemHeight=a+this.fontMetrics.h/2):a.options.bubbleLegend.autoRanges=
!0},setOptions:function(){var a=this,c=a.ranges,d=a.options,h=a.chart.series[d.seriesIndex],f=a.legend.baseline,e={"z-index":d.zIndex,"stroke-width":d.borderWidth},l={"z-index":d.zIndex,"stroke-width":d.connectorWidth},n=a.getLabelStyles(),r=h.options.marker.fillOpacity,m=a.chart.styledMode;c.forEach(function(g,p){m||(e.stroke=k(g.borderColor,d.borderColor,h.color),e.fill=k(g.color,d.color,1!==r?q(h.color).setOpacity(r).get("rgba"):h.color),l.stroke=k(g.connectorColor,d.connectorColor,h.color));c[p].radius=
a.getRangeRadius(g.value);c[p]=b(c[p],{center:c[0].radius-c[p].radius+f});m||b(!0,c[p],{bubbleStyle:b(!1,e),connectorStyle:b(!1,l),labelStyle:n})})},getLabelStyles:function(){var a=this.options,c={},d="left"===a.labels.align,h=this.legend.options.rtl;n(a.labels.style,function(a,b){"color"!==b&&"fontSize"!==b&&"z-index"!==b&&(c[b]=a)});return b(!1,c,{"font-size":a.labels.style.fontSize,fill:k(a.labels.style.color,"#000000"),"z-index":a.zIndex,align:h||d?"right":"left"})},getRangeRadius:function(a){var b=
this.options;return this.chart.series[this.options.seriesIndex].getRadius.call(this,b.ranges[b.ranges.length-1].value,b.ranges[0].value,b.minSize,b.maxSize,a)},render:function(){var a=this,b=a.chart.renderer,c=a.options.zThreshold;a.symbols||(a.symbols={connectors:[],bubbleItems:[],labels:[]});a.legendSymbol=b.g("bubble-legend");a.legendItem=b.g("bubble-legend-item");a.legendSymbol.translateX=0;a.legendSymbol.translateY=0;a.ranges.forEach(function(b){b.value>=c&&a.renderRange(b)});a.legendSymbol.add(a.legendItem);
a.legendItem.add(a.legendGroup);a.hideOverlappingLabels()},renderRange:function(a){var b=this.options,c=b.labels,d=this.chart.renderer,g=this.symbols,h=g.labels,f=a.center,e=Math.abs(a.radius),k=b.connectorDistance,l=c.align,r=c.style.fontSize,k=this.legend.options.rtl||"left"===l?-k:k,c=b.connectorWidth,n=this.ranges[0].radius,m=f-e-b.borderWidth/2+c/2,t,r=r/2-(this.fontMetrics.h-r)/2,q=d.styledMode;"center"===l&&(k=0,b.connectorDistance=0,a.labelStyle.align="center");l=m+b.labels.y;t=n+k+b.labels.x;
g.bubbleItems.push(d.circle(n,f+((m%1?1:.5)-(c%2?0:.5)),e).attr(q?{}:a.bubbleStyle).addClass((q?"highcharts-color-"+this.options.seriesIndex+" ":"")+"highcharts-bubble-legend-symbol "+(b.className||"")).add(this.legendSymbol));g.connectors.push(d.path(d.crispLine(["M",n,m,"L",n+k,m],b.connectorWidth)).attr(q?{}:a.connectorStyle).addClass((q?"highcharts-color-"+this.options.seriesIndex+" ":"")+"highcharts-bubble-legend-connectors "+(b.connectorClassName||"")).add(this.legendSymbol));a=d.text(this.formatLabel(a),
t,l+r).attr(q?{}:a.labelStyle).addClass("highcharts-bubble-legend-labels "+(b.labels.className||"")).add(this.legendSymbol);h.push(a);a.placed=!0;a.alignAttr={x:t,y:l+r}},getMaxLabelSize:function(){var a,b;this.symbols.labels.forEach(function(c){b=c.getBBox(!0);a=a?b.width>a.width?b:a:b});return a},formatLabel:function(b){var d=this.options,g=d.labels.formatter;return(d=d.labels.format)?a.format(d,b):g?g.call(b):c(b.value,1)},hideOverlappingLabels:function(){var a=this.chart,b=this.symbols;!this.options.labels.allowOverlap&&
b&&(a.hideOverlappingLabels(b.labels),b.labels.forEach(function(a,c){a.newOpacity?a.newOpacity!==a.oldOpacity&&b.connectors[c].show():b.connectors[c].hide()}))},getRanges:function(){var a=this.legend.bubbleLegend,c,h=a.options.ranges,f,e=Number.MAX_VALUE,l=-Number.MAX_VALUE;a.chart.series.forEach(function(a){a.isBubble&&!a.ignoreSeries&&(f=a.zData.filter(d),f.length&&(e=k(a.options.zMin,Math.min(e,Math.max(r(f),!1===a.options.displayNegative?a.options.zThreshold:-Number.MAX_VALUE))),l=k(a.options.zMax,
Math.max(l,t(f)))))});c=e===l?[{value:l}]:[{value:e},{value:(e+l)/2},{value:l,autoRanges:!0}];h.length&&h[0].radius&&c.reverse();c.forEach(function(a,d){h&&h[d]&&(c[d]=b(!1,h[d],a))});return c},predictBubbleSizes:function(){var a=this.chart,b=this.fontMetrics,c=a.legend.options,d="horizontal"===c.layout,h=d?a.legend.lastLineHeight:0,f=a.plotSizeX,e=a.plotSizeY,l=a.series[this.options.seriesIndex],a=Math.ceil(l.minPxSize),k=Math.ceil(l.maxPxSize),l=l.options.maxSize,r=Math.min(e,f);if(c.floating||
!/%$/.test(l))b=k;else if(l=parseFloat(l),b=(r+h-b.h/2)*l/100/(l/100+1),d&&e-b>=f||!d&&f-b>=e)b=k;return[a,Math.ceil(b)]},updateRanges:function(a,b){var c=this.legend.options.bubbleLegend;c.minSize=a;c.maxSize=b;c.ranges=this.getRanges()},correctSizes:function(){var a=this.legend,b=this.chart.series[this.options.seriesIndex];1<Math.abs(Math.ceil(b.maxPxSize)-this.options.maxSize)&&(this.updateRanges(this.options.minSize,b.maxPxSize),a.render())}};m(a.Legend,"afterGetAllItems",function(b){var c=this.bubbleLegend,
d=this.options,g=d.bubbleLegend,h=this.chart.getVisibleBubbleSeriesIndex();c&&c.ranges&&c.ranges.length&&(g.ranges.length&&(g.autoRanges=g.ranges[0].autoRanges?!0:!1),this.destroyItem(c));0<=h&&d.enabled&&g.enabled&&(g.seriesIndex=h,this.bubbleLegend=new a.BubbleLegend(g,this),this.bubbleLegend.addToLegend(b.allItems))});e.prototype.getVisibleBubbleSeriesIndex=function(){for(var a=this.series,b=0;b<a.length;){if(a[b]&&a[b].isBubble&&a[b].visible&&a[b].zData.length)return b;b++}return-1};w.prototype.getLinesHeights=
function(){var a=this.allItems,b=[],c,d=a.length,h,f=0;for(h=0;h<d;h++)if(a[h].legendItemHeight&&(a[h].itemHeight=a[h].legendItemHeight),a[h]===a[d-1]||a[h+1]&&a[h]._legendItemPos[1]!==a[h+1]._legendItemPos[1]){b.push({height:0});c=b[b.length-1];for(f;f<=h;f++)a[f].itemHeight>c.height&&(c.height=a[f].itemHeight);c.step=h}return b};w.prototype.retranslateItems=function(a){var b,c,d,h=this.options.rtl,f=0;this.allItems.forEach(function(g,e){b=g.legendGroup.translateX;c=g._legendItemPos[1];if((d=g.movementX)||
h&&g.ranges)d=h?b-g.options.maxSize/2:b+d,g.legendGroup.attr({translateX:d});e>a[f].step&&f++;g.legendGroup.attr({translateY:Math.round(c+a[f].height/2)});g._legendItemPos[1]=c+a[f].height/2})};m(v,"legendItemClick",function(){var a=this.chart,b=this.visible,c=this.chart.legend;c&&c.bubbleLegend&&(this.visible=!b,this.ignoreSeries=b,a=0<=a.getVisibleBubbleSeriesIndex(),c.bubbleLegend.visible!==a&&(c.update({bubbleLegend:{enabled:a}}),c.bubbleLegend.visible=a),this.visible=b)});u(e.prototype,"drawChartBox",
function(a,b,c){var d=this.legend,h=0<=this.getVisibleBubbleSeriesIndex(),f;d&&d.options.enabled&&d.bubbleLegend&&d.options.bubbleLegend.autoRanges&&h?(f=d.bubbleLegend.options,h=d.bubbleLegend.predictBubbleSizes(),d.bubbleLegend.updateRanges(h[0],h[1]),f.placed||(d.group.placed=!1,d.allItems.forEach(function(a){a.legendGroup.translateY=null})),d.render(),this.getMargins(),this.axes.forEach(function(a){a.render();f.placed||(a.setScale(),a.updateNames(),n(a.ticks,function(a){a.isNew=!0;a.isNewLabel=
!0}))}),f.placed=!0,this.getMargins(),a.call(this,b,c),d.bubbleLegend.correctSizes(),d.retranslateItems(d.getLinesHeights())):(a.call(this,b,c),d&&d.options.enabled&&d.bubbleLegend&&(d.render(),d.retranslateItems(d.getLinesHeights())))})})(y);(function(a){var v=a.arrayMax,w=a.arrayMin,e=a.Axis,m=a.color,u=a.isNumber,q=a.noop,d=a.pick,c=a.pInt,n=a.Point,b=a.Series,l=a.seriesType,k=a.seriesTypes;l("bubble","scatter",{dataLabels:{formatter:function(){return this.point.z},inside:!0,verticalAlign:"middle"},
animationLimit:250,marker:{lineColor:null,lineWidth:1,fillOpacity:.5,radius:null,states:{hover:{radiusPlus:0}},symbol:"circle"},minSize:8,maxSize:"20%",softThreshold:!1,states:{hover:{halo:{size:5}}},tooltip:{pointFormat:"({point.x}, {point.y}), Size: {point.z}"},turboThreshold:0,zThreshold:0,zoneAxis:"z"},{pointArrayMap:["y","z"],parallelArrays:["x","y","z"],trackerGroups:["group","dataLabelsGroup"],specialGroup:"group",bubblePadding:!0,zoneAxis:"z",directTouch:!0,isBubble:!0,pointAttribs:function(a,
c){var d=this.options.marker.fillOpacity;a=b.prototype.pointAttribs.call(this,a,c);1!==d&&(a.fill=m(a.fill).setOpacity(d).get("rgba"));return a},getRadii:function(a,b,c){var d,h=this.zData,f=c.minPxSize,e=c.maxPxSize,l=[],k;d=0;for(c=h.length;d<c;d++)k=h[d],l.push(this.getRadius(a,b,f,e,k));this.radii=l},getRadius:function(a,b,c,d,e){var h=this.options,f="width"!==h.sizeBy,g=h.zThreshold,l=b-a;h.sizeByAbsoluteValue&&null!==e&&(e=Math.abs(e-g),l=Math.max(b-g,Math.abs(a-g)),a=0);u(e)?e<a?c=c/2-1:(a=
0<l?(e-a)/l:.5,f&&0<=a&&(a=Math.sqrt(a)),c=Math.ceil(c+a*(d-c))/2):c=null;return c},animate:function(a){!a&&this.points.length<this.options.animationLimit&&(this.points.forEach(function(a){var b=a.graphic,c;b&&b.width&&(c={x:b.x,y:b.y,width:b.width,height:b.height},b.attr({x:a.plotX,y:a.plotY,width:1,height:1}),b.animate(c,this.options.animation))},this),this.animate=null)},translate:function(){var b,c=this.data,d,e,g=this.radii;k.scatter.prototype.translate.call(this);for(b=c.length;b--;)d=c[b],
e=g?g[b]:0,u(e)&&e>=this.minPxSize/2?(d.marker=a.extend(d.marker,{radius:e,width:2*e,height:2*e}),d.dlBox={x:d.plotX-e,y:d.plotY-e,width:2*e,height:2*e}):d.shapeArgs=d.plotY=d.dlBox=void 0},alignDataLabel:k.column.prototype.alignDataLabel,buildKDTree:q,applyZones:q},{haloPath:function(a){return n.prototype.haloPath.call(this,0===a?0:(this.marker?this.marker.radius||0:0)+a)},ttBelow:!1});e.prototype.beforePadding=function(){var b=this,e=this.len,l=this.chart,k=0,g=e,n=this.isXAxis,m=n?"xData":"yData",
q=this.min,y={},B=Math.min(l.plotWidth,l.plotHeight),z=Number.MAX_VALUE,A=-Number.MAX_VALUE,I=this.max-q,C=e/I,D=[];this.series.forEach(function(e){var f=e.options;!e.bubblePadding||!e.visible&&l.options.chart.ignoreHiddenSeries||(b.allowZoomOutside=!0,D.push(e),n&&(["minSize","maxSize"].forEach(function(a){var b=f[a],d=/%$/.test(b),b=c(b);y[a]=d?B*b/100:b}),e.minPxSize=y.minSize,e.maxPxSize=Math.max(y.maxSize,y.minSize),e=e.zData.filter(a.isNumber),e.length&&(z=d(f.zMin,Math.min(z,Math.max(w(e),
!1===f.displayNegative?f.zThreshold:-Number.MAX_VALUE))),A=d(f.zMax,Math.max(A,v(e))))))});D.forEach(function(a){var c=a[m],d=c.length,e;n&&a.getRadii(z,A,a);if(0<I)for(;d--;)u(c[d])&&b.dataMin<=c[d]&&c[d]<=b.dataMax&&(e=a.radii[d],k=Math.min((c[d]-q)*C-e,k),g=Math.max((c[d]-q)*C+e,g))});D.length&&0<I&&!this.isLog&&(g-=e,C*=(e+Math.max(0,k)-Math.min(g,e))/e,[["min","userMin",k],["max","userMax",g]].forEach(function(a){void 0===d(b.options[a[0]],b[a[1]])&&(b[a[0]]+=a[2]/C)}))}})(y);(function(a){var v=
a.seriesType,w=a.defined;v("packedbubble","bubble",{minSize:"10%",maxSize:"100%",sizeBy:"radius",zoneAxis:"y",tooltip:{pointFormat:"Value: {point.value}"}},{pointArrayMap:["value"],pointValKey:"value",isCartesian:!1,axisTypes:[],accumulateAllPoints:function(a){var e=a.chart,u=[],q,d;for(q=0;q<e.series.length;q++)if(a=e.series[q],a.visible||!e.options.chart.ignoreHiddenSeries)for(d=0;d<a.yData.length;d++)u.push([null,null,a.yData[d],a.index,d]);return u},translate:function(){var e,m=this.chart,u=this.data,
q=this.index,d,c,n;this.processedXData=this.xData;this.generatePoints();w(m.allDataPoints)||(m.allDataPoints=this.accumulateAllPoints(this),this.getPointRadius());e=this.placeBubbles(m.allDataPoints);for(n=0;n<e.length;n++)e[n][3]===q&&(d=u[e[n][4]],c=e[n][2],d.plotX=e[n][0]-m.plotLeft+m.diffX,d.plotY=e[n][1]-m.plotTop+m.diffY,d.marker=a.extend(d.marker,{radius:c,width:2*c,height:2*c}))},checkOverlap:function(a,m){var e=a[0]-m[0],q=a[1]-m[1];return-.001>Math.sqrt(e*e+q*q)-Math.abs(a[2]+m[2])},positionBubble:function(a,
m,u){var e=Math.sqrt,d=Math.asin,c=Math.acos,n=Math.pow,b=Math.abs,e=e(n(a[0]-m[0],2)+n(a[1]-m[1],2)),c=c((n(e,2)+n(u[2]+m[2],2)-n(u[2]+a[2],2))/(2*(u[2]+m[2])*e)),d=d(b(a[0]-m[0])/e);a=(0>a[1]-m[1]?0:Math.PI)+c+d*(0>(a[0]-m[0])*(a[1]-m[1])?1:-1);return[m[0]+(m[2]+u[2])*Math.sin(a),m[1]-(m[2]+u[2])*Math.cos(a),u[2],u[3],u[4]]},placeBubbles:function(a){var e=this.checkOverlap,u=this.positionBubble,q=[],d=1,c=0,n=0,b,l;b=a.sort(function(a,b){return b[2]-a[2]});if(!b.length)return[];if(2>b.length)return[0,
0,b[0][0],b[0][1],b[0][2]];q.push([[0,0,b[0][2],b[0][3],b[0][4]]]);q.push([[0,0-b[1][2]-b[0][2],b[1][2],b[1][3],b[1][4]]]);for(l=2;l<b.length;l++)b[l][2]=b[l][2]||1,a=u(q[d][c],q[d-1][n],b[l]),e(a,q[d][0])?(q.push([]),n=0,q[d+1].push(u(q[d][c],q[d][0],b[l])),d++,c=0):1<d&&q[d-1][n+1]&&e(a,q[d-1][n+1])?(n++,q[d].push(u(q[d][c],q[d-1][n],b[l])),c++):(c++,q[d].push(a));this.chart.stages=q;this.chart.rawPositions=[].concat.apply([],q);this.resizeRadius();return this.chart.rawPositions},resizeRadius:function(){var a=
this.chart,m=a.rawPositions,u=Math.min,q=Math.max,d=a.plotLeft,c=a.plotTop,n=a.plotHeight,b=a.plotWidth,l,k,h,f,r,t;l=h=Number.POSITIVE_INFINITY;k=f=Number.NEGATIVE_INFINITY;for(t=0;t<m.length;t++)r=m[t][2],l=u(l,m[t][0]-r),k=q(k,m[t][0]+r),h=u(h,m[t][1]-r),f=q(f,m[t][1]+r);t=[k-l,f-h];u=u.apply([],[(b-d)/t[0],(n-c)/t[1]]);if(1e-10<Math.abs(u-1)){for(t=0;t<m.length;t++)m[t][2]*=u;this.placeBubbles(m)}else a.diffY=n/2+c-h-(f-h)/2,a.diffX=b/2+d-l-(k-l)/2},getPointRadius:function(){var a=this,m=a.chart,
u=a.options,q=Math.min(m.plotWidth,m.plotHeight),d={},c=[],n=m.allDataPoints,b,l,k,h;["minSize","maxSize"].forEach(function(a){var b=parseInt(u[a],10),c=/%$/.test(b);d[a]=c?q*b/100:b});m.minRadius=b=d.minSize;m.maxRadius=l=d.maxSize;(n||[]).forEach(function(d,e){k=d[2];h=a.getRadius(b,l,b,l,k);0===k&&(h=null);n[e][2]=h;c.push(h)});this.radii=c},alignDataLabel:a.Series.prototype.alignDataLabel});a.addEvent(a.seriesTypes.packedbubble,"updatedData",function(){var a=this;this.chart.series.forEach(function(e){e.type===
a.type&&(e.isDirty=!0)})});a.addEvent(a.Chart,"beforeRedraw",function(){this.allDataPoints&&delete this.allDataPoints})})(y);(function(a){var v=a.pick,w=a.Series,e=a.seriesTypes,m=a.wrap,u=w.prototype,q=a.Pointer.prototype;a.polarExtended||(a.polarExtended=!0,u.searchPointByAngle=function(a){var c=this.chart,d=this.xAxis.pane.center;return this.searchKDTree({clientX:180+-180/Math.PI*Math.atan2(a.chartX-d[0]-c.plotLeft,a.chartY-d[1]-c.plotTop)})},u.getConnectors=function(a,c,e,b){var d,k,h,f,n,m,g,
p;k=b?1:0;d=0<=c&&c<=a.length-1?c:0>c?a.length-1+c:0;c=0>d-1?a.length-(1+k):d-1;k=d+1>a.length-1?k:d+1;h=a[c];k=a[k];f=h.plotX;h=h.plotY;n=k.plotX;m=k.plotY;k=a[d].plotX;d=a[d].plotY;f=(1.5*k+f)/2.5;h=(1.5*d+h)/2.5;n=(1.5*k+n)/2.5;g=(1.5*d+m)/2.5;m=Math.sqrt(Math.pow(f-k,2)+Math.pow(h-d,2));p=Math.sqrt(Math.pow(n-k,2)+Math.pow(g-d,2));f=Math.atan2(h-d,f-k);g=Math.PI/2+(f+Math.atan2(g-d,n-k))/2;Math.abs(f-g)>Math.PI/2&&(g-=Math.PI);f=k+Math.cos(g)*m;h=d+Math.sin(g)*m;n=k+Math.cos(Math.PI+g)*p;g=d+
Math.sin(Math.PI+g)*p;k={rightContX:n,rightContY:g,leftContX:f,leftContY:h,plotX:k,plotY:d};e&&(k.prevPointCont=this.getConnectors(a,c,!1,b));return k},m(u,"buildKDTree",function(a){this.chart.polar&&(this.kdByAngle?this.searchPoint=this.searchPointByAngle:this.options.findNearestPointBy="xy");a.apply(this)}),u.toXY=function(a){var c,d=this.chart,b=a.plotX;c=a.plotY;a.rectPlotX=b;a.rectPlotY=c;c=this.xAxis.postTranslate(a.plotX,this.yAxis.len-c);a.plotX=a.polarPlotX=c.x-d.plotLeft;a.plotY=a.polarPlotY=
c.y-d.plotTop;this.kdByAngle?(d=(b/Math.PI*180+this.xAxis.pane.options.startAngle)%360,0>d&&(d+=360),a.clientX=d):a.clientX=a.plotX},e.spline&&(m(e.spline.prototype,"getPointSpline",function(a,c,e,b){this.chart.polar?b?(a=this.getConnectors(c,b,!0,this.connectEnds),a=["C",a.prevPointCont.rightContX,a.prevPointCont.rightContY,a.leftContX,a.leftContY,a.plotX,a.plotY]):a=["M",e.plotX,e.plotY]:a=a.call(this,c,e,b);return a}),e.areasplinerange&&(e.areasplinerange.prototype.getPointSpline=e.spline.prototype.getPointSpline)),
a.addEvent(w,"afterTranslate",function(){var d=this.chart,c,e;if(d.polar){this.kdByAngle=d.tooltip&&d.tooltip.shared;if(!this.preventPostTranslate)for(c=this.points,e=c.length;e--;)this.toXY(c[e]);this.hasClipCircleSetter||(this.hasClipCircleSetter=!!a.addEvent(this,"afterRender",function(){var b;d.polar&&(b=this.yAxis.center,this.group.clip(d.renderer.clipCircle(b[0],b[1],b[2]/2)),this.setClip=a.noop)}))}},{order:2}),m(u,"getGraphPath",function(a,c){var d=this,b,e,k;if(this.chart.polar){c=c||this.points;
for(b=0;b<c.length;b++)if(!c[b].isNull){e=b;break}!1!==this.options.connectEnds&&void 0!==e&&(this.connectEnds=!0,c.splice(c.length,0,c[e]),k=!0);c.forEach(function(a){void 0===a.polarPlotY&&d.toXY(a)})}b=a.apply(this,[].slice.call(arguments,1));k&&c.pop();return b}),w=function(a,c){var d=this.chart,b=this.options.animation,e=this.group,k=this.markerGroup,h=this.xAxis.center,f=d.plotLeft,m=d.plotTop;d.polar?d.renderer.isSVG&&(!0===b&&(b={}),c?(a={translateX:h[0]+f,translateY:h[1]+m,scaleX:.001,scaleY:.001},
e.attr(a),k&&k.attr(a)):(a={translateX:f,translateY:m,scaleX:1,scaleY:1},e.animate(a,b),k&&k.animate(a,b),this.animate=null)):a.call(this,c)},m(u,"animate",w),e.column&&(e=e.column.prototype,e.polarArc=function(a,c,e,b){var d=this.xAxis.center,k=this.yAxis.len;return this.chart.renderer.symbols.arc(d[0],d[1],k-c,null,{start:e,end:b,innerR:k-v(a,k)})},m(e,"animate",w),m(e,"translate",function(a){var c=this.xAxis,d=c.startAngleRad,b,e,k;this.preventPostTranslate=!0;a.call(this);if(c.isRadial)for(b=
this.points,k=b.length;k--;)e=b[k],a=e.barX+d,e.shapeType="path",e.shapeArgs={d:this.polarArc(e.yBottom,e.plotY,a,a+e.pointWidth)},this.toXY(e),e.tooltipPos=[e.plotX,e.plotY],e.ttBelow=e.plotY>c.center[1]}),m(e,"alignDataLabel",function(a,c,e,b,l,k){this.chart.polar?(a=c.rectPlotX/Math.PI*180,null===b.align&&(b.align=20<a&&160>a?"left":200<a&&340>a?"right":"center"),null===b.verticalAlign&&(b.verticalAlign=45>a||315<a?"bottom":135<a&&225>a?"top":"middle"),u.alignDataLabel.call(this,c,e,b,l,k)):a.call(this,
c,e,b,l,k)})),m(q,"getCoordinates",function(a,c){var d=this.chart,b={xAxis:[],yAxis:[]};d.polar?d.axes.forEach(function(a){var e=a.isXAxis,h=a.center,f=c.chartX-h[0]-d.plotLeft,h=c.chartY-h[1]-d.plotTop;b[e?"xAxis":"yAxis"].push({axis:a,value:a.translate(e?Math.PI-Math.atan2(f,h):Math.sqrt(Math.pow(f,2)+Math.pow(h,2)),!0)})}):b=a.call(this,c);return b}),a.SVGRenderer.prototype.clipCircle=function(d,c,e){var b=a.uniqueKey(),l=this.createElement("clipPath").attr({id:b}).add(this.defs);d=this.circle(d,
c,e).add(l);d.id=b;d.clipPath=l;return d},a.addEvent(a.Chart,"getAxes",function(){this.pane||(this.pane=[]);a.splat(this.options.pane).forEach(function(d){new a.Pane(d,this)},this)}),a.addEvent(a.Chart,"afterDrawChartBox",function(){this.pane.forEach(function(a){a.render()})}),m(a.Chart.prototype,"get",function(d,c){return a.find(this.pane,function(a){return a.options.id===c})||d.call(this,c)}))})(y)});
//# sourceMappingURL=highcharts-more.js.map

`

var highchartsCSS string = `
/**
 * @license Highcharts
 *
 * (c) 2009-2016 Torstein Honsi
 *
 * License: www.highcharts.com/license
 */
.highcharts-container {
  position: relative;
  overflow: hidden;
  width: 100%;
  height: 100%;
  text-align: left;
  line-height: normal;
  z-index: 0;
  /* #1072 */
  -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
  font-family: "Lucida Grande", "Lucida Sans Unicode", Arial, Helvetica, sans-serif;
  font-size: 12px;
}

.highcharts-root {
  display: block;
}

.highcharts-root text {
  stroke-width: 0;
}

.highcharts-strong {
  font-weight: bold;
}

.highcharts-emphasized {
  font-style: italic;
}

.highcharts-anchor {
  cursor: pointer;
}

.highcharts-background {
  fill: #ffffff;
}

.highcharts-plot-border, .highcharts-plot-background {
  fill: none;
}

.highcharts-label-box {
  fill: none;
}

.highcharts-button-box {
  fill: inherit;
}

.highcharts-tracker-line {
  stroke-linejoin: round;
  stroke: rgba(192, 192, 192, 0.0001);
  stroke-width: 22;
  fill: none;
}

.highcharts-tracker-area {
  fill: rgba(192, 192, 192, 0.0001);
  stroke-width: 0;
}

/* Titles */
.highcharts-title {
  fill: #333333;
  font-size: 1.5em;
}

.highcharts-subtitle {
  fill: #666666;
}

/* Axes */
.highcharts-axis-line {
  fill: none;
  stroke: #ccd6eb;
}

.highcharts-yaxis .highcharts-axis-line {
  stroke-width: 0;
}

.highcharts-axis-title {
  fill: #666666;
}

.highcharts-axis-labels {
  fill: #666666;
  cursor: default;
  font-size: 0.9em;
}

.highcharts-grid-line {
  fill: none;
  stroke: #e6e6e6;
}

.highcharts-xaxis-grid .highcharts-grid-line {
  stroke-width: 0px;
}

.highcharts-tick {
  stroke: #ccd6eb;
}

.highcharts-yaxis .highcharts-tick {
  stroke-width: 0;
}

.highcharts-minor-grid-line {
  stroke: #f2f2f2;
}

.highcharts-crosshair-thin {
  stroke-width: 1px;
  stroke: #cccccc;
}

.highcharts-crosshair-category {
  stroke: #ccd6eb;
  stroke-opacity: 0.25;
}

/* Credits */
.highcharts-credits {
  cursor: pointer;
  fill: #999999;
  font-size: 0.7em;
  transition: fill 250ms, font-size 250ms;
}

.highcharts-credits:hover {
  fill: black;
  font-size: 1em;
}

/* Tooltip */
.highcharts-tooltip {
  cursor: default;
  pointer-events: none;
  white-space: nowrap;
  transition: stroke 150ms;
}

.highcharts-tooltip text {
  fill: #333333;
}

.highcharts-tooltip .highcharts-header {
  font-size: 0.85em;
}

.highcharts-tooltip-box {
  stroke-width: 1px;
  fill: #f7f7f7;
  fill-opacity: 0.85;
}

.highcharts-tooltip-box .highcharts-label-box {
  fill: #f7f7f7;
  fill-opacity: 0.85;
}

.highcharts-selection-marker {
  fill: #335cad;
  fill-opacity: 0.25;
}

.highcharts-graph {
  fill: none;
  stroke-width: 2px;
  stroke-linecap: round;
  stroke-linejoin: round;
}

.highcharts-state-hover .highcharts-graph {
  stroke-width: 3;
}

.highcharts-state-hover path {
  transition: stroke-width 50;
  /* quick in */
}

.highcharts-state-normal path {
  transition: stroke-width 250ms;
  /* slow out */
}

/* Legend hover affects points and series */
g.highcharts-series,
.highcharts-point,
.highcharts-markers,
.highcharts-data-labels {
  transition: opacity 250ms;
}

.highcharts-legend-series-active g.highcharts-series:not(.highcharts-series-hover),
.highcharts-legend-point-active .highcharts-point:not(.highcharts-point-hover),
.highcharts-legend-series-active .highcharts-markers:not(.highcharts-series-hover),
.highcharts-legend-series-active .highcharts-data-labels:not(.highcharts-series-hover) {
  opacity: 0.2;
}

/* Series options */
/* Default colors */
.highcharts-color-0 {
  fill: #7cb5ec;
  stroke: #7cb5ec;
}

.highcharts-color-1 {
  fill: #434348;
  stroke: #434348;
}

.highcharts-color-2 {
  fill: #90ed7d;
  stroke: #90ed7d;
}

.highcharts-color-3 {
  fill: #f7a35c;
  stroke: #f7a35c;
}

.highcharts-color-4 {
  fill: #8085e9;
  stroke: #8085e9;
}

.highcharts-color-5 {
  fill: #f15c80;
  stroke: #f15c80;
}

.highcharts-color-6 {
  fill: #e4d354;
  stroke: #e4d354;
}

.highcharts-color-7 {
  fill: #2b908f;
  stroke: #2b908f;
}

.highcharts-color-8 {
  fill: #f45b5b;
  stroke: #f45b5b;
}

.highcharts-color-9 {
  fill: #91e8e1;
  stroke: #91e8e1;
}

.highcharts-area {
  fill-opacity: 0.75;
  stroke-width: 0;
}

.highcharts-markers {
  stroke-width: 1px;
  stroke: #ffffff;
}

.highcharts-point {
  stroke-width: 1px;
}

.highcharts-dense-data .highcharts-point {
  stroke-width: 0;
}

.highcharts-data-label {
  font-size: 0.9em;
  font-weight: bold;
}

.highcharts-data-label-box {
  fill: none;
  stroke-width: 0;
}

.highcharts-data-label text, text.highcharts-data-label {
  fill: #333333;
}

.highcharts-data-label-connector {
  fill: none;
}

.highcharts-halo {
  fill-opacity: 0.25;
  stroke-width: 0;
}

.highcharts-series:not(.highcharts-pie-series) .highcharts-point-select,
.highcharts-markers .highcharts-point-select {
  fill: #cccccc;
  stroke: #000000;
}

.highcharts-column-series rect.highcharts-point {
  stroke: #ffffff;
}

.highcharts-column-series .highcharts-point {
  transition: fill-opacity 250ms;
}

.highcharts-column-series .highcharts-point-hover {
  fill-opacity: 0.75;
  transition: fill-opacity 50ms;
}

.highcharts-pie-series .highcharts-point {
  stroke-linejoin: round;
  stroke: #ffffff;
}

.highcharts-pie-series .highcharts-point-hover {
  fill-opacity: 0.75;
  transition: fill-opacity 50ms;
}

.highcharts-funnel-series .highcharts-point {
  stroke-linejoin: round;
  stroke: #ffffff;
}

.highcharts-funnel-series .highcharts-point-hover {
  fill-opacity: 0.75;
  transition: fill-opacity 50ms;
}

.highcharts-funnel-series .highcharts-point-select {
  fill: inherit;
  stroke: inherit;
}

.highcharts-pyramid-series .highcharts-point {
  stroke-linejoin: round;
  stroke: #ffffff;
}

.highcharts-pyramid-series .highcharts-point-hover {
  fill-opacity: 0.75;
  transition: fill-opacity 50ms;
}

.highcharts-pyramid-series .highcharts-point-select {
  fill: inherit;
  stroke: inherit;
}

.highcharts-solidgauge-series .highcharts-point {
  stroke-width: 0;
}

.highcharts-treemap-series .highcharts-point {
  stroke-width: 1px;
  stroke: #e6e6e6;
  transition: stroke 250ms, fill 250ms, fill-opacity 250ms;
}

.highcharts-treemap-series .highcharts-point-hover {
  stroke: #999999;
  transition: stroke 25ms, fill 25ms, fill-opacity 25ms;
}

.highcharts-treemap-series .highcharts-above-level {
  display: none;
}

.highcharts-treemap-series .highcharts-internal-node {
  fill: none;
}

.highcharts-treemap-series .highcharts-internal-node-interactive {
  fill-opacity: 0.15;
  cursor: pointer;
}

.highcharts-treemap-series .highcharts-internal-node-interactive:hover {
  fill-opacity: 0.75;
}

/* Legend */
.highcharts-legend-box {
  fill: none;
  stroke-width: 0;
}

.highcharts-legend-item > text {
  fill: #333333;
  font-weight: bold;
  font-size: 1em;
  cursor: pointer;
  stroke-width: 0;
}

.highcharts-legend-item:hover text {
  fill: #000000;
}

.highcharts-legend-item-hidden * {
  fill: #cccccc !important;
  stroke: #cccccc !important;
  transition: fill 250ms;
}

.highcharts-legend-nav-active {
  fill: #003399;
  cursor: pointer;
}

.highcharts-legend-nav-inactive {
  fill: #cccccc;
}

.highcharts-legend-title-box {
  fill: none;
  stroke-width: 0;
}

/* Bubble legend */
.highcharts-bubble-legend-symbol {
  stroke-width: 2;
  fill-opacity: 0.5;
}

.highcharts-bubble-legend-connectors {
  stroke-width: 1;
}

.highcharts-bubble-legend-labels {
  fill: #333333;
}

/* Loading */
.highcharts-loading {
  position: absolute;
  background-color: #ffffff;
  opacity: 0.5;
  text-align: center;
  z-index: 10;
  transition: opacity 250ms;
}

.highcharts-loading-hidden {
  height: 0 !important;
  opacity: 0;
  overflow: hidden;
  transition: opacity 250ms, height 250ms step-end;
}

.highcharts-loading-inner {
  font-weight: bold;
  position: relative;
  top: 45%;
}

/* Plot bands and polar pane backgrounds */
.highcharts-plot-band, .highcharts-pane {
  fill: #000000;
  fill-opacity: 0.05;
}

.highcharts-plot-line {
  fill: none;
  stroke: #999999;
  stroke-width: 1px;
}

/* Highcharts More and modules */
.highcharts-boxplot-box {
  fill: #ffffff;
}

.highcharts-boxplot-median {
  stroke-width: 2px;
}

.highcharts-bubble-series .highcharts-point {
  fill-opacity: 0.5;
}

.highcharts-errorbar-series .highcharts-point {
  stroke: #000000;
}

.highcharts-gauge-series .highcharts-data-label-box {
  stroke: #cccccc;
  stroke-width: 1px;
}

.highcharts-gauge-series .highcharts-dial {
  fill: #000000;
  stroke-width: 0;
}

.highcharts-polygon-series .highcharts-graph {
  fill: inherit;
  stroke-width: 0;
}

.highcharts-waterfall-series .highcharts-graph {
  stroke: #333333;
  stroke-dasharray: 1, 3;
}

.highcharts-sankey-series .highcharts-point {
  stroke-width: 0;
}

.highcharts-sankey-series .highcharts-link {
  transition: fill 250ms, fill-opacity 250ms;
  fill-opacity: 0.5;
}

.highcharts-sankey-series .highcharts-point-hover.highcharts-link {
  transition: fill 50ms, fill-opacity 50ms;
  fill-opacity: 1;
}

.highcharts-venn-series .highcharts-point {
  fill-opacity: 0.75;
  stroke: #cccccc;
  transition: stroke 250ms, fill-opacity 250ms;
}

.highcharts-venn-series .highcharts-point-hover {
  fill-opacity: 1;
  stroke: #cccccc;
}

/* Highstock */
.highcharts-navigator-mask-outside {
  fill-opacity: 0;
}

.highcharts-navigator-mask-inside {
  fill: #6685c2;
  /* navigator.maskFill option */
  fill-opacity: 0.25;
  cursor: ew-resize;
}

.highcharts-navigator-outline {
  stroke: #cccccc;
  fill: none;
}

.highcharts-navigator-handle {
  stroke: #cccccc;
  fill: #f2f2f2;
  cursor: ew-resize;
}

.highcharts-navigator-series {
  fill: #335cad;
  stroke: #335cad;
}

.highcharts-navigator-series .highcharts-graph {
  stroke-width: 1px;
}

.highcharts-navigator-series .highcharts-area {
  fill-opacity: 0.05;
}

.highcharts-navigator-xaxis .highcharts-axis-line {
  stroke-width: 0;
}

.highcharts-navigator-xaxis .highcharts-grid-line {
  stroke-width: 1px;
  stroke: #e6e6e6;
}

.highcharts-navigator-xaxis.highcharts-axis-labels {
  fill: #999999;
}

.highcharts-navigator-yaxis .highcharts-grid-line {
  stroke-width: 0;
}

.highcharts-scrollbar-thumb {
  fill: #cccccc;
  stroke: #cccccc;
  stroke-width: 1px;
}

.highcharts-scrollbar-button {
  fill: #e6e6e6;
  stroke: #cccccc;
  stroke-width: 1px;
}

.highcharts-scrollbar-arrow {
  fill: #666666;
}

.highcharts-scrollbar-rifles {
  stroke: #666666;
  stroke-width: 1px;
}

.highcharts-scrollbar-track {
  fill: #f2f2f2;
  stroke: #f2f2f2;
  stroke-width: 1px;
}

.highcharts-button {
  fill: #f7f7f7;
  stroke: #cccccc;
  cursor: default;
  stroke-width: 1px;
  transition: fill 250ms;
}

.highcharts-button text {
  fill: #333333;
}

.highcharts-button-hover {
  transition: fill 0ms;
  fill: #e6e6e6;
  stroke: #cccccc;
}

.highcharts-button-hover text {
  fill: #333333;
}

.highcharts-button-pressed {
  font-weight: bold;
  fill: #e6ebf5;
  stroke: #cccccc;
}

.highcharts-button-pressed text {
  fill: #333333;
  font-weight: bold;
}

.highcharts-button-disabled text {
  fill: #333333;
}

.highcharts-range-selector-buttons .highcharts-button {
  stroke-width: 0px;
}

.highcharts-range-label rect {
  fill: none;
}

.highcharts-range-label text {
  fill: #666666;
}

.highcharts-range-input rect {
  fill: none;
}

.highcharts-range-input text {
  fill: #333333;
}

.highcharts-range-input {
  stroke-width: 1px;
  stroke: #cccccc;
}

input.highcharts-range-selector {
  position: absolute;
  border: 0;
  width: 1px;
  /* Chrome needs a pixel to see it */
  height: 1px;
  padding: 0;
  text-align: center;
  left: -9em;
  /* #4798 */
}

.highcharts-crosshair-label text {
  fill: #ffffff;
  font-size: 1.1em;
}

.highcharts-crosshair-label .highcharts-label-box {
  fill: inherit;
}

.highcharts-candlestick-series .highcharts-point {
  stroke: #000000;
  stroke-width: 1px;
}

.highcharts-candlestick-series .highcharts-point-up {
  fill: #ffffff;
}

.highcharts-ohlc-series .highcharts-point-hover {
  stroke-width: 3px;
}

.highcharts-flags-series .highcharts-point .highcharts-label-box {
  stroke: #999999;
  fill: #ffffff;
  transition: fill 250ms;
}

.highcharts-flags-series .highcharts-point-hover .highcharts-label-box {
  stroke: #000000;
  fill: #ccd6eb;
}

.highcharts-flags-series .highcharts-point text {
  fill: #000000;
  font-size: 0.9em;
  font-weight: bold;
}

/* Highmaps */
.highcharts-map-series .highcharts-point {
  transition: fill 500ms, fill-opacity 500ms, stroke-width 250ms;
  stroke: #cccccc;
}

.highcharts-map-series .highcharts-point-hover {
  transition: fill 0ms, fill-opacity 0ms;
  fill-opacity: 0.5;
  stroke-width: 2px;
}

.highcharts-mapline-series .highcharts-point {
  fill: none;
}

.highcharts-heatmap-series .highcharts-point {
  stroke-width: 0;
}

.highcharts-map-navigation {
  font-size: 1.3em;
  font-weight: bold;
  text-align: center;
}

.highcharts-coloraxis {
  stroke-width: 0;
}

.highcharts-coloraxis-marker {
  fill: #999999;
}

.highcharts-null-point {
  fill: #f7f7f7;
}

/* 3d charts */
.highcharts-3d-frame {
  fill: transparent;
}

/* Exporting module */
.highcharts-contextbutton {
  fill: #ffffff;
  /* needed to capture hover */
  stroke: none;
  stroke-linecap: round;
}

.highcharts-contextbutton:hover {
  fill: #e6e6e6;
  stroke: #e6e6e6;
}

.highcharts-button-symbol {
  stroke: #666666;
  stroke-width: 3px;
}

.highcharts-menu {
  border: 1px solid #999999;
  background: #ffffff;
  padding: 5px 0;
  box-shadow: 3px 3px 10px #888;
}

.highcharts-menu-item {
  padding: 0.5em 1em;
  background: none;
  color: #333333;
  cursor: pointer;
  transition: background 250ms, color 250ms;
}

.highcharts-menu-item:hover {
  background: #335cad;
  color: #ffffff;
}

/* Drilldown module */
.highcharts-drilldown-point {
  cursor: pointer;
}

.highcharts-drilldown-data-label text,
text.highcharts-drilldown-data-label,
.highcharts-drilldown-axis-label {
  cursor: pointer;
  fill: #003399;
  font-weight: bold;
  text-decoration: underline;
}

/* No-data module */
.highcharts-no-data text {
  font-weight: bold;
  font-size: 12px;
  fill: #666666;
}

/* Drag-panes module */
.highcharts-axis-resizer {
  cursor: ns-resize;
  stroke: black;
  stroke-width: 2px;
}

/* Bullet type series */
.highcharts-bullet-target {
  stroke-width: 0;
}

/* Lineargauge type series */
.highcharts-lineargauge-target {
  stroke-width: 1px;
  stroke: #333333;
}

.highcharts-lineargauge-target-line {
  stroke-width: 1px;
  stroke: #333333;
}

/* Annotations module */
.highcharts-annotation-label-box {
  stroke-width: 1px;
  stroke: #000000;
  fill: #000000;
  fill-opacity: 0.75;
}

.highcharts-annotation-label text {
  fill: #e6e6e6;
}

/* Gantt */
.highcharts-treegrid-node-collapsed, .highcharts-treegrid-node-expanded {
  cursor: pointer;
}

.highcharts-point-connecting-path {
  fill: none;
}

.highcharts-grid-axis .highcharts-tick {
  stroke-width: 1px;
}

.highcharts-grid-axis .highcharts-axis-line {
  stroke-width: 1px;
}
`

var highchartsJS string = `
/*
 Highcharts JS v7.0.0 (2018-12-11)

 (c) 2009-2018 Torstein Honsi

 License: www.highcharts.com/license
*/
(function(O,J){"object"===typeof module&&module.exports?module.exports=O.document?J(O):J:"function"===typeof define&&define.amd?define(function(){return J(O)}):O.Highcharts=J(O)})("undefined"!==typeof window?window:this,function(O){var J=function(){var a="undefined"===typeof O?window:O,y=a.document,G=a.navigator&&a.navigator.userAgent||"",E=y&&y.createElementNS&&!!y.createElementNS("http://www.w3.org/2000/svg","svg").createSVGRect,h=/(edge|msie|trident)/i.test(G)&&!a.opera,d=-1!==G.indexOf("Firefox"),
r=-1!==G.indexOf("Chrome"),u=d&&4>parseInt(G.split("Firefox/")[1],10);return a.Highcharts?a.Highcharts.error(16,!0):{product:"Highcharts",version:"7.0.0",deg2rad:2*Math.PI/360,doc:y,hasBidiBug:u,hasTouch:y&&void 0!==y.documentElement.ontouchstart,isMS:h,isWebKit:-1!==G.indexOf("AppleWebKit"),isFirefox:d,isChrome:r,isSafari:!r&&-1!==G.indexOf("Safari"),isTouchDevice:/(Mobile|Android|Windows Phone)/.test(G),SVG_NS:"http://www.w3.org/2000/svg",chartCount:0,seriesTypes:{},symbolSizes:{},svg:E,win:a,marginNames:["plotTop",
"marginRight","marginBottom","plotLeft"],noop:function(){},charts:[]}}();(function(a){a.timers=[];var y=a.charts,G=a.doc,E=a.win;a.error=function(h,d,r){var u=a.isNumber(h)?"Highcharts error #"+h+": www.highcharts.com/errors/"+h:h;r&&a.fireEvent(r,"displayError",{code:h});if(d)throw Error(u);E.console&&console.log(u)};a.Fx=function(a,d,r){this.options=d;this.elem=a;this.prop=r};a.Fx.prototype={dSetter:function(){var a=this.paths[0],d=this.paths[1],r=[],u=this.now,v=a.length,w;if(1===u)r=this.toD;
else if(v===d.length&&1>u)for(;v--;)w=parseFloat(a[v]),r[v]=isNaN(w)?d[v]:u*parseFloat(d[v]-w)+w;else r=d;this.elem.attr("d",r,null,!0)},update:function(){var a=this.elem,d=this.prop,r=this.now,u=this.options.step;if(this[d+"Setter"])this[d+"Setter"]();else a.attr?a.element&&a.attr(d,r,null,!0):a.style[d]=r+this.unit;u&&u.call(a,r,this)},run:function(h,d,r){var u=this,v=u.options,w=function(a){return w.stopped?!1:u.step(a)},n=E.requestAnimationFrame||function(a){setTimeout(a,13)},g=function(){for(var c=
0;c<a.timers.length;c++)a.timers[c]()||a.timers.splice(c--,1);a.timers.length&&n(g)};h!==d||this.elem["forceAnimate:"+this.prop]?(this.startTime=+new Date,this.start=h,this.end=d,this.unit=r,this.now=this.start,this.pos=0,w.elem=this.elem,w.prop=this.prop,w()&&1===a.timers.push(w)&&n(g)):(delete v.curAnim[this.prop],v.complete&&0===Object.keys(v.curAnim).length&&v.complete.call(this.elem))},step:function(h){var d=+new Date,r,u=this.options,v=this.elem,w=u.complete,n=u.duration,g=u.curAnim;v.attr&&
!v.element?h=!1:h||d>=n+this.startTime?(this.now=this.end,this.pos=1,this.update(),r=g[this.prop]=!0,a.objectEach(g,function(a){!0!==a&&(r=!1)}),r&&w&&w.call(v),h=!1):(this.pos=u.easing((d-this.startTime)/n),this.now=this.start+(this.end-this.start)*this.pos,this.update(),h=!0);return h},initPath:function(h,d,r){function u(a){var b,k;for(f=a.length;f--;)b="M"===a[f]||"L"===a[f],k=/[a-zA-Z]/.test(a[f+3]),b&&k&&a.splice(f+1,0,a[f+1],a[f+2],a[f+1],a[f+2])}function v(a,l){for(;a.length<b;){a[0]=l[b-a.length];
var k=a.slice(0,p);[].splice.apply(a,[0,0].concat(k));x&&(k=a.slice(a.length-p),[].splice.apply(a,[a.length,0].concat(k)),f--)}a[0]="M"}function w(a,f){for(var k=(b-a.length)/p;0<k&&k--;)l=a.slice().splice(a.length/t-p,p*t),l[0]=f[b-p-k*p],m&&(l[p-6]=l[p-2],l[p-5]=l[p-1]),[].splice.apply(a,[a.length/t,0].concat(l)),x&&k--}d=d||"";var n,g=h.startX,c=h.endX,m=-1<d.indexOf("C"),p=m?7:3,b,l,f;d=d.split(" ");r=r.slice();var x=h.isArea,t=x?2:1,H;m&&(u(d),u(r));if(g&&c){for(f=0;f<g.length;f++)if(g[f]===
c[0]){n=f;break}else if(g[0]===c[c.length-g.length+f]){n=f;H=!0;break}void 0===n&&(d=[])}d.length&&a.isNumber(n)&&(b=r.length+n*t*p,H?(v(d,r),w(r,d)):(v(r,d),w(d,r)));return[d,r]},fillSetter:function(){a.Fx.prototype.strokeSetter.apply(this,arguments)},strokeSetter:function(){this.elem.attr(this.prop,a.color(this.start).tweenTo(a.color(this.end),this.pos),null,!0)}};a.merge=function(){var h,d=arguments,r,u={},v=function(d,n){"object"!==typeof d&&(d={});a.objectEach(n,function(g,c){!a.isObject(g,!0)||
a.isClass(g)||a.isDOMElement(g)?d[c]=n[c]:d[c]=v(d[c]||{},g)});return d};!0===d[0]&&(u=d[1],d=Array.prototype.slice.call(d,2));r=d.length;for(h=0;h<r;h++)u=v(u,d[h]);return u};a.pInt=function(a,d){return parseInt(a,d||10)};a.isString=function(a){return"string"===typeof a};a.isArray=function(a){a=Object.prototype.toString.call(a);return"[object Array]"===a||"[object Array Iterator]"===a};a.isObject=function(h,d){return!!h&&"object"===typeof h&&(!d||!a.isArray(h))};a.isDOMElement=function(h){return a.isObject(h)&&
"number"===typeof h.nodeType};a.isClass=function(h){var d=h&&h.constructor;return!(!a.isObject(h,!0)||a.isDOMElement(h)||!d||!d.name||"Object"===d.name)};a.isNumber=function(a){return"number"===typeof a&&!isNaN(a)&&Infinity>a&&-Infinity<a};a.erase=function(a,d){for(var h=a.length;h--;)if(a[h]===d){a.splice(h,1);break}};a.defined=function(a){return void 0!==a&&null!==a};a.attr=function(h,d,r){var u;a.isString(d)?a.defined(r)?h.setAttribute(d,r):h&&h.getAttribute&&((u=h.getAttribute(d))||"class"!==
d||(u=h.getAttribute(d+"Name"))):a.defined(d)&&a.isObject(d)&&a.objectEach(d,function(a,d){h.setAttribute(d,a)});return u};a.splat=function(h){return a.isArray(h)?h:[h]};a.syncTimeout=function(a,d,r){if(d)return setTimeout(a,d,r);a.call(0,r)};a.clearTimeout=function(h){a.defined(h)&&clearTimeout(h)};a.extend=function(a,d){var h;a||(a={});for(h in d)a[h]=d[h];return a};a.pick=function(){var a=arguments,d,r,u=a.length;for(d=0;d<u;d++)if(r=a[d],void 0!==r&&null!==r)return r};a.css=function(h,d){a.isMS&&
!a.svg&&d&&void 0!==d.opacity&&(d.filter="alpha(opacity\x3d"+100*d.opacity+")");a.extend(h.style,d)};a.createElement=function(h,d,r,u,v){h=G.createElement(h);var w=a.css;d&&a.extend(h,d);v&&w(h,{padding:0,border:"none",margin:0});r&&w(h,r);u&&u.appendChild(h);return h};a.extendClass=function(h,d){var r=function(){};r.prototype=new h;a.extend(r.prototype,d);return r};a.pad=function(a,d,r){return Array((d||2)+1-String(a).replace("-","").length).join(r||0)+a};a.relativeLength=function(a,d,r){return/%$/.test(a)?
d*parseFloat(a)/100+(r||0):parseFloat(a)};a.wrap=function(a,d,r){var h=a[d];a[d]=function(){var a=Array.prototype.slice.call(arguments),d=arguments,n=this;n.proceed=function(){h.apply(n,arguments.length?arguments:d)};a.unshift(h);a=r.apply(this,a);n.proceed=null;return a}};a.datePropsToTimestamps=function(h){a.objectEach(h,function(d,r){a.isObject(d)&&"function"===typeof d.getTime?h[r]=d.getTime():(a.isObject(d)||a.isArray(d))&&a.datePropsToTimestamps(d)})};a.formatSingle=function(h,d,r){var u=/\.([0-9])/,
v=a.defaultOptions.lang;/f$/.test(h)?(r=(r=h.match(u))?r[1]:-1,null!==d&&(d=a.numberFormat(d,r,v.decimalPoint,-1<h.indexOf(",")?v.thousandsSep:""))):d=(r||a.time).dateFormat(h,d);return d};a.format=function(h,d,r){for(var u="{",v=!1,w,n,g,c,m=[],p;h;){u=h.indexOf(u);if(-1===u)break;w=h.slice(0,u);if(v){w=w.split(":");n=w.shift().split(".");c=n.length;p=d;for(g=0;g<c;g++)p&&(p=p[n[g]]);w.length&&(p=a.formatSingle(w.join(":"),p,r));m.push(p)}else m.push(w);h=h.slice(u+1);u=(v=!v)?"}":"{"}m.push(h);
return m.join("")};a.getMagnitude=function(a){return Math.pow(10,Math.floor(Math.log(a)/Math.LN10))};a.normalizeTickInterval=function(h,d,r,u,v){var w,n=h;r=a.pick(r,1);w=h/r;d||(d=v?[1,1.2,1.5,2,2.5,3,4,5,6,8,10]:[1,2,2.5,5,10],!1===u&&(1===r?d=d.filter(function(a){return 0===a%1}):.1>=r&&(d=[1/r])));for(u=0;u<d.length&&!(n=d[u],v&&n*r>=h||!v&&w<=(d[u]+(d[u+1]||d[u]))/2);u++);return n=a.correctFloat(n*r,-Math.round(Math.log(.001)/Math.LN10))};a.stableSort=function(a,d){var h=a.length,u,v;for(v=0;v<
h;v++)a[v].safeI=v;a.sort(function(a,n){u=d(a,n);return 0===u?a.safeI-n.safeI:u});for(v=0;v<h;v++)delete a[v].safeI};a.arrayMin=function(a){for(var d=a.length,h=a[0];d--;)a[d]<h&&(h=a[d]);return h};a.arrayMax=function(a){for(var d=a.length,h=a[0];d--;)a[d]>h&&(h=a[d]);return h};a.destroyObjectProperties=function(h,d){a.objectEach(h,function(a,u){a&&a!==d&&a.destroy&&a.destroy();delete h[u]})};a.discardElement=function(h){var d=a.garbageBin;d||(d=a.createElement("div"));h&&d.appendChild(h);d.innerHTML=
""};a.correctFloat=function(a,d){return parseFloat(a.toPrecision(d||14))};a.setAnimation=function(h,d){d.renderer.globalAnimation=a.pick(h,d.options.chart.animation,!0)};a.animObject=function(h){return a.isObject(h)?a.merge(h):{duration:h?500:0}};a.timeUnits={millisecond:1,second:1E3,minute:6E4,hour:36E5,day:864E5,week:6048E5,month:24192E5,year:314496E5};a.numberFormat=function(h,d,r,u){h=+h||0;d=+d;var v=a.defaultOptions.lang,w=(h.toString().split(".")[1]||"").split("e")[0].length,n,g,c=h.toString().split("e");
-1===d?d=Math.min(w,20):a.isNumber(d)?d&&c[1]&&0>c[1]&&(n=d+ +c[1],0<=n?(c[0]=(+c[0]).toExponential(n).split("e")[0],d=n):(c[0]=c[0].split(".")[0]||0,h=20>d?(c[0]*Math.pow(10,c[1])).toFixed(d):0,c[1]=0)):d=2;g=(Math.abs(c[1]?c[0]:h)+Math.pow(10,-Math.max(d,w)-1)).toFixed(d);w=String(a.pInt(g));n=3<w.length?w.length%3:0;r=a.pick(r,v.decimalPoint);u=a.pick(u,v.thousandsSep);h=(0>h?"-":"")+(n?w.substr(0,n)+u:"");h+=w.substr(n).replace(/(\d{3})(?=\d)/g,"$1"+u);d&&(h+=r+g.slice(-d));c[1]&&0!==+h&&(h+=
"e"+c[1]);return h};Math.easeInOutSine=function(a){return-.5*(Math.cos(Math.PI*a)-1)};a.getStyle=function(h,d,r){if("width"===d)return Math.max(0,Math.min(h.offsetWidth,h.scrollWidth,h.getBoundingClientRect?Math.floor(h.getBoundingClientRect().width):Infinity)-a.getStyle(h,"padding-left")-a.getStyle(h,"padding-right"));if("height"===d)return Math.max(0,Math.min(h.offsetHeight,h.scrollHeight)-a.getStyle(h,"padding-top")-a.getStyle(h,"padding-bottom"));E.getComputedStyle||a.error(27,!0);if(h=E.getComputedStyle(h,
void 0))h=h.getPropertyValue(d),a.pick(r,"opacity"!==d)&&(h=a.pInt(h));return h};a.inArray=function(a,d,r){return d.indexOf(a,r)};a.find=Array.prototype.find?function(a,d){return a.find(d)}:function(a,d){var h,u=a.length;for(h=0;h<u;h++)if(d(a[h],h))return a[h]};a.keys=Object.keys;a.offset=function(a){var d=G.documentElement;a=a.parentElement||a.parentNode?a.getBoundingClientRect():{top:0,left:0};return{top:a.top+(E.pageYOffset||d.scrollTop)-(d.clientTop||0),left:a.left+(E.pageXOffset||d.scrollLeft)-
(d.clientLeft||0)}};a.stop=function(h,d){for(var r=a.timers.length;r--;)a.timers[r].elem!==h||d&&d!==a.timers[r].prop||(a.timers[r].stopped=!0)};a.objectEach=function(a,d,r){for(var h in a)a.hasOwnProperty(h)&&d.call(r||a[h],a[h],h,a)};a.objectEach({map:"map",each:"forEach",grep:"filter",reduce:"reduce",some:"some"},function(h,d){a[d]=function(a){return Array.prototype[h].apply(a,[].slice.call(arguments,1))}});a.addEvent=function(h,d,r,u){var v,w=h.addEventListener||a.addEventListenerPolyfill;v="function"===
typeof h&&h.prototype?h.prototype.protoEvents=h.prototype.protoEvents||{}:h.hcEvents=h.hcEvents||{};a.Point&&h instanceof a.Point&&h.series&&h.series.chart&&(h.series.chart.runTrackerClick=!0);w&&w.call(h,d,r,!1);v[d]||(v[d]=[]);v[d].push(r);u&&a.isNumber(u.order)&&(r.order=u.order,v[d].sort(function(a,g){return a.order-g.order}));return function(){a.removeEvent(h,d,r)}};a.removeEvent=function(h,d,r){function u(g,c){var m=h.removeEventListener||a.removeEventListenerPolyfill;m&&m.call(h,g,c,!1)}function v(g){var c,
m;h.nodeName&&(d?(c={},c[d]=!0):c=g,a.objectEach(c,function(a,b){if(g[b])for(m=g[b].length;m--;)u(b,g[b][m])}))}var w,n;["protoEvents","hcEvents"].forEach(function(a){var c=h[a];c&&(d?(w=c[d]||[],r?(n=w.indexOf(r),-1<n&&(w.splice(n,1),c[d]=w),u(d,r)):(v(c),c[d]=[])):(v(c),h[a]={}))})};a.fireEvent=function(h,d,r,u){var v,w,n,g,c;r=r||{};G.createEvent&&(h.dispatchEvent||h.fireEvent)?(v=G.createEvent("Events"),v.initEvent(d,!0,!0),a.extend(v,r),h.dispatchEvent?h.dispatchEvent(v):h.fireEvent(d,v)):["protoEvents",
"hcEvents"].forEach(function(m){if(h[m])for(w=h[m][d]||[],n=w.length,r.target||a.extend(r,{preventDefault:function(){r.defaultPrevented=!0},target:h,type:d}),g=0;g<n;g++)(c=w[g])&&!1===c.call(h,r)&&r.preventDefault()});u&&!r.defaultPrevented&&u.call(h,r)};a.animate=function(h,d,r){var u,v="",w,n,g;a.isObject(r)||(g=arguments,r={duration:g[2],easing:g[3],complete:g[4]});a.isNumber(r.duration)||(r.duration=400);r.easing="function"===typeof r.easing?r.easing:Math[r.easing]||Math.easeInOutSine;r.curAnim=
a.merge(d);a.objectEach(d,function(c,g){a.stop(h,g);n=new a.Fx(h,r,g);w=null;"d"===g?(n.paths=n.initPath(h,h.d,d.d),n.toD=d.d,u=0,w=1):h.attr?u=h.attr(g):(u=parseFloat(a.getStyle(h,g))||0,"opacity"!==g&&(v="px"));w||(w=c);w&&w.match&&w.match("px")&&(w=w.replace(/px/g,""));n.run(u,w,v)})};a.seriesType=function(h,d,r,u,v){var w=a.getOptions(),n=a.seriesTypes;w.plotOptions[h]=a.merge(w.plotOptions[d],r);n[h]=a.extendClass(n[d]||function(){},u);n[h].prototype.type=h;v&&(n[h].prototype.pointClass=a.extendClass(a.Point,
v));return n[h]};a.uniqueKey=function(){var a=Math.random().toString(36).substring(2,9),d=0;return function(){return"highcharts-"+a+"-"+d++}}();a.isFunction=function(a){return"function"===typeof a};E.jQuery&&(E.jQuery.fn.highcharts=function(){var h=[].slice.call(arguments);if(this[0])return h[0]?(new (a[a.isString(h[0])?h.shift():"Chart"])(this[0],h[0],h[1]),this):y[a.attr(this[0],"data-highcharts-chart")]})})(J);(function(a){var y=a.isNumber,G=a.merge,E=a.pInt;a.Color=function(h){if(!(this instanceof
a.Color))return new a.Color(h);this.init(h)};a.Color.prototype={parsers:[{regex:/rgba\(\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})\s*,\s*([0-9]?(?:\.[0-9]+)?)\s*\)/,parse:function(a){return[E(a[1]),E(a[2]),E(a[3]),parseFloat(a[4],10)]}},{regex:/rgb\(\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})\s*\)/,parse:function(a){return[E(a[1]),E(a[2]),E(a[3]),1]}}],names:{white:"#ffffff",black:"#000000"},init:function(h){var d,r,u,v;if((this.input=h=this.names[h&&h.toLowerCase?h.toLowerCase():
""]||h)&&h.stops)this.stops=h.stops.map(function(d){return new a.Color(d[1])});else if(h&&h.charAt&&"#"===h.charAt()&&(d=h.length,h=parseInt(h.substr(1),16),7===d?r=[(h&16711680)>>16,(h&65280)>>8,h&255,1]:4===d&&(r=[(h&3840)>>4|(h&3840)>>8,(h&240)>>4|h&240,(h&15)<<4|h&15,1])),!r)for(u=this.parsers.length;u--&&!r;)v=this.parsers[u],(d=v.regex.exec(h))&&(r=v.parse(d));this.rgba=r||[]},get:function(a){var d=this.input,h=this.rgba,u;this.stops?(u=G(d),u.stops=[].concat(u.stops),this.stops.forEach(function(d,
h){u.stops[h]=[u.stops[h][0],d.get(a)]})):u=h&&y(h[0])?"rgb"===a||!a&&1===h[3]?"rgb("+h[0]+","+h[1]+","+h[2]+")":"a"===a?h[3]:"rgba("+h.join(",")+")":d;return u},brighten:function(a){var d,h=this.rgba;if(this.stops)this.stops.forEach(function(d){d.brighten(a)});else if(y(a)&&0!==a)for(d=0;3>d;d++)h[d]+=E(255*a),0>h[d]&&(h[d]=0),255<h[d]&&(h[d]=255);return this},setOpacity:function(a){this.rgba[3]=a;return this},tweenTo:function(a,d){var h=this.rgba,u=a.rgba;u.length&&h&&h.length?(a=1!==u[3]||1!==
h[3],d=(a?"rgba(":"rgb(")+Math.round(u[0]+(h[0]-u[0])*(1-d))+","+Math.round(u[1]+(h[1]-u[1])*(1-d))+","+Math.round(u[2]+(h[2]-u[2])*(1-d))+(a?","+(u[3]+(h[3]-u[3])*(1-d)):"")+")"):d=a.input||"none";return d}};a.color=function(h){return new a.Color(h)}})(J);(function(a){var y,G,E=a.addEvent,h=a.animate,d=a.attr,r=a.charts,u=a.color,v=a.css,w=a.createElement,n=a.defined,g=a.deg2rad,c=a.destroyObjectProperties,m=a.doc,p=a.extend,b=a.erase,l=a.hasTouch,f=a.isArray,x=a.isFirefox,t=a.isMS,H=a.isObject,
F=a.isString,z=a.isWebKit,k=a.merge,A=a.noop,D=a.objectEach,C=a.pick,e=a.pInt,q=a.removeEvent,L=a.splat,I=a.stop,R=a.svg,K=a.SVG_NS,M=a.symbolSizes,S=a.win;y=a.SVGElement=function(){return this};p(y.prototype,{opacity:1,SVG_NS:K,textProps:"direction fontSize fontWeight fontFamily fontStyle color lineHeight width textAlign textDecoration textOverflow textOutline cursor".split(" "),init:function(a,e){this.element="span"===e?w(e):m.createElementNS(this.SVG_NS,e);this.renderer=a},animate:function(e,q,
b){q=a.animObject(C(q,this.renderer.globalAnimation,!0));0!==q.duration?(b&&(q.complete=b),h(this,e,q)):(this.attr(e,null,b),q.step&&q.step.call(this));return this},complexColor:function(e,q,b){var B=this.renderer,l,c,p,g,A,K,m,N,x,d,t,I=[],L;a.fireEvent(this.renderer,"complexColor",{args:arguments},function(){e.radialGradient?c="radialGradient":e.linearGradient&&(c="linearGradient");c&&(p=e[c],A=B.gradients,m=e.stops,d=b.radialReference,f(p)&&(e[c]=p={x1:p[0],y1:p[1],x2:p[2],y2:p[3],gradientUnits:"userSpaceOnUse"}),
"radialGradient"===c&&d&&!n(p.gradientUnits)&&(g=p,p=k(p,B.getRadialAttr(d,g),{gradientUnits:"userSpaceOnUse"})),D(p,function(a,e){"id"!==e&&I.push(e,a)}),D(m,function(a){I.push(a)}),I=I.join(","),A[I]?t=A[I].attr("id"):(p.id=t=a.uniqueKey(),A[I]=K=B.createElement(c).attr(p).add(B.defs),K.radAttr=g,K.stops=[],m.forEach(function(e){0===e[1].indexOf("rgba")?(l=a.color(e[1]),N=l.get("rgb"),x=l.get("a")):(N=e[1],x=1);e=B.createElement("stop").attr({offset:e[0],"stop-color":N,"stop-opacity":x}).add(K);
K.stops.push(e)})),L="url("+B.url+"#"+t+")",b.setAttribute(q,L),b.gradient=I,e.toString=function(){return L})})},applyTextOutline:function(e){var B=this.element,q,f,k,l,c;-1!==e.indexOf("contrast")&&(e=e.replace(/contrast/g,this.renderer.getContrast(B.style.fill)));e=e.split(" ");f=e[e.length-1];if((k=e[0])&&"none"!==k&&a.svg){this.fakeTS=!0;e=[].slice.call(B.getElementsByTagName("tspan"));this.ySetter=this.xSetter;k=k.replace(/(^[\d\.]+)(.*?)$/g,function(a,e,B){return 2*e+B});for(c=e.length;c--;)q=
e[c],"highcharts-text-outline"===q.getAttribute("class")&&b(e,B.removeChild(q));l=B.firstChild;e.forEach(function(a,e){0===e&&(a.setAttribute("x",B.getAttribute("x")),e=B.getAttribute("y"),a.setAttribute("y",e||0),null===e&&B.setAttribute("y",0));a=a.cloneNode(1);d(a,{"class":"highcharts-text-outline",fill:f,stroke:f,"stroke-width":k,"stroke-linejoin":"round"});B.insertBefore(a,l)})}},symbolCustomAttribs:"x y width height r start end innerR anchorX anchorY rounded".split(" "),attr:function(e,q,b,
f){var B,k=this.element,l,c=this,p,g,A=this.symbolCustomAttribs;"string"===typeof e&&void 0!==q&&(B=e,e={},e[B]=q);"string"===typeof e?c=(this[e+"Getter"]||this._defaultGetter).call(this,e,k):(D(e,function(B,q){p=!1;f||I(this,q);this.symbolName&&-1!==a.inArray(q,A)&&(l||(this.symbolAttr(e),l=!0),p=!0);!this.rotation||"x"!==q&&"y"!==q||(this.doTransform=!0);p||(g=this[q+"Setter"]||this._defaultSetter,g.call(this,B,q,k),!this.styledMode&&this.shadows&&/^(width|height|visibility|x|y|d|transform|cx|cy|r)$/.test(q)&&
this.updateShadows(q,B,g))},this),this.afterSetters());b&&b.call(this);return c},afterSetters:function(){this.doTransform&&(this.updateTransform(),this.doTransform=!1)},updateShadows:function(a,e,q){for(var B=this.shadows,b=B.length;b--;)q.call(B[b],"height"===a?Math.max(e-(B[b].cutHeight||0),0):"d"===a?this.d:e,a,B[b])},addClass:function(a,e){var B=this.attr("class")||"";-1===B.indexOf(a)&&(e||(a=(B+(B?" ":"")+a).replace("  "," ")),this.attr("class",a));return this},hasClass:function(a){return-1!==
(this.attr("class")||"").split(" ").indexOf(a)},removeClass:function(a){return this.attr("class",(this.attr("class")||"").replace(a,""))},symbolAttr:function(a){var e=this;"x y r start end width height innerR anchorX anchorY".split(" ").forEach(function(B){e[B]=C(a[B],e[B])});e.attr({d:e.renderer.symbols[e.symbolName](e.x,e.y,e.width,e.height,e)})},clip:function(a){return this.attr("clip-path",a?"url("+this.renderer.url+"#"+a.id+")":"none")},crisp:function(a,e){var B;e=e||a.strokeWidth||0;B=Math.round(e)%
2/2;a.x=Math.floor(a.x||this.x||0)+B;a.y=Math.floor(a.y||this.y||0)+B;a.width=Math.floor((a.width||this.width||0)-2*B);a.height=Math.floor((a.height||this.height||0)-2*B);n(a.strokeWidth)&&(a.strokeWidth=e);return a},css:function(a){var B=this.styles,q={},b=this.element,k,f="",l,c=!B,g=["textOutline","textOverflow","width"];a&&a.color&&(a.fill=a.color);B&&D(a,function(a,e){a!==B[e]&&(q[e]=a,c=!0)});c&&(B&&(a=p(B,q)),a&&(null===a.width||"auto"===a.width?delete this.textWidth:"text"===b.nodeName.toLowerCase()&&
a.width&&(k=this.textWidth=e(a.width))),this.styles=a,k&&!R&&this.renderer.forExport&&delete a.width,b.namespaceURI===this.SVG_NS?(l=function(a,e){return"-"+e.toLowerCase()},D(a,function(a,e){-1===g.indexOf(e)&&(f+=e.replace(/([A-Z])/g,l)+":"+a+";")}),f&&d(b,"style",f)):v(b,a),this.added&&("text"===this.element.nodeName&&this.renderer.buildText(this),a&&a.textOutline&&this.applyTextOutline(a.textOutline)));return this},getStyle:function(a){return S.getComputedStyle(this.element||this,"").getPropertyValue(a)},
strokeWidth:function(){if(!this.renderer.styledMode)return this["stroke-width"]||0;var a=this.getStyle("stroke-width"),q;a.indexOf("px")===a.length-2?a=e(a):(q=m.createElementNS(K,"rect"),d(q,{width:a,"stroke-width":0}),this.element.parentNode.appendChild(q),a=q.getBBox().width,q.parentNode.removeChild(q));return a},on:function(a,e){var q=this,B=q.element;l&&"click"===a?(B.ontouchstart=function(a){q.touchEventFired=Date.now();a.preventDefault();e.call(B,a)},B.onclick=function(a){(-1===S.navigator.userAgent.indexOf("Android")||
1100<Date.now()-(q.touchEventFired||0))&&e.call(B,a)}):B["on"+a]=e;return this},setRadialReference:function(a){var e=this.renderer.gradients[this.element.gradient];this.element.radialReference=a;e&&e.radAttr&&e.animate(this.renderer.getRadialAttr(a,e.radAttr));return this},translate:function(a,e){return this.attr({translateX:a,translateY:e})},invert:function(a){this.inverted=a;this.updateTransform();return this},updateTransform:function(){var a=this.translateX||0,e=this.translateY||0,q=this.scaleX,
b=this.scaleY,f=this.inverted,k=this.rotation,l=this.matrix,c=this.element;f&&(a+=this.width,e+=this.height);a=["translate("+a+","+e+")"];n(l)&&a.push("matrix("+l.join(",")+")");f?a.push("rotate(90) scale(-1,1)"):k&&a.push("rotate("+k+" "+C(this.rotationOriginX,c.getAttribute("x"),0)+" "+C(this.rotationOriginY,c.getAttribute("y")||0)+")");(n(q)||n(b))&&a.push("scale("+C(q,1)+" "+C(b,1)+")");a.length&&c.setAttribute("transform",a.join(" "))},toFront:function(){var a=this.element;a.parentNode.appendChild(a);
return this},align:function(a,e,q){var B,f,k,l,c={};f=this.renderer;k=f.alignedObjects;var p,g;if(a){if(this.alignOptions=a,this.alignByTranslate=e,!q||F(q))this.alignTo=B=q||"renderer",b(k,this),k.push(this),q=null}else a=this.alignOptions,e=this.alignByTranslate,B=this.alignTo;q=C(q,f[B],f);B=a.align;f=a.verticalAlign;k=(q.x||0)+(a.x||0);l=(q.y||0)+(a.y||0);"right"===B?p=1:"center"===B&&(p=2);p&&(k+=(q.width-(a.width||0))/p);c[e?"translateX":"x"]=Math.round(k);"bottom"===f?g=1:"middle"===f&&(g=
2);g&&(l+=(q.height-(a.height||0))/g);c[e?"translateY":"y"]=Math.round(l);this[this.placed?"animate":"attr"](c);this.placed=!0;this.alignAttr=c;return this},getBBox:function(a,e){var q,B=this.renderer,b,f=this.element,k=this.styles,l,c=this.textStr,A,K=B.cache,m=B.cacheKeys,x=f.namespaceURI===this.SVG_NS,d;e=C(e,this.rotation);b=e*g;l=B.styledMode?f&&y.prototype.getStyle.call(f,"font-size"):k&&k.fontSize;n(c)&&(d=c.toString(),-1===d.indexOf("\x3c")&&(d=d.replace(/[0-9]/g,"0")),d+=["",e||0,l,this.textWidth,
k&&k.textOverflow].join());d&&!a&&(q=K[d]);if(!q){if(x||B.forExport){try{(A=this.fakeTS&&function(a){[].forEach.call(f.querySelectorAll(".highcharts-text-outline"),function(e){e.style.display=a})})&&A("none"),q=f.getBBox?p({},f.getBBox()):{width:f.offsetWidth,height:f.offsetHeight},A&&A("")}catch(W){}if(!q||0>q.width)q={width:0,height:0}}else q=this.htmlGetBBox();B.isSVG&&(a=q.width,B=q.height,x&&(q.height=B={"11px,17":14,"13px,20":16}[k&&k.fontSize+","+Math.round(B)]||B),e&&(q.width=Math.abs(B*Math.sin(b))+
Math.abs(a*Math.cos(b)),q.height=Math.abs(B*Math.cos(b))+Math.abs(a*Math.sin(b))));if(d&&0<q.height){for(;250<m.length;)delete K[m.shift()];K[d]||m.push(d);K[d]=q}}return q},show:function(a){return this.attr({visibility:a?"inherit":"visible"})},hide:function(){return this.attr({visibility:"hidden"})},fadeOut:function(a){var e=this;e.animate({opacity:0},{duration:a||150,complete:function(){e.attr({y:-9999})}})},add:function(a){var e=this.renderer,q=this.element,B;a&&(this.parentGroup=a);this.parentInverted=
a&&a.inverted;void 0!==this.textStr&&e.buildText(this);this.added=!0;if(!a||a.handleZ||this.zIndex)B=this.zIndexSetter();B||(a?a.element:e.box).appendChild(q);if(this.onAdd)this.onAdd();return this},safeRemoveChild:function(a){var e=a.parentNode;e&&e.removeChild(a)},destroy:function(){var a=this,e=a.element||{},q=a.renderer,f=q.isSVG&&"SPAN"===e.nodeName&&a.parentGroup,k=e.ownerSVGElement,l=a.clipPath;e.onclick=e.onmouseout=e.onmouseover=e.onmousemove=e.point=null;I(a);l&&k&&([].forEach.call(k.querySelectorAll("[clip-path],[CLIP-PATH]"),
function(a){var e=a.getAttribute("clip-path"),q=l.element.id;(-1<e.indexOf("(#"+q+")")||-1<e.indexOf('("#'+q+'")'))&&a.removeAttribute("clip-path")}),a.clipPath=l.destroy());if(a.stops){for(k=0;k<a.stops.length;k++)a.stops[k]=a.stops[k].destroy();a.stops=null}a.safeRemoveChild(e);for(q.styledMode||a.destroyShadows();f&&f.div&&0===f.div.childNodes.length;)e=f.parentGroup,a.safeRemoveChild(f.div),delete f.div,f=e;a.alignTo&&b(q.alignedObjects,a);D(a,function(e,q){delete a[q]});return null},shadow:function(a,
e,q){var b=[],f,B,k=this.element,l,c,p,g;if(!a)this.destroyShadows();else if(!this.shadows){c=C(a.width,3);p=(a.opacity||.15)/c;g=this.parentInverted?"(-1,-1)":"("+C(a.offsetX,1)+", "+C(a.offsetY,1)+")";for(f=1;f<=c;f++)B=k.cloneNode(0),l=2*c+1-2*f,d(B,{stroke:a.color||"#000000","stroke-opacity":p*f,"stroke-width":l,transform:"translate"+g,fill:"none"}),B.setAttribute("class",(B.getAttribute("class")||"")+" highcharts-shadow"),q&&(d(B,"height",Math.max(d(B,"height")-l,0)),B.cutHeight=l),e?e.element.appendChild(B):
k.parentNode&&k.parentNode.insertBefore(B,k),b.push(B);this.shadows=b}return this},destroyShadows:function(){(this.shadows||[]).forEach(function(a){this.safeRemoveChild(a)},this);this.shadows=void 0},xGetter:function(a){"circle"===this.element.nodeName&&("x"===a?a="cx":"y"===a&&(a="cy"));return this._defaultGetter(a)},_defaultGetter:function(a){a=C(this[a+"Value"],this[a],this.element?this.element.getAttribute(a):null,0);/^[\-0-9\.]+$/.test(a)&&(a=parseFloat(a));return a},dSetter:function(a,e,q){a&&
a.join&&(a=a.join(" "));/(NaN| {2}|^$)/.test(a)&&(a="M 0 0");this[e]!==a&&(q.setAttribute(e,a),this[e]=a)},dashstyleSetter:function(a){var q,f=this["stroke-width"];"inherit"===f&&(f=1);if(a=a&&a.toLowerCase()){a=a.replace("shortdashdotdot","3,1,1,1,1,1,").replace("shortdashdot","3,1,1,1").replace("shortdot","1,1,").replace("shortdash","3,1,").replace("longdash","8,3,").replace(/dot/g,"1,3,").replace("dash","4,3,").replace(/,$/,"").split(",");for(q=a.length;q--;)a[q]=e(a[q])*f;a=a.join(",").replace(/NaN/g,
"none");this.element.setAttribute("stroke-dasharray",a)}},alignSetter:function(a){this.alignValue=a;this.element.setAttribute("text-anchor",{left:"start",center:"middle",right:"end"}[a])},opacitySetter:function(a,e,q){this[e]=a;q.setAttribute(e,a)},titleSetter:function(a){var e=this.element.getElementsByTagName("title")[0];e||(e=m.createElementNS(this.SVG_NS,"title"),this.element.appendChild(e));e.firstChild&&e.removeChild(e.firstChild);e.appendChild(m.createTextNode(String(C(a),"").replace(/<[^>]*>/g,
"").replace(/&lt;/g,"\x3c").replace(/&gt;/g,"\x3e")))},textSetter:function(a){a!==this.textStr&&(delete this.bBox,this.textStr=a,this.added&&this.renderer.buildText(this))},fillSetter:function(a,e,q){"string"===typeof a?q.setAttribute(e,a):a&&this.complexColor(a,e,q)},visibilitySetter:function(a,e,q){"inherit"===a?q.removeAttribute(e):this[e]!==a&&q.setAttribute(e,a);this[e]=a},zIndexSetter:function(a,q){var f=this.renderer,b=this.parentGroup,k=(b||f).element||f.box,l,c=this.element,B,p,f=k===f.box;
l=this.added;var g;n(a)?(c.setAttribute("data-z-index",a),a=+a,this[q]===a&&(l=!1)):n(this[q])&&c.removeAttribute("data-z-index");this[q]=a;if(l){(a=this.zIndex)&&b&&(b.handleZ=!0);q=k.childNodes;for(g=q.length-1;0<=g&&!B;g--)if(b=q[g],l=b.getAttribute("data-z-index"),p=!n(l),b!==c)if(0>a&&p&&!f&&!g)k.insertBefore(c,q[g]),B=!0;else if(e(l)<=a||p&&(!n(a)||0<=a))k.insertBefore(c,q[g+1]||null),B=!0;B||(k.insertBefore(c,q[f?3:0]||null),B=!0)}return B},_defaultSetter:function(a,e,q){q.setAttribute(e,a)}});
y.prototype.yGetter=y.prototype.xGetter;y.prototype.translateXSetter=y.prototype.translateYSetter=y.prototype.rotationSetter=y.prototype.verticalAlignSetter=y.prototype.rotationOriginXSetter=y.prototype.rotationOriginYSetter=y.prototype.scaleXSetter=y.prototype.scaleYSetter=y.prototype.matrixSetter=function(a,e){this[e]=a;this.doTransform=!0};y.prototype["stroke-widthSetter"]=y.prototype.strokeSetter=function(a,e,q){this[e]=a;this.stroke&&this["stroke-width"]?(y.prototype.fillSetter.call(this,this.stroke,
"stroke",q),q.setAttribute("stroke-width",this["stroke-width"]),this.hasStroke=!0):"stroke-width"===e&&0===a&&this.hasStroke&&(q.removeAttribute("stroke"),this.hasStroke=!1)};G=a.SVGRenderer=function(){this.init.apply(this,arguments)};p(G.prototype,{Element:y,SVG_NS:K,init:function(a,e,q,f,b,k,l){var c;c=this.createElement("svg").attr({version:"1.1","class":"highcharts-root"});l||c.css(this.getStyle(f));f=c.element;a.appendChild(f);d(a,"dir","ltr");-1===a.innerHTML.indexOf("xmlns")&&d(f,"xmlns",this.SVG_NS);
this.isSVG=!0;this.box=f;this.boxWrapper=c;this.alignedObjects=[];this.url=(x||z)&&m.getElementsByTagName("base").length?S.location.href.split("#")[0].replace(/<[^>]*>/g,"").replace(/([\('\)])/g,"\\$1").replace(/ /g,"%20"):"";this.createElement("desc").add().element.appendChild(m.createTextNode("Created with Highcharts 7.0.0"));this.defs=this.createElement("defs").add();this.allowHTML=k;this.forExport=b;this.styledMode=l;this.gradients={};this.cache={};this.cacheKeys=[];this.imgCount=0;this.setSize(e,
q,!1);var B;x&&a.getBoundingClientRect&&(e=function(){v(a,{left:0,top:0});B=a.getBoundingClientRect();v(a,{left:Math.ceil(B.left)-B.left+"px",top:Math.ceil(B.top)-B.top+"px"})},e(),this.unSubPixelFix=E(S,"resize",e))},definition:function(a){function e(a,f){var b;L(a).forEach(function(a){var k=q.createElement(a.tagName),l={};D(a,function(a,e){"tagName"!==e&&"children"!==e&&"textContent"!==e&&(l[e]=a)});k.attr(l);k.add(f||q.defs);a.textContent&&k.element.appendChild(m.createTextNode(a.textContent));
e(a.children||[],k);b=k});return b}var q=this;return e(a)},getStyle:function(a){return this.style=p({fontFamily:'"Lucida Grande", "Lucida Sans Unicode", Arial, Helvetica, sans-serif',fontSize:"12px"},a)},setStyle:function(a){this.boxWrapper.css(this.getStyle(a))},isHidden:function(){return!this.boxWrapper.getBBox().width},destroy:function(){var a=this.defs;this.box=null;this.boxWrapper=this.boxWrapper.destroy();c(this.gradients||{});this.gradients=null;a&&(this.defs=a.destroy());this.unSubPixelFix&&
this.unSubPixelFix();return this.alignedObjects=null},createElement:function(a){var e=new this.Element;e.init(this,a);return e},draw:A,getRadialAttr:function(a,e){return{cx:a[0]-a[2]/2+e.cx*a[2],cy:a[1]-a[2]/2+e.cy*a[2],r:e.r*a[2]}},truncate:function(a,e,q,f,b,k,l){var c=this,p=a.rotation,B,g=f?1:0,A=(q||f).length,K=A,d=[],x=function(a){e.firstChild&&e.removeChild(e.firstChild);a&&e.appendChild(m.createTextNode(a))},n=function(k,p){p=p||k;if(void 0===d[p])if(e.getSubStringLength)try{d[p]=b+e.getSubStringLength(0,
f?p+1:p)}catch(X){}else c.getSpanWidth&&(x(l(q||f,k)),d[p]=b+c.getSpanWidth(a,e));return d[p]},t,I;a.rotation=0;t=n(e.textContent.length);if(I=b+t>k){for(;g<=A;)K=Math.ceil((g+A)/2),f&&(B=l(f,K)),t=n(K,B&&B.length-1),g===A?g=A+1:t>k?A=K-1:g=K;0===A?x(""):q&&A===q.length-1||x(B||l(q||f,K))}f&&f.splice(0,K);a.actualWidth=t;a.rotation=p;return I},escapes:{"\x26":"\x26amp;","\x3c":"\x26lt;","\x3e":"\x26gt;","'":"\x26#39;",'"':"\x26quot;"},buildText:function(a){var q=a.element,f=this,k=f.forExport,b=C(a.textStr,
"").toString(),l=-1!==b.indexOf("\x3c"),c=q.childNodes,p,g=d(q,"x"),B=a.styles,A=a.textWidth,x=B&&B.lineHeight,n=B&&B.textOutline,t=B&&"ellipsis"===B.textOverflow,I=B&&"nowrap"===B.whiteSpace,L=B&&B.fontSize,z,H,h=c.length,B=A&&!a.added&&this.box,F=function(a){var b;f.styledMode||(b=/(px|em)$/.test(a&&a.style.fontSize)?a.style.fontSize:L||f.style.fontSize||12);return x?e(x):f.fontMetrics(b,a.getAttribute("style")?a:q).h},M=function(a,e){D(f.escapes,function(q,f){e&&-1!==e.indexOf(q)||(a=a.toString().replace(new RegExp(q,
"g"),f))});return a},w=function(a,e){var q;q=a.indexOf("\x3c");a=a.substring(q,a.indexOf("\x3e")-q);q=a.indexOf(e+"\x3d");if(-1!==q&&(q=q+e.length+1,e=a.charAt(q),'"'===e||"'"===e))return a=a.substring(q+1),a.substring(0,a.indexOf(e))};z=[b,t,I,x,n,L,A].join();if(z!==a.textCache){for(a.textCache=z;h--;)q.removeChild(c[h]);l||n||t||A||-1!==b.indexOf(" ")?(B&&B.appendChild(q),l?(b=f.styledMode?b.replace(/<(b|strong)>/g,'\x3cspan class\x3d"highcharts-strong"\x3e').replace(/<(i|em)>/g,'\x3cspan class\x3d"highcharts-emphasized"\x3e'):
b.replace(/<(b|strong)>/g,'\x3cspan style\x3d"font-weight:bold"\x3e').replace(/<(i|em)>/g,'\x3cspan style\x3d"font-style:italic"\x3e'),b=b.replace(/<a/g,"\x3cspan").replace(/<\/(b|strong|i|em|a)>/g,"\x3c/span\x3e").split(/<br.*?>/g)):b=[b],b=b.filter(function(a){return""!==a}),b.forEach(function(e,b){var l,c=0,B=0;e=e.replace(/^\s+|\s+$/g,"").replace(/<span/g,"|||\x3cspan").replace(/<\/span>/g,"\x3c/span\x3e|||");l=e.split("|||");l.forEach(function(e){if(""!==e||1===l.length){var x={},n=m.createElementNS(f.SVG_NS,
"tspan"),D,C;(D=w(e,"class"))&&d(n,"class",D);if(D=w(e,"style"))D=D.replace(/(;| |^)color([ :])/,"$1fill$2"),d(n,"style",D);(C=w(e,"href"))&&!k&&(d(n,"onclick",'location.href\x3d"'+C+'"'),d(n,"class","highcharts-anchor"),f.styledMode||v(n,{cursor:"pointer"}));e=M(e.replace(/<[a-zA-Z\/](.|\n)*?>/g,"")||" ");if(" "!==e){n.appendChild(m.createTextNode(e));c?x.dx=0:b&&null!==g&&(x.x=g);d(n,x);q.appendChild(n);!c&&H&&(!R&&k&&v(n,{display:"block"}),d(n,"dy",F(n)));if(A){var z=e.replace(/([^\^])-/g,"$1- ").split(" "),
x=!I&&(1<l.length||b||1<z.length);C=0;var h=F(n);if(t)p=f.truncate(a,n,e,void 0,0,Math.max(0,A-parseInt(L||12,10)),function(a,e){return a.substring(0,e)+"\u2026"});else if(x)for(;z.length;)z.length&&!I&&0<C&&(n=m.createElementNS(K,"tspan"),d(n,{dy:h,x:g}),D&&d(n,"style",D),n.appendChild(m.createTextNode(z.join(" ").replace(/- /g,"-"))),q.appendChild(n)),f.truncate(a,n,null,z,0===C?B:0,A,function(a,e){return z.slice(0,e).join(" ").replace(/- /g,"-")}),B=a.actualWidth,C++}c++}}});H=H||q.childNodes.length}),
t&&p&&a.attr("title",M(a.textStr,["\x26lt;","\x26gt;"])),B&&B.removeChild(q),n&&a.applyTextOutline&&a.applyTextOutline(n)):q.appendChild(m.createTextNode(M(b)))}},getContrast:function(a){a=u(a).rgba;a[0]*=1;a[1]*=1.2;a[2]*=.5;return 459<a[0]+a[1]+a[2]?"#000000":"#FFFFFF"},button:function(a,e,q,f,b,l,c,g,A){var B=this.label(a,e,q,A,null,null,null,null,"button"),K=0,n=this.styledMode;B.attr(k({padding:8,r:2},b));if(!n){var m,x,d,I;b=k({fill:"#f7f7f7",stroke:"#cccccc","stroke-width":1,style:{color:"#333333",
cursor:"pointer",fontWeight:"normal"}},b);m=b.style;delete b.style;l=k(b,{fill:"#e6e6e6"},l);x=l.style;delete l.style;c=k(b,{fill:"#e6ebf5",style:{color:"#000000",fontWeight:"bold"}},c);d=c.style;delete c.style;g=k(b,{style:{color:"#cccccc"}},g);I=g.style;delete g.style}E(B.element,t?"mouseover":"mouseenter",function(){3!==K&&B.setState(1)});E(B.element,t?"mouseout":"mouseleave",function(){3!==K&&B.setState(K)});B.setState=function(a){1!==a&&(B.state=K=a);B.removeClass(/highcharts-button-(normal|hover|pressed|disabled)/).addClass("highcharts-button-"+
["normal","hover","pressed","disabled"][a||0]);n||B.attr([b,l,c,g][a||0]).css([m,x,d,I][a||0])};n||B.attr(b).css(p({cursor:"default"},m));return B.on("click",function(a){3!==K&&f.call(B,a)})},crispLine:function(a,e){a[1]===a[4]&&(a[1]=a[4]=Math.round(a[1])-e%2/2);a[2]===a[5]&&(a[2]=a[5]=Math.round(a[2])+e%2/2);return a},path:function(a){var e=this.styledMode?{}:{fill:"none"};f(a)?e.d=a:H(a)&&p(e,a);return this.createElement("path").attr(e)},circle:function(a,e,q){a=H(a)?a:void 0===a?{}:{x:a,y:e,r:q};
e=this.createElement("circle");e.xSetter=e.ySetter=function(a,e,q){q.setAttribute("c"+e,a)};return e.attr(a)},arc:function(a,e,q,b,f,k){H(a)?(b=a,e=b.y,q=b.r,a=b.x):b={innerR:b,start:f,end:k};a=this.symbol("arc",a,e,q,q,b);a.r=q;return a},rect:function(a,e,q,b,f,k){f=H(a)?a.r:f;var l=this.createElement("rect");a=H(a)?a:void 0===a?{}:{x:a,y:e,width:Math.max(q,0),height:Math.max(b,0)};this.styledMode||(void 0!==k&&(a.strokeWidth=k,a=l.crisp(a)),a.fill="none");f&&(a.r=f);l.rSetter=function(a,e,q){d(q,
{rx:a,ry:a})};return l.attr(a)},setSize:function(a,e,q){var b=this.alignedObjects,f=b.length;this.width=a;this.height=e;for(this.boxWrapper.animate({width:a,height:e},{step:function(){this.attr({viewBox:"0 0 "+this.attr("width")+" "+this.attr("height")})},duration:C(q,!0)?void 0:0});f--;)b[f].align()},g:function(a){var e=this.createElement("g");return a?e.attr({"class":"highcharts-"+a}):e},image:function(a,e,q,b,f,k){var l={preserveAspectRatio:"none"},c,g=function(a,e){a.setAttributeNS?a.setAttributeNS("http://www.w3.org/1999/xlink",
"href",e):a.setAttribute("hc-svg-href",e)},A=function(e){g(c.element,a);k.call(c,e)};1<arguments.length&&p(l,{x:e,y:q,width:b,height:f});c=this.createElement("image").attr(l);k?(g(c.element,"data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw\x3d\x3d"),l=new S.Image,E(l,"load",A),l.src=a,l.complete&&A({})):g(c.element,a);return c},symbol:function(a,e,q,b,f,k){var l=this,c,g=/^url\((.*?)\)$/,A=g.test(a),K=!A&&(this.symbols[a]?a:"circle"),B=K&&this.symbols[K],x=n(e)&&B&&B.call(this.symbols,
Math.round(e),Math.round(q),b,f,k),d,t;B?(c=this.path(x),l.styledMode||c.attr("fill","none"),p(c,{symbolName:K,x:e,y:q,width:b,height:f}),k&&p(c,k)):A&&(d=a.match(g)[1],c=this.image(d),c.imgwidth=C(M[d]&&M[d].width,k&&k.width),c.imgheight=C(M[d]&&M[d].height,k&&k.height),t=function(){c.attr({width:c.width,height:c.height})},["width","height"].forEach(function(a){c[a+"Setter"]=function(a,e){var q={},b=this["img"+e],f="width"===e?"translateX":"translateY";this[e]=a;n(b)&&(this.element&&this.element.setAttribute(e,
b),this.alignByTranslate||(q[f]=((this[e]||0)-b)/2,this.attr(q)))}}),n(e)&&c.attr({x:e,y:q}),c.isImg=!0,n(c.imgwidth)&&n(c.imgheight)?t():(c.attr({width:0,height:0}),w("img",{onload:function(){var a=r[l.chartIndex];0===this.width&&(v(this,{position:"absolute",top:"-999em"}),m.body.appendChild(this));M[d]={width:this.width,height:this.height};c.imgwidth=this.width;c.imgheight=this.height;c.element&&t();this.parentNode&&this.parentNode.removeChild(this);l.imgCount--;if(!l.imgCount&&a&&a.onload)a.onload()},
src:d}),this.imgCount++));return c},symbols:{circle:function(a,e,q,b){return this.arc(a+q/2,e+b/2,q/2,b/2,{start:0,end:2*Math.PI,open:!1})},square:function(a,e,q,b){return["M",a,e,"L",a+q,e,a+q,e+b,a,e+b,"Z"]},triangle:function(a,e,q,b){return["M",a+q/2,e,"L",a+q,e+b,a,e+b,"Z"]},"triangle-down":function(a,e,q,b){return["M",a,e,"L",a+q,e,a+q/2,e+b,"Z"]},diamond:function(a,e,q,b){return["M",a+q/2,e,"L",a+q,e+b/2,a+q/2,e+b,a,e+b/2,"Z"]},arc:function(a,e,q,b,f){var k=f.start,c=f.r||q,l=f.r||b||q,p=f.end-
.001;q=f.innerR;b=C(f.open,.001>Math.abs(f.end-f.start-2*Math.PI));var g=Math.cos(k),A=Math.sin(k),K=Math.cos(p),p=Math.sin(p);f=.001>f.end-k-Math.PI?0:1;c=["M",a+c*g,e+l*A,"A",c,l,0,f,1,a+c*K,e+l*p];n(q)&&c.push(b?"M":"L",a+q*K,e+q*p,"A",q,q,0,f,0,a+q*g,e+q*A);c.push(b?"":"Z");return c},callout:function(a,e,q,b,f){var k=Math.min(f&&f.r||0,q,b),c=k+6,l=f&&f.anchorX;f=f&&f.anchorY;var p;p=["M",a+k,e,"L",a+q-k,e,"C",a+q,e,a+q,e,a+q,e+k,"L",a+q,e+b-k,"C",a+q,e+b,a+q,e+b,a+q-k,e+b,"L",a+k,e+b,"C",a,e+
b,a,e+b,a,e+b-k,"L",a,e+k,"C",a,e,a,e,a+k,e];l&&l>q?f>e+c&&f<e+b-c?p.splice(13,3,"L",a+q,f-6,a+q+6,f,a+q,f+6,a+q,e+b-k):p.splice(13,3,"L",a+q,b/2,l,f,a+q,b/2,a+q,e+b-k):l&&0>l?f>e+c&&f<e+b-c?p.splice(33,3,"L",a,f+6,a-6,f,a,f-6,a,e+k):p.splice(33,3,"L",a,b/2,l,f,a,b/2,a,e+k):f&&f>b&&l>a+c&&l<a+q-c?p.splice(23,3,"L",l+6,e+b,l,e+b+6,l-6,e+b,a+k,e+b):f&&0>f&&l>a+c&&l<a+q-c&&p.splice(3,3,"L",l-6,e,l,e-6,l+6,e,q-k,e);return p}},clipRect:function(e,q,b,f){var k=a.uniqueKey(),l=this.createElement("clipPath").attr({id:k}).add(this.defs);
e=this.rect(e,q,b,f,0).add(l);e.id=k;e.clipPath=l;e.count=0;return e},text:function(a,e,q,b){var f={};if(b&&(this.allowHTML||!this.forExport))return this.html(a,e,q);f.x=Math.round(e||0);q&&(f.y=Math.round(q));n(a)&&(f.text=a);a=this.createElement("text").attr(f);b||(a.xSetter=function(a,e,q){var b=q.getElementsByTagName("tspan"),f,k=q.getAttribute(e),l;for(l=0;l<b.length;l++)f=b[l],f.getAttribute(e)===k&&f.setAttribute(e,a);q.setAttribute(e,a)});return a},fontMetrics:function(a,q){a=this.styledMode?
q&&y.prototype.getStyle.call(q,"font-size"):a||q&&q.style&&q.style.fontSize||this.style&&this.style.fontSize;a=/px/.test(a)?e(a):/em/.test(a)?parseFloat(a)*(q?this.fontMetrics(null,q.parentNode).f:16):12;q=24>a?a+3:Math.round(1.2*a);return{h:q,b:Math.round(.8*q),f:a}},rotCorr:function(a,e,q){var b=a;e&&q&&(b=Math.max(b*Math.cos(e*g),4));return{x:-a/3*Math.sin(e*g),y:b}},label:function(e,b,f,l,c,g,A,K,d){var m=this,x=m.styledMode,t=m.g("button"!==d&&"label"),I=t.text=m.text("",0,0,A).attr({zIndex:1}),
D,L,B=0,C=3,z=0,H,h,F,M,R,w={},r,u,S=/^url\((.*?)\)$/.test(l),v=x||S,N=function(){return x?D.strokeWidth()%2/2:(r?parseInt(r,10):0)%2/2},P,T,E;d&&t.addClass("highcharts-"+d);P=function(){var a=I.element.style,e={};L=(void 0===H||void 0===h||R)&&n(I.textStr)&&I.getBBox();t.width=(H||L.width||0)+2*C+z;t.height=(h||L.height||0)+2*C;u=C+Math.min(m.fontMetrics(a&&a.fontSize,I).b,L?L.height:Infinity);v&&(D||(t.box=D=m.symbols[l]||S?m.symbol(l):m.rect(),D.addClass(("button"===d?"":"highcharts-label-box")+
(d?" highcharts-"+d+"-box":"")),D.add(t),a=N(),e.x=a,e.y=(K?-u:0)+a),e.width=Math.round(t.width),e.height=Math.round(t.height),D.attr(p(e,w)),w={})};T=function(){var a=z+C,e;e=K?0:u;n(H)&&L&&("center"===R||"right"===R)&&(a+={center:.5,right:1}[R]*(H-L.width));if(a!==I.x||e!==I.y)I.attr("x",a),I.hasBoxWidthChanged&&(L=I.getBBox(!0),P()),void 0!==e&&I.attr("y",e);I.x=a;I.y=e};E=function(a,e){D?D.attr(a,e):w[a]=e};t.onAdd=function(){I.add(t);t.attr({text:e||0===e?e:"",x:b,y:f});D&&n(c)&&t.attr({anchorX:c,
anchorY:g})};t.widthSetter=function(e){H=a.isNumber(e)?e:null};t.heightSetter=function(a){h=a};t["text-alignSetter"]=function(a){R=a};t.paddingSetter=function(a){n(a)&&a!==C&&(C=t.padding=a,T())};t.paddingLeftSetter=function(a){n(a)&&a!==z&&(z=a,T())};t.alignSetter=function(a){a={left:0,center:.5,right:1}[a];a!==B&&(B=a,L&&t.attr({x:F}))};t.textSetter=function(a){void 0!==a&&I.textSetter(a);P();T()};t["stroke-widthSetter"]=function(a,e){a&&(v=!0);r=this["stroke-width"]=a;E(e,a)};x?t.rSetter=function(a,
e){E(e,a)}:t.strokeSetter=t.fillSetter=t.rSetter=function(a,e){"r"!==e&&("fill"===e&&a&&(v=!0),t[e]=a);E(e,a)};t.anchorXSetter=function(a,e){c=t.anchorX=a;E(e,Math.round(a)-N()-F)};t.anchorYSetter=function(a,e){g=t.anchorY=a;E(e,a-M)};t.xSetter=function(a){t.x=a;B&&(a-=B*((H||L.width)+2*C),t["forceAnimate:x"]=!0);F=Math.round(a);t.attr("translateX",F)};t.ySetter=function(a){M=t.y=Math.round(a);t.attr("translateY",M)};var Q=t.css;A={css:function(a){if(a){var e={};a=k(a);t.textProps.forEach(function(q){void 0!==
a[q]&&(e[q]=a[q],delete a[q])});I.css(e);"width"in e&&P();"fontSize"in e&&(P(),T())}return Q.call(t,a)},getBBox:function(){return{width:L.width+2*C,height:L.height+2*C,x:L.x-C,y:L.y-C}},destroy:function(){q(t.element,"mouseenter");q(t.element,"mouseleave");I&&(I=I.destroy());D&&(D=D.destroy());y.prototype.destroy.call(t);t=m=P=T=E=null}};x||(A.shadow=function(a){a&&(P(),D&&D.shadow(a));return t});return p(t,A)}});a.Renderer=G})(J);(function(a){var y=a.attr,G=a.createElement,E=a.css,h=a.defined,d=
a.extend,r=a.isFirefox,u=a.isMS,v=a.isWebKit,w=a.pick,n=a.pInt,g=a.SVGRenderer,c=a.win,m=a.wrap;d(a.SVGElement.prototype,{htmlCss:function(a){var b="SPAN"===this.element.tagName&&a&&"width"in a,l=w(b&&a.width,void 0),f;b&&(delete a.width,this.textWidth=l,f=!0);a&&"ellipsis"===a.textOverflow&&(a.whiteSpace="nowrap",a.overflow="hidden");this.styles=d(this.styles,a);E(this.element,a);f&&this.htmlUpdateTransform();return this},htmlGetBBox:function(){var a=this.element;return{x:a.offsetLeft,y:a.offsetTop,
width:a.offsetWidth,height:a.offsetHeight}},htmlUpdateTransform:function(){if(this.added){var a=this.renderer,b=this.element,l=this.translateX||0,f=this.translateY||0,c=this.x||0,g=this.y||0,m=this.textAlign||"left",d={left:0,center:.5,right:1}[m],z=this.styles,k=z&&z.whiteSpace;E(b,{marginLeft:l,marginTop:f});!a.styledMode&&this.shadows&&this.shadows.forEach(function(a){E(a,{marginLeft:l+1,marginTop:f+1})});this.inverted&&b.childNodes.forEach(function(e){a.invertChild(e,b)});if("SPAN"===b.tagName){var z=
this.rotation,A=this.textWidth&&n(this.textWidth),D=[z,m,b.innerHTML,this.textWidth,this.textAlign].join(),C;(C=A!==this.oldTextWidth)&&!(C=A>this.oldTextWidth)&&((C=this.textPxLength)||(E(b,{width:"",whiteSpace:k||"nowrap"}),C=b.offsetWidth),C=C>A);C&&(/[ \-]/.test(b.textContent||b.innerText)||"ellipsis"===b.style.textOverflow)?(E(b,{width:A+"px",display:"block",whiteSpace:k||"normal"}),this.oldTextWidth=A,this.hasBoxWidthChanged=!0):this.hasBoxWidthChanged=!1;D!==this.cTT&&(k=a.fontMetrics(b.style.fontSize,
b).b,!h(z)||z===(this.oldRotation||0)&&m===this.oldAlign||this.setSpanRotation(z,d,k),this.getSpanCorrection(!h(z)&&this.textPxLength||b.offsetWidth,k,d,z,m));E(b,{left:c+(this.xCorr||0)+"px",top:g+(this.yCorr||0)+"px"});this.cTT=D;this.oldRotation=z;this.oldAlign=m}}else this.alignOnAdd=!0},setSpanRotation:function(a,b,l){var f={},c=this.renderer.getTransformKey();f[c]=f.transform="rotate("+a+"deg)";f[c+(r?"Origin":"-origin")]=f.transformOrigin=100*b+"% "+l+"px";E(this.element,f)},getSpanCorrection:function(a,
b,l){this.xCorr=-a*l;this.yCorr=-b}});d(g.prototype,{getTransformKey:function(){return u&&!/Edge/.test(c.navigator.userAgent)?"-ms-transform":v?"-webkit-transform":r?"MozTransform":c.opera?"-o-transform":""},html:function(c,b,l){var f=this.createElement("span"),g=f.element,p=f.renderer,n=p.isSVG,h=function(a,b){["opacity","visibility"].forEach(function(f){m(a,f+"Setter",function(a,e,q,f){a.call(this,e,q,f);b[q]=e})});a.addedSetters=!0},z=a.charts[p.chartIndex],z=z&&z.styledMode;f.textSetter=function(a){a!==
g.innerHTML&&delete this.bBox;this.textStr=a;g.innerHTML=w(a,"");f.doTransform=!0};n&&h(f,f.element.style);f.xSetter=f.ySetter=f.alignSetter=f.rotationSetter=function(a,b){"align"===b&&(b="textAlign");f[b]=a;f.doTransform=!0};f.afterSetters=function(){this.doTransform&&(this.htmlUpdateTransform(),this.doTransform=!1)};f.attr({text:c,x:Math.round(b),y:Math.round(l)}).css({position:"absolute"});z||f.css({fontFamily:this.style.fontFamily,fontSize:this.style.fontSize});g.style.whiteSpace="nowrap";f.css=
f.htmlCss;n&&(f.add=function(a){var b,l=p.box.parentNode,c=[];if(this.parentGroup=a){if(b=a.div,!b){for(;a;)c.push(a),a=a.parentGroup;c.reverse().forEach(function(a){function e(e,q){a[q]=e;"translateX"===q?k.left=e+"px":k.top=e+"px";a.doTransform=!0}var k,g=y(a.element,"class");g&&(g={className:g});b=a.div=a.div||G("div",g,{position:"absolute",left:(a.translateX||0)+"px",top:(a.translateY||0)+"px",display:a.display,opacity:a.opacity,pointerEvents:a.styles&&a.styles.pointerEvents},b||l);k=b.style;
d(a,{classSetter:function(a){return function(e){this.element.setAttribute("class",e);a.className=e}}(b),on:function(){c[0].div&&f.on.apply({element:c[0].div},arguments);return a},translateXSetter:e,translateYSetter:e});a.addedSetters||h(a,k)})}}else b=l;b.appendChild(g);f.added=!0;f.alignOnAdd&&f.htmlUpdateTransform();return f});return f}})})(J);(function(a){var y=a.defined,G=a.extend,E=a.merge,h=a.pick,d=a.timeUnits,r=a.win;a.Time=function(a){this.update(a,!1)};a.Time.prototype={defaultOptions:{},
update:function(a){var d=h(a&&a.useUTC,!0),w=this;this.options=a=E(!0,this.options||{},a);this.Date=a.Date||r.Date;this.timezoneOffset=(this.useUTC=d)&&a.timezoneOffset;this.getTimezoneOffset=this.timezoneOffsetFunction();(this.variableTimezone=!(d&&!a.getTimezoneOffset&&!a.timezone))||this.timezoneOffset?(this.get=function(a,g){var c=g.getTime(),m=c-w.getTimezoneOffset(g);g.setTime(m);a=g["getUTC"+a]();g.setTime(c);return a},this.set=function(a,g,c){var m;if("Milliseconds"===a||"Seconds"===a||"Minutes"===
a&&0===g.getTimezoneOffset()%60)g["set"+a](c);else m=w.getTimezoneOffset(g),m=g.getTime()-m,g.setTime(m),g["setUTC"+a](c),a=w.getTimezoneOffset(g),m=g.getTime()+a,g.setTime(m)}):d?(this.get=function(a,g){return g["getUTC"+a]()},this.set=function(a,g,c){return g["setUTC"+a](c)}):(this.get=function(a,g){return g["get"+a]()},this.set=function(a,g,c){return g["set"+a](c)})},makeTime:function(d,r,w,n,g,c){var m,p,b;this.useUTC?(m=this.Date.UTC.apply(0,arguments),p=this.getTimezoneOffset(m),m+=p,b=this.getTimezoneOffset(m),
p!==b?m+=b-p:p-36E5!==this.getTimezoneOffset(m-36E5)||a.isSafari||(m-=36E5)):m=(new this.Date(d,r,h(w,1),h(n,0),h(g,0),h(c,0))).getTime();return m},timezoneOffsetFunction:function(){var d=this,h=this.options,w=r.moment;if(!this.useUTC)return function(a){return 6E4*(new Date(a)).getTimezoneOffset()};if(h.timezone){if(w)return function(a){return 6E4*-w.tz(a,h.timezone).utcOffset()};a.error(25)}return this.useUTC&&h.getTimezoneOffset?function(a){return 6E4*h.getTimezoneOffset(a)}:function(){return 6E4*
(d.timezoneOffset||0)}},dateFormat:function(d,h,w){if(!a.defined(h)||isNaN(h))return a.defaultOptions.lang.invalidDate||"";d=a.pick(d,"%Y-%m-%d %H:%M:%S");var n=this,g=new this.Date(h),c=this.get("Hours",g),m=this.get("Day",g),p=this.get("Date",g),b=this.get("Month",g),l=this.get("FullYear",g),f=a.defaultOptions.lang,x=f.weekdays,t=f.shortWeekdays,H=a.pad,g=a.extend({a:t?t[m]:x[m].substr(0,3),A:x[m],d:H(p),e:H(p,2," "),w:m,b:f.shortMonths[b],B:f.months[b],m:H(b+1),o:b+1,y:l.toString().substr(2,2),
Y:l,H:H(c),k:c,I:H(c%12||12),l:c%12||12,M:H(n.get("Minutes",g)),p:12>c?"AM":"PM",P:12>c?"am":"pm",S:H(g.getSeconds()),L:H(Math.floor(h%1E3),3)},a.dateFormats);a.objectEach(g,function(a,b){for(;-1!==d.indexOf("%"+b);)d=d.replace("%"+b,"function"===typeof a?a.call(n,h):a)});return w?d.substr(0,1).toUpperCase()+d.substr(1):d},resolveDTLFormat:function(d){return a.isObject(d,!0)?d:(d=a.splat(d),{main:d[0],from:d[1],to:d[2]})},getTimeTicks:function(a,r,w,n){var g=this,c=[],m,p={},b;m=new g.Date(r);var l=
a.unitRange,f=a.count||1,x;n=h(n,1);if(y(r)){g.set("Milliseconds",m,l>=d.second?0:f*Math.floor(g.get("Milliseconds",m)/f));l>=d.second&&g.set("Seconds",m,l>=d.minute?0:f*Math.floor(g.get("Seconds",m)/f));l>=d.minute&&g.set("Minutes",m,l>=d.hour?0:f*Math.floor(g.get("Minutes",m)/f));l>=d.hour&&g.set("Hours",m,l>=d.day?0:f*Math.floor(g.get("Hours",m)/f));l>=d.day&&g.set("Date",m,l>=d.month?1:Math.max(1,f*Math.floor(g.get("Date",m)/f)));l>=d.month&&(g.set("Month",m,l>=d.year?0:f*Math.floor(g.get("Month",
m)/f)),b=g.get("FullYear",m));l>=d.year&&g.set("FullYear",m,b-b%f);l===d.week&&(b=g.get("Day",m),g.set("Date",m,g.get("Date",m)-b+n+(b<n?-7:0)));b=g.get("FullYear",m);n=g.get("Month",m);var t=g.get("Date",m),H=g.get("Hours",m);r=m.getTime();g.variableTimezone&&(x=w-r>4*d.month||g.getTimezoneOffset(r)!==g.getTimezoneOffset(w));r=m.getTime();for(m=1;r<w;)c.push(r),r=l===d.year?g.makeTime(b+m*f,0):l===d.month?g.makeTime(b,n+m*f):!x||l!==d.day&&l!==d.week?x&&l===d.hour&&1<f?g.makeTime(b,n,t,H+m*f):r+
l*f:g.makeTime(b,n,t+m*f*(l===d.day?1:7)),m++;c.push(r);l<=d.hour&&1E4>c.length&&c.forEach(function(a){0===a%18E5&&"000000000"===g.dateFormat("%H%M%S%L",a)&&(p[a]="day")})}c.info=G(a,{higherRanks:p,totalRange:l*f});return c}}})(J);(function(a){var y=a.color,G=a.merge;a.defaultOptions={colors:"#7cb5ec #434348 #90ed7d #f7a35c #8085e9 #f15c80 #e4d354 #2b908f #f45b5b #91e8e1".split(" "),symbols:["circle","diamond","square","triangle","triangle-down"],lang:{loading:"Loading...",months:"January February March April May June July August September October November December".split(" "),
shortMonths:"Jan Feb Mar Apr May Jun Jul Aug Sep Oct Nov Dec".split(" "),weekdays:"Sunday Monday Tuesday Wednesday Thursday Friday Saturday".split(" "),decimalPoint:".",numericSymbols:"kMGTPE".split(""),resetZoom:"Reset zoom",resetZoomTitle:"Reset zoom level 1:1",thousandsSep:" "},global:{},time:a.Time.prototype.defaultOptions,chart:{styledMode:!1,borderRadius:0,colorCount:10,defaultSeriesType:"line",ignoreHiddenSeries:!0,spacing:[10,10,15,10],resetZoomButton:{theme:{zIndex:6},position:{align:"right",
x:-10,y:10}},width:null,height:null,borderColor:"#335cad",backgroundColor:"#ffffff",plotBorderColor:"#cccccc"},title:{text:"Chart title",align:"center",margin:15,widthAdjust:-44},subtitle:{text:"",align:"center",widthAdjust:-44},plotOptions:{},labels:{style:{position:"absolute",color:"#333333"}},legend:{enabled:!0,align:"center",alignColumns:!0,layout:"horizontal",labelFormatter:function(){return this.name},borderColor:"#999999",borderRadius:0,navigation:{activeColor:"#003399",inactiveColor:"#cccccc"},
itemStyle:{color:"#333333",cursor:"pointer",fontSize:"12px",fontWeight:"bold",textOverflow:"ellipsis"},itemHoverStyle:{color:"#000000"},itemHiddenStyle:{color:"#cccccc"},shadow:!1,itemCheckboxStyle:{position:"absolute",width:"13px",height:"13px"},squareSymbol:!0,symbolPadding:5,verticalAlign:"bottom",x:0,y:0,title:{style:{fontWeight:"bold"}}},loading:{labelStyle:{fontWeight:"bold",position:"relative",top:"45%"},style:{position:"absolute",backgroundColor:"#ffffff",opacity:.5,textAlign:"center"}},tooltip:{enabled:!0,
animation:a.svg,borderRadius:3,dateTimeLabelFormats:{millisecond:"%A, %b %e, %H:%M:%S.%L",second:"%A, %b %e, %H:%M:%S",minute:"%A, %b %e, %H:%M",hour:"%A, %b %e, %H:%M",day:"%A, %b %e, %Y",week:"Week from %A, %b %e, %Y",month:"%B %Y",year:"%Y"},footerFormat:"",padding:8,snap:a.isTouchDevice?25:10,headerFormat:'\x3cspan style\x3d"font-size: 10px"\x3e{point.key}\x3c/span\x3e\x3cbr/\x3e',pointFormat:'\x3cspan style\x3d"color:{point.color}"\x3e\u25cf\x3c/span\x3e {series.name}: \x3cb\x3e{point.y}\x3c/b\x3e\x3cbr/\x3e',
backgroundColor:y("#f7f7f7").setOpacity(.85).get(),borderWidth:1,shadow:!0,style:{color:"#333333",cursor:"default",fontSize:"12px",pointerEvents:"none",whiteSpace:"nowrap"}},credits:{enabled:!0,href:"https://www.highcharts.com?credits",position:{align:"right",x:-10,verticalAlign:"bottom",y:-5},style:{cursor:"pointer",color:"#999999",fontSize:"9px"},text:"Highcharts.com"}};a.setOptions=function(y){a.defaultOptions=G(!0,a.defaultOptions,y);a.time.update(G(a.defaultOptions.global,a.defaultOptions.time),
!1);return a.defaultOptions};a.getOptions=function(){return a.defaultOptions};a.defaultPlotOptions=a.defaultOptions.plotOptions;a.time=new a.Time(G(a.defaultOptions.global,a.defaultOptions.time));a.dateFormat=function(y,h,d){return a.time.dateFormat(y,h,d)}})(J);(function(a){var y=a.correctFloat,G=a.defined,E=a.destroyObjectProperties,h=a.fireEvent,d=a.isNumber,r=a.merge,u=a.pick,v=a.deg2rad;a.Tick=function(a,d,g,c,m){this.axis=a;this.pos=d;this.type=g||"";this.isNewLabel=this.isNew=!0;this.parameters=
m||{};this.tickmarkOffset=this.parameters.tickmarkOffset;this.options=this.parameters.options;g||c||this.addLabel()};a.Tick.prototype={addLabel:function(){var d=this,n=d.axis,g=n.options,c=n.chart,m=n.categories,p=n.names,b=d.pos,l=u(d.options&&d.options.labels,g.labels),f=n.tickPositions,x=b===f[0],t=b===f[f.length-1],m=this.parameters.category||(m?u(m[b],p[b],b):b),h=d.label,f=f.info,F,z,k,A;n.isDatetimeAxis&&f&&(z=c.time.resolveDTLFormat(g.dateTimeLabelFormats[!g.grid&&f.higherRanks[b]||f.unitName]),
F=z.main);d.isFirst=x;d.isLast=t;d.formatCtx={axis:n,chart:c,isFirst:x,isLast:t,dateTimeLabelFormat:F,tickPositionInfo:f,value:n.isLog?y(n.lin2log(m)):m,pos:b};g=n.labelFormatter.call(d.formatCtx,this.formatCtx);if(A=z&&z.list)d.shortenLabel=function(){for(k=0;k<A.length;k++)if(h.attr({text:n.labelFormatter.call(a.extend(d.formatCtx,{dateTimeLabelFormat:A[k]}))}),h.getBBox().width<n.getSlotWidth(d)-2*u(l.padding,5))return;h.attr({text:""})};if(G(h))h&&h.textStr!==g&&(!h.textWidth||l.style&&l.style.width||
h.styles.width||h.css({width:null}),h.attr({text:g}));else{if(d.label=h=G(g)&&l.enabled?c.renderer.text(g,0,0,l.useHTML).add(n.labelGroup):null)c.styledMode||h.css(r(l.style)),h.textPxLength=h.getBBox().width;d.rotation=0}},getLabelSize:function(){return this.label?this.label.getBBox()[this.axis.horiz?"height":"width"]:0},handleOverflow:function(a){var d=this.axis,g=d.options.labels,c=a.x,m=d.chart.chartWidth,p=d.chart.spacing,b=u(d.labelLeft,Math.min(d.pos,p[3])),p=u(d.labelRight,Math.max(d.isRadial?
0:d.pos+d.len,m-p[1])),l=this.label,f=this.rotation,x={left:0,center:.5,right:1}[d.labelAlign||l.attr("align")],t=l.getBBox().width,h=d.getSlotWidth(this),F=h,z=1,k,A={};if(f||"justify"!==u(g.overflow,"justify"))0>f&&c-x*t<b?k=Math.round(c/Math.cos(f*v)-b):0<f&&c+x*t>p&&(k=Math.round((m-c)/Math.cos(f*v)));else if(m=c+(1-x)*t,c-x*t<b?F=a.x+F*(1-x)-b:m>p&&(F=p-a.x+F*x,z=-1),F=Math.min(h,F),F<h&&"center"===d.labelAlign&&(a.x+=z*(h-F-x*(h-Math.min(t,F)))),t>F||d.autoRotation&&(l.styles||{}).width)k=F;
k&&(this.shortenLabel?this.shortenLabel():(A.width=Math.floor(k),(g.style||{}).textOverflow||(A.textOverflow="ellipsis"),l.css(A)))},getPosition:function(d,n,g,c){var m=this.axis,p=m.chart,b=c&&p.oldChartHeight||p.chartHeight;d={x:d?a.correctFloat(m.translate(n+g,null,null,c)+m.transB):m.left+m.offset+(m.opposite?(c&&p.oldChartWidth||p.chartWidth)-m.right-m.left:0),y:d?b-m.bottom+m.offset-(m.opposite?m.height:0):a.correctFloat(b-m.translate(n+g,null,null,c)-m.transB)};h(this,"afterGetPosition",{pos:d});
return d},getLabelPosition:function(a,d,g,c,m,p,b,l){var f=this.axis,x=f.transA,t=f.reversed,n=f.staggerLines,F=f.tickRotCorr||{x:0,y:0},z=m.y,k=c||f.reserveSpaceDefault?0:-f.labelOffset*("center"===f.labelAlign?.5:1),A={};G(z)||(z=0===f.side?g.rotation?-8:-g.getBBox().height:2===f.side?F.y+8:Math.cos(g.rotation*v)*(F.y-g.getBBox(!1,0).height/2));a=a+m.x+k+F.x-(p&&c?p*x*(t?-1:1):0);d=d+z-(p&&!c?p*x*(t?1:-1):0);n&&(g=b/(l||1)%n,f.opposite&&(g=n-g-1),d+=f.labelOffset/n*g);A.x=a;A.y=Math.round(d);h(this,
"afterGetLabelPosition",{pos:A});return A},getMarkPath:function(a,d,g,c,m,p){return p.crispLine(["M",a,d,"L",a+(m?0:-g),d+(m?g:0)],c)},renderGridLine:function(a,d,g){var c=this.axis,m=c.options,p=this.gridLine,b={},l=this.pos,f=this.type,x=u(this.tickmarkOffset,c.tickmarkOffset),t=c.chart.renderer,n=f?f+"Grid":"grid",h=m[n+"LineWidth"],z=m[n+"LineColor"],m=m[n+"LineDashStyle"];p||(c.chart.styledMode||(b.stroke=z,b["stroke-width"]=h,m&&(b.dashstyle=m)),f||(b.zIndex=1),a&&(d=0),this.gridLine=p=t.path().attr(b).addClass("highcharts-"+
(f?f+"-":"")+"grid-line").add(c.gridGroup));if(p&&(g=c.getPlotLinePath(l+x,p.strokeWidth()*g,a,"pass")))p[a||this.isNew?"attr":"animate"]({d:g,opacity:d})},renderMark:function(a,d,g){var c=this.axis,m=c.options,p=c.chart.renderer,b=this.type,l=b?b+"Tick":"tick",f=c.tickSize(l),x=this.mark,t=!x,n=a.x;a=a.y;var h=u(m[l+"Width"],!b&&c.isXAxis?1:0),m=m[l+"Color"];f&&(c.opposite&&(f[0]=-f[0]),t&&(this.mark=x=p.path().addClass("highcharts-"+(b?b+"-":"")+"tick").add(c.axisGroup),c.chart.styledMode||x.attr({stroke:m,
"stroke-width":h})),x[t?"attr":"animate"]({d:this.getMarkPath(n,a,f[0],x.strokeWidth()*g,c.horiz,p),opacity:d}))},renderLabel:function(a,n,g,c){var m=this.axis,p=m.horiz,b=m.options,l=this.label,f=b.labels,x=f.step,m=u(this.tickmarkOffset,m.tickmarkOffset),t=!0,h=a.x;a=a.y;l&&d(h)&&(l.xy=a=this.getLabelPosition(h,a,l,p,f,m,c,x),this.isFirst&&!this.isLast&&!u(b.showFirstLabel,1)||this.isLast&&!this.isFirst&&!u(b.showLastLabel,1)?t=!1:!p||f.step||f.rotation||n||0===g||this.handleOverflow(a),x&&c%x&&
(t=!1),t&&d(a.y)?(a.opacity=g,l[this.isNewLabel?"attr":"animate"](a),this.isNewLabel=!1):(l.attr("y",-9999),this.isNewLabel=!0))},render:function(d,n,g){var c=this.axis,m=c.horiz,p=this.pos,b=u(this.tickmarkOffset,c.tickmarkOffset),p=this.getPosition(m,p,b,n),b=p.x,l=p.y,c=m&&b===c.pos+c.len||!m&&l===c.pos?-1:1;g=u(g,1);this.isActive=!0;this.renderGridLine(n,g,c);this.renderMark(p,g,c);this.renderLabel(p,n,g,d);this.isNew=!1;a.fireEvent(this,"afterRender")},destroy:function(){E(this,this.axis)}}})(J);
var V=function(a){var y=a.addEvent,G=a.animObject,E=a.arrayMax,h=a.arrayMin,d=a.color,r=a.correctFloat,u=a.defaultOptions,v=a.defined,w=a.deg2rad,n=a.destroyObjectProperties,g=a.extend,c=a.fireEvent,m=a.format,p=a.getMagnitude,b=a.isArray,l=a.isNumber,f=a.isString,x=a.merge,t=a.normalizeTickInterval,H=a.objectEach,F=a.pick,z=a.removeEvent,k=a.splat,A=a.syncTimeout,D=a.Tick,C=function(){this.init.apply(this,arguments)};a.extend(C.prototype,{defaultOptions:{dateTimeLabelFormats:{millisecond:{main:"%H:%M:%S.%L",
range:!1},second:{main:"%H:%M:%S",range:!1},minute:{main:"%H:%M",range:!1},hour:{main:"%H:%M",range:!1},day:{main:"%e. %b"},week:{main:"%e. %b"},month:{main:"%b '%y"},year:{main:"%Y"}},endOnTick:!1,labels:{enabled:!0,indentation:10,x:0,style:{color:"#666666",cursor:"default",fontSize:"11px"}},maxPadding:.01,minorTickLength:2,minorTickPosition:"outside",minPadding:.01,startOfWeek:1,startOnTick:!1,tickLength:10,tickPixelInterval:100,tickmarkPlacement:"between",tickPosition:"outside",title:{align:"middle",
style:{color:"#666666"}},type:"linear",minorGridLineColor:"#f2f2f2",minorGridLineWidth:1,minorTickColor:"#999999",lineColor:"#ccd6eb",lineWidth:1,gridLineColor:"#e6e6e6",tickColor:"#ccd6eb"},defaultYAxisOptions:{endOnTick:!0,maxPadding:.05,minPadding:.05,tickPixelInterval:72,showLastLabel:!0,labels:{x:-8},startOnTick:!0,title:{rotation:270,text:"Values"},stackLabels:{allowOverlap:!1,enabled:!1,formatter:function(){return a.numberFormat(this.total,-1)},style:{color:"#000000",fontSize:"11px",fontWeight:"bold",
textOutline:"1px contrast"}},gridLineWidth:1,lineWidth:0},defaultLeftAxisOptions:{labels:{x:-15},title:{rotation:270}},defaultRightAxisOptions:{labels:{x:15},title:{rotation:90}},defaultBottomAxisOptions:{labels:{autoRotation:[-45],x:0},title:{rotation:0}},defaultTopAxisOptions:{labels:{autoRotation:[-45],x:0},title:{rotation:0}},init:function(a,q){var e=q.isX,b=this;b.chart=a;b.horiz=a.inverted&&!b.isZAxis?!e:e;b.isXAxis=e;b.coll=b.coll||(e?"xAxis":"yAxis");c(this,"init",{userOptions:q});b.opposite=
q.opposite;b.side=q.side||(b.horiz?b.opposite?0:2:b.opposite?1:3);b.setOptions(q);var f=this.options,l=f.type;b.labelFormatter=f.labels.formatter||b.defaultLabelFormatter;b.userOptions=q;b.minPixelPadding=0;b.reversed=f.reversed;b.visible=!1!==f.visible;b.zoomEnabled=!1!==f.zoomEnabled;b.hasNames="category"===l||!0===f.categories;b.categories=f.categories||b.hasNames;b.names||(b.names=[],b.names.keys={});b.plotLinesAndBandsGroups={};b.isLog="logarithmic"===l;b.isDatetimeAxis="datetime"===l;b.positiveValuesOnly=
b.isLog&&!b.allowNegativeLog;b.isLinked=v(f.linkedTo);b.ticks={};b.labelEdge=[];b.minorTicks={};b.plotLinesAndBands=[];b.alternateBands={};b.len=0;b.minRange=b.userMinRange=f.minRange||f.maxZoom;b.range=f.range;b.offset=f.offset||0;b.stacks={};b.oldStacks={};b.stacksTouched=0;b.max=null;b.min=null;b.crosshair=F(f.crosshair,k(a.options.tooltip.crosshairs)[e?0:1],!1);q=b.options.events;-1===a.axes.indexOf(b)&&(e?a.axes.splice(a.xAxis.length,0,b):a.axes.push(b),a[b.coll].push(b));b.series=b.series||
[];a.inverted&&!b.isZAxis&&e&&void 0===b.reversed&&(b.reversed=!0);H(q,function(a,e){y(b,e,a)});b.lin2log=f.linearToLogConverter||b.lin2log;b.isLog&&(b.val2lin=b.log2lin,b.lin2val=b.lin2log);c(this,"afterInit")},setOptions:function(a){this.options=x(this.defaultOptions,"yAxis"===this.coll&&this.defaultYAxisOptions,[this.defaultTopAxisOptions,this.defaultRightAxisOptions,this.defaultBottomAxisOptions,this.defaultLeftAxisOptions][this.side],x(u[this.coll],a));c(this,"afterSetOptions",{userOptions:a})},
defaultLabelFormatter:function(){var e=this.axis,q=this.value,b=e.chart.time,f=e.categories,l=this.dateTimeLabelFormat,c=u.lang,k=c.numericSymbols,c=c.numericSymbolMagnitude||1E3,g=k&&k.length,d,p=e.options.labels.format,e=e.isLog?Math.abs(q):e.tickInterval;if(p)d=m(p,this,b);else if(f)d=q;else if(l)d=b.dateFormat(l,q);else if(g&&1E3<=e)for(;g--&&void 0===d;)b=Math.pow(c,g+1),e>=b&&0===10*q%b&&null!==k[g]&&0!==q&&(d=a.numberFormat(q/b,-1)+k[g]);void 0===d&&(d=1E4<=Math.abs(q)?a.numberFormat(q,-1):
a.numberFormat(q,-1,void 0,""));return d},getSeriesExtremes:function(){var a=this,q=a.chart;c(this,"getSeriesExtremes",null,function(){a.hasVisibleSeries=!1;a.dataMin=a.dataMax=a.threshold=null;a.softThreshold=!a.isXAxis;a.buildStacks&&a.buildStacks();a.series.forEach(function(e){if(e.visible||!q.options.chart.ignoreHiddenSeries){var b=e.options,f=b.threshold,c;a.hasVisibleSeries=!0;a.positiveValuesOnly&&0>=f&&(f=null);if(a.isXAxis)b=e.xData,b.length&&(e=h(b),c=E(b),l(e)||e instanceof Date||(b=b.filter(l),
e=h(b),c=E(b)),b.length&&(a.dataMin=Math.min(F(a.dataMin,b[0],e),e),a.dataMax=Math.max(F(a.dataMax,b[0],c),c)));else if(e.getExtremes(),c=e.dataMax,e=e.dataMin,v(e)&&v(c)&&(a.dataMin=Math.min(F(a.dataMin,e),e),a.dataMax=Math.max(F(a.dataMax,c),c)),v(f)&&(a.threshold=f),!b.softThreshold||a.positiveValuesOnly)a.softThreshold=!1}})});c(this,"afterGetSeriesExtremes")},translate:function(a,q,b,f,c,k){var e=this.linkedParent||this,d=1,g=0,p=f?e.oldTransA:e.transA;f=f?e.oldMin:e.min;var A=e.minPixelPadding;
c=(e.isOrdinal||e.isBroken||e.isLog&&c)&&e.lin2val;p||(p=e.transA);b&&(d*=-1,g=e.len);e.reversed&&(d*=-1,g-=d*(e.sector||e.len));q?(a=(a*d+g-A)/p+f,c&&(a=e.lin2val(a))):(c&&(a=e.val2lin(a)),a=l(f)?d*(a-f)*p+g+d*A+(l(k)?p*k:0):void 0);return a},toPixels:function(a,q){return this.translate(a,!1,!this.horiz,null,!0)+(q?0:this.pos)},toValue:function(a,q){return this.translate(a-(q?0:this.pos),!0,!this.horiz,null,!0)},getPlotLinePath:function(a,q,b,f,c){var e=this.chart,k=this.left,d=this.top,g,p,A=b&&
e.oldChartHeight||e.chartHeight,t=b&&e.oldChartWidth||e.chartWidth,m;g=this.transB;var x=function(a,e,q){if("pass"!==f&&a<e||a>q)f?a=Math.min(Math.max(e,a),q):m=!0;return a};c=F(c,this.translate(a,null,null,b));c=Math.min(Math.max(-1E5,c),1E5);a=b=Math.round(c+g);g=p=Math.round(A-c-g);l(c)?this.horiz?(g=d,p=A-this.bottom,a=b=x(a,k,k+this.width)):(a=k,b=t-this.right,g=p=x(g,d,d+this.height)):(m=!0,f=!1);return m&&!f?null:e.renderer.crispLine(["M",a,g,"L",b,p],q||1)},getLinearTickPositions:function(a,
q,b){var e,f=r(Math.floor(q/a)*a);b=r(Math.ceil(b/a)*a);var c=[],k;r(f+a)===f&&(k=20);if(this.single)return[q];for(q=f;q<=b;){c.push(q);q=r(q+a,k);if(q===e)break;e=q}return c},getMinorTickInterval:function(){var a=this.options;return!0===a.minorTicks?F(a.minorTickInterval,"auto"):!1===a.minorTicks?null:a.minorTickInterval},getMinorTickPositions:function(){var a=this,q=a.options,b=a.tickPositions,f=a.minorTickInterval,c=[],k=a.pointRangePadding||0,l=a.min-k,k=a.max+k,g=k-l;if(g&&g/f<a.len/3)if(a.isLog)this.paddedTicks.forEach(function(e,
q,b){q&&c.push.apply(c,a.getLogTickPositions(f,b[q-1],b[q],!0))});else if(a.isDatetimeAxis&&"auto"===this.getMinorTickInterval())c=c.concat(a.getTimeTicks(a.normalizeTimeTickInterval(f),l,k,q.startOfWeek));else for(q=l+(b[0]-l)%f;q<=k&&q!==c[0];q+=f)c.push(q);0!==c.length&&a.trimTicks(c);return c},adjustForMinRange:function(){var a=this.options,q=this.min,b=this.max,f,c,k,l,g,d,p,A;this.isXAxis&&void 0===this.minRange&&!this.isLog&&(v(a.min)||v(a.max)?this.minRange=null:(this.series.forEach(function(a){d=
a.xData;for(l=p=a.xIncrement?1:d.length-1;0<l;l--)if(g=d[l]-d[l-1],void 0===k||g<k)k=g}),this.minRange=Math.min(5*k,this.dataMax-this.dataMin)));b-q<this.minRange&&(c=this.dataMax-this.dataMin>=this.minRange,A=this.minRange,f=(A-b+q)/2,f=[q-f,F(a.min,q-f)],c&&(f[2]=this.isLog?this.log2lin(this.dataMin):this.dataMin),q=E(f),b=[q+A,F(a.max,q+A)],c&&(b[2]=this.isLog?this.log2lin(this.dataMax):this.dataMax),b=h(b),b-q<A&&(f[0]=b-A,f[1]=F(a.min,b-A),q=E(f)));this.min=q;this.max=b},getClosest:function(){var a;
this.categories?a=1:this.series.forEach(function(e){var q=e.closestPointRange,b=e.visible||!e.chart.options.chart.ignoreHiddenSeries;!e.noSharedTooltip&&v(q)&&b&&(a=v(a)?Math.min(a,q):q)});return a},nameToX:function(a){var e=b(this.categories),f=e?this.categories:this.names,c=a.options.x,k;a.series.requireSorting=!1;v(c)||(c=!1===this.options.uniqueNames?a.series.autoIncrement():e?f.indexOf(a.name):F(f.keys[a.name],-1));-1===c?e||(k=f.length):k=c;void 0!==k&&(this.names[k]=a.name,this.names.keys[a.name]=
k);return k},updateNames:function(){var a=this,q=this.names;0<q.length&&(Object.keys(q.keys).forEach(function(a){delete q.keys[a]}),q.length=0,this.minRange=this.userMinRange,(this.series||[]).forEach(function(e){e.xIncrement=null;if(!e.points||e.isDirtyData)a.max=Math.max(a.max,e.xData.length-1),e.processData(),e.generatePoints();e.data.forEach(function(q,b){var f;q&&q.options&&void 0!==q.name&&(f=a.nameToX(q),void 0!==f&&f!==q.x&&(q.x=f,e.xData[b]=f))})}))},setAxisTranslation:function(a){var e=
this,b=e.max-e.min,k=e.axisPointRange||0,l,g=0,d=0,p=e.linkedParent,A=!!e.categories,t=e.transA,m=e.isXAxis;if(m||A||k)l=e.getClosest(),p?(g=p.minPointOffset,d=p.pointRangePadding):e.series.forEach(function(a){var b=A?1:m?F(a.options.pointRange,l,0):e.axisPointRange||0;a=a.options.pointPlacement;k=Math.max(k,b);e.single||(g=Math.max(g,f(a)?0:b/2),d=Math.max(d,"on"===a?0:b))}),p=e.ordinalSlope&&l?e.ordinalSlope/l:1,e.minPointOffset=g*=p,e.pointRangePadding=d*=p,e.pointRange=Math.min(k,b),m&&(e.closestPointRange=
l);a&&(e.oldTransA=t);e.translationSlope=e.transA=t=e.staticScale||e.len/(b+d||1);e.transB=e.horiz?e.left:e.bottom;e.minPixelPadding=t*g;c(this,"afterSetAxisTranslation")},minFromRange:function(){return this.max-this.range},setTickInterval:function(e){var b=this,f=b.chart,k=b.options,g=b.isLog,d=b.isDatetimeAxis,A=b.isXAxis,m=b.isLinked,x=k.maxPadding,n=k.minPadding,D,h=k.tickInterval,C=k.tickPixelInterval,z=b.categories,H=l(b.threshold)?b.threshold:null,w=b.softThreshold,u,y,E;d||z||m||this.getTickAmount();
y=F(b.userMin,k.min);E=F(b.userMax,k.max);m?(b.linkedParent=f[b.coll][k.linkedTo],D=b.linkedParent.getExtremes(),b.min=F(D.min,D.dataMin),b.max=F(D.max,D.dataMax),k.type!==b.linkedParent.options.type&&a.error(11,1,f)):(!w&&v(H)&&(b.dataMin>=H?(D=H,n=0):b.dataMax<=H&&(u=H,x=0)),b.min=F(y,D,b.dataMin),b.max=F(E,u,b.dataMax));g&&(b.positiveValuesOnly&&!e&&0>=Math.min(b.min,F(b.dataMin,b.min))&&a.error(10,1,f),b.min=r(b.log2lin(b.min),15),b.max=r(b.log2lin(b.max),15));b.range&&v(b.max)&&(b.userMin=b.min=
y=Math.max(b.dataMin,b.minFromRange()),b.userMax=E=b.max,b.range=null);c(b,"foundExtremes");b.beforePadding&&b.beforePadding();b.adjustForMinRange();!(z||b.axisPointRange||b.usePercentage||m)&&v(b.min)&&v(b.max)&&(f=b.max-b.min)&&(!v(y)&&n&&(b.min-=f*n),!v(E)&&x&&(b.max+=f*x));l(k.softMin)&&!l(b.userMin)&&(b.min=Math.min(b.min,k.softMin));l(k.softMax)&&!l(b.userMax)&&(b.max=Math.max(b.max,k.softMax));l(k.floor)&&(b.min=Math.min(Math.max(b.min,k.floor),Number.MAX_VALUE));l(k.ceiling)&&(b.max=Math.max(Math.min(b.max,
k.ceiling),F(b.userMax,-Number.MAX_VALUE)));w&&v(b.dataMin)&&(H=H||0,!v(y)&&b.min<H&&b.dataMin>=H?b.min=H:!v(E)&&b.max>H&&b.dataMax<=H&&(b.max=H));b.tickInterval=b.min===b.max||void 0===b.min||void 0===b.max?1:m&&!h&&C===b.linkedParent.options.tickPixelInterval?h=b.linkedParent.tickInterval:F(h,this.tickAmount?(b.max-b.min)/Math.max(this.tickAmount-1,1):void 0,z?1:(b.max-b.min)*C/Math.max(b.len,C));A&&!e&&b.series.forEach(function(a){a.processData(b.min!==b.oldMin||b.max!==b.oldMax)});b.setAxisTranslation(!0);
b.beforeSetTickPositions&&b.beforeSetTickPositions();b.postProcessTickInterval&&(b.tickInterval=b.postProcessTickInterval(b.tickInterval));b.pointRange&&!h&&(b.tickInterval=Math.max(b.pointRange,b.tickInterval));e=F(k.minTickInterval,b.isDatetimeAxis&&b.closestPointRange);!h&&b.tickInterval<e&&(b.tickInterval=e);d||g||h||(b.tickInterval=t(b.tickInterval,null,p(b.tickInterval),F(k.allowDecimals,!(.5<b.tickInterval&&5>b.tickInterval&&1E3<b.max&&9999>b.max)),!!this.tickAmount));this.tickAmount||(b.tickInterval=
b.unsquish());this.setTickPositions()},setTickPositions:function(){var e=this.options,b,f=e.tickPositions;b=this.getMinorTickInterval();var k=e.tickPositioner,l=e.startOnTick,g=e.endOnTick;this.tickmarkOffset=this.categories&&"between"===e.tickmarkPlacement&&1===this.tickInterval?.5:0;this.minorTickInterval="auto"===b&&this.tickInterval?this.tickInterval/5:b;this.single=this.min===this.max&&v(this.min)&&!this.tickAmount&&(parseInt(this.min,10)===this.min||!1!==e.allowDecimals);this.tickPositions=
b=f&&f.slice();!b&&(!this.ordinalPositions&&(this.max-this.min)/this.tickInterval>Math.max(2*this.len,200)?(b=[this.min,this.max],a.error(19,!1,this.chart)):b=this.isDatetimeAxis?this.getTimeTicks(this.normalizeTimeTickInterval(this.tickInterval,e.units),this.min,this.max,e.startOfWeek,this.ordinalPositions,this.closestPointRange,!0):this.isLog?this.getLogTickPositions(this.tickInterval,this.min,this.max):this.getLinearTickPositions(this.tickInterval,this.min,this.max),b.length>this.len&&(b=[b[0],
b.pop()],b[0]===b[1]&&(b.length=1)),this.tickPositions=b,k&&(k=k.apply(this,[this.min,this.max])))&&(this.tickPositions=b=k);this.paddedTicks=b.slice(0);this.trimTicks(b,l,g);this.isLinked||(this.single&&2>b.length&&(this.min-=.5,this.max+=.5),f||k||this.adjustTickAmount());c(this,"afterSetTickPositions")},trimTicks:function(a,b,f){var e=a[0],k=a[a.length-1],q=this.minPointOffset||0;if(!this.isLinked){if(b&&-Infinity!==e)this.min=e;else for(;this.min-q>a[0];)a.shift();if(f)this.max=k;else for(;this.max+
q<a[a.length-1];)a.pop();0===a.length&&v(e)&&!this.options.tickPositions&&a.push((k+e)/2)}},alignToOthers:function(){var a={},b,f=this.options;!1===this.chart.options.chart.alignTicks||!1===f.alignTicks||!1===f.startOnTick||!1===f.endOnTick||this.isLog||this.chart[this.coll].forEach(function(e){var f=e.options,f=[e.horiz?f.left:f.top,f.width,f.height,f.pane].join();e.series.length&&(a[f]?b=!0:a[f]=1)});return b},getTickAmount:function(){var a=this.options,b=a.tickAmount,f=a.tickPixelInterval;!v(a.tickInterval)&&
this.len<f&&!this.isRadial&&!this.isLog&&a.startOnTick&&a.endOnTick&&(b=2);!b&&this.alignToOthers()&&(b=Math.ceil(this.len/f)+1);4>b&&(this.finalTickAmt=b,b=5);this.tickAmount=b},adjustTickAmount:function(){var a=this.tickInterval,b=this.tickPositions,f=this.tickAmount,k=this.finalTickAmt,c=b&&b.length,l=F(this.threshold,this.softThreshold?0:null);if(this.hasData()){if(c<f){for(;b.length<f;)b.length%2||this.min===l?b.push(r(b[b.length-1]+a)):b.unshift(r(b[0]-a));this.transA*=(c-1)/(f-1);this.min=
b[0];this.max=b[b.length-1]}else c>f&&(this.tickInterval*=2,this.setTickPositions());if(v(k)){for(a=f=b.length;a--;)(3===k&&1===a%2||2>=k&&0<a&&a<f-1)&&b.splice(a,1);this.finalTickAmt=void 0}}},setScale:function(){var a,b;this.oldMin=this.min;this.oldMax=this.max;this.oldAxisLength=this.len;this.setAxisSize();b=this.len!==this.oldAxisLength;this.series.forEach(function(e){if(e.isDirtyData||e.isDirty||e.xAxis.isDirty)a=!0});b||a||this.isLinked||this.forceRedraw||this.userMin!==this.oldUserMin||this.userMax!==
this.oldUserMax||this.alignToOthers()?(this.resetStacks&&this.resetStacks(),this.forceRedraw=!1,this.getSeriesExtremes(),this.setTickInterval(),this.oldUserMin=this.userMin,this.oldUserMax=this.userMax,this.isDirty||(this.isDirty=b||this.min!==this.oldMin||this.max!==this.oldMax)):this.cleanStacks&&this.cleanStacks();c(this,"afterSetScale")},setExtremes:function(a,b,f,k,l){var e=this,q=e.chart;f=F(f,!0);e.series.forEach(function(a){delete a.kdTree});l=g(l,{min:a,max:b});c(e,"setExtremes",l,function(){e.userMin=
a;e.userMax=b;e.eventArgs=l;f&&q.redraw(k)})},zoom:function(a,b){var e=this.dataMin,f=this.dataMax,k=this.options,c=Math.min(e,F(k.min,e)),k=Math.max(f,F(k.max,f));if(a!==this.min||b!==this.max)this.allowZoomOutside||(v(e)&&(a<c&&(a=c),a>k&&(a=k)),v(f)&&(b<c&&(b=c),b>k&&(b=k))),this.displayBtn=void 0!==a||void 0!==b,this.setExtremes(a,b,!1,void 0,{trigger:"zoom"});return!0},setAxisSize:function(){var e=this.chart,b=this.options,f=b.offsets||[0,0,0,0],k=this.horiz,c=this.width=Math.round(a.relativeLength(F(b.width,
e.plotWidth-f[3]+f[1]),e.plotWidth)),l=this.height=Math.round(a.relativeLength(F(b.height,e.plotHeight-f[0]+f[2]),e.plotHeight)),g=this.top=Math.round(a.relativeLength(F(b.top,e.plotTop+f[0]),e.plotHeight,e.plotTop)),b=this.left=Math.round(a.relativeLength(F(b.left,e.plotLeft+f[3]),e.plotWidth,e.plotLeft));this.bottom=e.chartHeight-l-g;this.right=e.chartWidth-c-b;this.len=Math.max(k?c:l,0);this.pos=k?b:g},getExtremes:function(){var a=this.isLog;return{min:a?r(this.lin2log(this.min)):this.min,max:a?
r(this.lin2log(this.max)):this.max,dataMin:this.dataMin,dataMax:this.dataMax,userMin:this.userMin,userMax:this.userMax}},getThreshold:function(a){var e=this.isLog,b=e?this.lin2log(this.min):this.min,e=e?this.lin2log(this.max):this.max;null===a||-Infinity===a?a=b:Infinity===a?a=e:b>a?a=b:e<a&&(a=e);return this.translate(a,0,1,0,1)},autoLabelAlign:function(a){a=(F(a,0)-90*this.side+720)%360;return 15<a&&165>a?"right":195<a&&345>a?"left":"center"},tickSize:function(a){var e=this.options,b=e[a+"Length"],
f=F(e[a+"Width"],"tick"===a&&this.isXAxis?1:0);if(f&&b)return"inside"===e[a+"Position"]&&(b=-b),[b,f]},labelMetrics:function(){var a=this.tickPositions&&this.tickPositions[0]||0;return this.chart.renderer.fontMetrics(this.options.labels.style&&this.options.labels.style.fontSize,this.ticks[a]&&this.ticks[a].label)},unsquish:function(){var a=this.options.labels,b=this.horiz,f=this.tickInterval,k=f,c=this.len/(((this.categories?1:0)+this.max-this.min)/f),l,g=a.rotation,d=this.labelMetrics(),p,A=Number.MAX_VALUE,
t,m=function(a){a/=c||1;a=1<a?Math.ceil(a):1;return r(a*f)};b?(t=!a.staggerLines&&!a.step&&(v(g)?[g]:c<F(a.autoRotationLimit,80)&&a.autoRotation))&&t.forEach(function(a){var e;if(a===g||a&&-90<=a&&90>=a)p=m(Math.abs(d.h/Math.sin(w*a))),e=p+Math.abs(a/360),e<A&&(A=e,l=a,k=p)}):a.step||(k=m(d.h));this.autoRotation=t;this.labelRotation=F(l,g);return k},getSlotWidth:function(a){var e=this.chart,b=this.horiz,f=this.options.labels,k=Math.max(this.tickPositions.length-(this.categories?0:1),1),c=e.margin[3];
return a&&a.slotWidth||b&&2>(f.step||0)&&!f.rotation&&(this.staggerLines||1)*this.len/k||!b&&(f.style&&parseInt(f.style.width,10)||c&&c-e.spacing[3]||.33*e.chartWidth)},renderUnsquish:function(){var a=this.chart,b=a.renderer,k=this.tickPositions,c=this.ticks,l=this.options.labels,g=l&&l.style||{},d=this.horiz,p=this.getSlotWidth(),A=Math.max(1,Math.round(p-2*(l.padding||5))),t={},m=this.labelMetrics(),x=l.style&&l.style.textOverflow,n,D,h=0,C;f(l.rotation)||(t.rotation=l.rotation||0);k.forEach(function(a){(a=
c[a])&&a.label&&a.label.textPxLength>h&&(h=a.label.textPxLength)});this.maxLabelLength=h;if(this.autoRotation)h>A&&h>m.h?t.rotation=this.labelRotation:this.labelRotation=0;else if(p&&(n=A,!x))for(D="clip",A=k.length;!d&&A--;)if(C=k[A],C=c[C].label)C.styles&&"ellipsis"===C.styles.textOverflow?C.css({textOverflow:"clip"}):C.textPxLength>p&&C.css({width:p+"px"}),C.getBBox().height>this.len/k.length-(m.h-m.f)&&(C.specificTextOverflow="ellipsis");t.rotation&&(n=h>.5*a.chartHeight?.33*a.chartHeight:h,x||
(D="ellipsis"));if(this.labelAlign=l.align||this.autoLabelAlign(this.labelRotation))t.align=this.labelAlign;k.forEach(function(a){var b=(a=c[a])&&a.label,e=g.width,f={};b&&(b.attr(t),a.shortenLabel?a.shortenLabel():n&&!e&&"nowrap"!==g.whiteSpace&&(n<b.textPxLength||"SPAN"===b.element.tagName)?(f.width=n,x||(f.textOverflow=b.specificTextOverflow||D),b.css(f)):b.styles&&b.styles.width&&!f.width&&!e&&b.css({width:null}),delete b.specificTextOverflow,a.rotation=t.rotation)},this);this.tickRotCorr=b.rotCorr(m.b,
this.labelRotation||0,0!==this.side)},hasData:function(){return this.hasVisibleSeries||v(this.min)&&v(this.max)&&this.tickPositions&&0<this.tickPositions.length},addTitle:function(a){var b=this.chart.renderer,e=this.horiz,f=this.opposite,k=this.options.title,c,l=this.chart.styledMode;this.axisTitle||((c=k.textAlign)||(c=(e?{low:"left",middle:"center",high:"right"}:{low:f?"right":"left",middle:"center",high:f?"left":"right"})[k.align]),this.axisTitle=b.text(k.text,0,0,k.useHTML).attr({zIndex:7,rotation:k.rotation||
0,align:c}).addClass("highcharts-axis-title"),l||this.axisTitle.css(x(k.style)),this.axisTitle.add(this.axisGroup),this.axisTitle.isNew=!0);l||k.style.width||this.isRadial||this.axisTitle.css({width:this.len});this.axisTitle[a?"show":"hide"](!0)},generateTick:function(a){var b=this.ticks;b[a]?b[a].addLabel():b[a]=new D(this,a)},getOffset:function(){var a=this,b=a.chart,f=b.renderer,k=a.options,l=a.tickPositions,g=a.ticks,d=a.horiz,p=a.side,A=b.inverted&&!a.isZAxis?[1,0,3,2][p]:p,t,m,x=0,n,D=0,h=k.title,
C=k.labels,z=0,r=b.axisOffset,b=b.clipOffset,w=[-1,1,1,-1][p],u=k.className,y=a.axisParent;t=a.hasData();a.showAxis=m=t||F(k.showEmpty,!0);a.staggerLines=a.horiz&&C.staggerLines;a.axisGroup||(a.gridGroup=f.g("grid").attr({zIndex:k.gridZIndex||1}).addClass("highcharts-"+this.coll.toLowerCase()+"-grid "+(u||"")).add(y),a.axisGroup=f.g("axis").attr({zIndex:k.zIndex||2}).addClass("highcharts-"+this.coll.toLowerCase()+" "+(u||"")).add(y),a.labelGroup=f.g("axis-labels").attr({zIndex:C.zIndex||7}).addClass("highcharts-"+
a.coll.toLowerCase()+"-labels "+(u||"")).add(y));t||a.isLinked?(l.forEach(function(b,e){a.generateTick(b,e)}),a.renderUnsquish(),a.reserveSpaceDefault=0===p||2===p||{1:"left",3:"right"}[p]===a.labelAlign,F(C.reserveSpace,"center"===a.labelAlign?!0:null,a.reserveSpaceDefault)&&l.forEach(function(a){z=Math.max(g[a].getLabelSize(),z)}),a.staggerLines&&(z*=a.staggerLines),a.labelOffset=z*(a.opposite?-1:1)):H(g,function(a,b){a.destroy();delete g[b]});h&&h.text&&!1!==h.enabled&&(a.addTitle(m),m&&!1!==h.reserveSpace&&
(a.titleOffset=x=a.axisTitle.getBBox()[d?"height":"width"],n=h.offset,D=v(n)?0:F(h.margin,d?5:10)));a.renderLine();a.offset=w*F(k.offset,r[p]);a.tickRotCorr=a.tickRotCorr||{x:0,y:0};f=0===p?-a.labelMetrics().h:2===p?a.tickRotCorr.y:0;D=Math.abs(z)+D;z&&(D=D-f+w*(d?F(C.y,a.tickRotCorr.y+8*w):C.x));a.axisTitleMargin=F(n,D);a.getMaxLabelDimensions&&(a.maxLabelDimensions=a.getMaxLabelDimensions(g,l));d=this.tickSize("tick");r[p]=Math.max(r[p],a.axisTitleMargin+x+w*a.offset,D,t&&l.length&&d?d[0]+w*a.offset:
0);k=k.offset?0:2*Math.floor(a.axisLine.strokeWidth()/2);b[A]=Math.max(b[A],k);c(this,"afterGetOffset")},getLinePath:function(a){var b=this.chart,e=this.opposite,f=this.offset,k=this.horiz,c=this.left+(e?this.width:0)+f,f=b.chartHeight-this.bottom-(e?this.height:0)+f;e&&(a*=-1);return b.renderer.crispLine(["M",k?this.left:c,k?f:this.top,"L",k?b.chartWidth-this.right:c,k?f:b.chartHeight-this.bottom],a)},renderLine:function(){this.axisLine||(this.axisLine=this.chart.renderer.path().addClass("highcharts-axis-line").add(this.axisGroup),
this.chart.styledMode||this.axisLine.attr({stroke:this.options.lineColor,"stroke-width":this.options.lineWidth,zIndex:7}))},getTitlePosition:function(){var a=this.horiz,b=this.left,f=this.top,k=this.len,c=this.options.title,l=a?b:f,g=this.opposite,d=this.offset,p=c.x||0,A=c.y||0,t=this.axisTitle,m=this.chart.renderer.fontMetrics(c.style&&c.style.fontSize,t),t=Math.max(t.getBBox(null,0).height-m.h-1,0),k={low:l+(a?0:k),middle:l+k/2,high:l+(a?k:0)}[c.align],b=(a?f+this.height:b)+(a?1:-1)*(g?-1:1)*this.axisTitleMargin+
[-t,t,m.f,-t][this.side];return{x:a?k+p:b+(g?this.width:0)+d+p,y:a?b+A-(g?this.height:0)+d:k+A}},renderMinorTick:function(a){var b=this.chart.hasRendered&&l(this.oldMin),e=this.minorTicks;e[a]||(e[a]=new D(this,a,"minor"));b&&e[a].isNew&&e[a].render(null,!0);e[a].render(null,!1,1)},renderTick:function(a,b){var e=this.isLinked,f=this.ticks,k=this.chart.hasRendered&&l(this.oldMin);if(!e||a>=this.min&&a<=this.max)f[a]||(f[a]=new D(this,a)),k&&f[a].isNew&&f[a].render(b,!0,-1),f[a].render(b)},render:function(){var b=
this,f=b.chart,k=b.options,g=b.isLog,d=b.isLinked,p=b.tickPositions,t=b.axisTitle,m=b.ticks,x=b.minorTicks,n=b.alternateBands,h=k.stackLabels,C=k.alternateGridColor,z=b.tickmarkOffset,F=b.axisLine,r=b.showAxis,w=G(f.renderer.globalAnimation),u,v;b.labelEdge.length=0;b.overlap=!1;[m,x,n].forEach(function(a){H(a,function(a){a.isActive=!1})});if(b.hasData()||d)b.minorTickInterval&&!b.categories&&b.getMinorTickPositions().forEach(function(a){b.renderMinorTick(a)}),p.length&&(p.forEach(function(a,e){b.renderTick(a,
e)}),z&&(0===b.min||b.single)&&(m[-1]||(m[-1]=new D(b,-1,null,!0)),m[-1].render(-1))),C&&p.forEach(function(e,k){v=void 0!==p[k+1]?p[k+1]+z:b.max-z;0===k%2&&e<b.max&&v<=b.max+(f.polar?-z:z)&&(n[e]||(n[e]=new a.PlotLineOrBand(b)),u=e+z,n[e].options={from:g?b.lin2log(u):u,to:g?b.lin2log(v):v,color:C},n[e].render(),n[e].isActive=!0)}),b._addedPlotLB||((k.plotLines||[]).concat(k.plotBands||[]).forEach(function(a){b.addPlotBandOrLine(a)}),b._addedPlotLB=!0);[m,x,n].forEach(function(a){var b,e=[],k=w.duration;
H(a,function(a,b){a.isActive||(a.render(b,!1,0),a.isActive=!1,e.push(b))});A(function(){for(b=e.length;b--;)a[e[b]]&&!a[e[b]].isActive&&(a[e[b]].destroy(),delete a[e[b]])},a!==n&&f.hasRendered&&k?k:0)});F&&(F[F.isPlaced?"animate":"attr"]({d:this.getLinePath(F.strokeWidth())}),F.isPlaced=!0,F[r?"show":"hide"](!0));t&&r&&(k=b.getTitlePosition(),l(k.y)?(t[t.isNew?"attr":"animate"](k),t.isNew=!1):(t.attr("y",-9999),t.isNew=!0));h&&h.enabled&&b.renderStackTotals();b.isDirty=!1;c(this,"afterRender")},redraw:function(){this.visible&&
(this.render(),this.plotLinesAndBands.forEach(function(a){a.render()}));this.series.forEach(function(a){a.isDirty=!0})},keepProps:"extKey hcEvents names series userMax userMin".split(" "),destroy:function(a){var b=this,e=b.stacks,f=b.plotLinesAndBands,k;c(this,"destroy",{keepEvents:a});a||z(b);H(e,function(a,b){n(a);e[b]=null});[b.ticks,b.minorTicks,b.alternateBands].forEach(function(a){n(a)});if(f)for(a=f.length;a--;)f[a].destroy();"stackTotalGroup axisLine axisTitle axisGroup gridGroup labelGroup cross scrollbar".split(" ").forEach(function(a){b[a]&&
(b[a]=b[a].destroy())});for(k in b.plotLinesAndBandsGroups)b.plotLinesAndBandsGroups[k]=b.plotLinesAndBandsGroups[k].destroy();H(b,function(a,e){-1===b.keepProps.indexOf(e)&&delete b[e]})},drawCrosshair:function(a,b){var e,f=this.crosshair,k=F(f.snap,!0),l,g=this.cross;c(this,"drawCrosshair",{e:a,point:b});a||(a=this.cross&&this.cross.e);if(this.crosshair&&!1!==(v(b)||!k)){k?v(b)&&(l=F(b.crosshairPos,this.isXAxis?b.plotX:this.len-b.plotY)):l=a&&(this.horiz?a.chartX-this.pos:this.len-a.chartY+this.pos);
v(l)&&(e=this.getPlotLinePath(b&&(this.isXAxis?b.x:F(b.stackY,b.y)),null,null,null,l)||null);if(!v(e)){this.hideCrosshair();return}k=this.categories&&!this.isRadial;g||(this.cross=g=this.chart.renderer.path().addClass("highcharts-crosshair highcharts-crosshair-"+(k?"category ":"thin ")+f.className).attr({zIndex:F(f.zIndex,2)}).add(),this.chart.styledMode||(g.attr({stroke:f.color||(k?d("#ccd6eb").setOpacity(.25).get():"#cccccc"),"stroke-width":F(f.width,1)}).css({"pointer-events":"none"}),f.dashStyle&&
g.attr({dashstyle:f.dashStyle})));g.show().attr({d:e});k&&!f.width&&g.attr({"stroke-width":this.transA});this.cross.e=a}else this.hideCrosshair();c(this,"afterDrawCrosshair",{e:a,point:b})},hideCrosshair:function(){this.cross&&this.cross.hide()}});return a.Axis=C}(J);(function(a){var y=a.Axis,G=a.getMagnitude,E=a.normalizeTickInterval,h=a.timeUnits;y.prototype.getTimeTicks=function(){return this.chart.time.getTimeTicks.apply(this.chart.time,arguments)};y.prototype.normalizeTimeTickInterval=function(a,
r){var d=r||[["millisecond",[1,2,5,10,20,25,50,100,200,500]],["second",[1,2,5,10,15,30]],["minute",[1,2,5,10,15,30]],["hour",[1,2,3,4,6,8,12]],["day",[1,2]],["week",[1,2]],["month",[1,2,3,4,6]],["year",null]];r=d[d.length-1];var v=h[r[0]],w=r[1],n;for(n=0;n<d.length&&!(r=d[n],v=h[r[0]],w=r[1],d[n+1]&&a<=(v*w[w.length-1]+h[d[n+1][0]])/2);n++);v===h.year&&a<5*v&&(w=[1,2,5]);a=E(a/v,w,"year"===r[0]?Math.max(G(a/v),1):1);return{unitRange:v,count:a,unitName:r[0]}}})(J);(function(a){var y=a.Axis,G=a.getMagnitude,
E=a.normalizeTickInterval,h=a.pick;y.prototype.getLogTickPositions=function(a,r,u,v){var d=this.options,n=this.len,g=[];v||(this._minorAutoInterval=null);if(.5<=a)a=Math.round(a),g=this.getLinearTickPositions(a,r,u);else if(.08<=a)for(var n=Math.floor(r),c,m,p,b,l,d=.3<a?[1,2,4]:.15<a?[1,2,4,6,8]:[1,2,3,4,5,6,7,8,9];n<u+1&&!l;n++)for(m=d.length,c=0;c<m&&!l;c++)p=this.log2lin(this.lin2log(n)*d[c]),p>r&&(!v||b<=u)&&void 0!==b&&g.push(b),b>u&&(l=!0),b=p;else r=this.lin2log(r),u=this.lin2log(u),a=v?this.getMinorTickInterval():
d.tickInterval,a=h("auto"===a?null:a,this._minorAutoInterval,d.tickPixelInterval/(v?5:1)*(u-r)/((v?n/this.tickPositions.length:n)||1)),a=E(a,null,G(a)),g=this.getLinearTickPositions(a,r,u).map(this.log2lin),v||(this._minorAutoInterval=a/5);v||(this.tickInterval=a);return g};y.prototype.log2lin=function(a){return Math.log(a)/Math.LN10};y.prototype.lin2log=function(a){return Math.pow(10,a)}})(J);(function(a,y){var G=a.arrayMax,E=a.arrayMin,h=a.defined,d=a.destroyObjectProperties,r=a.erase,u=a.merge,
v=a.pick;a.PlotLineOrBand=function(a,d){this.axis=a;d&&(this.options=d,this.id=d.id)};a.PlotLineOrBand.prototype={render:function(){a.fireEvent(this,"render");var d=this,n=d.axis,g=n.horiz,c=d.options,m=c.label,p=d.label,b=c.to,l=c.from,f=c.value,x=h(l)&&h(b),t=h(f),H=d.svgElem,F=!H,z=[],k=c.color,A=v(c.zIndex,0),D=c.events,z={"class":"highcharts-plot-"+(x?"band ":"line ")+(c.className||"")},C={},e=n.chart.renderer,q=x?"bands":"lines";n.isLog&&(l=n.log2lin(l),b=n.log2lin(b),f=n.log2lin(f));n.chart.styledMode||
(t?(z.stroke=k,z["stroke-width"]=c.width,c.dashStyle&&(z.dashstyle=c.dashStyle)):x&&(k&&(z.fill=k),c.borderWidth&&(z.stroke=c.borderColor,z["stroke-width"]=c.borderWidth)));C.zIndex=A;q+="-"+A;(k=n.plotLinesAndBandsGroups[q])||(n.plotLinesAndBandsGroups[q]=k=e.g("plot-"+q).attr(C).add());F&&(d.svgElem=H=e.path().attr(z).add(k));if(t)z=n.getPlotLinePath(f,H.strokeWidth());else if(x)z=n.getPlotBandPath(l,b,c);else return;F&&z&&z.length?(H.attr({d:z}),D&&a.objectEach(D,function(a,b){H.on(b,function(a){D[b].apply(d,
[a])})})):H&&(z?(H.show(),H.animate({d:z})):(H.hide(),p&&(d.label=p=p.destroy())));m&&h(m.text)&&z&&z.length&&0<n.width&&0<n.height&&!z.isFlat?(m=u({align:g&&x&&"center",x:g?!x&&4:10,verticalAlign:!g&&x&&"middle",y:g?x?16:10:x?6:-4,rotation:g&&!x&&90},m),this.renderLabel(m,z,x,A)):p&&p.hide();return d},renderLabel:function(a,d,g,c){var m=this.label,p=this.axis.chart.renderer;m||(m={align:a.textAlign||a.align,rotation:a.rotation,"class":"highcharts-plot-"+(g?"band":"line")+"-label "+(a.className||
"")},m.zIndex=c,this.label=m=p.text(a.text,0,0,a.useHTML).attr(m).add(),this.axis.chart.styledMode||m.css(a.style));c=d.xBounds||[d[1],d[4],g?d[6]:d[1]];d=d.yBounds||[d[2],d[5],g?d[7]:d[2]];g=E(c);p=E(d);m.align(a,!1,{x:g,y:p,width:G(c)-g,height:G(d)-p});m.show()},destroy:function(){r(this.axis.plotLinesAndBands,this);delete this.axis;d(this)}};a.extend(y.prototype,{getPlotBandPath:function(a,d){var g=this.getPlotLinePath(d,null,null,!0),c=this.getPlotLinePath(a,null,null,!0),m=[],p=this.horiz,b=
1,l;a=a<this.min&&d<this.min||a>this.max&&d>this.max;if(c&&g)for(a&&(l=c.toString()===g.toString(),b=0),a=0;a<c.length;a+=6)p&&g[a+1]===c[a+1]?(g[a+1]+=b,g[a+4]+=b):p||g[a+2]!==c[a+2]||(g[a+2]+=b,g[a+5]+=b),m.push("M",c[a+1],c[a+2],"L",c[a+4],c[a+5],g[a+4],g[a+5],g[a+1],g[a+2],"z"),m.isFlat=l;return m},addPlotBand:function(a){return this.addPlotBandOrLine(a,"plotBands")},addPlotLine:function(a){return this.addPlotBandOrLine(a,"plotLines")},addPlotBandOrLine:function(d,n){var g=(new a.PlotLineOrBand(this,
d)).render(),c=this.userOptions;g&&(n&&(c[n]=c[n]||[],c[n].push(d)),this.plotLinesAndBands.push(g));return g},removePlotBandOrLine:function(a){for(var d=this.plotLinesAndBands,g=this.options,c=this.userOptions,m=d.length;m--;)d[m].id===a&&d[m].destroy();[g.plotLines||[],c.plotLines||[],g.plotBands||[],c.plotBands||[]].forEach(function(c){for(m=c.length;m--;)c[m].id===a&&r(c,c[m])})},removePlotBand:function(a){this.removePlotBandOrLine(a)},removePlotLine:function(a){this.removePlotBandOrLine(a)}})})(J,
V);(function(a){var y=a.doc,G=a.extend,E=a.format,h=a.isNumber,d=a.merge,r=a.pick,u=a.splat,v=a.syncTimeout,w=a.timeUnits;a.Tooltip=function(){this.init.apply(this,arguments)};a.Tooltip.prototype={init:function(a,d){this.chart=a;this.options=d;this.crosshairs=[];this.now={x:0,y:0};this.isHidden=!0;this.split=d.split&&!a.inverted;this.shared=d.shared||this.split;this.outside=d.outside&&!this.split},cleanSplit:function(a){this.chart.series.forEach(function(d){var c=d&&d.tt;c&&(!c.isActive||a?d.tt=c.destroy():
c.isActive=!1)})},applyFilter:function(){var a=this.chart;a.renderer.definition({tagName:"filter",id:"drop-shadow-"+a.index,opacity:.5,children:[{tagName:"feGaussianBlur",in:"SourceAlpha",stdDeviation:1},{tagName:"feOffset",dx:1,dy:1},{tagName:"feComponentTransfer",children:[{tagName:"feFuncA",type:"linear",slope:.3}]},{tagName:"feMerge",children:[{tagName:"feMergeNode"},{tagName:"feMergeNode",in:"SourceGraphic"}]}]});a.renderer.definition({tagName:"style",textContent:".highcharts-tooltip-"+a.index+
"{filter:url(#drop-shadow-"+a.index+")}"})},getLabel:function(){var d=this.chart.renderer,g=this.chart.styledMode,c=this.options,m;this.label||(this.outside&&(this.container=m=a.doc.createElement("div"),m.className="highcharts-tooltip-container",a.css(m,{position:"absolute",top:"1px",pointerEvents:c.style&&c.style.pointerEvents}),a.doc.body.appendChild(m),this.renderer=d=new a.Renderer(m,0,0)),this.split?this.label=d.g("tooltip"):(this.label=d.label("",0,0,c.shape||"callout",null,null,c.useHTML,null,
"tooltip").attr({padding:c.padding,r:c.borderRadius}),g||this.label.attr({fill:c.backgroundColor,"stroke-width":c.borderWidth}).css(c.style).shadow(c.shadow)),g&&(this.applyFilter(),this.label.addClass("highcharts-tooltip-"+this.chart.index)),this.outside&&(this.label.attr({x:this.distance,y:this.distance}),this.label.xSetter=function(a){m.style.left=a+"px"},this.label.ySetter=function(a){m.style.top=a+"px"}),this.label.attr({zIndex:8}).add());return this.label},update:function(a){this.destroy();
d(!0,this.chart.options.tooltip.userOptions,a);this.init(this.chart,d(!0,this.options,a))},destroy:function(){this.label&&(this.label=this.label.destroy());this.split&&this.tt&&(this.cleanSplit(this.chart,!0),this.tt=this.tt.destroy());this.renderer&&(this.renderer=this.renderer.destroy(),a.discardElement(this.container));a.clearTimeout(this.hideTimer);a.clearTimeout(this.tooltipTimeout)},move:function(d,g,c,m){var p=this,b=p.now,l=!1!==p.options.animation&&!p.isHidden&&(1<Math.abs(d-b.x)||1<Math.abs(g-
b.y)),f=p.followPointer||1<p.len;G(b,{x:l?(2*b.x+d)/3:d,y:l?(b.y+g)/2:g,anchorX:f?void 0:l?(2*b.anchorX+c)/3:c,anchorY:f?void 0:l?(b.anchorY+m)/2:m});p.getLabel().attr(b);l&&(a.clearTimeout(this.tooltipTimeout),this.tooltipTimeout=setTimeout(function(){p&&p.move(d,g,c,m)},32))},hide:function(d){var g=this;a.clearTimeout(this.hideTimer);d=r(d,this.options.hideDelay,500);this.isHidden||(this.hideTimer=v(function(){g.getLabel()[d?"fadeOut":"hide"]();g.isHidden=!0},d))},getAnchor:function(a,d){var c=
this.chart,g=c.pointer,p=c.inverted,b=c.plotTop,l=c.plotLeft,f=0,x=0,t,h;a=u(a);this.followPointer&&d?(void 0===d.chartX&&(d=g.normalize(d)),a=[d.chartX-c.plotLeft,d.chartY-b]):a[0].tooltipPos?a=a[0].tooltipPos:(a.forEach(function(a){t=a.series.yAxis;h=a.series.xAxis;f+=a.plotX+(!p&&h?h.left-l:0);x+=(a.plotLow?(a.plotLow+a.plotHigh)/2:a.plotY)+(!p&&t?t.top-b:0)}),f/=a.length,x/=a.length,a=[p?c.plotWidth-x:f,this.shared&&!p&&1<a.length&&d?d.chartY-b:p?c.plotHeight-f:x]);return a.map(Math.round)},getPosition:function(a,
d,c){var g=this.chart,p=this.distance,b={},l=g.inverted&&c.h||0,f,x=this.outside,t=x?y.documentElement.clientWidth-2*p:g.chartWidth,h=x?Math.max(y.body.scrollHeight,y.documentElement.scrollHeight,y.body.offsetHeight,y.documentElement.offsetHeight,y.documentElement.clientHeight):g.chartHeight,n=g.pointer.chartPosition,z=["y",h,d,(x?n.top-p:0)+c.plotY+g.plotTop,x?0:g.plotTop,x?h:g.plotTop+g.plotHeight],k=["x",t,a,(x?n.left-p:0)+c.plotX+g.plotLeft,x?0:g.plotLeft,x?t:g.plotLeft+g.plotWidth],A=!this.followPointer&&
r(c.ttBelow,!g.inverted===!!c.negative),D=function(a,e,f,k,c,d){var g=f<k-p,q=k+p+f<e,t=k-p-f;k+=p;if(A&&q)b[a]=k;else if(!A&&g)b[a]=t;else if(g)b[a]=Math.min(d-f,0>t-l?t:t-l);else if(q)b[a]=Math.max(c,k+l+f>e?k:k+l);else return!1},C=function(a,e,f,k){var c;k<p||k>e-p?c=!1:b[a]=k<f/2?1:k>e-f/2?e-f-2:k-f/2;return c},e=function(a){var b=z;z=k;k=b;f=a},q=function(){!1!==D.apply(0,z)?!1!==C.apply(0,k)||f||(e(!0),q()):f?b.x=b.y=0:(e(!0),q())};(g.inverted||1<this.len)&&e();q();return b},defaultFormatter:function(a){var d=
this.points||u(this),c;c=[a.tooltipFooterHeaderFormatter(d[0])];c=c.concat(a.bodyFormatter(d));c.push(a.tooltipFooterHeaderFormatter(d[0],!0));return c},refresh:function(d,g){var c,m=this.options,p,b=d,l,f={},x=[];c=m.formatter||this.defaultFormatter;var f=this.shared,t,h=this.chart.styledMode;m.enabled&&(a.clearTimeout(this.hideTimer),this.followPointer=u(b)[0].series.tooltipOptions.followPointer,l=this.getAnchor(b,g),g=l[0],p=l[1],!f||b.series&&b.series.noSharedTooltip?f=b.getLabelConfig():(b.forEach(function(a){a.setState("hover");
x.push(a.getLabelConfig())}),f={x:b[0].category,y:b[0].y},f.points=x,b=b[0]),this.len=x.length,f=c.call(f,this),t=b.series,this.distance=r(t.tooltipOptions.distance,16),!1===f?this.hide():(c=this.getLabel(),this.isHidden&&c.attr({opacity:1}).show(),this.split?this.renderSplit(f,u(d)):(m.style.width&&!h||c.css({width:this.chart.spacingBox.width}),c.attr({text:f&&f.join?f.join(""):f}),c.removeClass(/highcharts-color-[\d]+/g).addClass("highcharts-color-"+r(b.colorIndex,t.colorIndex)),h||c.attr({stroke:m.borderColor||
b.color||t.color||"#666666"}),this.updatePosition({plotX:g,plotY:p,negative:b.negative,ttBelow:b.ttBelow,h:l[2]||0})),this.isHidden=!1))},renderSplit:function(d,g){var c=this,m=[],p=this.chart,b=p.renderer,l=!0,f=this.options,x=0,t,h=this.getLabel(),n=p.plotTop;a.isString(d)&&(d=[!1,d]);d.slice(0,g.length+1).forEach(function(a,k){if(!1!==a&&""!==a){k=g[k-1]||{isHeader:!0,plotX:g[0].plotX,plotY:p.plotHeight};var d=k.series||c,D=d.tt,C=k.series||{},e="highcharts-color-"+r(k.colorIndex,C.colorIndex,
"none");D||(D={padding:f.padding,r:f.borderRadius},p.styledMode||(D.fill=f.backgroundColor,D.stroke=f.borderColor||k.color||C.color||"#333333",D["stroke-width"]=f.borderWidth),d.tt=D=b.label(null,null,null,(k.isHeader?f.headerShape:f.shape)||"callout",null,null,f.useHTML).addClass("highcharts-tooltip-box "+e).attr(D).add(h));D.isActive=!0;D.attr({text:a});p.styledMode||D.css(f.style).shadow(f.shadow);a=D.getBBox();C=a.width+D.strokeWidth();k.isHeader?(x=a.height,p.xAxis[0].opposite&&(t=!0,n-=x),C=
Math.max(0,Math.min(k.plotX+p.plotLeft-C/2,p.chartWidth+(p.scrollablePixels?p.scrollablePixels-p.marginRight:0)-C))):C=k.plotX+p.plotLeft-r(f.distance,16)-C;0>C&&(l=!1);a=(k.series&&k.series.yAxis&&k.series.yAxis.pos)+(k.plotY||0);a-=n;k.isHeader&&(a=t?-x:p.plotHeight+x);m.push({target:a,rank:k.isHeader?1:0,size:d.tt.getBBox().height+1,point:k,x:C,tt:D})}});this.cleanSplit();f.positioner&&m.forEach(function(a){var b=f.positioner.call(c,a.tt.getBBox().width,a.size,a.point);a.x=b.x;a.align=0;a.target=
b.y;a.rank=r(b.rank,a.rank)});a.distribute(m,p.plotHeight+x);m.forEach(function(a){var b=a.point,c=b.series;a.tt.attr({visibility:void 0===a.pos?"hidden":"inherit",x:l||b.isHeader||f.positioner?a.x:b.plotX+p.plotLeft+r(f.distance,16),y:a.pos+n,anchorX:b.isHeader?b.plotX+p.plotLeft:b.plotX+c.xAxis.pos,anchorY:b.isHeader?p.plotTop+p.plotHeight/2:b.plotY+c.yAxis.pos})})},updatePosition:function(a){var d=this.chart,c=this.getLabel(),m=(this.options.positioner||this.getPosition).call(this,c.width,c.height,
a),p=a.plotX+d.plotLeft;a=a.plotY+d.plotTop;var b;this.outside&&(b=(this.options.borderWidth||0)+2*this.distance,this.renderer.setSize(c.width+b,c.height+b,!1),p+=d.pointer.chartPosition.left-m.x,a+=d.pointer.chartPosition.top-m.y);this.move(Math.round(m.x),Math.round(m.y||0),p,a)},getDateFormat:function(a,d,c,m){var g=this.chart.time,b=g.dateFormat("%m-%d %H:%M:%S.%L",d),l,f,x={millisecond:15,second:12,minute:9,hour:6,day:3},t="millisecond";for(f in w){if(a===w.week&&+g.dateFormat("%w",d)===c&&"00:00:00.000"===
b.substr(6)){f="week";break}if(w[f]>a){f=t;break}if(x[f]&&b.substr(x[f])!=="01-01 00:00:00.000".substr(x[f]))break;"week"!==f&&(t=f)}f&&(l=g.resolveDTLFormat(m[f]).main);return l},getXDateFormat:function(a,d,c){d=d.dateTimeLabelFormats;var g=c&&c.closestPointRange;return(g?this.getDateFormat(g,a.x,c.options.startOfWeek,d):d.day)||d.year},tooltipFooterHeaderFormatter:function(a,d){d=d?"footer":"header";var c=a.series,g=c.tooltipOptions,p=g.xDateFormat,b=c.xAxis,l=b&&"datetime"===b.options.type&&h(a.key),
f=g[d+"Format"];l&&!p&&(p=this.getXDateFormat(a,g,b));l&&p&&(a.point&&a.point.tooltipDateKeys||["key"]).forEach(function(a){f=f.replace("{point."+a+"}","{point."+a+":"+p+"}")});c.chart.styledMode&&(f=this.styledModeFormat(f));return E(f,{point:a,series:c},this.chart.time)},bodyFormatter:function(a){return a.map(function(a){var c=a.series.tooltipOptions;return(c[(a.point.formatPrefix||"point")+"Formatter"]||a.point.tooltipFormatter).call(a.point,c[(a.point.formatPrefix||"point")+"Format"]||"")})},
styledModeFormat:function(a){return a.replace('style\x3d"font-size: 10px"','class\x3d"highcharts-header"').replace(/style="color:{(point|series)\.color}"/g,'class\x3d"highcharts-color-{$1.colorIndex}"')}}})(J);(function(a){var y=a.addEvent,G=a.attr,E=a.charts,h=a.color,d=a.css,r=a.defined,u=a.extend,v=a.find,w=a.fireEvent,n=a.isNumber,g=a.isObject,c=a.offset,m=a.pick,p=a.splat,b=a.Tooltip;a.Pointer=function(a,b){this.init(a,b)};a.Pointer.prototype={init:function(a,f){this.options=f;this.chart=a;this.runChartClick=
f.chart.events&&!!f.chart.events.click;this.pinchDown=[];this.lastValidTouch={};b&&(a.tooltip=new b(a,f.tooltip),this.followTouchMove=m(f.tooltip.followTouchMove,!0));this.setDOMEvents()},zoomOption:function(a){var b=this.chart,c=b.options.chart,l=c.zoomType||"",b=b.inverted;/touch/.test(a.type)&&(l=m(c.pinchType,l));this.zoomX=a=/x/.test(l);this.zoomY=l=/y/.test(l);this.zoomHor=a&&!b||l&&b;this.zoomVert=l&&!b||a&&b;this.hasZoom=a||l},normalize:function(a,b){var f;f=a.touches?a.touches.length?a.touches.item(0):
a.changedTouches[0]:a;b||(this.chartPosition=b=c(this.chart.container));return u(a,{chartX:Math.round(f.pageX-b.left),chartY:Math.round(f.pageY-b.top)})},getCoordinates:function(a){var b={xAxis:[],yAxis:[]};this.chart.axes.forEach(function(f){b[f.isXAxis?"xAxis":"yAxis"].push({axis:f,value:f.toValue(a[f.horiz?"chartX":"chartY"])})});return b},findNearestKDPoint:function(a,b,c){var f;a.forEach(function(a){var l=!(a.noSharedTooltip&&b)&&0>a.options.findNearestPointBy.indexOf("y");a=a.searchPoint(c,
l);if((l=g(a,!0))&&!(l=!g(f,!0)))var l=f.distX-a.distX,d=f.dist-a.dist,k=(a.series.group&&a.series.group.zIndex)-(f.series.group&&f.series.group.zIndex),l=0<(0!==l&&b?l:0!==d?d:0!==k?k:f.series.index>a.series.index?-1:1);l&&(f=a)});return f},getPointFromEvent:function(a){a=a.target;for(var b;a&&!b;)b=a.point,a=a.parentNode;return b},getChartCoordinatesFromPoint:function(a,b){var f=a.series,c=f.xAxis,f=f.yAxis,l=m(a.clientX,a.plotX),d=a.shapeArgs;if(c&&f)return b?{chartX:c.len+c.pos-l,chartY:f.len+
f.pos-a.plotY}:{chartX:l+c.pos,chartY:a.plotY+f.pos};if(d&&d.x&&d.y)return{chartX:d.x,chartY:d.y}},getHoverData:function(a,b,c,d,p,h,n){var f,l=[],t=n&&n.isBoosting;d=!(!d||!a);n=b&&!b.stickyTracking?[b]:c.filter(function(a){return a.visible&&!(!p&&a.directTouch)&&m(a.options.enableMouseTracking,!0)&&a.stickyTracking});b=(f=d?a:this.findNearestKDPoint(n,p,h))&&f.series;f&&(p&&!b.noSharedTooltip?(n=c.filter(function(a){return a.visible&&!(!p&&a.directTouch)&&m(a.options.enableMouseTracking,!0)&&!a.noSharedTooltip}),
n.forEach(function(a){var b=v(a.points,function(a){return a.x===f.x&&!a.isNull});g(b)&&(t&&(b=a.getPoint(b)),l.push(b))})):l.push(f));return{hoverPoint:f,hoverSeries:b,hoverPoints:l}},runPointActions:function(b,f){var c=this.chart,d=c.tooltip&&c.tooltip.options.enabled?c.tooltip:void 0,l=d?d.shared:!1,g=f||c.hoverPoint,p=g&&g.series||c.hoverSeries,p=this.getHoverData(g,p,c.series,"touchmove"!==b.type&&(!!f||p&&p.directTouch&&this.isDirectTouch),l,b,{isBoosting:c.isBoosting}),k,g=p.hoverPoint;k=p.hoverPoints;
f=(p=p.hoverSeries)&&p.tooltipOptions.followPointer;l=l&&p&&!p.noSharedTooltip;if(g&&(g!==c.hoverPoint||d&&d.isHidden)){(c.hoverPoints||[]).forEach(function(a){-1===k.indexOf(a)&&a.setState()});(k||[]).forEach(function(a){a.setState("hover")});if(c.hoverSeries!==p)p.onMouseOver();c.hoverPoint&&c.hoverPoint.firePointEvent("mouseOut");if(!g.series)return;g.firePointEvent("mouseOver");c.hoverPoints=k;c.hoverPoint=g;d&&d.refresh(l?k:g,b)}else f&&d&&!d.isHidden&&(g=d.getAnchor([{}],b),d.updatePosition({plotX:g[0],
plotY:g[1]}));this.unDocMouseMove||(this.unDocMouseMove=y(c.container.ownerDocument,"mousemove",function(b){var f=E[a.hoverChartIndex];if(f)f.pointer.onDocumentMouseMove(b)}));c.axes.forEach(function(f){var c=m(f.crosshair.snap,!0),d=c?a.find(k,function(a){return a.series[f.coll]===f}):void 0;d||!c?f.drawCrosshair(b,d):f.hideCrosshair()})},reset:function(a,b){var f=this.chart,c=f.hoverSeries,d=f.hoverPoint,l=f.hoverPoints,g=f.tooltip,k=g&&g.shared?l:d;a&&k&&p(k).forEach(function(b){b.series.isCartesian&&
void 0===b.plotX&&(a=!1)});if(a)g&&k&&(g.refresh(k),g.shared&&l?l.forEach(function(a){a.setState(a.state,!0);a.series.isCartesian&&(a.series.xAxis.crosshair&&a.series.xAxis.drawCrosshair(null,a),a.series.yAxis.crosshair&&a.series.yAxis.drawCrosshair(null,a))}):d&&(d.setState(d.state,!0),f.axes.forEach(function(a){a.crosshair&&a.drawCrosshair(null,d)})));else{if(d)d.onMouseOut();l&&l.forEach(function(a){a.setState()});if(c)c.onMouseOut();g&&g.hide(b);this.unDocMouseMove&&(this.unDocMouseMove=this.unDocMouseMove());
f.axes.forEach(function(a){a.hideCrosshair()});this.hoverX=f.hoverPoints=f.hoverPoint=null}},scaleGroups:function(a,b){var f=this.chart,c;f.series.forEach(function(d){c=a||d.getPlotBox();d.xAxis&&d.xAxis.zoomEnabled&&d.group&&(d.group.attr(c),d.markerGroup&&(d.markerGroup.attr(c),d.markerGroup.clip(b?f.clipRect:null)),d.dataLabelsGroup&&d.dataLabelsGroup.attr(c))});f.clipRect.attr(b||f.clipBox)},dragStart:function(a){var b=this.chart;b.mouseIsDown=a.type;b.cancelClick=!1;b.mouseDownX=this.mouseDownX=
a.chartX;b.mouseDownY=this.mouseDownY=a.chartY},drag:function(a){var b=this.chart,c=b.options.chart,d=a.chartX,l=a.chartY,g=this.zoomHor,p=this.zoomVert,k=b.plotLeft,m=b.plotTop,D=b.plotWidth,C=b.plotHeight,e,q=this.selectionMarker,n=this.mouseDownX,r=this.mouseDownY,u=c.panKey&&a[c.panKey+"Key"];q&&q.touch||(d<k?d=k:d>k+D&&(d=k+D),l<m?l=m:l>m+C&&(l=m+C),this.hasDragged=Math.sqrt(Math.pow(n-d,2)+Math.pow(r-l,2)),10<this.hasDragged&&(e=b.isInsidePlot(n-k,r-m),b.hasCartesianSeries&&(this.zoomX||this.zoomY)&&
e&&!u&&!q&&(this.selectionMarker=q=b.renderer.rect(k,m,g?1:D,p?1:C,0).attr({"class":"highcharts-selection-marker",zIndex:7}).add(),b.styledMode||q.attr({fill:c.selectionMarkerFill||h("#335cad").setOpacity(.25).get()})),q&&g&&(d-=n,q.attr({width:Math.abs(d),x:(0<d?0:d)+n})),q&&p&&(d=l-r,q.attr({height:Math.abs(d),y:(0<d?0:d)+r})),e&&!q&&c.panning&&b.pan(a,c.panning)))},drop:function(a){var b=this,c=this.chart,l=this.hasPinched;if(this.selectionMarker){var g={originalEvent:a,xAxis:[],yAxis:[]},p=this.selectionMarker,
m=p.attr?p.attr("x"):p.x,k=p.attr?p.attr("y"):p.y,A=p.attr?p.attr("width"):p.width,D=p.attr?p.attr("height"):p.height,h;if(this.hasDragged||l)c.axes.forEach(function(e){if(e.zoomEnabled&&r(e.min)&&(l||b[{xAxis:"zoomX",yAxis:"zoomY"}[e.coll]])){var f=e.horiz,c="touchend"===a.type?e.minPixelPadding:0,d=e.toValue((f?m:k)+c),f=e.toValue((f?m+A:k+D)-c);g[e.coll].push({axis:e,min:Math.min(d,f),max:Math.max(d,f)});h=!0}}),h&&w(c,"selection",g,function(a){c.zoom(u(a,l?{animation:!1}:null))});n(c.index)&&
(this.selectionMarker=this.selectionMarker.destroy());l&&this.scaleGroups()}c&&n(c.index)&&(d(c.container,{cursor:c._cursor}),c.cancelClick=10<this.hasDragged,c.mouseIsDown=this.hasDragged=this.hasPinched=!1,this.pinchDown=[])},onContainerMouseDown:function(a){a=this.normalize(a);2!==a.button&&(this.zoomOption(a),a.preventDefault&&a.preventDefault(),this.dragStart(a))},onDocumentMouseUp:function(b){E[a.hoverChartIndex]&&E[a.hoverChartIndex].pointer.drop(b)},onDocumentMouseMove:function(a){var b=this.chart,
c=this.chartPosition;a=this.normalize(a,c);!c||this.inClass(a.target,"highcharts-tracker")||b.isInsidePlot(a.chartX-b.plotLeft,a.chartY-b.plotTop)||this.reset()},onContainerMouseLeave:function(b){var f=E[a.hoverChartIndex];f&&(b.relatedTarget||b.toElement)&&(f.pointer.reset(),f.pointer.chartPosition=null)},onContainerMouseMove:function(b){var f=this.chart;r(a.hoverChartIndex)&&E[a.hoverChartIndex]&&E[a.hoverChartIndex].mouseIsDown||(a.hoverChartIndex=f.index);b=this.normalize(b);b.returnValue=!1;
"mousedown"===f.mouseIsDown&&this.drag(b);!this.inClass(b.target,"highcharts-tracker")&&!f.isInsidePlot(b.chartX-f.plotLeft,b.chartY-f.plotTop)||f.openMenu||this.runPointActions(b)},inClass:function(a,b){for(var f;a;){if(f=G(a,"class")){if(-1!==f.indexOf(b))return!0;if(-1!==f.indexOf("highcharts-container"))return!1}a=a.parentNode}},onTrackerMouseOut:function(a){var b=this.chart.hoverSeries;a=a.relatedTarget||a.toElement;this.isDirectTouch=!1;if(!(!b||!a||b.stickyTracking||this.inClass(a,"highcharts-tooltip")||
this.inClass(a,"highcharts-series-"+b.index)&&this.inClass(a,"highcharts-tracker")))b.onMouseOut()},onContainerClick:function(a){var b=this.chart,c=b.hoverPoint,d=b.plotLeft,l=b.plotTop;a=this.normalize(a);b.cancelClick||(c&&this.inClass(a.target,"highcharts-tracker")?(w(c.series,"click",u(a,{point:c})),b.hoverPoint&&c.firePointEvent("click",a)):(u(a,this.getCoordinates(a)),b.isInsidePlot(a.chartX-d,a.chartY-l)&&w(b,"click",a)))},setDOMEvents:function(){var b=this,f=b.chart.container,c=f.ownerDocument;
f.onmousedown=function(a){b.onContainerMouseDown(a)};f.onmousemove=function(a){b.onContainerMouseMove(a)};f.onclick=function(a){b.onContainerClick(a)};this.unbindContainerMouseLeave=y(f,"mouseleave",b.onContainerMouseLeave);a.unbindDocumentMouseUp||(a.unbindDocumentMouseUp=y(c,"mouseup",b.onDocumentMouseUp));a.hasTouch&&(f.ontouchstart=function(a){b.onContainerTouchStart(a)},f.ontouchmove=function(a){b.onContainerTouchMove(a)},a.unbindDocumentTouchEnd||(a.unbindDocumentTouchEnd=y(c,"touchend",b.onDocumentTouchEnd)))},
destroy:function(){var b=this;b.unDocMouseMove&&b.unDocMouseMove();this.unbindContainerMouseLeave();a.chartCount||(a.unbindDocumentMouseUp&&(a.unbindDocumentMouseUp=a.unbindDocumentMouseUp()),a.unbindDocumentTouchEnd&&(a.unbindDocumentTouchEnd=a.unbindDocumentTouchEnd()));clearInterval(b.tooltipTimeout);a.objectEach(b,function(a,c){b[c]=null})}}})(J);(function(a){var y=a.charts,G=a.extend,E=a.noop,h=a.pick;G(a.Pointer.prototype,{pinchTranslate:function(a,h,u,v,w,n){this.zoomHor&&this.pinchTranslateDirection(!0,
a,h,u,v,w,n);this.zoomVert&&this.pinchTranslateDirection(!1,a,h,u,v,w,n)},pinchTranslateDirection:function(a,h,u,v,w,n,g,c){var d=this.chart,p=a?"x":"y",b=a?"X":"Y",l="chart"+b,f=a?"width":"height",x=d["plot"+(a?"Left":"Top")],t,r,F=c||1,z=d.inverted,k=d.bounds[a?"h":"v"],A=1===h.length,D=h[0][l],C=u[0][l],e=!A&&h[1][l],q=!A&&u[1][l],L;u=function(){!A&&20<Math.abs(D-e)&&(F=c||Math.abs(C-q)/Math.abs(D-e));r=(x-C)/F+D;t=d["plot"+(a?"Width":"Height")]/F};u();h=r;h<k.min?(h=k.min,L=!0):h+t>k.max&&(h=
k.max-t,L=!0);L?(C-=.8*(C-g[p][0]),A||(q-=.8*(q-g[p][1])),u()):g[p]=[C,q];z||(n[p]=r-x,n[f]=t);n=z?1/F:F;w[f]=t;w[p]=h;v[z?a?"scaleY":"scaleX":"scale"+b]=F;v["translate"+b]=n*x+(C-n*D)},pinch:function(a){var d=this,u=d.chart,v=d.pinchDown,w=a.touches,n=w.length,g=d.lastValidTouch,c=d.hasZoom,m=d.selectionMarker,p={},b=1===n&&(d.inClass(a.target,"highcharts-tracker")&&u.runTrackerClick||d.runChartClick),l={};1<n&&(d.initiated=!0);c&&d.initiated&&!b&&a.preventDefault();[].map.call(w,function(a){return d.normalize(a)});
"touchstart"===a.type?([].forEach.call(w,function(a,b){v[b]={chartX:a.chartX,chartY:a.chartY}}),g.x=[v[0].chartX,v[1]&&v[1].chartX],g.y=[v[0].chartY,v[1]&&v[1].chartY],u.axes.forEach(function(a){if(a.zoomEnabled){var b=u.bounds[a.horiz?"h":"v"],f=a.minPixelPadding,c=a.toPixels(h(a.options.min,a.dataMin)),d=a.toPixels(h(a.options.max,a.dataMax)),l=Math.max(c,d);b.min=Math.min(a.pos,Math.min(c,d)-f);b.max=Math.max(a.pos+a.len,l+f)}}),d.res=!0):d.followTouchMove&&1===n?this.runPointActions(d.normalize(a)):
v.length&&(m||(d.selectionMarker=m=G({destroy:E,touch:!0},u.plotBox)),d.pinchTranslate(v,w,p,m,l,g),d.hasPinched=c,d.scaleGroups(p,l),d.res&&(d.res=!1,this.reset(!1,0)))},touch:function(d,r){var u=this.chart,v,w;if(u.index!==a.hoverChartIndex)this.onContainerMouseLeave({relatedTarget:!0});a.hoverChartIndex=u.index;1===d.touches.length?(d=this.normalize(d),(w=u.isInsidePlot(d.chartX-u.plotLeft,d.chartY-u.plotTop))&&!u.openMenu?(r&&this.runPointActions(d),"touchmove"===d.type&&(r=this.pinchDown,v=r[0]?
4<=Math.sqrt(Math.pow(r[0].chartX-d.chartX,2)+Math.pow(r[0].chartY-d.chartY,2)):!1),h(v,!0)&&this.pinch(d)):r&&this.reset()):2===d.touches.length&&this.pinch(d)},onContainerTouchStart:function(a){this.zoomOption(a);this.touch(a,!0)},onContainerTouchMove:function(a){this.touch(a)},onDocumentTouchEnd:function(d){y[a.hoverChartIndex]&&y[a.hoverChartIndex].pointer.drop(d)}})})(J);(function(a){var y=a.addEvent,G=a.charts,E=a.css,h=a.doc,d=a.extend,r=a.noop,u=a.Pointer,v=a.removeEvent,w=a.win,n=a.wrap;
if(!a.hasTouch&&(w.PointerEvent||w.MSPointerEvent)){var g={},c=!!w.PointerEvent,m=function(){var b=[];b.item=function(a){return this[a]};a.objectEach(g,function(a){b.push({pageX:a.pageX,pageY:a.pageY,target:a.target})});return b},p=function(b,c,f,d){"touch"!==b.pointerType&&b.pointerType!==b.MSPOINTER_TYPE_TOUCH||!G[a.hoverChartIndex]||(d(b),d=G[a.hoverChartIndex].pointer,d[c]({type:f,target:b.currentTarget,preventDefault:r,touches:m()}))};d(u.prototype,{onContainerPointerDown:function(a){p(a,"onContainerTouchStart",
"touchstart",function(a){g[a.pointerId]={pageX:a.pageX,pageY:a.pageY,target:a.currentTarget}})},onContainerPointerMove:function(a){p(a,"onContainerTouchMove","touchmove",function(a){g[a.pointerId]={pageX:a.pageX,pageY:a.pageY};g[a.pointerId].target||(g[a.pointerId].target=a.currentTarget)})},onDocumentPointerUp:function(a){p(a,"onDocumentTouchEnd","touchend",function(a){delete g[a.pointerId]})},batchMSEvents:function(a){a(this.chart.container,c?"pointerdown":"MSPointerDown",this.onContainerPointerDown);
a(this.chart.container,c?"pointermove":"MSPointerMove",this.onContainerPointerMove);a(h,c?"pointerup":"MSPointerUp",this.onDocumentPointerUp)}});n(u.prototype,"init",function(a,c,f){a.call(this,c,f);this.hasZoom&&E(c.container,{"-ms-touch-action":"none","touch-action":"none"})});n(u.prototype,"setDOMEvents",function(a){a.apply(this);(this.hasZoom||this.followTouchMove)&&this.batchMSEvents(y)});n(u.prototype,"destroy",function(a){this.batchMSEvents(v);a.call(this)})}})(J);(function(a){var y=a.addEvent,
G=a.css,E=a.discardElement,h=a.defined,d=a.fireEvent,r=a.isFirefox,u=a.marginNames,v=a.merge,w=a.pick,n=a.setAnimation,g=a.stableSort,c=a.win,m=a.wrap;a.Legend=function(a,b){this.init(a,b)};a.Legend.prototype={init:function(a,b){this.chart=a;this.setOptions(b);b.enabled&&(this.render(),y(this.chart,"endResize",function(){this.legend.positionCheckboxes()}),this.proximate?this.unchartrender=y(this.chart,"render",function(){this.legend.proximatePositions();this.legend.positionItems()}):this.unchartrender&&
this.unchartrender())},setOptions:function(a){var b=w(a.padding,8);this.options=a;this.chart.styledMode||(this.itemStyle=a.itemStyle,this.itemHiddenStyle=v(this.itemStyle,a.itemHiddenStyle));this.itemMarginTop=a.itemMarginTop||0;this.padding=b;this.initialItemY=b-5;this.symbolWidth=w(a.symbolWidth,16);this.pages=[];this.proximate="proximate"===a.layout&&!this.chart.inverted},update:function(a,b){var c=this.chart;this.setOptions(v(!0,this.options,a));this.destroy();c.isDirtyLegend=c.isDirtyBox=!0;
w(b,!0)&&c.redraw();d(this,"afterUpdate")},colorizeItem:function(a,b){a.legendGroup[b?"removeClass":"addClass"]("highcharts-legend-item-hidden");if(!this.chart.styledMode){var c=this.options,f=a.legendItem,g=a.legendLine,p=a.legendSymbol,m=this.itemHiddenStyle.color,c=b?c.itemStyle.color:m,h=b?a.color||m:m,n=a.options&&a.options.marker,k={fill:h};f&&f.css({fill:c,color:c});g&&g.attr({stroke:h});p&&(n&&p.isMarker&&(k=a.pointAttribs(),b||(k.stroke=k.fill=m)),p.attr(k))}d(this,"afterColorizeItem",{item:a,
visible:b})},positionItems:function(){this.allItems.forEach(this.positionItem,this);this.chart.isResizing||this.positionCheckboxes()},positionItem:function(a){var b=this.options,c=b.symbolPadding,b=!b.rtl,f=a._legendItemPos,d=f[0],f=f[1],g=a.checkbox;if((a=a.legendGroup)&&a.element)a[h(a.translateY)?"animate":"attr"]({translateX:b?d:this.legendWidth-d-2*c-4,translateY:f});g&&(g.x=d,g.y=f)},destroyItem:function(a){var b=a.checkbox;["legendItem","legendLine","legendSymbol","legendGroup"].forEach(function(b){a[b]&&
(a[b]=a[b].destroy())});b&&E(a.checkbox)},destroy:function(){function a(a){this[a]&&(this[a]=this[a].destroy())}this.getAllItems().forEach(function(b){["legendItem","legendGroup"].forEach(a,b)});"clipRect up down pager nav box title group".split(" ").forEach(a,this);this.display=null},positionCheckboxes:function(){var a=this.group&&this.group.alignAttr,b,c=this.clipHeight||this.legendHeight,f=this.titleHeight;a&&(b=a.translateY,this.allItems.forEach(function(d){var g=d.checkbox,l;g&&(l=b+f+g.y+(this.scrollOffset||
0)+3,G(g,{left:a.translateX+d.checkboxOffset+g.x-20+"px",top:l+"px",display:this.proximate||l>b-6&&l<b+c-6?"":"none"}))},this))},renderTitle:function(){var a=this.options,b=this.padding,c=a.title,f=0;c.text&&(this.title||(this.title=this.chart.renderer.label(c.text,b-3,b-4,null,null,null,a.useHTML,null,"legend-title").attr({zIndex:1}),this.chart.styledMode||this.title.css(c.style),this.title.add(this.group)),a=this.title.getBBox(),f=a.height,this.offsetWidth=a.width,this.contentGroup.attr({translateY:f}));
this.titleHeight=f},setText:function(c){var b=this.options;c.legendItem.attr({text:b.labelFormat?a.format(b.labelFormat,c,this.chart.time):b.labelFormatter.call(c)})},renderItem:function(a){var b=this.chart,c=b.renderer,f=this.options,d=this.symbolWidth,g=f.symbolPadding,p=this.itemStyle,m=this.itemHiddenStyle,h="horizontal"===f.layout?w(f.itemDistance,20):0,k=!f.rtl,A=a.legendItem,D=!a.series,n=!D&&a.series.drawLegendSymbol?a.series:a,e=n.options,e=this.createCheckboxForItem&&e&&e.showCheckbox,h=
d+g+h+(e?20:0),q=f.useHTML,r=a.options.className;A||(a.legendGroup=c.g("legend-item").addClass("highcharts-"+n.type+"-series highcharts-color-"+a.colorIndex+(r?" "+r:"")+(D?" highcharts-series-"+a.index:"")).attr({zIndex:1}).add(this.scrollGroup),a.legendItem=A=c.text("",k?d+g:-g,this.baseline||0,q),b.styledMode||A.css(v(a.visible?p:m)),A.attr({align:k?"left":"right",zIndex:2}).add(a.legendGroup),this.baseline||(this.fontMetrics=c.fontMetrics(b.styledMode?12:p.fontSize,A),this.baseline=this.fontMetrics.f+
3+this.itemMarginTop,A.attr("y",this.baseline)),this.symbolHeight=f.symbolHeight||this.fontMetrics.f,n.drawLegendSymbol(this,a),this.setItemEvents&&this.setItemEvents(a,A,q),e&&this.createCheckboxForItem(a));this.colorizeItem(a,a.visible);!b.styledMode&&p.width||A.css({width:(f.itemWidth||f.width||b.spacingBox.width)-h});this.setText(a);b=A.getBBox();a.itemWidth=a.checkboxOffset=f.itemWidth||a.legendItemWidth||b.width+h;this.maxItemWidth=Math.max(this.maxItemWidth,a.itemWidth);this.totalItemWidth+=
a.itemWidth;this.itemHeight=a.itemHeight=Math.round(a.legendItemHeight||b.height||this.symbolHeight)},layoutItem:function(a){var b=this.options,c=this.padding,f="horizontal"===b.layout,d=a.itemHeight,g=b.itemMarginBottom||0,p=this.itemMarginTop,m=f?w(b.itemDistance,20):0,h=b.width,k=h||this.chart.spacingBox.width-2*c-b.x,b=b.alignColumns&&this.totalItemWidth>k?this.maxItemWidth:a.itemWidth;f&&this.itemX-c+b>k&&(this.itemX=c,this.itemY+=p+this.lastLineHeight+g,this.lastLineHeight=0);this.lastItemY=
p+this.itemY+g;this.lastLineHeight=Math.max(d,this.lastLineHeight);a._legendItemPos=[this.itemX,this.itemY];f?this.itemX+=b:(this.itemY+=p+d+g,this.lastLineHeight=d);this.offsetWidth=h||Math.max((f?this.itemX-c-(a.checkbox?0:m):b)+c,this.offsetWidth)},getAllItems:function(){var a=[];this.chart.series.forEach(function(b){var c=b&&b.options;b&&w(c.showInLegend,h(c.linkedTo)?!1:void 0,!0)&&(a=a.concat(b.legendItems||("point"===c.legendType?b.data:b)))});d(this,"afterGetAllItems",{allItems:a});return a},
getAlignment:function(){var a=this.options;return this.proximate?a.align.charAt(0)+"tv":a.floating?"":a.align.charAt(0)+a.verticalAlign.charAt(0)+a.layout.charAt(0)},adjustMargins:function(a,b){var c=this.chart,f=this.options,d=this.getAlignment();d&&[/(lth|ct|rth)/,/(rtv|rm|rbv)/,/(rbh|cb|lbh)/,/(lbv|lm|ltv)/].forEach(function(g,l){g.test(d)&&!h(a[l])&&(c[u[l]]=Math.max(c[u[l]],c.legend[(l+1)%2?"legendHeight":"legendWidth"]+[1,-1,-1,1][l]*f[l%2?"x":"y"]+w(f.margin,12)+b[l]+(0===l&&void 0!==c.options.title.margin?
c.titleOffset+c.options.title.margin:0)))})},proximatePositions:function(){var c=this.chart,b=[],d="left"===this.options.align;this.allItems.forEach(function(f){var g,l;g=d;f.xAxis&&f.points&&(f.xAxis.options.reversed&&(g=!g),g=a.find(g?f.points:f.points.slice(0).reverse(),function(b){return a.isNumber(b.plotY)}),l=f.legendGroup.getBBox().height,b.push({target:f.visible?(g?g.plotY:f.xAxis.height)-.3*l:c.plotHeight,size:l,item:f}))},this);a.distribute(b,c.plotHeight);b.forEach(function(a){a.item._legendItemPos[1]=
c.plotTop-c.spacing[0]+a.pos})},render:function(){var a=this.chart,b=a.renderer,c=this.group,f,d,m,h=this.box,n=this.options,z=this.padding;this.itemX=z;this.itemY=this.initialItemY;this.lastItemY=this.offsetWidth=0;c||(this.group=c=b.g("legend").attr({zIndex:7}).add(),this.contentGroup=b.g().attr({zIndex:1}).add(c),this.scrollGroup=b.g().add(this.contentGroup));this.renderTitle();f=this.getAllItems();g(f,function(a,b){return(a.options&&a.options.legendIndex||0)-(b.options&&b.options.legendIndex||
0)});n.reversed&&f.reverse();this.allItems=f;this.display=d=!!f.length;this.itemHeight=this.totalItemWidth=this.maxItemWidth=this.lastLineHeight=0;f.forEach(this.renderItem,this);f.forEach(this.layoutItem,this);f=(n.width||this.offsetWidth)+z;m=this.lastItemY+this.lastLineHeight+this.titleHeight;m=this.handleOverflow(m);m+=z;h||(this.box=h=b.rect().addClass("highcharts-legend-box").attr({r:n.borderRadius}).add(c),h.isNew=!0);a.styledMode||h.attr({stroke:n.borderColor,"stroke-width":n.borderWidth||
0,fill:n.backgroundColor||"none"}).shadow(n.shadow);0<f&&0<m&&(h[h.isNew?"attr":"animate"](h.crisp.call({},{x:0,y:0,width:f,height:m},h.strokeWidth())),h.isNew=!1);h[d?"show":"hide"]();a.styledMode&&"none"===c.getStyle("display")&&(f=m=0);this.legendWidth=f;this.legendHeight=m;d&&(b=a.spacingBox,/(lth|ct|rth)/.test(this.getAlignment())&&(b=v(b,{y:b.y+a.titleOffset+a.options.title.margin})),c.align(v(n,{width:f,height:m,verticalAlign:this.proximate?"top":n.verticalAlign}),!0,b));this.proximate||this.positionItems()},
handleOverflow:function(a){var b=this,c=this.chart,f=c.renderer,d=this.options,g=d.y,m=this.padding,g=c.spacingBox.height+("top"===d.verticalAlign?-g:g)-m,p=d.maxHeight,h,k=this.clipRect,A=d.navigation,D=w(A.animation,!0),n=A.arrowSize||12,e=this.nav,q=this.pages,r,u=this.allItems,v=function(a){"number"===typeof a?k.attr({height:a}):k&&(b.clipRect=k.destroy(),b.contentGroup.clip());b.contentGroup.div&&(b.contentGroup.div.style.clip=a?"rect("+m+"px,9999px,"+(m+a)+"px,0)":"auto")};"horizontal"!==d.layout||
"middle"===d.verticalAlign||d.floating||(g/=2);p&&(g=Math.min(g,p));q.length=0;a>g&&!1!==A.enabled?(this.clipHeight=h=Math.max(g-20-this.titleHeight-m,0),this.currentPage=w(this.currentPage,1),this.fullHeight=a,u.forEach(function(a,b){var e=a._legendItemPos[1],c=Math.round(a.legendItem.getBBox().height),f=q.length;if(!f||e-q[f-1]>h&&(r||e)!==q[f-1])q.push(r||e),f++;a.pageIx=f-1;r&&(u[b-1].pageIx=f-1);b===u.length-1&&e+c-q[f-1]>h&&e!==r&&(q.push(e),a.pageIx=f);e!==r&&(r=e)}),k||(k=b.clipRect=f.clipRect(0,
m,9999,0),b.contentGroup.clip(k)),v(h),e||(this.nav=e=f.g().attr({zIndex:1}).add(this.group),this.up=f.symbol("triangle",0,0,n,n).on("click",function(){b.scroll(-1,D)}).add(e),this.pager=f.text("",15,10).addClass("highcharts-legend-navigation"),c.styledMode||this.pager.css(A.style),this.pager.add(e),this.down=f.symbol("triangle-down",0,0,n,n).on("click",function(){b.scroll(1,D)}).add(e)),b.scroll(0),a=g):e&&(v(),this.nav=e.destroy(),this.scrollGroup.attr({translateY:1}),this.clipHeight=0);return a},
scroll:function(a,b){var c=this.pages,f=c.length;a=this.currentPage+a;var d=this.clipHeight,g=this.options.navigation,m=this.pager,p=this.padding;a>f&&(a=f);0<a&&(void 0!==b&&n(b,this.chart),this.nav.attr({translateX:p,translateY:d+this.padding+7+this.titleHeight,visibility:"visible"}),this.up.attr({"class":1===a?"highcharts-legend-nav-inactive":"highcharts-legend-nav-active"}),m.attr({text:a+"/"+f}),this.down.attr({x:18+this.pager.getBBox().width,"class":a===f?"highcharts-legend-nav-inactive":"highcharts-legend-nav-active"}),
this.chart.styledMode||(this.up.attr({fill:1===a?g.inactiveColor:g.activeColor}).css({cursor:1===a?"default":"pointer"}),this.down.attr({fill:a===f?g.inactiveColor:g.activeColor}).css({cursor:a===f?"default":"pointer"})),this.scrollOffset=-c[a-1]+this.initialItemY,this.scrollGroup.animate({translateY:this.scrollOffset}),this.currentPage=a,this.positionCheckboxes())}};a.LegendSymbolMixin={drawRectangle:function(a,b){var c=a.symbolHeight,f=a.options.squareSymbol;b.legendSymbol=this.chart.renderer.rect(f?
(a.symbolWidth-c)/2:0,a.baseline-c+1,f?c:a.symbolWidth,c,w(a.options.symbolRadius,c/2)).addClass("highcharts-point").attr({zIndex:3}).add(b.legendGroup)},drawLineMarker:function(a){var b=this.options,c=b.marker,f=a.symbolWidth,d=a.symbolHeight,g=d/2,m=this.chart.renderer,p=this.legendGroup;a=a.baseline-Math.round(.3*a.fontMetrics.b);var h={};this.chart.styledMode||(h={"stroke-width":b.lineWidth||0},b.dashStyle&&(h.dashstyle=b.dashStyle));this.legendLine=m.path(["M",0,a,"L",f,a]).addClass("highcharts-graph").attr(h).add(p);
c&&!1!==c.enabled&&f&&(b=Math.min(w(c.radius,g),g),0===this.symbol.indexOf("url")&&(c=v(c,{width:d,height:d}),b=0),this.legendSymbol=c=m.symbol(this.symbol,f/2-b,a-b,2*b,2*b,c).addClass("highcharts-point").add(p),c.isMarker=!0)}};(/Trident\/7\.0/.test(c.navigator.userAgent)||r)&&m(a.Legend.prototype,"positionItem",function(a,b){var c=this,f=function(){b._legendItemPos&&a.call(c,b)};f();c.bubbleLegend||setTimeout(f)})})(J);(function(a){var y=a.addEvent,G=a.animate,E=a.animObject,h=a.attr,d=a.doc,r=
a.Axis,u=a.createElement,v=a.defaultOptions,w=a.discardElement,n=a.charts,g=a.css,c=a.defined,m=a.extend,p=a.find,b=a.fireEvent,l=a.isNumber,f=a.isObject,x=a.isString,t=a.Legend,H=a.marginNames,F=a.merge,z=a.objectEach,k=a.Pointer,A=a.pick,D=a.pInt,C=a.removeEvent,e=a.seriesTypes,q=a.splat,L=a.syncTimeout,I=a.win,R=a.Chart=function(){this.getArgs.apply(this,arguments)};a.chart=function(a,b,e){return new R(a,b,e)};m(R.prototype,{callbacks:[],getArgs:function(){var a=[].slice.call(arguments);if(x(a[0])||
a[0].nodeName)this.renderTo=a.shift();this.init(a[0],a[1])},init:function(e,c){var f,k,d=e.series,g=e.plotOptions||{};b(this,"init",{args:arguments},function(){e.series=null;f=F(v,e);for(k in f.plotOptions)f.plotOptions[k].tooltip=g[k]&&F(g[k].tooltip)||void 0;f.tooltip.userOptions=e.chart&&e.chart.forExport&&e.tooltip.userOptions||e.tooltip;f.series=e.series=d;this.userOptions=e;var q=f.chart,l=q.events;this.margin=[];this.spacing=[];this.bounds={h:{},v:{}};this.labelCollectors=[];this.callback=
c;this.isResizing=0;this.options=f;this.axes=[];this.series=[];this.time=e.time&&Object.keys(e.time).length?new a.Time(e.time):a.time;this.styledMode=q.styledMode;this.hasCartesianSeries=q.showAxes;var m=this;m.index=n.length;n.push(m);a.chartCount++;l&&z(l,function(a,b){y(m,b,a)});m.xAxis=[];m.yAxis=[];m.pointCount=m.colorCounter=m.symbolCounter=0;b(m,"afterInit");m.firstRender()})},initSeries:function(b){var c=this.options.chart;(c=e[b.type||c.type||c.defaultSeriesType])||a.error(17,!0,this);c=
new c;c.init(this,b);return c},orderSeries:function(a){var b=this.series;for(a=a||0;a<b.length;a++)b[a]&&(b[a].index=a,b[a].name=b[a].getName())},isInsidePlot:function(a,b,e){var c=e?b:a;a=e?a:b;return 0<=c&&c<=this.plotWidth&&0<=a&&a<=this.plotHeight},redraw:function(e){b(this,"beforeRedraw");var c=this.axes,f=this.series,k=this.pointer,d=this.legend,g=this.userOptions.legend,q=this.isDirtyLegend,l,p,A=this.hasCartesianSeries,h=this.isDirtyBox,D,t=this.renderer,n=t.isHidden(),C=[];this.setResponsive&&
this.setResponsive(!1);a.setAnimation(e,this);n&&this.temporaryDisplay();this.layOutTitles();for(e=f.length;e--;)if(D=f[e],D.options.stacking&&(l=!0,D.isDirty)){p=!0;break}if(p)for(e=f.length;e--;)D=f[e],D.options.stacking&&(D.isDirty=!0);f.forEach(function(a){a.isDirty&&("point"===a.options.legendType?(a.updateTotals&&a.updateTotals(),q=!0):g&&(g.labelFormatter||g.labelFormat)&&(q=!0));a.isDirtyData&&b(a,"updatedData")});q&&d&&d.options.enabled&&(d.render(),this.isDirtyLegend=!1);l&&this.getStacks();
A&&c.forEach(function(a){a.updateNames();a.updateYNames&&a.updateYNames();a.setScale()});this.getMargins();A&&(c.forEach(function(a){a.isDirty&&(h=!0)}),c.forEach(function(a){var e=a.min+","+a.max;a.extKey!==e&&(a.extKey=e,C.push(function(){b(a,"afterSetExtremes",m(a.eventArgs,a.getExtremes()));delete a.eventArgs}));(h||l)&&a.redraw()}));h&&this.drawChartBox();b(this,"predraw");f.forEach(function(a){(h||a.isDirty)&&a.visible&&a.redraw();a.isDirtyData=!1});k&&k.reset(!0);t.draw();b(this,"redraw");
b(this,"render");n&&this.temporaryDisplay(!0);C.forEach(function(a){a.call()})},get:function(a){function b(b){return b.id===a||b.options&&b.options.id===a}var e,c=this.series,f;e=p(this.axes,b)||p(this.series,b);for(f=0;!e&&f<c.length;f++)e=p(c[f].points||[],b);return e},getAxes:function(){var a=this,e=this.options,c=e.xAxis=q(e.xAxis||{}),e=e.yAxis=q(e.yAxis||{});b(this,"getAxes");c.forEach(function(a,b){a.index=b;a.isX=!0});e.forEach(function(a,b){a.index=b});c.concat(e).forEach(function(b){new r(a,
b)});b(this,"afterGetAxes")},getSelectedPoints:function(){var a=[];this.series.forEach(function(b){a=a.concat((b.data||[]).filter(function(a){return a.selected}))});return a},getSelectedSeries:function(){return this.series.filter(function(a){return a.selected})},setTitle:function(a,b,e){var c=this,f=c.options,k=c.styledMode,d;d=f.title=F(!k&&{style:{color:"#333333",fontSize:f.isStock?"16px":"18px"}},f.title,a);f=f.subtitle=F(!k&&{style:{color:"#666666"}},f.subtitle,b);[["title",a,d],["subtitle",b,
f]].forEach(function(a,b){var e=a[0],f=c[e],d=a[1];a=a[2];f&&d&&(c[e]=f=f.destroy());a&&!f&&(c[e]=c.renderer.text(a.text,0,0,a.useHTML).attr({align:a.align,"class":"highcharts-"+e,zIndex:a.zIndex||4}).add(),c[e].update=function(a){c.setTitle(!b&&a,b&&a)},k||c[e].css(a.style))});c.layOutTitles(e)},layOutTitles:function(a){var b=0,e,c=this.renderer,f=this.spacingBox;["title","subtitle"].forEach(function(a){var e=this[a],k=this.options[a];a="title"===a?-3:k.verticalAlign?0:b+2;var d;e&&(this.styledMode||
(d=k.style.fontSize),d=c.fontMetrics(d,e).b,e.css({width:(k.width||f.width+k.widthAdjust)+"px"}).align(m({y:a+d},k),!1,"spacingBox"),k.floating||k.verticalAlign||(b=Math.ceil(b+e.getBBox(k.useHTML).height)))},this);e=this.titleOffset!==b;this.titleOffset=b;!this.isDirtyBox&&e&&(this.isDirtyBox=this.isDirtyLegend=e,this.hasRendered&&A(a,!0)&&this.isDirtyBox&&this.redraw())},getChartSize:function(){var b=this.options.chart,e=b.width,b=b.height,f=this.renderTo;c(e)||(this.containerWidth=a.getStyle(f,
"width"));c(b)||(this.containerHeight=a.getStyle(f,"height"));this.chartWidth=Math.max(0,e||this.containerWidth||600);this.chartHeight=Math.max(0,a.relativeLength(b,this.chartWidth)||(1<this.containerHeight?this.containerHeight:400))},temporaryDisplay:function(b){var e=this.renderTo;if(b)for(;e&&e.style;)e.hcOrigStyle&&(a.css(e,e.hcOrigStyle),delete e.hcOrigStyle),e.hcOrigDetached&&(d.body.removeChild(e),e.hcOrigDetached=!1),e=e.parentNode;else for(;e&&e.style;){d.body.contains(e)||e.parentNode||
(e.hcOrigDetached=!0,d.body.appendChild(e));if("none"===a.getStyle(e,"display",!1)||e.hcOricDetached)e.hcOrigStyle={display:e.style.display,height:e.style.height,overflow:e.style.overflow},b={display:"block",overflow:"hidden"},e!==this.renderTo&&(b.height=0),a.css(e,b),e.offsetWidth||e.style.setProperty("display","block","important");e=e.parentNode;if(e===d.body)break}},setClassName:function(a){this.container.className="highcharts-container "+(a||"")},getContainer:function(){var e,c=this.options,
f=c.chart,k,q;e=this.renderTo;var p=a.uniqueKey(),A,t;e||(this.renderTo=e=f.renderTo);x(e)&&(this.renderTo=e=d.getElementById(e));e||a.error(13,!0,this);k=D(h(e,"data-highcharts-chart"));l(k)&&n[k]&&n[k].hasRendered&&n[k].destroy();h(e,"data-highcharts-chart",this.index);e.innerHTML="";f.skipClone||e.offsetWidth||this.temporaryDisplay();this.getChartSize();k=this.chartWidth;q=this.chartHeight;g(e,{overflow:"hidden"});this.styledMode||(A=m({position:"relative",overflow:"hidden",width:k+"px",height:q+
"px",textAlign:"left",lineHeight:"normal",zIndex:0,"-webkit-tap-highlight-color":"rgba(0,0,0,0)"},f.style));this.container=e=u("div",{id:p},A,e);this._cursor=e.style.cursor;this.renderer=new (a[f.renderer]||a.Renderer)(e,k,q,null,f.forExport,c.exporting&&c.exporting.allowHTML,this.styledMode);this.setClassName(f.className);if(this.styledMode)for(t in c.defs)this.renderer.definition(c.defs[t]);else this.renderer.setStyle(f.style);this.renderer.chartIndex=this.index;b(this,"afterGetContainer")},getMargins:function(a){var e=
this.spacing,f=this.margin,k=this.titleOffset;this.resetMargins();k&&!c(f[0])&&(this.plotTop=Math.max(this.plotTop,k+this.options.title.margin+e[0]));this.legend&&this.legend.display&&this.legend.adjustMargins(f,e);b(this,"getMargins");a||this.getAxisMargins()},getAxisMargins:function(){var a=this,b=a.axisOffset=[0,0,0,0],e=a.margin;a.hasCartesianSeries&&a.axes.forEach(function(a){a.visible&&a.getOffset()});H.forEach(function(f,k){c(e[k])||(a[f]+=b[k])});a.setChartSize()},reflow:function(b){var e=
this,f=e.options.chart,k=e.renderTo,g=c(f.width)&&c(f.height),q=f.width||a.getStyle(k,"width"),f=f.height||a.getStyle(k,"height"),k=b?b.target:I;if(!g&&!e.isPrinting&&q&&f&&(k===I||k===d)){if(q!==e.containerWidth||f!==e.containerHeight)a.clearTimeout(e.reflowTimeout),e.reflowTimeout=L(function(){e.container&&e.setSize(void 0,void 0,!1)},b?100:0);e.containerWidth=q;e.containerHeight=f}},setReflow:function(a){var b=this;!1===a||this.unbindReflow?!1===a&&this.unbindReflow&&(this.unbindReflow=this.unbindReflow()):
(this.unbindReflow=y(I,"resize",function(a){b.reflow(a)}),y(this,"destroy",this.unbindReflow))},setSize:function(e,c,f){var k=this,d=k.renderer,q;k.isResizing+=1;a.setAnimation(f,k);k.oldChartHeight=k.chartHeight;k.oldChartWidth=k.chartWidth;void 0!==e&&(k.options.chart.width=e);void 0!==c&&(k.options.chart.height=c);k.getChartSize();k.styledMode||(q=d.globalAnimation,(q?G:g)(k.container,{width:k.chartWidth+"px",height:k.chartHeight+"px"},q));k.setChartSize(!0);d.setSize(k.chartWidth,k.chartHeight,
f);k.axes.forEach(function(a){a.isDirty=!0;a.setScale()});k.isDirtyLegend=!0;k.isDirtyBox=!0;k.layOutTitles();k.getMargins();k.redraw(f);k.oldChartHeight=null;b(k,"resize");L(function(){k&&b(k,"endResize",null,function(){--k.isResizing})},E(q).duration)},setChartSize:function(a){var e=this.inverted,c=this.renderer,f=this.chartWidth,k=this.chartHeight,d=this.options.chart,g=this.spacing,q=this.clipOffset,l,m,p,A;this.plotLeft=l=Math.round(this.plotLeft);this.plotTop=m=Math.round(this.plotTop);this.plotWidth=
p=Math.max(0,Math.round(f-l-this.marginRight));this.plotHeight=A=Math.max(0,Math.round(k-m-this.marginBottom));this.plotSizeX=e?A:p;this.plotSizeY=e?p:A;this.plotBorderWidth=d.plotBorderWidth||0;this.spacingBox=c.spacingBox={x:g[3],y:g[0],width:f-g[3]-g[1],height:k-g[0]-g[2]};this.plotBox=c.plotBox={x:l,y:m,width:p,height:A};f=2*Math.floor(this.plotBorderWidth/2);e=Math.ceil(Math.max(f,q[3])/2);c=Math.ceil(Math.max(f,q[0])/2);this.clipBox={x:e,y:c,width:Math.floor(this.plotSizeX-Math.max(f,q[1])/
2-e),height:Math.max(0,Math.floor(this.plotSizeY-Math.max(f,q[2])/2-c))};a||this.axes.forEach(function(a){a.setAxisSize();a.setAxisTranslation()});b(this,"afterSetChartSize",{skipAxes:a})},resetMargins:function(){b(this,"resetMargins");var a=this,e=a.options.chart;["margin","spacing"].forEach(function(b){var c=e[b],k=f(c)?c:[c,c,c,c];["Top","Right","Bottom","Left"].forEach(function(c,f){a[b][f]=A(e[b+c],k[f])})});H.forEach(function(b,e){a[b]=A(a.margin[e],a.spacing[e])});a.axisOffset=[0,0,0,0];a.clipOffset=
[0,0,0,0]},drawChartBox:function(){var a=this.options.chart,e=this.renderer,c=this.chartWidth,f=this.chartHeight,k=this.chartBackground,d=this.plotBackground,g=this.plotBorder,q,l=this.styledMode,m=this.plotBGImage,p=a.backgroundColor,A=a.plotBackgroundColor,h=a.plotBackgroundImage,D,t=this.plotLeft,n=this.plotTop,C=this.plotWidth,x=this.plotHeight,r=this.plotBox,z=this.clipRect,u=this.clipBox,v="animate";k||(this.chartBackground=k=e.rect().addClass("highcharts-background").add(),v="attr");if(l)q=
D=k.strokeWidth();else{q=a.borderWidth||0;D=q+(a.shadow?8:0);p={fill:p||"none"};if(q||k["stroke-width"])p.stroke=a.borderColor,p["stroke-width"]=q;k.attr(p).shadow(a.shadow)}k[v]({x:D/2,y:D/2,width:c-D-q%2,height:f-D-q%2,r:a.borderRadius});v="animate";d||(v="attr",this.plotBackground=d=e.rect().addClass("highcharts-plot-background").add());d[v](r);l||(d.attr({fill:A||"none"}).shadow(a.plotShadow),h&&(m?m.animate(r):this.plotBGImage=e.image(h,t,n,C,x).add()));z?z.animate({width:u.width,height:u.height}):
this.clipRect=e.clipRect(u);v="animate";g||(v="attr",this.plotBorder=g=e.rect().addClass("highcharts-plot-border").attr({zIndex:1}).add());l||g.attr({stroke:a.plotBorderColor,"stroke-width":a.plotBorderWidth||0,fill:"none"});g[v](g.crisp({x:t,y:n,width:C,height:x},-g.strokeWidth()));this.isDirtyBox=!1;b(this,"afterDrawChartBox")},propFromSeries:function(){var a=this,b=a.options.chart,c,f=a.options.series,k,d;["inverted","angular","polar"].forEach(function(g){c=e[b.type||b.defaultSeriesType];d=b[g]||
c&&c.prototype[g];for(k=f&&f.length;!d&&k--;)(c=e[f[k].type])&&c.prototype[g]&&(d=!0);a[g]=d})},linkSeries:function(){var a=this,e=a.series;e.forEach(function(a){a.linkedSeries.length=0});e.forEach(function(b){var e=b.options.linkedTo;x(e)&&(e=":previous"===e?a.series[b.index-1]:a.get(e))&&e.linkedParent!==b&&(e.linkedSeries.push(b),b.linkedParent=e,b.visible=A(b.options.visible,e.options.visible,b.visible))});b(this,"afterLinkSeries")},renderSeries:function(){this.series.forEach(function(a){a.translate();
a.render()})},renderLabels:function(){var a=this,b=a.options.labels;b.items&&b.items.forEach(function(e){var c=m(b.style,e.style),f=D(c.left)+a.plotLeft,k=D(c.top)+a.plotTop+12;delete c.left;delete c.top;a.renderer.text(e.html,f,k).attr({zIndex:2}).css(c).add()})},render:function(){var a=this.axes,b=this.renderer,e=this.options,c,f,k;this.setTitle();this.legend=new t(this,e.legend);this.getStacks&&this.getStacks();this.getMargins(!0);this.setChartSize();e=this.plotWidth;c=this.plotHeight=Math.max(this.plotHeight-
21,0);a.forEach(function(a){a.setScale()});this.getAxisMargins();f=1.1<e/this.plotWidth;k=1.05<c/this.plotHeight;if(f||k)a.forEach(function(a){(a.horiz&&f||!a.horiz&&k)&&a.setTickInterval(!0)}),this.getMargins();this.drawChartBox();this.hasCartesianSeries&&a.forEach(function(a){a.visible&&a.render()});this.seriesGroup||(this.seriesGroup=b.g("series-group").attr({zIndex:3}).add());this.renderSeries();this.renderLabels();this.addCredits();this.setResponsive&&this.setResponsive();this.hasRendered=!0},
addCredits:function(a){var b=this;a=F(!0,this.options.credits,a);a.enabled&&!this.credits&&(this.credits=this.renderer.text(a.text+(this.mapCredits||""),0,0).addClass("highcharts-credits").on("click",function(){a.href&&(I.location.href=a.href)}).attr({align:a.position.align,zIndex:8}),b.styledMode||this.credits.css(a.style),this.credits.add().align(a.position),this.credits.update=function(a){b.credits=b.credits.destroy();b.addCredits(a)})},destroy:function(){var e=this,c=e.axes,f=e.series,k=e.container,
d,g=k&&k.parentNode;b(e,"destroy");e.renderer.forExport?a.erase(n,e):n[e.index]=void 0;a.chartCount--;e.renderTo.removeAttribute("data-highcharts-chart");C(e);for(d=c.length;d--;)c[d]=c[d].destroy();this.scroller&&this.scroller.destroy&&this.scroller.destroy();for(d=f.length;d--;)f[d]=f[d].destroy();"title subtitle chartBackground plotBackground plotBGImage plotBorder seriesGroup clipRect credits pointer rangeSelector legend resetZoomButton tooltip renderer".split(" ").forEach(function(a){var b=e[a];
b&&b.destroy&&(e[a]=b.destroy())});k&&(k.innerHTML="",C(k),g&&w(k));z(e,function(a,b){delete e[b]})},firstRender:function(){var a=this,e=a.options;if(!a.isReadyToRender||a.isReadyToRender()){a.getContainer();a.resetMargins();a.setChartSize();a.propFromSeries();a.getAxes();(e.series||[]).forEach(function(b){a.initSeries(b)});a.linkSeries();b(a,"beforeRender");k&&(a.pointer=new k(a,e));a.render();if(!a.renderer.imgCount&&a.onload)a.onload();a.temporaryDisplay(!0)}},onload:function(){[this.callback].concat(this.callbacks).forEach(function(a){a&&
void 0!==this.index&&a.apply(this,[this])},this);b(this,"load");b(this,"render");c(this.index)&&this.setReflow(this.options.chart.reflow);this.onload=null}})})(J);(function(a){var y=a.addEvent,G=a.Chart;y(G,"afterSetChartSize",function(y){var h=this.options.chart.scrollablePlotArea;(h=h&&h.minWidth)&&!this.renderer.forExport&&(this.scrollablePixels=h=Math.max(0,h-this.chartWidth))&&(this.plotWidth+=h,this.clipBox.width+=h,y.skipAxes||this.axes.forEach(function(d){1===d.side?d.getPlotLinePath=function(){var h=
this.right,u;this.right=h-d.chart.scrollablePixels;u=a.Axis.prototype.getPlotLinePath.apply(this,arguments);this.right=h;return u}:(d.setAxisSize(),d.setAxisTranslation())}))});y(G,"render",function(){this.scrollablePixels?(this.setUpScrolling&&this.setUpScrolling(),this.applyFixed()):this.fixedDiv&&this.applyFixed()});G.prototype.setUpScrolling=function(){this.scrollingContainer=a.createElement("div",{className:"highcharts-scrolling"},{overflowX:"auto",WebkitOverflowScrolling:"touch"},this.renderTo);
this.innerContainer=a.createElement("div",{className:"highcharts-inner-container"},null,this.scrollingContainer);this.innerContainer.appendChild(this.container);this.setUpScrolling=null};G.prototype.applyFixed=function(){var y=this.container,h,d,r=!this.fixedDiv;r&&(this.fixedDiv=a.createElement("div",{className:"highcharts-fixed"},{position:"absolute",overflow:"hidden",pointerEvents:"none",zIndex:2},null,!0),this.renderTo.insertBefore(this.fixedDiv,this.renderTo.firstChild),this.fixedRenderer=h=
new a.Renderer(this.fixedDiv,0,0),this.scrollableMask=h.path().attr({fill:a.color(this.options.chart.backgroundColor||"#fff").setOpacity(.85).get(),zIndex:-1}).addClass("highcharts-scrollable-mask").add(),[this.inverted?".highcharts-xaxis":".highcharts-yaxis",this.inverted?".highcharts-xaxis-labels":".highcharts-yaxis-labels",".highcharts-contextbutton",".highcharts-credits",".highcharts-legend",".highcharts-subtitle",".highcharts-title",".highcharts-legend-checkbox"].forEach(function(a){[].forEach.call(y.querySelectorAll(a),
function(a){(a.namespaceURI===h.SVG_NS?h.box:h.box.parentNode).appendChild(a);a.style.pointerEvents="auto"})}));this.fixedRenderer.setSize(this.chartWidth,this.chartHeight);d=this.chartWidth+this.scrollablePixels;a.stop(this.container);this.container.style.width=d+"px";this.renderer.boxWrapper.attr({width:d,height:this.chartHeight,viewBox:[0,0,d,this.chartHeight].join(" ")});this.chartBackground.attr({width:d});r&&(d=this.options.chart.scrollablePlotArea,d.scrollPositionX&&(this.scrollingContainer.scrollLeft=
this.scrollablePixels*d.scrollPositionX));r=this.axisOffset;d=this.plotTop-r[0]-1;var r=this.plotTop+this.plotHeight+r[2],u=this.plotLeft+this.plotWidth-this.scrollablePixels;this.scrollableMask.attr({d:this.scrollablePixels?["M",0,d,"L",this.plotLeft-1,d,"L",this.plotLeft-1,r,"L",0,r,"Z","M",u,d,"L",this.chartWidth,d,"L",this.chartWidth,r,"L",u,r,"Z"]:["M",0,0]})}})(J);(function(a){var y,G=a.extend,E=a.erase,h=a.fireEvent,d=a.format,r=a.isArray,u=a.isNumber,v=a.pick,w=a.uniqueKey,n=a.defined,g=a.removeEvent;
a.Point=y=function(){};a.Point.prototype={init:function(a,d,g){var b;b=a.chart.options.chart.colorCount;var c=a.chart.styledMode;this.series=a;c||(this.color=a.color);this.applyOptions(d,g);this.id=n(this.id)?this.id:w();a.options.colorByPoint?(c||(b=a.options.colors||a.chart.options.colors,this.color=this.color||b[a.colorCounter],b=b.length),d=a.colorCounter,a.colorCounter++,a.colorCounter===b&&(a.colorCounter=0)):d=a.colorIndex;this.colorIndex=v(this.colorIndex,d);a.chart.pointCount++;h(this,"afterInit");
return this},applyOptions:function(a,d){var c=this.series,b=c.options.pointValKey||c.pointValKey;a=y.prototype.optionsToObject.call(this,a);G(this,a);this.options=this.options?G(this.options,a):a;a.group&&delete this.group;a.dataLabels&&delete this.dataLabels;b&&(this.y=this[b]);this.isNull=v(this.isValid&&!this.isValid(),null===this.x||!u(this.y,!0));this.selected&&(this.state="select");"name"in this&&void 0===d&&c.xAxis&&c.xAxis.hasNames&&(this.x=c.xAxis.nameToX(this));void 0===this.x&&c&&(this.x=
void 0===d?c.autoIncrement(this):d);return this},setNestedProperty:function(c,d,g){g.split(".").reduce(function(b,c,f,g){b[c]=g.length-1===f?d:a.isObject(b[c],!0)?b[c]:{};return b[c]},c);return c},optionsToObject:function(c){var d={},g=this.series,b=g.options.keys,l=b||g.pointArrayMap||["y"],f=l.length,h=0,t=0;if(u(c)||null===c)d[l[0]]=c;else if(r(c))for(!b&&c.length>f&&(g=typeof c[0],"string"===g?d.name=c[0]:"number"===g&&(d.x=c[0]),h++);t<f;)b&&void 0===c[h]||(0<l[t].indexOf(".")?a.Point.prototype.setNestedProperty(d,
c[h],l[t]):d[l[t]]=c[h]),h++,t++;else"object"===typeof c&&(d=c,c.dataLabels&&(g._hasPointLabels=!0),c.marker&&(g._hasPointMarkers=!0));return d},getClassName:function(){return"highcharts-point"+(this.selected?" highcharts-point-select":"")+(this.negative?" highcharts-negative":"")+(this.isNull?" highcharts-null-point":"")+(void 0!==this.colorIndex?" highcharts-color-"+this.colorIndex:"")+(this.options.className?" "+this.options.className:"")+(this.zone&&this.zone.className?" "+this.zone.className.replace("highcharts-negative",
""):"")},getZone:function(){var a=this.series,d=a.zones,a=a.zoneAxis||"y",g=0,b;for(b=d[g];this[a]>=b.value;)b=d[++g];this.nonZonedColor||(this.nonZonedColor=this.color);this.color=b&&b.color&&!this.options.color?b.color:this.nonZonedColor;return b},destroy:function(){var a=this.series.chart,d=a.hoverPoints,h;a.pointCount--;d&&(this.setState(),E(d,this),d.length||(a.hoverPoints=null));if(this===a.hoverPoint)this.onMouseOut();if(this.graphic||this.dataLabel||this.dataLabels)g(this),this.destroyElements();
this.legendItem&&a.legend.destroyItem(this);for(h in this)this[h]=null},destroyElements:function(){for(var a=["graphic","dataLabel","dataLabelUpper","connector","shadowGroup"],d,g=6;g--;)d=a[g],this[d]&&(this[d]=this[d].destroy());this.dataLabels&&(this.dataLabels.forEach(function(a){a.element&&a.destroy()}),delete this.dataLabels);this.connectors&&(this.connectors.forEach(function(a){a.element&&a.destroy()}),delete this.connectors)},getLabelConfig:function(){return{x:this.category,y:this.y,color:this.color,
colorIndex:this.colorIndex,key:this.name||this.category,series:this.series,point:this,percentage:this.percentage,total:this.total||this.stackTotal}},tooltipFormatter:function(a){var c=this.series,g=c.tooltipOptions,b=v(g.valueDecimals,""),l=g.valuePrefix||"",f=g.valueSuffix||"";c.chart.styledMode&&(a=c.chart.tooltip.styledModeFormat(a));(c.pointArrayMap||["y"]).forEach(function(c){c="{point."+c;if(l||f)a=a.replace(RegExp(c+"}","g"),l+c+"}"+f);a=a.replace(RegExp(c+"}","g"),c+":,."+b+"f}")});return d(a,
{point:this,series:this.series},c.chart.time)},firePointEvent:function(a,d,g){var b=this,c=this.series.options;(c.point.events[a]||b.options&&b.options.events&&b.options.events[a])&&this.importEvents();"click"===a&&c.allowPointSelect&&(g=function(a){b.select&&b.select(null,a.ctrlKey||a.metaKey||a.shiftKey)});h(this,a,d,g)},visible:!0}})(J);(function(a){var y=a.addEvent,G=a.animObject,E=a.arrayMax,h=a.arrayMin,d=a.correctFloat,r=a.defaultOptions,u=a.defaultPlotOptions,v=a.defined,w=a.erase,n=a.extend,
g=a.fireEvent,c=a.isArray,m=a.isNumber,p=a.isString,b=a.merge,l=a.objectEach,f=a.pick,x=a.removeEvent,t=a.splat,H=a.SVGElement,F=a.syncTimeout,z=a.win;a.Series=a.seriesType("line",null,{lineWidth:2,allowPointSelect:!1,showCheckbox:!1,animation:{duration:1E3},events:{},marker:{lineWidth:0,lineColor:"#ffffff",enabledThreshold:2,radius:4,states:{normal:{animation:!0},hover:{animation:{duration:50},enabled:!0,radiusPlus:2,lineWidthPlus:1},select:{fillColor:"#cccccc",lineColor:"#000000",lineWidth:2}}},
point:{events:{}},dataLabels:{align:"center",formatter:function(){return null===this.y?"":a.numberFormat(this.y,-1)},style:{fontSize:"11px",fontWeight:"bold",color:"contrast",textOutline:"1px contrast"},verticalAlign:"bottom",x:0,y:0,padding:5},cropThreshold:300,pointRange:0,softThreshold:!0,states:{normal:{animation:!0},hover:{animation:{duration:50},lineWidthPlus:1,marker:{},halo:{size:10,opacity:.25}},select:{}},stickyTracking:!0,turboThreshold:1E3,findNearestPointBy:"x"},{isCartesian:!0,pointClass:a.Point,
sorted:!0,requireSorting:!0,directTouch:!1,axisTypes:["xAxis","yAxis"],colorCounter:0,parallelArrays:["x","y"],coll:"series",init:function(a,b){g(this,"init",{options:b});var c=this,k,e=a.series,d;c.chart=a;c.options=b=c.setOptions(b);c.linkedSeries=[];c.bindAxes();n(c,{name:b.name,state:"",visible:!1!==b.visible,selected:!0===b.selected});k=b.events;l(k,function(a,b){y(c,b,a)});if(k&&k.click||b.point&&b.point.events&&b.point.events.click||b.allowPointSelect)a.runTrackerClick=!0;c.getColor();c.getSymbol();
c.parallelArrays.forEach(function(a){c[a+"Data"]=[]});c.setData(b.data,!1);c.isCartesian&&(a.hasCartesianSeries=!0);e.length&&(d=e[e.length-1]);c._i=f(d&&d._i,-1)+1;a.orderSeries(this.insert(e));g(this,"afterInit")},insert:function(a){var b=this.options.index,c;if(m(b)){for(c=a.length;c--;)if(b>=f(a[c].options.index,a[c]._i)){a.splice(c+1,0,this);break}-1===c&&a.unshift(this);c+=1}else a.push(this);return f(c,a.length-1)},bindAxes:function(){var b=this,c=b.options,f=b.chart,d;(b.axisTypes||[]).forEach(function(e){f[e].forEach(function(a){d=
a.options;if(c[e]===d.index||void 0!==c[e]&&c[e]===d.id||void 0===c[e]&&0===d.index)b.insert(a.series),b[e]=a,a.isDirty=!0});b[e]||b.optionalAxis===e||a.error(18,!0,f)})},updateParallelArrays:function(a,b){var c=a.series,f=arguments,e=m(b)?function(e){var f="y"===e&&c.toYData?c.toYData(a):a[e];c[e+"Data"][b]=f}:function(a){Array.prototype[b].apply(c[a+"Data"],Array.prototype.slice.call(f,2))};c.parallelArrays.forEach(e)},autoIncrement:function(){var a=this.options,b=this.xIncrement,c,d=a.pointIntervalUnit,
e=this.chart.time,b=f(b,a.pointStart,0);this.pointInterval=c=f(this.pointInterval,a.pointInterval,1);d&&(a=new e.Date(b),"day"===d?e.set("Date",a,e.get("Date",a)+c):"month"===d?e.set("Month",a,e.get("Month",a)+c):"year"===d&&e.set("FullYear",a,e.get("FullYear",a)+c),c=a.getTime()-b);this.xIncrement=b+c;return b},setOptions:function(a){var c=this.chart,d=c.options,k=d.plotOptions,e=(c.userOptions||{}).plotOptions||{},q=k[this.type],l=c.styledMode;this.userOptions=a;c=b(q,k.series,a);this.tooltipOptions=
b(r.tooltip,r.plotOptions.series&&r.plotOptions.series.tooltip,r.plotOptions[this.type].tooltip,d.tooltip.userOptions,k.series&&k.series.tooltip,k[this.type].tooltip,a.tooltip);this.stickyTracking=f(a.stickyTracking,e[this.type]&&e[this.type].stickyTracking,e.series&&e.series.stickyTracking,this.tooltipOptions.shared&&!this.noSharedTooltip?!0:c.stickyTracking);null===q.marker&&delete c.marker;this.zoneAxis=c.zoneAxis;a=this.zones=(c.zones||[]).slice();!c.negativeColor&&!c.negativeFillColor||c.zones||
(d={value:c[this.zoneAxis+"Threshold"]||c.threshold||0,className:"highcharts-negative"},l||(d.color=c.negativeColor,d.fillColor=c.negativeFillColor),a.push(d));a.length&&v(a[a.length-1].value)&&a.push(l?{}:{color:this.color,fillColor:this.fillColor});g(this,"afterSetOptions",{options:c});return c},getName:function(){return this.name||"Series "+(this.index+1)},getCyclic:function(a,b,c){var d,e=this.chart,k=this.userOptions,g=a+"Index",l=a+"Counter",m=c?c.length:f(e.options.chart[a+"Count"],e[a+"Count"]);
b||(d=f(k[g],k["_"+g]),v(d)||(e.series.length||(e[l]=0),k["_"+g]=d=e[l]%m,e[l]+=1),c&&(b=c[d]));void 0!==d&&(this[g]=d);this[a]=b},getColor:function(){this.chart.styledMode?this.getCyclic("color"):this.options.colorByPoint?this.options.color=null:this.getCyclic("color",this.options.color||u[this.type].color,this.chart.options.colors)},getSymbol:function(){this.getCyclic("symbol",this.options.marker.symbol,this.chart.options.symbols)},drawLegendSymbol:a.LegendSymbolMixin.drawLineMarker,updateData:function(b){var c=
this.options,f=this.points,d=[],e,k,g,l=this.requireSorting;this.xIncrement=null;b.forEach(function(b){var k,q,h;k=a.defined(b)&&this.pointClass.prototype.optionsToObject.call({series:this},b)||{};h=k.x;if((k=k.id)||m(h))k&&(q=(q=this.chart.get(k))&&q.x),void 0===q&&m(h)&&(q=this.xData.indexOf(h,g)),-1===q||void 0===q||f[q].touched?d.push(b):b!==c.data[q]?(f[q].update(b,!1,null,!1),f[q].touched=!0,l&&(g=q+1)):f[q]&&(f[q].touched=!0),e=!0},this);if(e)for(b=f.length;b--;)k=f[b],k.touched||k.remove(!1),
k.touched=!1;else if(b.length===f.length)b.forEach(function(a,b){f[b].update&&a!==c.data[b]&&f[b].update(a,!1,null,!1)});else return!1;d.forEach(function(a){this.addPoint(a,!1)},this);return!0},setData:function(b,d,g,l){var e=this,k=e.points,h=k&&k.length||0,A,t=e.options,n=e.chart,D=null,C=e.xAxis,x=t.turboThreshold,r=this.xData,z=this.yData,u=(A=e.pointArrayMap)&&A.length,v;b=b||[];A=b.length;d=f(d,!0);!1!==l&&A&&h&&!e.cropped&&!e.hasGroupedData&&e.visible&&!e.isSeriesBoosting&&(v=this.updateData(b));
if(!v){e.xIncrement=null;e.colorCounter=0;this.parallelArrays.forEach(function(a){e[a+"Data"].length=0});if(x&&A>x){for(g=0;null===D&&g<A;)D=b[g],g++;if(m(D))for(g=0;g<A;g++)r[g]=this.autoIncrement(),z[g]=b[g];else if(c(D))if(u)for(g=0;g<A;g++)D=b[g],r[g]=D[0],z[g]=D.slice(1,u+1);else for(g=0;g<A;g++)D=b[g],r[g]=D[0],z[g]=D[1];else a.error(12,!1,n)}else for(g=0;g<A;g++)void 0!==b[g]&&(D={series:e},e.pointClass.prototype.applyOptions.apply(D,[b[g]]),e.updateParallelArrays(D,g));z&&p(z[0])&&a.error(14,
!0,n);e.data=[];e.options.data=e.userOptions.data=b;for(g=h;g--;)k[g]&&k[g].destroy&&k[g].destroy();C&&(C.minRange=C.userMinRange);e.isDirty=n.isDirtyBox=!0;e.isDirtyData=!!k;g=!1}"point"===t.legendType&&(this.processData(),this.generatePoints());d&&n.redraw(g)},processData:function(b){var c=this.xData,f=this.yData,d=c.length,e;e=0;var k,g,l=this.xAxis,m,h=this.options;m=h.cropThreshold;var p=this.getExtremesFromAll||h.getExtremesFromAll,t=this.isCartesian,h=l&&l.val2lin,n=l&&l.isLog,x=this.requireSorting,
r,z;if(t&&!this.isDirty&&!l.isDirty&&!this.yAxis.isDirty&&!b)return!1;l&&(b=l.getExtremes(),r=b.min,z=b.max);t&&this.sorted&&!p&&(!m||d>m||this.forceCrop)&&(c[d-1]<r||c[0]>z?(c=[],f=[]):this.yData&&(c[0]<r||c[d-1]>z)&&(e=this.cropData(this.xData,this.yData,r,z),c=e.xData,f=e.yData,e=e.start,k=!0));for(m=c.length||1;--m;)d=n?h(c[m])-h(c[m-1]):c[m]-c[m-1],0<d&&(void 0===g||d<g)?g=d:0>d&&x&&(a.error(15,!1,this.chart),x=!1);this.cropped=k;this.cropStart=e;this.processedXData=c;this.processedYData=f;this.closestPointRange=
g},cropData:function(a,b,c,d,e){var k=a.length,g=0,l=k,m;e=f(e,this.cropShoulder,1);for(m=0;m<k;m++)if(a[m]>=c){g=Math.max(0,m-e);break}for(c=m;c<k;c++)if(a[c]>d){l=c+e;break}return{xData:a.slice(g,l),yData:b.slice(g,l),start:g,end:l}},generatePoints:function(){var a=this.options,b=a.data,c=this.data,f,e=this.processedXData,d=this.processedYData,g=this.pointClass,l=e.length,m=this.cropStart||0,h,p=this.hasGroupedData,a=a.keys,x,r=[],z;c||p||(c=[],c.length=b.length,c=this.data=c);a&&p&&(this.options.keys=
!1);for(z=0;z<l;z++)h=m+z,p?(x=(new g).init(this,[e[z]].concat(t(d[z]))),x.dataGroup=this.groupMap[z],x.dataGroup.options&&(x.options=x.dataGroup.options,n(x,x.dataGroup.options))):(x=c[h])||void 0===b[h]||(c[h]=x=(new g).init(this,b[h],e[z])),x&&(x.index=h,r[z]=x);this.options.keys=a;if(c&&(l!==(f=c.length)||p))for(z=0;z<f;z++)z!==m||p||(z+=l),c[z]&&(c[z].destroyElements(),c[z].plotX=void 0);this.data=c;this.points=r},getExtremes:function(a){var b=this.yAxis,f=this.processedXData,d,e=[],k=0;d=this.xAxis.getExtremes();
var g=d.min,l=d.max,p,t,n=this.requireSorting?1:0,x,z;a=a||this.stackedYData||this.processedYData||[];d=a.length;for(z=0;z<d;z++)if(t=f[z],x=a[z],p=(m(x,!0)||c(x))&&(!b.positiveValuesOnly||x.length||0<x),t=this.getExtremesFromAll||this.options.getExtremesFromAll||this.cropped||(f[z+n]||t)>=g&&(f[z-n]||t)<=l,p&&t)if(p=x.length)for(;p--;)"number"===typeof x[p]&&(e[k++]=x[p]);else e[k++]=x;this.dataMin=h(e);this.dataMax=E(e)},translate:function(){this.processedXData||this.processData();this.generatePoints();
var a=this.options,b=a.stacking,c=this.xAxis,l=c.categories,e=this.yAxis,q=this.points,h=q.length,p=!!this.modifyValue,t=a.pointPlacement,n="between"===t||m(t),x=a.threshold,z=a.startFromThreshold?x:0,r,u,F,H,w=Number.MAX_VALUE;"between"===t&&(t=.5);m(t)&&(t*=f(a.pointRange||c.pointRange));for(a=0;a<h;a++){var y=q[a],E=y.x,G=y.y;u=y.low;var J=b&&e.stacks[(this.negStacks&&G<(z?0:x)?"-":"")+this.stackKey],U;e.positiveValuesOnly&&null!==G&&0>=G&&(y.isNull=!0);y.plotX=r=d(Math.min(Math.max(-1E5,c.translate(E,
0,0,0,1,t,"flags"===this.type)),1E5));b&&this.visible&&!y.isNull&&J&&J[E]&&(H=this.getStackIndicator(H,E,this.index),U=J[E],G=U.points[H.key],u=G[0],G=G[1],u===z&&H.key===J[E].base&&(u=f(m(x)&&x,e.min)),e.positiveValuesOnly&&0>=u&&(u=null),y.total=y.stackTotal=U.total,y.percentage=U.total&&y.y/U.total*100,y.stackY=G,U.setOffset(this.pointXOffset||0,this.barW||0));y.yBottom=v(u)?Math.min(Math.max(-1E5,e.translate(u,0,1,0,1)),1E5):null;p&&(G=this.modifyValue(G,y));y.plotY=u="number"===typeof G&&Infinity!==
G?Math.min(Math.max(-1E5,e.translate(G,0,1,0,1)),1E5):void 0;y.isInside=void 0!==u&&0<=u&&u<=e.len&&0<=r&&r<=c.len;y.clientX=n?d(c.translate(E,0,0,0,1,t)):r;y.negative=y.y<(x||0);y.category=l&&void 0!==l[y.x]?l[y.x]:y.x;y.isNull||(void 0!==F&&(w=Math.min(w,Math.abs(r-F))),F=r);y.zone=this.zones.length&&y.getZone()}this.closestPointRangePx=w;g(this,"afterTranslate")},getValidPoints:function(a,b){var c=this.chart;return(a||this.points||[]).filter(function(a){return b&&!c.isInsidePlot(a.plotX,a.plotY,
c.inverted)?!1:!a.isNull})},setClip:function(a){var b=this.chart,c=this.options,f=b.renderer,e=b.inverted,d=this.clipBox,k=d||b.clipBox,g=this.sharedClipKey||["_sharedClip",a&&a.duration,a&&a.easing,k.height,c.xAxis,c.yAxis].join(),l=b[g],m=b[g+"m"];l||(a&&(k.width=0,e&&(k.x=b.plotSizeX),b[g+"m"]=m=f.clipRect(e?b.plotSizeX+99:-99,e?-b.plotLeft:-b.plotTop,99,e?b.chartWidth:b.chartHeight)),b[g]=l=f.clipRect(k),l.count={length:0});a&&!l.count[this.index]&&(l.count[this.index]=!0,l.count.length+=1);!1!==
c.clip&&(this.group.clip(a||d?l:b.clipRect),this.markerGroup.clip(m),this.sharedClipKey=g);a||(l.count[this.index]&&(delete l.count[this.index],--l.count.length),0===l.count.length&&g&&b[g]&&(d||(b[g]=b[g].destroy()),b[g+"m"]&&(b[g+"m"]=b[g+"m"].destroy())))},animate:function(a){var b=this.chart,c=G(this.options.animation),f;a?this.setClip(c):(f=this.sharedClipKey,(a=b[f])&&a.animate({width:b.plotSizeX,x:0},c),b[f+"m"]&&b[f+"m"].animate({width:b.plotSizeX+99,x:0},c),this.animate=null)},afterAnimate:function(){this.setClip();
g(this,"afterAnimate");this.finishedAnimating=!0},drawPoints:function(){var a=this.points,b=this.chart,c,d,e,g,l=this.options.marker,m,h,p,t=this[this.specialGroup]||this.markerGroup;c=this.xAxis;var n,x=f(l.enabled,!c||c.isRadial?!0:null,this.closestPointRangePx>=l.enabledThreshold*l.radius);if(!1!==l.enabled||this._hasPointMarkers)for(c=0;c<a.length;c++)d=a[c],g=d.graphic,m=d.marker||{},h=!!d.marker,e=x&&void 0===m.enabled||m.enabled,p=!1!==d.isInside,e&&!d.isNull?(e=f(m.symbol,this.symbol),n=this.markerAttribs(d,
d.selected&&"select"),g?g[p?"show":"hide"](!0).animate(n):p&&(0<n.width||d.hasImage)&&(d.graphic=g=b.renderer.symbol(e,n.x,n.y,n.width,n.height,h?m:l).add(t)),g&&!b.styledMode&&g.attr(this.pointAttribs(d,d.selected&&"select")),g&&g.addClass(d.getClassName(),!0)):g&&(d.graphic=g.destroy())},markerAttribs:function(a,b){var c=this.options.marker,d=a.marker||{},e=d.symbol||c.symbol,k=f(d.radius,c.radius);b&&(c=c.states[b],b=d.states&&d.states[b],k=f(b&&b.radius,c&&c.radius,k+(c&&c.radiusPlus||0)));a.hasImage=
e&&0===e.indexOf("url");a.hasImage&&(k=0);a={x:Math.floor(a.plotX)-k,y:a.plotY-k};k&&(a.width=a.height=2*k);return a},pointAttribs:function(a,b){var c=this.options.marker,d=a&&a.options,e=d&&d.marker||{},g=this.color,k=d&&d.color,l=a&&a.color,d=f(e.lineWidth,c.lineWidth);a=a&&a.zone&&a.zone.color;g=k||a||l||g;a=e.fillColor||c.fillColor||g;g=e.lineColor||c.lineColor||g;b&&(c=c.states[b],b=e.states&&e.states[b]||{},d=f(b.lineWidth,c.lineWidth,d+f(b.lineWidthPlus,c.lineWidthPlus,0)),a=b.fillColor||c.fillColor||
a,g=b.lineColor||c.lineColor||g);return{stroke:g,"stroke-width":d,fill:a}},destroy:function(){var b=this,c=b.chart,f=/AppleWebKit\/533/.test(z.navigator.userAgent),d,e,q=b.data||[],m,h;g(b,"destroy");x(b);(b.axisTypes||[]).forEach(function(a){(h=b[a])&&h.series&&(w(h.series,b),h.isDirty=h.forceRedraw=!0)});b.legendItem&&b.chart.legend.destroyItem(b);for(e=q.length;e--;)(m=q[e])&&m.destroy&&m.destroy();b.points=null;a.clearTimeout(b.animationTimeout);l(b,function(a,b){a instanceof H&&!a.survive&&(d=
f&&"group"===b?"hide":"destroy",a[d]())});c.hoverSeries===b&&(c.hoverSeries=null);w(c.series,b);c.orderSeries();l(b,function(a,e){delete b[e]})},getGraphPath:function(a,b,c){var f=this,e=f.options,d=e.step,g,k=[],l=[],m;a=a||f.points;(g=a.reversed)&&a.reverse();(d={right:1,center:2}[d]||d&&3)&&g&&(d=4-d);!e.connectNulls||b||c||(a=this.getValidPoints(a));a.forEach(function(g,q){var h=g.plotX,p=g.plotY,t=a[q-1];(g.leftCliff||t&&t.rightCliff)&&!c&&(m=!0);g.isNull&&!v(b)&&0<q?m=!e.connectNulls:g.isNull&&
!b?m=!0:(0===q||m?q=["M",g.plotX,g.plotY]:f.getPointSpline?q=f.getPointSpline(a,g,q):d?(q=1===d?["L",t.plotX,p]:2===d?["L",(t.plotX+h)/2,t.plotY,"L",(t.plotX+h)/2,p]:["L",h,t.plotY],q.push("L",h,p)):q=["L",h,p],l.push(g.x),d&&(l.push(g.x),2===d&&l.push(g.x)),k.push.apply(k,q),m=!1)});k.xMap=l;return f.graphPath=k},drawGraph:function(){var a=this,b=this.options,c=(this.gappedPath||this.getGraphPath).call(this),f=this.chart.styledMode,e=[["graph","highcharts-graph"]];f||e[0].push(b.lineColor||this.color,
b.dashStyle);e=a.getZonesGraphs(e);e.forEach(function(e,d){var g=e[0],k=a[g];k?(k.endX=a.preventGraphAnimation?null:c.xMap,k.animate({d:c})):c.length&&(a[g]=a.chart.renderer.path(c).addClass(e[1]).attr({zIndex:1}).add(a.group),f||(k={stroke:e[2],"stroke-width":b.lineWidth,fill:a.fillGraph&&a.color||"none"},e[3]?k.dashstyle=e[3]:"square"!==b.linecap&&(k["stroke-linecap"]=k["stroke-linejoin"]="round"),k=a[g].attr(k).shadow(2>d&&b.shadow)));k&&(k.startX=c.xMap,k.isArea=c.isArea)})},getZonesGraphs:function(a){this.zones.forEach(function(b,
c){c=["zone-graph-"+c,"highcharts-graph highcharts-zone-graph-"+c+" "+(b.className||"")];this.chart.styledMode||c.push(b.color||this.color,b.dashStyle||this.options.dashStyle);a.push(c)},this);return a},applyZones:function(){var a=this,b=this.chart,c=b.renderer,d=this.zones,e,g,l=this.clips||[],m,h=this.graph,p=this.area,t=Math.max(b.chartWidth,b.chartHeight),n=this[(this.zoneAxis||"y")+"Axis"],x,z,r=b.inverted,u,v,F,H,w=!1;d.length&&(h||p)&&n&&void 0!==n.min&&(z=n.reversed,u=n.horiz,h&&!this.showLine&&
h.hide(),p&&p.hide(),x=n.getExtremes(),d.forEach(function(d,k){e=z?u?b.plotWidth:0:u?0:n.toPixels(x.min)||0;e=Math.min(Math.max(f(g,e),0),t);g=Math.min(Math.max(Math.round(n.toPixels(f(d.value,x.max),!0)||0),0),t);w&&(e=g=n.toPixels(x.max));v=Math.abs(e-g);F=Math.min(e,g);H=Math.max(e,g);n.isXAxis?(m={x:r?H:F,y:0,width:v,height:t},u||(m.x=b.plotHeight-m.x)):(m={x:0,y:r?H:F,width:t,height:v},u&&(m.y=b.plotWidth-m.y));r&&c.isVML&&(m=n.isXAxis?{x:0,y:z?F:H,height:m.width,width:b.chartWidth}:{x:m.y-b.plotLeft-
b.spacingBox.x,y:0,width:m.height,height:b.chartHeight});l[k]?l[k].animate(m):(l[k]=c.clipRect(m),h&&a["zone-graph-"+k].clip(l[k]),p&&a["zone-area-"+k].clip(l[k]));w=d.value>x.max;a.resetZones&&0===g&&(g=void 0)}),this.clips=l)},invertGroups:function(a){function b(){["group","markerGroup"].forEach(function(b){c[b]&&(f.renderer.isVML&&c[b].attr({width:c.yAxis.len,height:c.xAxis.len}),c[b].width=c.yAxis.len,c[b].height=c.xAxis.len,c[b].invert(a))})}var c=this,f=c.chart,e;c.xAxis&&(e=y(f,"resize",b),
y(c,"destroy",e),b(a),c.invertGroups=b)},plotGroup:function(a,b,c,f,e){var d=this[a],g=!d;g&&(this[a]=d=this.chart.renderer.g().attr({zIndex:f||.1}).add(e));d.addClass("highcharts-"+b+" highcharts-series-"+this.index+" highcharts-"+this.type+"-series "+(v(this.colorIndex)?"highcharts-color-"+this.colorIndex+" ":"")+(this.options.className||"")+(d.hasClass("highcharts-tracker")?" highcharts-tracker":""),!0);d.attr({visibility:c})[g?"attr":"animate"](this.getPlotBox());return d},getPlotBox:function(){var a=
this.chart,b=this.xAxis,c=this.yAxis;a.inverted&&(b=c,c=this.xAxis);return{translateX:b?b.left:a.plotLeft,translateY:c?c.top:a.plotTop,scaleX:1,scaleY:1}},render:function(){var a=this,b=a.chart,c,f=a.options,e=!!a.animate&&b.renderer.isSVG&&G(f.animation).duration,d=a.visible?"inherit":"hidden",l=f.zIndex,m=a.hasRendered,h=b.seriesGroup,p=b.inverted;c=a.plotGroup("group","series",d,l,h);a.markerGroup=a.plotGroup("markerGroup","markers",d,l,h);e&&a.animate(!0);c.inverted=a.isCartesian?p:!1;a.drawGraph&&
(a.drawGraph(),a.applyZones());a.drawDataLabels&&a.drawDataLabels();a.visible&&a.drawPoints();a.drawTracker&&!1!==a.options.enableMouseTracking&&a.drawTracker();a.invertGroups(p);!1===f.clip||a.sharedClipKey||m||c.clip(b.clipRect);e&&a.animate();m||(a.animationTimeout=F(function(){a.afterAnimate()},e));a.isDirty=!1;a.hasRendered=!0;g(a,"afterRender")},redraw:function(){var a=this.chart,b=this.isDirty||this.isDirtyData,c=this.group,d=this.xAxis,e=this.yAxis;c&&(a.inverted&&c.attr({width:a.plotWidth,
height:a.plotHeight}),c.animate({translateX:f(d&&d.left,a.plotLeft),translateY:f(e&&e.top,a.plotTop)}));this.translate();this.render();b&&delete this.kdTree},kdAxisArray:["clientX","plotY"],searchPoint:function(a,b){var c=this.xAxis,f=this.yAxis,e=this.chart.inverted;return this.searchKDTree({clientX:e?c.len-a.chartY+c.pos:a.chartX-c.pos,plotY:e?f.len-a.chartX+f.pos:a.chartY-f.pos},b)},buildKDTree:function(){function a(c,e,f){var d,g;if(g=c&&c.length)return d=b.kdAxisArray[e%f],c.sort(function(a,
b){return a[d]-b[d]}),g=Math.floor(g/2),{point:c[g],left:a(c.slice(0,g),e+1,f),right:a(c.slice(g+1),e+1,f)}}this.buildingKdTree=!0;var b=this,c=-1<b.options.findNearestPointBy.indexOf("y")?2:1;delete b.kdTree;F(function(){b.kdTree=a(b.getValidPoints(null,!b.directTouch),c,c);b.buildingKdTree=!1},b.options.kdNow?0:1)},searchKDTree:function(a,b){function c(a,b,k,l){var m=b.point,h=f.kdAxisArray[k%l],q,p,t=m;p=v(a[e])&&v(m[e])?Math.pow(a[e]-m[e],2):null;q=v(a[d])&&v(m[d])?Math.pow(a[d]-m[d],2):null;
q=(p||0)+(q||0);m.dist=v(q)?Math.sqrt(q):Number.MAX_VALUE;m.distX=v(p)?Math.sqrt(p):Number.MAX_VALUE;h=a[h]-m[h];q=0>h?"left":"right";p=0>h?"right":"left";b[q]&&(q=c(a,b[q],k+1,l),t=q[g]<t[g]?q:m);b[p]&&Math.sqrt(h*h)<t[g]&&(a=c(a,b[p],k+1,l),t=a[g]<t[g]?a:t);return t}var f=this,e=this.kdAxisArray[0],d=this.kdAxisArray[1],g=b?"distX":"dist";b=-1<f.options.findNearestPointBy.indexOf("y")?2:1;this.kdTree||this.buildingKdTree||this.buildKDTree();if(this.kdTree)return c(a,this.kdTree,b,b)}})})(J);(function(a){var y=
a.Axis,G=a.Chart,E=a.correctFloat,h=a.defined,d=a.destroyObjectProperties,r=a.format,u=a.objectEach,v=a.pick,w=a.Series;a.StackItem=function(a,d,c,m,h){var b=a.chart.inverted;this.axis=a;this.isNegative=c;this.options=d;this.x=m;this.total=null;this.points={};this.stack=h;this.rightCliff=this.leftCliff=0;this.alignOptions={align:d.align||(b?c?"left":"right":"center"),verticalAlign:d.verticalAlign||(b?"middle":c?"bottom":"top"),y:v(d.y,b?4:c?14:-6),x:v(d.x,b?c?-6:6:0)};this.textAlign=d.textAlign||
(b?c?"right":"left":"center")};a.StackItem.prototype={destroy:function(){d(this,this.axis)},render:function(a){var d=this.axis.chart,c=this.options,m=c.format,m=m?r(m,this,d.time):c.formatter.call(this);this.label?this.label.attr({text:m,visibility:"hidden"}):this.label=d.renderer.text(m,null,null,c.useHTML).css(c.style).attr({align:this.textAlign,rotation:c.rotation,visibility:"hidden"}).add(a);this.label.labelrank=d.plotHeight},setOffset:function(a,d){var c=this.axis,g=c.chart,p=c.translate(c.usePercentage?
100:this.total,0,0,0,1),b=c.translate(0),b=h(p)&&Math.abs(p-b);a=g.xAxis[0].translate(this.x)+a;c=h(p)&&this.getStackBox(g,this,a,p,d,b,c);(d=this.label)&&c&&(d.align(this.alignOptions,null,c),c=d.alignAttr,d[!1===this.options.crop||g.isInsidePlot(c.x,c.y)?"show":"hide"](!0))},getStackBox:function(a,d,c,m,h,b,l){var f=d.axis.reversed,g=a.inverted;a=l.height+l.pos-(g?a.plotLeft:a.plotTop);d=d.isNegative&&!f||!d.isNegative&&f;return{x:g?d?m:m-b:c,y:g?a-c-h:d?a-m-b:a-m,width:g?b:h,height:g?h:b}}};G.prototype.getStacks=
function(){var a=this;a.yAxis.forEach(function(a){a.stacks&&a.hasVisibleSeries&&(a.oldStacks=a.stacks)});a.series.forEach(function(d){!d.options.stacking||!0!==d.visible&&!1!==a.options.chart.ignoreHiddenSeries||(d.stackKey=d.type+v(d.options.stack,""))})};y.prototype.buildStacks=function(){var a=this.series,d=v(this.options.reversedStacks,!0),c=a.length,m;if(!this.isXAxis){this.usePercentage=!1;for(m=c;m--;)a[d?m:c-m-1].setStackedPoints();for(m=0;m<c;m++)a[m].modifyStacks()}};y.prototype.renderStackTotals=
function(){var a=this.chart,d=a.renderer,c=this.stacks,m=this.stackTotalGroup;m||(this.stackTotalGroup=m=d.g("stack-labels").attr({visibility:"visible",zIndex:6}).add());m.translate(a.plotLeft,a.plotTop);u(c,function(a){u(a,function(a){a.render(m)})})};y.prototype.resetStacks=function(){var a=this,d=a.stacks;a.isXAxis||u(d,function(c){u(c,function(d,g){d.touched<a.stacksTouched?(d.destroy(),delete c[g]):(d.total=null,d.cumulative=null)})})};y.prototype.cleanStacks=function(){var a;this.isXAxis||(this.oldStacks&&
(a=this.stacks=this.oldStacks),u(a,function(a){u(a,function(a){a.cumulative=a.total})}))};w.prototype.setStackedPoints=function(){if(this.options.stacking&&(!0===this.visible||!1===this.chart.options.chart.ignoreHiddenSeries)){var d=this.processedXData,g=this.processedYData,c=[],m=g.length,p=this.options,b=p.threshold,l=v(p.startFromThreshold&&b,0),f=p.stack,p=p.stacking,x=this.stackKey,t="-"+x,r=this.negStacks,u=this.yAxis,z=u.stacks,k=u.oldStacks,A,D,C,e,q,w,y;u.stacksTouched+=1;for(q=0;q<m;q++)w=
d[q],y=g[q],A=this.getStackIndicator(A,w,this.index),e=A.key,C=(D=r&&y<(l?0:b))?t:x,z[C]||(z[C]={}),z[C][w]||(k[C]&&k[C][w]?(z[C][w]=k[C][w],z[C][w].total=null):z[C][w]=new a.StackItem(u,u.options.stackLabels,D,w,f)),C=z[C][w],null!==y?(C.points[e]=C.points[this.index]=[v(C.cumulative,l)],h(C.cumulative)||(C.base=e),C.touched=u.stacksTouched,0<A.index&&!1===this.singleStacks&&(C.points[e][0]=C.points[this.index+","+w+",0"][0])):C.points[e]=C.points[this.index]=null,"percent"===p?(D=D?x:t,r&&z[D]&&
z[D][w]?(D=z[D][w],C.total=D.total=Math.max(D.total,C.total)+Math.abs(y)||0):C.total=E(C.total+(Math.abs(y)||0))):C.total=E(C.total+(y||0)),C.cumulative=v(C.cumulative,l)+(y||0),null!==y&&(C.points[e].push(C.cumulative),c[q]=C.cumulative);"percent"===p&&(u.usePercentage=!0);this.stackedYData=c;u.oldStacks={}}};w.prototype.modifyStacks=function(){var a=this,d=a.stackKey,c=a.yAxis.stacks,m=a.processedXData,h,b=a.options.stacking;a[b+"Stacker"]&&[d,"-"+d].forEach(function(d){for(var f=m.length,g,l;f--;)if(g=
m[f],h=a.getStackIndicator(h,g,a.index,d),l=(g=c[d]&&c[d][g])&&g.points[h.key])a[b+"Stacker"](l,g,f)})};w.prototype.percentStacker=function(a,d,c){d=d.total?100/d.total:0;a[0]=E(a[0]*d);a[1]=E(a[1]*d);this.stackedYData[c]=a[1]};w.prototype.getStackIndicator=function(a,d,c,m){!h(a)||a.x!==d||m&&a.key!==m?a={x:d,index:0,key:m}:a.index++;a.key=[c,d,a.index].join();return a}})(J);(function(a){var y=a.addEvent,G=a.animate,E=a.Axis,h=a.Chart,d=a.createElement,r=a.css,u=a.defined,v=a.erase,w=a.extend,n=
a.fireEvent,g=a.isNumber,c=a.isObject,m=a.isArray,p=a.merge,b=a.objectEach,l=a.pick,f=a.Point,x=a.Series,t=a.seriesTypes,H=a.setAnimation,F=a.splat;a.cleanRecursively=function(d,f){var g=0,k=0;b(d,function(b,e){c(d[e],!0)&&f[e]?a.cleanRecursively(d[e],f[e])&&delete d[e]:c(d[e])||d[e]!==f[e]||(delete d[e],k++);g++});return g===k};w(h.prototype,{addSeries:function(a,b,c){var d,f=this;a&&(b=l(b,!0),n(f,"addSeries",{options:a},function(){d=f.initSeries(a);f.isDirtyLegend=!0;f.linkSeries();n(f,"afterAddSeries");
b&&f.redraw(c)}));return d},addAxis:function(a,b,c,d){var f=b?"xAxis":"yAxis",e=this.options;a=p(a,{index:this[f].length,isX:b});b=new E(this,a);e[f]=F(e[f]||{});e[f].push(a);l(c,!0)&&this.redraw(d);return b},showLoading:function(a){var b=this,c=b.options,f=b.loadingDiv,g=c.loading,e=function(){f&&r(f,{left:b.plotLeft+"px",top:b.plotTop+"px",width:b.plotWidth+"px",height:b.plotHeight+"px"})};f||(b.loadingDiv=f=d("div",{className:"highcharts-loading highcharts-loading-hidden"},null,b.container),b.loadingSpan=
d("span",{className:"highcharts-loading-inner"},null,f),y(b,"redraw",e));f.className="highcharts-loading";b.loadingSpan.innerHTML=a||c.lang.loading;b.styledMode||(r(f,w(g.style,{zIndex:10})),r(b.loadingSpan,g.labelStyle),b.loadingShown||(r(f,{opacity:0,display:""}),G(f,{opacity:g.style.opacity||.5},{duration:g.showDuration||0})));b.loadingShown=!0;e()},hideLoading:function(){var a=this.options,b=this.loadingDiv;b&&(b.className="highcharts-loading highcharts-loading-hidden",this.styledMode||G(b,{opacity:0},
{duration:a.loading.hideDuration||100,complete:function(){r(b,{display:"none"})}}));this.loadingShown=!1},propsRequireDirtyBox:"backgroundColor borderColor borderWidth margin marginTop marginRight marginBottom marginLeft spacing spacingTop spacingRight spacingBottom spacingLeft borderRadius plotBackgroundColor plotBackgroundImage plotBorderColor plotBorderWidth plotShadow shadow".split(" "),propsRequireUpdateSeries:"chart.inverted chart.polar chart.ignoreHiddenSeries chart.type colors plotOptions time tooltip".split(" "),
collectionsWithUpdate:"xAxis yAxis zAxis series colorAxis pane".split(" "),update:function(c,d,f,m){var k=this,e={credits:"addCredits",title:"setTitle",subtitle:"setSubtitle"},h,t,x,r=[];n(k,"update",{options:c});a.cleanRecursively(c,k.options);if(h=c.chart){p(!0,k.options.chart,h);"className"in h&&k.setClassName(h.className);"reflow"in h&&k.setReflow(h.reflow);if("inverted"in h||"polar"in h||"type"in h)k.propFromSeries(),t=!0;"alignTicks"in h&&(t=!0);b(h,function(a,b){-1!==k.propsRequireUpdateSeries.indexOf("chart."+
b)&&(x=!0);-1!==k.propsRequireDirtyBox.indexOf(b)&&(k.isDirtyBox=!0)});!k.styledMode&&"style"in h&&k.renderer.setStyle(h.style)}!k.styledMode&&c.colors&&(this.options.colors=c.colors);c.plotOptions&&p(!0,this.options.plotOptions,c.plotOptions);b(c,function(a,b){if(k[b]&&"function"===typeof k[b].update)k[b].update(a,!1);else if("function"===typeof k[e[b]])k[e[b]](a);"chart"!==b&&-1!==k.propsRequireUpdateSeries.indexOf(b)&&(x=!0)});this.collectionsWithUpdate.forEach(function(a){var b;c[a]&&("series"===
a&&(b=[],k[a].forEach(function(a,c){a.options.isInternal||b.push(l(a.options.index,c))})),F(c[a]).forEach(function(c,e){(e=u(c.id)&&k.get(c.id)||k[a][b?b[e]:e])&&e.coll===a&&(e.update(c,!1),f&&(e.touched=!0));if(!e&&f)if("series"===a)k.addSeries(c,!1).touched=!0;else if("xAxis"===a||"yAxis"===a)k.addAxis(c,"xAxis"===a,!1).touched=!0}),f&&k[a].forEach(function(a){a.touched||a.options.isInternal?delete a.touched:r.push(a)}))});r.forEach(function(a){a.remove&&a.remove(!1)});t&&k.axes.forEach(function(a){a.update({},
!1)});x&&k.series.forEach(function(a){a.update({},!1)});c.loading&&p(!0,k.options.loading,c.loading);t=h&&h.width;h=h&&h.height;g(t)&&t!==k.chartWidth||g(h)&&h!==k.chartHeight?k.setSize(t,h,m):l(d,!0)&&k.redraw(m);n(k,"afterUpdate",{options:c})},setSubtitle:function(a){this.setTitle(void 0,a)}});w(f.prototype,{update:function(a,b,d,f){function g(){e.applyOptions(a);null===e.y&&h&&(e.graphic=h.destroy());c(a,!0)&&(h&&h.element&&a&&a.marker&&void 0!==a.marker.symbol&&(e.graphic=h.destroy()),a&&a.dataLabels&&
e.dataLabel&&(e.dataLabel=e.dataLabel.destroy()),e.connector&&(e.connector=e.connector.destroy()));m=e.index;k.updateParallelArrays(e,m);t.data[m]=c(t.data[m],!0)||c(a,!0)?e.options:l(a,t.data[m]);k.isDirty=k.isDirtyData=!0;!k.fixedBox&&k.hasCartesianSeries&&(p.isDirtyBox=!0);"point"===t.legendType&&(p.isDirtyLegend=!0);b&&p.redraw(d)}var e=this,k=e.series,h=e.graphic,m,p=k.chart,t=k.options;b=l(b,!0);!1===f?g():e.firePointEvent("update",{options:a},g)},remove:function(a,b){this.series.removePoint(this.series.data.indexOf(this),
a,b)}});w(x.prototype,{addPoint:function(a,b,c,d){var f=this.options,e=this.data,g=this.chart,k=this.xAxis,k=k&&k.hasNames&&k.names,h=f.data,m,p,t=this.xData,n,x;b=l(b,!0);m={series:this};this.pointClass.prototype.applyOptions.apply(m,[a]);x=m.x;n=t.length;if(this.requireSorting&&x<t[n-1])for(p=!0;n&&t[n-1]>x;)n--;this.updateParallelArrays(m,"splice",n,0,0);this.updateParallelArrays(m,n);k&&m.name&&(k[x]=m.name);h.splice(n,0,a);p&&(this.data.splice(n,0,null),this.processData());"point"===f.legendType&&
this.generatePoints();c&&(e[0]&&e[0].remove?e[0].remove(!1):(e.shift(),this.updateParallelArrays(m,"shift"),h.shift()));this.isDirtyData=this.isDirty=!0;b&&g.redraw(d)},removePoint:function(a,b,c){var d=this,f=d.data,e=f[a],g=d.points,k=d.chart,h=function(){g&&g.length===f.length&&g.splice(a,1);f.splice(a,1);d.options.data.splice(a,1);d.updateParallelArrays(e||{series:d},"splice",a,1);e&&e.destroy();d.isDirty=!0;d.isDirtyData=!0;b&&k.redraw()};H(c,k);b=l(b,!0);e?e.firePointEvent("remove",null,h):
h()},remove:function(a,b,c){function d(){f.destroy();f.remove=null;e.isDirtyLegend=e.isDirtyBox=!0;e.linkSeries();l(a,!0)&&e.redraw(b)}var f=this,e=f.chart;!1!==c?n(f,"remove",null,d):d()},update:function(b,c){a.cleanRecursively(b,this.userOptions);var d=this,f=d.chart,g=d.userOptions,e=d.oldType||d.type,k=b.type||g.type||f.options.chart.type,h=t[e].prototype,m,x=["group","markerGroup","dataLabelsGroup"],r=["navigatorSeries","baseSeries"],u=d.finishedAnimating&&{animation:!1},z=["data","name","turboThreshold"],
v=Object.keys(b),F=0<v.length;v.forEach(function(a){-1===z.indexOf(a)&&(F=!1)});if(F)b.data&&this.setData(b.data,!1),b.name&&this.setName(b.name,!1);else{r=x.concat(r);r.forEach(function(a){r[a]=d[a];delete d[a]});b=p(g,u,{index:d.index,pointStart:l(g.pointStart,d.xData[0])},{data:d.options.data},b);d.remove(!1,null,!1);for(m in h)d[m]=void 0;t[k||e]?w(d,t[k||e].prototype):a.error(17,!0,f);r.forEach(function(a){d[a]=r[a]});d.init(f,b);b.zIndex!==g.zIndex&&x.forEach(function(a){d[a]&&d[a].attr({zIndex:b.zIndex})});
d.oldType=e;f.linkSeries()}n(this,"afterUpdate");l(c,!0)&&f.redraw(F?void 0:!1)},setName:function(a){this.name=this.options.name=this.userOptions.name=a;this.chart.isDirtyLegend=!0}});w(E.prototype,{update:function(a,c){var d=this.chart,f=a&&a.events||{};a=p(this.userOptions,a);d.options[this.coll].indexOf&&(d.options[this.coll][d.options[this.coll].indexOf(this.userOptions)]=a);b(d.options[this.coll].events,function(a,b){"undefined"===typeof f[b]&&(f[b]=void 0)});this.destroy(!0);this.init(d,w(a,
{events:f}));d.isDirtyBox=!0;l(c,!0)&&d.redraw()},remove:function(a){for(var b=this.chart,c=this.coll,d=this.series,f=d.length;f--;)d[f]&&d[f].remove(!1);v(b.axes,this);v(b[c],this);m(b.options[c])?b.options[c].splice(this.options.index,1):delete b.options[c];b[c].forEach(function(a,b){a.options.index=a.userOptions.index=b});this.destroy();b.isDirtyBox=!0;l(a,!0)&&b.redraw()},setTitle:function(a,b){this.update({title:a},b)},setCategories:function(a,b){this.update({categories:a},b)}})})(J);(function(a){var y=
a.color,G=a.pick,E=a.Series,h=a.seriesType;h("area","line",{softThreshold:!1,threshold:0},{singleStacks:!1,getStackPoints:function(d){var h=[],u=[],v=this.xAxis,w=this.yAxis,n=w.stacks[this.stackKey],g={},c=this.index,m=w.series,p=m.length,b,l=G(w.options.reversedStacks,!0)?1:-1,f;d=d||this.points;if(this.options.stacking){for(f=0;f<d.length;f++)d[f].leftNull=d[f].rightNull=null,g[d[f].x]=d[f];a.objectEach(n,function(a,b){null!==a.total&&u.push(b)});u.sort(function(a,b){return a-b});b=m.map(function(a){return a.visible});
u.forEach(function(a,d){var m=0,t,x;if(g[a]&&!g[a].isNull)h.push(g[a]),[-1,1].forEach(function(k){var h=1===k?"rightNull":"leftNull",m=0,r=n[u[d+k]];if(r)for(f=c;0<=f&&f<p;)t=r.points[f],t||(f===c?g[a][h]=!0:b[f]&&(x=n[a].points[f])&&(m-=x[1]-x[0])),f+=l;g[a][1===k?"rightCliff":"leftCliff"]=m});else{for(f=c;0<=f&&f<p;){if(t=n[a].points[f]){m=t[1];break}f+=l}m=w.translate(m,0,1,0,1);h.push({isNull:!0,plotX:v.translate(a,0,0,0,1),x:a,plotY:m,yBottom:m})}})}return h},getGraphPath:function(a){var d=E.prototype.getGraphPath,
h=this.options,v=h.stacking,w=this.yAxis,n,g,c=[],m=[],p=this.index,b,l=w.stacks[this.stackKey],f=h.threshold,x=w.getThreshold(h.threshold),t,h=h.connectNulls||"percent"===v,H=function(d,g,k){var h=a[d];d=v&&l[h.x].points[p];var t=h[k+"Null"]||0;k=h[k+"Cliff"]||0;var n,e,h=!0;k||t?(n=(t?d[0]:d[1])+k,e=d[0]+k,h=!!t):!v&&a[g]&&a[g].isNull&&(n=e=f);void 0!==n&&(m.push({plotX:b,plotY:null===n?x:w.getThreshold(n),isNull:h,isCliff:!0}),c.push({plotX:b,plotY:null===e?x:w.getThreshold(e),doCurve:!1}))};a=
a||this.points;v&&(a=this.getStackPoints(a));for(n=0;n<a.length;n++)if(g=a[n].isNull,b=G(a[n].rectPlotX,a[n].plotX),t=G(a[n].yBottom,x),!g||h)h||H(n,n-1,"left"),g&&!v&&h||(m.push(a[n]),c.push({x:n,plotX:b,plotY:t})),h||H(n,n+1,"right");n=d.call(this,m,!0,!0);c.reversed=!0;g=d.call(this,c,!0,!0);g.length&&(g[0]="L");g=n.concat(g);d=d.call(this,m,!1,h);g.xMap=n.xMap;this.areaPath=g;return d},drawGraph:function(){this.areaPath=[];E.prototype.drawGraph.apply(this);var a=this,h=this.areaPath,u=this.options,
v=[["area","highcharts-area",this.color,u.fillColor]];this.zones.forEach(function(d,h){v.push(["zone-area-"+h,"highcharts-area highcharts-zone-area-"+h+" "+d.className,d.color||a.color,d.fillColor||u.fillColor])});v.forEach(function(d){var n=d[0],g=a[n];g?(g.endX=a.preventGraphAnimation?null:h.xMap,g.animate({d:h})):(g={zIndex:0},a.chart.styledMode||(g.fill=G(d[3],y(d[2]).setOpacity(G(u.fillOpacity,.75)).get())),g=a[n]=a.chart.renderer.path(h).addClass(d[1]).attr(g).add(a.group),g.isArea=!0);g.startX=
h.xMap;g.shiftUnit=u.step?2:1})},drawLegendSymbol:a.LegendSymbolMixin.drawRectangle})})(J);(function(a){var y=a.pick;a=a.seriesType;a("spline","line",{},{getPointSpline:function(a,E,h){var d=E.plotX,r=E.plotY,u=a[h-1];h=a[h+1];var v,w,n,g;if(u&&!u.isNull&&!1!==u.doCurve&&!E.isCliff&&h&&!h.isNull&&!1!==h.doCurve&&!E.isCliff){a=u.plotY;n=h.plotX;h=h.plotY;var c=0;v=(1.5*d+u.plotX)/2.5;w=(1.5*r+a)/2.5;n=(1.5*d+n)/2.5;g=(1.5*r+h)/2.5;n!==v&&(c=(g-w)*(n-d)/(n-v)+r-g);w+=c;g+=c;w>a&&w>r?(w=Math.max(a,r),
g=2*r-w):w<a&&w<r&&(w=Math.min(a,r),g=2*r-w);g>h&&g>r?(g=Math.max(h,r),w=2*r-g):g<h&&g<r&&(g=Math.min(h,r),w=2*r-g);E.rightContX=n;E.rightContY=g}E=["C",y(u.rightContX,u.plotX),y(u.rightContY,u.plotY),y(v,d),y(w,r),d,r];u.rightContX=u.rightContY=null;return E}})})(J);(function(a){var y=a.seriesTypes.area.prototype,G=a.seriesType;G("areaspline","spline",a.defaultPlotOptions.area,{getStackPoints:y.getStackPoints,getGraphPath:y.getGraphPath,drawGraph:y.drawGraph,drawLegendSymbol:a.LegendSymbolMixin.drawRectangle})})(J);
(function(a){var y=a.animObject,G=a.color,E=a.extend,h=a.defined,d=a.isNumber,r=a.merge,u=a.pick,v=a.Series,w=a.seriesType,n=a.svg;w("column","line",{borderRadius:0,crisp:!0,groupPadding:.2,marker:null,pointPadding:.1,minPointLength:0,cropThreshold:50,pointRange:null,states:{hover:{halo:!1,brightness:.1},select:{color:"#cccccc",borderColor:"#000000"}},dataLabels:{align:null,verticalAlign:null,y:null},softThreshold:!1,startFromThreshold:!0,stickyTracking:!1,tooltip:{distance:6},threshold:0,borderColor:"#ffffff"},
{cropShoulder:0,directTouch:!0,trackerGroups:["group","dataLabelsGroup"],negStacks:!0,init:function(){v.prototype.init.apply(this,arguments);var a=this,c=a.chart;c.hasRendered&&c.series.forEach(function(c){c.type===a.type&&(c.isDirty=!0)})},getColumnMetrics:function(){var a=this,c=a.options,d=a.xAxis,h=a.yAxis,b=d.options.reversedStacks,b=d.reversed&&!b||!d.reversed&&b,l,f={},n=0;!1===c.grouping?n=1:a.chart.series.forEach(function(b){var c=b.options,d=b.yAxis,g;b.type!==a.type||!b.visible&&a.chart.options.chart.ignoreHiddenSeries||
h.len!==d.len||h.pos!==d.pos||(c.stacking?(l=b.stackKey,void 0===f[l]&&(f[l]=n++),g=f[l]):!1!==c.grouping&&(g=n++),b.columnIndex=g)});var t=Math.min(Math.abs(d.transA)*(d.ordinalSlope||c.pointRange||d.closestPointRange||d.tickInterval||1),d.len),r=t*c.groupPadding,v=(t-2*r)/(n||1),c=Math.min(c.maxPointWidth||d.len,u(c.pointWidth,v*(1-2*c.pointPadding)));a.columnMetrics={width:c,offset:(v-c)/2+(r+((a.columnIndex||0)+(b?1:0))*v-t/2)*(b?-1:1)};return a.columnMetrics},crispCol:function(a,c,d,h){var b=
this.chart,g=this.borderWidth,f=-(g%2?.5:0),g=g%2?.5:1;b.inverted&&b.renderer.isVML&&(g+=1);this.options.crisp&&(d=Math.round(a+d)+f,a=Math.round(a)+f,d-=a);h=Math.round(c+h)+g;f=.5>=Math.abs(c)&&.5<h;c=Math.round(c)+g;h-=c;f&&h&&(--c,h+=1);return{x:a,y:c,width:d,height:h}},translate:function(){var a=this,c=a.chart,d=a.options,p=a.dense=2>a.closestPointRange*a.xAxis.transA,p=a.borderWidth=u(d.borderWidth,p?0:1),b=a.yAxis,l=d.threshold,f=a.translatedThreshold=b.getThreshold(l),n=u(d.minPointLength,
5),t=a.getColumnMetrics(),r=t.width,F=a.barW=Math.max(r,1+2*p),z=a.pointXOffset=t.offset;c.inverted&&(f-=.5);d.pointPadding&&(F=Math.ceil(F));v.prototype.translate.apply(a);a.points.forEach(function(d){var g=u(d.yBottom,f),k=999+Math.abs(g),m=r,k=Math.min(Math.max(-k,d.plotY),b.len+k),e=d.plotX+z,q=F,p=Math.min(k,g),t,x=Math.max(k,g)-p;n&&Math.abs(x)<n&&(x=n,t=!b.reversed&&!d.negative||b.reversed&&d.negative,d.y===l&&a.dataMax<=l&&b.min<l&&(t=!t),p=Math.abs(p-f)>n?g-n:f-(t?n:0));h(d.options.pointWidth)&&
(m=q=Math.ceil(d.options.pointWidth),e-=Math.round((m-r)/2));d.barX=e;d.pointWidth=m;d.tooltipPos=c.inverted?[b.len+b.pos-c.plotLeft-k,a.xAxis.len-e-q/2,x]:[e+q/2,k+b.pos-c.plotTop,x];d.shapeType=d.shapeType||"rect";d.shapeArgs=a.crispCol.apply(a,d.isNull?[e,f,q,0]:[e,p,q,x])})},getSymbol:a.noop,drawLegendSymbol:a.LegendSymbolMixin.drawRectangle,drawGraph:function(){this.group[this.dense?"addClass":"removeClass"]("highcharts-dense-data")},pointAttribs:function(a,c){var d=this.options,g,b=this.pointAttrToOptions||
{};g=b.stroke||"borderColor";var l=b["stroke-width"]||"borderWidth",f=a&&a.color||this.color,h=a&&a[g]||d[g]||this.color||f,t=a&&a[l]||d[l]||this[l]||0,b=d.dashStyle;a&&this.zones.length&&(f=a.getZone(),f=a.options.color||f&&f.color||this.color);c&&(a=r(d.states[c],a.options.states&&a.options.states[c]||{}),c=a.brightness,f=a.color||void 0!==c&&G(f).brighten(a.brightness).get()||f,h=a[g]||h,t=a[l]||t,b=a.dashStyle||b);g={fill:f,stroke:h,"stroke-width":t};b&&(g.dashstyle=b);return g},drawPoints:function(){var a=
this,c=this.chart,h=a.options,p=c.renderer,b=h.animationLimit||250,l;a.points.forEach(function(f){var g=f.graphic,m=g&&c.pointCount<b?"animate":"attr";if(d(f.plotY)&&null!==f.y){l=f.shapeArgs;if(g)g[m](r(l));else f.graphic=g=p[f.shapeType](l).add(f.group||a.group);h.borderRadius&&g.attr({r:h.borderRadius});c.styledMode||g[m](a.pointAttribs(f,f.selected&&"select")).shadow(h.shadow,null,h.stacking&&!h.borderRadius);g.addClass(f.getClassName(),!0)}else g&&(f.graphic=g.destroy())})},animate:function(a){var c=
this,d=this.yAxis,g=c.options,b=this.chart.inverted,h={},f=b?"translateX":"translateY",x;n&&(a?(h.scaleY=.001,a=Math.min(d.pos+d.len,Math.max(d.pos,d.toPixels(g.threshold))),b?h.translateX=a-d.len:h.translateY=a,c.group.attr(h)):(x=c.group.attr(f),c.group.animate({scaleY:1},E(y(c.options.animation),{step:function(a,b){h[f]=x+b.pos*(d.pos-x);c.group.attr(h)}})),c.animate=null))},remove:function(){var a=this,c=a.chart;c.hasRendered&&c.series.forEach(function(c){c.type===a.type&&(c.isDirty=!0)});v.prototype.remove.apply(a,
arguments)}})})(J);(function(a){a=a.seriesType;a("bar","column",null,{inverted:!0})})(J);(function(a){var y=a.Series;a=a.seriesType;a("scatter","line",{lineWidth:0,findNearestPointBy:"xy",marker:{enabled:!0},tooltip:{headerFormat:'\x3cspan style\x3d"color:{point.color}"\x3e\u25cf\x3c/span\x3e \x3cspan style\x3d"font-size: 10px"\x3e {series.name}\x3c/span\x3e\x3cbr/\x3e',pointFormat:"x: \x3cb\x3e{point.x}\x3c/b\x3e\x3cbr/\x3ey: \x3cb\x3e{point.y}\x3c/b\x3e\x3cbr/\x3e"}},{sorted:!1,requireSorting:!1,
noSharedTooltip:!0,trackerGroups:["group","markerGroup","dataLabelsGroup"],takeOrdinalPosition:!1,drawGraph:function(){this.options.lineWidth&&y.prototype.drawGraph.call(this)}})})(J);(function(a){var y=a.deg2rad,G=a.isNumber,E=a.pick,h=a.relativeLength;a.CenteredSeriesMixin={getCenter:function(){var a=this.options,r=this.chart,u=2*(a.slicedOffset||0),v=r.plotWidth-2*u,r=r.plotHeight-2*u,w=a.center,w=[E(w[0],"50%"),E(w[1],"50%"),a.size||"100%",a.innerSize||0],n=Math.min(v,r),g,c;for(g=0;4>g;++g)c=
w[g],a=2>g||2===g&&/%$/.test(c),w[g]=h(c,[v,r,n,w[2]][g])+(a?u:0);w[3]>w[2]&&(w[3]=w[2]);return w},getStartAndEndRadians:function(a,h){a=G(a)?a:0;h=G(h)&&h>a&&360>h-a?h:a+360;return{start:y*(a+-90),end:y*(h+-90)}}}})(J);(function(a){var y=a.addEvent,G=a.CenteredSeriesMixin,E=a.defined,h=a.extend,d=G.getStartAndEndRadians,r=a.noop,u=a.pick,v=a.Point,w=a.Series,n=a.seriesType,g=a.setAnimation;n("pie","line",{center:[null,null],clip:!1,colorByPoint:!0,dataLabels:{allowOverlap:!0,connectorPadding:5,distance:30,
enabled:!0,formatter:function(){return this.point.isNull?void 0:this.point.name},softConnector:!0,x:0,connectorShape:"fixedOffset",crookDistance:"70%"},ignoreHiddenPoint:!0,legendType:"point",marker:null,size:null,showInLegend:!1,slicedOffset:10,stickyTracking:!1,tooltip:{followPointer:!0},borderColor:"#ffffff",borderWidth:1,states:{hover:{brightness:.1}}},{isCartesian:!1,requireSorting:!1,directTouch:!0,noSharedTooltip:!0,trackerGroups:["group","dataLabelsGroup"],axisTypes:[],pointAttribs:a.seriesTypes.column.prototype.pointAttribs,
animate:function(a){var c=this,d=c.points,b=c.startAngleRad;a||(d.forEach(function(a){var d=a.graphic,g=a.shapeArgs;d&&(d.attr({r:a.startR||c.center[3]/2,start:b,end:b}),d.animate({r:g.r,start:g.start,end:g.end},c.options.animation))}),c.animate=null)},updateTotals:function(){var a,d=0,g=this.points,b=g.length,h,f=this.options.ignoreHiddenPoint;for(a=0;a<b;a++)h=g[a],d+=f&&!h.visible?0:h.isNull?0:h.y;this.total=d;for(a=0;a<b;a++)h=g[a],h.percentage=0<d&&(h.visible||!f)?h.y/d*100:0,h.total=d},generatePoints:function(){w.prototype.generatePoints.call(this);
this.updateTotals()},getX:function(a,d,g){var b=this.center,c=this.radii?this.radii[g.index]:b[2]/2;return b[0]+(d?-1:1)*Math.cos(Math.asin(Math.max(Math.min((a-b[1])/(c+g.labelDistance),1),-1)))*(c+g.labelDistance)+(0<g.labelDistance?(d?-1:1)*this.options.dataLabels.padding:0)},translate:function(a){this.generatePoints();var c=0,g=this.options,b=g.slicedOffset,h=b+(g.borderWidth||0),f,n,t=d(g.startAngle,g.endAngle),r=this.startAngleRad=t.start,t=(this.endAngleRad=t.end)-r,v=this.points,z,k,A=g.dataLabels.distance,
g=g.ignoreHiddenPoint,w,C=v.length,e;a||(this.center=a=this.getCenter());for(w=0;w<C;w++){e=v[w];e.labelDistance=u(e.options.dataLabels&&e.options.dataLabels.distance,A);this.maxLabelDistance=Math.max(this.maxLabelDistance||0,e.labelDistance);f=r+c*t;if(!g||e.visible)c+=e.percentage/100;n=r+c*t;e.shapeType="arc";e.shapeArgs={x:a[0],y:a[1],r:a[2]/2,innerR:a[3]/2,start:Math.round(1E3*f)/1E3,end:Math.round(1E3*n)/1E3};n=(n+f)/2;n>1.5*Math.PI?n-=2*Math.PI:n<-Math.PI/2&&(n+=2*Math.PI);e.slicedTranslation=
{translateX:Math.round(Math.cos(n)*b),translateY:Math.round(Math.sin(n)*b)};z=Math.cos(n)*a[2]/2;k=Math.sin(n)*a[2]/2;e.tooltipPos=[a[0]+.7*z,a[1]+.7*k];e.half=n<-Math.PI/2||n>Math.PI/2?1:0;e.angle=n;f=Math.min(h,e.labelDistance/5);e.labelPosition={natural:{x:a[0]+z+Math.cos(n)*e.labelDistance,y:a[1]+k+Math.sin(n)*e.labelDistance},final:{},alignment:0>e.labelDistance?"center":e.half?"right":"left",connectorPosition:{breakAt:{x:a[0]+z+Math.cos(n)*f,y:a[1]+k+Math.sin(n)*f},touchingSliceAt:{x:a[0]+z,
y:a[1]+k}}}}},drawGraph:null,drawPoints:function(){var a=this,d=a.chart,g=d.renderer,b,l,f,n,t=a.options.shadow;!t||a.shadowGroup||d.styledMode||(a.shadowGroup=g.g("shadow").add(a.group));a.points.forEach(function(c){l=c.graphic;if(c.isNull)l&&(c.graphic=l.destroy());else{n=c.shapeArgs;b=c.getTranslate();if(!d.styledMode){var m=c.shadowGroup;t&&!m&&(m=c.shadowGroup=g.g("shadow").add(a.shadowGroup));m&&m.attr(b);f=a.pointAttribs(c,c.selected&&"select")}l?(l.setRadialReference(a.center),d.styledMode||
l.attr(f),l.animate(h(n,b))):(c.graphic=l=g[c.shapeType](n).setRadialReference(a.center).attr(b).add(a.group),d.styledMode||l.attr(f).attr({"stroke-linejoin":"round"}).shadow(t,m));l.attr({visibility:c.visible?"inherit":"hidden"});l.addClass(c.getClassName())}})},searchPoint:r,sortByAngle:function(a,d){a.sort(function(a,b){return void 0!==a.angle&&(b.angle-a.angle)*d})},drawLegendSymbol:a.LegendSymbolMixin.drawRectangle,getCenter:G.getCenter,getSymbol:r},{init:function(){v.prototype.init.apply(this,
arguments);var a=this,d;a.name=u(a.name,"Slice");d=function(c){a.slice("select"===c.type)};y(a,"select",d);y(a,"unselect",d);return a},isValid:function(){return a.isNumber(this.y,!0)&&0<=this.y},setVisible:function(a,d){var c=this,b=c.series,g=b.chart,f=b.options.ignoreHiddenPoint;d=u(d,f);a!==c.visible&&(c.visible=c.options.visible=a=void 0===a?!c.visible:a,b.options.data[b.data.indexOf(c)]=c.options,["graphic","dataLabel","connector","shadowGroup"].forEach(function(b){if(c[b])c[b][a?"show":"hide"](!0)}),
c.legendItem&&g.legend.colorizeItem(c,a),a||"hover"!==c.state||c.setState(""),f&&(b.isDirty=!0),d&&g.redraw())},slice:function(a,d,h){var b=this.series;g(h,b.chart);u(d,!0);this.sliced=this.options.sliced=E(a)?a:!this.sliced;b.options.data[b.data.indexOf(this)]=this.options;this.graphic.animate(this.getTranslate());this.shadowGroup&&this.shadowGroup.animate(this.getTranslate())},getTranslate:function(){return this.sliced?this.slicedTranslation:{translateX:0,translateY:0}},haloPath:function(a){var c=
this.shapeArgs;return this.sliced||!this.visible?[]:this.series.chart.renderer.symbols.arc(c.x,c.y,c.r+a,c.r+a,{innerR:this.shapeArgs.r-1,start:c.start,end:c.end})},connectorShapes:{fixedOffset:function(a,d,g){var b=d.breakAt;d=d.touchingSliceAt;return["M",a.x,a.y].concat(g.softConnector?["C",a.x+("left"===a.alignment?-5:5),a.y,2*b.x-d.x,2*b.y-d.y,b.x,b.y]:["L",b.x,b.y]).concat(["L",d.x,d.y])},straight:function(a,d){d=d.touchingSliceAt;return["M",a.x,a.y,"L",d.x,d.y]},crookedLine:function(c,d,g){d=
d.touchingSliceAt;var b=this.series,h=b.center[0],f=b.chart.plotWidth,m=b.chart.plotLeft,b=c.alignment,t=this.shapeArgs.r;g=a.relativeLength(g.crookDistance,1);g="left"===b?h+t+(f+m-h-t)*(1-g):m+(h-t)*g;h=["L",g,c.y];if("left"===b?g>c.x||g<d.x:g<c.x||g>d.x)h=[];return["M",c.x,c.y].concat(h).concat(["L",d.x,d.y])}},getConnectorPath:function(){var a=this.labelPosition,d=this.series.options.dataLabels,g=d.connectorShape,b=this.connectorShapes;b[g]&&(g=b[g]);return g.call(this,{x:a.final.x,y:a.final.y,
alignment:a.alignment},a.connectorPosition,d)}})})(J);(function(a){var y=a.addEvent,G=a.arrayMax,E=a.defined,h=a.extend,d=a.format,r=a.merge,u=a.noop,v=a.pick,w=a.relativeLength,n=a.Series,g=a.seriesTypes,c=a.stableSort,m=a.isArray,p=a.splat;a.distribute=function(b,d,f){function g(a,b){return a.target-b.target}var h,l=!0,m=b,p=[],k;k=0;var n=m.reducedLen||d;for(h=b.length;h--;)k+=b[h].size;if(k>n){c(b,function(a,b){return(b.rank||0)-(a.rank||0)});for(k=h=0;k<=n;)k+=b[h].size,h++;p=b.splice(h-1,b.length)}c(b,
g);for(b=b.map(function(a){return{size:a.size,targets:[a.target],align:v(a.align,.5)}});l;){for(h=b.length;h--;)l=b[h],k=(Math.min.apply(0,l.targets)+Math.max.apply(0,l.targets))/2,l.pos=Math.min(Math.max(0,k-l.size*l.align),d-l.size);h=b.length;for(l=!1;h--;)0<h&&b[h-1].pos+b[h-1].size>b[h].pos&&(b[h-1].size+=b[h].size,b[h-1].targets=b[h-1].targets.concat(b[h].targets),b[h-1].align=.5,b[h-1].pos+b[h-1].size>d&&(b[h-1].pos=d-b[h-1].size),b.splice(h,1),l=!0)}m.push.apply(m,p);h=0;b.some(function(b){var c=
0;if(b.targets.some(function(){m[h].pos=b.pos+c;if(Math.abs(m[h].pos-m[h].target)>f)return m.slice(0,h+1).forEach(function(a){delete a.pos}),m.reducedLen=(m.reducedLen||d)-.1*d,m.reducedLen>.1*d&&a.distribute(m,d,f),!0;c+=m[h].size;h++}))return!0});c(m,g)};n.prototype.drawDataLabels=function(){function b(a,b){var d=b.filter;return d?(b=d.operator,a=a[d.property],d=d.value,"\x3e"===b&&a>d||"\x3c"===b&&a<d||"\x3e\x3d"===b&&a>=d||"\x3c\x3d"===b&&a<=d||"\x3d\x3d"===b&&a==d||"\x3d\x3d\x3d"===b&&a===d?
!0:!1):!0}function c(a,b){var d=[],c;if(m(a)&&!m(b))d=a.map(function(a){return r(a,b)});else if(m(b)&&!m(a))d=b.map(function(b){return r(a,b)});else if(m(a)||m(b))for(c=Math.max(a.length,b.length);c--;)d[c]=r(a[c],b[c]);else d=r(a,b);return d}var f=this,g=f.chart,h=f.options,n=h.dataLabels,u=f.points,z,k=f.hasRendered||0,A,w=v(n.defer,!!h.animation),C=g.renderer,n=c(c(g.options.plotOptions&&g.options.plotOptions.series&&g.options.plotOptions.series.dataLabels,g.options.plotOptions&&g.options.plotOptions[f.type]&&
g.options.plotOptions[f.type].dataLabels),n);a.fireEvent(this,"drawDataLabels");if(m(n)||n.enabled||f._hasPointLabels)A=f.plotGroup("dataLabelsGroup","data-labels",w&&!k?"hidden":"visible",n.zIndex||6),w&&(A.attr({opacity:+k}),k||y(f,"afterAnimate",function(){f.visible&&A.show(!0);A[h.animation?"animate":"attr"]({opacity:1},{duration:200})})),u.forEach(function(e){z=p(c(n,e.dlOptions||e.options&&e.options.dataLabels));z.forEach(function(c,k){var l=c.enabled&&!e.isNull&&b(e,c),m,q,t,p,n=e.dataLabels?
e.dataLabels[k]:e.dataLabel,r=e.connectors?e.connectors[k]:e.connector,u=!n;l&&(m=e.getLabelConfig(),q=c[e.formatPrefix+"Format"]||c.format,m=E(q)?d(q,m,g.time):(c[e.formatPrefix+"Formatter"]||c.formatter).call(m,c),q=c.style,t=c.rotation,g.styledMode||(q.color=v(c.color,q.color,f.color,"#000000"),"contrast"===q.color&&(e.contrastColor=C.getContrast(e.color||f.color),q.color=c.inside||0>v(c.distance,e.labelDistance)||h.stacking?e.contrastColor:"#000000"),h.cursor&&(q.cursor=h.cursor)),p={r:c.borderRadius||
0,rotation:t,padding:c.padding,zIndex:1},g.styledMode||(p.fill=c.backgroundColor,p.stroke=c.borderColor,p["stroke-width"]=c.borderWidth),a.objectEach(p,function(a,b){void 0===a&&delete p[b]}));!n||l&&E(m)?l&&E(m)&&(n?p.text=m:(e.dataLabels=e.dataLabels||[],n=e.dataLabels[k]=t?C.text(m,0,-9999).addClass("highcharts-data-label"):C.label(m,0,-9999,c.shape,null,null,c.useHTML,null,"data-label"),k||(e.dataLabel=n),n.addClass(" highcharts-data-label-color-"+e.colorIndex+" "+(c.className||"")+(c.useHTML?
" highcharts-tracker":""))),n.options=c,n.attr(p),g.styledMode||n.css(q).shadow(c.shadow),n.added||n.add(A),f.alignDataLabel(e,n,c,null,u)):(e.dataLabel=e.dataLabel&&e.dataLabel.destroy(),e.dataLabels&&(1===e.dataLabels.length?delete e.dataLabels:delete e.dataLabels[k]),k||delete e.dataLabel,r&&(e.connector=e.connector.destroy(),e.connectors&&(1===e.connectors.length?delete e.connectors:delete e.connectors[k])))})});a.fireEvent(this,"afterDrawDataLabels")};n.prototype.alignDataLabel=function(a,c,
d,g,m){var b=this.chart,f=this.isCartesian&&b.inverted,l=v(a.dlBox&&a.dlBox.centerX,a.plotX,-9999),k=v(a.plotY,-9999),n=c.getBBox(),p,t=d.rotation,e=d.align,q=this.visible&&(a.series.forceDL||b.isInsidePlot(l,Math.round(k),f)||g&&b.isInsidePlot(l,f?g.x+1:g.y+g.height-1,f)),r="justify"===v(d.overflow,"justify");if(q&&(p=b.renderer.fontMetrics(b.styledMode?void 0:d.style.fontSize,c).b,g=h({x:f?this.yAxis.len-k:l,y:Math.round(f?this.xAxis.len-l:k),width:0,height:0},g),h(d,{width:n.width,height:n.height}),
t?(r=!1,l=b.renderer.rotCorr(p,t),l={x:g.x+d.x+g.width/2+l.x,y:g.y+d.y+{top:0,middle:.5,bottom:1}[d.verticalAlign]*g.height},c[m?"attr":"animate"](l).attr({align:e}),k=(t+720)%360,k=180<k&&360>k,"left"===e?l.y-=k?n.height:0:"center"===e?(l.x-=n.width/2,l.y-=n.height/2):"right"===e&&(l.x-=n.width,l.y-=k?0:n.height),c.placed=!0,c.alignAttr=l):(c.align(d,null,g),l=c.alignAttr),r&&0<=g.height?a.isLabelJustified=this.justifyDataLabel(c,d,l,n,g,m):v(d.crop,!0)&&(q=b.isInsidePlot(l.x,l.y)&&b.isInsidePlot(l.x+
n.width,l.y+n.height)),d.shape&&!t))c[m?"attr":"animate"]({anchorX:f?b.plotWidth-a.plotY:a.plotX,anchorY:f?b.plotHeight-a.plotX:a.plotY});q||(c.attr({y:-9999}),c.placed=!1)};n.prototype.justifyDataLabel=function(a,c,d,g,h,m){var b=this.chart,f=c.align,k=c.verticalAlign,l,n,p=a.box?0:a.padding||0;l=d.x+p;0>l&&("right"===f?c.align="left":c.x=-l,n=!0);l=d.x+g.width-p;l>b.plotWidth&&("left"===f?c.align="right":c.x=b.plotWidth-l,n=!0);l=d.y+p;0>l&&("bottom"===k?c.verticalAlign="top":c.y=-l,n=!0);l=d.y+
g.height-p;l>b.plotHeight&&("top"===k?c.verticalAlign="bottom":c.y=b.plotHeight-l,n=!0);n&&(a.placed=!m,a.align(c,null,h));return n};g.pie&&(g.pie.prototype.dataLabelPositioners={radialDistributionY:function(a){return a.top+a.distributeBox.pos},radialDistributionX:function(a,c,d,g){return a.getX(d<c.top+2||d>c.bottom-2?g:d,c.half,c)},justify:function(a,c,d){return d[0]+(a.half?-1:1)*(c+a.labelDistance)},alignToPlotEdges:function(a,c,d,g){a=a.getBBox().width;return c?a+g:d-a-g},alignToConnectors:function(a,
c,d,g){var b=0,f;a.forEach(function(a){f=a.dataLabel.getBBox().width;f>b&&(b=f)});return c?b+g:d-b-g}},g.pie.prototype.drawDataLabels=function(){var b=this,c=b.data,d,g=b.chart,h=b.options.dataLabels,m=h.connectorPadding,p=v(h.connectorWidth,1),r=g.plotWidth,k=g.plotHeight,u=g.plotLeft,w=Math.round(g.chartWidth/3),y,e=b.center,q=e[2]/2,L=e[1],I,J,K,M,S=[[],[]],B,N,P,T,Q=[0,0,0,0],O=b.dataLabelPositioners;b.visible&&(h.enabled||b._hasPointLabels)&&(c.forEach(function(a){a.dataLabel&&a.visible&&a.dataLabel.shortened&&
(a.dataLabel.attr({width:"auto"}).css({width:"auto",textOverflow:"clip"}),a.dataLabel.shortened=!1)}),n.prototype.drawDataLabels.apply(b),c.forEach(function(a){a.dataLabel&&(a.visible?(S[a.half].push(a),a.dataLabel._pos=null,!E(h.style.width)&&!E(a.options.dataLabels&&a.options.dataLabels.style&&a.options.dataLabels.style.width)&&a.dataLabel.getBBox().width>w&&(a.dataLabel.css({width:.7*w}),a.dataLabel.shortened=!0)):(a.dataLabel=a.dataLabel.destroy(),a.dataLabels&&1===a.dataLabels.length&&delete a.dataLabels))}),
S.forEach(function(c,f){var l,n,p=c.length,t=[],x;if(p)for(b.sortByAngle(c,f-.5),0<b.maxLabelDistance&&(l=Math.max(0,L-q-b.maxLabelDistance),n=Math.min(L+q+b.maxLabelDistance,g.plotHeight),c.forEach(function(a){0<a.labelDistance&&a.dataLabel&&(a.top=Math.max(0,L-q-a.labelDistance),a.bottom=Math.min(L+q+a.labelDistance,g.plotHeight),x=a.dataLabel.getBBox().height||21,a.distributeBox={target:a.labelPosition.natural.y-a.top+x/2,size:x,rank:a.y},t.push(a.distributeBox))}),l=n+x-l,a.distribute(t,l,l/5)),
T=0;T<p;T++){d=c[T];K=d.labelPosition;I=d.dataLabel;P=!1===d.visible?"hidden":"inherit";N=l=K.natural.y;t&&E(d.distributeBox)&&(void 0===d.distributeBox.pos?P="hidden":(M=d.distributeBox.size,N=O.radialDistributionY(d)));delete d.positionIndex;if(h.justify)B=O.justify(d,q,e);else switch(h.alignTo){case "connectors":B=O.alignToConnectors(c,f,r,u);break;case "plotEdges":B=O.alignToPlotEdges(I,f,r,u);break;default:B=O.radialDistributionX(b,d,N,l)}I._attr={visibility:P,align:K.alignment};I._pos={x:B+
h.x+({left:m,right:-m}[K.alignment]||0),y:N+h.y-10};K.final.x=B;K.final.y=N;v(h.crop,!0)&&(J=I.getBBox().width,l=null,B-J<m&&1===f?(l=Math.round(J-B+m),Q[3]=Math.max(l,Q[3])):B+J>r-m&&0===f&&(l=Math.round(B+J-r+m),Q[1]=Math.max(l,Q[1])),0>N-M/2?Q[0]=Math.max(Math.round(-N+M/2),Q[0]):N+M/2>k&&(Q[2]=Math.max(Math.round(N+M/2-k),Q[2])),I.sideOverflow=l)}}),0===G(Q)||this.verifyDataLabelOverflow(Q))&&(this.placeDataLabels(),p&&this.points.forEach(function(a){var c;y=a.connector;if((I=a.dataLabel)&&I._pos&&
a.visible&&0<a.labelDistance){P=I._attr.visibility;if(c=!y)a.connector=y=g.renderer.path().addClass("highcharts-data-label-connector  highcharts-color-"+a.colorIndex+(a.className?" "+a.className:"")).add(b.dataLabelsGroup),g.styledMode||y.attr({"stroke-width":p,stroke:h.connectorColor||a.color||"#666666"});y[c?"attr":"animate"]({d:a.getConnectorPath()});y.attr("visibility",P)}else y&&(a.connector=y.destroy())}))},g.pie.prototype.placeDataLabels=function(){this.points.forEach(function(a){var b=a.dataLabel;
b&&a.visible&&((a=b._pos)?(b.sideOverflow&&(b._attr.width=b.getBBox().width-b.sideOverflow,b.css({width:b._attr.width+"px",textOverflow:(this.options.dataLabels.style||{}).textOverflow||"ellipsis"}),b.shortened=!0),b.attr(b._attr),b[b.moved?"animate":"attr"](a),b.moved=!0):b&&b.attr({y:-9999}))},this)},g.pie.prototype.alignDataLabel=u,g.pie.prototype.verifyDataLabelOverflow=function(a){var b=this.center,c=this.options,d=c.center,g=c.minSize||80,h,m=null!==c.size;m||(null!==d[0]?h=Math.max(b[2]-Math.max(a[1],
a[3]),g):(h=Math.max(b[2]-a[1]-a[3],g),b[0]+=(a[3]-a[1])/2),null!==d[1]?h=Math.max(Math.min(h,b[2]-Math.max(a[0],a[2])),g):(h=Math.max(Math.min(h,b[2]-a[0]-a[2]),g),b[1]+=(a[0]-a[2])/2),h<b[2]?(b[2]=h,b[3]=Math.min(w(c.innerSize||0,h),h),this.translate(b),this.drawDataLabels&&this.drawDataLabels()):m=!0);return m});g.column&&(g.column.prototype.alignDataLabel=function(a,c,d,g,h){var b=this.chart.inverted,f=a.series,l=a.dlBox||a.shapeArgs,k=v(a.below,a.plotY>v(this.translatedThreshold,f.yAxis.len)),
m=v(d.inside,!!this.options.stacking);l&&(g=r(l),0>g.y&&(g.height+=g.y,g.y=0),l=g.y+g.height-f.yAxis.len,0<l&&(g.height-=l),b&&(g={x:f.yAxis.len-g.y-g.height,y:f.xAxis.len-g.x-g.width,width:g.height,height:g.width}),m||(b?(g.x+=k?0:g.width,g.width=0):(g.y+=k?g.height:0,g.height=0)));d.align=v(d.align,!b||m?"center":k?"right":"left");d.verticalAlign=v(d.verticalAlign,b||m?"middle":k?"top":"bottom");n.prototype.alignDataLabel.call(this,a,c,d,g,h);a.isLabelJustified&&a.contrastColor&&c.css({color:a.contrastColor})})})(J);
(function(a){var y=a.Chart,G=a.isArray,E=a.objectEach,h=a.pick,d=a.addEvent,r=a.fireEvent;d(y,"render",function(){var a=[];(this.labelCollectors||[]).forEach(function(d){a=a.concat(d())});(this.yAxis||[]).forEach(function(d){d.options.stackLabels&&!d.options.stackLabels.allowOverlap&&E(d.stacks,function(d){E(d,function(d){a.push(d.label)})})});(this.series||[]).forEach(function(d){var r=d.options.dataLabels;d.visible&&(!1!==r.enabled||d._hasPointLabels)&&d.points.forEach(function(d){d.visible&&(G(d.dataLabels)?
d.dataLabels:d.dataLabel?[d.dataLabel]:[]).forEach(function(g){var c=g.options;g.labelrank=h(c.labelrank,d.labelrank,d.shapeArgs&&d.shapeArgs.height);c.allowOverlap||a.push(g)})})});this.hideOverlappingLabels(a)});y.prototype.hideOverlappingLabels=function(a){var d=this,h=a.length,n=d.renderer,g,c,m,p,b,l,f=function(a,b,d,c,f,g,h,l){return!(f>a+d||f+h<a||g>b+c||g+l<b)};m=function(a){var b,d,c,f=a.box?0:a.padding||0;c=0;if(a&&(!a.alignAttr||a.placed))return b=a.alignAttr||{x:a.attr("x"),y:a.attr("y")},
d=a.parentGroup,a.width||(c=a.getBBox(),a.width=c.width,a.height=c.height,c=n.fontMetrics(null,a.element).h),{x:b.x+(d.translateX||0)+f,y:b.y+(d.translateY||0)+f-c,width:a.width-2*f,height:a.height-2*f}};for(c=0;c<h;c++)if(g=a[c])g.oldOpacity=g.opacity,g.newOpacity=1,g.absoluteBox=m(g);a.sort(function(a,b){return(b.labelrank||0)-(a.labelrank||0)});for(c=0;c<h;c++)for(l=(m=a[c])&&m.absoluteBox,g=c+1;g<h;++g)if(b=(p=a[g])&&p.absoluteBox,l&&b&&m!==p&&0!==m.newOpacity&&0!==p.newOpacity&&(b=f(l.x,l.y,
l.width,l.height,b.x,b.y,b.width,b.height)))(m.labelrank<p.labelrank?m:p).newOpacity=0;a.forEach(function(a){var b,c;a&&(c=a.newOpacity,a.oldOpacity!==c&&(a.alignAttr&&a.placed?(c?a.show(!0):b=function(){a.hide()},a.alignAttr.opacity=c,a[a.isOld?"animate":"attr"](a.alignAttr,null,b),r(d,"afterHideOverlappingLabels")):a.attr({opacity:c})),a.isOld=!0)})}})(J);(function(a){var y=a.addEvent,G=a.Chart,E=a.createElement,h=a.css,d=a.defaultOptions,r=a.defaultPlotOptions,u=a.extend,v=a.fireEvent,w=a.hasTouch,
n=a.isObject,g=a.Legend,c=a.merge,m=a.pick,p=a.Point,b=a.Series,l=a.seriesTypes,f=a.svg,x;x=a.TrackerMixin={drawTrackerPoint:function(){var a=this,b=a.chart,d=b.pointer,c=function(a){var b=d.getPointFromEvent(a);void 0!==b&&(d.isDirectTouch=!0,b.onMouseOver(a))};a.points.forEach(function(a){a.graphic&&(a.graphic.element.point=a);a.dataLabel&&(a.dataLabel.div?a.dataLabel.div.point=a:a.dataLabel.element.point=a)});a._hasTracking||(a.trackerGroups.forEach(function(f){if(a[f]){a[f].addClass("highcharts-tracker").on("mouseover",
c).on("mouseout",function(a){d.onTrackerMouseOut(a)});if(w)a[f].on("touchstart",c);!b.styledMode&&a.options.cursor&&a[f].css(h).css({cursor:a.options.cursor})}}),a._hasTracking=!0);v(this,"afterDrawTracker")},drawTrackerGraph:function(){var a=this,b=a.options,d=b.trackByArea,c=[].concat(d?a.areaPath:a.graphPath),g=c.length,h=a.chart,l=h.pointer,m=h.renderer,e=h.options.tooltip.snap,q=a.tracker,p,n=function(){if(h.hoverSeries!==a)a.onMouseOver()},r="rgba(192,192,192,"+(f?.0001:.002)+")";if(g&&!d)for(p=
g+1;p--;)"M"===c[p]&&c.splice(p+1,0,c[p+1]-e,c[p+2],"L"),(p&&"M"===c[p]||p===g)&&c.splice(p,0,"L",c[p-2]+e,c[p-1]);q?q.attr({d:c}):a.graph&&(a.tracker=m.path(c).attr({visibility:a.visible?"visible":"hidden",zIndex:2}).addClass(d?"highcharts-tracker-area":"highcharts-tracker-line").add(a.group),h.styledMode||a.tracker.attr({"stroke-linejoin":"round",stroke:r,fill:d?r:"none","stroke-width":a.graph.strokeWidth()+(d?0:2*e)}),[a.tracker,a.markerGroup].forEach(function(a){a.addClass("highcharts-tracker").on("mouseover",
n).on("mouseout",function(a){l.onTrackerMouseOut(a)});b.cursor&&!h.styledMode&&a.css({cursor:b.cursor});if(w)a.on("touchstart",n)}));v(this,"afterDrawTracker")}};l.column&&(l.column.prototype.drawTracker=x.drawTrackerPoint);l.pie&&(l.pie.prototype.drawTracker=x.drawTrackerPoint);l.scatter&&(l.scatter.prototype.drawTracker=x.drawTrackerPoint);u(g.prototype,{setItemEvents:function(a,b,d){var f=this,g=f.chart.renderer.boxWrapper,h="highcharts-legend-"+(a instanceof p?"point":"series")+"-active",l=f.chart.styledMode;
(d?b:a.legendGroup).on("mouseover",function(){a.setState("hover");g.addClass(h);l||b.css(f.options.itemHoverStyle)}).on("mouseout",function(){f.styledMode||b.css(c(a.visible?f.itemStyle:f.itemHiddenStyle));g.removeClass(h);a.setState()}).on("click",function(b){var d=function(){a.setVisible&&a.setVisible()};g.removeClass(h);b={browserEvent:b};a.firePointEvent?a.firePointEvent("legendItemClick",b,d):v(a,"legendItemClick",b,d)})},createCheckboxForItem:function(a){a.checkbox=E("input",{type:"checkbox",
className:"highcharts-legend-checkbox",checked:a.selected,defaultChecked:a.selected},this.options.itemCheckboxStyle,this.chart.container);y(a.checkbox,"click",function(b){v(a.series||a,"checkboxClick",{checked:b.target.checked,item:a},function(){a.select()})})}});u(G.prototype,{showResetZoom:function(){function a(){b.zoomOut()}var b=this,c=d.lang,f=b.options.chart.resetZoomButton,g=f.theme,h=g.states,l="chart"===f.relativeTo?null:"plotBox";v(this,"beforeShowResetZoom",null,function(){b.resetZoomButton=
b.renderer.button(c.resetZoom,null,null,a,g,h&&h.hover).attr({align:f.position.align,title:c.resetZoomTitle}).addClass("highcharts-reset-zoom").add().align(f.position,!1,l)})},zoomOut:function(){v(this,"selection",{resetSelection:!0},this.zoom)},zoom:function(a){var b,d=this.pointer,c=!1,f;!a||a.resetSelection?(this.axes.forEach(function(a){b=a.zoom()}),d.initiated=!1):a.xAxis.concat(a.yAxis).forEach(function(a){var f=a.axis;d[f.isXAxis?"zoomX":"zoomY"]&&(b=f.zoom(a.min,a.max),f.displayBtn&&(c=!0))});
f=this.resetZoomButton;c&&!f?this.showResetZoom():!c&&n(f)&&(this.resetZoomButton=f.destroy());b&&this.redraw(m(this.options.chart.animation,a&&a.animation,100>this.pointCount))},pan:function(a,b){var c=this,d=c.hoverPoints,f;d&&d.forEach(function(a){a.setState()});("xy"===b?[1,0]:[1]).forEach(function(b){b=c[b?"xAxis":"yAxis"][0];var d=b.horiz,g=a[d?"chartX":"chartY"],d=d?"mouseDownX":"mouseDownY",e=c[d],h=(b.pointRange||0)/2,k=b.reversed&&!c.inverted||!b.reversed&&c.inverted?-1:1,l=b.getExtremes(),
m=b.toValue(e-g,!0)+h*k,k=b.toValue(e+b.len-g,!0)-h*k,p=k<m,e=p?k:m,m=p?m:k,k=Math.min(l.dataMin,h?l.min:b.toValue(b.toPixels(l.min)-b.minPixelPadding)),h=Math.max(l.dataMax,h?l.max:b.toValue(b.toPixels(l.max)+b.minPixelPadding)),p=k-e;0<p&&(m+=p,e=k);p=m-h;0<p&&(m=h,e-=p);b.series.length&&e!==l.min&&m!==l.max&&(b.setExtremes(e,m,!1,!1,{trigger:"pan"}),f=!0);c[d]=g});f&&c.redraw(!1);h(c.container,{cursor:"move"})}});u(p.prototype,{select:function(a,b){var c=this,d=c.series,f=d.chart;a=m(a,!c.selected);
c.firePointEvent(a?"select":"unselect",{accumulate:b},function(){c.selected=c.options.selected=a;d.options.data[d.data.indexOf(c)]=c.options;c.setState(a&&"select");b||f.getSelectedPoints().forEach(function(a){a.selected&&a!==c&&(a.selected=a.options.selected=!1,d.options.data[d.data.indexOf(a)]=a.options,a.setState(""),a.firePointEvent("unselect"))})})},onMouseOver:function(a){var b=this.series.chart,c=b.pointer;a=a?c.normalize(a):c.getChartCoordinatesFromPoint(this,b.inverted);c.runPointActions(a,
this)},onMouseOut:function(){var a=this.series.chart;this.firePointEvent("mouseOut");(a.hoverPoints||[]).forEach(function(a){a.setState()});a.hoverPoints=a.hoverPoint=null},importEvents:function(){if(!this.hasImportedEvents){var b=this,d=c(b.series.options.point,b.options).events;b.events=d;a.objectEach(d,function(a,c){y(b,c,a)});this.hasImportedEvents=!0}},setState:function(a,b){var c=Math.floor(this.plotX),d=this.plotY,f=this.series,g=f.options.states[a||"normal"]||{},h=r[f.type].marker&&f.options.marker,
l=h&&!1===h.enabled,e=h&&h.states&&h.states[a||"normal"]||{},p=!1===e.enabled,n=f.stateMarkerGraphic,t=this.marker||{},w=f.chart,x=f.halo,y,E=h&&f.markerAttribs;a=a||"";if(!(a===this.state&&!b||this.selected&&"select"!==a||!1===g.enabled||a&&(p||l&&!1===e.enabled)||a&&t.states&&t.states[a]&&!1===t.states[a].enabled)){E&&(y=f.markerAttribs(this,a));if(this.graphic)this.state&&this.graphic.removeClass("highcharts-point-"+this.state),a&&this.graphic.addClass("highcharts-point-"+a),w.styledMode||this.graphic.animate(f.pointAttribs(this,
a),m(w.options.chart.animation,g.animation)),y&&this.graphic.animate(y,m(w.options.chart.animation,e.animation,h.animation)),n&&n.hide();else{if(a&&e){h=t.symbol||f.symbol;n&&n.currentSymbol!==h&&(n=n.destroy());if(n)n[b?"animate":"attr"]({x:y.x,y:y.y});else h&&(f.stateMarkerGraphic=n=w.renderer.symbol(h,y.x,y.y,y.width,y.height).add(f.markerGroup),n.currentSymbol=h);!w.styledMode&&n&&n.attr(f.pointAttribs(this,a))}n&&(n[a&&w.isInsidePlot(c,d,w.inverted)?"show":"hide"](),n.element.point=this)}(c=
g.halo)&&c.size?(x||(f.halo=x=w.renderer.path().add((this.graphic||n).parentGroup)),x.show()[b?"animate":"attr"]({d:this.haloPath(c.size)}),x.attr({"class":"highcharts-halo highcharts-color-"+m(this.colorIndex,f.colorIndex)+(this.className?" "+this.className:""),zIndex:-1}),x.point=this,w.styledMode||x.attr(u({fill:this.color||f.color,"fill-opacity":c.opacity},c.attributes))):x&&x.point&&x.point.haloPath&&x.animate({d:x.point.haloPath(0)},null,x.hide);this.state=a;v(this,"afterSetState")}},haloPath:function(a){return this.series.chart.renderer.symbols.circle(Math.floor(this.plotX)-
a,this.plotY-a,2*a,2*a)}});u(b.prototype,{onMouseOver:function(){var a=this.chart,b=a.hoverSeries;if(b&&b!==this)b.onMouseOut();this.options.events.mouseOver&&v(this,"mouseOver");this.setState("hover");a.hoverSeries=this},onMouseOut:function(){var a=this.options,b=this.chart,c=b.tooltip,d=b.hoverPoint;b.hoverSeries=null;if(d)d.onMouseOut();this&&a.events.mouseOut&&v(this,"mouseOut");!c||this.stickyTracking||c.shared&&!this.noSharedTooltip||c.hide();this.setState()},setState:function(a){var b=this,
c=b.options,d=b.graph,f=c.states,g=c.lineWidth,c=0;a=a||"";if(b.state!==a&&([b.group,b.markerGroup,b.dataLabelsGroup].forEach(function(c){c&&(b.state&&c.removeClass("highcharts-series-"+b.state),a&&c.addClass("highcharts-series-"+a))}),b.state=a,!(b.chart.styledMode||f[a]&&!1===f[a].enabled)&&(a&&(g=f[a].lineWidth||g+(f[a].lineWidthPlus||0)),d&&!d.dashstyle)))for(g={"stroke-width":g},d.animate(g,m(f[a||"normal"]&&f[a||"normal"].animation,b.chart.options.chart.animation));b["zone-graph-"+c];)b["zone-graph-"+
c].attr(g),c+=1},setVisible:function(a,b){var c=this,d=c.chart,f=c.legendItem,g,h=d.options.chart.ignoreHiddenSeries,l=c.visible;g=(c.visible=a=c.options.visible=c.userOptions.visible=void 0===a?!l:a)?"show":"hide";["group","dataLabelsGroup","markerGroup","tracker","tt"].forEach(function(a){if(c[a])c[a][g]()});if(d.hoverSeries===c||(d.hoverPoint&&d.hoverPoint.series)===c)c.onMouseOut();f&&d.legend.colorizeItem(c,a);c.isDirty=!0;c.options.stacking&&d.series.forEach(function(a){a.options.stacking&&
a.visible&&(a.isDirty=!0)});c.linkedSeries.forEach(function(b){b.setVisible(a,!1)});h&&(d.isDirtyBox=!0);v(c,g);!1!==b&&d.redraw()},show:function(){this.setVisible(!0)},hide:function(){this.setVisible(!1)},select:function(a){this.selected=a=this.options.selected=void 0===a?!this.selected:a;this.checkbox&&(this.checkbox.checked=a);v(this,a?"select":"unselect")},drawTracker:x.drawTrackerGraph})})(J);(function(a){var y=a.Chart,G=a.isArray,E=a.isObject,h=a.pick,d=a.splat;y.prototype.setResponsive=function(d){var h=
this.options.responsive,r=[],w=this.currentResponsive;h&&h.rules&&h.rules.forEach(function(g){void 0===g._id&&(g._id=a.uniqueKey());this.matchResponsiveRule(g,r,d)},this);var n=a.merge.apply(0,r.map(function(d){return a.find(h.rules,function(a){return a._id===d}).chartOptions})),r=r.toString()||void 0;r!==(w&&w.ruleIds)&&(w&&this.update(w.undoOptions,d),r?(this.currentResponsive={ruleIds:r,mergedOptions:n,undoOptions:this.currentOptions(n)},this.update(n,d)):this.currentResponsive=void 0)};y.prototype.matchResponsiveRule=
function(a,d){var r=a.condition;(r.callback||function(){return this.chartWidth<=h(r.maxWidth,Number.MAX_VALUE)&&this.chartHeight<=h(r.maxHeight,Number.MAX_VALUE)&&this.chartWidth>=h(r.minWidth,0)&&this.chartHeight>=h(r.minHeight,0)}).call(this)&&d.push(a._id)};y.prototype.currentOptions=function(h){function r(h,n,g,c){var m;a.objectEach(h,function(a,b){if(!c&&-1<["series","xAxis","yAxis"].indexOf(b))for(a=d(a),g[b]=[],m=0;m<a.length;m++)n[b][m]&&(g[b][m]={},r(a[m],n[b][m],g[b][m],c+1));else E(a)?
(g[b]=G(a)?[]:{},r(a,n[b]||{},g[b],c+1)):g[b]=n[b]||null})}var v={};r(h,this.options,v,0);return v}})(J);return J});
//# sourceMappingURL=highcharts.js.map
`

var highchartsDarkUnicaThemeJS string = `/* *
 *
 *  (c) 2010-2019 Torstein Honsi
 *
 *  License: www.highcharts.com/license
 *
 *  Dark theme for Highcharts JS
 *
 * */

'use strict';

/* global document */


Highcharts.theme = {
    colors: ['#2b908f', '#90ee7e', '#f45b5b', '#7798BF', '#aaeeee', '#ff0066',
        '#eeaaee', '#55BF3B', '#DF5353', '#7798BF', '#aaeeee'],
    chart: {
        backgroundColor: {
            linearGradient: { x1: 0, y1: 0, x2: 1, y2: 1 },
            stops: [
                [0, '#2a2a2b'],
                [1, '#3e3e40']
            ]
        },
        style: {
            fontFamily: '\'Unica One\', sans-serif'
        },
        plotBorderColor: '#606063'
    },
    title: {
        style: {
            color: '#E0E0E3',
            textTransform: 'uppercase',
            fontSize: '20px'
        }
    },
    subtitle: {
        style: {
            color: '#E0E0E3',
            textTransform: 'uppercase'
        }
    },
    xAxis: {
        gridLineColor: '#707073',
        labels: {
            style: {
                color: '#E0E0E3'
            }
        },
        lineColor: '#707073',
        minorGridLineColor: '#505053',
        tickColor: '#707073',
        title: {
            style: {
                color: '#A0A0A3'

            }
        }
    },
    yAxis: {
        gridLineColor: '#707073',
        labels: {
            style: {
                color: '#E0E0E3'
            }
        },
        lineColor: '#707073',
        minorGridLineColor: '#505053',
        tickColor: '#707073',
        tickWidth: 1,
        title: {
            style: {
                color: '#A0A0A3'
            }
        }
    },
    tooltip: {
        backgroundColor: 'rgba(0, 0, 0, 0.85)',
        style: {
            color: '#F0F0F0'
        }
    },
    plotOptions: {
        series: {
            dataLabels: {
                color: '#F0F0F3',
                style: {
                    fontSize: '13px'
                }
            },
            marker: {
                lineColor: '#333'
            }
        },
        boxplot: {
            fillColor: '#505053'
        },
        candlestick: {
            lineColor: 'white'
        },
        errorbar: {
            color: 'white'
        }
    },
    legend: {
        backgroundColor: 'rgba(0, 0, 0, 0.5)',
        itemStyle: {
            color: '#E0E0E3'
        },
        itemHoverStyle: {
            color: '#FFF'
        },
        itemHiddenStyle: {
            color: '#606063'
        },
        title: {
            style: {
                color: '#C0C0C0'
            }
        }
    },
    credits: {
        style: {
            color: '#666'
        }
    },
    labels: {
        style: {
            color: '#707073'
        }
    },

    drilldown: {
        activeAxisLabelStyle: {
            color: '#F0F0F3'
        },
        activeDataLabelStyle: {
            color: '#F0F0F3'
        }
    },

    navigation: {
        buttonOptions: {
            symbolStroke: '#DDDDDD',
            theme: {
                fill: '#505053'
            }
        }
    },

    // scroll charts
    rangeSelector: {
        buttonTheme: {
            fill: '#505053',
            stroke: '#000000',
            style: {
                color: '#CCC'
            },
            states: {
                hover: {
                    fill: '#707073',
                    stroke: '#000000',
                    style: {
                        color: 'white'
                    }
                },
                select: {
                    fill: '#000003',
                    stroke: '#000000',
                    style: {
                        color: 'white'
                    }
                }
            }
        },
        inputBoxBorderColor: '#505053',
        inputStyle: {
            backgroundColor: '#333',
            color: 'silver'
        },
        labelStyle: {
            color: 'silver'
        }
    },

    navigator: {
        handles: {
            backgroundColor: '#666',
            borderColor: '#AAA'
        },
        outlineColor: '#CCC',
        maskFill: 'rgba(255,255,255,0.1)',
        series: {
            color: '#7798BF',
            lineColor: '#A6C7ED'
        },
        xAxis: {
            gridLineColor: '#505053'
        }
    },

    scrollbar: {
        barBackgroundColor: '#808083',
        barBorderColor: '#808083',
        buttonArrowColor: '#CCC',
        buttonBackgroundColor: '#606063',
        buttonBorderColor: '#606063',
        rifleColor: '#FFF',
        trackBackgroundColor: '#404043',
        trackBorderColor: '#404043'
    }
};

// Apply the theme
Highcharts.setOptions(Highcharts.theme);`
