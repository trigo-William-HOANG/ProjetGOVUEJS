<script setup>
import '../styles/Layout.css'
import { ref, onMounted } from "vue";

const currentPath = ref("");
const hideMenu = ref(false);

// Set current path when component is mounted
onMounted(() => {
  currentPath.value = window.location.pathname;
  components: { loader: true }
});

// Function to check if a route is active
const isActive = (path) => {
  return currentPath.value === path ? "active" : "";
};

</script>

<template>
  <div class="layout">
<div class="sidebar" :class="{ 'close': hideMenu }">
      <div class="logo">
    
        <button @click="hideMenu = !hideMenu">
          <i class="bx bx-menu"></i>
        </button>
        <img src="/image/trigo-logo.png" />
      </div>
     
      <ul class="side-menu">
        
        <li :class="isActive('/')">
          <a href="/"> <i class="bx bx-home-alt-2"></i> Accueil </a>
        </li>
        <li :class="isActive('/commande')">
          <a href="/commande"> <i class="bx bx-package"></i> Mes commandes </a>
        </li>
        <li :class="isActive('/dashboard')">
          <a href="/dashboard"> <i class="bx bx-desktop"></i> Mes dashboards </a>
        </li>
        <li :class="isActive('/mission')">
          <a href="/mission"> <i class="bx bx-spreadsheet"></i> Mes missions </a>
        </li>
        <li :class="isActive('/document')">
          <a href="/document"> <i class="bx bx-file"></i> Mes documents </a>
        </li>
        <li :class="isActive('/application')">
          <a href="/application">
            <i class="bx bxs-briefcase"></i> Applications métiers
          </a>
        </li>
        <li :class="isActive('/rapport')">
          <a href="/rapport"> <i class="bx bx-compass"></i> Rapports internes </a>
        </li>
        <li :class="isActive('/retour')">
          <a href="/retour"> <i class="bx bx-conversation"></i> Mes retours </a>
        </li>
        <li :class="isActive('/instruction')">
          <a href="/instruction">
            <i class="bx bx-chalkboard"></i> Mes Instructions de travail
          </a>
        </li>
        <li :class="isActive('/contact')">
          <a href="/contact">
            <i class="bx bx-message-square-dots"></i> Nous contacter
          </a>
        </li>
        <li :class="isActive('/annonce')">
          <a href="/annonce"> <i class="bx bxs-megaphone"></i> Annonces </a>
        </li>
      </ul>
    </div>

    <div class="navbar content">
      <nav>
        <div class="slot-link">
          <slot name="link"></slot>
        </div>
        <a href="#"><i class="bx bx-help-circle"></i></a>
        <a href="#" class="circle-user-dropdown">WH</a>
      </nav>
    </div>

    <main>
      <div class="content-header">
        <slot name="content-header"></slot>
        
      </div>
      <div class="content-main">
        <slot name="content"></slot>
      </div>
    </main>

    <footer>
      <slot name="footer"></slot>
    </footer>
  </div>
</template>
