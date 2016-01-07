<div class="panel panel-primary" id="div_admin_label">
  <div class="panel-heading">标签管理</div>
  <div class="panel-body">
    <div class="container-fluid">
      <div class="row">
        <div class="col-xs-9">
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon-label">标签:</span>
          <select id="select_admin_label" name="select_label" class="form-control"></select>
          <span class="input-group-btn">
            <button id="btn_admin_add_label" name="btn_add_label" class="btn btn-default">增加</button>
          </span>
          <span class="input-group-btn">
            <button class="btn btn-default" data-toggle="modal" data-target="#modal_delete">删除</button>
          </span>
          <span class="input-group-btn">
            <button id="btn_admin_clean_label" name="btn_clean_label" class="btn btn-default">清除</button>
          </span>
        </div>
      </div>
    </div>
  </div>
</div>
<div class="panel-footer">
  <div class="container-fluid">
    
    
  </div>
</div>
<!-- Modal -->
<div class="modal fade" id="modal_delete" tabindex="-1" role="dialog" aria-labelledby="modal_delete_label">
  <div class="modal-dialog" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        <h4 class="modal-title" id="modal_delete_label">确认删除?</h4>
      </div>
      <div class="modal-body">
        请再次确认是否要删除这个标签!
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
        <button type="button" id="btn_admin_delete_label" name="btn_delete_label" class="btn btn-primary">确认</button>
      </div>
    </div>
  </div>
</div>
<footer>
</footer>
<script type="text/javascript" src="../static/js/cust/func_admin.js"></script>
<script type="text/javascript" src="../static/js/cust/control_admin.js"></script>