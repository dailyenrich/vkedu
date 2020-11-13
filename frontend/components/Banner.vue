<template>
  <div class="swiper-container" v-swiper="swiperOption">
    <div class="swiper-wrapper">
      <div v-for="(item, index) in dataImage" :key="index" class="swiper-slide">
        <img :src="item.imgUrl" />
      </div>
    </div>
    <div class="swiper-pagination"></div><!--分页器。如果放置在swiper-container外面，需要自定义样式。-->
  </div>
</template>
<script>
export default {
  data() {
    return {
      dataImage: [
        {
          imgUrl: "https://img.mukewang.com/5eeb15170001b36418720764.jpg"
        },
        {
          imgUrl: "https://img.mukewang.com/5eeb15170001b36418720764.jpg"
        },
        {
          imgUrl: "https://img.mukewang.com/5eeb15170001b36418720764.jpg"
        },
        {
          imgUrl: "https://img.mukewang.com/5eeb15170001b36418720764.jpg"
        },
      ],
      swiperOption: { // 配置说明直接看官网
        loop: true,
        lazy: {
          loadPrevNext: true
        },
        pagination: {
          el: ".swiper-pagination",
          clickable: true,
        },
      }
    };
  },
};
</script>

<style lang="scss">
@import url("swiper/css/swiper.min.css");
img{
  width: 100%;
}
</style>