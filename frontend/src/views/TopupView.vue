<template>
  <div class="container my-5">
    <form id="form" @submit.prevent="Submit" @click="Validate" @keypress="Validate">
      <div class="mb-3 row justify-content-center">
        <label for="amount" class="col-1 col-form-label">Amount</label>
        <div class="col-3">
          <input type="number" v-model="form.amount" class="form-control" id="amount" autocomplete="off" required min="1">
          <div class="invalid-feedback">
            Please enter the desired amount to top up
          </div>
        </div>
      </div>
      <div class="mb-3 row justify-content-center">
        <div class="col-3 offset-1">
          <button type="submit" class="btn btn-primary">Pay {{ this.form.amount ? "$" + this.form.amount : "" }}</button>
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
        amount: ""
      }
    }
  },
  methods: {
    Submit(){
      let token = this.$store.state.token;

      if (token) {
        axios({
          method: "POST",
          url: "http://127.0.0.1:8080/v1/user/topup",
          data: {
            amount: this.form.amount.toString(),
          },
          headers: {
            "Authorization": token,
          },
        })
        .then(response => {
          this.form.amount = ""
          document.getElementById("form").setAttribute("class", "")
          this.$notify({type: "success", title: "Success", text: response.data?.message});
        })
        .catch(error => {
          this.$notify({type: "error", title: "Oopsie", text: error?.response.data?.message || "Oops!"});
        });
      }
    },
    Validate() {
      document.getElementById("form").setAttribute("class", "was-validated")
    },
  }
}
</script>