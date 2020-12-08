<template>
  <div>
    <div class="main">
      <div class="content">
        <div class="detail">
          <vue-markdown>i am a ~~tast~~ **test**.</vue-markdown>
        </div>
      </div>
      <div class="side">
        <Adimg />
        <HotVquan />
      </div>
    </div>
  </div>
</template>

<script>
import VueMarkdown from "vue-markdown"

export default {
  data() {
    return {
      current: 3,
    };
  },
  components: {
    VueMarkdown
  },
  methods: {
    onChange(current) {
      this.current = current;
    },
  },
}
</script>

<style lang="scss">
  .content{
    width: 70%;
    img{
      max-width: 100%!important;
    }
  }
</style>
