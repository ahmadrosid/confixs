<script lang="ts">
  import { createQuery, createMutation } from "@tanstack/svelte-query";
  import RetroOutline from "./ui/RetroOutline.svelte";
  import Button from "./ui/Button.svelte";
  import CodeEditor from "./CodeEditor.svelte";
  import Dialog from "./ui/Dialog.svelte";
  import Loading from "./ui/Loading.svelte";

  let dialogOpen = false;
  let filePath = "";

  const closeDialog = () => {
    dialogOpen = false;
  }

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
    mutationFn: fetchConfigDetail
  });

  $: if ($configDetailMutation.isSuccess) {
    let data = $configDetailMutation.data;
    filePath = data.path;
    dialogOpen = true;
  }
</script>

<RetroOutline className="w-full" childClassName="p-0">
  <div class="border-b border-gray-800 p-3">
    <h2 class="text-xl font-bold">Nginx Config</h2>
  </div>
  <div class="px-4">
    <ul class="space-y-2 py-4 list-inside">
      {#if $configFilesQuery.isLoading}
        <div class="p-4 flex justify-center">
          <Loading />
        </div>
      {:else if $configFilesQuery.isError}
        <li class="text-rose-500">Error: {$configFilesQuery.error.message}</li>
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

    <Dialog bind:open={dialogOpen} on:close={closeDialog} title={`Edit Config: ${filePath}`}>
      <div class="border border-gray-400 grid">
        <div class="overflow-y-auto h-full min-h-[400px] max-h-[400px]">
          {#if $configDetailMutation.isPending}
            <div class="w-full h-full grid place-content-center">
              <Loading />
            </div>
          {:else if $configDetailMutation.isError}
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
