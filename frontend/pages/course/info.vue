<template>
  <div>
    <div class="main">
      <div class="content">
        <div class="cource-list">
          <Intra />
          <Toc />
        </div>
      </div>
      <div class="side">
        <HotArticle />
        <Adimg />
        <HotUser />
        <HotVquan />
      </div>
    </div>
  </div>
</template>

<script>
import Toc from "~/components/Toc/Default";
import Intra from "~/components/Course/Intra";

export default {
  components: {
    Toc,
    Intra
  }
}
</script>

<style lang="scss">
.banner {
  width: 100%;
}

.main {
  display: flex;
  justify-content: space-between;
  .content {
    width: 70%;
    .cource-list .list .item:first-child {
      margin-top: 0px;
    }
  }
  .content-title {
    display: flex;
    margin-top: 20px;
    justify-content: space-between;

    .left, .right {
      display: flex;
    }
    .left {
      align-items: center;
      h4 {
        color: $main-title-color;
        font-size: 18px;
      }
      span {
        font-size: 14px;
        color: $main-desc-color;
        margin-left: 10px;
        margin-top: 2px;
      }
    }
    .right {
      align-items: center;
      a {
        color: $main-desc-color;
        font-size: 15px;
      }
    }
  }
}
.pagination {
  display: flex;
  justify-content: center;
  margin: 40px 0!important;
}
.ant-pagination-item-active {
  border-color: $main-font-color-active!important;
  a {
    color: $main-font-color-active!important;
  }
}
</style>
