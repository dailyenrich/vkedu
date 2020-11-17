<template>
  <div>
    <div class="main">
      <div class="content">
        <div class="banner">
          <Banner />
        </div>
        <div class="hot-cource">
          <div class="content-title">
            <div class="left">
              <h4>热门课程</h4>
              <span>新鲜热辣的课程，你值得拥有</span>
            </div>
            <div class="right"><a href="">查看更多 ></a></div>
          </div>
          <CourseList />
        </div>
        <div class="hot-cource">
          <div class="content-title">
            <div class="left">
              <h4>最新话题</h4>
              <span>新鲜出炉的话题，快来围观</span>
            </div>
            <div class="right"><a href="">查看更多 ></a></div>
          </div>
          <ArticleList />
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
import CourseList from "~/components/Course/HotList";
import ArticleList from "~/components/Article/List";

export default {
  components: {
    CourseList,
    ArticleList
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
</style>
