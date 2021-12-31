<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
      <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
      <title>NI - XSS 任务|记录</title>
      
      <!-- Favicon -->
      <link href="/static/index/assets/img/apple-touch-icon.png" rel="icon">
      
      <link rel="stylesheet" href="/static/userProfile/assets/css/backend.min.css?v=1.0.0">
      <link rel="stylesheet" href="/static/userProfile/assets/vendor/@fortawesome/fontawesome-free/css/all.min.css">
      <link rel="stylesheet" href="/static/userProfile/assets/vendor/line-awesome/dist/line-awesome/css/line-awesome.min.css">
      <link rel="stylesheet" href="/static/userProfile/assets/vendor/remixicon/fonts/remixicon.css">
      <link rel="stylesheet" href="/static/userProfile/assets/vendor/@icon/dripicons/dripicons.css">
      
      <link rel='stylesheet' href='/static/userProfile/assets/vendor/fullcalendar/core/main.css' />
      <link rel='stylesheet' href='/static/userProfile/assets/vendor/fullcalendar/daygrid/main.css' />
      <link rel='stylesheet' href='/static/userProfile/assets/vendor/fullcalendar/timegrid/main.css' />
      <link rel='stylesheet' href='/static/userProfile/assets/vendor/fullcalendar/list/main.css' />
      <link rel="stylesheet" href="/static/userProfile/assets/vendor/mapbox/mapbox-gl.css">  </head>
  <body class="noteplus-layout  ">
    <!-- loader Start -->
    <div id="loading">
          <div id="loading-center">
			<input type="hidden" id="taskID" value="{{ .taskID }}"/>
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
                     <a href="#"><img src="/static/userProfile/assets/images/logo.png" class="ri-menu-line wrapper-menu" alt="logo"></a>
                  </div>
              </nav>
          </div>
      </div>       
       <div class="iq-sidebar  sidebar-default ">
          <div class="iq-sidebar-logo d-flex align-items-center justify-content-between">
              <a href="index.html" class="header-logo">
                  <img src="/static/userProfile/assets/images/logo.png" class="img-fluid rounded-normal light-logo" alt="logo"> <h4 class="logo-title ml-3">NI XSS</h4>
              </a>
              <div class="iq-menu-bt-sidebar">
                  <i class="las la-times wrapper-menu"></i>
              </div>
          </div>    
		  
         <div class="data-scrollbar" data-scroll="1">
              <nav class="iq-sidebar-menu">
                  <ul id="iq-sidebar-toggle" class="iq-menu">
					<li class="">
						<a class="svg-icon" href="/userProfile/{{ .authID }}/home">
							<i class="">
                                 <svg class="svg-icon" id="iq-ui-1-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
									<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 5v2m0 4v2m0 4v2M5 5a2 2 0 00-2 2v3a2 2 0 110 4v3a2 2 0 002 2h14a2 2 0 002-2v-3a2 2 0 110-4V7a2 2 0 00-2-2H5z" style="stroke-dasharray: 72, 92; stroke-dashoffset: 0;"></path>
                                  </svg>
                            </i>
							<span>主页</span>
						</a>
					</li>
					<li class="">
						<a class="svg-icon" href="/userProfile/{{ .authID }}/profile">
							<i class="">
							<svg class="svg-icon" id="iq-user-1" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" style="stroke-dasharray: 62, 82; stroke-dashoffset: 0;"></path>
                             </svg>
							 </i>
							<span>个人</span>
						</a>
					</li>
					<li class="">
						<a class="svg-icon" href="/userProfile/{{ .authID }}/tasks">
							<i class="">
                                    <svg class="svg-icon" id="iq-ui-1-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01" style="stroke-dasharray: 74, 94; stroke-dashoffset: 0;"></path>
                                    </svg>
                            </i>
							<span>任务</span>
						</a>
					</li>
					<li class="">
						<a class="svg-icon" href="/userProfile/{{ .authID }}/modules">
							<i class="">
                                 <svg class="svg-icon" id="iq-ui-1-7" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01" style="stroke-dasharray: 97, 117; stroke-dashoffset: 0;"></path>
                                 </svg>
                            </i>
							<span>模块</span>
						</a>
					</li>
					<li class="">
						<a class="svg-icon" href="/userProfile/{{ .authID }}/logout">
							<i class="">
                                <svg class="svg-icon feather feather-lock" id="iq-auth-1-5" xmlns="http://www.w3.org/2000/svg" width="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path width="18" height="11" d="M 5,11 L 19,11 A 2,2,0,0,1,21,13 L 21,20 A 2,2,0,0,1,19,22 L 5,22 A 2,2,0,0,1,3,20 L 3,13 A 2,2,0,0,1,5,11" style="stroke-dasharray: 55, 75; stroke-dashoffset: 0;"></path><path d="M7 11V7a5 5 0 0 1 10 0v4" style="stroke-dasharray: 24, 44; stroke-dashoffset: 0;"></path></svg>
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
                          <h4 class="text-capitalize" id="taskTitle">{{ .taskName }}</h4>
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
	
						<div class="form-group">
                            <label class="label-control">任务名称</label>
                            <input type="text" class="form-control" id="taskName" placeholder=" -- 请输入任务名称 -- " value="{{ .taskName }}" data-change="input" >
                        </div>
						
						<div class="form-group">
                            <label class="label-control">任务描述</label>
                            <textarea type="text" class="form-control" id="taskData" rows="3" data-change="input"  placeholder=" -- 请输入任务描述 -- ">{{ .taskData }}</textarea>
                        </div>
						
						<div class="col-sm-12 text-right">
							<button type="button" class="btn btn-outline-primary ml-1" id="taskSave" data-dismiss="modal">保存</button>
						</div>
	
                  </div>
               </div>
			</div>
			<div class="col-sm-12">
                  <div class="card">
                     <div class="card-header d-flex justify-content-between">
                        <div class="header-title">
                           <h4 class="card-title">记录列表</h4>
                        </div>
                     </div>
                     <div class="card-body">
                        <div class="table-responsive">
                           <div class="row justify-content-between">
                              <div class="col-sm-6 col-md-6">
                                 <div id="user_list_datatable_info" class="dataTables_filter">
                                    <form class="mr-3 position-relative">
                                       <div class="form-group mb-0">
                                          <input type="search" class="form-control" id="recordSearch" placeholder="记录ID" aria-controls="user-list-table">
                                       </div>
                                    </form>
                                 </div>
                              </div>
                              <div class="col-sm-6 col-md-6">
                                 <div class="user-list-files d-flex">
									<a class="bg-info" href="#" id="viewBack">
                                       返回
                                    </a>
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
                                 <th>记录ID</th>
                                 <th>请求方法</th>
                                 <th>请求IP</th>
                                 <th>请求时间</th>
                                 <th style="min-width: 100px">操作</th>
                              </tr>
                           </thead>
                           <tbody id="viewTableRows">
                                 
								{{ .taskRow }}
							
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
	
	
	<div class="modal fade" id="delete-record" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered" role="document" style="width:300px">
            <div class="modal-content">
                <div class="modal-body">
                    <div class="popup text-center">
                        <h4 class="mb-3">确认删除</h4>
                        <div class="content create-workform">
                            <p class="mb-2">你确定删除这个记录吗 ?</p>
							<input type="hidden" id="deleteRecordID" value="" />
                            <div class="row">  
                                <div class="col-lg-12 mt-4">
                                    <div class="d-flex flex-wrap align-items-ceter justify-content-center">
                                        <div class="btn btn-outline-primary mr-4" data-dismiss="modal" >取消</div>
                                        <div class="btn btn-outline-primary" data-dismiss="modal" id="recordDeleteConfirm" >确定</div>
                                    </div>
                                </div>                                           
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
	
	<div class="modal fade" id="default-record" tabindex="-1" role="dialog" aria-hidden="true" >
            <div class="modal-dialog modal-dialog-centered" role="document" style="width:1000px">
                <div class="modal-content">
					<div class="modal-header">
						<h4 class="modal-title" id="defaultRecordID" >记录ID</h4>
						<div class="btn-cancel p-0" data-dismiss="modal"><i class="las la-times"></i></div>
					</div>
                    <div class="modal-body">
                        <div class="popup text-left">
                            <div class="content create-workform">
							
								<div class="form-group">
									<label class="label-control">请求方法</label>
									<textarea type="text" class="form-control"  rows="1" data-change="input"  id="defaultMethod" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">IP</label>
									<textarea type="text" class="form-control"  rows="1" data-change="input"  id="defaultIP" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">请求时间</label>
									<textarea type="text" class="form-control"  rows="1" data-change="input"  id="defaultTime" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">代理</label>
									<textarea type="text" class="form-control"  rows="2" data-change="input"  id="defaultAgent" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">请求头(json)</label>
									<textarea type="text" class="form-control"  rows="5" data-change="input"  id="defaultHeaders" disabled></textarea>
								</div>
								
								<hr>
								
								<div id="defaultExtra">
									
								</div>
								
                            </div>
                        </div>
                    </div>
					<div class="modal-footer">
						<button type="button" class="btn btn-outline-primary ml-1 " data-dismiss="modal">关闭</button>
					</div>
                </div>
            </div>
        </div> 
		
			<div class="modal fade" id="custom-record" tabindex="-1" role="dialog" aria-hidden="true" >
            <div class="modal-dialog modal-dialog-centered" role="document" style="width:600px">
                <div class="modal-content">
					<div class="modal-header">
						<h4 class="modal-title" id="customRecordID" >记录ID</h4>
						<div class="btn-cancel p-0" data-dismiss="modal"><i class="las la-times"></i></div>
					</div>
                    <div class="modal-body">
                        <div class="popup text-left">
                            <div class="content create-workform" >
								
								<div class="form-group">
									<label class="label-control">请求方法</label>
									<textarea type="text" class="form-control"  rows="1" data-change="input"  id="customMethod" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">IP</label>
									<textarea type="text" class="form-control"  rows="1" data-change="input"  id="customIP" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">请求时间</label>
									<textarea type="text" class="form-control"  rows="1" data-change="input"  id="customTime" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">代理</label>
									<textarea type="text" class="form-control"  rows="2" data-change="input"  id="customAgent" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">请求头(json)</label>
									<textarea type="text" class="form-control"  rows="5" data-change="input"  id="customHeaders" disabled></textarea>
								</div>
								
                            </div>
                        </div>
                    </div>
					<div class="modal-footer">
						<button type="button" class="btn btn-outline-primary ml-1 " data-dismiss="modal">关闭</button>
					</div>
                </div>
            </div>
        </div> 
		
		<div class="modal fade" id="request-record" tabindex="-1" role="dialog" aria-hidden="true" >
            <div class="modal-dialog modal-dialog-centered" role="document" style="width:600px">
                <div class="modal-content">
					<div class="modal-header">
						<h4 class="modal-title" id="requestRecordID" >记录ID</h4>
						<div class="btn-cancel p-0" data-dismiss="modal"><i class="las la-times"></i></div>
					</div>
                    <div class="modal-body">
                        <div class="popup text-left">
                            <div class="content create-workform" >
								
								<div class="form-group">
									<label class="label-control">请求方法</label>
									<textarea type="text" class="form-control"  rows="1" data-change="input"  id="requestMethod" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">IP</label>
									<textarea type="text" class="form-control"  rows="1" data-change="input"  id="requestIP" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">请求时间</label>
									<textarea type="text" class="form-control"  rows="1" data-change="input"  id="requestTime" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">代理</label>
									<textarea type="text" class="form-control"  rows="2" data-change="input"  id="requestAgent" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">请求头(json)</label>
									<textarea type="text" class="form-control"  rows="5" data-change="input"  id="requestHeaders" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">请求头(json)</label>
									<textarea type="text" class="form-control"  rows="5" data-change="input"  id="customHeaders" disabled></textarea>
								</div>

                            </div>
                        </div>
                    </div>
					<div class="modal-footer">
						<button type="button" class="btn btn-outline-primary ml-1 " data-dismiss="modal">关闭</button>
					</div>
                </div>
            </div>
        </div> 
		
		<div class="modal fade" id="fish-record" tabindex="-1" role="dialog" aria-hidden="true" >
            <div class="modal-dialog modal-dialog-centered" role="document" style="width:600px">
                <div class="modal-content">
					<div class="modal-header">
						<h4 class="modal-title" id="fishRecordID" >记录ID</h4>
						<div class="btn-cancel p-0" data-dismiss="modal"><i class="las la-times"></i></div>
					</div>
                    <div class="modal-body">
                        <div class="popup text-left">
                            <div class="content create-workform" >
								
								<div class="form-group">
									<label class="label-control">请求方法</label>
									<textarea type="text" class="form-control"  rows="1" data-change="input"  id="fishMethod" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">IP</label>
									<textarea type="text" class="form-control"  rows="1" data-change="input"  id="fishIP" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">请求时间</label>
									<textarea type="text" class="form-control"  rows="1" data-change="input"  id="fishTime" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">代理</label>
									<textarea type="text" class="form-control"  rows="2" data-change="input"  id="fishAgent" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">请求头(json)</label>
									<textarea type="text" class="form-control"  rows="5" data-change="input"  id="fishHeaders" disabled></textarea>
								</div>
								
                            </div>
                        </div>
                    </div>
					<div class="modal-footer">
						<button type="button" class="btn btn-outline-primary ml-1 " data-dismiss="modal">关闭</button>
					</div>
                </div>
            </div>
        </div> 
		
		<div class="modal fade" id="port-record" tabindex="-1" role="dialog" aria-hidden="true" >
            <div class="modal-dialog modal-dialog-centered" role="document" style="width:600px">
                <div class="modal-content">
					<div class="modal-header">
						<h4 class="modal-title" id="portRecordID" >记录ID</h4>
						<div class="btn-cancel p-0" data-dismiss="modal"><i class="las la-times"></i></div>
					</div>
                    <div class="modal-body">
                        <div class="popup text-left">
                            <div class="content create-workform" >
								
								<div class="form-group">
									<label class="label-control">请求方法</label>
									<textarea type="text" class="form-control"  rows="1" data-change="input"  id="portMethod" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">IP</label>
									<textarea type="text" class="form-control"  rows="1" data-change="input"  id="portIP" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">请求时间</label>
									<textarea type="text" class="form-control"  rows="1" data-change="input"  id="portTime" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">代理</label>
									<textarea type="text" class="form-control"  rows="2" data-change="input"  id="portAgent" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">请求头(json)</label>
									<textarea type="text" class="form-control"  rows="5" data-change="input"  id="portHeaders" disabled></textarea>
								</div>
								<hr>
								<div class="form-group">
									<label class="label-control">开放端口</label>
									<textarea type="text" class="form-control"  rows="3" data-change="input"  id="portOpen" disabled></textarea>
								</div>
								<div class="form-group">
									<label class="label-control">关闭端口</label>
									<textarea type="text" class="form-control"  rows="3" data-change="input"  id="portClosed" disabled></textarea>
								</div>
								
                            </div>
                        </div>
                    </div>
					<div class="modal-footer">
						<button type="button" class="btn btn-outline-primary ml-1 " data-dismiss="modal">关闭</button>
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
    <script src="/static/userProfile/assets/js/backend-bundle.min.js"></script>
    
    <!-- Flextree Javascript-->
    <script src="/static/userProfile/assets/js/flex-tree.min.js"></script>
    <script src="/static/userProfile/assets/js/tree.js"></script>
    
    <!-- Table Treeview JavaScript -->
    <script src="/static/userProfile/assets/js/table-treeview.js"></script>
    
    <!-- Masonary Gallery Javascript -->
    <script src="/static/userProfile/assets/js/masonry.pkgd.min.js"></script>
    <script src="/static/userProfile/assets/js/imagesloaded.pkgd.min.js"></script>
    
    <!-- Mapbox Javascript -->
    <script src="/static/userProfile/assets/js/mapbox-gl.js"></script>
    <script src="/static/userProfile/assets/js/mapbox.js"></script>
    
    <!-- Fullcalender Javascript -->
    <script src='/static/userProfile/assets/vendor/fullcalendar/core/main.js'></script>
    <script src='/static/userProfile/assets/vendor/fullcalendar/daygrid/main.js'></script>
    <script src='/static/userProfile/assets/vendor/fullcalendar/timegrid/main.js'></script>
    <script src='/static/userProfile/assets/vendor/fullcalendar/list/main.js'></script>
    
    <!-- SweetAlert JavaScript -->
    <script src="/static/userProfile/assets/js/sweetalert.js"></script>
    
    <!-- Vectoe Map JavaScript -->
    <script src="/static/userProfile/assets/js/vector-map-custom.js"></script>
    
    <!-- Chart Custom JavaScript -->
    <script src="/static/userProfile/assets/js/customizer.js"></script>
    
    <!-- Chart Custom JavaScript -->
    <script src="/static/userProfile/assets/js/chart-custom.js"></script>
    
    <!-- slider JavaScript -->
    <script src="/static/userProfile/assets/js/slider.js"></script>
    
    <!-- app JavaScript -->
    <script src="/static/userProfile/assets/js/app.js"></script>
	<script src="/static/userProfile/assets/js/alert.js"></script>
	<script src="/static/userProfile/assets/js/view.js"></script>
	
  </body>
</html>