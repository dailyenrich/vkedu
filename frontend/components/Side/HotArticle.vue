<template>
  <div class="hot-article side-padding-20 side-border side-margin-bottom-20">
    <div class="title">
      热门文章
      <div></div>
    </div>
    <ul class="hot-list">
      <li>
        <div class="rank-tag no-shrink top3">1</div>
        <a href="">小程序用户行为数据采集器讲解</a>
      </li>
      <li>
        <div class="rank-tag no-shrink top3">2</div>
        <a href="">小程序用户行为数据采集器讲解</a>
      </li>
      <li>
        <div class="rank-tag no-shrink top3">3</div>
        <a href="">小程序用户行为数据采集器讲解</a>
      </li>
      <li>
        <div class="rank-tag no-shrink">4</div>
        <a href="">小程序用户行为数据采集器讲解</a>
      </li>
      <li>
        <div class="rank-tag no-shrink">5</div>
        <a href="">小程序用户行为数据采集器讲解</a>
      </li>
      <li>
        <div class="rank-tag no-shrink">6</div>
        <a href="">小程序用户行为数据采集器讲解</a>
      </li>
    </ul>
  </div>
</template>

<style lang="scss" scoped>
.hot-list {
  li {
    display: flex;
    margin-bottom: 20px;
    a {
      color: $main-desc-color;
    }
    &:last-child {
      margin-bottom: 0;
    }
  }
}
.rank-tag {
  width: 18px;
  height: 18px;
  color: rgba(170,176,193,0.3);
  font-family: bebas-regular, sans-serif;
  font-size: 16px;
  font-weight: bold;
  text-align: center;
  line-height: 18px;
  flex-shrink: 0;
  margin-right: 15px;
}
.top3 {
  background: orange;
  border-radius: 4px;
  color: #fff;
}
.title {
  padding: 0px 0 20px;
  color: $main-title-color;
}
</style>
