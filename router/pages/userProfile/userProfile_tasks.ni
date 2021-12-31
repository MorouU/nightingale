<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
      <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
      <title>NI - XSS 任务</title>
      
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
                          <h4 class="text-capitalize">任务</h4>
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
                           <h4>创建任务</h4>
                           <div class="media flex-wrap align-items-center">
                              <a href="page-project-plans.html#" class="btn btn-danger add-btn" data-toggle="modal" data-target="#new-task"><i class="las la-plus pr-2"></i>新建任务</a>
                           </div>
                     </div>
                  </div>
               </div>
			</div>
			<div class="col-sm-12">
                  <div class="card">
                     <div class="card-header d-flex justify-content-between">
                        <div class="header-title">
                           <h4 class="card-title">任务列表</h4>
                        </div>
                     </div>
                     <div class="card-body">
                        <div class="table-responsive">
                           <div class="row justify-content-between">
                              <div class="col-sm-6 col-md-6">
                                 <div id="user_list_datatable_info" class="dataTables_filter">
                                    <form class="mr-3 position-relative">
                                       <div class="form-group mb-0">
                                          <input type="search" class="form-control" id="findInputSearch" placeholder="任务ID" aria-controls="user-list-table">
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
                                 <th>任务ID</th>
                                 <th>任务名称</th>
                                 <th>模块名</th>
                                 <th>记录数</th>
                                 <th>状态</th>
								 <th>参数</th>
                                 <th>创建时间</th>
                                 <th style="min-width: 100px">操作</th>
                              </tr>
                           </thead>
                           <tbody id="taskTableRows">
                                 
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
    <div class="modal fade" id="new-task" tabindex="-1" role="dialog" aria-hidden="true">
		<form action="#">
        <div class="modal-dialog modal-dialog-centered" style="width:1000px">
            <div class="modal-content">
				<div class="modal-header">
					<h4 class="modal-title" >新建任务</h4>
					<div class="btn-cancel p-0" data-dismiss="modal"><i class="las la-times"></i></div>
				</div>
                <div class="modal-body">
                    <div class="popup text-left">
                        <div class="content create-workform">
                    <div class="card-body write-card pb-8">
                        <div class="row">
                            <div class="col-md-12">
                                    <div class="form-group">
                                        <label class="label-control">任务名称</label>
                                        <input type="text" class="form-control" id="newTaskName" placeholder=" -- 请输入任务名称 -- " value="" data-change="input" >
                                    </div>
                                    <div class="form-group">
                                        <label class="label-control">任务描述</label>
                                        <textarea type="text" class="form-control" id="newTaskData" rows="3" data-change="input"  placeholder=" -- 请输入任务描述 -- "></textarea>
                                    </div>
                                    <div class="form-group">
                                        <label class="label-control">任务状态</label>
                                        <div>
                                            <select name="priority" id="newTaskStatus" class="form-control" data-change="select" >
												<option value="-1"> -- 任务是否启用 -- </option>
                                                <option value="0">停用</option>
                                                <option value="1">启用</option>
                                            </select>
                                        </div>
                                    </div>
									<div class="form-group">
                                        <label class="label-control">任务模块</label>
                                        <div>
                                            <select name="priority" id="newTaskModule" class="form-control" data-change="select" >
												<option selected value="-1"> -- 请选择任务模块 -- </option>
                                                <optgroup label="个人模块">
													{{ .personModule }}
												</optgroup>
												<optgroup label="公共模块">
													{{ .publicModule }}
												</optgroup>
                                            </select>
                                        </div>
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
					<button type="button" class="btn btn-outline-primary ml-1" id="newTaskNext" data-dismiss="modal">下一步</button>
				</div>
			  </div>
            </div>
        </div>
	  </form>
    </div>
	
	<div class="modal fade" id="control-task" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered" role="document" style="width:300px">
            <div class="modal-content">
                <div class="modal-body">
                    <div class="popup text-center">
                        <h4 class="mb-3">改变状态</h4>
                        <div class="content create-workform">
                            <p class="mb-2" id="statusTask">你确定启用这个任务吗 ?</p>
							<input type="hidden" id="controlTaskID" value="" />
                            <div class="row">  
                                <div class="col-lg-12 mt-4">
                                    <div class="d-flex flex-wrap align-items-ceter justify-content-center">
                                        <div class="btn btn-outline-primary mr-4" data-dismiss="modal" >取消</div>
                                        <div class="btn btn-outline-primary" data-dismiss="modal" id="taskControlConfirm" >确定</div>
                                    </div>
                                </div>                                           
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
	<div class="modal fade" id="delete-task" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered" role="document" style="width:300px">
            <div class="modal-content">
                <div class="modal-body">
                    <div class="popup text-center">
                        <h4 class="mb-3">确认删除</h4>
                        <div class="content create-workform">
                            <p class="mb-2">你确定删除这个任务吗 ?</p>
							<input type="hidden" id="deleteTaskID" value="" />
                            <div class="row">  
                                <div class="col-lg-12 mt-4">
                                    <div class="d-flex flex-wrap align-items-ceter justify-content-center">
                                        <div class="btn btn-outline-primary mr-4" data-dismiss="modal" >取消</div>
                                        <div class="btn btn-outline-primary" data-dismiss="modal" id="taskDeleteConfirm" >确定</div>
                                    </div>
                                </div>                                           
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
	<div class="modal fade" id="getApi-task" tabindex="-1" role="dialog" aria-hidden="true" >
            <div class="modal-dialog modal-dialog-centered" role="document" style="width:600px">
                <div class="modal-content">
					<div class="modal-header">
						<h4 class="modal-title" id="taskApiName" >任务名称</h4>
						<div class="btn-cancel p-0" data-dismiss="modal"><i class="las la-times"></i></div>
					</div>
                    <div class="modal-body">
                        <div class="popup text-left">
                            <div class="content create-workform">
								<textarea type="text" class="form-control" style="background-color:#f3f6fd" rows="6" disabled id="taskApiData">内容</textarea>
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
    </div>
	
    <div class="modal fade" id="new-custom-task" tabindex="-1" role="dialog" aria-hidden="true">
		<form action="#">
        <div class="modal-dialog modal-dialog-centered" style="width:1000px">
            <div class="modal-content">
				<div class="modal-header">
					<h4 class="modal-title" id="customTaskName">任务名称</h4>
					<div class="btn-cancel p-0" data-dismiss="modal"><i class="las la-times"></i></div>
				</div>
                <div class="modal-body">
                    <div class="popup text-left">
                        <div class="content create-workform">
                    <div class="card-body write-card pb-8">
                        <div class="row">
                            <div class="col-md-12">
                                
                                    <div class="form-group">
										<input type="hidden" id="customModuleID" value="" />
                                        <label class="label-control">任务代码</label>
                                        <textarea type="text" class="form-control"  rows="5" data-change="input"  placeholder=" -- 请输入任务代码 -- " id="customCode"></textarea>
                                    </div>
									
                            </div>
                        </div>
                    </div>
                  </div>
                </div>
              </div>
			  
			  <div class="modal-footer">
				<div class="col-sm-12 text-right">
					<button type="button" class="btn btn-outline-primary ml-1" data-dismiss="modal" id="customBack">上一步</button>
					<button type="reset" class="btn btn-outline-primary ml-1" >重置</button>
					<button type="button" class="btn btn-outline-primary ml-1" id="customOK">完成</button>
				</div>
			  </div>
            </div>
        </div>
	  </form>
    </div>
	
	<div class="modal fade" id="default-task" tabindex="-1" role="dialog" aria-hidden="true" >
            <div class="modal-dialog modal-dialog-centered" role="document" style="width:1000px">
                <div class="modal-content">
					<div class="modal-header">
						<h4 class="modal-title" id="defaultTaskName" >任务名称</h4>
						<div class="btn-cancel p-0" data-dismiss="modal"><i class="las la-times"></i></div>
					</div>
                    <div class="modal-body">
                        <div class="popup text-left">
                            <div class="content create-workform">
								<input type="hidden" id="defaultModuleID" value="" />
								<div class="data-privacy">
                                    <h4 class="mb-2"> 选项 </h4>
                                    <div class="custom-control custom-checkbox">
                                        <input type="checkbox" id="defaultCurrentUrl" class="custom-control-input"
                                            checked>
                                        <label class="custom-control-label" for="defaultCurrentUrl"> 获取对方当前URL </label>
                                    </div>
									<div class="custom-control custom-checkbox">
                                        <input type="checkbox" id="defaultRefererUrl" class="custom-control-input"
                                            checked>
                                        <label class="custom-control-label" for="defaultRefererUrl"> 获取对方来路URL </label>
                                    </div>
                                    <div class="custom-control custom-checkbox">
                                        <input type="checkbox" id="defaultCookie" class="custom-control-input"
											checked>
                                        <label class="custom-control-label" for="defaultCookie"> 获取对方Cookie数据 </label>
                                    </div>
                                    <div class="custom-control custom-checkbox">
                                        <input type="checkbox" id="defaultOS" class="custom-control-input">
                                        <label class="custom-control-label" for="defaultOS"> 获取对方操作系统 </label>
                                    </div>
									<div class="custom-control custom-checkbox">
                                        <input type="checkbox" id="defaultBrowser" class="custom-control-input"
											checked>
                                        <label class="custom-control-label" for="defaultBrowser"> 获取对方浏览器信息 </label>
                                    </div>
									<div class="custom-control custom-checkbox">
                                        <input type="checkbox" id="defaultScreenResolution" class="custom-control-input">
                                        <label class="custom-control-label" for="defaultScreenResolution"> 获取对方屏幕分辨率 </label>
                                    </div>
									<div class="custom-control custom-checkbox">
                                        <input type="checkbox" id="defaultWebPage" class="custom-control-input">
                                        <label class="custom-control-label" for="defaultWebPage"> 获取对方网页内容 <code>(会以json格式发送至邮箱)</code> </label>
                                    </div>
									<div class="custom-control custom-checkbox">
                                        <input type="checkbox" id="defaultScreenShot" class="custom-control-input">
                                        <label class="custom-control-label" for="defaultScreenShot"> 获取对方网页截图 <code>(会以json格式发送至邮箱)</code> </label>
                                    </div>
									<br>
                                    <p>新建任务完成后请将 API 嵌入可能存在XSS的点即可.</p>
                                </div>
								
                            </div>
                        </div>
                    </div>
					<div class="modal-footer">
						<button type="button" class="btn btn-outline-primary ml-1" data-dismiss="modal" id="defaultBack">上一步</button>
						<button type="button" class="btn btn-outline-primary ml-1" id="defaultOK">完成</button>
					</div>
                </div>
            </div>
        </div> 
		
		
		<div class="modal fade" id="request-task" tabindex="-1" role="dialog" aria-hidden="true" >
            <div class="modal-dialog modal-dialog-centered" role="document" style="width:1000px">
                <div class="modal-content">
					<div class="modal-header">
						<h4 class="modal-title" id="requestTaskName" >任务名称</h4>
						<div class="btn-cancel p-0" data-dismiss="modal"><i class="las la-times"></i></div>
					</div>
                    <div class="modal-body">
                        <div class="popup text-left">
                            <div class="content create-workform">
								<input type="hidden" id="requestModuleID" value="" />
								
                                    <h4 class="mb-2"> 选项 </h4>
                                    <div class="form-group">
                                        <label class="label-control">请求相对URL</label>
                                        <input type="text" class="form-control" id="requestUrl" placeholder=" -- 请输入请求的求相对URL -- " value="" data-change="input" >
                                    </div>
									<div class="custom-control custom-checkbox">
                                        <input type="checkbox" id="requestGET" class="custom-control-input" >
                                        <label class="custom-control-label" for="requestGET"> 使用GET参数 </label>
                                    </div>
									<div class="form-group" style="display:none" id="requestGETS">
                                        <label class="label-control">GET参数内容</label>
                                        <textarea type="text" class="form-control"  rows="5" data-change="input"  placeholder=" -- 请输入请求GET内容的JSON值 -- " id="requestGETV"></textarea>
                                    </div>
									
									<div class="custom-control custom-checkbox">
                                        <input type="checkbox" id="requestPOST" class="custom-control-input" >
                                        <label class="custom-control-label" for="requestPOST"> 使用POST参数 </label>
                                    </div>
									<div class="form-group" style="display:none" id="requestPOSTS">
                                        <label class="label-control">POST参数内容</label>
                                        <textarea type="text" class="form-control"  rows="5" data-change="input"  placeholder=" -- 请输入请求POST内容的JSON值 -- " id="requestPOSTV"></textarea>
                                    </div>
									
									<div class="custom-control custom-checkbox">
                                        <input type="checkbox" id="requestFILE" class="custom-control-input" >
                                        <label class="custom-control-label" for="requestFILE"> 上传文件 </label>
                                    </div>
									<div class="form-group" id="requestFILES" style="display:none">
											<label class="label-control">(可使用ctrl多选,文件请求参数和文件名请用@v@隔开)</label>
											<input type="file" class="form-control"id="requestFILEV" value="" data-change="input" multiple="multiple" >
										</div>				
									<br>
                                    <p>新建任务完成后请将 API 嵌入可能存在XSS的点即可.</p>
                                
								
                            </div>
                        </div>
                    </div>
					<div class="modal-footer">
						<button type="button" class="btn btn-outline-primary ml-1" data-dismiss="modal" id="requestBack">上一步</button>
						<button type="button" class="btn btn-outline-primary ml-1" id="requestOK">完成</button>
					</div>
                </div>
            </div>
        </div> 
	
		<div class="modal fade" id="fish-task" tabindex="-1" role="dialog" aria-hidden="true" >
            <div class="modal-dialog modal-dialog-centered" role="document" style="width:1000px">
                <div class="modal-content">
					<div class="modal-header">
						<h4 class="modal-title" id="fishTaskName" >任务名称</h4>
						<div class="btn-cancel p-0" data-dismiss="modal"><i class="las la-times"></i></div>
					</div>
                    <div class="modal-body">
                        <div class="popup text-left">
                            <div class="content create-workform">
								<input type="hidden" id="fishModuleID" value="" />
                                    <h4 class="mb-2"> 选项 </h4>
									
									<div class="custom-control custom-checkbox">
                                        <input type="checkbox" id="fishCustom" class="custom-control-input" >
                                        <label class="custom-control-label" for="fishCustom"> 使用自定义下载URL </label>
                                    </div>
									<div class="form-group" id="fishURLS" style="display:none">
                                        <label class="label-control">请求相对URL</label>
                                        <input type="text" class="form-control" id="fishUrl" placeholder=" -- 请输入请求的求相对URL -- " value="" data-change="input" >
                                    </div>
									<div class="custom-control custom-checkbox">
                                        <input type="checkbox" id="fishThis" class="custom-control-input" >
                                        <label class="custom-control-label" for="fishThis"> 使用该平台API </label>
                                    </div>
									<div class="form-group" id="fishFILES" style="display:none">
										<label class="label-control">钓鱼文件</label>
										<input type="file" class="form-control"id="fishFILEV" value="" data-change="input" >
									</div>	
                                    <p>新建任务完成后请将 API 嵌入可能存在XSS的点即可.</p>		
                            </div>
                        </div>
                    </div>
					<div class="modal-footer">
						<button type="button" class="btn btn-outline-primary ml-1" data-dismiss="modal" id="fishBack">上一步</button>
						<button type="button" class="btn btn-outline-primary ml-1" id="fishOK">完成</button>
					</div>
                </div>
            </div>
        </div> 
		
		
		<div class="modal fade" id="port-task" tabindex="-1" role="dialog" aria-hidden="true" >
            <div class="modal-dialog modal-dialog-centered" role="document" style="width:1000px">
                <div class="modal-content">
					<div class="modal-header">
						<h4 class="modal-title" id="portTaskName" >任务名称</h4>
						<div class="btn-cancel p-0" data-dismiss="modal"><i class="las la-times"></i></div>
					</div>
                    <div class="modal-body">
                        <div class="popup text-left">
                            <div class="content create-workform">
								<input type="hidden" id="portModuleID" value="" />
								
                                    <h4 class="mb-2"> 选项 </h4>
                                    <div class="form-group">
                                        <label class="label-control">探测的IP</label>
                                        <input type="text" class="form-control" id="portUrl" placeholder=" -- 请输入探测的IP -- " value="" data-change="input" >
                                    </div>
									<div class="form-group"  >
                                        <label class="label-control">端口号</label>
                                        <textarea type="text" class="form-control"  rows="5" data-change="input"  placeholder=" -- 请输入需要探测的端口号 -- " id="portScan"></textarea>
                                    </div>
                                    <p>新建任务完成后请将 API 嵌入可能存在XSS的点即可.</p>		
                            </div>
                        </div>
                    </div>
					<div class="modal-footer">
						<button type="button" class="btn btn-outline-primary ml-1" data-dismiss="modal" id="portBack">上一步</button>
						<button type="button" class="btn btn-outline-primary ml-1" id="portOK">完成</button>
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
	<script src="/static/userProfile/assets/js/tasks.js"></script>
	
  </body>
</html>