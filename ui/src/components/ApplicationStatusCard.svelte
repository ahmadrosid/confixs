<script>
  import { CheckCircle, XCircle } from "lucide-svelte";
  import RetroOutline from "./ui/RetroOutline.svelte";

  export let isNginxInstalled = true;
  export let isCertbotInstalled = false;
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
        {:else}
          <XCircle class="text-red-500" size={24} />
        {/if}
      </li>
      <li class="flex items-center space-x-2">
        <span class="font-semibold flex-grow">Certbot</span>
        {#if isCertbotInstalled}
          <CheckCircle class="text-green-500" size={24} />
        {:else}
          <XCircle class="text-red-500" size={24} />
        {/if}
      </li>
    </ul>
  </div>
</RetroOutline>
