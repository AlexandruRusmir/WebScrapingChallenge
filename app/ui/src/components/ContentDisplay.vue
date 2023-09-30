<template>
    <div v-if="contentData.wordCount" class="alert">
        <h3>
            Word Count: {{ contentData.wordCount }}
        </h3>
    </div>
    <div v-if="contentData.content">
        <h3>
            Tags:
        </h3>
        <div v-for="(item, key) in contentData.content" :key="key">
            <div v-if="item && item.length">
            <h4 class="mt-4">{{ key.toUpperCase() }}</h4>
            <div v-if="key === 'a' || key === 'img'">
                <ul>
                <li v-for="(subItem, index) in item" :key="index">
                    <div v-if="key === 'a'">
                    <a :href="subItem.href">{{ subItem.text ?? 'text' }}</a>
                    </div>
                    <div v-if="key === 'img'">
                    Image #{{ index + 1 }}
                    <ul>
                        <li> src: {{ subItem.src ?? null }} </li>
                        <li> alt: {{ subItem.alt ?? null }} </li>
                    </ul>
                    </div>
                </li>
                </ul>
            </div>
            <div v-else>
                <ul>
                <li v-for="(text, index) in item" :key="index">{{ text }}</li>
                </ul>
            </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        props: ["contentData"]
    };
</script>
  