<script>
  import Sidebar from "../components/common/Sidebar.svelte";
  import Topbar from "../components/common/Topbar.svelte";
</script>

<div class="flex h-screen overflow-hidden">
  <Sidebar />

  <main class="flex-1 flex flex-col min-w-0 overflow-hidden">
    <Topbar />

    <div
      class="flex-1 overflow-y-auto p-8 bg-background-light dark:bg-background-dark/50"
    >
      <slot />
    </div>
  </main>
</div>
