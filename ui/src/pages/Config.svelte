<script>
    import BaseLayout from "../components/BaseLayout.svelte";
    import { pb } from "../lib/pocketbase";
  
    import { onMount } from "svelte";
    import { navigate } from "svelte-routing";
    import NginxConfigCard from "../components/NginxConfigCard.svelte";
  
    onMount(async () => {
      if (!pb.authStore.isValid) {
        navigate("/login");
        return;
      }
    });
  </script>
  
  <BaseLayout>
    <div
      class="bg-pattern absolute inset-x-0 top-0 -z-0 h-[280px] border-b border-gray-500"
    ></div>
  
    <div class="container mx-auto max-w-6xl px-4 pb-12 relative">
      <div class="py-16 pb-4">
        <h1 class="text-4xl font-bold">Nginx config list</h1>
      </div>

      <div class="py-4 min-h-[60vh]">
        <NginxConfigCard />
      </div>
    </div>
  </BaseLayout>
  