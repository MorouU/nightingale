

<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
      <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
      <title>NI - XSS 主页</title>
      
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
                     <a href="#"><img src="/static/userProfile/assets/images/logo.png" class="ri-menu-line wrapper-menu" alt="logo"></a>
                  </div>
              </nav>
          </div>
      </div>      
       <div class="iq-sidebar  sidebar-default ">
          <div class="iq-sidebar-logo d-flex align-items-center justify-content-between">
              <a href="/preview/1097306/2021-01-28/note/backend/index.html" class="header-logo">
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
                          <h4 class="text-capitalize">主页</h4>
                      </div>
                  </div>
              </div>
          </div>        
		</div>
      <div class="container-fluid">
         <div class="row">
            <div class="col-lg-4 col-md-4 col-sm-4">
               <div class="card card-block card-stretch card-height">
                  <div class="card-body text-center">
                     <svg width="36" height="48" viewBox="0 0 36 48" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path
                           d="M9.10495 33.9964C8.29026 33.1817 8.71495 33.4114 6.74995 32.8855C5.85838 32.6464 5.07463 32.1871 4.36588 31.6367L0.112441 42.0655C-0.299122 43.0752 0.469629 44.1721 1.559 44.1308L6.4987 43.9424L9.8962 47.5311C10.6462 48.3224 11.9624 48.0758 12.374 47.0661L17.2537 35.1017C16.2375 35.668 15.1096 35.9999 13.9434 35.9999C12.1153 35.9999 10.3978 35.2883 9.10495 33.9964V33.9964ZM35.8875 42.0655L31.634 31.6367C30.9253 32.188 30.1415 32.6464 29.25 32.8855C27.2746 33.4142 27.7078 33.1836 26.895 33.9964C25.6021 35.2883 23.8837 35.9999 22.0556 35.9999C20.8893 35.9999 19.7615 35.6671 18.7453 35.1017L23.625 47.0661C24.0365 48.0758 25.3537 48.3224 26.1028 47.5311L29.5012 43.9424L34.4409 44.1308C35.5303 44.1721 36.299 43.0742 35.8875 42.0655V42.0655ZM24.6562 31.8749C26.0887 30.4171 26.2528 30.5427 28.2928 29.9867C29.595 29.6314 30.6131 28.5955 30.9618 27.2699C31.6631 24.6074 31.4812 24.9289 33.3946 22.9808C34.3481 22.0105 34.7203 20.5958 34.3715 19.2702C33.6712 16.6096 33.6703 16.9808 34.3715 14.3174C34.7203 12.9917 34.3481 11.5771 33.3946 10.6067C31.4812 8.65862 31.6631 8.97925 30.9618 6.31768C30.6131 4.99206 29.595 3.95612 28.2928 3.60081C25.679 2.88737 25.994 3.07393 24.0787 1.12487C23.1253 0.154558 21.735 -0.225129 20.4328 0.130183C17.82 0.842683 18.1846 0.843621 15.5671 0.130183C14.2649 -0.225129 12.8746 0.153621 11.9212 1.12487C10.0078 3.073 10.3228 2.88737 7.70807 3.60081C6.40588 3.95612 5.38776 4.99206 5.03901 6.31768C4.33869 8.97925 4.51963 8.65862 2.60619 10.6067C1.65275 11.5771 1.27963 12.9917 1.62932 14.3174C2.32963 16.9761 2.33057 16.6049 1.62932 19.2692C1.28057 20.5949 1.65275 22.0096 2.60619 22.9808C4.51963 24.9289 4.33776 24.6074 5.03901 27.2699C5.38776 28.5955 6.40588 29.6314 7.70807 29.9867C9.8062 30.5586 9.96276 30.4686 11.3437 31.8749C12.584 33.1377 14.5162 33.3636 16.0068 32.4205C16.6029 32.0421 17.2944 31.8411 18.0004 31.8411C18.7065 31.8411 19.3979 32.0421 19.994 32.4205C21.4837 33.3636 23.4159 33.1377 24.6562 31.8749ZM9.15557 16.4961C9.15557 11.5246 13.1156 7.49425 18 7.49425C22.8843 7.49425 26.8443 11.5246 26.8443 16.4961C26.8443 21.4677 22.8843 25.498 18 25.498C13.1156 25.498 9.15557 21.4677 9.15557 16.4961V16.4961Z"
                           fill="#88DFFF" />
                     </svg>
                     <h2 class="mb-2 mt-3">{{ .taskTotal }}</h2>
                     <h4>任务的总数</h4>
                  </div>
               </div>
            </div>
            <div class="col-lg-4 col-md-4 col-sm-4">
               <div class="card card-block card-stretch card-height">
                  <div class="card-body text-center">
                     <svg width="60" height="48" viewBox="0 0 60 48" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path
                           d="M23.9091 24.5297C24.495 25.1156 25.4447 25.1156 26.0306 24.5297L27.0909 23.4694C27.6769 22.8834 27.6769 21.9338 27.0909 21.3478L23.7422 18L27.09 14.6512C27.6759 14.0653 27.6759 13.1156 27.09 12.5297L26.0297 11.4694C25.4437 10.8834 24.4941 10.8834 23.9081 11.4694L18.4387 16.9387C17.8528 17.5247 17.8528 18.4744 18.4387 19.0603L23.9091 24.5297V24.5297ZM32.91 23.4703L33.9703 24.5306C34.5563 25.1166 35.5059 25.1166 36.0919 24.5306L41.5613 19.0613C42.1472 18.4753 42.1472 17.5256 41.5613 16.9397L36.0919 11.4703C35.5059 10.8844 34.5563 10.8844 33.9703 11.4703L32.91 12.5306C32.3241 13.1166 32.3241 14.0662 32.91 14.6522L36.2578 18L32.91 21.3488C32.3241 21.9347 32.3241 22.8844 32.91 23.4703V23.4703ZM58.5 39H35.7694C35.7 40.8572 34.3903 42 32.7 42H27C25.2478 42 23.9044 40.3622 23.9278 39H1.5C0.675 39 0 39.675 0 40.5V42C0 45.3 2.7 48 6 48H54C57.3 48 60 45.3 60 42V40.5C60 39.675 59.325 39 58.5 39ZM54 4.5C54 2.025 51.975 0 49.5 0H10.5C8.025 0 6 2.025 6 4.5V36H54V4.5ZM48 30H12V6H48V30Z"
                           fill="#81F4B2" />
                     </svg>
                     <h2 class="mb-2 mt-3">{{ .enableTaskNum }}</h2>
                     <h4>启用的任务数</h4>
                  </div>
               </div>
            </div>
            <div class="col-lg-4 col-md-4 col-sm-4">
               <div class="card card-block card-stretch card-height">
                  <div class="card-body text-center">
                     <svg width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path
                           d="M35.3676 11.2517C34.8398 11.2372 34.3256 11.3198 33.8438 11.4898V7.03125C33.8438 4.70503 31.9512 2.8125 29.625 2.8125C29.0759 2.8125 28.5517 2.91909 28.0701 3.11072C27.5821 1.32047 25.9428 0 24 0C22.0572 0 20.4179 1.32047 19.9299 3.11072C19.4483 2.91909 18.9241 2.8125 18.375 2.8125C16.0488 2.8125 14.1562 4.70503 14.1562 7.03125V11.4895C13.6747 11.3198 13.1607 11.2372 12.6324 11.2517C10.3711 11.3136 8.53125 13.2316 8.53125 15.5272V48H36.6562V41.2395L38.5637 36.4704C39.1643 34.9689 39.4688 33.3877 39.4688 31.7705V15.5272C39.4688 13.2316 37.6289 11.3136 35.3676 11.2517V11.2517ZM11.3438 45.1875V42.375H33.8438V45.1875H11.3438ZM36.6562 31.7705C36.6562 33.0283 36.4194 34.2581 35.9523 35.4261L34.2979 39.5625H11.3438V15.5272C11.3438 14.7405 11.9564 14.0837 12.7095 14.0631C13.0926 14.0504 13.4561 14.1937 13.7305 14.4607C14.0051 14.7278 14.1563 15.0858 14.1563 15.4687V21.9843H16.9688V7.03125C16.9688 6.25584 17.5997 5.625 18.3751 5.625C19.1505 5.625 19.7813 6.25584 19.7813 7.03125V21.9844H22.5938V4.21875C22.5938 3.44334 23.2247 2.8125 24.0001 2.8125C24.7755 2.8125 25.4063 3.44334 25.4063 4.21875V21.9844H28.2188V7.03125C28.2188 6.25584 28.8497 5.625 29.6251 5.625C30.4005 5.625 31.0313 6.25584 31.0313 7.03125V24.1714C24.712 24.8732 19.7812 30.2467 19.7812 36.75H22.5938C22.5938 31.3222 27.0097 26.9062 32.4375 26.9062H33.8438V15.4688C33.8438 15.0859 33.995 14.7278 34.2696 14.4608C34.544 14.1938 34.9067 14.0508 35.2906 14.0632C36.0436 14.0838 36.6562 14.7406 36.6562 15.5273V31.7705Z"
                           fill="#FF9BBD" />
                     </svg>
                     <h2 class="mb-2 mt-3">{{ .disableTaskNum }}</h2>
                     <h4>停用的任务数</h4>
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
									<a class="bg-info" href="#" id="findBack" >
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
	<script src="/static/userProfile/assets/js/home.js"></script>
	
  </body>
</html>