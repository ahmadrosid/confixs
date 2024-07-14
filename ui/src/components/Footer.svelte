<script lang="ts">
  import Twitter from "./BrandIcon/Twitter.svelte";
  import Instagram from "./BrandIcon/Instagram.svelte";
  import Github from "./BrandIcon/Github.svelte";

  const currentYear = new Date().getFullYear();
  const socialIcons = [
    { icon: Twitter, href: "#" },
    { icon: Instagram, href: "#" },
    { icon: Github, href: "#" },
  ];
</script>

<footer class="bg-black text-gray-300">
  <div class="container max-w-6xl mx-auto px-4 py-12">

    <div
      class="flex flex-col md:flex-row justify-between items-center"
    >
      <div class="mb-4 md:mb-0">
        <p>&copy; {currentYear} Confixs. All rights reserved.</p>
      </div>
      <div class="flex space-x-4">
        {#each socialIcons as { icon, href }}
          <a
            {href}
            class="text-gray-400 hover:text-white transition-colors duration-300"
          >
            <svelte:component this={icon} class="w-5 h-5 text-white" />
          </a>
        {/each}
      </div>
    </div>
  </div>
</footer>
