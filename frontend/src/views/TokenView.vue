<template>
  <div class="container my-5">
    <div class="alert alert-info">Your API token: {{ token }}</div>
  </div>
</template>

<script>
export default {
  data(){
    return {
      token: this.$store.state.token,
    }
  }
}
</script>