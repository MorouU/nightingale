<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <meta content="width=device-width, initial-scale=1.0" name="viewport">

  <title>NI - XSS 攻击平台</title>
  <meta content="" name="descriptison">
  <meta content="" name="keywords">

  <script src="/static/index/assets/js/query.js"></script>

  <!-- Favicons -->
  <link href="/static/index/assets/img/apple-touch-icon.png" rel="icon">
  <link href="/static/index/assets/img/apple-touch-icon.png" rel="apple-touch-icon">

  <!-- Google Fonts -->
  <link href="https://fonts.googleapis.com/css?family=Open+Sans:300,300i,400,400i,600,600i,700,700i|Raleway:300,300i,400,400i,600,600i,700,700i,900" rel="stylesheet">

  <!-- Vendor CSS Files -->
  <link href="/static/index/assets/vendor/bootstrap/css/bootstrap.min.css" rel="stylesheet">
  <link href="/static/index/assets/vendor/icofont/icofont.min.css" rel="stylesheet">
  <link href="/static/index/assets/vendor/boxicons/css/boxicons.min.css" rel="stylesheet">
  <link href="/static/index/assets/vendor/animate.css/animate.min.css" rel="stylesheet">
  <link href="/static/index/assets/vendor/venobox/venobox.css" rel="stylesheet">
  <link href="/static/index/assets/vendor/owl.carousel/assets/owl.carousel.min.css" rel="stylesheet">


  <!-- Template Main CSS File -->
  <link href="/static/index/assets/css/style.css" rel="stylesheet">

  <!-- =======================================================
  * Template Name: Mamba - v2.4.0
  * Template URL: https://bootstrapmade.com/mamba-one-page-bootstrap-template-free/
  * Author: BootstrapMade.com
  * License: https://bootstrapmade.com/license/
  ======================================================== -->
</head>

