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
        <label for="first_name" class="col-1 col-form-label">First Name</label>
        <div class="col-3">
          <input type="text" v-model="form.first_name" class="form-control" id="first_name" autocomplete="off" required>
          <div class="invalid-feedback">
            Please enter first name
          </div>
        </div>
      </div>
      <div class="mb-3 row justify-content-center">
        <label for="last_name" class="col-1 col-form-label">Last Name</label>
        <div class="col-3">
          <input type="text" v-model="form.last_name" class="form-control" id="last_name" autocomplete="off" required>
          <div class="invalid-feedback">
            Please enter last name
          </div>
        </div>
      </div>
      <div class="mb-3 row justify-content-center">
        <div class="col-3 offset-1">
          <button type="submit" class="btn btn-primary">Register</button>
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
        first_name: "",
        last_name: "",
      }
    }
  },
  methods: {
    Submit(){
      axios({
        method: "POST",
        url: "http://127.0.0.1:8080/v1/user/register",
        data: {
          email: this.form.email,
          first_name: this.form.first_name,
          last_name: this.form.last_name,
        }
      })
      .then(response => {
        if (response.data?.code === 201) {
          this.$notify({type: "success", title: "Success", text: response.data?.message || "Oops!"});
          this.$router.push("/login")
        }
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