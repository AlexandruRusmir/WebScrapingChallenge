<template>
    <div class="mb-4">
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
        height: 25rem;
        overflow-y: auto;
        width: 20rem;
    }
</style>