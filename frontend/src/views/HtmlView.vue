<template>
  <div class="container my-5">
    <form id="form" @submit.prevent="Submit" @click="Validate" @keypress="Validate">
      <div class="mb-3 row">
        <div class="col-8 offset-2">
          <label for="html" class="col-form-label">HTML</label>
          <textarea name="html" rows="5" v-model="form.html" class="form-control" id="html" required></textarea>
          <div><small class="text-muted">Copy and paste HTML contents here</small></div>
          <div class="invalid-feedback">
            Please copy and paste some HTML
          </div>
        </div>
      </div>
      <div class="mb-3 row">
        <div class="col-8 offset-2">
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
        html: "",
      }
    }
  },
  methods: {
    Submit(){
      axios({
        method: "POST",
        url: "http://127.0.0.1:8080/v1/app/html",
        headers: {
          "Authorization": this.$store.state.token,
        },
        data: {
          payload: this.form.html,
        }
      })
      .then(response => {
        this.$store.commit("addJob", response.data?.id)
        this.$router.push("/result")

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