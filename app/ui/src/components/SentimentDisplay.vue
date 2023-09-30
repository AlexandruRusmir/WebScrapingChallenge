<template>
    <div v-if="sentiment" class="alert">
        <h3>
            Sentiment: {{ sentiment }}
        </h3>
    </div>
</template>

<script>

export default {
    props: ["sentiment"],
};
</script>
