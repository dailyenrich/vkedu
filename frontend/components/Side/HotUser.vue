<template>
  <div class="hot-quan side-padding-20 side-border side-margin-bottom-20">
    <div class="title">
      活跃学员
    </div>
    <ul class="hot-list">
      <li>
        <a href="">
          <div><img src="https://images.gitbook.cn/4e9470c0-b75e-11e8-b809-d915783fbfe8?imageView2/1/w/400/h/400" alt=""></div>
          <div>
            程序员小闲
            <div>大数据专家</div>
            <div>6篇文章</div>
          </div>
        </a>
      </li>
      <li>
        <a href="">
          <div><img src="https://images.gitbook.cn/4e9470c0-b75e-11e8-b809-d915783fbfe8?imageView2/1/w/400/h/400" alt=""></div>
          <div>
            程序员小闲
            <div>大数据专家</div>
            <div>6篇文章</div>
          </div>
        </a>
      </li>
      <li>
        <a href="">
          <div><img src="https://images.gitbook.cn/4e9470c0-b75e-11e8-b809-d915783fbfe8?imageView2/1/w/400/h/400" alt=""></div>
          <div>
            程序员小闲
            <div>大数据专家</div>
            <div>6篇文章</div>
          </div>
        </a>
      </li>
      <li>
        <a href="">
          <div><img src="https://images.gitbook.cn/4e9470c0-b75e-11e8-b809-d915783fbfe8?imageView2/1/w/400/h/400" alt=""></div>
          <div>
            程序员小闲
            <div>大数据专家</div>
            <div>6篇文章</div>
          </div>
        </a>
      </li>
      <li>
        <a href="">
          <div><img src="https://images.gitbook.cn/4e9470c0-b75e-11e8-b809-d915783fbfe8?imageView2/1/w/400/h/400" alt=""></div>
          <div>
            程序员小闲
            <div>大数据专家</div>
            <div>6篇文章</div>
          </div>
        </a>
      </li>
    </ul>
  </div>
</template>

<style lang="scss" scoped>
.hot-list {
  li {
    margin-bottom: 20px;
    a {
      display: flex;
      color: $main-desc-color;
      padding: 10px 0;
      & > div {
        display: flex;
        flex-direction: column;
        justify-content: center;
        & > div {
          font-size: 12px;
          margin-top: 8px;
        }
      }
    }
    &:last-child {
      margin-bottom: 0;
    }
    img {
      border-radius: 100%;
      width: 60px;
      height: 60px;
      margin-right: 20px;
    }
  }
}
.title {
  padding: 0px 0 20px;
  color: $main-title-color;
}
</style>
