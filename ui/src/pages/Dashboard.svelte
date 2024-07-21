<script>
  import { onMount } from "svelte";
  import { navigate } from "svelte-routing";

  import RetroOutline from "../components/ui/RetroOutline.svelte";
  import BaseLayout from "../components/BaseLayout.svelte";
  import { pb } from "../lib/pocketbase";
  import NginxConfigCard from "../components/NginxConfigCard.svelte";
  import ResouceUsageChart from "../components/ResouceUsageChart.svelte";
  import ApplicationStatusCard from "../components/ApplicationStatusCard.svelte";
  import SiteStatusCard from "../components/SiteStatusCard.svelte";

  let name = "";

  onMount(async () => {
    if (!pb.authStore.isValid) {
      navigate("/login");
      return;
    }

    const user = pb.authStore.model;
    if (user) {
      name = user.name;
    }
  });
</script>

<BaseLayout>
  <div
    class="bg-pattern absolute inset-x-0 top-0 -z-0 h-[280px] border-b border-gray-500"
  ></div>

  <div class="container mx-auto max-w-6xl px-4 pb-12 relative">
    <div class="py-16 pb-4">
      <h1 class="text-4xl font-bold">Welcome back {name}!</h1>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-3 gap-6 py-4">
      <RetroOutline>
        <div class="p-6">
          <div class="text-xl font-bold">Total Sites</div>
          <div class="text-4xl font-bold">3</div>
        </div>
      </RetroOutline>
      <RetroOutline>
        <div class="p-6">
          <div class="text-xl font-bold">Active Sites</div>
          <div class="text-4xl font-bold">2</div>
        </div>
      </RetroOutline>
      <RetroOutline>
        <div class="p-6">
          <div class="text-xl font-bold">SSL Secured</div>
          <div class="text-4xl font-bold">0</div>
        </div>
      </RetroOutline>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 gap-8 py-4">
      <SiteStatusCard />
      <ResouceUsageChart />
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 gap-8 py-4">
      <ApplicationStatusCard />
      <NginxConfigCard />
    </div>
  </div>
</BaseLayout>