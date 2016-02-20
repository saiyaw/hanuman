<!DOCTYPE html>
<html  lang="en">
	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
		<title>萝卜</title>
		<link type="text/css" rel="stylesheet" href="../static/css/bootstrap.min.css" >
		<link type="text/css" rel="stylesheet" href="../static/css/dataTables.bootstrap.min.css">
		<link type="text/css" rel="stylesheet" href="../static/css/fileinput.min.css">
		<link type="text/css" rel="stylesheet" href="../static/css/select2.min.css">
		<link type="text/css" rel="stylesheet" href="../static/css/select2-bootstrap.min.css">
		<link type="text/css" rel="stylesheet" href="../static/css/bootstrap-datetimepicker.min.css">

		<script type="text/javascript" src="../static/js/jquery-2.1.4.min.js"></script>
		<script type="text/javascript" src="../static/js/moment/moment.min.js"></script>
		<script type="text/javascript" src="../static/js/moment/locale/zh-cn.js"></script>
		<script type="text/javascript" src="../static/js/plugins/canvas-to-blob.min.js"></script>
		<script type="text/javascript" src="../static/js/fileinput.min.js"></script>
		<script type="text/javascript" src="../static/js/fileinput_locale_zh.js"></script>
		<script type="text/javascript" src="../static/js/bootstrap.min.js"></script>
		<script type="text/javascript" src="../static/js/highcharts-custom.js"></script>
		<script type="text/javascript" src="../static/js/jquery.dataTables.min.js"></script>
		<script type="text/javascript" src="../static/js/dataTables.bootstrap.min.js"></script>
		<script type="text/javascript" src="../static/js/i18n/zh-CN.js"></script>
		<script type="text/javascript" src="../static/js/jquery.cookie.js"></script>
		<script type="text/javascript" src="../static/js/select2.full.min.js"></script>

		<script type="text/javascript" src="../static/js/bootstrap-datetimepicker.min.js"></script>

		<script type="text/javascript" src="../static/js/cust/func.js"></script>		
		
	</head>
	<body>
		<nav class="navbar navbar-default">
			<div class="container-fluid">
				<div class="navbar-header">
					<a class="navbar-brand" href="#">简历</a>
				</div>
				<div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
					<ul class="nav navbar-nav">
						<li><a href="/list">列表</a></li>
						<li><a href="/new">新建</a></li>
						<li><a href="/admin">管理</a></li>
					</ul>
				</div>
			</div>
		</nav>
		{{.LayoutContent}}
	</body>
</html>