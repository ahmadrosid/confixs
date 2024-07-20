<script lang="ts">
  import { createQuery, createMutation } from "@tanstack/svelte-query";
  import RetroOutline from "./ui/RetroOutline.svelte";
  import Button from "./ui/Button.svelte";
  import PlusIcon from "lucide-svelte/icons/plus";
  import CodeEditor from "./CodeEditor.svelte";
  import Dialog from "./ui/Dialog.svelte";

  let dialogOpen = false;
  let fileContent = "";

  const fetchConfigFiles = async () => {
    const response = await fetch("/api/config/list");
    if (!response.ok) {
      throw new Error("Failed to fetch NGINX configs");
    }
    const data = await response.json();
    return data.data;
  };

  const fetchConfigDetail = async ({ path }: {path: string}) => {
    const response = await fetch("/api/config/get", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        filePath: path,
      }),
    });
    if (!response.ok) {
      throw new Error("Failed to fetch NGINX config detail.");
    }
    const data = await response.json();
    return data.data;
  };

  const configFilesQuery = createQuery({
    queryKey: ["configFiles"],
    queryFn: fetchConfigFiles,
  });

  const configDetailMutation = createMutation({
    mutationKey: ["configDetail"],
    mutationFn: fetchConfigDetail,
    onSuccess: (data) => {
      console.log(data);
      fileContent = data.data;
      dialogOpen = true;
    },
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
            <Button size="sm" variant="warning" on:click={() => $configDetailMutation.mutate({ path: configFile })}>Edit</Button>
          </li>
        {/each}
      {/if}
    </ul>
    
    {#if $configDetailMutation.isError}
      <div class="pb-4">Error: {$configDetailMutation.error.message}</div>
    {:else if $configDetailMutation.isSuccess}
      <Dialog bind:open={dialogOpen} title="Edit Config">
        <div class="border">
          <div class="overflow-y-auto h-full max-h-[400px]">
            <CodeEditor value={$configDetailMutation.data} />
          </div>
        </div>
      </Dialog>
    {/if}

  </div>
</RetroOutline>
