<template>
  <ul class="list">
    <li class="item" v-for="(item, index) in data" :key="index">
      <a :href="item.url">
        <img :src="item.img" alt="" :style="imgAttr">
        <div class="info">
          <div class="left">
            <div class="title hover" :style="titleAttr">{{item.title}}</div>
            <div class="desc" :style="descAttr">{{item.desc}}</div>
            <div class="author" :style="authorAttr">{{item.author}}</div>
            <div class="price-info" :style="priceInfoAttr"><span class="price" :style="priceAttr">￥{{item.price}}</span> | {{item.mark}} 订阅</div>
          </div>
          <div class="right">
            <a-button type="danger" :size="size" style="background: #ff7000">
              立即订阅
            </a-button>
          </div>
        </div>
      </a>
    </li>
  </ul>
</template>

<script>
export default {
  data() {
    return {
    }
  },
  props: {
    "data": {
      type: Array,
      default: function() {
        return [
          {
            img: "https://images.gitbook.cn/Fu479gYM0LhbsbMeGJ32xs3MkrtJ?imageMogr2/thumbnail/400x",
            title: "大数据开发",
            desc: "备战大数据，一起跨入5G时代",
            author: "SOCKSTACK - 高级工程师",
            price: 19,
            mark: 1000,
            url: "/info",
          },
          {
            url: "/info",
            img: "https://images.gitbook.cn/Fu479gYM0LhbsbMeGJ32xs3MkrtJ?imageMogr2/thumbnail/400x",
            title: "大数据开发",
            desc: "备战大数据，一起跨入5G时代",
            author: "SOCKSTACK - 高级工程师",
            price: 19,
            mark: 1000
          },
          {
            url: "/info",
            img: "https://images.gitbook.cn/Fu479gYM0LhbsbMeGJ32xs3MkrtJ?imageMogr2/thumbnail/400x",
            title: "大数据开发",
            desc: "备战大数据，一起跨入5G时代",
            author: "SOCKSTACK - 高级工程师",
            price: 19,
            mark: 1000
          },
          {
            url: "/info",
            img: "https://images.gitbook.cn/Fu479gYM0LhbsbMeGJ32xs3MkrtJ?imageMogr2/thumbnail/400x",
            title: "大数据开发",
            desc: "备战大数据，一起跨入5G时代",
            author: "SOCKSTACK - 高级工程师",
            price: 19,
            mark: 1000
          },
          {
            url: "/info",
            img: "https://images.gitbook.cn/Fu479gYM0LhbsbMeGJ32xs3MkrtJ?imageMogr2/thumbnail/400x",
            title: "大数据开发",
            desc: "备战大数据，一起跨入5G时代",
            author: "SOCKSTACK - 高级工程师",
            price: 19,
            mark: 1000
          },
        ]
      }
    },
    "imgAttr": {
      type: Object,
      default: function() {
        return {
          "width": "120px",
          "max-width": "120px",
        }
      }
    },
    "titleAttr": {
      type: Object,
      default: function() {
        return {
        }
      }
    },
    "descAttr": {
      type: Object,
      default: function() {
        return {
        }
      }
    },
    "authorAttr": {
      type: Object,
      default: function() {
        return {
        }
      }
    },
    "priceAttr": {
      type: Object,
      default: function() {
        return {
        }
      }
    },
    "priceInfoAttr": {
      type: Object,
      default: function() {
        return {
        }
      }
    },
    "size": {
      type: String,
      default: function() {
        return "small"
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.list {
  .item {
    margin-top: 20px;
    a {
      text-decoration: none;
      color: $main-title-color;
      display: flex;
    }
  }
  .info {
    width: 100%;
    padding: 20px;
    border: 1px solid $main-border-color;
    border-left: none;
    display: flex;
    justify-content: space-between;
    align-items: center;

    .left {
      align-self: flex-start;
    }
  }
  .title {
    display: inline-block;
    vertical-align: middle;
    color: #222;
    font-size: 18px;
    font-weight: 500;
    max-width: 394px;
    overflow: hidden;
    white-space: nowrap;
    margin: 0;
    line-height: 22px;
  }
  .desc {
    color: $main-desc-color;
    font-size: 14px;
    line-height: 14px;
    margin-top: 20px;
    overflow: hidden;
    text-overflow: ellipsis;
    -ms-text-overflow: ellipsis;
    white-space: nowrap;
  }
  .author {
    color: $main-desc-color;
    font-size: 12px;
    line-height: 12px;
    margin-top: 20px;
  }
  .price-info {
    margin-top: 10px;
    color: $main-desc-color;
    font-size: 12px;
    line-height: 12px;
  }
}
</style>>
