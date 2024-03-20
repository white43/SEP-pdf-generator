<template v-if="this.$store.state.jobId">
  <div id="container" class="container my-5"></div>
</template>

<script>
import axios from 'axios'

export default {
  methods: {
    initialize() {
      let id = this.$store.state.jobId;
      let token = this.$store.state.token;

      if (id && token) {
        let interval = setInterval(function () {
          axios({
            method: "GET",
            url: "http://127.0.0.1:8080/v1/app/result/" + id + "/",
            headers: {
              "Authorization": token,
            }
          })
          .then(response => {
            clearInterval(interval)

            let result = response.data?.result;

            if (result) {
              document.getElementById("alert")?.remove()

              let container = document.getElementById("container");
              let fragment = document.createDocumentFragment();

              let iframe = document.createElement("iframe")
              iframe.src = "data:application/pdf;base64," + result;
              iframe.style.width = "100%"
              iframe.style.height = "80vh"

              fragment.appendChild(iframe)
              container.appendChild(fragment)
            }
          })
          .catch(error => {
            let div = document.getElementById("alert")

            if (div) {
              div.innerText = error.response.data?.message || "Oops!"
            } else {
              let container = document.getElementById("container");
              let fragment = document.createDocumentFragment();

              let div = document.createElement("div")
              div.id = "alert"
              div.className = "alert alert-warning"
              div.innerText = error.response.data?.message || "Oops!"

              fragment.appendChild(div)
              container.appendChild(fragment)
            }
          })
        }, 1000);
      }
    },
  },
  beforeMount() {
    this.initialize()
  }
}
</script>