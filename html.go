package main

const letter = `
<!DOCTYPE html>
<html>
<head>
<title>Temporary text sharing - Temporary Web note - 临时文本 - 临时笔记</title>
<meta name="viewport" content="width=device-width, initial-scale=1" />
<meta name="description" content="Temporary text web sharing Tool, Temporary note web sharing  - 临时文本分享工具, 临时笔记">
</head>
<body>
<textarea style="width:100%;height:800px;font-size: 20px;" placeholder="输入 Input" name="data">{{.Data}}</textarea>
<div style="
    position: fixed;
    bottom: 10%;
    display: inline;
    left: 50%;
"><button id="share" style="font-size: 20px;">Share Page</button></div>
<button id="copy" style="display:none" data-clipboard-text=""></button>
<script src="https://cdnjs.cloudflare.com/ajax/libs/clipboard.js/2.0.4/clipboard.min.js"></script>
<script>
window.onload=function(){

	clipboard = new ClipboardJS('#copy');
	clipboard.on('success', function(e) {
		console.info('Action:', e.action);
		console.info('Text:', e.text);
		console.info('Trigger:', e.trigger);

	alert('Copied Share Url');
	
		e.clearSelection();
	});
	
	clipboard.on('error', function(e) {
		console.error('Action:', e.action);
		console.error('Trigger:', e.trigger);
	});

let s=document.querySelector('#share');
let c=document.querySelector('textarea');
let key='';

c.onkeypress=function(e){
}

s.onclick=function(e){
	s.setAttribute('disabled','');

	let content=c.value;
	var xmlhttp;
if (window.XMLHttpRequest)
  {// code for IE7+, Firefox, Chrome, Opera, Safari
  xmlhttp=new XMLHttpRequest();
  }
else
  {// code for IE6, IE5
  xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
  }
xmlhttp.onreadystatechange=function()
  {
  if (xmlhttp.readyState==4 && xmlhttp.status==200)
    {
		key=xmlhttp.responseText;
	s.removeAttribute('disabled');
	document.querySelector('#copy').setAttribute('data-clipboard-text',location.protocol+'//'+location.host+'/'+key);
	document.querySelector('#copy').click();
    }
  }
xmlhttp.open("POST","/save",true);
xmlhttp.send(content);
}

}
</script>
</body>
</html>
`
