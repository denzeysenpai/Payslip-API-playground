<script lang="ts">
  import Router from "svelte-spa-router";
  import routes from "./routes";
  import NavSection from "./components/NavSection/NavSection.svelte";
  import Header from "./components/Header/Header.svelte";
</script>

<main>
  <Header></Header>
  <div class="main">
    <NavSection></NavSection>
    <Router {routes} />
  </div>
</main>

<style>
  main {
    padding: 0px;
    width: 100%;
    height: 100vh;
    margin: 0px;
    overflow: hidden;
    position: relative;
  }
  .main {
    padding: 0px;
    width: 100%;
    height: 100%;
    margin: 0px;
    display: flex;
    overflow: hidden;
  }
</style>
