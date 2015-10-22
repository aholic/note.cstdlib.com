<!DOCTYPE html>
<html>
<head>
<meta content='text/html; charset=utf-8' http-equiv='Content-Type'>
<meta content='Easy way to share note' name='description'>
<title>Share Your Note</title>
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
<h1 style='margin:0px;'>Share Your Note</h1>
<p>
how to use: [write down your note] -> [generate url] -> [share it]
</p>
<textarea id="noteContent" rows="35" style="background-color:#EEEEEE;width:100%;resize:none"></textarea>
<!--
<p>The server is written in golang, explore <a href="https://github.com/aholic/gojsonfmt/blob/master/format.go">more</a>.</p>
-->
<input type="button" name="generateUrl" id="generateUrl" value="Generate Url"/>
<input type="text" id="urlToShare" name="urlToShare" value="" style="width:300px" readonly>
<div style="height:150px;"> </div>
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
