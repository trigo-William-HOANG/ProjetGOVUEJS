<script setup>
import { ref, computed } from "vue";
import { defineAsyncComponent } from "vue";
import Home from "./pages/Home.vue";
import Application from "./pages/Application.vue";
import Commande from "./pages/Commande.vue";
import Dashboard from "./pages/Dashboard.vue";
import Mission from "./pages/Mission.vue";
import Document from "./pages/Document.vue";
import Rapport from "./pages/Rapport.vue";
import Retour from "./pages/Retour.vue";
import Instruction from "./pages/Instruction.vue";
import Contact from "./pages/Contact.vue";
import Annonce from "./pages/Annonce.vue";
import Application1 from "./pages/Application1.vue";
import Application2 from "./pages/Application2.vue";
import Application3 from "./pages/Application3.vue";

const routes = {
  "/": Home,
  "/application": Application,
  "/commande": Commande,
  "/dashboard": Dashboard,
  "/mission": Mission,
  "/document": Document,
  "/rapport": Rapport,
  "/retour": Retour,
  "/instruction": Instruction,
  "/contact": Contact,
  "/annonce": Annonce,
  "/application-1": Application1,
  "/application-2": Application2,
  "/application-3": Application3,
};

// Reactive variable for the current path
const currentPath = ref(window.location.pathname);

// Function to update currentPath on URL change
const updatePath = () => {
  currentPath.value = window.location.pathname;
};

// Listen for changes to the URL
window.addEventListener("popstate", updatePath);

// Compute the component to render based on the current path
const currentView = computed(() => {
  return routes[currentPath.value] || NotFound;
});
</script>

<template>
  <component :is="currentView" />
</template>
