<template>
  <ul class="list">
    <li class="item" v-for="(item, index) in data" :key="index">
      <a href="">
        <img :src="item.img" alt="" :style="imgAttr">
        <div class="info">
          <div class="left">
            <div class="title hover" :style="titleAttr"><div class="article-tag no-shrink">最新</div>{{item.title}}</div>
            <div class="desc" :style="descAttr">{{item.desc}}</div>
            <div class="author" :style="authorAttr"><img src="https://images.gitbook.cn/4e9470c0-b75e-11e8-b809-d915783fbfe8?imageView2/1/w/400/h/400" alt="" data-v-0661afbb=""> {{item.author}}</div>
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
            img: "https://images.gitbook.cn/88807d40-1cc1-11eb-9777-a528c454293d?imageslim",
            title: "大数据开发",
            desc: "备战大数据，一起跨入5G时代",
            author: "SOCKSTACK - 高级工程师",
          },
          {
            img: "https://images.gitbook.cn/88807d40-1cc1-11eb-9777-a528c454293d?imageslim",
            title: "大数据开发",
            desc: "备战大数据，一起跨入5G时代",
            author: "SOCKSTACK - 高级工程师",
          },
          {
            img: "https://images.gitbook.cn/88807d40-1cc1-11eb-9777-a528c454293d?imageslim",
            title: "大数据开发",
            desc: "备战大数据，一起跨入5G时代",
            author: "SOCKSTACK - 高级工程师",
          },
          {
            img: "https://images.gitbook.cn/88807d40-1cc1-11eb-9777-a528c454293d?imageslim",
            title: "大数据开发",
            desc: "备战大数据，一起跨入5G时代",
            author: "SOCKSTACK - 高级工程师",
          },
          {
            img: "https://images.gitbook.cn/88807d40-1cc1-11eb-9777-a528c454293d?imageslim",
            title: "大数据开发",
            desc: "备战大数据，一起跨入5G时代",
            author: "SOCKSTACK - 高级工程师",
          },
        ]
      }
    },
    "imgAttr": {
      type: Object,
      default: function() {
        return {
          "width": "100px",
          "height": "100px",
          "max-width": "100px",
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
        return "large"
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.list {
  border-right: 1px solid #f2f5fa;
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
    padding: 10px 20px;
    border-left: none;
    display: flex;
    justify-content: space-between;
    align-items: flex-start;

    .left {
      align-self: flex-start;
    }
  }
  .title {
    display: flex;
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
    margin: 8px 0 10px;
    overflow: hidden;
    text-overflow: ellipsis;
    -ms-text-overflow: ellipsis;
    white-space: nowrap;
  }
  .author {
    color: $main-desc-color;
    font-size: 12px;
    display: flex;
    align-items: center;
    img {
        width: 20px;
        height: 20px;
        border-radius: 100%;
        margin-right: 10px;
    }
  }
  .price-info {
    margin-top: 25px;
    color: $main-desc-color;
  }
  .article-tag {
      color: #FF700A;
      border: 1px solid #FF700A;
      font-size: 12px;
      padding: 0px 5px;
      border-radius: 5px;
      margin-right: 10px;
  }
}
</style>>
