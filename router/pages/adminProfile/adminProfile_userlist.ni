<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
      <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
      <title>NI - XSS 用户管理</title>
      
      <!-- Favicon -->
      <link href="/static/index/assets/img/apple-touch-icon.png" rel="icon">
      
      <link rel="stylesheet" href="/static/adminProfile/assets/css/backend.min.css?v=1.0.0">
      <link rel="stylesheet" href="/static/adminProfile/assets/vendor/@fortawesome/fontawesome-free/css/all.min.css">
      <link rel="stylesheet" href="/static/adminProfile/assets/vendor/line-awesome/dist/line-awesome/css/line-awesome.min.css">
      <link rel="stylesheet" href="/static/adminProfile/assets/vendor/remixicon/fonts/remixicon.css">
      <link rel="stylesheet" href="/static/adminProfile/assets/vendor/@icon/dripicons/dripicons.css">
      
      <link rel='stylesheet' href='/static/adminProfile/assets/vendor/fullcalendar/core/main.css' />
      <link rel='stylesheet' href='/static/adminProfile/assets/vendor/fullcalendar/daygrid/main.css' />
      <link rel='stylesheet' href='/static/adminProfile/assets/vendor/fullcalendar/timegrid/main.css' />
      <link rel='stylesheet' href='/static/adminProfile/assets/vendor/fullcalendar/list/main.css' />
      <link rel="stylesheet" href="/static/adminProfile/assets/vendor/mapbox/mapbox-gl.css">  </head>
  <body class="noteplus-layout  ">
    <!-- loader Start -->
    <div id="loading">
          <div id="loading-center">
			<input type="hidden" id="authID" value="{{ .authID }}"/>
			<input type="hidden" id="userID" value="{{ .userID }}"/>
          </div>
    </div>
	<!-- loader END -->
    <!-- Wrapper Start -->
    <div class="wrapper">
       <div class="iq-top-navbar">
          <div class="iq-navbar-custom">
              <nav class="navbar navbar-expand-lg navbar-light p-0">
                  <div class="iq-navbar-logo d-flex align-items-center justify-content-between">
                     <a href="#"><img src="/static/adminProfile/assets/images/logo.png" class="ri-menu-line wrapper-menu" alt="logo"></a>
                  </div>
              </nav>
          </div>
      </div>       
       <div class="iq-sidebar  sidebar-default ">
          <div class="iq-sidebar-logo d-flex align-items-center justify-content-between">
              <a href="index.html" class="header-logo">
                  <img src="/static/adminProfile/assets/images/logo.png" class="img-fluid rounded-normal light-logo" alt="logo"> <h4 class="logo-title ml-3">NI XSS</h4>
              </a>
              <div class="iq-menu-bt-sidebar">
                  <i class="las la-times wrapper-menu"></i>
              </div>
          </div>    
		  
         <div class="data-scrollbar" data-scroll="1">
              <nav class="iq-sidebar-menu">
                  <ul id="iq-sidebar-toggle" class="iq-menu">
					<li class="">
						<a class="svg-icon" href="/adminProfile/{{ .authID }}/profile">
							<i class="">
                                    <svg class="svg-icon" id="iq-ui-1-0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                         <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.121 17.804A13.937 13.937 0 0112 16c2.5 0 4.847.655 6.879 1.804M15 10a3 3 0 11-6 0 3 3 0 016 0zm6 2a9 9 0 11-18 0 9 9 0 0118 0z" style="stroke-dasharray: 90, 110; stroke-dashoffset: 0;"></path>
                                    </svg>
							</i>
							<span>个人管理</span>
						</a>
					</li>
					<li class="">
						<a class="svg-icon" href="/adminProfile/{{ .authID }}/userList">
							<i class="">
                                     <svg class="svg-icon feather feather-hard-drive" id="iq-ui-1-19" xmlns="http://www.w3.org/2000/svg" width="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22,12L2,12" style="stroke-dasharray: 20, 40; stroke-dashoffset: 0;"></path><path d="M5.45 5.11L2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11z" style="stroke-dasharray: 64, 84; stroke-dashoffset: 0;"></path><path d="M6,16L6.01,16" style="stroke-dasharray: 1, 21; stroke-dashoffset: 0;"></path><path d="M10,16L10.01,16" style="stroke-dasharray: 1, 21; stroke-dashoffset: 0;"></path>
                                   </svg>
                            </i>
							<span>用户管理</span>
						</a>
					</li>
					<li class="">
						<a class="svg-icon" href="/adminProfile/{{ .authID }}/logout">
							<i class="">
                                  <svg class="svg-icon" id="iq-auth-1-1" width="20" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                       <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" style="stroke-dasharray: 65, 85; stroke-dashoffset: 0;"></path>
                                  </svg>
                            </i>
							<span>登出</span>
						</a>
					</li>
                  </ul>
              </nav>
              <div class="p-2"></div>
          </div>
		  
      </div>         
	  <div class="content-page">
     <div class="container-fluid">        
        <div class="desktop-header">
            <div class="card card-block topnav-left">
                  <div class="card-body d-flex align-items-center">
                      <div class="d-flex justify-content-between">
                          <h4 class="text-capitalize">用户管理</h4>
                      </div>
                  </div>
              </div>
          </div>
        </div>
		<div class="container-fluid">   
        <div class="row">
            <div class="col-lg-12">
				<div class="card card-block card-stretch card-height">
                  <div class="card-body text-center">
					<div class="d-flex flex-wrap align-items-center justify-content-between">
                           <h4>新建用户</h4>
                           <div class="media flex-wrap align-items-center">
                              <a href="page-project-plans.html#" class="btn btn-info add-btn" data-toggle="modal" data-target="#new-user"><i class="las la-plus pr-2"></i>新建用户</a>
                           </div>
                     </div>
                  </div>
               </div>
			</div>
			<div class="col-sm-12">
                  <div class="card">
                     <div class="card-header d-flex justify-content-between">
                        <div class="header-title">
                           <h4 class="card-title">用户列表</h4>
                        </div>
                     </div>
                     <div class="card-body">
                        <div class="table-responsive">
                           <div class="row justify-content-between">
                              <div class="col-sm-6 col-md-6">
                                 <div id="user_list_datatable_info" class="dataTables_filter">
                                    <form class="mr-3 position-relative">
                                       <div class="form-group mb-0">
                                          <input type="search" class="form-control" id="findInputSearch" placeholder="用户ID/用户名称/用户邮箱" aria-controls="user-list-table">
                                       </div>
                                    </form>
                                 </div>
                              </div>
                              <div class="col-sm-6 col-md-6">
                                 <div class="user-list-files d-flex">
									<a class="bg-info" href="#" id="findBack">
                                       重置
                                    </a>
                                    <a class="bg-info" href="#" id="findGo">
                                       查找
                                    </a>
                                 </div>
                              </div>
                           </div>
                           <table id="user-list-table" class="table table-striped tbl-server-info mt-4" role="grid" aria-describedby="user-list-page-info">
                           <thead>
                              <tr class="ligth" style="text-align:center">
                                 <th>用户ID</th>
                                 <th>用户名称</th>
                                 <th>用户邮箱</th>
                                 <th>用户状态</th>
								 <th>用户等级</th>
                                 <th>登录IP</th>
								 <th>最后登录时间</th>
                                 <th>创建时间</th>
                                 <th style="min-width: 100px">操作</th>
                              </tr>
                           </thead>
                           <tbody id="userTableRows">
                                 
								{{ .userRow }}
							
                           </tbody>
                           </table>
                        </div>
                           <div class="row justify-content-between mt-3">
                              <div id="user-list-page-info" class="col-md-6">
                                 第 <span id="currentPage">1</span> 页 / 共 <span id="totalPage">{{ .pagesTotal }}</span> 页
                              </div>
                              <div class="col-md-6">
                                 <nav aria-label="Page navigation example">
                                    <ul class="pagination justify-content-end mb-0">
                                       <li class="page-item">
                                          <a class="page-link" href="#" id="previous">上一页</a>
                                       </li>
                                       <li class="page-item">
                                          <a class="page-link" href="#" id="next">下一页</a>
                                       </li>
                                    </ul>
                                 </nav>
                              </div>
                           </div>
                     </div>
                  </div>
            </div>
        </div>
	</div>
        <!-- Page end  -->
    </div>
    <!-- Modal -->
    <div class="modal fade" id="new-user" tabindex="-1" role="dialog" aria-hidden="true">
		<form action="#">
        <div class="modal-dialog modal-dialog-centered" style="width:1000px">
            <div class="modal-content">
				<div class="modal-header">
					<h4 class="modal-title" >新建用户</h4>
					<div class="btn-cancel p-0" data-dismiss="modal"><i class="las la-times"></i></div>
				</div>
                <div class="modal-body">
                    <div class="popup text-left">
                        <div class="content create-workform">
                    <div class="card-body write-card pb-8">
                        <div class="row">
                            <div class="col-md-12">
                           
                                    <div class="form-group">
                                        <label class="label-control">用户名称</label>
                                        <input type="text" class="form-control" id="newUserName" placeholder=" -- 请输入用户名称 -- " value="" data-change="input" >
                                    </div>
									<div class="form-group">
                                        <label class="label-control">用户邮箱</label>
                                        <input type="text" class="form-control" id="newUserEmail" placeholder=" -- 请输入用户邮箱 -- " value="" data-change="input" >
                                    </div>
									<div class="form-group">
                                        <label class="label-control">用户手机</label>
                                        <input type="text" class="form-control" id="newUserPhone" placeholder=" -- 请输入用户手机 -- " value="" data-change="input" >
                                    </div>
									<div class="form-group">
                                        <label class="label-control">用户密码</label>
                                        <input type="text" class="form-control" id="newUserPass" placeholder=" -- 请输入用户密码 -- " value="" data-change="input" >
                                    </div>
									<div class="form-group">
                                        <label class="label-control">用户状态</label>
                                        <select name="priority" id="newUserStatus" class="form-control" data-change="select" >
												<option value="-1"> -- 用户是否启用 -- </option>
                                                <option value="0">停用</option>
                                                <option value="1">启用</option>
                                            </select>
                                    </div>
									<div class="form-group">
                                        <label class="label-control">用户等级</label>
                                        <select name="priority" id="newUserLevel" class="form-control" data-change="select" >
												<option value="-1"> -- 用户等级 -- </option>
                                                <option value="0">低级用户</option>
                                                <option value="1">普通用户</option>
												<option value="2">高级用户</option>
                                            </select>
                                    </div>
                            </div>
                        </div>
                    </div>
                  </div>
                </div>
              </div>
			  <div class="modal-footer">
				<div class="col-sm-12 text-right">
					<button type="button" class="btn btn-outline-primary ml-1" data-dismiss="modal">关闭</button>
					<button type="reset" class="btn btn-outline-primary ml-1" >重置</button>
					<button type="button" class="btn btn-outline-primary ml-1" id="newUserOK" data-dismiss="modal">完成创建</button>
				</div>
			  </div>
            </div>
        </div>
	  </form>
    </div>
	    <div class="modal fade" id="edit-user" tabindex="-1" role="dialog" aria-hidden="true">
		<form action="#">
        <div class="modal-dialog modal-dialog-centered" style="width:1000px">
            <div class="modal-content">
				<div class="modal-header">
					<h4 class="modal-title" id="editUserTitle">编辑用户</h4>
					<div class="btn-cancel p-0" data-dismiss="modal"><i class="las la-times"></i></div>
				</div>
                <div class="modal-body">
                    <div class="popup text-left">
                        <div class="content create-workform">
                    <div class="card-body write-card pb-8">
                        <div class="row">
                            <div class="col-md-12">
									<input type="hidden" id="editUserID" value="">
                                    <div class="form-group">
                                        <label class="label-control">用户名称</label>
                                        <input type="text" class="form-control" id="editUserName" placeholder=" -- 请输入用户名称 -- " value="" data-change="input" >
                                    </div>
									<div class="form-group">
                                        <label class="label-control">用户邮箱</label>
                                        <input type="text" class="form-control" id="editUserEmail" placeholder=" -- 请输入用户邮箱 -- " value="" data-change="input" >
                                    </div>
									<div class="form-group">
                                        <label class="label-control">用户手机</label>
                                        <input type="text" class="form-control" id="editUserPhone" placeholder=" -- 请输入用户手机 -- " value="" data-change="input" >
                                    </div>
									<div class="form-group">
                                        <label class="label-control">用户密码</label>
                                        <input type="text" class="form-control" id="editUserPass" placeholder=" -- 请输入用户密码 -- " value="" data-change="input" >
                                    </div>
									<div class="form-group">
                                        <label class="label-control">用户状态</label>
                                        <select name="priority" id="editUserStatus" class="form-control" data-change="select" >
												<option value="-1"> -- 用户是否启用 -- </option>
                                                <option value="0" id="userStatus_0">停用</option>
                                                <option value="1" id="userStatus_1">启用</option>
                                            </select>
                                    </div>
									<div class="form-group">
                                        <label class="label-control">用户等级</label>
                                        <select name="priority" id="editUserLevel" class="form-control" data-change="select" >
												<option value="-1"> -- 用户等级 -- </option>
                                                <option value="0" id="userLevel_0">低级用户</option>
                                                <option value="1" id="userLevel_1">普通用户</option>
												<option value="2" id="userLevel_2">高级用户</option>
                                            </select>
                                    </div>
                            </div>
                        </div>
                    </div>
                  </div>
                </div>
              </div>  
			  <div class="modal-footer">
				<div class="col-sm-12 text-right">
					<button type="button" class="btn btn-outline-primary ml-1" data-dismiss="modal">关闭</button>
					<button type="button" class="btn btn-outline-primary ml-1" id="editUserSave" data-dismiss="modal">保存</button>
				</div>
			  </div>
            </div>
        </div>
	  </form>
    </div>
	
	<div class="modal fade" id="delete-user" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered" role="document" style="width:300px">
            <div class="modal-content">
                <div class="modal-body">
                    <div class="popup text-center">
                        <h4 class="mb-3">确认删除</h4>
                        <div class="content create-workform">
                            <p class="mb-2">你确定删除这个用户吗 ?</p>
							<input type="hidden" id="deleteUserID" value="" />
                            <div class="row">  
                                <div class="col-lg-12 mt-4">
                                    <div class="d-flex flex-wrap align-items-ceter justify-content-center">
                                        <div class="btn btn-outline-primary mr-4" data-dismiss="modal" >取消</div>
                                        <div class="btn btn-outline-primary" data-dismiss="modal" id="userDeleteConfirm" >确定</div>
                                    </div>
                                </div>                                           
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div> 
	
      </div>
    </div>
	
	
	<!-- Alert -->
	<div class="modal fade" id="ThisAlert" tabindex="-1" role="dialog" aria-labelledby="Alert" aria-hidden="true">
	<div class="modal-dialog">
		<div class="modal-content modal-dialog-centered">
			<div class="modal-header">
				<h4 class="modal-title" id="AlertTitle">
					
				</h4>
				<button type="button" class="close" data-dismiss="modal" aria-hidden="true">
					&times;
				</button>
			</div>
			<div class="modal-body" id="AlertBody">
				
			</div>
			<div class="modal-footer">
				<button type="button" class="btn btn-default" data-dismiss="modal">关闭
				</button>
			</div>
		</div>
	</div>
	</div>
	
    <!-- Wrapper End-->
    <footer class="iq-footer">
        <div class="container-fluid">
            <div class="row">
                <div class="col-lg-12 text-right">
                    <span class="text-secondary mr-1"><script>document.write(new Date().getFullYear())</script>©</span> <a href="user-profile.html#" class="">NI XSS</a>.
                </div>
            </div>
        </div>
    </footer>
    <!-- Backend Bundle JavaScript -->
    <script src="/static/adminProfile/assets/js/backend-bundle.min.js"></script>
    
    <!-- Flextree Javascript-->
    <script src="/static/adminProfile/assets/js/flex-tree.min.js"></script>
    <script src="/static/adminProfile/assets/js/tree.js"></script>
    
    <!-- Table Treeview JavaScript -->
    <script src="/static/adminProfile/assets/js/table-treeview.js"></script>
    
    <!-- Masonary Gallery Javascript -->
    <script src="/static/adminProfile/assets/js/masonry.pkgd.min.js"></script>
    <script src="/static/adminProfile/assets/js/imagesloaded.pkgd.min.js"></script>
    
    <!-- Mapbox Javascript -->
    <script src="/static/adminProfile/assets/js/mapbox-gl.js"></script>
    <script src="/static/adminProfile/assets/js/mapbox.js"></script>
    
    <!-- Fullcalender Javascript -->
    <script src='/static/adminProfile/assets/vendor/fullcalendar/core/main.js'></script>
    <script src='/static/adminProfile/assets/vendor/fullcalendar/daygrid/main.js'></script>
    <script src='/static/adminProfile/assets/vendor/fullcalendar/timegrid/main.js'></script>
    <script src='/static/adminProfile/assets/vendor/fullcalendar/list/main.js'></script>
    
    <!-- SweetAlert JavaScript -->
    <script src="/static/adminProfile/assets/js/sweetalert.js"></script>
    
    <!-- Vectoe Map JavaScript -->
    <script src="/static/adminProfile/assets/js/vector-map-custom.js"></script>
    
    <!-- Chart Custom JavaScript -->
    <script src="/static/adminProfile/assets/js/customizer.js"></script>
    
    <!-- Chart Custom JavaScript -->
    <script src="/static/adminProfile/assets/js/chart-custom.js"></script>
    
    <!-- slider JavaScript -->
    <script src="/static/adminProfile/assets/js/slider.js"></script>
    
    <!-- app JavaScript -->
    <script src="/static/adminProfile/assets/js/app.js"></script>
	<script src="/static/adminProfile/assets/js/alert.js"></script>
	<script src="/static/adminProfile/assets/js/userlist.js"></script>
	
  </body>
</html>