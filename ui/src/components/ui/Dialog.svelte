<script lang="ts">
  import { XIcon } from "lucide-svelte";
  import RetroOutline from "./RetroOutline.svelte";
  import { onMount, onDestroy, createEventDispatcher } from 'svelte';

  export let open = false;
  export let title = "Dialog";

  const dispatch = createEventDispatcher();

  function close() {
    dispatch('close');
  }

  function handleKeydown(event: KeyboardEvent) {
    if (event.key === 'Escape' && open) {
      close();
    }
  }

  onMount(() => {
    document.addEventListener('keydown', handleKeydown);
  });

  onDestroy(() => {
    document.removeEventListener('keydown', handleKeydown);
  });

  function clickOutside(node: HTMLElement) {
    const handleClick = (event: MouseEvent) => {
      if (
        node &&
        !node.contains(event.target as Node) &&
        !event.defaultPrevented
      ) {
        node.dispatchEvent(new CustomEvent("outclick"));
      }
    };

    document.addEventListener("click", handleClick, true);
    return {
      destroy() {
        document.removeEventListener("click", handleClick, true);
      },
    };
  }

  function handleOutclick() {
    if (open) {
      close();
    }
  }
</script>

{#if open}
  <div
    class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-20"
  >
    <div
      class="bg-transparent p-4 max-w-3xl w-full"
      use:clickOutside
      on:outclick={handleOutclick}
    >
      <RetroOutline className="w-full">
        <div class="p-4">
          <div class="flex items-center justify-between mb-4">
            <h2 class="text-2xl font-bold">{title}</h2>
            <button class="hover:text-gray-800" on:click={close}>
              <XIcon />
            </button>
          </div>
          <slot></slot>
        </div>
      </RetroOutline>
    </div>
  </div>
{/if}
