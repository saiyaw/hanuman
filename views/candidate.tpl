


{{template "basic.tpl"}}

<div class="panel panel-info" id="div_candidate_attachment">
<div class="panel-heading">附件</div>
<div class="panel-body">
	<ul class="list-group" id="ul_candidate_attachment_list">
	</ul>
</div>
<div class="panel-footer">
	<div class="container-fluid">
		<input type="file" id="input_candidate_attachment" name="attachment_upload" multiple class="file-loading">
	</div>
</div>
</div>
<div class="panel panel-primary" id="div_candidate_comment">
<div class="panel-heading">备注</div>
<div class="panel-body">
	<div class="container-fluid">
		<div class="form-group">
			<textarea class="form-control" rows="3" id="textarea_candidate_comment" name="textarea_comment"></textarea>
		</div>
	</div>
</div>
<div class="panel-footer">
	<div class="container-fluid">
		<button id="btn_candidate_save_comment" name="btn_comment_save" class="btn btn-default">保存</button>
	</div>
</div>
</div>
<div class="panel panel-default" id="div_candidate_history">
<div class="panel-heading">记录</div>
<div class="panel-body">
	<ul class="list-group" id="ul_candidate_comment_list">
	</ul>
</div>
</div>
<footer>
</footer>

<script type="text/javascript" src="../static/js/cust/func_candidate.js"></script>
<script type="text/javascript" src="../static/js/cust/control_candidate.js"></script>