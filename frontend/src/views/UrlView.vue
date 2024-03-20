<template>
  <div class="container my-5">
    <form id="form" @submit.prevent="Submit" @click="Validate" @keypress="Validate">
      <div class="mb-3 row justify-content-center">
        <div class="col-1 col-form-label">
          <label for="url">URL</label>
        </div>
        <div class="col-6">
          <input type="url" v-model="form.url" class="form-control" id="url" autocomplete="off" required>
          <div class="invalid-feedback">
            Please copy and paste a link
          </div>
        </div>
        <div class="col-1">
          <button type="submit" class="btn btn-primary">Generate</button>
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
        url: "",
      }
    }
  },
  methods: {
    Submit(){
      axios({
        method: "POST",
        url: "http://127.0.0.1:8080/v1/app/url",
        headers: {
          "Authorization": this.$store.state.token,
        },
        data: {
          payload: this.form.url,
        }
      })
      .then(response => {
        this.$store.commit("addJob", response.data?.id)
        this.$router.push("/result")
      })
      .catch(error => {
        this.$notify({type: "error", title: "Oopsie", text: error?.response.data?.message || "Oops!"});
      })
    },
    Validate() {
      document.getElementById("form").setAttribute("class", "was-validated")
    },
  }
}
</script>