<template>
  <div class="container my-5">
    <div class="alert alert-info">Your Balance: ${{ balance }}</div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data(){
    return {
      balance: 0,
    }
  },
  methods: {
    init() {
      let token = this.$store.state.token;

      if (token) {
        axios({
          method: "GET",
          url: "http://127.0.0.1:8080/v1/user/balance",
          headers: {
            "Authorization": token,
          },
        })
        .then(response => {
          this.balance = response.data?.balance;
        })
        .catch(error => {
          this.$notify({type: "error", title: "Oopsie", text: error?.response.data?.message || "Oops!"});
        });
      }
    },
  },
  beforeMount() {
    this.init()
  }
}
</script>