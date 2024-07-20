<script>
  import { Bar } from "svelte-chartjs";
  import {
    Chart,
    Title,
    Tooltip,
    Legend,
    BarElement,
    CategoryScale,
    LinearScale,
  } from "chart.js";

  import RetroOutline from "../components/ui/RetroOutline.svelte";
  import BaseLayout from "../components/BaseLayout.svelte";
  import { pb } from "../lib/pocketbase";

  import { onMount } from "svelte";
  import { navigate } from "svelte-routing";
  import NginxConfigCard from "../components/NginxConfigCard.svelte";

  let name = "";
  Chart.register(
    Title,
    Tooltip,
    Legend,
    BarElement,
    CategoryScale,
    LinearScale
  );
  let chartData = {
    labels: ["CPU", "Memory", "Disk"],
    datasets: [
      {
        label: "Usage",
        data: [65, 80, 45],
        backgroundColor: "#000",
      },
    ],
  };

  let chartOptions = {
    responsive: true,
    maintainAspectRatio: false,
  };

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
      <RetroOutline childClassName="p-0">
        <div class="text-xl font-bold border-b border-gray-800 p-3">
          Site Status
        </div>
        <div class="p-4">
          <ul class="space-y-2 list-inside">
            <li>
              <h3 class="font-semibold">example.com</h3>
              <p class="text-sm">Lesson intro database course</p>
            </li>
            <li>
              <h3 class="font-semibold">test.com</h3>
              <p class="text-sm">Setup database schema</p>
            </li>
            <li>
              <h3 class="font-semibold">dev.example.com</h3>
              <p class="text-sm">Finish deployment</p>
            </li>
          </ul>
        </div>
      </RetroOutline>
      <RetroOutline childClassName="p-0">
        <div class="text-xl font-bold border-b border-gray-800 p-3">
          Server Resource Usage
        </div>
        <div class="p-4">
          <Bar data={chartData} options={chartOptions} />
        </div>
      </RetroOutline>
    </div>

    <div class="py-4">
      <NginxConfigCard />
    </div>
  </div>
</BaseLayout>
