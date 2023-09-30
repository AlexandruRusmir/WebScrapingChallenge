<template>
    <div v-if="shouldDisplay">
      <button @click="downloadResults" class="btn btn-secondary mt-3">Save Results as JSON</button>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      sentiment: {
        type: String,
        default: ""
      },
      contentData: {
        type: Object,
        default: () => ({})
      }
    },
  
    computed: {
      shouldDisplay() {
        return this.sentiment && (Object.keys(this.contentData).length > 0);
      }
    },
  
    methods: {
      downloadResults() {
        const dataStr = "data:text/json;charset=utf-8," + encodeURIComponent(JSON.stringify({
          sentiment: this.sentiment,
          contentData: this.contentData
        }));
        const downloadAnchorNode = document.createElement('a');
        downloadAnchorNode.setAttribute("href", dataStr);
        downloadAnchorNode.setAttribute("download", "api_results.json");
        document.body.appendChild(downloadAnchorNode);
        downloadAnchorNode.click();
        downloadAnchorNode.remove();
      }
    }
  };
  </script>
  