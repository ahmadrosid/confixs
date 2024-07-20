<script lang="ts">
    import { createQuery } from '@tanstack/svelte-query';
    import RetroOutline from "./ui/RetroOutline.svelte";
    import Button from "./ui/Button.svelte";
    import PlusIcon from "lucide-svelte/icons/plus";
    import CodeEditor from './CodeEditor.svelte';
    import Dialog from './ui/Dialog.svelte';
    let dialogOpen = false;

    const fetchConfigFiles = async () => {
      const response = await fetch("/api/config/list");
      if (!response.ok) {
        throw new Error("Failed to fetch NGINX configs");
      }
      const data = await response.json();
      return data.data;
    };
  
    const configFilesQuery = createQuery({
      queryKey: ['configFiles'],
      queryFn: fetchConfigFiles,
    });
  </script>
  
  <RetroOutline className="w-full" childClassName="p-0">
    <div class="border-b border-gray-800 p-3">
      <div class="flex items-center justify-between">
        <h2 class="text-xl font-bold">Nginx Config</h2>
        <div>
          <Button className="inline-flex items-center gap-2 pl-3">
            <PlusIcon size="20" /> New Config
          </Button>
        </div>
      </div>
    </div>
    <div class="px-4">
      <ul class="space-y-2 py-4 list-inside">
        {#if $configFilesQuery.isLoading}
          <li>Loading...</li>
        {:else if $configFilesQuery.isError}
          <li>Error: {$configFilesQuery.error.message}</li>
        {:else if $configFilesQuery.isSuccess}
          {#each $configFilesQuery.data as configFile}
            <li class="flex items-center justify-between pb-1">
              <h3 class="font-semibold">{configFile}</h3>
              <Button size="sm" variant="warning">Edit</Button>
            </li>
          {/each}
        {/if}
      </ul>
      <button on:click={() => dialogOpen = true}>Edit Config</button>

      <Dialog bind:open={dialogOpen} title="Edit Config">
        <div class="border">
          <div class="overflow-y-auto h-full max-h-[400px]">
            <CodeEditor />
          </div>
        </div>
      </Dialog>
    </div>
  </RetroOutline>