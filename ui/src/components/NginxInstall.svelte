<script lang="ts">
  import { onDestroy } from "svelte";
  import Button from "./ui/Button.svelte";
  import RetroOutline from "./ui/RetroOutline.svelte";

  let installationOutput: string[] = [];
  let isInstalling: boolean = false;
  let installSuccess: boolean = false;
  let error: string | null = null;

  let eventSource: EventSource | null = null;

  function startInstallation(): void {
    isInstalling = true;
    installationOutput = [];
    error = null;

    eventSource = new EventSource("/api/install/nginx");

    eventSource.onmessage = (event: MessageEvent) => {
      installationOutput = [...installationOutput, event.data];
    };

    eventSource.onerror = (err: Event) => {
      error = "An error occurred during installation";
      isInstalling = false;
      eventSource?.close();
    };

    eventSource.addEventListener("close", (e) => {
      const result = JSON.parse(e.data);
      if (result.status === "success") {
        installSuccess = true;
      }
      isInstalling = false;
      eventSource?.close();
    });
  }

  onDestroy(() => {
    if (eventSource) {
      eventSource.close();
    }
  });
</script>

<div class="space-y-4">
  <h1 class="font-bold text-xl">You have not installed Nginx yet.</h1>
  <p class="text-sm">Installing Nginx is required to run this application.</p>

  <Button disabled={isInstalling} on:click={startInstallation}>
    {#if !isInstalling}
      Install Nginx
    {:else}
      Installing Nginx...
    {/if}
  </Button>

  {#if error}
    <p class="error">{error}</p>
  {/if}

  {#if installationOutput.length > 0}
    <RetroOutline>
      <div class="h-80 p-2 overflow-y-auto bg-neutral-700 text-white">
        {#each installationOutput as line}
          <pre class="text-sm">{line}</pre>
        {/each}
      </div>
    </RetroOutline>
  {/if}


  {#if installSuccess}
  <div class="flex justify-end">
    <Button disabled={isInstalling} variant="success" on:click={() => window.location.href = "/dashboard"}>
        Continue
      </Button>
  </div>
  {/if}

</div>
