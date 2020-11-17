<template>
  <div class="hot-quan side-padding-20 side-border side-margin-bottom-20">
    <div class="title">
      热门微圈
    </div>
    <ul class="hot-list">
      <li>
        <a href="">小程序填坑</a>
      </li>
      <li>
        <a href="">韭菜聊技术</a>
      </li>
      <li>
        <a href="">前端交流</a>
      </li>
      <li>
        <a href="">后端大杂烩</a>
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
      font-size: 14px;
    }
    &:last-child {
      margin-bottom: 0;
    }
  }
}
.title {
  padding: 0px 0 20px;
  color: $main-title-color;
  font-size: 16px;
}
</style>
