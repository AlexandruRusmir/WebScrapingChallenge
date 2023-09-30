<template>
  <div class="container mt-5">
    <InputForm :url="url" @processData="processData($event)" />
    <TagSelector :selectedTags="selectedTags" @update:selectedTags="updateSelectedTags" />
    <SentimentDisplay :sentiment="sentiment" />
    <ContentDisplay :contentData="content" />
  </div>
</template>

<script>
  import { ref } from "vue";
  import { scrapeContent, analyzeSentiment } from "./api.js";
  import InputForm from "./components/InputForm.vue";
  import SentimentDisplay from "./components/SentimentDisplay.vue";
  import ContentDisplay from "./components/ContentDisplay.vue";
  import TagSelector from "./components/TagSelector.vue";

  export default {
    components: {
      InputForm,
      SentimentDisplay,
      ContentDisplay,
      TagSelector
    },
    setup() {
      const url = ref("");
      const content = ref({});
      const sentiment = ref("");
      const selectedTags = ref([]);
      
      const processData = async (urlValue) => {
        const scrapedData = await scrapeContent(urlValue, selectedTags.value);
        content.value.content = scrapedData.content;
        content.value.wordCount = scrapedData.wordCount;

        const sentimentData = await analyzeSentiment(urlValue);
        sentiment.value = sentimentData.sentiment;
      };

      const updateSelectedTags = (tags) => {
        selectedTags.value = tags;
      };

      

      return {
        url,
        processData,
        content,
        sentiment,
        selectedTags,
        updateSelectedTags
      };
    }
  };
</script>
