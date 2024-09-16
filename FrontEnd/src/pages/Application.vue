<script setup>
import "./../styles/App.css";
import Layout from "./../components/Layout.vue";
import Cards from "./../components/Cards.vue";
import { ref, computed } from "vue";

const cards = [
  {
    imagesrc: "image/exp-appimage.jpg",
    title: "Application 1",
    description: "Petite Description",
    link: "/application-1"
  },
  {
    imagesrc: "image/exp-appimage.jpg",
    title: "Application 2",
    description: "Petite Description",
    link: "/application-2"
  },
  {
    imagesrc: "image/exp-appimage.jpg",
    title: "Application 3",
    description: "Petite Description",
    link: "/application-3"
  },
  {
    imagesrc: "image/exp-appimage2.jpg",
    title: "Application 4",
    description: "Petite Description",
    link: "/application-4"
  },
  {
    imagesrc: "image/exp-appimage2.jpg",
    title: "application 5",
    description: "Petite Description",
    link: "/application-5"
  },
  {
    imagesrc: "image/exp-appimage2.jpg",
    title: "Application 6",
    description: "Petite Description",
    link: "/application-6"
  },
];

const currentPage = ref(0);
const itemsPerPage = 3;

const paginatedCards = computed(() => {
  const start = currentPage.value * itemsPerPage;
  const end = start + itemsPerPage;
  return cards.slice(start, end);
});
const totalPages = computed(() => Math.ceil(cards.length/itemsPerPage))
const next = () => {
  if (currentPage.value < totalPages.value - 1) {
    currentPage.value++;
  }
};

const previous = () => {
  if (currentPage.value > 0) {
    currentPage.value--;
  }
};

const goToPage = (page) => {
  currentPage.value = page;
};

</script>

<template>
  <Layout>
    <template #link> </template>
    <template #content-header>
      <h1>Bienvenue sur l'environnement digital TRIGO FRANCE </h1>
    </template>

    <template #content >
    <div class="main">
      <span> Lien direct vers les outils digitaux </span>
      <div class="Cards">
     
        <Cards
          v-for="(card, index) in paginatedCards"
          :key="index"
          :imagesrc="card.imagesrc"
          :title="card.title"
          :description="card.description"
          :link="card.link"
        />
  
      </div>
      <div class="page">
        <button @click="previous" :disabled="currentPage === 0">Précédent</button>
        <button 
          v-for="page in totalPages"
          :key="page"
          @click="goToPage(page - 1)"
          :class="{ active: currentPage === page - 1 }"
        >
          {{ page }}
        </button>
        <button @click="next" :disabled="currentPage === Math.ceil(cards.length / itemsPerPage) - 1">Suivant</button>
        

      </div>
      </div>
    
    
    </template>    
  </Layout>
</template>

