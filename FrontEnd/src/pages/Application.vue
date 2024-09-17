<script setup>
import { ref, computed, onMounted, watch } from "vue";
import Layout from "./../components/Layout.vue";
import Cards from "./../components/Cards.vue";
import axios from 'axios'; 


const cards = ref([]);

const currentPage = ref(0);
const itemsPerPage = 3;

const fetchCards = async (page) => {
  try {
    const startId = page * itemsPerPage;
    const endId = startId + itemsPerPage;
    const fetchedCards = [];

    for (let id = startId; id < endId; id++) {
      const response = await axios.get(`http://localhost:8080/application-data/?Id=${id}`);
      if (response.data) {
        fetchedCards.push(response.data);
      }
    }

    cards.value = fetchedCards;
  } catch (error) {
    console.error('Error fetching cards:', error);
  }
};

watch(currentPage, (newPage) => {
  fetchCards(newPage);
});

onMounted(() => {
  fetchCards(currentPage.value);
});

const paginatedCards = computed(() => cards.value);

const totalPages = computed(() => Math.ceil(cards.value.length / itemsPerPage));

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
        <span>Lien direct vers les outils digitaux</span>
        <div class="Cards">
          <Cards
            v-for="(card, index) in paginatedCards"
            :key="index"
            :imagesrc="card.Image"
            :title="card.Title"
            :description="card.Description"
            :link="card.Link"
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
          <button @click="next" :disabled="currentPage === totalPages - 1">Suivant</button>
        </div>
      </div>
    </template>
  </Layout>
</template>
