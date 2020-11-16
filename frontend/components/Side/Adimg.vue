<template>
  <div class="ad-img side-margin-bottom-20">
    <img :src="adimg.img" alt="">
  </div>
</template>

<script>
export default {
  props: {
    adimg: {
      type: Object,
      default: function() {
          return {img: "https://images.gitbook.cn/FoyiuT7b9raIvxPKhVPpjbVKFDDj?imageView2"}
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.ad-img {
    width: 100%;
    img {
        width: 100%;
    }
}
</style>