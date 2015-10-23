<!DOCTYPE html>
<html>
<head>
<meta content='text/html; charset=utf-8' http-equiv='Content-Type'>
<meta content='便笺快速分享' name='description'>
<title>随享记</title>
</head>

<style type="text/css">
.align-center{
    margin:0 auto;      
    width:800px;        
    text-align:center;  
}
</style>
<body style='font-family:Helvetica'>
<div class="align-center">
<div style="height:50px;"> </div>
<h1 style='margin:0px;'>随享记</h1>
<div style="height:90px;"> </div>
<textarea id="noteContent" rows="35" style="background-color:#EEEEEE;width:100%;resize:none"></textarea>
<p>
使用步骤：输入便笺内容、生成url、将url发给你要分享的人
</p>
<input type="button" name="generateUrl" id="generateUrl" value="生成url"/>
<input type="text" id="urlToShare" name="urlToShare" value="" style="width:300px" readonly>
<div style="height:120px;"> </div>
<a style='width:800px;' href="http://www.miitbeian.gov.cn/">
 京ICP备14040458号
</a>
</div>

<script src="/static/js/jquery-2.1.4.min.js" type="text/javascript"></script>
<script type="text/javascript">
    $(document).ready(function() {
        $("#urlToShare").click(function() {
            $(this).select();
        });
        $("#generateUrl").click(function() {
            if ($("#noteContent").val().trim() == "") return;
            $.post("/note/submit", {"noteContent" : $("#noteContent").val()}, function(rst) {
                if (!rst.flag) {
                    alert(rst.msg);
                    return;
                } 
                $("#urlToShare").val("http://note.cstdlib.com/note/show/" + rst.data)
            },"json");
        });
    });
</script>
</body>
</html>
