<template>
  <div class="info">
    <div class="left"><img src="https://images.gitbook.cn/Fu479gYM0LhbsbMeGJ32xs3MkrtJ?imageMogr2/thumbnail/400x" alt=""></div>
    <div class="right">
      <ul>
        <li class="title">大数据开发</li>
        <li class="desc">备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代备战大数据，一起步入5G时代</li>
        <li class="author">SOCKSTACK - 高级工程师</li>
        <li class="price">¥ 19 - 1000 订阅</li>
      </ul>
      <a-button type="danger" size="small" style="background: #ff7000;">
        立即订阅
      </a-button>
    </div>
  </div>
</template>

<script>
export default {
}
</script>

<style lang="scss" scoped>
  .info {
    display: flex;
    padding-bottom: 15px;
    border-bottom: 1px solid $main-border-color;
    .left {
      width: 120px;
      max-width: 120px;
      img {
        width: 100%;
      }
    }
    .right {
      padding: 10px 20px;
      display: flex;
      justify-content: space-between;
      width: calc(100% - 120px);
      align-items: center;

      ul {
        display: flex;
        flex-direction: column;
        height: 100%;
        width: 100%;
      }

      .title {
        font-size: 18px;
      }
      .desc, .author, .price {
        font-size: 14px;
        color: $main-desc-color;
      }
      .desc {
        margin-top: 16px;
        max-width: 80%;
        overflow : hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 4;
        -webkit-box-orient: vertical;
      }
      .author {
        margin-top: 20px;
        font-size: 13px;
      }
      .price {
        margin-top: 10px;
        font-size: 13px;
      }
    }
  }
</style>
