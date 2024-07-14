<script lang="ts">
  import XIcon from "lucide-svelte/icons/x";
  import RetroOutline from "../ui/RetroOutline.svelte";
  import { getToastState } from "../../lib/toast";

  let toast = getToastState();
  let demoVisible = false;

  function clickOutside(node: HTMLElement) {
    const handleClick = (event: MouseEvent) => {
      if (node && !node.contains(event.target as Node) && !event.defaultPrevented) {
        node.dispatchEvent(new CustomEvent('outclick'));
      }
    };

    document.addEventListener('click', handleClick, true);
    return {
      destroy() {
        document.removeEventListener('click', handleClick, true);
      }
    };
  }

  function toggleDemo() {
    demoVisible = !demoVisible;
  }

  function handleOutclick() {
    demoVisible = false;
  }
</script>

<div class="container mx-auto max-w-6xl px-4 relative">
  <main>
    <div class="py-16">
      <h1 class="text-6xl font-bold text-center text-balance">
        Simplify Your Nginx Configuration Management!
      </h1>
      <p class="text-xl text-center mt-6">
        Nginx Made Easy: Configure, Secure, and Monitor with Confixs
      </p>
      <div class="text-center pt-12">
        <div class="relative inline-block mr-4">
          <button
            on:click={toggleDemo}
            class="bg-black text-white px-4 py-2 hover:underline relative z-10"
          >
            See demo
          </button>
          <div
            class="absolute left-1 top-1 -right-1 -bottom-1 bg-white border-2 border-gray-800"
            aria-hidden="true"
          ></div>
        </div>
        <button
          on:click={() => toast.add("Documentation", "Coming soon", "info")}
          class="hover:underline">Documentation</button
        >
      </div>
    </div>

    <div class="pb-12">
      <RetroOutline>
        <img src="/demo.png" alt="Intersight" />
      </RetroOutline>
    </div>
  </main>
</div>

{#if demoVisible}
  <div
    class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-20"
  >
    <div class="bg-transparent p-4 max-w-3xl w-full" use:clickOutside on:outclick={handleOutclick}>
      <RetroOutline className="w-full">
        <div class="p-6">
          <div class="flex items-center justify-between mb-4">
            <h2 class="text-2xl font-bold">Confixs Demo</h2>
            <button class="hover:text-gray-800" on:click={toggleDemo}>
              <XIcon />
            </button>
          </div>
          <div>Iframe content goes here!</div>
        </div>
      </RetroOutline>
    </div>
  </div>
{/if}