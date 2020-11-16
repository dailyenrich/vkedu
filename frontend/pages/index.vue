<template>
  <div>
    <div class="main">
      <div class="content">
        <div class="banner">
          <Banner />
        </div>
        <List />
      </div>
      <div class="side">
        1
      </div>
    </div>
  </div>
</template>

<script>
export default {}
</script>

<style lang="scss">
.banner {
  width: 100%;
}

.main {
  display: flex;
  justify-content: space-between;
  .content {
    width: 70%;
  }
  .side {
    width: 29%;
    padding: 20px;
    border: 1px solid $main-border-color;
  }
}
</style>
