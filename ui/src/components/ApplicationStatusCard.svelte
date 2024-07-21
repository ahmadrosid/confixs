<script>
  import { CheckCircle, XCircle } from "lucide-svelte";
  import RetroOutline from "./ui/RetroOutline.svelte";
  import Button from "./ui/Button.svelte";

  // These should be props or fetched from an API
  export let isNginxInstalled = true;
  export let isCertbotInstalled = false;

  function handleNginxAction() {
    // Add logic for Nginx action (e.g., installation, configuration)
    console.log("Nginx action triggered");
  }

  function handleCertbotAction() {
    // Add logic for Certbot action (e.g., installation, SSL certificate management)
    console.log("Certbot action triggered");
  }
</script>

<RetroOutline childClassName="p-0">
  <div class="text-xl font-bold border-b border-gray-800 p-3">
    Application Status
  </div>
  <div class="p-4">
    <ul class="space-y-4">
      <li class="flex items-center space-x-2">
        <span class="font-semibold flex-grow">Nginx</span>
        {#if isNginxInstalled}
          <CheckCircle class="text-green-500" size={24} />
          <Button on:click={handleNginxAction}>Configure</Button>
        {:else}
          <XCircle class="text-red-500" size={24} />
          <Button on:click={handleNginxAction}>Install</Button>
        {/if}
      </li>
      <li class="flex items-center space-x-2">
        <span class="font-semibold flex-grow">Certbot</span>
        {#if isCertbotInstalled}
          <CheckCircle class="text-green-500" size={24} />
          <Button on:click={handleCertbotAction}>Manage SSL</Button>
        {:else}
          <XCircle class="text-red-500" size={24} />
          <Button on:click={handleCertbotAction} variant="danger">
            Install
          </Button>
        {/if}
      </li>
    </ul>
  </div>
</RetroOutline>
