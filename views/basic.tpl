<div class="panel panel-primary" id="div_candidate_basic">
  <div class="panel-heading">基本信息</div>
  <div class="panel-body">
    <div class="container-fluid">
      <div class="row">
        <div class="col-xs-3">
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon-name">名字:</span>
            <input type="text" id="input_candidate_name" name="input_candidate_name" class="form-control" placeholder="名字" aria-describedby="basic-addon-name">
          </div>
        </div>
        <div class="col-xs-3">
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon-birthday">生日:</span>
            <input type="text" id="input_candidate_birthday" name="input_candidate_birthday" class="form-control" placeholder="生日" aria-describedby="basic-addon-birthday">
          </div>
        </div>
        <div class="col-xs-3">
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon-gender">性别:</span>
            <select id="select_candidate_gender" class="form-control">
              <option value="男">男</option>
              <option value="女">女</option>
            </select>
          </div>
        </div>
        <div class="col-xs-3">
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon-workyear">工作年限:</span>
            <input type="number" id="input_candidate_workyear" name="input_candidate_workyear" class="form-control" placeholder="工作年限" aria-describedby="basic-addon-workyear">
          </div>
        </div>
      </div>
    </div>
    <div class="clearfix" style="margin-bottom: 10px;"></div>
    <div class="container-fluid">
      <div class="row">
        <div class="col-xs-3">
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon-email">邮箱:</span>
            <input type="email" id="input_candidate_email" name="input_candidate_email" class="form-control" placeholder="邮箱" aria-describedby="basic-addon-email">
          </div>
        </div>
        <div class="col-xs-3">
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon-mobile">电话:</span>
            <input type="text" id="input_candidate_mobile" name="input_candidate_mobile" class="form-control" placeholder="电话" aria-describedby="basic-addon-mobile">
          </div>
        </div>
        <div class="col-xs-3">
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon-post">职位:</span>
            <input type="text" id="input_candidate_post" name="input_candidate_post" class="form-control" placeholder="职位" aria-describedby="basic-addon-post">
          </div>
        </div>
        <div class="col-xs-3">
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon-city">城市:</span>
            <input type="text" id="input_candidate_city" name="input_candidate_city" class="form-control" placeholder="城市" aria-describedby="basic-addon-city">
          </div>
        </div>
      </div>
    </div>
    <div class="clearfix" style="margin-bottom: 10px;"></div>
    <div class="container-fluid">
      <div class="row">
        <div class="col-xs-3">
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon-company">公司:</span>
            <input type="text" id="input_candidate_company" name="input_candidate_company" class="form-control" placeholder="公司" aria-describedby="basic-addon-company">
          </div>
        </div>
        <div class="col-xs-3">
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon-education">学历:</span>
            <input type="text" id="input_candidate_education" name="input_candidate_education" class="form-control" placeholder="学历" aria-describedby="basic-addon-education">
          </div>
        </div>
        <div class="col-xs-3">
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon-age">年龄:</span>
            <input type="number" id="input_candidate_age" name="input_candidate_age" class="form-control" placeholder="年龄" aria-describedby="basic-addon-age">
          </div>
        </div>
        
      </div>
    </div>
    <div class="clearfix" style="margin-bottom: 10px;"></div>
    <div class="container-fluid">
      <div class="row">
        <div class="col-xs-6">
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon-label">标签:</span>
          <select id="select_candidate_label" name="select_label" class="form-control" multiple></select>
        </div>
      </div>
    </div>
  </div>
  
</div>
<div class="panel-footer">
  <div class="container-fluid">
    <button id="btn_candidate_save" name="btn_candidate_save" class="btn btn-default">保存</button>
  </div>
</div>
</div>
<script type="text/javascript" src="../static/js/jquery.validate.min.js"></script>
<script type="text/javascript" src="../static/js/localization/messages_zh.min.js"></script>
<script type="text/javascript" src="../static/js/jquery-validate.bootstrap-tooltip.min.js"></script>
<script type="text/javascript" src="../static/js/cust/func_basic.js"></script>
<script type="text/javascript" src="../static/js/cust/control_basic.js"></script>