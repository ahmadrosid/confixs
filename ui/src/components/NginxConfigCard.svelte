<script lang="ts">
  import { createQuery, createMutation } from "@tanstack/svelte-query";
  import RetroOutline from "./ui/RetroOutline.svelte";
  import Button from "./ui/Button.svelte";
  import PlusIcon from "lucide-svelte/icons/plus";
  import CodeEditor from "./CodeEditor.svelte";
  import Dialog from "./ui/Dialog.svelte";

  let dialogOpen = false;
  let filePath = "";

  const fetchConfigFiles = async () => {
    const response = await fetch("/api/config/list");
    if (!response.ok) {
      throw new Error("Failed to fetch NGINX configs");
    }
    const data = await response.json();
    return data.data;
  };

  const fetchConfigDetail = async ({ path }: { path: string }) => {
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
    return data;
  };

  const configFilesQuery = createQuery({
    queryKey: ["configFiles"],
    queryFn: fetchConfigFiles,
  });

  const configDetailMutation = createMutation({
    mutationKey: ["configDetail"],
    mutationFn: fetchConfigDetail,
    onSuccess: (data) => {
      console.log({data});
      filePath = data.path;
      dialogOpen = true;
    },
  });
</script>

<RetroOutline className="w-full" childClassName="p-0">
  <div class="border-b border-gray-800 p-3">
    <h2 class="text-xl font-bold">Nginx Config</h2>
  </div>
  <div class="px-4">
    <ul class="space-y-2 py-4 list-inside">
      {#if $configFilesQuery.isLoading}
        <li>Loading...</li>
      {:else if $configFilesQuery.isError}
        <li>Error: {$configFilesQuery.error.message}</li>
      {:else if $configFilesQuery.isSuccess}

      <li class="flex items-center justify-between pb-1">
        <h3 class="font-semibold">/etc/nginx/nginx.conf</h3>
        <Button
          size="sm"
          variant="warning"
          on:click={() =>{
            dialogOpen = true;
            $configDetailMutation.mutate({ path: "/etc/nginx/nginx.conf" });
          }}>Edit</Button
        >
      </li>
        {#each $configFilesQuery.data as configFile}
          <li class="flex items-center justify-between pb-1">
            <h3 class="font-semibold">{configFile}</h3>
            <Button
              size="sm"
              variant="warning"
              on:click={() =>{
                dialogOpen = true;
                $configDetailMutation.mutate({ path: configFile });
              }}>Edit</Button
            >
          </li>
        {/each}
      {/if}
    </ul>

    <Dialog bind:open={dialogOpen} title={`Edit Config: ${filePath}`}>
      <div class="border border-gray-400">
        <div class="overflow-y-auto h-full min-h-[400px] max-h-[400px]">
          {#if $configDetailMutation.isError}
            <div class="pb-4">Error: {$configDetailMutation.error.message}</div>
          {:else if $configDetailMutation.isSuccess}
            <CodeEditor value={$configDetailMutation.data.content} />
          {/if}
        </div>
      </div>
      <div class="pt-4">
        <div class="flex gap-2 justify-end">
          <Button on:click={() => (dialogOpen = false)} variant="secondary">Cancel</Button>
          <Button>Save</Button>
        </div>
      </div>
    </Dialog>
  </div>
</RetroOutline>
