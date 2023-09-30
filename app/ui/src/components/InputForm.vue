<template>
  <div class="mb-3">
    <input v-model="localUrl" class="form-control" placeholder="Enter URL" />
    <button @click="emitProcessData" class="btn btn-primary mt-3">Process Data</button>
    <div v-if="error" class="mt-2 text-danger">{{ error }}</div>
  </div>
</template>

<script>
import { ref } from "vue";

export default {
  props: ["url"],
  setup(props, { emit }) {
    const localUrl = ref(props.url);
    const error = ref(""); 

    const emitProcessData = () => {
      if (localUrl.value.trim() === "") {
        error.value = "Please enter a URL.";
      } else {
        error.value = ""; 
        emit("processData", localUrl.value);
      }
    };

    return {
      localUrl,
      emitProcessData,
      error  // Exposing error ref
    };
  }
};
</script>
