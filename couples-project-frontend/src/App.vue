<template>
  <div class="min-h-screen bg-gray-100">
    <!-- Navbar -->
    <nav class="bg-white shadow p-4 flex justify-between">
      <div class="flex space-x-4">
        <button class="px-3 py-2 rounded bg-blue-500 text-white hover:bg-blue-600">Home</button>
        <button class="px-3 py-2 rounded bg-green-500 text-white hover:bg-green-600">Gallery</button>
        <button class="px-3 py-2 rounded bg-purple-500 text-white hover:bg-purple-600">Profile</button>
      </div>
    </nav>

    <!-- Hero Section -->
    <header class="bg-white shadow my-8 p-8 text-center">
      <h1 class="text-4xl font-bold text-gray-800">Welcome to Your Journey Together</h1>
      <p class="mt-4 text-gray-600">Track your special moments, share memories, and celebrate your time together!</p>
      <button class="mt-6 px-6 py-3 bg-pink-500 text-white font-bold rounded-full hover:bg-pink-600">Get Started</button>
    </header>

    <!-- Gallery Section -->
    <section class="container mx-auto px-4">
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
        <div v-for="n in 6" :key="n" class="bg-white shadow-md rounded-lg overflow-hidden">
          <img class="w-full h-48 object-cover" :src="`https://via.placeholder.com/300?text=Photo+${n}`" alt="Placeholder Image">
          <div class="p-4">
            <h2 class="text-xl font-bold text-gray-800">Moment {{n}}</h2>
            <p class="text-gray-600 mt-2">This is a placeholder description for a memorable moment.</p>
            <button class="mt-4 px-4 py-2 bg-red-500 text-white rounded-full hover:bg-red-600">
              <span>&hearts;</span> Like
            </button>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'App',
}
</script>

<style>
/* Global reset e estilo base */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Inter', sans-serif;
  background-color: #f0f4f8;
  color: #333;
}

/* Navbar */
nav {
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

nav .nav-button {
  display: inline-block;
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  color: #fff;
  margin-right: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.nav-button-home {
  background-color: #007bff;
}

.nav-button-home:hover {
  background-color: #0056b3;
}

.nav-button-gallery {
  background-color: #28a745;
}

.nav-button-gallery:hover {
  background-color: #218838;
}

.nav-button-profile {
  background-color: #6f42c1;
}

.nav-button-profile:hover {
  background-color: #5a32a3;
}

/* Hero section */
header {
  text-align: center;
  padding: 4rem 0;
  background-color: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

header h1 {
  font-size: 3rem;
  color: #333;
}

header p {
  font-size: 1.25rem;
  color: #666;
  margin-top: 1rem;
}

header .cta-button {
  margin-top: 2rem;
  padding: 0.75rem 2rem;
  background-color: #ff5a5f;
  color: #fff;
  font-weight: bold;
  border-radius: 9999px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.cta-button:hover {
  background-color: #ff3b3b;
}

/* Gallery Section */
section {
  margin: 2rem auto;
  padding: 0 1.5rem;
  max-width: 1200px;
}

.gallery-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
}

.gallery-card {
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  transition: transform 0.3s ease;
}

.gallery-card:hover {
  transform: translateY(-5px);
}

.gallery-card img {
  width: 100%;
  height: 200px;
  object-fit: cover;
}

.gallery-card .content {
  padding: 1.5rem;
}

.gallery-card h2 {
  font-size: 1.5rem;
  color: #333;
}

.gallery-card p {
  margin-top: 0.5rem;
  font-size: 1rem;
  color: #666;
}

.like-button {
  display: inline-block;
  margin-top: 1rem;
  padding: 0.5rem 1rem;
  background-color: #e74c3c;
  color: #fff;
  border-radius: 9999px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.like-button:hover {
  background-color: #c0392b;
}

</style>
