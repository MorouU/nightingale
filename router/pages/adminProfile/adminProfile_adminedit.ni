<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
      <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
      <title>NI - XSS 个人管理</title>
      
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
  <body class=" ">
    <!-- loader Start -->
    <div id="loading">
          <div id="loading-center">
			<input type="hidden" id="authID" value="{{ .authID }}"/>
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
              <a href="/preview/1097306/2021-01-28/note/backend/index.html" class="header-logo">
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
                          <h4 class="text-capitalize">个人</h4>
                      </div>
                  </div>
              </div>
          </div>        
		</div>
      <div class="container-fluid">
         <div class="row">
            <div class="col-lg-12">
               <div class="iq-edit-list-data">
                  <div class="tab-content">
                     <div class="tab-pane fade active show" id="personal-information" role="tabpanel">
                        <div class="card">
                           <div class="card-header d-flex justify-content-between">
                              <div class="iq-header-title">
                                 <h4 class="card-title">个人信息</h4>
                              </div>
                           </div>
                           <div class="card-body">
                              <form>
                                 <div class=" row align-items-center">
                                    <div class="form-group col-sm-6">
                                       <label for="editAdminName">用户名:</label>
                                       <input type="text" class="form-control" id="editAdminName" value="{{ .adminName }}">
                                    </div>
                                    <div class="form-group col-sm-6">
                                       <label for="editAdminEmail">邮箱:</label>
                                       <input type="text" class="form-control" id="editAdminEmail" value="{{ .adminEmail }}">
                                    </div>
                                    <div class="form-group col-sm-6">
                                       <label for="editAdminPhone">手机号:</label>
                                       <input type="text" class="form-control" id="editAdminPhone" value="{{ .adminPhone }}">
                                    </div>
                                 </div>
								 <div class="form-group col-sm-12 text-right">
									<button type="reset" class="btn btn-primary mr-2" >重置</button>
									<button type="button" class="btn btn-info mr-2 " id="editSubmit">提交</button>
								 </div>
                              </form>
                           </div>
						   </div>
						   <div class="card">
						   <div class="card-header d-flex justify-content-between">
                              <div class="iq-header-title">
                                 <h4 class="card-title">修改密码</h4>
                              </div>
                           </div>
						   <div class="card-body">
                              <form>
                                 <div class="form-group">
                                    <label for="changeAdminCPass">当前密码:</label>
                                    <input type="Password" class="form-control" id="changeAdminCPass" value="">
                                 </div>
                                 <div class="form-group">
                                    <label for="changeAdminNPass">新密码:</label>
                                    <input type="Password" class="form-control" id="changeAdminNPass" value="">
                                 </div>
                                 <div class="form-group">
                                    <label for="changeAdminVPass">确认新密码:</label>
                                    <input type="Password" class="form-control" id="changeAdminVPass" value="">
                                 </div>
								 <div class="form-group col-sm-12 text-right">
									<button type="reset" class="btn btn-primary mr-2">重置</button>
									<button type="button" class="btn btn-info mr-2" id="changeSubmit">提交</button>
								</div>
                              </form>
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
	<script src="/static/adminProfile/assets/js/profile.js"></script>
	
  </body>
</html>