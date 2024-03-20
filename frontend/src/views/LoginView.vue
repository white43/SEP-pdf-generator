<template>
  <div class="container my-5">
    <form id="form" @submit.prevent="Submit" @click="Validate" @keypress="Validate">
      <div class="mb-3 row justify-content-center">
        <label for="email" class="col-1 col-form-label">Email</label>
        <div class="col-3">
          <input type="email" v-model="form.email" class="form-control" id="email" autocomplete="off" required>
          <div class="invalid-feedback">
            Please enter email
          </div>
        </div>
      </div>
      <div class="mb-3 row justify-content-center">
        <label for="password" class="col-1 col-form-label">Password</label>
        <div class="col-3">
          <input type="password" v-model="form.password" class="form-control" id="password" required>
          <div class="invalid-feedback">
            Please enter password
          </div>
        </div>
      </div>
      <div class="mb-3 row justify-content-center">
        <div class="col-3 offset-1">
          <button type="submit" class="btn btn-primary">Login</button>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data(){
    return {
      form: {
        email: "",
        password: "",
      }
    }
  },
  methods: {
    Submit(){
      axios({
        method: "POST",
        url: "http://127.0.0.1:8080/v1/user/login",
        data: {
          email: this.form.email,
          password: this.form.password,
        }
      })
      .then(response => {
        this.$store.commit("login", response.data?.access_token)
        this.$router.push("/balance")
      })
      .catch(error => {
        this.$notify({type: "error", title: "Oopsie", text: error?.response.data?.message || "Oops!"});
      });
    },
    Validate() {

      document.getElementById("form").setAttribute("class", "was-validated")
    },
  }
}
</script>