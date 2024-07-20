<script lang="ts">
    import { createQuery } from '@tanstack/svelte-query';
    import RetroOutline from "./ui/RetroOutline.svelte";
    import Button from "./ui/Button.svelte";
    import PlusIcon from "lucide-svelte/icons/plus";
  
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
            <li>
              <h3 class="font-semibold">{configFile}</h3>
            </li>
          {/each}
        {/if}
      </ul>
    </div>
  </RetroOutline>