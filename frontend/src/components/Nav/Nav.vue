<template>
  <header class="header">
    <!--导航条-->
    <nav>
      <!--logo-->
      <a href="/" class="navbar-brand home">
        <img src="img/logo.png" alt="Vaniot logo" class="d-none d-md-inline-block">
        <img src="img/logo-small.png" alt="Vaniot logo" class="d-inline-block d-md-none">
        <span class="sr-only"> Vaniot-go to homepage</span>
      </a>

      <!--小屏幕的button-->
      <div class="navbar-buttons">
        <button type="button" data-toggle="collapse" data-target="#navigation" class="btn btn-outline-secondary navbar-toggler">
          <span class="sr-only">Toggle navigation</span>
          <i class="fa fa-align-justify"></i>
        </button>
        <button type="button" data-toggle="collapse" data-target="#search" class="btn btn-outline-secondary navbar-toggler">
          <span class="sr-only">Toggle search</span>
          <i class="fa fa-search"></i>
        </button>
      </div>
      <!--大屏幕的导航条-->
      <div class="navigation">
        <ul>
          <li class="nav-item">
            <a href="#" class="nav-link active">首页</a>
          </li>
          <li class="nav-item">
            <a href="#" class="nav-link active">分类</a>
          </li>
          <li class="nav-item">
            <a href="#" class="nav-link active">归档</a>
          </li>
        </ul>
        <!--点击出现搜索栏-->
        <div class="navbar-buttons d-flex justify-content-end">
          <!-- /.nav-collapse 搜索-->
          <div id="search-not-mobile" class="navbar-collapse collapse"></div><a data-toggle="collapse" href="#search" class="btn navbar-btn btn-primary d-none d-lg-inline-block"><span class="sr-only">Toggle search</span><i class="fa fa-search"></i></a>
        </div>
      </div>
    </nav>
    <!--下拉搜索框-->
    <div id="search" class="collapse">
      <div class="container">
        <form role="search" class="ml-auto">
          <div class="input-group">
            <input type="text" placeholder="Search" class="form-control">
            <div class="input-group-append">
              <button type="button" class="btn btn-primary"><i class="fa fa-search"></i></button>
            </div>
          </div>
        </form>
      </div>
    </div>
  </header>
</template>

<script>
    export default {
        name: "Nav"
    }
</script>

<style scoped>
  .navbar {
    border-bottom: 1px solid #e6e6e6;
    background-color: #fff;
  }

  .navbar .navbar-nav > .nav-item > .nav-link {
    text-transform: uppercase;
    letter-spacing: 0.1em;
    font-weight: 700;
    padding: 30px 15px !important;
    color: #555;
  }
</style>