<body>

  <!-- ======= Top Bar ======= -->


  <!-- ======= Header ======= -->
  <header id="header">
    <div class="container">

      <div class="logo float-left">
        <h1 class="text-light"><a href="/"><span>NI XSS PLATFORM </span></a></h1>
        <!-- Uncomment below if you prefer to use an image logo -->
        <!-- <a href="/"><img src="/static/index/assets/img/logo.png" alt="" class="img-fluid"></a>-->
      </div>

      <nav class="nav-menu float-right d-none d-lg-block">
        <ul>
          <li class="active"><a href="/#">主页</a></li>
          <li><a href="/userEnter/login">用户登录</a></li>
		  <li><a href="/adminEnter/login">管理员登录</a></li>
          <li><a href="/#contact">给我们反馈</a></li>
        </ul>
      </nav><!-- .nav-menu -->

    </div>
  </header><!-- End Header -->

  <!-- ======= Hero Section ======= -->
  <section id="hero">
    <div class="hero-container">
      <div id="heroCarousel" class="carousel slide carousel-fade" data-ride="carousel">

        <ol class="carousel-indicators" id="hero-carousel-indicators"></ol>

        <div class="carousel-inner" role="listbox">

          <!-- Slide 1 -->
          <div class="carousel-item active" style="background-image: url('/static/index/assets/img/slide/slide-1.jpg');">
            <div class="carousel-container">
              <div class="carousel-content container">
                <h2 class="animate__animated animate__fadeInDown">欢迎使用<span>NI XSS攻击平台</span></h2>
                <p class="animate__animated animate__fadeInUp">使用该产品须知：</br>
				</br><span style="font-size:larger">1. 该产品由 Morouu 独立完成</span>
				</br><span style="font-size:larger">2. 该产品仅限学习使用，请勿投入于真实攻击中</span>
				</p>
                <a href="/#about" class="btn-get-started animate__animated animate__fadeInUp scrollto">更多该产品详情</a>
              </div>
            </div>
          </div>

        </div>

        <a class="carousel-control-prev" href="/#heroCarousel" role="button" data-slide="prev">
          <span class="carousel-control-prev-icon icofont-rounded-left" aria-hidden="true"></span>
          <span class="sr-only">Previous</span>
        </a>
        <a class="carousel-control-next" href="/#heroCarousel" role="button" data-slide="next">
          <span class="carousel-control-next-icon icofont-rounded-right" aria-hidden="true"></span>
          <span class="sr-only">Next</span>
        </a>

      </div>
    </div>
  </section><!-- End Hero -->

  <main id="main">



    <!-- ======= Counts Section ======= -->
    <section class="counts section-bg">
      <div class="container">

        <div class="row">

          <div class="col-lg-3 col-md-6 text-center" data-aos="fade-up">
            <div class="count-box">
              <i class="icofont-simple-smile" style="color: #20b38e;"></i>
              <span data-toggle="counter-up">23,333</span>
              <p>好评数</p>
            </div>
          </div>

          <div class="col-lg-3 col-md-6 text-center" data-aos="fade-up" data-aos-delay="200">
            <div class="count-box">
              <i class="icofont-document-folder" style="color: #c042ff;"></i>
              <span data-toggle="counter-up">0</span>
              <p>退款数</p>
            </div>
          </div>

          <div class="col-lg-3 col-md-6 text-center" data-aos="fade-up" data-aos-delay="400">
            <div class="count-box">
              <i class="icofont-live-support" style="color: #46d1ff;"></i>
              <span data-toggle="counter-up">1,646,122</span>
              <p>下单量</p>
            </div>
          </div>

          <div class="col-lg-3 col-md-6 text-center" data-aos="fade-up" data-aos-delay="600">
            <div class="count-box">
              <i class="icofont-users-alt-5" style="color: #ffb459;"></i>
              <span data-toggle="counter-up">150,000</span>
              <p>专业团队人数</p>
            </div>
          </div>

        </div>

      </div>
    </section><!-- End Counts Section -->

    <!-- ======= Our Team Section ======= -->
    <section id="team" class="team">
      <div class="container">

        <div class="section-title">
          <h2>科研人员</h2>
          <p>包括但不仅限于以下几人.</p>
        </div>

        <div class="row">


          <div class="col-xl-3 col-lg-4 col-md-6" data-aos="fade-up" data-aos-delay="100">
            <div class="member">
              <div class="pic"><img src="/static/index/assets/img/team/team-2.jpg" class="img-fluid" alt=""></div>
              <div class="member-info">
                <h4>Morouu</h4>
                <span>总工程师</span>
                <div class="social">
                  <a href="/"><i class="icofont-twitter"></i></a>
                  <a href="/"><i class="icofont-facebook"></i></a>
                  <a href="/"><i class="icofont-instagram"></i></a>
                  <a href="/"><i class="icofont-linkedin"></i></a>
                </div>
              </div>
            </div>
          </div>




        </div>

      </div>
    </section><!-- End Our Team Section -->

    <!-- ======= Frequently Asked Questions Section ======= -->
    <section id="faq" class="faq section-bg">
      <div class="container">

        <div class="section-title">
          <h2>有奖竞答</h2>
        </div>

        <div class="row  d-flex align-items-stretch" id="questions">

          <div class="col-lg-6 faq-item" data-aos="fade-up">
            <h4>为啥没有人退款?</h4>
            <p>
              我们也不清楚哦，可能是产品质量太好了呢.
            </p>
          </div>

          <div class="col-lg-6 faq-item" data-aos="fade-up" data-aos-delay="100">
            <h4>开发人员就你一个啊?</h4>
            <p>
			  是的哦！
			  </p>
          </div>

          <div class="col-lg-6 faq-item" data-aos="fade-up" data-aos-delay="200">
            <h4>啥是专业团队?</h4>
            <p>
              专业提供售后服务哦.
            </p>
          </div>

          <div class="col-lg-6 faq-item" data-aos="fade-up" data-aos-delay="300">
            <h4>NI有啥含义吗?</h4>
            <p>
			   就是夜莺啦，没想到啥好名字 呜呜呜/(ㄒoㄒ)/~~
			 </p>
          </div>


          <div class="col-lg-6 faq-item" data-aos="fade-up" data-aos-delay="500">
            <h4>我要如何与你们联系?</h4>
            <p>
             联系在下边哦，您可以随时反馈任何问题的.
            </p>
          </div>

        </div>

      </div>
    </section><!-- End Frequently Asked Questions Section -->

    <!-- ======= Contact Us Section ======= -->
    <section id="contact" class="contact">
      <div class="container">

        <div class="section-title">
          <h2>联系我们</h2>
        </div>

        <div class="row">

          <div class="col-lg-6 d-flex align-items-stretch" data-aos="fade-up">
            <div class="info-box">
              <i class="bx bx-map"></i>
              <h3>我们的地址</h3>
              <p>C星系, 新一星区, NY535022 星</p>
            </div>
          </div>

          <div class="col-lg-3 d-flex align-items-stretch" data-aos="fade-up" data-aos-delay="100">
            <div class="info-box">
              <i class="bx bx-envelope"></i>
              <h3>邮件联系</h3>
              <p><a href="#" class="__cf_email__" data-cfemail="9df4f3fbf2ddf8e5fcf0edf1f8b3fef2f0">[123456@654321.boom]</a><br><a href="#" class="__cf_email__" data-cfemail="593a36372d383a2d193c21383429353c773a3634">[abcdef@fedcba.baozha]</a></p>
            </div>
          </div>

          <div class="col-lg-3 d-flex align-items-stretch" data-aos="fade-up" data-aos-delay="200">
            <div class="info-box ">
              <i class="bx bx-phone-call"></i>
              <h3>电话联系</h3>
              <p>+C 00001 53502 250<br>+C 00001 05202 535</p>
            </div>
          </div>

          <div class="col-lg-12" data-aos="fade-up" data-aos-delay="300">
            <form  action="#contact" class="php-email-form">
              <div class="form-row" >
                <div class="col-lg-6 form-group">
                  <input type="text" class="form-control" id="name" placeholder="名字" data-rule="minlen:4" data-msg="Please enter at least 4 chars" />
                  <div class="validate"></div>
                </div>
                <div class="col-lg-6 form-group">
                  <input type="email" class="form-control"  id="email" placeholder="邮箱" data-rule="email" data-msg="Please enter a valid email" />
                  <div class="validate"></div>
                </div>
              </div>
              <div class="form-group">
                <input type="text" class="form-control"  id="subject" placeholder="您的问题" data-rule="minlen:4" data-msg="Please enter at least 8 chars of subject" />
                <div class="validate"></div>
              </div>
              <div class="form-group">
                <textarea class="form-control"  id="message" rows="5" data-rule="required" data-msg="Please write something for us" placeholder="请详细描述您的问题"></textarea>
                <div class="validate"></div>
              </div>
              <div class="mb-3">
                <div class="loading">Loading</div>
                <div class="error-message"></div>
                <div class="sent-message">您的反馈已经传达. 非常感谢您!</div>
              </div>
              <div class="text-center"><input type="button" onclick="feedback();" value="发送反馈"></div>
            </form>
          </div>

        </div>

      </div>
    </section><!-- End Contact Us Section -->

  </main><!-- End #main -->

  <!-- ======= Footer ======= -->
  <footer id="footer">
    <div class="footer-top">
      <div class="container">
        <div class="row">

          <div class="col-lg-3 col-md-6 footer-info">
            <h3>Mamba</h3>
            <p>
              C星系, <br>
              新一星区, NY535022 星<br><br>
              <strong>Phone:</strong> +1 5589 55488 55<br>
              <strong>Email:</strong> <a href="#" class="__cf_email__" data-cfemail="85ecebe3eac5e0fde4e8f5e9e0abe6eae8">[123456@654321.boom]</a><br>
            </p>
            <div class="social-links mt-3">
              <a href="/#" class="twitter"><i class="bx bxl-twitter"></i></a>
              <a href="/#" class="facebook"><i class="bx bxl-facebook"></i></a>
              <a href="/#" class="instagram"><i class="bx bxl-instagram"></i></a>
              <a href="/#" class="google-plus"><i class="bx bxl-skype"></i></a>
              <a href="/#" class="linkedin"><i class="bx bxl-linkedin"></i></a>
            </div>
          </div>

          <div class="col-lg-2 col-md-6 footer-links">
            <h4>有用的链接</h4>
            <ul>
				<li><i class="bx bx-chevron-right"></i> <a href="/#">主页</a></li>
				<li><i class="bx bx-chevron-right"></i><a href="/userEnter/login">用户登录</a></li>
				<li><i class="bx bx-chevron-right"></i><a href="/adminEnter/login">管理员登录</a></li>
				<li><i class="bx bx-chevron-right"></i><a href="/#contact">给我们反馈</a></li>
            </ul>
          </div>

        </div>
      </div>
    </div>

    <div class="container">
      <div class="copyright">
        &copy; Copyright <strong><span>Morouu</span></strong>. All Rights Reserved
      </div>
      <div class="credits">
        <!-- All the links in the footer should remain intact. -->
        <!-- You can delete the links only if you purchased the pro version. -->
        <!-- Licensing information: https://bootstrapmade.com/license/ -->
        <!-- Purchase the pro version with working PHP/AJAX contact form: https://bootstrapmade.com/mamba-one-page-bootstrap-template-free/ -->
        Designed by <a href="https://bootstrapmade.com/">BootstrapMade</a>
      </div>
    </div>
  </footer><!-- End Footer -->

  <a href="/#" class="back-to-top"><i class="icofont-simple-up"></i></a>

  <!-- Vendor JS Files -->
  <script src="/static/index/assets/vendor/bootstrap/js/bootstrap.bundle.min.js"></script>
  <script src="/static/index/assets/vendor/jquery.easing/jquery.easing.min.js"></script>
  <script src="/static/index/assets/vendor/php-email-form/validate.js"></script>
  <script src="/static/index/assets/vendor/jquery-sticky/jquery.sticky.js"></script>
  <script src="/static/index/assets/vendor/venobox/venobox.min.js"></script>
  <script src="/static/index/assets/vendor/waypoints/jquery.waypoints.min.js"></script>
  <script src="/static/index/assets/vendor/counterup/counterup.min.js"></script>
  <script src="/static/index/assets/vendor/owl.carousel/owl.carousel.min.js"></script>
  <script src="/static/index/assets/vendor/isotope-layout/isotope.pkgd.min.js"></script>
  <script src="/static/index/assets/vendor/aos/aos.js"></script>


  <!-- Template Main JS File -->
  <script src="/static/index/assets/js/main.js"></script>

  <script>if( window.self == window.top ) { (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){ (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o), m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m) })(window,document,'script','//www.google-analytics.com/analytics.js','ga'); ga('create', 'UA-55234356-4', 'auto'); ga('send', 'pageview'); } </script>
</body>

</html>

