<template>
  <nav class="btn-group btn-group-lg d-block text-center">
    <router-link v-if="this.$store.state.token" to="/balance" class="btn btn-light">Balance</router-link>
    <router-link v-if="this.$store.state.token" to="/top-up" class="btn btn-light">Top Up</router-link>
    <router-link v-if="this.$store.state.token" to="/token" class="btn btn-light">API Token</router-link>
    <router-link v-if="this.$store.state.token" to="/html" class="btn btn-light">From HTML</router-link>
    <router-link v-if="this.$store.state.token" to="/url" class="btn btn-light">From URL</router-link>
    <router-link v-if="this.$store.state.token && this.$store.state.jobId" to="/result" class="btn btn-light">Result</router-link>
    <router-link v-if="!this.$store.state.token" to="/registration" class="btn btn-light">Registration</router-link>
    <router-link v-if="!this.$store.state.token" to="/login" class="btn btn-light">Login</router-link>
    <a v-if="this.$store.state.token" @click="logout" class="btn btn-light">Logout</a>
  </nav>
  <notifications duration=10 />
  <router-view/>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

body {
  background-color: #fbfbfb !important;
}

nav {
  padding: 30px;
}

nav a.router-link-exact-active {
  color: #42b983;
}
.vue-notification-group {
  right: 2.5vw !important;
  top: 5vh !important;
}
.vue-notification {
  font-size: 15px !important;
}
</style>
<script>
export default {
  methods: {
    logout(){
      this.$store.commit("logout")
      this.$router.push("/")
    }
  },
  mounted() {
    this.$store.commit("initializationStore")
  }
}
</script>
