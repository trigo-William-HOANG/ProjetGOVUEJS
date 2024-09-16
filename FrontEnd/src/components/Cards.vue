<script setup>
import "../styles/Cards.css";

const props = defineProps({
  imagesrc: String,
  title: String,
  description: String,
  link: String,
});

const navigateTo = (url) => {
  window.location.href = url;
};
</script>

<template>
  <div class="Card" @click="navigateTo(props.link)">
    <div class="Card-image">
      <img :src="props.imagesrc" />
    </div>
    <div class="Card-title">
      <h2>{{ props.title }}</h2>
      <p>{{ props.description }}</p>
    </div>
  </div>
</template>
