import { createStore } from 'vuex'

export default createStore({
  state: {
    token: "",
    jobId: "",
  },
  getters: {
  },
  mutations: {
    login(state, token) {
      state.token = token;
      state.jobId = ""
      localStorage.setItem("token", token);
    },
    logout(state) {
      state.token = ""
      state.jobId = ""
      localStorage.removeItem("token")
      localStorage.removeItem("jobId")
    },
    initializationStore(state) {
      let token = localStorage.getItem("token")
      if (token) {
        state.token = token
      }
    },
    addJob(state, id) {
      state.jobId = id
      localStorage.setItem("jobId", id)
    },
    removeJob(state) {
      state.jobId = ""
      localStorage.removeItem("jobId")
    }
  },
  actions: {
  },
  modules: {
  }
})
