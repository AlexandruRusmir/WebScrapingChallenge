<template>
    <div>
        <label for="tagSelector">Select tags:</label>
        <br>
        <select id="tagSelector" v-model="selectedTags" multiple class="custom-select">
            <option v-for="tag in availableTags" :key="tag" :value="tag">{{ tag }}</option>
        </select>
    </div>
</template>

<script>
    import { ref, watch, onMounted } from "vue";
    import { fetchAvailableTags } from "@/api";

    export default {
        setup(_, { emit }) {
            const availableTags = ref([]);
            const selectedTags = ref([]);

            const fetchTags = async () => {
                const tags = await fetchAvailableTags();
                availableTags.value = tags;
            };

            onMounted(fetchTags);

            watch(selectedTags, (newVal) => {
                emit("update:selectedTags", newVal);
            });

            return {
                availableTags,
                selectedTags
            };
        }
    };
</script>
<style scoped>
    .custom-select {
        height: 150px; /* Adjust to your preference */
        overflow-y: auto; /* Show scrollbar when the number of options exceeds the visible area */
    }
</style>